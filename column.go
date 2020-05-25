package querybuilder

import "strings"

// Column abstraction representation
type Column interface {
	Name() string
}

// Columns a list of fields
type Columns []Column

// Format all the columns to get queried
func (c Columns) Format() string {
	cols := []string{}
	for _, col := range c {
		cols = append(cols, col.Name())
	}
	return strings.Join(cols, ", ")
}

// CommonColumn are the common columns
type CommonColumn string

const (
	// Count common column
	Count    CommonColumn = "COUNT(*)"
	// Wildcard common column
	Wildcard CommonColumn = "*"
)

// Name retrieves the name of the column
func (c CommonColumn) Name() string {
	return string(c)
}
