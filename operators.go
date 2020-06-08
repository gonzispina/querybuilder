package querybuilder

import (
	"fmt"
	"strings"
)

type logicalType string

const (
	// None is for the first one
	none logicalType = ""
	// And logicalType operator
	and logicalType = "AND"
	// Or logicalType operator
	or logicalType = "OR"
)

func(t logicalType) format(value string) string {
	methods := map[logicalType]func (t logicalType, value string) string{
		and: baseFormatter,
		or:  baseFormatter,
		none: noneFormatter,

	}

	formatter, _ := methods[t]
	return formatter(t, value)
}

func baseFormatter(t logicalType, value string) string {
	return fmt.Sprintf("%s %s", t, value)
}

func noneFormatter(t logicalType, value string) string {
	return value
}

type relationalType string

const (
	// Equal operator
	equal relationalType = "="
	// NotEqual Operator
	notEqual relationalType = "<>"
	// Grater operator
	greater relationalType = ">"
	// LesserEqual Operator
	greaterEqual relationalType = ">="
	// Lesser Operator
	lesser relationalType = "<"
	// LesserEqual Operator
	lesserEqual relationalType = "<="
	// In operator
	in relationalType = "IN"
	// Between operator
	between = "BETWEEN"
	// IsNull operator
	isNull = "IS"
	// IsNotNULL operator
	isNotNull = "IS NOT"
)

func(t relationalType) newRelational(fieldType Type, values ...interface{}) relational {
	base := &baseRelational{
		relation: t,
		values:   values,
		fieldType: fieldType,
	}

	relations := map[relationalType]relational{
		equal: base,
		notEqual: base,
		greater: base,
		greaterEqual: base,
		lesser: base,
		lesserEqual: base,
		in: &inRelational{base},
		between: &betweenRelational{base},
		isNull: &isNullRelational{base},
		isNotNull: &isNotNullRelational{base},
	}

	relational, _ := relations[t]
	return relational
}

type relational interface {
	format(starter int) (string, []interface{})
}

type baseRelational struct {
	relation relationalType
	values   []interface{}
	fieldType Type
}

func (b *baseRelational) format(starter int) (string, []interface{}) {
	if len(b.values) == 0 {
		return "", b.values
	}

	dollarValue := fmt.Sprintf("$%v", starter)
	return fmt.Sprintf("%s %s", b.relation, b.fieldType.format(dollarValue)), b.values
}

type inRelational struct {
	*baseRelational
}

func (i *inRelational) format(starter int) (string, []interface{}) {
	if len(i.values) == 0 {
		return "", i.values
	}

	formatted := make([]string, len(i.values))
	for index, _ := range i.values {
		dollarValue := fmt.Sprintf("$%v", starter + index )
		formatted[index] = i.fieldType.format(dollarValue)
	}

	return "IN (" + strings.Join(formatted, ", ") + ")", i.values
}

type betweenRelational struct {
	*baseRelational
}

func (b *betweenRelational) format(starter int) (string, []interface{}) {
	if len(b.values) < 2 {
		return "", b.values
	}

	first := b.fieldType.format(fmt.Sprintf("$%v", starter))
	second := b.fieldType.format(fmt.Sprintf("$%v", starter + 1))
	return fmt.Sprintf("BETWEEN (%s AND %s)", first, second), b.values
}

type isNullRelational struct {
	*baseRelational
}

func (i *isNullRelational) format(starter int) (string, []interface{}) {
	return "IS NULL", i.values
}

type isNotNullRelational struct {
	*baseRelational
}

func (i *isNotNullRelational) format(starter int) (string, []interface{}) {
	return "IS NOT NULL", i.values
}

