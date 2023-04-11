package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/JitenPalaparthi/vehicleMaster/database"
	"github.com/JitenPalaparthi/vehicleMaster/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

var (
	PORT string
	DSN  string
	// dsn := "host=localhost user=admin password=admin123 dbname=customersdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// KAFKA_CONN []string = []string{"localhost:29092"}
	ctx       context.Context
	ginLambda *ginadapter.GinLambda
	db        interface{}
	err       error

	GOOGLE_APPLICATION_CREDENTIALS,
	INSTANCE_CONNECTION_NAME,
	DB_PORT,
	DB_NAME,
	DB_USER,
	DB_PASS string
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARNING|ERROR|FATAL] -log_dir=[string]\n")
	flag.PrintDefaults()
	//os.Exit(2)
}

func init() {
	flag.Usage = usage
	ctx = context.Background()
}

func main() {

	flag.StringVar(&PORT, "port", "50090", "--port=50090")
	//flag.StringVar(&DSN, "db", "host=localhost user=admin password=admin123 dbname=vehicle_master_db port=5432 sslmode=disable TimeZone=Asia/Shanghai", "--db=host=localhost user=admin password=admin123 dbname=customersdb port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	flag.StringVar(&DSN, "db", "host=127.0.0.1 user=postgres password=postgres dbname=vehicle_master_bd port=28932 sslmode=disable TimeZone=Asia/Shanghai", "--db=host=localhost user=admin password=admin123 dbname=customersdb port=5432 sslmode=disable TimeZone=Asia/Shanghai")

	flag.Set("stderrthreshold", "INFO") // can set up the glog
	flag.Parse()
	defer glog.Flush()
	if os.Getenv("PORT") != "" {
		PORT = os.Getenv("PORT")
	}
	if os.Getenv("DB_CONN") != "" {
		DSN = os.Getenv("DB_CONN")
	}

	ctx, cancel := context.WithCancel(ctx)

	glog.Info("Connecting to the database--")
	db, err = database.GetConnection(DSN)

	if err != nil {
		cancel()
		glog.Fatal(err)
	}

	r := gin.Default()

	r.GET("/greet", handlers.Greet("Hello World!"))
	r.GET("/health", handlers.Health())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	glog.Info("creating new instance of VehicleSpecDB object")
	vsdb := new(database.VehicleSpecs)
	vsdb.DB = db
	vsHandler := new(handlers.VehicleSpec)
	vsHandler.IVehicleSpec = vsdb
	//vsHandler.Conn = KAFKA_CONN

	v1_specs := r.Group("v1/private/vehiclespec")
	v1_specs.POST("/add", vsHandler.Create(ctx))
	v1_specs.GET("/:id", vsHandler.GetBy(ctx))
	v1_specs.GET("/all/:offset/:limit", vsHandler.GetAllByOffset(ctx))
	v1_specs.PUT("/:id", vsHandler.UpdateBy(ctx))
	v1_specs.DELETE("/:id", vsHandler.DeleteBy(ctx))

	glog.Info("creating new instance of CompanyDB object")
	cdb := new(database.Company)
	cdb.DB = db
	cHandler := new(handlers.Company)
	cHandler.ICompany = cdb
	v1_company := r.Group("v1/private/company")
	v1_company.POST("/add", cHandler.Create(ctx))
	v1_company.GET("/:id", cHandler.GetBy(ctx))
	v1_company.GET("/all/:offset/:limit", cHandler.GetAllByOffset(ctx))
	v1_company.GET("/all/:offset/:limit/:search", cHandler.Search(ctx))
	v1_company.PUT("/:id", cHandler.UpdateBy(ctx))
	v1_company.DELETE("/:id", cHandler.DeleteBy(ctx))

	glog.Info("creating new instance of VehicleDB object")
	vdb := new(database.Vehicle)
	vdb.DB = db
	vHandler := new(handlers.Vehicle)
	vHandler.IVehicle = vdb

	v1_vehicle := r.Group("v1/private/vehicle")
	v1_vehicle.POST("/add", vHandler.Create(ctx))
	v1_vehicle.GET("/:id", vHandler.GetBy(ctx))
	v1_vehicle.GET("/all/:offset/:limit", vHandler.GetAllByOffset(ctx))
	v1_vehicle.PUT("/:id", vHandler.UpdateBy(ctx))
	v1_vehicle.DELETE("/:id", vHandler.DeleteBy(ctx))
	env := os.Getenv("GIN_MODE")
	if env == "release" {
		ginLambda = ginadapter.New(r)
		lambda.Start(Handler)
	} else {
		r.Run(":" + PORT) // http.ListenAndServe()
	}

}
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}
