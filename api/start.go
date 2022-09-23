package handler

import (
	"net/http"

	. "github.com/denk0403/TintAPI2/utils"
)

// Handles "/api/start" endpoint. Awakes the API from idle to prepare to receive requests.
//
// Deprecated: This endpoint is redundant now because all handlers are "serverless".
func HandleStart(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	SendOutputResponse(w, "")
}
