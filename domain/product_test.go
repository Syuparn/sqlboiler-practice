package domain

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewProduct(t *testing.T) {
	tests := []struct {
		name        string
		productName ProductName
		categoryID  CategoryID
		expected    *Product
	}{
		{
			"create a new product",
			"pencil",
			"1123456789ABCDEFGHJKMNPQRS",
			&Product{
				ID:         "0123456789ABCDEFGHJKMNPQRS",
				Name:       "pencil",
				CategoryID: "1123456789ABCDEFGHJKMNPQRS",
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
			actual, err := NewProduct(tt.productName, tt.categoryID)

			// assert
			require.NoError(t, err)
			require.Equal(t, tt.expected, actual)
		})
	}
}

func TestNewProductError(t *testing.T) {
	tests := []struct {
		name        string
		productName ProductName
		CategoryID  CategoryID
	}{
		{
			"name must not be nil",
			"",
			"1123456789ABCDEFGHJKMNPQRS",
		},
		{
			"len(name) must not exceed 40",
			ProductName(strings.Repeat("a", 41)),
			"1123456789ABCDEFGHJKMNPQRS",
		},
		{
			"categoryID must not be nil",
			"pencil",
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// run
			actual, err := NewProduct(tt.productName, tt.CategoryID)

			// assert
			require.Error(t, err)
			require.Nil(t, actual)
		})
	}
}
