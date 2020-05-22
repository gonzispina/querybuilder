package querybuilder

type relational string

const unknownOperator = ""

const (
	// Equal operator
	equal relational = "="
	// Grater operator
	greater relational = ">"
	// Lesser Operator
	lesser relational = "<"
	// In operator
	in relational = "IN"
	// Between logic operator
	between = "BETWEEN"
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
