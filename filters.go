package querybuilder

import (
	"fmt"
	"strings"
)

type bracket string
const (
	noBracket     bracket = ""
	opener bracket = "("
	closer bracket = ")"
)

// Filter is a query filter
type filter struct {
	logic      logical
	relational relational
	bracket    bracket
}

func (f filter) format(starter int) (string, []interface{}) {
	if f.relational == nil {
		return "", nil
	}

	relation, values := f.relational.format(starter)
	if relation == "" {
		return "", values
	}

	return fmt.Sprintf("%s %s", f.logic.format(), relation), values
}

// New Creates a new filter
func New(field Field) Filters {
	v := filter{
		logic: none.newLogical(field),
		bracket: opener,
	}

	return Filters{v}
}

// Filters an array of filters
type Filters []filter

func (f Filters) addLogical(operator logicalType, field Field) Filters {
	if !f.canAddOperation() {
		return f
	}

	return append(f, filter{
		logic: operator.newLogical(field),
	})
}

func (f Filters) addRelational(operator relationalType, values ...interface{}) Filters {
	if !f.canAddOperation() {
		return f
	}

	lastIndex := len(f) - 1
	lastFilter := f[lastIndex]
	lastFilter.relational = operator.newRelational(values...)

	return append(f[:lastIndex], lastFilter)
}

func (f Filters) canAddOperation() bool {
	lastIndex := len(f) - 1
	return lastIndex >= 0
}
// And adds an and condition
func (f Filters) And(field Field) Filters {
	return f.addLogical(and, field)
}

// Or adds an or condition
func (f Filters) Or(field Field) Filters {
	return f.addLogical(or, field)
}

// Equals adds an equality condition
func (f Filters) EqualTo(value interface{}) Filters {
	return f.addRelational(equal, value)
}

// NotEqualTo adds an inequality condition
func (f Filters) NotEqualTo(value interface{}) Filters {
	return f.addRelational(notEqual, value)
}

// Greater adds a greater condition
func (f Filters) GreaterThan(value interface{}) Filters {
	return f.addRelational(greater, value)
}

// GreaterEqualThan adds a greater or equal condition
func (f Filters) GreaterEqualThan(value interface{}) Filters {
	return f.addRelational(greaterEqual, value)
}

// Lesser adds a lesser condition
func (f Filters) LesserThan(value interface{}) Filters {
	return f.addRelational(lesser, value)
}

// LesserEqualThan adds a lesser or equal condition
func (f Filters) LesserEqualThan(value interface{}) Filters {
	return f.addRelational(lesserEqual, value)
}

// In adds multiple values to equality
func (f Filters) In(values ...interface{}) Filters {
	return f.addRelational(in, values...)
}

// Between adds a between twe values condition
func (f Filters) Between(first interface{}, second interface{}) Filters {
	return f.addRelational(between, first, second)
}

// IsNull adds a null equality condition
func (f Filters) IsNull() Filters {
	return f.addRelational(isNull)
}

// IsNotNull adds a null inequality condition
func (f Filters) IsNotNull() Filters {
	return f.addRelational(isNotNull)
}

// Format returns the filters formatted
func (f Filters) Format() (string, []interface{}) {
	var values []interface{}
	formatted := make([]string, len(f))
	dollarSignCount := 1
	for _, filter := range f {
		filterFormatted, filterValues := filter.format(dollarSignCount)
		if filterFormatted == "" {
			continue
		}

		formatted = append(formatted, filterFormatted)
		values = append(values, filterValues...)
		dollarSignCount += len(filterValues)
	}

	joined := strings.Join(formatted, " ")
	replaced := strings.ReplaceAll(joined, "  ", " ")
	trimmed := strings.TrimLeft(replaced, " ")

	return trimmed, values
}

