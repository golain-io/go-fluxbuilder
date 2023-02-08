package gofluxbuilder

type RangeBuilder struct {
	Start string
	Stop  string
}

func (b RangeBuilder) Validate() error {
	if b.Start == "" {
		return throwError(rangeBuilderError, "range needs a start value")
	}
	return nil
}

func (b RangeBuilder) Parse() string {
	return rangeGenerator(b)
}
