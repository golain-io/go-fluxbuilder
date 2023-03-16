package gofluxbuilder

import (
	"fmt"
	"reflect"
	"strings"
)

func filterGenerator(filter Filter) string {
	generator := "filter(fn: (r) =>"
	for _, unit := range filter {
		if unit.isConditional {
			generator += fmt.Sprintf(" %s", unit.measurement)
			continue
		}
		generator += fmt.Sprintf(" r.%s %s \"%s\"", unit.key, unit.measurement,
			unit.value)
	}
	generator += ")"
	return generator
}

func pipeGenerator() string {
	return "\n\t|> "
}

func rangeGenerator(data RangeBuilder) string {
	generator := "range("
	v := reflect.ValueOf(data)
	key := v.Type()
	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i).Interface()
		if value == "" {
			continue
		}
		if i == v.NumField()-1 {
			generator += fmt.Sprintf("%s: %v", strings.ToLower(key.Field(i).
				Name),
				value)
			break
		}
		generator += fmt.Sprintf("%s: %s, ", strings.ToLower(key.Field(i).
			Name),
			value)
	}
	generator += ")"
	return generator
}

func commonGenerator(name string, data interface{}) string {
	generator := fmt.Sprintf("%s(", name)
	v := reflect.ValueOf(data)
	key := v.Type()
	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i).Interface()
		if value == "" {
			continue
		}
		if i == v.NumField()-1 {
			generator += fmt.Sprintf("%s: \"%v\"", strings.ToLower(key.Field(i).
				Name),
				value)
			break
		}
		generator += fmt.Sprintf("%s: \"%s\", ", strings.ToLower(key.Field(i).
			Name),
			value)
	}
	generator += ")"
	return generator
}

func sortGenerator(data SortBuilder) string {
	generator := "sort("
	if len(data.Columns) > 0 {
		generator += "columns: ["
		for i, value := range data.Columns {
			if i == len(data.Columns)-1 {
				generator += fmt.Sprintf("\"%v\"", value)
				break
			}
			generator += fmt.Sprintf("\"%v\", ", value)
		}
		generator += "], "
		generator += fmt.Sprintf("desc: %t", data.Desc)
	}
	generator += ")"
	return generator
}

func groupGenerator(data GroupBuilder) string {
	generator := "group("
	if len(data.Columns) > 0 {
		generator += "columns: ["
		for i, value := range data.Columns {
			if i == len(data.Columns)-1 {
				generator += fmt.Sprintf("\"%v\"", value)
				break
			}
			generator += fmt.Sprintf("\"%v\", ", value)
		}
		generator += "], "
		if data.Mode != "" {
			generator += fmt.Sprintf("mode: %s", data.Mode)
		}
	}
	generator += ")"
	return generator
}

func limitGenerator(data LimitBuilder) string {
	generator := fmt.Sprintf("limit(n: %d,", data.N)
	if data.Offset > 0 {
		generator += fmt.Sprintf(" offset: %d", data.Offset)
	}
	generator += ")"
	return generator
}
