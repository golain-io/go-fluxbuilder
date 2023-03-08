package gofluxbuilder

type MaxBuilder struct {
	Column string
}

// Validate is the validation impl of Builder for MaxBuilder
func (f MaxBuilder) Validate() error {
	return nil
}

// Parse is the parsing impl of Builder for MaxBuilder
func (f MaxBuilder) Parse() string {
	return commonGenerator("max", f)
}
