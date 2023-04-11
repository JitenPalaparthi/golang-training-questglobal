package interfaces

import (
	"context"

	"github.com/JitenPalaparthi/vehicleMaster/models"
)

type ICompany interface {
	Create(ctx context.Context, company *models.Company) (*models.Company, error)
	UpdateBy(ctx context.Context, id string, data map[string]interface{}) (*models.Company, error)
	GetBy(ctx context.Context, id string) (*models.Company, error)
	GetAllByOffSet(ctx context.Context, offset, limit int) ([]models.Company, error)
	DeleteBy(ctx context.Context, id string) (interface{}, error)
	Search(ctx context.Context, offset, limit int, search string) ([]models.Company, error)
}
