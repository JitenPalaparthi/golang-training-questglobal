package models

import "fmt"

type Contact struct {
	ID           uint64 `json:"id"` // automatically taken as primary key
	Name         string `json:"name"`
	Email        string `json:"email"`
	Address      string `json:"address"`
	PhoneNumber  string `json:"phoneNumber" gorm:"column:phoneNumber"` // Otherwise(if not gorm:"column:") the name in the database table is given as phone_number
	Status       string `json:"status"`
	LastModified int64  `json:"lastModified" gorm:"column:lastModified"`
}

func (c *Contact) Validate() error {
	if c.Name == "" {
		return fmt.Errorf("empty name field")
	}
	if c.Email == "" {
		return fmt.Errorf("empty email field")
	}
	if c.PhoneNumber == "" {
		return fmt.Errorf("empty phoneNumber field")
	}
	return nil
}
