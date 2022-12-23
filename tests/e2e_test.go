// go test -v ./tests/e2e_test.go
package tests

import (
	"io"
	"net/http"
	"testing"
)

func TestTwoWayTME2E(t *testing.T) {
	res, _ := http.Post("https://tint-api2.vercel.app/api/two-way-tm", "application/json", nil)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(string(body))
	}
}
