package gofluxbuilder

type LimitBuilder struct {
	N      int64
	Offset int64
}

// Validate is the validation impl of Builder for LimitBuilder
func (f LimitBuilder) Validate() error {
	if f.N < 0 || f.N == 0 {
		return throwError(limitBuilderError, "provided "+
			"N is invalid")
	}
	return nil
}

// Parse is the parsing impl of Builder for LimitBuilder
func (f LimitBuilder) Parse() string {
	return limitGenerator(f)
}
