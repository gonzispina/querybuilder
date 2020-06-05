package querybuilder

type joins []*join

type joinType string

const (
	right joinType = "RIGHT"
	left  joinType = "LEFT"
)

type joinTable struct {
	*Table
	father *join
}

func (j *joinTable) As(as string) *join {
	j.as = as
	return j.father
}

type join struct {
	table   *joinTable
	filters Filters
	t       joinType
	father  *sel
}

func (j *join) format() string {
	if j.table.name == "" {
		return ""
	}

	str := string(j.t) + " " + j.table.name
	if j.table.as != "" {
		str += " " + j.table.as
	}

	str += " ON (" +  ")"
	return str
}

func (j *join) On(filters Filters) *sel {
	j.filters = filters
	return j.father
}
