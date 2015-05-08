package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func init() {
	InitDb()
}

func TestHandleIndexReturnsWithStatusOK(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	handleIndex(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Response body did not contain expected %v:\n\tbody: %v", "200", response.Code)
	}
}

func TestHandleIndexReturnsJSON(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	handleIndex(response, request)

	ct := response.HeaderMap["Content-Type"][0]
	if !strings.EqualFold(ct, "application/json") {
		t.Fatalf("Content-Type does not equal 'application/json'")
	}
}
