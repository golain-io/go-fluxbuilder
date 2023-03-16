package gofluxbuilder

type CountBuilder struct {
	Column string
}

// Validate is the validation impl of Builder for MaxBuilder
func (f CountBuilder) Validate() error {
	return nil
}

// Parse is the parsing impl of Builder for MaxBuilder
func (f CountBuilder) Parse() string {
	return commonGenerator("count", f)
}
