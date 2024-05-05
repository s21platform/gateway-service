package auth

import (
	"net/http/httptest"
	"testing"
)

func TestGetAuth(t *testing.T) {
	r := httptest.NewRequest("GET", "/auth/login", nil)
	w := httptest.NewRecorder()

	GetAuth(w, r)

	if status := w.Code; status != 200 {
		t.Errorf("Status code not match. Expected 200: received: %d", status)
	}

	expected := "Work"
	if w.Body.String() != expected {
		t.Errorf("Response not match. Expected %s: received: %s", expected, w.Body.String())
	}
}
