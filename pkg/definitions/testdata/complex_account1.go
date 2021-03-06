package testdata

// Product ...
type Product struct {
	Code  string
	Price uint
}

// ComplexAccount1 ...
type ComplexAccount1 struct {
	Product
}

// IsModel ..
func (s *ComplexAccount1) IsModel() bool {
	return true
}
