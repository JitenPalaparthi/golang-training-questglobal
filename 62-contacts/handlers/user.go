package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.stackroute.in/JitenP/golang-training-questglobal/database"
	"gitlab.stackroute.in/JitenP/golang-training-questglobal/models"
)

type User struct {
	database.User // promoted
}

func (c *User) Register(ch chan<- string) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var err error
		if ctx.Request.Method != "POST" {
			ctx.String(http.StatusNotImplemented, "http method not implementd")
			ctx.Abort()
			return
		}

		user := new(models.User)

		if err := ctx.Bind(user); err != nil { // json.Unmarshal
			ctx.String(http.StatusBadRequest, err.Error())
			ctx.Abort()
			return
		}

		if err := user.ValidateRegister(); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			ctx.Abort()
			return
		}

		user.Status = "created"
		user.LastModified = time.Now().Unix() // brainfeed: What is Unix time.How to convert from Unix time to normal data and time and vice versa

		if user, err = c.User.Register(user); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			ctx.Abort()
			return
		} else {
			ch <- "successfully regsterd at " + time.Now().String()
			ctx.JSON(http.StatusCreated, user)
			ctx.Abort()
			return
		}

	}
}

func (c *User) Login(ch chan<- string) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var err error
		if ctx.Request.Method != "POST" {
			ctx.String(http.StatusNotImplemented, "http method not implementd")
			ctx.Abort()
			return
		}

		user := new(models.User)

		if err := ctx.Bind(&user); err != nil { // json.Unmarshal
			ctx.String(http.StatusBadRequest, err.Error())
			ctx.Abort()
			return
		}

		if err := user.ValidateLogin(); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			ctx.Abort()
			return
		}

		user.Status = "created"
		user.LastModified = time.Now().Unix() // brainfeed: What is Unix time.How to convert from Unix time to normal data and time and vice versa

		if err = c.User.Login(user.UName, user.Password); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			ctx.Abort()
			return
		} else {
			ch <- "successfully regsterd at " + time.Now().String()
			ctx.String(http.StatusOK, "Sucessfully logged in")
			ctx.Abort()
			return
		}

	}
}
