package domain

import (
	"context"
	"errors"
	"fmt"
)

type Product struct {
	ID         ProductID
	Name       ProductName
	CategoryID CategoryID
}

type ProductID string
type ProductName string

const productNameLenLimit = 40

func NewProduct(name ProductName, categoryID CategoryID) (*Product, error) {
	if name == "" {
		return nil, errors.New("name must not be nil")
	}
	if len(name) > productNameLenLimit {
		return nil, fmt.Errorf("name is too long (%d > %d)", len(name), productNameLenLimit)
	}
	if categoryID == "" {
		return nil, errors.New("categoryID must not be nil")
	}

	return &Product{
		ID:         ProductID(generateID()),
		Name:       name,
		CategoryID: categoryID,
	}, nil
}

type ProductRepository interface {
	Register(context.Context, *Product) error
	Get(context.Context, ProductID) (*Product, error)
	List(context.Context) ([]*Product, error)
	Delete(context.Context, *Product) error
}
