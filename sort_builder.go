package gofluxbuilder

type SortBuilder struct {
	Columns []string
	Desc    bool
}

// Validate is the validation impl of Builder for SortBuilder
func (f SortBuilder) Validate() error {
	return nil
}

// Parse is the parsing impl of Builder for SortBuilder
func (f SortBuilder) Parse() string {
	return sortGenerator(f)
}
