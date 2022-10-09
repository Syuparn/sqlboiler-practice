package usecase

import (
	"context"
	"fmt"

	"github.com/syuparn/sqlboilerpractice/domain"
)

type ListProductInputData struct {
}

type ListProductOutputData struct {
	Products []*domain.Product
}

type ListProductInputPort interface {
	Handle(context.Context, *ListProductInputData) (*ListProductOutputData, error)
}

type listProductInteractor struct {
	ProductRepository domain.ProductRepository
}

func NewListProductInputPort(
	productRepository domain.ProductRepository,
) ListProductInputPort {
	return &listProductInteractor{
		ProductRepository: productRepository,
	}
}

func (i *listProductInteractor) Handle(
	ctx context.Context,
	_ *ListProductInputData,
) (*ListProductOutputData, error) {
	products, err := i.ProductRepository.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list product: %w", err)
	}

	return &ListProductOutputData{
		Products: products,
	}, nil
}
