package testdata

import (
	"time"

	"gorm.io/gorm"
)

// ComplexAccount4 ...
type ComplexAccount4 struct {
	gorm.Model
	CreatedAt time.Time
	UpdatedAt time.Time
	GUID      *string `gorm:"not null;unique;column:guid"`
	FirstName *string `gorm:"not null;column:first_name"`
}

// IsModel ..
func (s *ComplexAccount4) IsModel() bool {
	return true
}
