package database

import (
	"context"
	"strconv"

	"github.com/JitenPalaparthi/vehicleMaster/models"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type VehicleSpecs struct {
	DB interface{}
}

func (vs *VehicleSpecs) Create(ctx context.Context, spec *models.VehicleSpec) (*models.VehicleSpec, error) {
	// Need to check this implementation. What if second time it is called
	if err := vs.DB.(*gorm.DB).AutoMigrate(&models.VehicleSpec{}); err != nil {
		return nil, err
	}
	tx := vs.DB.(*gorm.DB).Create(spec)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return spec, nil
}

func (vs *VehicleSpecs) UpdateBy(ctx context.Context, id string, data map[string]interface{}) (*models.VehicleSpec, error) {
	spec := new(models.VehicleSpec)
	_id, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	type TD = struct {
		TechDetails datatypes.JSONMap `json:"techDetails" gorm:"column:techDetails"`
	}
	var td TD
	val, ok := data["techDetails"]
	if ok {

		if vs.DB.(*gorm.DB).Model(spec).Find(&td, id).Error == nil {
			for k, v := range val.(map[string]interface{}) {
				td.TechDetails[k] = v
			}
		} else {
			return nil, ErrNotUpdateMoreDetails
		}
		data["techDetails"] = td.TechDetails
	}
	spec.ID = uint(_id)
	tx := vs.DB.(*gorm.DB).Model(spec).Updates(data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	spec, err = vs.GetBy(ctx, id)
	if err != nil {
		return nil, err
	}

	return spec, nil
}
func (vs *VehicleSpecs) GetBy(ctx context.Context, id string) (*models.VehicleSpec, error) {
	spec := new(models.VehicleSpec)
	tx := vs.DB.(*gorm.DB).First(spec, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return spec, nil
}

func (vs *VehicleSpecs) DeleteBy(ctx context.Context, id string) (interface{}, error) {
	tx := vs.DB.(*gorm.DB).Delete(&models.VehicleSpec{}, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx.RowsAffected, nil
}

func (vs *VehicleSpecs) GetAllByOffSet(ctx context.Context, offset, limit int) ([]models.VehicleSpec, error) {
	specs := make([]models.VehicleSpec, 0)
	tx := vs.DB.(*gorm.DB).Limit(limit).Offset(offset).Find(&specs)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return specs, nil
}
