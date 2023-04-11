package models

import (
	"encoding/json"
	"fmt"

	"gorm.io/datatypes"
)

type Vehicle struct {
	ID                  uint              `json:"id"`
	Ref                 string            `json:"ref"`
	RegistrationDetails datatypes.JSONMap `json:"registrationDetails" gorm:"column:registrationDetails"`
	Features            datatypes.JSONMap `json:"features"`
	VehicleSpecID       uint              `json:"vehicleSpecID" gorm:"column:vehicleSpecID"`
	VehicleSpec         VehicleSpec       `json:"vehicleSpec" gorm:"foreignKey:vehicleSpecID"`
	//Images              []string          `json:"images"`
	Year            uint16            `json:"year" gorm:"column:year"`
	KMReading       uint64            `json:"kmReading" gorm:"column:kmReading"`
	FuelReading     uint8             `json:"fuelReading" gorm:"column:fuelReading"`
	ReadingsTakenOn uint64            `json:"readingsTakenOn" gorm:"column:readingsTakenOn"`
	MoreDetails     datatypes.JSONMap `json:"moreDetails" gorm:"column:moreDetails"`
	DetailsOn       uint64            `json:"detailsOn" gorm:"column:detailsOn"`
	CompanyID       uint              `json:"companyID" gorm:"column:companyID"`
	Company         Company           `json:"company" gorm:"foreignKey:companyID"`
	Tags            string            `json:"tags"` // comma separated values
	Status          string            `json:"status"`
	LastModified    int64             `json:"lastModified" gorm:"column:lastModified"`
}

func (v *Vehicle) ToByte() ([]byte, error) {
	return json.Marshal(v)
}

func (v *Vehicle) ToJSONString() (string, error) {
	buf, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}
func (v *Vehicle) Validate() error {
	//TODO
	return nil
}
func (v *Vehicle) ToString() string {
	return fmt.Sprintln(*v)
}
