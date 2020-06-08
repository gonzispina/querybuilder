package querybuilder

import "fmt"

type Type int

const (
	// String type
	String Type = iota
	// Date type
	Date
	// Numeric
	Numeric
	// Bool
	Bool
)

func (f Type) format(value string) string {
	formatters := map[Type]formatter{
		Date:   dateFormat,
	}

	format, found := formatters[f]
	if !found {
		return value
	}

	return format(value)
}

type formatter func(value string) string

func dateFormat(value string) string   { return fmt.Sprintf("to_timestamp(%s)", value) }
