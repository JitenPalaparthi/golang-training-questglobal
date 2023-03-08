package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		m := &Message{Status: "pong"}
		c.JSON(http.StatusOK, m)
		//c.XML(http.StatusOK, gin.H{"message": "pong"})
		//c.String(http.StatusOK, "pong")
		// c.XML(http.StatusOK, m)
	})

	r.GET("/greet/:name", Authenticate(), func(c *gin.Context) {
		if name, ok := c.Params.Get("name"); !ok {
			c.String(http.StatusBadRequest, "name parameter has not been provided")
			c.Abort()
		} else {
			c.String(http.StatusOK, "Hello "+name)
		}
		//c.Abort()
		c.Next()
	}, Audit)

	// http.HandleFunc("/ping",func(w http.ResponseWriter, r *http.Request) {
	// })
	//http.ListenAndServe(":50090", nil)
	go func() {
		r.Run(":50090")
	}()

	count := 0
	for {
		count++
		time.Sleep(time.Second * 10)
		fmt.Println("This is being executed for the past-->", count*10, "seconds")
	}

}

func Authenticate() func(*gin.Context) {
	return func(c *gin.Context) {
		email := c.Request.Header.Get("email")
		password := c.Request.Header.Get("password")

		if email == "" || password == "" {
			c.String(http.StatusUnauthorized, "bad credentials")
			c.Abort()
		} else if email == "admin@questglobal.com" && password == "admin@123" {
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
// 2- commandline arguments. flags pa ckae
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
