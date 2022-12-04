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
		price       ProducePriceYen
		categoryID  CategoryID
		expected    *Product
	}{
		{
			"create a new product",
			"pencil",
			100,
			"1123456789ABCDEFGHJKMNPQRS",
			&Product{
				ID:         "0123456789ABCDEFGHJKMNPQRS",
				Name:       "pencil",
				Price:      100,
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
			actual, err := NewProduct(tt.productName, tt.price, tt.categoryID)

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
		price       ProducePriceYen
		CategoryID  CategoryID
	}{
		{
			"name must not be nil",
			"",
			100,
			"1123456789ABCDEFGHJKMNPQRS",
		},
		{
			"len(name) must not exceed 40",
			ProductName(strings.Repeat("a", 41)),
			100,
			"1123456789ABCDEFGHJKMNPQRS",
		},
		{
			"price must be positive",
			"pencil",
			-2,
			"1123456789ABCDEFGHJKMNPQRS",
		},
		{
			"categoryID must not be nil",
			"pencil",
			100,
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// run
			actual, err := NewProduct(tt.productName, tt.price, tt.CategoryID)

			// assert
			require.Error(t, err)
			require.Nil(t, actual)
		})
	}
}
