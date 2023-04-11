package database

import (
	"context"
	"errors"
	"strconv"

	"github.com/JitenPalaparthi/vehicleMaster/models"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Vehicle struct {
	DB interface{}
}

var (
	ErrVehicleSpecNotFound = errors.New("Vehicle Spec not found")
)

func (v *Vehicle) Create(ctx context.Context, vehicle *models.Vehicle) (*models.Vehicle, error) {
	// Need to check this implementation. What if second time it is called

	if err := v.DB.(*gorm.DB).AutoMigrate(&models.Vehicle{}); err != nil {
		return nil, err
	}

	var count1 int64
	v.DB.(*gorm.DB).Model(&models.VehicleSpec{}).Where("id = ?", vehicle.VehicleSpecID, "status = ?", "active").Count(&count1)

	if count1 == 0 {
		return nil, ErrVehicleSpecNotFound
	}

	var count2 int64
	v.DB.(*gorm.DB).Model(&models.Company{}).Where("id = ?", vehicle.CompanyID, "status = ?", "active").Count(&count2)

	if count2 == 0 {
		return nil, ErrVehicleSpecNotFound
	}

	tx := v.DB.(*gorm.DB).Create(vehicle)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return vehicle, nil
}

func (v *Vehicle) UpdateBy(ctx context.Context, id string, data map[string]interface{}) (*models.Vehicle, error) {
	vehicle := new(models.Vehicle)
	_id, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	type RD = struct {
		RegistrationDetails datatypes.JSONMap `json:"registrationDetails" gorm:"column:registrationDetails"`
	}
	var rd RD
	val, ok := data["registrationDetails"]
	if ok {

		if v.DB.(*gorm.DB).Model(vehicle).Find(&rd, id).Error == nil {
			for k, v := range val.(map[string]interface{}) {
				rd.RegistrationDetails[k] = v
			}
		} else {
			return nil, ErrNotUpdateMoreDetails
		}
		data["registrationDetails"] = rd.RegistrationDetails
	}
	type F = struct {
		Features datatypes.JSONMap `json:"features"`
	}
	var f F
	val1, ok1 := data["features"]
	if ok1 {

		if v.DB.(*gorm.DB).Model(vehicle).Find(&rd, id).Error == nil {
			for k, v := range val1.(map[string]interface{}) {
				f.Features[k] = v
			}
		} else {
			return nil, ErrNotUpdateMoreDetails
		}
		data["features"] = f.Features
	}

	type MD = struct {
		MoreDetails datatypes.JSONMap `json:"moreDetails" gorm:"column:moreDetails"`
	}
	var md MD
	val2, ok2 := data["moreDetails"]
	if ok2 {

		if v.DB.(*gorm.DB).Model(vehicle).Find(&md, id).Error == nil {
			for k, v := range val2.(map[string]interface{}) {
				md.MoreDetails[k] = v
			}
		} else {
			return nil, ErrNotUpdateMoreDetails
		}
		data["moreDetails"] = md.MoreDetails
	}

	vehicle.ID = uint(_id)
	tx := v.DB.(*gorm.DB).Model(vehicle).Updates(data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	vehicle, err = v.GetBy(ctx, id)
	if err != nil {
		return nil, err
	}

	return vehicle, nil
}

func (v *Vehicle) GetBy(ctx context.Context, id string) (*models.Vehicle, error) {
	vehicle := new(models.Vehicle)
	tx := v.DB.(*gorm.DB).Where("id=?", id).Preload("VehicleSpec").Preload("Company").First(&vehicle)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return vehicle, nil
}

func (v *Vehicle) DeleteBy(ctx context.Context, id string) (interface{}, error) {
	tx := v.DB.(*gorm.DB).Delete(&models.Vehicle{}, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx.RowsAffected, nil
}

func (v *Vehicle) GetAllByOffSet(ctx context.Context, offset, limit int) ([]models.Vehicle, error) {
	vehicles := make([]models.Vehicle, 0)
	tx := v.DB.(*gorm.DB).Preload("VehicleSpec").Preload("Company").Limit(limit).Offset(offset).Find(&vehicles)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return vehicles, nil
}
