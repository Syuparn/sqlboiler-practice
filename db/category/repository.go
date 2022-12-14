package category

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/samber/lo"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/syuparn/sqlboilerpractice/domain"
	"github.com/syuparn/sqlboilerpractice/models"
)

type categoryRepository struct {
	db *sql.DB
}

var _ domain.CategoryRepository = &categoryRepository{}

func NewCategoryRepository(db *sql.DB) domain.CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) Register(ctx context.Context, category *domain.Category) error {
	c := &models.Category{
		ID:   string(category.ID),
		Name: string(category.Name),
	}

	if err := c.Insert(ctx, r.db, boil.Infer()); err != nil {
		return fmt.Errorf("failed to insert category: %w", err)
	}

	return nil
}

func (r *categoryRepository) List(ctx context.Context) ([]*domain.Category, error) {
	categories, err := models.Categories().All(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("failed to list categories: %w", err)
	}

	return lo.Map(categories, func(c *models.Category, _ int) *domain.Category {
		return &domain.Category{
			ID:   domain.CategoryID(c.ID),
			Name: domain.CategoryName(c.Name),
		}
	}), nil
}

func (r *categoryRepository) Get(ctx context.Context, id domain.CategoryID) (*domain.Category, error) {
	category, err := models.Categories(
		models.CategoryWhere.ID.EQ(string(id)),
	).One(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("failed to get category (id: %s): %w", id, err)
	}

	return &domain.Category{
		ID:   domain.CategoryID(category.ID),
		Name: domain.CategoryName(category.Name),
	}, nil
}

func (r *categoryRepository) Delete(ctx context.Context, category *domain.Category) error {
	c := &models.Category{
		ID:   string(category.ID),
		Name: string(category.Name),
	}

	if _, err := c.Delete(ctx, r.db); err != nil {
		return fmt.Errorf("failed to delete category: %w", err)
	}

	return nil
}
