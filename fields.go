package querybuilder

import (
	"strings"
)

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

// Field are the possible fields to create filters
type Field interface {
	Column
	Type() Type
}