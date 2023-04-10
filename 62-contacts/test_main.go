package main

import (
	"flag"
	"fmt"
	"log"

	"gitlab.stackroute.in/JitenP/golang-training-questglobal/database"
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

	err = database.Init(db)
	if err != nil {
		fmt.Println(err)
	}
	// if err := db.AutoMigrate(&models.Contact{}, &models.User{}, &models.Technology{}, &models.Tip{}, &models.UserInterests{}, &models.TestMaster{}, &models.QuestionMaster{}, &models.AnswerMaster{}, &models.UserTest{}); err != nil { // f the dbtable is not existed it creates a table
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("all tables are created")
	// }

	// contact := new(models.Contact)
	// contact.Name = "Jiten"
	// contact.Address = "Bangalore"
	// contact.Email = "Jitenp@outlook.Com"
	// contact.Status = "Active"
	// contact.LastModified = time.Now().Unix()

	// cdb := new(database.Contact)
	// cdb.DB = db
	// fmt.Println(cdb.Add(contact))

	// data := make(map[string]any)
	// data["phoneNumber"] = "9618558500"
	// data["address"] = "Bangalore, Yeshvantpur"
	// fmt.Println(cdb.Update(1, data))

	// fmt.Println(cdb.DeleteBy(7))

}
