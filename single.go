package querybuilder

func newSingleFilter(field Field, main *Filters, father *group) *single {
	return &single{
		field:      field,
		main:       main,
		group:      father,
	}
}

// single is a query single
type single struct {
	logical    logicalType
	relational relational
	field      Field
	values     []interface{}
	group      *group
	main       *Filters
}

func (s *single) addLogical(operator logicalType) {
	s.logical = operator
}

func (s *single) addRelational(operator relationalType, values ...interface{}) *Filters {
	s.relational = operator.newRelational(s.field.Type(), values...)
	return s.main
}

func (s *single) father() *group {
	return s.group
}

func (s *single) format(starter int) (string, []interface{}) {
	if s.relational == nil {
		return "", nil
	}

	relation, values := s.relational.format(starter)
	if relation == "" {
		return "", values
	}

	return s.logical.format(s.field.Name() + " " + relation), values
}


// Equals adds an equality condition
func (s *single) EqualTo(value interface{}) *Filters {
	return s.addRelational(equal, value)
}

// NotEqualTo adds an inequality condition
func (s *single) NotEqualTo(value interface{}) *Filters {
	return s.addRelational(notEqual, value)
}

// Greater adds a greater condition
func (s *single) GreaterThan(value interface{}) *Filters {
	return s.addRelational(greater, value)
}

// GreaterEqualThan adds a greater or equal condition
func (s *single) GreaterEqualThan(value interface{}) *Filters {
	return s.addRelational(greaterEqual, value)
}

// Lesser adds a lesser condition
func (s *single) LesserThan(value interface{}) *Filters {
	return s.addRelational(lesser, value)
}

// LesserEqualThan adds a lesser or equal condition
func (s *single) LesserEqualThan(value interface{}) *Filters {
	return s.addRelational(lesserEqual, value)
}

// In adds multiple values to equality
func (s *single) In(values ...interface{}) *Filters {
	return s.addRelational(in, values...)
}

// Between adds a between twe values condition
func (s *single) Between(first interface{}, second interface{}) *Filters {
	return s.addRelational(between, first, second)
}

// IsNull adds a null equality condition
func (s *single) IsNull() *Filters {
	return s.addRelational(isNull)
}

// IsNotNull adds a null inequality condition
func (s *single) IsNotNull() *Filters {
	return s.addRelational(isNotNull)
}
