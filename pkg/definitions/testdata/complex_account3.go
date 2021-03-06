package testdata

import "time"

// ComplexAccount3 ...
type ComplexAccount3 struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

// IsModel ..
func (s *ComplexAccount3) IsModel() bool {
	return true
}
