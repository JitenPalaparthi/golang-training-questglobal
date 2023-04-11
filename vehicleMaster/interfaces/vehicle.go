package interfaces

import (
	"context"

	"github.com/JitenPalaparthi/vehicleMaster/models"
)

type IVehicle interface {
	Create(ctx context.Context, vehicle *models.Vehicle) (*models.Vehicle, error)
	UpdateBy(ctx context.Context, id string, data map[string]interface{}) (*models.Vehicle, error)
	GetBy(ctx context.Context, id string) (*models.Vehicle, error)
	GetAllByOffSet(ctx context.Context, offset, limit int) ([]models.Vehicle, error)
	DeleteBy(ctx context.Context, id string) (interface{}, error)
}
