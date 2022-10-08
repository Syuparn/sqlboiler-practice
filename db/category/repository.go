package category

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

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
	return nil, errors.New("err")
}

func (r *categoryRepository) Delete(ctx context.Context, category *domain.Category) error {
	return errors.New("err")
}
