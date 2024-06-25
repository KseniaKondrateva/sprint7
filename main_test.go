package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCorrectRequest(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "expected status code: %d, got %d", http.StatusOK, responseRecorder.Code)

	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body, "expected a non-empty response body")
}
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4

	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // здесь нужно создать запрос к сервису
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	assert.Equal(t, http.StatusOK, responseRecorder.Code, "expected status code: %d, got %d", http.StatusOK, responseRecorder.Code)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")
	assert.Len(t, list, totalCount, "expected cafe count: %d, got %d", totalCount, len(list))
}
func TestMainHandlerWhenInvalidCity(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=4&city=ivanovo", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "expected status code: %d, got %d", http.StatusBadRequest, responseRecorder.Code)

	require.Equal(t, "wrong city value", responseRecorder.Body.String())
}
