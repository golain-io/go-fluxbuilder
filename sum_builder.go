package gofluxbuilder

type SumBuilder struct {
	Column string
}

// Validate is the validation impl of Builder for SumBuilder
func (f SumBuilder) Validate() error {
	return nil
}

// Parse is the parsing impl of Builder for SumBuilder
func (f SumBuilder) Parse() string {
	return commonGenerator("sum", f)
}
