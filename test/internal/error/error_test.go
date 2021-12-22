package error

import (
	"github.com/stretchr/testify/assert"
	errors "github.com/tavomartinez88/marvel/internal/error"
	"net/http"
	"testing"
)

func TestError(t *testing.T) {
	e := errors.ClientError{
		HttpStatus: http.StatusInternalServerError,
		Message: "msg",
	}

	assert.NotNil(t, e)
	assert.True(t, true, len(e.Error())>0)
}
