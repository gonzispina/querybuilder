package querybuilder_test

import (
	"fmt"
	"testing"

	"bitbucket.org/brubank/libs/util"
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

// Type returns the field type
func (f field) Type() Type {
	types := map[field]filter.Type{
		id:       filter.String,
		userID:   filter.String,
		amount:   filter.Numeric,
		dueDate:  filter.Date,
		isActive: filter.Bool,
	}

	t, _ := types[f]
	return t
}

func TestFilters(test *testing.T) {
	test.Run("TestFilters - Equal String Type", func(t *testing.T) {
		uuid := util.UUID()
		got := filter.New(userID).EqualTo(uuid).Format()

		expected := fmt.Sprintf("user_id = '%s'", uuid)
		assert.Equal(t, expected, got)
	})

	test.Run("TestFilters - Equal Date Type", func(t *testing.T) {
		date := "2020-01-01"
		got := filter.New(dueDate).EqualTo(date).Format()

		expected := fmt.Sprintf("due_date = to_timestamp('%s')", date)
		assert.Equal(t, expected, got)
	})

	test.Run("TestFilters - Equal Bool Type", func(t *testing.T) {
		got := filter.New(isActive).EqualTo("true").Format()
		expected := fmt.Sprintf("is_active = true")
		assert.Equal(t, expected, got)
	})

	test.Run("TestFilters - Equal AND In", func(t *testing.T) {
		uuid1 := util.UUID()
		uuid2 := util.UUID()

		got := filter.New(userID).
			EqualTo(uuid1).
			And(id).
			In(uuid1, uuid2).
			Format()

		expected := fmt.Sprintf("user_id = '%s' AND id IN ('%s', '%s')", uuid1, uuid1, uuid2)
		assert.Equal(t, expected, got)
	})

	test.Run("TestFilters - Between OR In", func(t *testing.T) {
		got := filter.New(amount).
			Between("0", "2").
			Or(amount).
			In("5", "6", "7").
			Format()

		expected := fmt.Sprintf("amount BETWEEN (%s AND %s) OR amount IN (%s, %s, %s)", "0", "2", "5", "6", "7")
		assert.Equal(t, expected, got)
	})

	test.Run("TestFilters - Lesser OR Greater", func(t *testing.T) {
		got := filter.New(amount).
			LesserThan("7").
			Or(amount).
			GreaterThan("14").
			Format()

		expected := fmt.Sprintf("amount < 7 OR amount > 14")
		assert.Equal(t, expected, got)
	})

	test.Run("TestFilters - Consecutive relational operations", func(t *testing.T) {
		uuid1 := util.UUID()
		uuid2 := util.UUID()
		got := filter.New(userID).
			EqualTo(uuid1).
			EqualTo("").
			LesserThan(uuid2).
			Format()

		expected := fmt.Sprintf("user_id < '%s'", uuid2)
		assert.Equal(t, expected, got)
	})

	test.Run("TestFilters - Consecutive logical operations", func(t *testing.T) {
		uuid1 := util.UUID()
		uuid2 := util.UUID()
		got := filter.New(userID).
			EqualTo(uuid1).
			Or(userID).
			And(userID).
			EqualTo(uuid2).
			Format()

		expected := fmt.Sprintf("user_id = '%s' AND user_id = '%s'", uuid1, uuid2)
		assert.Equal(t, expected, got)
	})
}
