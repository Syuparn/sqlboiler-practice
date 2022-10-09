package di

import (
	"go.uber.org/dig"

	"github.com/syuparn/sqlboilerpractice/db"
	"github.com/syuparn/sqlboilerpractice/db/category"
	"github.com/syuparn/sqlboilerpractice/db/product"
	"github.com/syuparn/sqlboilerpractice/usecase"
)

func NewContainer() *dig.Container {
	c := dig.New()

	c.Provide(db.NewClient)

	c.Provide(category.NewCategoryRepository)
	c.Provide(product.NewProductRepository)

	c.Provide(usecase.NewCreateCategoryInputPort)
	c.Provide(usecase.NewListCategoryInputPort)
	c.Provide(usecase.NewDeleteCategoryInputPort)
	c.Provide(usecase.NewCreateProductInputPort)

	return c
}
