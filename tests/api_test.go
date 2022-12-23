// go test -v ./tests/api_test.go
package tests

import (
	"io"
	"net/http/httptest"
	"testing"

	handler "github.com/denk0403/TintAPI2/api"
)

func TestDFA(t *testing.T) {
	req := httptest.NewRequest("POST", "https://tint-api2.vercel.app/api/dfa", nil)
	w := httptest.NewRecorder()
	handler.HandleDFA(w, req)

	body, err := io.ReadAll(w.Result().Body)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(string(body))
	}
}

func TestOneWayTM(t *testing.T) {
	req := httptest.NewRequest("POST", "https://tint-api2.vercel.app/api/one-way-tm", nil)
	w := httptest.NewRecorder()
	handler.HandleDFA(w, req)

	body, err := io.ReadAll(w.Result().Body)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(string(body))
	}
}

func TestTwoWayTM(t *testing.T) {
	req := httptest.NewRequest("POST", "https://tint-api2.vercel.app/api/two-way-tm", nil)
	w := httptest.NewRecorder()
	handler.HandleDFA(w, req)

	body, err := io.ReadAll(w.Result().Body)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(string(body))
	}
}
