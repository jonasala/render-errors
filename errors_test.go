package rendererr

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrInvalidRequest(t *testing.T) {
	err := ErrInvalidRequest(errors.New("invalid request")).(*ErrResponse)
	assert.Equal(t, http.StatusBadRequest, err.HTTPStatusCode)
	assert.Equal(t, http.StatusText(http.StatusBadRequest), err.StatusText)
	assert.Equal(t, "invalid request", err.ErrorText)
}

func TestErrRender(t *testing.T) {
	err := ErrRender(errors.New("unprocessable entity")).(*ErrResponse)
	assert.Equal(t, http.StatusUnprocessableEntity, err.HTTPStatusCode)
	assert.Equal(t, http.StatusText(http.StatusUnprocessableEntity), err.StatusText)
	assert.Equal(t, "unprocessable entity", err.ErrorText)
}
