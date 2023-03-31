package main

import (
	"net/http"

	_ "demo/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host	localhost:8080

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	r := gin.Default()

	r.GET("/helloworld", Helloworld)
	r.GET("/greet/:name", Helloworld)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}

// Helloworld godoc
//
//	@Summary	helloworld example
//	@Schemes
//	@Description	call helloworld it returns helloworld
//	@Tags			helloworld ping
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	Helloworld
//	@Router			/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// GreetExample godoc
//
//	@Summary	greet example
//	@Schemes
//	@Description	greet with name
//	@Tags			greet
//	@Accept			json
//	@Produce		json
//	@Success		200		{string}	Hello	{name}
//	@Param			name	path		string	true	"Name"
//	@Router			/greet/:name [get]
func Greet(g *gin.Context) {
	name, _ := g.Params.Get("name")
	g.String(200, "Hello %s", name)
}
