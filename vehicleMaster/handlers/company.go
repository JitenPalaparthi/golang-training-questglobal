package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/JitenPalaparthi/vehicleMaster/interfaces"
	"github.com/JitenPalaparthi/vehicleMaster/models"

	"github.com/gin-gonic/gin"
)

type Company struct {
	ICompany interfaces.ICompany
}

func (cp *Company) Create(ctx context.Context) func(*gin.Context) {
	return func(c *gin.Context) {
		company := new(models.Company)
		err := c.Bind(company)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": "bad request"})
			return
		}
		if err := company.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": err.Error()})
			return
		}

		company.Status = "active" //TODO
		company.LastModified = time.Now().Unix()

		company.Tags = company.ToString() // adding tags

		if con, err := cp.ICompany.Create(ctx, company); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": "Error in creating contact info"})
			return
		} else {

			c.JSON(http.StatusCreated, con)
			return
		}
	}
}

func (cp *Company) GetBy(ctx context.Context) func(*gin.Context) {
	return func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": "invalid id parameter", "message": "invalid id parameter"})
			return
		}

		company, err := cp.ICompany.GetBy(ctx, id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": "bad request"})
			return
		}
		c.JSON(http.StatusOK, company)
	}
}

func (cp *Company) GetAllByOffset(ctx context.Context) func(*gin.Context) {
	return func(c *gin.Context) {

		limit, ok := c.Params.Get("limit")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": "invalid id parameter", "message": "invalid id parameter"})
			return
		}
		_limit, err := strconv.Atoi(limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": "bad request"})
			return
		}
		offset, ok := c.Params.Get("offset")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": "invalid id parameter", "message": "invalid id parameter"})
			return
		}
		_offset, err := strconv.Atoi(offset)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": "bad request"})
			return
		}
		companies, err := cp.ICompany.GetAllByOffSet(ctx, _offset, _limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": "bad request"})
			return
		}
		c.JSON(http.StatusOK, companies)
	}
}

func (cp *Company) Search(ctx context.Context) func(*gin.Context) {
	return func(c *gin.Context) {

		limit, ok := c.Params.Get("limit")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": "invalid id parameter", "message": "invalid id parameter"})
			return
		}
		_limit, err := strconv.Atoi(limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": "bad request"})
			return
		}
		offset, ok := c.Params.Get("offset")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": "invalid id parameter", "message": "invalid id parameter"})
			return
		}
		_offset, err := strconv.Atoi(offset)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": "bad request"})
			return
		}

		search, ok := c.Params.Get("search")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": "invalid id parameter", "message": "invalid id parameter"})
			return
		}

		companies, err := cp.ICompany.Search(ctx, _offset, _limit, search)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": "bad request"})
			return
		}
		c.JSON(http.StatusOK, companies)
	}
}

func (cp *Company) DeleteBy(ctx context.Context) func(*gin.Context) {
	return func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": "invalid id parameter", "message": "invalid id parameter"})
			return
		}

		noOfRecords, err := cp.ICompany.DeleteBy(ctx, id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": "bad request"})
			return
		}
		c.JSON(http.StatusAccepted, noOfRecords)
	}
}

func (cp *Company) UpdateBy(ctx context.Context) func(*gin.Context) {
	return func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": "invalid id parameter", "message": "invalid id parameter"})
			return
		}
		data := make(map[string]interface{})
		err := c.Bind(&data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": "bad request"})
			return
		}

		company, err := cp.ICompany.UpdateBy(ctx, id, data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": "bad request"})
			return
		}
		c.JSON(http.StatusOK, company)
	}
}
