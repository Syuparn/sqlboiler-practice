package di

import (
	"go.uber.org/dig"

	"github.com/syuparn/sqlboilerpractice/config"
	"github.com/syuparn/sqlboilerpractice/db"
	"github.com/syuparn/sqlboilerpractice/db/category"
	"github.com/syuparn/sqlboilerpractice/db/product"
	"github.com/syuparn/sqlboilerpractice/usecase"
)

func NewContainer() *dig.Container {
	c := dig.New()

	c.Provide(config.NewConfig)

	c.Provide(db.NewClient)

	c.Provide(category.NewCategoryRepository)
	c.Provide(product.NewProductRepository)
	c.Provide(product.NewSummarizeProductService)

	c.Provide(usecase.NewCreateCategoryInputPort)
	c.Provide(usecase.NewListCategoryInputPort)
	c.Provide(usecase.NewDeleteCategoryInputPort)
	c.Provide(usecase.NewCreateProductInputPort)
	c.Provide(usecase.NewListProductInputPort)
	c.Provide(usecase.NewDeleteProductInputPort)
	c.Provide(usecase.NewCountProductInputPort)

	return c
}
