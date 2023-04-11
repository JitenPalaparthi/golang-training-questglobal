package models

import (
	"encoding/json"
	"fmt"

	"gorm.io/datatypes"
)

type Company struct {
	ID             uint              `json:"id"`
	Name           string            `json:"name"`
	CompID         string            `json:"compID" gorm:"column:compID"`
	City           string            `json:"city"`
	MoreDetails    datatypes.JSONMap `json:"moreDetails" gorm:"column:moreDetails"`
	Email          string            `json:"email"`
	ContactDetails string            `json:"contactDetails" gorm:"column:contactDetails"`
	Tags           string            `json:"tags"` // comma separated values
	Status         string            `json:"status"`
	LastModified   int64             `json:"lastModified" gorm:"column:lastModified"`
}

func (c *Company) ToByte() ([]byte, error) {
	return json.Marshal(c)
}

func (c *Company) ToJSONString() (string, error) {
	buf, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func (c *Company) ToString() string {
	return fmt.Sprintln(*c)
}

func (c *Company) Validate() error {
	//TODO
	return nil
}
