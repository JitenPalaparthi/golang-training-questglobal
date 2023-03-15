package database

import (
	"fmt"

	"gitlab.stackroute.in/JitenP/golang-training-questglobal/models"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

// Two reasons of creating a type.
// 1- I can give methods to the type
// 2- If I write methods to a type, I can write same methods for other types.

func (c *User) Register(user *models.User) (*models.User, error) {
	if err := c.DB.AutoMigrate(&models.User{}); err != nil { // f the dbtable is not existed it creates a table
		return nil, err
	}
	tx := c.DB.Create(user) // inserts a record into a table

	if tx.Error != nil {
		return nil, tx.Error
	}
	user.Password = "***********"
	return user, nil
}

func (c *User) Login(username, password string) (err error) {
	user := new(models.User)
	tx := c.DB.Where("uname=? and password=?", username, password).First(&user)
	if tx.Error != nil {
		return tx.Error
	}
	if user.ID == 0 || user.UName == "" {
		return fmt.Errorf("failed to login")
	}
	return nil
}
