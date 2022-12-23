package handler

import (
	"bytes"
	"io"
	"net/http"

	. "github.com/denk0403/TintAPI2/utils"
	"github.com/goccy/go-graphviz"
)

// Handles "/api/dot" endpoint. Generates a state diagram from DOT syntax.
func HandleDot(w http.ResponseWriter, r *http.Request) {
	var err error

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Invalid HTTP method", http.StatusMethodNotAllowed)
		return
	}

	body_bytes, err := io.ReadAll(r.Body)

	if err != nil {
		WriteClientError(w, err.Error())
		return
	}

	graph, err := graphviz.ParseBytes(body_bytes)
	if err != nil {
		WriteClientError(w, err.Error())
		return
	}

	var buf bytes.Buffer
	if err := graphviz.New().Render(graph, graphviz.SVG, &buf); err != nil {
		WriteClientError(w, err.Error())
		return
	}

	w.Write(buf.Bytes())
}
