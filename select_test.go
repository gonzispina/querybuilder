package querybuilder_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectQuery(test *testing.T) {
	test.Run("TestSelectQuery - Select with columns", func(t *testing.T) {
		query, _ := querybuilder.Select(dueDate, userID).From("results").Done()
		expected := "SELECT due_date, user_id FROM results;"
		assert.Equal(t, expected, query)
	})

	test.Run("TestSelectQuery - Select without columns", func(t *testing.T) {
		query, _ := querybuilder.Select().From("results").Done()
		expected := "SELECT * FROM results;"
		assert.Equal(t, expected, query)
	})

	test.Run("TestSelectQuery - Select with columns and orde asc", func(t *testing.T) {
		query, _ := querybuilder.Select().
			From("results").
			OrderBy(dueDate).Asc().
			Done()

		expected := "SELECT * FROM results ORDER BY due_date ASC;"
		assert.Equal(t, expected, query)
	})

	test.Run("TestSelectQuery - Select with columns and orde desc", func(t *testing.T) {
		query, _ := querybuilder.Select().
			From("results").
			OrderBy(dueDate).Desc().
			Done()

		expected := "SELECT * FROM results ORDER BY due_date DESC;"
		assert.Equal(t, expected, query)
	})

	test.Run("TestSelectQuery - Select without columns and no filters and with limit and offset", func(t *testing.T) {
		query, _ := querybuilder.Select().
			From("results").
			Limit(2).
			Offset(5).
			Done()

		expected := "SELECT * FROM results LIMIT 2 OFFSET 5;"
		assert.Equal(t, expected, query)
	})

	test.Run("TestSelectQuery - Select join", func(t *testing.T) {
		query, _ := querybuilder.Select().
			From("results").
			Limit(2).
			Offset(5).
			Done()

		expected := "SELECT * FROM results LIMIT 2 OFFSET 5;"
		assert.Equal(t, expected, query)
	})
}
