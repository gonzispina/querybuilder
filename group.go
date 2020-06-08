package querybuilder

import (
	"strings"
)

func newFilterGroup(father *group) *group {
	return &group{
		group: father,
	}
}

type group struct {
	filters []filter
	logical logicalType
	group   *group
}

func (g *group) addFilter(filter filter) filter {
	g.filters = append(g.filters, filter)
	return g
}

func (g *group) addLogical(operator logicalType) {
	g.logical = operator
}

func (g *group) father() *group {
	return g.group
}

func (g *group) format(starter int) (string, []interface{}) {
	var args []interface{}
	var formatted []string
	for _, filter := range g.filters {
		filterFormatted, filterArgs := filter.format(starter)
		if filterFormatted == "" {
			continue
		}

		starter += len(filterArgs)
		args = append(args, filterArgs...)
		formatted = append(formatted, filterFormatted)
	}

	if len(formatted) == 0 {
		return "", nil
	}

	return g.logical.format("(" + strings.Join(formatted, " ") + ")"), args
}

