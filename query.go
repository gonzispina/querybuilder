package querybuilder

type Query interface {
	Done() (string, []interface{})
	hasTable() bool
}

type baseQuery struct {
	table Table
}

func (b baseQuery) hasTable() bool {
	return b.table.name != ""
}

type insert struct {
	baseQuery
	fields Fields
	values []string
}

type update struct {
	baseQuery
	fields  Fields
	values  []string
	filters Filters
}

type delete struct {
	baseQuery
	filters Filters
}

func Select(columns ...Column) *sel {
	s := &sel{}
	s.table.columns = columns
	return s
}

func InsertInto(table Table) insert {
	return insert{
		baseQuery: baseQuery{table: table},
	}
}
