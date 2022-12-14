package domain

import (
	"context"
	"errors"
	"fmt"
)

type Category struct {
	ID   CategoryID
	Name CategoryName
}

type CategoryID string
type CategoryName string

const categoryNameLenLimit = 40

func NewCategory(name CategoryName) (*Category, error) {
	if name == "" {
		return nil, errors.New("name must not be nil")
	}
	if len(name) > categoryNameLenLimit {
		return nil, fmt.Errorf("name is too long (%d > %d)", len(name), categoryNameLenLimit)
	}

	return &Category{
		ID:   CategoryID(generateID()),
		Name: CategoryName(name),
	}, nil
}

type CategoryRepository interface {
	Register(context.Context, *Category) error
	Get(context.Context, CategoryID) (*Category, error)
	List(context.Context) ([]*Category, error)
	Delete(context.Context, *Category) error
}
