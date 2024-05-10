package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func jsonEqual(a, b string) bool {
	var objA, objB interface{}

	err := json.Unmarshal([]byte(a), &objA)
	if err != nil {
		return false
	}
	err = json.Unmarshal([]byte(b), &objB)
	if err != nil {
		return false
	}
	return reflect.DeepEqual(objA, objB)
}

func makeRequest(handler http.HandlerFunc, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}

func TestReadinessHandler(t *testing.T) {
	rr := makeRequest(readinessHandler, "GET", "/v1/readiness")

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	if !jsonEqual(rr.Body.String(), `{"status":"ok"}`) {
		t.Errorf("handler returned wrong body: got %v want %v",
			rr.Body.String(), `{"status":"ok"}`)
	}
}

func TestErrorHandler(t *testing.T) {
	rr := makeRequest(errHandler, "GET", "/v1/err")

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
	if !jsonEqual(rr.Body.String(), `{"error":"Internal Server Error"}`) {
		t.Errorf("handler returned wrong body: got %v want %v",
			rr.Body.String(), `{"error":"Internal Server Error"}`)
	}
}
