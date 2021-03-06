package testdata

// SimpleAccount ...
type SimpleAccount struct {
	GUID              *string  `gorm:"not null;unique;column:guid"`
	FirstName         *string  `gorm:"not null;column:first_name"`
	LastName          *string  `gorm:"not null;column:last_name"`
	OrganisationID    *string  `gorm:"not null;column:organisation_id"`
	IsAccountBillable *bool    `gorm:"default:true;column:is_account_billable"`
	APIKey            string   `gorm:"not null;type:text;column:api_key"`
	Type              string   `gorm:"not null;type:varchar(255);column:type"`
	Active            *bool    `gorm:"default:true"`
	HasAcceptedTerms  bool     `gorm:"default:false"`
	AmountPaid        float32  `gorm:"null;type:numeric"`
	AmountDeducted    *float32 `gorm:"type:numeric;default:0"`
}

// IsModel ..
func (s *SimpleAccount) IsModel() bool {
	return true
}
