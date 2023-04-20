package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.stackroute.in/JitenP/golang-training-questglobal/database"
	"gitlab.stackroute.in/JitenP/golang-training-questglobal/mb"
	"gitlab.stackroute.in/JitenP/golang-training-questglobal/models"
)

type Contact struct {
	database.Contact // promoted
	mb.Kafka         // kafka related messages
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
			bytes, _ := json.Marshal(contact)
			err := c.Kafka.WriteMessage(bytes, "Contact-Created") // producer

			fmt.Println(err)
			ctx.JSON(http.StatusCreated, contact)
			//ctx.Abort()
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

func (c *Contact) DeleteBy(ch chan<- string) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var (
			err error
			id  string
			_id int
			ok  bool
		)

		if ctx.Request.Method != "DELETE" {
			ctx.String(http.StatusNotImplemented, "http method not implementd")
			ctx.Abort()
			return
		}

		if id, ok = ctx.Params.Get("id"); !ok {
			ctx.String(http.StatusBadRequest, "invalid id")
			ctx.Abort()
			return
		}

		_id, err = strconv.Atoi(id)
		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			ctx.Abort()
			return
		}

		if rowAffected, err := c.Delete(_id); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			ctx.Abort()
			return
		} else if rowAffected == 0 {
			ctx.String(http.StatusBadRequest, "no data found with the given id->"+id)
			ctx.Abort()
			return
		} else {
			ch <- "successfully deleted on " + time.Now().String()
			ctx.JSON(http.StatusAccepted, rowAffected)
			ctx.Abort()
			return
		}

	}
}

func (c *Contact) GetByID(ch chan<- string) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var (
			err error
			id  string
			_id int
			ok  bool
		)

		if ctx.Request.Method != "GET" {
			ctx.String(http.StatusNotImplemented, "http method not implementd")
			ctx.Abort()
			return
		}

		if id, ok = ctx.Params.Get("id"); !ok {
			ctx.String(http.StatusBadRequest, "id not found")
			ctx.Abort()
			return
		}

		_id, err = strconv.Atoi(id)
		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			ctx.Abort()
			return
		}

		if contact, err := c.GetBy(_id); err != nil {
			if err.Error() == "record not found" {
				ctx.String(http.StatusBadRequest, "no data with the given id->"+id)
				ctx.Abort()
				return
			}
			ctx.String(http.StatusBadRequest, err.Error())
			ctx.Abort()
			return
		} else {
			ch <- "successfully fetched on " + time.Now().String()
			ctx.JSON(http.StatusAccepted, contact)
			ctx.Abort()
			return
		}

	}
}

func (c *Contact) GetAll(ch chan<- string) func(*gin.Context) {
	return func(ctx *gin.Context) {

		if ctx.Request.Method != "GET" {
			ctx.String(http.StatusNotImplemented, "http method not implementd")
			ctx.Abort()
			return
		}

		if contacts, err := c.Contact.GetAll(); err != nil {

			ctx.String(http.StatusBadRequest, err.Error())
			ctx.Abort()
			return
		} else {
			ch <- "successfully fetched all recordson " + time.Now().String()
			ctx.JSON(http.StatusAccepted, contacts)
			ctx.Abort()
			return
		}

	}
}

func (c *Contact) GetAllBy(ch chan<- string) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var (
			err           error
			skip, limit   string
			_skip, _limit int
			ok            bool
		)

		if ctx.Request.Method != "GET" {
			ctx.String(http.StatusNotImplemented, "http method not implementd")
			ctx.Abort()
			return
		}

		if limit, ok = ctx.Params.Get("limit"); !ok {
			ctx.String(http.StatusBadRequest, "limit not found")
			ctx.Abort()
			return
		}

		if skip, ok = ctx.Params.Get("skip"); !ok {
			ctx.String(http.StatusBadRequest, "skip not found")
			ctx.Abort()
			return
		}

		_skip, err = strconv.Atoi(skip)
		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			ctx.Abort()
			return
		}
		_limit, err = strconv.Atoi(limit)
		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			ctx.Abort()
			return
		}

		if ctx.Request.Method != "GET" {
			ctx.String(http.StatusNotImplemented, "http method not implementd")
			ctx.Abort()
			return
		}

		if contacts, err := c.Contact.GetAllBy(_skip, _limit); err != nil {

			ctx.String(http.StatusBadRequest, err.Error())
			ctx.Abort()
			return
		} else {
			ch <- "successfully fetched all records by skip and limit on " + time.Now().String()
			ctx.JSON(http.StatusAccepted, contacts)
			ctx.Abort()
			return
		}

	}
}
