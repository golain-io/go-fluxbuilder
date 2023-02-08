package gofluxbuilder

type RangeBuilder struct {
	Start string
	Stop  string
}

// Validate is the validation impl of Builder for RangeBuilder
func (b RangeBuilder) Validate() error {
	if b.Start == "" {
		return throwError(rangeBuilderError, "range needs a start value")
	}
	return nil
}

// Parse is the parsing impl of Builder for RangeBuilder
func (b RangeBuilder) Parse() string {
	return rangeGenerator(b)
}
