package querybuilder

import (
	"strings"
)

type CommonColumn string

const (
	Count    CommonColumn = "COUNT(*)"
	Wildcard CommonColumn = "*"
)

func (c CommonColumn) Name() string {
	return string(c)
}

// Fields a list of fields
type Fields []Field

// Format all the fields to get queried
func (f Fields) Format() string {
	fields := []string{}
	for _, field := range f {
		fields = append(fields, field.Name())
	}
	return strings.Join(fields, ", ")
}

// Fields a list of fields
type Columns []Column

// Format all the columns to get queried
func (c Columns) Format() string {
	cols := []string{}
	for _, col := range c {
		cols = append(cols, col.Name())
	}
	return strings.Join(cols, ", ")
}

// Field are the possible fields to create filters
type Field interface {
	Name() string
	Type() Type
}

type Column interface {
	Name() string
}
