package database

import (
	"gitlab.stackroute.in/JitenP/golang-training-questglobal/models"
	"gorm.io/gorm"
)

type Contact struct {
	DB *gorm.DB
}

// Two reasons of creating a type.
// 1- I can give methods to the type
// 2- If I write methods to a type, I can write same methods for other types.

func (c *Contact) Add(contact *models.Contact) (*models.Contact, error) {
	if err := c.DB.AutoMigrate(&models.Contact{}); err != nil { // f the dbtable is not existed it creates a table
		return nil, err
	}
	tx := c.DB.Create(contact) // inserts a record into a table

	if tx.Error != nil {
		return nil, tx.Error
	}

	return contact, nil
}
