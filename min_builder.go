package gofluxbuilder

type MinBuilder struct {
	Column string
}

// Validate is the validation impl of Builder for MinBuilder
func (f MinBuilder) Validate() error {
	return nil
}

// Parse is the parsing impl of Builder for MinBuilder
func (f MinBuilder) Parse() string {
	return commonGenerator("min", f)
}
