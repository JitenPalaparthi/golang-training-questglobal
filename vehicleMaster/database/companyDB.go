package database

import (
	"context"
	"errors"
	"strconv"

	"github.com/JitenPalaparthi/vehicleMaster/models"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Company struct {
	DB interface{}
}

var (
	ErrNotUpdateMoreDetails = errors.New("unable to update moreDetails field")
)

func (c *Company) Create(ctx context.Context, company *models.Company) (*models.Company, error) {
	// Need to check this implementation. What if second time it is called
	if err := c.DB.(*gorm.DB).AutoMigrate(&models.Company{}); err != nil {
		return nil, err
	}
	tx := c.DB.(*gorm.DB).Create(company)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return company, nil
}

func (c *Company) UpdateBy(ctx context.Context, id string, data map[string]interface{}) (*models.Company, error) {
	company := new(models.Company)
	_id, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	type MD = struct {
		MoreDetails datatypes.JSONMap `json:"moreDetails" gorm:"column:moreDetails"`
	}
	var md MD
	val, ok := data["moreDetails"]
	if ok {

		if c.DB.(*gorm.DB).Model(company).Find(&md, id).Error == nil {
			for k, v := range val.(map[string]interface{}) {
				md.MoreDetails[k] = v
			}
		} else {
			return nil, ErrNotUpdateMoreDetails
		}
		data["moreDetails"] = md.MoreDetails
	}
	company.ID = uint(_id)
	tx := c.DB.(*gorm.DB).Model(company).Updates(data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	company, err = c.GetBy(ctx, id)
	if err != nil {
		return nil, err
	}

	return company, nil
}
func (c *Company) GetBy(ctx context.Context, id string) (*models.Company, error) {
	company := new(models.Company)
	tx := c.DB.(*gorm.DB).First(company, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return company, nil
}

func (c *Company) DeleteBy(ctx context.Context, id string) (interface{}, error) {
	tx := c.DB.(*gorm.DB).Delete(&models.Company{}, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx.RowsAffected, nil
}

func (c *Company) GetAllByOffSet(ctx context.Context, offset, limit int) ([]models.Company, error) {
	companies := make([]models.Company, 0)
	tx := c.DB.(*gorm.DB).Limit(limit).Offset(offset).Find(&companies)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return companies, nil
}

func (c *Company) Search(ctx context.Context, offset, limit int, search string) ([]models.Company, error) {
	companies := make([]models.Company, 0)
	//tx := c.DB.(*gorm.DB).Limit(limit).Offset(offset).Find(&companies)
	tx := c.DB.(*gorm.DB).Where("tags @@ to_tsquery(?)", search).Limit(limit).Offset(offset).Find(&companies)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return companies, nil
}
