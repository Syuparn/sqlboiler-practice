package test

import (
	"context"
	"testing"

	"github.com/dolthub/go-mysql-server/sql"
	"github.com/stretchr/testify/require"
	"github.com/syuparn/sqlboilerpractice/usecase"
)

func TestHandleCategory(t *testing.T) {
	sqlSim, err := newSimulator(port)
	if err != nil {
		t.Fatal(err)
	}
	sqlSim.Run()
	defer sqlSim.Stop()

	c := mockContainer()
	ctx := context.Background()

	// create
	create := &usecase.CreateCategoryInputData{
		Name: "stationary",
	}
	err = c.Invoke(func(p usecase.CreateCategoryInputPort) {
		_, err = p.Handle(ctx, create)
		require.NoError(t, err)
	})

	// list
	list := &usecase.ListCategoryInputData{}
	var id string
	err = c.Invoke(func(p usecase.ListCategoryInputPort) {
		out, err := p.Handle(ctx, list)

		require.NoError(t, err)
		require.Equal(t, 1, len(out.Categories))
		require.Equal(t, "stationary", string(out.Categories[0].Name))
		id = string(out.Categories[0].ID)
	})

	// delete
	delete := &usecase.DeleteCategoryInputData{
		ID: id,
	}
	err = c.Invoke(func(p usecase.DeleteCategoryInputPort) {
		_, err := p.Handle(ctx, delete)

		require.NoError(t, err)
	})

	// list again
	list = &usecase.ListCategoryInputData{}
	err = c.Invoke(func(p usecase.ListCategoryInputPort) {
		out, err := p.Handle(ctx, list)

		require.NoError(t, err)
		require.Equal(t, 0, len(out.Categories))
	})
}

func TestHandleProduct(t *testing.T) {
	sqlSim, err := newSimulator(port)
	if err != nil {
		t.Fatal(err)
	}

	// prepare category
	sqlSim.DB.CategoryTable.Insert(sql.NewEmptyContext(), sql.NewRow(
		"0123456789ABCDEFGHJKMNPQRS",
		"stationary",
	))

	sqlSim.Run()
	defer sqlSim.Stop()

	c := mockContainer()
	ctx := context.Background()

	// create
	create := &usecase.CreateProductInputData{
		Name:       "pencil",
		Price:      100,
		CategoryID: "0123456789ABCDEFGHJKMNPQRS",
	}
	err = c.Invoke(func(p usecase.CreateProductInputPort) {
		_, err = p.Handle(ctx, create)
		require.NoError(t, err)
	})

	// list
	list := &usecase.ListProductInputData{}
	var id string
	err = c.Invoke(func(p usecase.ListProductInputPort) {
		out, err := p.Handle(ctx, list)

		require.NoError(t, err)
		require.Equal(t, 1, len(out.Products))
		require.Equal(t, "pencil", string(out.Products[0].Name))
		require.Equal(t, 100, int(out.Products[0].Price))
		require.Equal(t, "0123456789ABCDEFGHJKMNPQRS", string(out.Products[0].CategoryID))
		id = string(out.Products[0].ID)
	})

	// delete
	delete := &usecase.DeleteProductInputData{
		ID: id,
	}
	err = c.Invoke(func(p usecase.DeleteProductInputPort) {
		_, err := p.Handle(ctx, delete)

		require.NoError(t, err)
	})

	// list again
	list = &usecase.ListProductInputData{}
	err = c.Invoke(func(p usecase.ListProductInputPort) {
		out, err := p.Handle(ctx, list)

		require.NoError(t, err)
		require.Equal(t, 0, len(out.Products))
	})
}
