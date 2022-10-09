package di

import (
	"go.uber.org/dig"

	"github.com/syuparn/sqlboilerpractice/db"
	"github.com/syuparn/sqlboilerpractice/db/category"
	"github.com/syuparn/sqlboilerpractice/usecase"
)

func NewContainer() *dig.Container {
	c := dig.New()

	c.Provide(db.NewClient)

	c.Provide(category.NewCategoryRepository)

	c.Provide(usecase.NewCreateCategoryInputPort)
	c.Provide(usecase.NewListCategoryInputPort)
	c.Provide(usecase.NewDeleteCategoryInputPort)

	return c
}
