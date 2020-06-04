package querybuilder

type orderType string

const (
	asc  orderType = "ASC"
	desc orderType = "DESC"
)

type order struct {
	t      orderType
	column Field
	father *sel
}

func (o *order) format() string {
	if o.column == nil {
		return ""
	}

	if o.t == "" {
		o.t = asc
	}

	return " ORDER BY " + o.column.Name() + " " + string(o.t)
}

func (o *order) Asc() *sel {
	o.t = asc
	return o.father
}

func (o *order) Desc() *sel {
	o.t = desc
	return o.father
}
