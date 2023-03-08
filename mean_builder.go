package gofluxbuilder

type MeanBuilder struct {
	Column string
}

// Validate is the validation impl of Builder for MeanBuilder
func (f MeanBuilder) Validate() error {
	return nil
}

// Parse is the parsing impl of Builder for MeanBuilder
func (f MeanBuilder) Parse() string {
	return commonGenerator("mean", f)
}
