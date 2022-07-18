package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_api_server_HAndlehello(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	s.Handlehello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello")
}
