package models

import (
	"encoding/json"
	"fmt"
)

type RateMaster struct {
	ID            uint        `json:"id"`
	CompanyID     uint        `json:"companyID" gorm:"column:companyID"`
	Company       Company     `json:"company" gorm:"foreignKey:companyID"`
	VehicleSpecID uint        `json:"vehicleSpecID" gorm:"column:vehicleSpecID"`
	VehicleSpec   VehicleSpec `json:"vehicleSpec" gorm:"foreignKey:vehicleSpecID"`
	Code          string      `json:"code"` //
	// Code fileld has a simple formula
	// It follows [Hour(s) Day(s) Week(s) Month(s) Year(s)]
	// [H D W M Y]
	// [5****] --> rate for 5 hours
	// [*2***] --> rate for 2 days
	// [*2-5***] -> rate for 2 to 5 days
	// [**1-2**] --> rate for 1 -2 weeks
	Description  string  `json:"description"`
	Rate         float32 `json:"rate"`
	ProratedRate float32 `json:"proratedRate" gorm:"column:proratedRate"`
	Currency     string  `json:"currency"`
	Tags         string  `json:"tags"` // comma separated values
	Status       string  `json:"status"`
	LastModified int64   `json:"lastModified" gorm:"column:lastModified"`
}

func (v *RateMaster) ToByte() ([]byte, error) {
	return json.Marshal(v)
}

func (v *RateMaster) ToJSONString() (string, error) {
	buf, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}
func (v *RateMaster) Validate() error {
	//TODO
	return nil
}
func (v *RateMaster) ToString() string {
	return fmt.Sprintln(*v)
}
