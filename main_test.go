package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGreet(t *testing.T) {
	got := Greet("Jenkins")
	want := "Hello, Jenkins!"
	if got != want {
		t.Errorf("Greet() = %q, want %q", got, want)
	}
}

func TestRoutes(t *testing.T) {
	tests := []struct {
		path string
		want string
	}{
		{"/", "Hello, Jenkins!\n"},
		{"/health", "ok\n"},
	}
	for _, tt := range tests {
		rec := httptest.NewRecorder()
		newMux().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, tt.path, nil))
		if rec.Code != http.StatusOK {
			t.Errorf("GET %s: status = %d, want 200", tt.path, rec.Code)
		}
		if got := rec.Body.String(); got != tt.want {
			t.Errorf("GET %s: body = %q, want %q", tt.path, got, tt.want)
		}
	}
}
