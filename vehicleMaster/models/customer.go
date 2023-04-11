package models

import (
	"encoding/json"
	"fmt"

	"gorm.io/datatypes"
)

type Customer struct {
	ID           uint   `json:"id"`
	Ref          string `json:"ref"`
	FirstName    string `json:"firstName" gorm:"column:firstName"`
	LastName     string `json:"lastName" gorm:"column:lastName"`
	Email        string `json:"email"`
	ContactNo    string `json:"contactNo" gorm:"column:contactNo"`
	Address      `json:"address"`
	MoreDetails  datatypes.JSONMap `json:"moreDetails" gorm:"column:moreDetails"`
	Tags         string            `json:"tags"` // comma separated values
	Status       string            `json:"status"`
	LastModified int64             `json:"lastModified" gorm:"column:lastModified"`
}
type Address struct {
	Line1    string `json:"line1"`
	Line2    string `json:"line2"`
	Street   string `json:"street"`
	City     string `json:"city"`
	Province string `json:"province"`
	Country  string `json:"country"`
	ZipCode  string `json:"zipCode" gorm:"column:zipCode"`
}

func (c *Customer) ToByte() ([]byte, error) {
	return json.Marshal(c)
}

func (c *Customer) ToJSONString() (string, error) {
	buf, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}
func (c *Customer) Validate() error {
	//TODO
	return nil
}
func (c *Customer) ToString() string {
	return fmt.Sprintln(*c)
}
