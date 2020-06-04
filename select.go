package querybuilder

type sel struct {
	baseQuery
	fields Fields
	joins  joins
	filter *Filters
	order  *order
	limit  *limit
}

func (s *sel) From(table string) *sel {
	s.table.name = table
	return s
}

func (s *sel) As(as string) *sel {
	s.table.as = as
	return s
}

func (s *sel) LeftJoin(table string) *joinTable {
	j := &join{t: left}
	j.table = &joinTable{
		Table: &Table{
			name:    table,
			as:      "",
		},
		father: j,
	}

	s.joins = append(s.joins, j)
	return j.table
}

func (s *sel) RightJoin(table string) *joinTable {
	j := &join{t: right}
	j.table = &joinTable{
		Table: &Table{
			name:    table,
			as:      "",
		},
		father: j,
	}

	s.joins = append(s.joins, j)
	return j.table
}

func (s *sel) Where(filters Filters) *sel {
	if !s.hasTable() {
		return s
	}

	s.filter = &filters
	return s
}

func (s *sel) OrderBy(column Field) *order {
	s.order = &order{
		column: column,
		father: s,
	}

	return s.order
}

func (s *sel) Limit(limit int) *sel {
	if !s.hasTable() {
		return s
	}

	s.limit.limit = limit
	return s
}

func (s *sel) Offset(offset int) *sel {
	if !s.hasTable() {
		return s
	}

	s.limit.offset = offset
	return s
}

func (s *sel) Done() (string, interface{}) {
	if !s.hasTable() {
		return "", nil
	}

	if len(s.fields) == 0 {
		s.fields = Fields{wildcard}
	}

	query := "SELECT " + s.fields.Format()
	query += " FROM " + s.table.format()

	if s.joins != nil && len(s.joins) > 0 {
		for _, j := range s.joins {
			query += j.format()
		}
	}

	if s.order != nil {
		query += s.order.format()
	}

	if s.limit != nil {
		query += s.limit.format()
	}

	query += ";"
	return query, nil
}
