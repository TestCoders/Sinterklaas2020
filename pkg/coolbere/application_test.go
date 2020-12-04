package coolbere

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewApplication(t *testing.T) {
	app := NewApplication()
	assert.NotNil(t, app.infoLog)
	assert.NotNil(t, app.errorLog)
	assert.NotNil(t, app.products)
}

func TestApplication_Error(t *testing.T) {
	w := &bytes.Buffer{}
	app := &Application{
		errorLog: log.New(w, "ERROR:\t", 0),
		infoLog:  nil,
		products: nil,
	}

	app.Error(errors.New("test error"))
	assert.Equal(t, "ERROR:\ttest error\n", w.String())
}

func TestApplication_Error_emptyErr(t *testing.T) {
	w := &bytes.Buffer{}
	app := &Application{
		errorLog: log.New(w, "ERROR:\t", 0),
		infoLog:  nil,
		products: nil,
	}

	app.Error(nil)
	assert.Equal(t, "", w.String())
}

func TestApplication_serverError(t *testing.T) {
	logger := &bytes.Buffer{}

	app := &Application{
		errorLog: log.New(logger, "ERROR:\t", 0),
		infoLog:  log.New(logger, "INFO:\t", 0),
		products: nil,
	}

	w := httptest.NewRecorder()
	app.serverError(w, errors.New("test"))
	assert.True(t, strings.Contains(logger.String(), "ERROR:\ttest"))
	assert.Equal(t, 500, w.Result().StatusCode)
}
