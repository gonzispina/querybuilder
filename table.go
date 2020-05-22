package querybuilder

import "strings"

type Table struct {
	name    string
	columns Columns
	as      string
}

func (t Table) format() string {
	if t.name == "" {
		return ""
	}

	str := strings.ToLower(t.name)

	if t.as != "" {
		str += " AS " + t.as
	}

	return str
}

func (t Table) prefix() string {
	if t.as != "" {
		return t.as
	}

	return t.name
}
