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

type Vehicle struct {
	IVehicle interfaces.IVehicle
}

func (vs *Vehicle) Create(ctx context.Context) func(*gin.Context) {
	return func(c *gin.Context) {
		vehicle := new(models.Vehicle)
		err := c.Bind(vehicle)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": "bad request"})
			return
		}
		if err := vehicle.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": err.Error()})
			return
		}
		vehicle.Status = "active" //TODO
		vehicle.LastModified = time.Now().Unix()
		if con, err := vs.IVehicle.Create(ctx, vehicle); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": "Error in creating contact info"})
			return
		} else {

			c.JSON(http.StatusCreated, con)
			return
		}
	}
}

func (vs *Vehicle) GetBy(ctx context.Context) func(*gin.Context) {
	return func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": "invalid id parameter", "message": "invalid id parameter"})
			return
		}

		vehicle, err := vs.IVehicle.GetBy(ctx, id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": "bad request"})
			return
		}
		c.JSON(http.StatusOK, vehicle)
	}
}

func (vs *Vehicle) GetAllByOffset(ctx context.Context) func(*gin.Context) {
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
		vehicle, err := vs.IVehicle.GetAllByOffSet(ctx, _offset, _limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": "bad request"})
			return
		}
		c.JSON(http.StatusOK, vehicle)
	}
}

func (vs *Vehicle) DeleteBy(ctx context.Context) func(*gin.Context) {
	return func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": "invalid id parameter", "message": "invalid id parameter"})
			return
		}

		noOfRecords, err := vs.IVehicle.DeleteBy(ctx, id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": "bad request"})
			return
		}
		c.JSON(http.StatusAccepted, noOfRecords)
	}
}

func (vs *Vehicle) UpdateBy(ctx context.Context) func(*gin.Context) {
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

		vehicle, err := vs.IVehicle.UpdateBy(ctx, id, data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": "bad request"})
			return
		}
		c.JSON(http.StatusOK, vehicle)
	}
}
