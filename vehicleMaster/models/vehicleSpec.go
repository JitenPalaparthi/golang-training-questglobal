package models

import (
	"encoding/json"
	"fmt"

	"gorm.io/datatypes"
)

type VehicleSpec struct {
	ID               uint   `json:"id"`
	Make             string `json:"make"`
	Model            string `json:"model"`
	Brand            string `json:"brand"`
	BodyType         string `json:"bodyType" gorm:"column:bodyType"`
	TransmissionType string `json:"transmissionType" gorm:"column:transmissionType"`
	VehicleType      string `json:"vehicleType" gorm:"column:vehicleType"` // Car Type
	FuelType         string `json:"fuelType" gorm:"column:fuelType"`
	Color            string `json:"color"`
	TechDetails      datatypes.JSONMap/*map[string]interface{}*/ `json:"techDetails" gorm:"column:techDetails"`
	Tags             string `json:"tags"` // comma separated values
	Status           string `json:"status"`
	LastModified     int64  `json:"lastModified" gorm:"column:lastModified"`
}

func (v *VehicleSpec) ToByte() ([]byte, error) {
	return json.Marshal(v)
}

func (v *VehicleSpec) ToJSONString() (string, error) {
	buf, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func (v *VehicleSpec) Validate() error {
	//TODO
	return nil
}

func (vs *VehicleSpec) ToString() string {
	return fmt.Sprintln(*vs)
}
