package category

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

func TestRegisterCategory(t *testing.T) {
	tests := []struct {
		name          string
		category      *domain.Category
		expectedQuery string
		expectedArgs  []driver.Value
	}{
		{
			"resister a new category",
			&domain.Category{
				ID:   "0123456789ABCDEFGHJKMNPQRS",
				Name: "stationary",
			},
			"INSERT INTO `category` (`id`,`name`) VALUES (?,?)",
			[]driver.Value{
				"0123456789ABCDEFGHJKMNPQRS", "stationary",
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
			r := NewCategoryRepository(db)
			err := r.Register(context.TODO(), tt.category)

			// assert
			require.NoError(t, err)
		})
	}
}

func TestListCategory(t *testing.T) {
	columns := []string{"id", "name"}

	tests := []struct {
		name     string
		query    string
		mockRows [][]driver.Value
		expected []*domain.Category
	}{
		{
			"list all categories",
			"SELECT `category`.* FROM `category`",
			[][]driver.Value{
				{"0123456789ABCDEFGHJKMNPQRS", "stationary"},
				{"1123456789ABCDEFGHJKMNPQRS", "book"},
			},
			[]*domain.Category{
				{
					ID:   "0123456789ABCDEFGHJKMNPQRS",
					Name: "stationary",
				},
				{
					ID:   "1123456789ABCDEFGHJKMNPQRS",
					Name: "book",
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
			r := NewCategoryRepository(db)
			actual, err := r.List(context.TODO())

			// assert
			require.NoError(t, err)
			require.Equal(t, tt.expected, actual)
		})
	}
}
