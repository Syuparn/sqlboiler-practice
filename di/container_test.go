package di

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/syuparn/sqlboilerpractice/usecase"
)

func TestNewContainer(t *testing.T) {
	c := NewContainer()

	err := c.Invoke(func(i usecase.CreateCategoryInputPort) {
		require.NotNil(t, i)
	})
	require.NoError(t, err)
}
