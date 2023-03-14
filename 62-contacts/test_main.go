package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"gitlab.stackroute.in/JitenP/golang-training-questglobal/database"
	"gitlab.stackroute.in/JitenP/golang-training-questglobal/models"
)

var (
	//port string
	dsn string //= "host=localhost user=postgres password=postgres dbname=contactsbd port=55432 sslmode=disable TimeZone=Asia/Shanghai"
)

func main() {

	flag.StringVar(&dsn, "dsn", "host=localhost user=postgres password=postgres dbname=contactsbd port=55432 sslmode=disable TimeZone=Asia/Shanghai", "--dsn=host=localhost user=postgres password=postgres dbname=contactsbd port=55432 sslmode=disable TimeZone=Asia/Shanghai")
	flag.Parse()

	db, err := database.Connect(dsn)
	if err != nil {
		log.Fatalln(err) //Fatals cannot be recoverd but panics can. Fatal exit the application with os.Exit(1)
	} else {
		log.Println(db)
	}

	contact := new(models.Contact)
	contact.Name = "Jiten"
	contact.Address = "Bangalore"
	contact.Email = "Jitenp@outlook.Com"
	contact.Status = "Active"
	contact.LastModified = time.Now().Unix()

	cdb := new(database.Contact)
	cdb.DB = db
	fmt.Println(cdb.Add(contact))

	data := make(map[string]any)
	data["phoneNumber"] = "9618558500"
	data["address"] = "Bangalore, Yeshvantpur"
	fmt.Println(cdb.Update(1, data))

	fmt.Println(cdb.DeleteBy(7))

}
