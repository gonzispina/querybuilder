package querybuilder_test

import (
	"fmt"
	"testing"

	"github.com/gonzispina/querybuilder"
	"github.com/stretchr/testify/assert"
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
func (f field) Type() querybuilder.Type {
	types := map[field]querybuilder.Type{
		id:       querybuilder.String,
		userID:   querybuilder.String,
		amount:   querybuilder.Numeric,
		dueDate:  querybuilder.Date,
		isActive: querybuilder.Bool,
	}

	t, _ := types[f]
	return t
}


func TestFilters(test *testing.T) {
	test.Run("Equal String Type", func(t *testing.T) {
		id := "1234"
		got := querybuilder.New(userID).EqualTo(id).Format()

		expected := fmt.Sprintf("user_id = '%s'", id)
		assert.Equal(t, expected, got)
	})

	test.Run("Equal Date Type", func(t *testing.T) {
		date := "2020-01-01"
		got := querybuilder.New(dueDate).EqualTo(date).Format()

		expected := fmt.Sprintf("due_date = to_timestamp('%s')", date)
		assert.Equal(t, expected, got)
	})

	test.Run("Equal Bool Type", func(t *testing.T) {
		got := querybuilder.New(isActive).EqualTo("true").Format()
		expected := fmt.Sprintf("is_active = true")
		assert.Equal(t, expected, got)
	})

	test.Run("Not Equal Date Type", func(t *testing.T) {
		date := "01-01-1995"
		got := querybuilder.New(dueDate).NotEqualTo(date).Format()
		expected := fmt.Sprintf("due_date <> to_timestamp('%s')", date)
		assert.Equal(t, expected, got)
	})

	test.Run("Equal AND In", func(t *testing.T) {
		id1 := "1234"
		id2 := "5678"

		got := querybuilder.New(userID).
			EqualTo(id1).
			And(id).
			In(id1, id2).
			Format()

		expected := fmt.Sprintf("user_id = '%s' AND id IN ('%s', '%s')", id1, id1, id2)
		assert.Equal(t, expected, got)
	})

	test.Run("Between OR In", func(t *testing.T) {
		got := querybuilder.New(amount).
			Between("0", "2").
			Or(amount).
			In("5", "6", "7").
			Format()

		expected := fmt.Sprintf("amount BETWEEN (%s AND %s) OR amount IN (%s, %s, %s)", "0", "2", "5", "6", "7")
		assert.Equal(t, expected, got)
	})

	test.Run("Lesser OR Greater Equal", func(t *testing.T) {
		got := querybuilder.New(amount).
			LesserThan("7").
			Or(amount).
			GreaterEqualThan("14").
			Format()

		expected := fmt.Sprintf("amount < 7 OR amount >= 14")
		assert.Equal(t, expected, got)
	})

	test.Run("Lesser Equal Or Greater", func(t *testing.T) {
		got := querybuilder.New(amount).
			LesserEqualThan("7").
			Or(amount).
			GreaterThan("14").
			Format()

		expected := fmt.Sprintf("amount <= 7 OR amount > 14")
		assert.Equal(t, expected, got)
	})

	test.Run("Is null", func(t *testing.T) {
		got := querybuilder.New(userID).IsNull().Format()
		assert.Equal(t, "user_id IS NULL", got)
	})

	test.Run("Is not null", func(t *testing.T) {
		got := querybuilder.New(userID).IsNotNull().Format()
		assert.Equal(t, "user_id IS NOT NULL", got)
	})

	test.Run("Is null And Is not null", func(t *testing.T) {
		got := querybuilder.
			New(userID).IsNull().
			And(id).IsNotNull().
			Format()

		assert.Equal(t, "user_id IS NULL AND id IS NOT NULL", got)
	})

	test.Run("Is NOT null consecutive conditions", func(t *testing.T) {
		got := querybuilder.New(userID).
			IsNotNull().
			IsNull().
			IsNotNull().
			Format()
		
		assert.Equal(t, "user_id IS NOT NULL", got)
	})

	test.Run("Consecutive relational operations", func(t *testing.T) {
		id1 := "1234"
		id2 := "5678"
		got := querybuilder.New(userID).
			EqualTo(id1).
			EqualTo("").
			LesserThan(id2).
			Format()

		expected := fmt.Sprintf("user_id < '%s'", id2)
		assert.Equal(t, expected, got)
	})

	test.Run("Consecutive logical operations", func(t *testing.T) {
		id1 := "1234"
		id2 := "5678"
		got := querybuilder.New(userID).
			EqualTo(id1).
			Or(userID).
			And(userID).
			EqualTo(id2).
			Format()

		expected := fmt.Sprintf("user_id = '%s' AND user_id = '%s'", id1, id2)
		assert.Equal(t, expected, got)
	})
}
