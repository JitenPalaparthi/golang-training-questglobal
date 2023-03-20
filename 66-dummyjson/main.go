package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	bytes, err := os.ReadFile("dummy.json")
	if err != nil {
		fmt.Println(err)
	}

	pd := new(PersonData)

	err = json.Unmarshal(bytes, pd)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pd)
}

type PersonData struct {
	Id        string `json:"_id"`
	Name      string `json:"name"`
	DOB       string `json:"dob"`
	Telephone string `json:"telephone"`
	Email     string `json:"email"`
	Url       string `json:"url"`

	Score       float32 `json:"score"`
	Description string  `json:"description"`
	Verified    bool    `json:"verified"`
	Salary      float32 `json:"salary"`

	Address struct {
		Street   string `json:"street"`
		Town     string `json:"town"`
		PostCode string `json:"postcode"`
	} `json:"address"`

	Pets []string `json:"pets"`
}
