package testdata

// ComplexAccount2 ...
type ComplexAccount2 struct {
	Product
	GUID      *string `gorm:"not null;unique;column:guid"`
	FirstName *string `gorm:"not null;column:first_name"`
}

// IsModel ..
func (s *ComplexAccount2) IsModel() bool {
	return true
}
