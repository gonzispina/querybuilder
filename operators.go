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

func(t logicalType) newLogical(field Field) logical {
	base := &baseLogical{
		logic: t,
		field: field,
	}

	logics := map[logicalType]logical{
		and: base,
		or:  base,
		none: &noneLogical{base},

	}

	logical, _ := logics[t]
	return logical
}

type logical interface {
	format() string
}

type baseLogical struct {
	logic logicalType
	field Field
}

func (b *baseLogical) format() string {
	return fmt.Sprintf("%s %s", b.logic, b.field.Name())
}

type noneLogical struct {
	*baseLogical
}

func (n *noneLogical) format() string {
	return n.field.Name()
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

func(t relationalType) newRelational(values ...interface{}) relational {
	base := &baseRelational{
		relation: t,
		values:   values,
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
}

func (b *baseRelational) format(starter int) (string, []interface{}) {
	if len(b.values) == 0 {
		return "", b.values
	}

	return fmt.Sprintf("%s $%v", b.relation, starter), b.values
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
		 formatted[index] = fmt.Sprintf("$%v", starter + index )
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
	return fmt.Sprintf("BETWEEN ($%v AND $%v", starter, starter + 1), b.values
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

