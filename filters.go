package querybuilder

import (
	"fmt"
	"strings"
)

// Filter is a query filter
type filter struct {
	logic      logic
	relational relational
	field      Field
	value      string
}

// New Creates a new filter
func New(field Field) Filters {
	v := filter{
		field: field,
		logic: none,
	}

	return Filters{v}
}

// Filters an array of filters
type Filters []filter

// Format returns the filters formatted
func (f Filters) Format() string {
	formatted := []string{}
	for _, filter := range f {
		if filter.relational == unknownOperator || filter.logic == unknownOperator {
			continue
		}

		formatted = append(formatted, fmt.Sprintf("%s %s %s %s", filter.logic, filter.field, filter.relational, filter.value))
	}

	joined := strings.Join(formatted, " ")
	replaced := strings.ReplaceAll(joined, "  ", " ")
	trimmed := strings.TrimLeft(replaced, " ")

	return trimmed
}

func (f Filters) canAddOperation() bool {
	lastIndex := len(f) - 1
	return lastIndex >= 0
}

// And adds an and condition
func (f Filters) And(field Field) Filters {
	if !f.canAddOperation() {
		return f
	}

	return append(f, filter{
		logic: and,
		field: field,
	})
}

// Or adds an or condition
func (f Filters) Or(field Field) Filters {
	if !f.canAddOperation() {
		return f
	}

	return append(f, filter{
		logic: or,
		field: field,
	})
}

// Equals adds an equality condition
func (f Filters) EqualTo(value string) Filters {
	if !f.canAddOperation() {
		return f
	}

	lastIndex := len(f) - 1
	lastFilter := f[lastIndex]
	lastFilter.relational = equal
	lastFilter.value = lastFilter.field.Type().format(value)

	return append(f[:lastIndex], lastFilter)
}

// Greater adds a greater condition
func (f Filters) GreaterThan(value string) Filters {
	if !f.canAddOperation() {
		return f
	}

	lastIndex := len(f) - 1
	lastFilter := f[lastIndex]
	lastFilter.relational = greater
	lastFilter.value = lastFilter.field.Type().format(value)

	return append(f[:lastIndex], lastFilter)
}

// Lesser adds a lesser condition
func (f Filters) LesserThan(value string) Filters {
	if !f.canAddOperation() {
		return f
	}

	lastIndex := len(f) - 1
	lastFilter := f[lastIndex]
	lastFilter.relational = lesser
	lastFilter.value = lastFilter.field.Type().format(value)

	return append(f[:lastIndex], lastFilter)
}

// In adds multiple values to equality
func (f Filters) In(values ...string) Filters {
	if !f.canAddOperation() {
		return f
	}

	if len(values) == 0 {
		return f
	}

	lastIndex := len(f) - 1
	lastFilter := f[lastIndex]
	lastFilter.relational = in

	var formatted []string
	for _, value := range values {
		formatted = append(formatted, lastFilter.field.Type().format(value))
	}

	lastFilter.value = "(" + strings.Join(formatted, ", ") + ")"

	return append(f[:lastIndex], lastFilter)
}

// In adds multiple values to equality
func (f Filters) Between(first string, second string) Filters {
	if !f.canAddOperation() {
		return f
	}

	lastIndex := len(f) - 1
	lastFilter := f[lastIndex]
	lastFilter.relational = between
	firstFormatted := lastFilter.field.Type().format(first)
	secondFormatted := lastFilter.field.Type().format(second)
	lastFilter.value = "(" + firstFormatted + " AND " + secondFormatted + ")"

	return append(f[:lastIndex], lastFilter)
}
