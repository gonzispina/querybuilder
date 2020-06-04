package querybuilder

type relational string

const unknownOperator = ""

const (
	// Equal operator
	equal relational = "="
	// NotEqual Operator
	notEqual relational = "<>"
	// Grater operator
	greater relational = ">"
	// LesserEqual Operator
	greaterEqual relational = ">="
	// Lesser Operator
	lesser relational = "<"
	// LesserEqual Operator
	lesserEqual relational = "<="
	// In operator
	in relational = "IN"
	// Between operator
	between = "BETWEEN"
	// IsNull operator
	isNull = "IS"
	// IsNotNULL operator
	isNotNull = "IS NOT"
)

type logic string

const (
	// None is for the first one
	none logic = " "
	// And logic operator
	and logic = "AND"
	// Or logic operator
	or logic = "OR"
)
