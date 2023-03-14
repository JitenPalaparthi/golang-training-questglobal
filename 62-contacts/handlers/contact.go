package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.stackroute.in/JitenP/golang-training-questglobal/database"
	"gitlab.stackroute.in/JitenP/golang-training-questglobal/models"
)

type Contact struct {
	database.Contact // promoted
}

func (c *Contact) Create(ch chan<- string) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var err error
		if ctx.Request.Method != "POST" {
			ctx.String(http.StatusNotImplemented, "http method not implementd")
			ctx.Abort()
			return
		}

		contact := new(models.Contact)

		// if err := json.NewDecoder(ctx.Request.Body).Decode(contact); err != nil {
		// 	ctx.String(http.StatusBadRequest, err.Error())
		// 	ctx.Abort()
		// 	return
		// }

		if err := ctx.Bind(contact); err != nil { // json.Unmarshal
			ctx.String(http.StatusBadRequest, err.Error())
			ctx.Abort()
			return
		}

		if err := contact.Validate(); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			ctx.Abort()
			return
		}

		contact.Status = "created"
		contact.LastModified = time.Now().Unix() // brainfeed: What is Unix time.How to convert from Unix time to normal data and time and vice versa

		if contact, err = c.Add(contact); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			ctx.Abort()
			return
		} else {
			ch <- "successfully updated on " + time.Now().String()
			ctx.JSON(http.StatusCreated, contact)
			ctx.Abort()
			return
		}

	}
}

// func (c *Contact) UpdateByD(ctx *gin.Context) {
// 	var (
// 		err error
// 		id  string
// 		_id int
// 		ok  bool
// 	)
// 	if ctx.Request.Method != "PUT" {
// 		ctx.String(http.StatusNotImplemented, "http method not implementd")
// 		ctx.Abort()
// 		return
// 	}

// 	if id, ok = ctx.Params.Get("id"); !ok {
// 		ctx.String(http.StatusBadRequest, "id not found")
// 		ctx.Abort()
// 		return
// 	}
// 	data := make(map[string]any)

// 	if err := ctx.Bind(data); err != nil { // json.Unmarshal
// 		ctx.String(http.StatusBadRequest, err.Error())
// 		ctx.Abort()
// 		return
// 	}

// 	data["status"] = "updated"

// 	data["lastModified"] = time.Now().Unix() // brainfeed: What is Unix time.How to convert from Unix time to normal data and time and vice versa
// 	_id, err = strconv.Atoi(id)
// 	if err != nil {
// 		ctx.String(http.StatusBadRequest, err.Error())
// 		ctx.Abort()
// 		return
// 	}

// 	if contact, err := c.Update(_id, data); err != nil {
// 		ctx.String(http.StatusBadRequest, err.Error())
// 		ctx.Abort()
// 		return
// 	} else {

// 		ctx.JSON(http.StatusCreated, contact)
// 		ctx.Abort()
// 		return
// 	}
// }

func (c *Contact) UpdateBy(ch chan<- string) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var (
			err error
			id  string
			_id int
			ok  bool
		)

		if ctx.Request.Method != "PUT" {
			ctx.String(http.StatusNotImplemented, "http method not implementd")
			ctx.Abort()
			return
		}

		if id, ok = ctx.Params.Get("id"); !ok {
			ctx.String(http.StatusBadRequest, "id not found")
			ctx.Abort()
			return
		}

		data := make(map[string]any)

		if err := ctx.Bind(&data); err != nil { // json.Unmarshal
			ctx.String(http.StatusBadRequest, err.Error())
			ctx.Abort()
			return
		}

		data["status"] = "updated"

		data["lastModified"] = time.Now().Unix() // brainfeed: What is Unix time.How to convert from Unix time to normal data and time and vice versa
		_id, err = strconv.Atoi(id)
		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			ctx.Abort()
			return
		}

		if contact, err := c.Update(_id, data); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			ctx.Abort()
			return
		} else {
			ch <- "successfully updated on " + time.Now().String()
			ctx.JSON(http.StatusCreated, contact)
			ctx.Abort()
			return
		}

	}
}
