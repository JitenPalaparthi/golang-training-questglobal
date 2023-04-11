package interfaces

import (
	"context"

	"github.com/JitenPalaparthi/vehicleMaster/models"
)

type IVehicleSpec interface {
	Create(ctx context.Context, contact *models.VehicleSpec) (*models.VehicleSpec, error)
	UpdateBy(ctx context.Context, id string, data map[string]interface{}) (*models.VehicleSpec, error)
	GetBy(ctx context.Context, id string) (*models.VehicleSpec, error)
	//GetAll(ctx context.Context, filter map[string]interface{}) ([]models.VehicleSpec, error)
	GetAllByOffSet(ctx context.Context, offset, limit int) ([]models.VehicleSpec, error)
	DeleteBy(ctx context.Context, id string) (interface{}, error)
}
