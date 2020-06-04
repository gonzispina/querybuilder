package querybuilder

import (
	"fmt"
	"strings"
)

// Field abstraction representation
type Field interface {
	Name() string
	Table() string
	Type() Type
}

// Columns a list of fields
type Fields []Field

// Format all the columns to get queried
func (c Fields) Format() string {
	cols := make([]string, len(c))
	for _, f := range c {
		table := f.Table()
		if table != "" {
			table += table + "."
		}

		cols = append(cols, fmt.Sprintf("%s%s", table, f.Name()))
	}
	return strings.Join(cols, ", ")
}

type commonField string

const (
	count    commonField = "COUNT(*)"
	wildcard commonField = "*"
)

// Name retrieves the name of the column
func (c commonField) Name() string {
	return string(c)
}

// Table retrieves white
func (c commonField) Table() string {
	return ""
}

// Type if not specified retrieves a string
func (c commonField) Type() Type {
	return String
}