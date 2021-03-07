package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

// User ...
type User struct {
	gorm.Model
	Name         string
	Age          uint
	CompanyRefer int
	Company      Company `gorm:"foreignKey:CompanyRefer"`
}

// // IsModel ..
// func (m *User) IsModel() bool {
// 	return true
// }

// Organisations is the parent of each individual account
type Organisations struct {
	GUID             *string `gorm:"not null;unique;column:guid"`
	Name             *string `gorm:"not null;column:name"`
	OrganisationCode *int32  `gorm:"not null;unique;column:organisation_code"`
}

// // IsModel ..
// func (m *Organisations) IsModel() bool {
// 	return true
// }

// Accounts is the child/children of an organisation. An organisation must have at least one
// acoount
type Accounts struct {
	gorm.Model
	GUID              *string        `gorm:"not null;unique;column:guid"`
	FirstName         *string        `gorm:"not null;column:first_name"`
	LastName          *string        `gorm:"not null;column:last_name"`
	OrganisationID    *string        `gorm:"not null;column:organisation_id"`
	IsAccountBillable *bool          `gorm:"default:true;column:is_account_billable"`
	APIKey            *string        `gorm:"not null;type:text;column:api_key"`
	Type              *string        `gorm:"not null;type:varchar(255);column:type"`
	Active            *bool          `gorm:"default:true"`
	HasAcceptedTerms  bool           `gorm:"default:false"`
	AmountPaid        float32        `gorm:"null;type:numeric"`
	AmountDeducted    *float32       `gorm:"type:numeric;default:0"`
	Date              time.Time      `gorm:"not null"`
	Grouped           pq.StringArray `gorm:"type:varchar(64)[];"`
	Company           interface{}    `gorm:"foreignKey:CompanyRefer"`
}

// IsModel ..
// func (m *Accounts) IsModel() bool {
// 	return true
// }

// Credentials holds auth credentials when user logs in via afya notes console
// This table is populated on first sign up
type Credentials struct {
	GUID      string `gorm:"not null;unique;column:guid"`
	AccountID string `gorm:"not null;unique;column:account_id"`
	Password  string `gorm:"type:varchar(255)"`
}

// // IsModel ..
// func (m *Credentials) IsModel() bool {
// 	return true
// }

// Product ...
type Product struct {
	Code    string
	Price   uint
	Company interface{} `gorm:"foreignKey:CompanyRefer;foreignKeyRefField:guid"`
}

// IsModel ..
func (m *Product) IsModel() bool {
	return true
}

// Company ....
type Company struct {
	ID   int
	GUID string `gorm:"not null;unique;column:guid"`
	Name string
}

// IsModel ..
func (m *Company) IsModel() bool {
	return true
}
