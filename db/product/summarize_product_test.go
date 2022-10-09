package product

import (
	"context"
	"database/sql/driver"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"

	"github.com/syuparn/sqlboilerpractice/domain"
)

func TestSummarizeProduct(t *testing.T) {
	columns := []string{"count"}

	tests := []struct {
		name     string
		category *domain.Category
		query    string
		mockRow  []driver.Value
		expected *domain.CategoryStatistics
	}{
		{
			"summarize product of category book",
			&domain.Category{
				ID:   "C123456789ABCDEFGHJKMNPQRS",
				Name: "book",
			},
			"SELECT COUNT(*) FROM `product` WHERE (`product`.`category_id` = ?)",
			[]driver.Value{5},
			&domain.CategoryStatistics{
				CategoryID:   "C123456789ABCDEFGHJKMNPQRS",
				CategoryName: "book",
				NumProducts:  5,
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
			r := NewSummarizeProductService(db)
			actual, err := r.Summarize(context.TODO(), tt.category)

			// assert
			require.NoError(t, err)
			require.Equal(t, tt.expected, actual)
		})
	}
}
