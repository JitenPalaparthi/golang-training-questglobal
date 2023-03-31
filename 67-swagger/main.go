package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/helloworld", Helloworld)
	r.GET("/greet/:name", Helloworld)

	r.Run(":8080")
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// GreetExample godoc
// @Summary greet example
// @Schemes
// @Description greet with name
// @Tags greet
// @Accept json
// @Produce json
// @Success 200 {string} Hello {name}
// @Param        name   path      string  true  "Name"
// @Router /greet/:name [get]
func Greet(g *gin.Context) {
	name, _ := g.Params.Get("name")
	g.String(200, "Hello %s", name)
}
