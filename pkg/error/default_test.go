package error

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestIsHttpError(t *testing.T) {
	t.Run("is HttpError", func(t *testing.T) {
		err := NewHttpError(http.StatusBadRequest, fmt.Errorf("test"))
		isHttpError := IsHttpError(err)
		assert.True(t, isHttpError)
	})
	t.Run("is not HttpError", func(t *testing.T) {
		err := fmt.Errorf("test")
		isHttpError := IsHttpError(err)
		assert.False(t, isHttpError)
	})
}
