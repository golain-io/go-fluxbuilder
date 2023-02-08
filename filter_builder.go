package gofluxbuilder

type filterUnit struct {
	key           string
	value         string
	measurement   string
	isConditional bool
}

type Filter []filterUnit

type FilterBuilder struct {
	filters []filterUnit
}

// NewFilterBuilder allows to create Filters for flux query
func NewFilterBuilder() *FilterBuilder {
	return &FilterBuilder{}
}

// And conditional allows to define "and" in flux filter
func (f *FilterBuilder) And() *FilterBuilder {
	f.filters = append(f.filters, filterUnit{
		measurement:   "and",
		isConditional: true,
	})
	return f
}

// Or conditional allows to define "or" in flux filter
func (f *FilterBuilder) Or() *FilterBuilder {
	f.filters = append(f.filters, filterUnit{
		measurement:   "or",
		isConditional: true,
	})
	return f
}

// Equal measurement allows to define "==" in flux filter
func (f *FilterBuilder) Equal(key string, value string) *FilterBuilder {
	f.filters = append(f.filters, filterUnit{
		key:           key,
		value:         value,
		measurement:   "==",
		isConditional: false,
	})
	return f
}

// NotEqual measurement allows to define "!=" in flux filter
func (f *FilterBuilder) NotEqual(key string, value string) *FilterBuilder {
	f.filters = append(f.filters, filterUnit{
		key:           key,
		value:         value,
		measurement:   "!=",
		isConditional: false,
	})
	return f
}

// LesserThan measurement allows to define "<" in flux filter
func (f *FilterBuilder) LesserThan(key string, value string) *FilterBuilder {
	f.filters = append(f.filters, filterUnit{
		key:           key,
		value:         value,
		measurement:   "<",
		isConditional: false,
	})
	return f
}

// LesserThanEqualTo measurement allows to define "<=" in flux filter
func (f *FilterBuilder) LesserThanEqualTo(key string,
	value string) *FilterBuilder {
	f.filters = append(f.filters, filterUnit{
		key:           key,
		value:         value,
		measurement:   "<=",
		isConditional: false,
	})
	return f
}

// GreaterThan measurement allows to define ">" in flux filter
func (f *FilterBuilder) GreaterThan(key string, value string) *FilterBuilder {
	f.filters = append(f.filters, filterUnit{
		key:           key,
		value:         value,
		measurement:   ">",
		isConditional: false,
	})
	return f
}

// GreaterThanEqualTo measurement allows to define ">=" in flux filter
func (f *FilterBuilder) GreaterThanEqualTo(key string,
	value string) *FilterBuilder {
	f.filters = append(f.filters, filterUnit{
		key:           key,
		value:         value,
		measurement:   ">=",
		isConditional: false,
	})
	return f
}

// Build allows to build the filter
func (f *FilterBuilder) Build() Filter {
	var filter []filterUnit
	for _, filters := range f.filters {
		filter = append(filter, filters)
	}
	return filter
}

// Validate is the validation impl of Builder for FilterBuilder
func (f *FilterBuilder) Validate() error {
	if f.filters == nil {
		return throwError(filterBuilderError, "nil filters, "+
			"no filters applied")
	}
	return nil
}

// Parse is the parsing impl of Builder for FilterBuilder
func (f *FilterBuilder) Parse() string {
	return filterGenerator(f.Build())
}
