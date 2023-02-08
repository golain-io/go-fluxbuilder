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
