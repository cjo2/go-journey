package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_ParentEndpoint(t *testing.T) {
	server := &Server{}
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	assert.NoError(t, err)

	handler := http.HandlerFunc(server.ParentHandler)

	handler.ServeHTTP(rr, req)
	fmt.Println(rr.Body)
	assert.Equal(t, http.StatusOK, rr.Code)
}
