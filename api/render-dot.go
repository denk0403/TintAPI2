package handler

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	. "github.com/denk0403/TintAPI2/utils"
	"github.com/goccy/go-graphviz"
)

const MAX_NODES = 250
const MAX_EDGES = MAX_NODES * (MAX_NODES - 1) / 2

// Handles "/api/render-dot" endpoint. Generates a state diagram from DOT syntax.
func HandleRenderDot(w http.ResponseWriter, r *http.Request) {
	var (
		err    error
		layout graphviz.Layout
		format graphviz.Format
	)

	query_params := r.URL.Query()
	layout_str := query_params.Get("layout")
	if layout_str == "" {
		layout = graphviz.SFDP
	} else {
		layout = graphviz.Layout(layout_str)
	}

	format_str := query_params.Get("format")
	if format_str == "" {
		format = graphviz.SVG
	} else {
		format = graphviz.Format(format_str)
	}

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

	node_count, edge_count := graph.NumberNodes(), graph.NumberEdges()
	if node_count > MAX_NODES {
		err_msg := fmt.Sprintf("Too many nodes. A max of %d nodes are allowed. Received %d.", MAX_NODES, node_count)
		WriteClientError(w, err_msg)
		return
	}
	if edge_count > MAX_EDGES {
		err_msg := fmt.Sprintf("Too many edges. A max of %d edges are allowed. Received %d.", MAX_EDGES, edge_count)
		WriteClientError(w, err_msg)
		return
	}

	var buf bytes.Buffer
	err = graphviz.New().SetLayout(layout).Render(graph, format, &buf)

	if err != nil {
		WriteClientError(w, err.Error())
		return
	}

	w.Write(buf.Bytes())
}
