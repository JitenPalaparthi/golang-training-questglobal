package main

import (
	"fmt"
	"net/http"
	"time"

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
//
