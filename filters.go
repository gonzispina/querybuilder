package querybuilder

import (
	"fmt"
	"strings"
)

type formatterFunc func(fieldType Type, values ...string) string

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

func (f Filters) addLogical(operator logic, field Field) Filters {
	if !f.canAddOperation() {
		return f
	}

	return append(f, filter{
		logic: operator,
		field: field,
	})
}

func (f Filters) addRelational(operator relational, formatter formatterFunc, values ...string) Filters {
	if !f.canAddOperation() {
		return f
	}

	if len(values) == 0 {
		return f
	}

	lastIndex := len(f) - 1
	lastFilter := f[lastIndex]
	lastFilter.relational = operator
	lastFilter.value = formatter(lastFilter.field.Type(), values...)

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
func (f Filters) EqualTo(value string) Filters {
	return f.addRelational(equal, func(fieldType Type, values ...string) string {
		return fieldType.format(values[0])
	}, value)
}

// NotEqualTo adds an inequality condition
func (f Filters) NotEqualTo(value string) Filters {
	return f.addRelational(notEqual, func(fieldType Type, values ...string) string {
		return fieldType.format(values[0])
	}, value)
}

// Greater adds a greater condition
func (f Filters) GreaterThan(value string) Filters {
	return f.addRelational(greater, func(fieldType Type, values ...string) string {
		return fieldType.format(values[0])
	}, value)
}

// GreaterEqualThan adds a greater or equal condition
func (f Filters) GreaterEqualThan(value string) Filters {
	return f.addRelational(greaterEqual, func(fieldType Type, values ...string) string {
		return fieldType.format(values[0])
	}, value)
}

// Lesser adds a lesser condition
func (f Filters) LesserThan(value string) Filters {
	return f.addRelational(lesser, func(fieldType Type, values ...string) string {
		return fieldType.format(values[0])
	}, value)
}

// LesserEqualThan adds a lesser or equal condition
func (f Filters) LesserEqualThan(value string) Filters {
	return f.addRelational(lesserEqual, func(fieldType Type, values ...string) string {
		return fieldType.format(values[0])
	}, value)
}


// In adds multiple values to equality
func (f Filters) In(values ...string) Filters {
	return f.addRelational(in, func(fieldType Type, v ...string) string {
		var formatted []string
		for _, value := range v {
			formatted = append(formatted, fieldType.format(value))
		}

		return "(" + strings.Join(formatted, ", ") + ")"
	}, values...)
}

// Between adds a between twe values condition
func (f Filters) Between(first string, second string) Filters {
	return f.addRelational(between, func(fieldType Type, values ...string) string {
		firstFormatted := fieldType.format(values[0])
		secondFormatted := fieldType.format(values[1])
		return "(" + firstFormatted + " AND " + secondFormatted + ")"
	}, first, second)
}

// IsNull adds a null equality condition
func (f Filters) IsNull() Filters {
	return f.addRelational(isNull, func(fieldType Type, values ...string) string {
		return "NULL"
	}, "")
}

// IsNotNull adds a null inequality condition
func (f Filters) IsNotNull() Filters {
	return f.addRelational(isNotNull, func(fieldType Type, values ...string) string {
		return "NULL"
	}, "")
}


