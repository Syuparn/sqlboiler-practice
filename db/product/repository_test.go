package product

import (
	"context"
	"database/sql/driver"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"

	"github.com/syuparn/sqlboilerpractice/domain"
)

func TestRegisterProduct(t *testing.T) {
	tests := []struct {
		name          string
		product       *domain.Product
		expectedQuery string
		expectedArgs  []driver.Value
	}{
		{
			"resister a new product",
			&domain.Product{
				ID:         "0123456789ABCDEFGHJKMNPQRS",
				Name:       "stationary",
				CategoryID: "1123456789ABCDEFGHJKMNPQRS",
			},
			"INSERT INTO `product` (`id`,`name`,`category_id`) VALUES (?,?,?)",
			[]driver.Value{
				"0123456789ABCDEFGHJKMNPQRS", "stationary", "1123456789ABCDEFGHJKMNPQRS",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock
			db, mock, teardown := prepareDB(t)
			defer teardown()
			mock.ExpectExec(regexp.QuoteMeta(tt.expectedQuery)).
				WithArgs(tt.expectedArgs...).
				WillReturnResult(sqlmock.NewResult(0, 1))

			// run
			r := NewProductRepository(db)
			err := r.Register(context.TODO(), tt.product)

			// assert
			require.NoError(t, err)
		})
	}
}

func TestListProduct(t *testing.T) {
	columns := []string{"id", "name", "category_id"}

	tests := []struct {
		name     string
		query    string
		mockRows [][]driver.Value
		expected []*domain.Product
	}{
		{
			"list all products",
			"SELECT `product`.* FROM `product`",
			[][]driver.Value{
				{"0123456789ABCDEFGHJKMNPQRS", "pencil", "C123456789ABCDEFGHJKMNPQRS"},
				{"1123456789ABCDEFGHJKMNPQRS", "novel", "C223456789ABCDEFGHJKMNPQRS"},
			},
			[]*domain.Product{
				{
					ID:         "0123456789ABCDEFGHJKMNPQRS",
					Name:       "pencil",
					CategoryID: "C123456789ABCDEFGHJKMNPQRS",
				},
				{
					ID:         "1123456789ABCDEFGHJKMNPQRS",
					Name:       "novel",
					CategoryID: "C223456789ABCDEFGHJKMNPQRS",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock
			db, mock, teardown := prepareDB(t)
			defer teardown()
			rows := sqlmock.NewRows(columns)
			lo.ForEach(tt.mockRows, func(row []driver.Value, _ int) {
				rows.AddRow(row...)
			})
			mock.ExpectQuery(regexp.QuoteMeta(tt.query)).
				WillReturnRows(rows)

			// run
			r := NewProductRepository(db)
			actual, err := r.List(context.TODO())

			// assert
			require.NoError(t, err)
			require.Equal(t, tt.expected, actual)
		})
	}
}

func TestGetProduct(t *testing.T) {
	columns := []string{"id", "name", "category_id"}

	tests := []struct {
		name     string
		id       domain.ProductID
		query    string
		mockRow  []driver.Value
		expected *domain.Product
	}{
		{
			"get a product",
			"0123456789ABCDEFGHJKMNPQRS",
			"SELECT `product`.* FROM `product` WHERE (`product`.`id` = ?) LIMIT 1",
			[]driver.Value{"0123456789ABCDEFGHJKMNPQRS", "pencil", "C123456789ABCDEFGHJKMNPQRS"},
			&domain.Product{
				ID:         "0123456789ABCDEFGHJKMNPQRS",
				Name:       "pencil",
				CategoryID: "C123456789ABCDEFGHJKMNPQRS",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock
			db, mock, teardown := prepareDB(t)
			defer teardown()
			rows := sqlmock.NewRows(columns).AddRow(tt.mockRow...)
			mock.ExpectQuery(regexp.QuoteMeta(tt.query)).
				WillReturnRows(rows)

			// run
			r := NewProductRepository(db)
			actual, err := r.Get(context.TODO(), tt.id)

			// assert
			require.NoError(t, err)
			require.Equal(t, tt.expected, actual)
		})
	}
}

func TestDeleteProduct(t *testing.T) {
	tests := []struct {
		name          string
		product       *domain.Product
		expectedQuery string
		expectedArgs  []driver.Value
	}{
		{
			"delete a product",
			&domain.Product{
				ID:         "0123456789ABCDEFGHJKMNPQRS",
				Name:       "pencil",
				CategoryID: "C123456789ABCDEFGHJKMNPQRS",
			},
			"DELETE FROM `product` WHERE `id`=?",
			[]driver.Value{
				"0123456789ABCDEFGHJKMNPQRS",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock
			db, mock, teardown := prepareDB(t)
			defer teardown()
			mock.ExpectExec(regexp.QuoteMeta(tt.expectedQuery)).
				WithArgs(tt.expectedArgs...).
				WillReturnResult(sqlmock.NewResult(0, 1))

			// run
			r := NewProductRepository(db)
			err := r.Delete(context.TODO(), tt.product)

			// assert
			require.NoError(t, err)
		})
	}
}
