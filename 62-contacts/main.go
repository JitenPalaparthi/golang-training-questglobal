package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gitlab.stackroute.in/JitenP/golang-training-questglobal/database"
	"gitlab.stackroute.in/JitenP/golang-training-questglobal/handlers"
)

var (
	port string
	dsn  string //= "host=localhost user=postgres password=postgres dbname=contactsbd port=55432 sslmode=disable TimeZone=Asia/Shanghai"
	// for mysql: --dsn="user:password@tcp(127.0.0.1:53306)/contactsdb?charset=utf8mb4&parseTime=True&loc=Local"
	mb string = "localhost:29092"
	ch chan string

	secretCode = "IAmTheSecretCode"
)

func main() {
	port = os.Getenv("PORT")
	if port == "" {
		flag.StringVar(&port, "port", "8080", "--port=8080 or --port 8080 or -port 8080")
	}
	dsn = os.Getenv("DBCONN")
	if dsn == "" {
		flag.StringVar(&dsn, "dsn", "host=localhost user=postgres password=postgres dbname=contactsbd port=55432 sslmode=disable TimeZone=Asia/Shanghai", "--dsn=host=localhost user=postgres password=postgres dbname=contactsbd port=55432 sslmode=disable TimeZone=Asia/Shanghai")
	}
	if mb == "" {
		flag.StringVar(&mb, "mb", "localhost:29092", "--mb=localhost:29092")

	}
	flag.Parse()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))
	db, err := database.Connect(dsn)
	if err != nil {
		log.Fatalln(err) //Fatals cannot be recoverd but panics can. Fatal exit the application with os.Exit(1)
	} else {
		log.Println(db)
	}
	ch = make(chan string, 10)

	uHandler := new(handlers.User)
	uHandler.DB = db

	public_v1 := r.Group("/v1/public")
	{
		public_v1.GET("/ping", func(c *gin.Context) {
			m := &Message{Status: "pong"}
			c.JSON(http.StatusOK, m)
		})
		public_v1.POST("/login", uHandler.Login(ch))
		public_v1.POST("/register", uHandler.Register(ch))

	}

	cHandler := new(handlers.Contact) //creating new instance of handler
	cHandler.DB = db
	cHandler.Conns = []string{mb}
	private_v1 := r.Group("v1/private")
	{
		private_v1.POST("/contact/add", cHandler.Create(ch))
		private_v1.PUT("/contact/update/:id", cHandler.UpdateBy(ch))
		private_v1.DELETE("/contact/delete/:id", cHandler.DeleteBy(ch))
		private_v1.GET("/contact/get/:id", cHandler.GetByID(ch))
		private_v1.GET("/contact/get/all", cHandler.GetAll(ch))
		private_v1.GET("/contact/get/all/:skip/:limit", cHandler.GetAllBy(ch))
		private_v1.GET("/greet/:name", Authenticate("123456"), func(c *gin.Context) {
			if name, ok := c.Params.Get("name"); !ok {
				c.String(http.StatusBadRequest, "name parameter has not been provided")
				c.Abort()
			} else {
				c.String(http.StatusOK, "Hello "+name)
			}
			//c.Abort()
			c.Next()
		}, Audit)
	}
	go func() {
		for v := range ch {
			file, err := os.OpenFile("audit.io", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
			if err != nil {
				//fmt.Println(err)
				log.Println(err)
			}
			defer file.Close()

			_, err = file.WriteString(v + "\n")
			if err != nil {
				log.Println(err)
			}
		}
	}()

	//log.Fatal(r.RunTLS(":"+port, "security/localhost.crt", "security/localhost.key"))
	log.Fatal(r.Run(":" + port))
}

func Authenticate(token string) func(*gin.Context) {
	return func(c *gin.Context) {
		email := c.Request.Header.Get("email")
		password := c.Request.Header.Get("password")

		if email == "" || password == "" {
			c.String(http.StatusUnauthorized, "bad credentials")
			c.Abort()
		} else if email == "admin@questglobal.com" && password == "admin@123" && token == "123456" {
			c.Next()
		} else {
			c.String(http.StatusUnauthorized, "bad credentials")
			c.Abort()
		}
	}
}

func Audit(c *gin.Context) {
	url := c.Request.URL
	file, err := os.OpenFile("audit.io", os.O_RDWR, 777)
	if err != nil {
		//fmt.Println(err)
		log.Println(err)
	}
	defer file.Close()

	_, err = file.Write([]byte(url.String()))
	if err != nil {
		log.Println(err)
	}
}

type Message struct {
	Status string
}

// 1- gin package
// 2- commandline arguments. flags package
// 3- env variables. os package
// 4- datbase retry logic
// 5- logging // glog
// 6- gorm package // gorm package implementation
// 7- kafka message broker
// 8- dockerize this repo

// What this micro service is for
// This does CRUD operations on contacts. Creating/ Reading/ Updating and Deleting contects in a database
// Dependency injectin and also follow solid principles

// In order to pass the request , user has to pass valid email and password
