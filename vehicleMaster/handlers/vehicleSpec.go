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

type VehicleSpec struct {
	IVehicleSpec interfaces.IVehicleSpec
}

func (vs *VehicleSpec) Create(ctx context.Context) func(*gin.Context) {
	return func(c *gin.Context) {
		spec := new(models.VehicleSpec)
		err := c.Bind(spec)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": "bad request"})
			return
		}
		if err := spec.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": err.Error()})
			return
		}
		spec.Status = "active" //TODO
		spec.LastModified = time.Now().Unix()
		if con, err := vs.IVehicleSpec.Create(ctx, spec); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": "Error in creating contact info"})
			return
		} else {

			c.JSON(http.StatusCreated, con)
			return
		}
	}
}

func (vs *VehicleSpec) GetBy(ctx context.Context) func(*gin.Context) {
	return func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": "invalid id parameter", "message": "invalid id parameter"})
			return
		}

		spec, err := vs.IVehicleSpec.GetBy(ctx, id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": "bad request"})
			return
		}
		c.JSON(http.StatusOK, spec)
	}
}

func (vs *VehicleSpec) GetAllByOffset(ctx context.Context) func(*gin.Context) {
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
		specs, err := vs.IVehicleSpec.GetAllByOffSet(ctx, _offset, _limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": "bad request"})
			return
		}
		c.JSON(http.StatusOK, specs)
	}
}

func (vs *VehicleSpec) DeleteBy(ctx context.Context) func(*gin.Context) {
	return func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": "invalid id parameter", "message": "invalid id parameter"})
			return
		}

		noOfRecords, err := vs.IVehicleSpec.DeleteBy(ctx, id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": "bad request"})
			return
		}
		c.JSON(http.StatusAccepted, noOfRecords)
	}
}

func (vs *VehicleSpec) UpdateBy(ctx context.Context) func(*gin.Context) {
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

		spec, err := vs.IVehicleSpec.UpdateBy(ctx, id, data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "innerError": err.Error(), "message": "bad request"})
			return
		}
		c.JSON(http.StatusOK, spec)
	}
}
