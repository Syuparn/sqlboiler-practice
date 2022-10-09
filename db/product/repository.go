package product

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/samber/lo"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/syuparn/sqlboilerpractice/domain"
	"github.com/syuparn/sqlboilerpractice/models"
)

type productRepository struct {
	db *sql.DB
}

var _ domain.ProductRepository = &productRepository{}

func NewProductRepository(db *sql.DB) domain.ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) Register(ctx context.Context, product *domain.Product) error {
	c := &models.Product{
		ID:         string(product.ID),
		Name:       string(product.Name),
		CategoryID: null.StringFrom(string(product.CategoryID)),
	}

	if err := c.Insert(ctx, r.db, boil.Infer()); err != nil {
		return fmt.Errorf("failed to insert product: %w", err)
	}

	return nil
}

func (r *productRepository) List(ctx context.Context) ([]*domain.Product, error) {
	products, err := models.Products().All(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("failed to list products: %w", err)
	}

	return lo.Map(products, func(c *models.Product, _ int) *domain.Product {
		return &domain.Product{
			ID:         domain.ProductID(c.ID),
			Name:       domain.ProductName(c.Name),
			CategoryID: domain.CategoryID(c.CategoryID.String),
		}
	}), nil
}

func (r *productRepository) Get(ctx context.Context, id domain.ProductID) (*domain.Product, error) {
	product, err := models.Products(
		models.ProductWhere.ID.EQ(string(id)),
	).One(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("failed to get product (id: %s): %w", id, err)
	}

	return &domain.Product{
		ID:         domain.ProductID(product.ID),
		Name:       domain.ProductName(product.Name),
		CategoryID: domain.CategoryID(product.CategoryID.String),
	}, nil
}

func (r *productRepository) Delete(ctx context.Context, product *domain.Product) error {
	c := &models.Product{
		ID:         string(product.ID),
		Name:       string(product.Name),
		CategoryID: null.StringFrom(string(product.CategoryID)),
	}

	if _, err := c.Delete(ctx, r.db); err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	return nil
}
