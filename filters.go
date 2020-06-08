package querybuilder

type filter interface {
	addLogical(operator logicalType)
	father() *group
	format(starter int) (string, []interface{})
}

func New() *Filters {
	mainGroup := newFilterGroup(nil)

	return &Filters{
		mainGroup:     mainGroup,
		currentGroup:  mainGroup,
	}
}

// Filters an array of filters
type Filters struct {
	mainGroup    *group
	currentGroup *group
	lastLogical  logicalType
}

func (f *Filters) addLogical(operator logicalType) *Filters {
	f.lastLogical = operator
	return f
}

// New Creates a new single
func (f *Filters) Field(field Field) *single {
	newFilter := newSingleFilter(field, f, f.currentGroup)
	newFilter.addLogical(f.lastLogical)
	f.currentGroup.addFilter(newFilter)
	return newFilter
}

// OpenBracket opens a bracket
func (f *Filters) OpenBracket() *Filters {
	newGroup := newFilterGroup(f.currentGroup)
	newGroup.addLogical(f.lastLogical)
	f.currentGroup.addFilter(newGroup)
	f.lastLogical = none
	return f
}

// CloseBracket closes a group
func (f *Filters) CloseBracket() *Filters {
	if f.currentGroup == f.mainGroup {
		return f
	}

	f.currentGroup = f.currentGroup.father()
	return f
}

// And adds an and condition
func (f *Filters) And() *Filters {
	return f.addLogical(and)
}

// Or adds an or condition
func (f *Filters) Or() *Filters {
	return f.addLogical(or)
}

// Format returns the filters formatted
func (f *Filters) Format() (string, []interface{}) {
	return f.mainGroup.format(1)
}

