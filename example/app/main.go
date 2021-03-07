package main

import (
	"time"

	"github.com/kineticengines/gorm-migrations/example/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Africa/Nairobi"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	_ = db.AutoMigrate(&models.Accounts{}, &models.Company{}, &models.User{}, &models.Organisations{}, &models.Credentials{}, &models.Company{})

	// Create
	guid := "0afca2aa-7de1-11eb-8398-434ca8dedd68"
	fname := "dx"
	lname := "ter"
	db.Create(&models.Accounts{FirstName: &fname, LastName: &lname, GUID: &guid})

	// Read
	var account models.Accounts
	db.First(&account, 1)
	db.First(&account, "guid = ?", guid)

	db.Model(&account).Update("last_name", "new-last-name")

	db.Model(&account).Updates(models.Accounts{HasAcceptedTerms: true, Date: time.Now()})

	// Delete - delete product
	db.Delete(&account, 1)
}
