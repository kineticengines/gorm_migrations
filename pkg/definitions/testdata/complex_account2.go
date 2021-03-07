package testdata

// Product2 ...
type Product2 struct {
	Code  string
	Price uint
}

// ComplexAccount2 ...
type ComplexAccount2 struct {
	Product2
	GUID      *string `gorm:"not null;unique;column:guid"`
	FirstName *string `gorm:"not null;column:first_name"`
}

// IsModel ..
func (s *ComplexAccount2) IsModel() bool {
	return true
}
