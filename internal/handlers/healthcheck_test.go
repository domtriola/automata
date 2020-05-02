package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/domtriola/automata/internal/handlers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHealthCheck(t *testing.T) {
	req, err := http.NewRequest("GET", "/healthcheck", nil)
	require.NoError(t, err)

	respRec := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.Healthcheck)

	handler.ServeHTTP(respRec, req)

	assert.HTTPSuccess(t, handler, "GET", "/healthcheck", nil)
}
