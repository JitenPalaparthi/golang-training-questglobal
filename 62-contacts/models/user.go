package models

import "fmt"

type User struct {
	ID           uint64 `json:"id"` // automatically taken as primary key
	UName        string `json:"uName" gorm:"column:uname"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Status       string `json:"status"`
	LastModified int64  `json:"lastModified" gorm:"column:lastModified"`
}

func (u *User) ValidateRegister() error {

	if u.UName == "" {
		return fmt.Errorf("empty UName field")
	}
	if u.Email == "" {
		return fmt.Errorf("empty email field")
	}
	if u.Password == "" {
		return fmt.Errorf("empty password field")
	}
	return nil
}

func (u *User) ValidateLogin() error {

	if u.UName == "" {
		return fmt.Errorf("empty UName field")
	}
	// if u.Email == "" {
	// 	return fmt.Errorf("empty email field")
	// }
	if u.Password == "" {
		return fmt.Errorf("empty password field")
	}
	return nil
}
