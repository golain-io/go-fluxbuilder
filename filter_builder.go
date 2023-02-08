package gofluxbuilder

import "errors"

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

func NewFilterBuilder() *FilterBuilder {
	return &FilterBuilder{}
}

func (f *FilterBuilder) And() *FilterBuilder {
	f.filters = append(f.filters, filterUnit{
		measurement:   "and",
		isConditional: true,
	})
	return f
}
func (f *FilterBuilder) Or() *FilterBuilder {
	f.filters = append(f.filters, filterUnit{
		measurement:   "or",
		isConditional: true,
	})
	return f
}
func (f *FilterBuilder) Equal(key string, value string) *FilterBuilder {
	f.filters = append(f.filters, filterUnit{
		key:           key,
		value:         value,
		measurement:   "==",
		isConditional: false,
	})
	return f
}

func (f *FilterBuilder) NotEqual(key string, value string) *FilterBuilder {
	f.filters = append(f.filters, filterUnit{
		key:           key,
		value:         value,
		measurement:   "!=",
		isConditional: false,
	})
	return f
}
func (f *FilterBuilder) LesserThan(key string, value string) *FilterBuilder {
	f.filters = append(f.filters, filterUnit{
		key:           key,
		value:         value,
		measurement:   "<",
		isConditional: false,
	})
	return f
}
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
func (f *FilterBuilder) GreaterThan(key string, value string) *FilterBuilder {
	f.filters = append(f.filters, filterUnit{
		key:           key,
		value:         value,
		measurement:   ">",
		isConditional: false,
	})
	return f
}
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

func (f *FilterBuilder) Build() Filter {
	var filter []filterUnit
	for _, filters := range f.filters {
		filter = append(filter, filters)
	}
	return filter
}

func (f *FilterBuilder) Validate() error {
	if f.filters == nil {
		return errors.New("nil filters, no filters applied")
	}
	return nil
}

func (f *FilterBuilder) Parse() string {
	return filterGenerator(f.Build())
}
