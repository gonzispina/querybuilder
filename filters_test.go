package querybuilder_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	qb "github.com/gonzispina/querybuilder"
)

type field string

const (
	id       field = "id"
	userID   field = "user_id"
	amount   field = "amount"
	dueDate  field = "due_date"
	isActive field = "is_active"
)

// Name returns a proper string representing the column name
func (f field) Name() string {
	return string(f)
}

// Table returns the table name of the field
func (f field) Table() string {
	return "payments"
}

// Type returns the field type
func (f field) Type() qb.Type {
	types := map[field]qb.Type{
		id:       qb.String,
		userID:   qb.String,
		amount:   qb.Numeric,
		dueDate:  qb.Date,
		isActive: qb.Bool,
	}

	t, _ := types[f]
	return t
}


func TestFilters(test *testing.T) {
	test.Run("Equal String Type", func(t *testing.T) {
		id := "1234"
		got, args := qb.New().Field(userID).EqualTo(id).Format()
		value := args[0].(string)

		assert.Equal(t, "(user_id = $1)", got)
		assert.Equal(t, 1, len(args))
		assert.Equal(t, id, value)
	})

	test.Run("Equal Date Type", func(t *testing.T) {
		date := "2020-01-01"
		got, args := qb.New().Field(dueDate).EqualTo(date).Format()
		value := args[0].(string)

		assert.Equal(t, "(due_date = to_timestamp($1))", got)
		assert.Equal(t, 1, len(args))
		assert.Equal(t, date, value)
	})

	test.Run("Equal Bool Type", func(t *testing.T) {
		got, args := qb.New().Field(isActive).EqualTo(true).Format()
		value := args[0].(bool)

		assert.Equal(t, "(is_active = $1)", got)
		assert.Equal(t, 1, len(args))
		assert.Equal(t, true, value)
	})

	test.Run("Not Equal Date Type", func(t *testing.T) {
		date := "01-01-1995"
		got, args := qb.New().
			Field(dueDate).NotEqualTo(date).
			Format()

		value := args[0].(string)

		assert.Equal(t, "(due_date <> to_timestamp($1))", got)
		assert.Equal(t, 1, len(args))
		assert.Equal(t, date, value)
	})

	test.Run("Equal AND In", func(t *testing.T) {
		id1 := "1234"
		id2 := "5678"

		got, args := qb.New().
			Field(userID).EqualTo(id1).
			And().
			Field(id).In(id1, id2).
			Format()

		value0 := args[0].(string)
		value1 := args[1].(string)
		value2 := args[2].(string)

		assert.Equal(t, "(user_id = $1 AND id IN ($2, $3))", got)
		assert.Equal(t, 3, len(args))
		assert.Equal(t, id1, value0)
		assert.Equal(t, id1, value1)
		assert.Equal(t, id2, value2)
	})

	test.Run("Between OR In", func(t *testing.T) {
		got, args := qb.New().
			Field(amount).Between(1, 2).
			Or().
			Field(amount).In(5, 6, 7).
			Format()

		value0 := args[0].(int)
		value1 := args[1].(int)
		value2 := args[2].(int)
		value3 := args[3].(int)
		value4 := args[4].(int)

		assert.Equal(t,"(amount BETWEEN ($1 AND $2) OR amount IN ($3, $4, $5))", got)
		assert.Equal(t, 5, len(args))
		assert.Equal(t, 1, value0)
		assert.Equal(t, 2, value1)
		assert.Equal(t, 5, value2)
		assert.Equal(t, 6, value3)
		assert.Equal(t, 7, value4)
	})

	test.Run("Lesser OR Greater Equal", func(t *testing.T) {
		got, args := qb.New().
			Field(amount).LesserThan(7).
			Or().
			Field(amount).GreaterEqualThan(14).
			Format()

		value0 := args[0].(int)
		value1 := args[1].(int)

		assert.Equal(t, "(amount < $1 OR amount >= $2)", got)
		assert.Equal(t, 2, len(args))
		assert.Equal(t, 7, value0)
		assert.Equal(t, 14, value1)

	})

	test.Run("Lesser Equal Or Greater", func(t *testing.T) {
		got, args := qb.New().
			Field(amount).LesserEqualThan(7).
			Or().
			Field(amount).GreaterThan(14).
			Format()

		value0 := args[0].(int)
		value1 := args[1].(int)

		assert.Equal(t, "(amount <= $1 OR amount > $2)", got)
		assert.Equal(t, 2, len(args))
		assert.Equal(t, 7, value0)
		assert.Equal(t, 14, value1)
	})

	test.Run("Is null", func(t *testing.T) {
		got, args := qb.New().Field(userID).IsNull().Format()
		assert.Equal(t, 0, len(args))
		assert.Equal(t, "(user_id IS NULL)", got)
	})

	test.Run("Is not null", func(t *testing.T) {
		got, args := qb.New().Field(userID).IsNotNull().Format()
		assert.Equal(t, 0, len(args))
		assert.Equal(t, "(user_id IS NOT NULL)", got)
	})

	test.Run("Is null And Is not null", func(t *testing.T) {
		got, args := qb.New().
			Field(userID).IsNull().
			And().
			Field(id).IsNotNull().
			Format()

		assert.Equal(t, "(user_id IS NULL AND id IS NOT NULL)", got)
		assert.Equal(t, 0, len(args))
	})

	test.Run("Is NOT null consecutive conditions", func(t *testing.T) {
		got, args := qb.New().
			Field(userID).
			IsNotNull().
			Format()

		assert.Equal(t, "(user_id IS NOT NULL)", got)
		assert.Equal(t, 0, len(args))
	})


	test.Run("Consecutive logical operations", func(t *testing.T) {
		id1 := "1234"
		id2 := "5678"
		got, args := qb.New().
			Field(userID).EqualTo(id1).
			Or().
			And().
			Field(userID).EqualTo(id2).
			Format()

		value0 := args[0].(string)
		value1 := args[1].(string)

		assert.Equal(t, "(user_id = $1 AND user_id = $2)", got)
		assert.Equal(t, 2, len(args))
		assert.Equal(t, id1, value0)
		assert.Equal(t, id2, value1)
	})

	test.Run("New with bracket", func(t *testing.T) {
		got, args := qb.New().
			OpenBracket().
				Field(userID).IsNotNull().
			CloseBracket().
			Format()

		assert.Equal(t, "(user_id IS NOT NULL)", got)
		assert.Equal(t, 0, len(args))
	})

	test.Run("Closer without opener", func(t *testing.T) {
		got, args := qb.New().
			Field(userID).IsNotNull().
			CloseBracket().
			Format()

		assert.Equal(t, "(user_id IS NOT NULL)", got)
		assert.Equal(t, 0, len(args))
	})
}
