package gofluxbuilder

type GroupBuilder struct {
	Columns []string
	Mode    string
}

// Validate is the validation impl of Builder for SortBuilder
func (f GroupBuilder) Validate() error {
	return nil
}

// Parse is the parsing impl of Builder for SortBuilder
func (f GroupBuilder) Parse() string {
	return groupGenerator(f)
}
