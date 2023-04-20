package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.stackroute.in/JitenP/golang-training-questglobal/database"
	"gitlab.stackroute.in/JitenP/golang-training-questglobal/models"

	jwt_lib "github.com/dgrijalva/jwt-go"
)

type User struct {
	database.User // promoted
}

var (
	secretCode = "IAmTheSecretCode"
)

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

			// jwt work

			token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256")) // Create a new instance of the token with a given algo
			token.Claims = jwt_lib.MapClaims{
				"id":  user.UName,
				"exp": time.Now().Add(time.Hour * 1).Unix(),
			}

			tokenstring, err := token.SignedString([]byte(secretCode))
			if err != nil {
				ctx.String(http.StatusBadRequest, err.Error())
				ctx.Abort()
			}
			ctx.String(http.StatusOK, tokenstring)
			ctx.Abort()
			return
		}

	}
}
