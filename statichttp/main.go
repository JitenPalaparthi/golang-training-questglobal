//package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	fs := http.FileServer(http.Dir("static"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))
// 	fmt.Println(http.ListenAndServe(":8085", nil))
// }

package main

import (
	"log"
	"web/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	controller.Router(r)
	r.Static("/static", "./static")
	r.Static("/css", "./static/css")
	// r.Static("/img", "./static/img")
	// r.Static("/scss", "./static/scss")
	// r.Static("/vendor", "./static/vendor")
	r.Static("/js", "./static/js")
	//r.StaticFile("/favicon.ico", "./img/favicon.ico")

	//r.LoadHTMLGlob("templates/**/*")
	log.Println("Server started")
	r.Run(":8085")
}
