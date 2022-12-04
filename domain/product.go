package domain

import (
	"context"
	"errors"
	"fmt"
)

type Product struct {
	ID         ProductID
	Name       ProductName
	Price      ProducePriceYen
	CategoryID CategoryID
}

type ProductID string
type ProductName string
type ProducePriceYen int64

const productNameLenLimit = 40

func NewProduct(name ProductName, price ProducePriceYen, categoryID CategoryID) (*Product, error) {
	if name == "" {
		return nil, errors.New("name must not be nil")
	}
	if len(name) > productNameLenLimit {
		return nil, fmt.Errorf("name is too long (%d > %d)", len(name), productNameLenLimit)
	}
	if price < 0 {
		return nil, errors.New("price must be positive")
	}
	if categoryID == "" {
		return nil, errors.New("categoryID must not be nil")
	}

	return &Product{
		ID:         ProductID(generateID()),
		Name:       name,
		Price:      ProducePriceYen(price),
		CategoryID: categoryID,
	}, nil
}

type ProductRepository interface {
	Register(context.Context, *Product) error
	Get(context.Context, ProductID) (*Product, error)
	List(context.Context) ([]*Product, error)
	Delete(context.Context, *Product) error
}
