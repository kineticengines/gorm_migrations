package main

import (
	"time"

	"github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Accounts ...
type Accounts struct {
	gorm.Model
	GUID              *string `gorm:"not null;unique;column:guid"`
	FirstName         *string `gorm:"not null;column:first_name"`
	LastName          string  `gorm:"not null;column:last_name"`
	IsAccountBillable *bool   `gorm:"default:true;column:is_account_billable"`

	Active           *bool          `gorm:"default:true"`
	HasAcceptedTerms bool           `gorm:"default:false"`
	AmountPaid       float32        `gorm:"null;type:numeric"`
	AmountDeducted   *float32       `gorm:"type:numeric;default:0"`
	Date             time.Time      `gorm:"not null"`
	Grouped          pq.StringArray `gorm:"type:varchar(64)[];"`
}

func main() {

	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Africa/Nairobi"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	_ = db.AutoMigrate(&Accounts{})

	// Create
	guid := "0afca2aa-7de1-11eb-8398-434ca8dedd68"
	fname := "dx"
	db.Create(&Accounts{FirstName: &fname, LastName: "ter", GUID: &guid})

	// Read
	var account Accounts
	db.First(&account, 1)
	db.First(&account, "guid = ?", guid)

	db.Model(&account).Update("last_name", "new-last-name")

	db.Model(&account).Updates(Accounts{HasAcceptedTerms: true, Date: time.Now()})

	// Delete - delete product
	db.Delete(&account, 1)
}
