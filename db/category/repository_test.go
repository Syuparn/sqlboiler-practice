package category

import (
	"context"
	"database/sql/driver"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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
