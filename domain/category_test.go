package domain

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCategory(t *testing.T) {
	tests := []struct {
		name         string
		categoryName CategoryName
		expected     *Category
	}{
		{
			"create a new category",
			"stationary",
			&Category{
				ID:   "0123456789ABCDEFGHJKMNPQRS",
				Name: "stationary",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// patch ULID generator
			origGen := generateID
			defer func() {
				generateID = origGen
			}()
			generateID = func() string { return "0123456789ABCDEFGHJKMNPQRS" }

			// run
			actual, err := NewCategory(tt.categoryName)

			// assert
			require.NoError(t, err)
			require.Equal(t, tt.expected, actual)
		})
	}
}

func TestNewCategoryError(t *testing.T) {
	tests := []struct {
		name         string
		categoryName CategoryName
	}{
		{
			"name must not be nil",
			"",
		},
		{
			"len(name) must not exceed 40",
			CategoryName(strings.Repeat("a", 41)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// run
			actual, err := NewCategory(tt.categoryName)

			// assert
			require.Error(t, err)
			require.Nil(t, actual)
		})
	}
}
