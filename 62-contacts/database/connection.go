package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	MAX_NUMBEROF_RETRIES uint8 = 10
	MAX_RETRY_DURATION   uint8 = 10
)

// retry logic is very important for microservice based architecture.
// services can be offline at any point of time so . there should be retry log.
func Connect(dsn string) (db *gorm.DB, err error) {
	var count uint8
retry:
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) //gorm.Open(mysql.Open(dsn), &gorm.Config{}) //gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		count++
		if count <= MAX_NUMBEROF_RETRIES {
			log.Println("Trying to connect to database-->", count, "times")
			time.Sleep(time.Second * time.Duration(MAX_RETRY_DURATION))
			goto retry
		}
	}
	return db, err
}