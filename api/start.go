package handler

import (
	"net/http"

	. "github.com/denk0403/TintAPI2/utils"
)

func HandleStart(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	SendOutputResponse(w, "")
}
