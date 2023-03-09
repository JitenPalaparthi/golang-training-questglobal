package models

import "fmt"

type Contact struct {
	ID           uint64 `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Address      string `json:"address"`
	PhoneNumber  string `json:"phoneNumber"`
	Status       string `json:"status"`
	LastModified uint64 `json:"lastModified"`
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
