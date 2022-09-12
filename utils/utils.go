package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/cjcodell1/tint/machine"
)

type Input struct {
	Program string `json:"program"`
	Tests   string `json:"tests"`
	Verbose bool   `json:"verbose"`
}

type Result struct {
	Status int    `json:"status"`
	Output string `json:"output"`
}

func WriteClientError(w http.ResponseWriter, err string) {
	http.Error(w, err, http.StatusBadRequest)
}

func WriteServerError(w http.ResponseWriter, err string) {
	http.Error(w, err, http.StatusInternalServerError)
}

func DoDumbHTTPStuff(w http.ResponseWriter, r *http.Request) (Input, bool) {
	var input Input

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return input, false
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		WriteClientError(w, "Missing request body")
		return input, false
	}

	err = json.Unmarshal(body, &input)
	if err != nil {
		WriteClientError(w, "Invalid input")
		return input, false
	}

	return input, true
}

func SendOutputResponse(w http.ResponseWriter, output string) {
	result := Result{
		Status: 200,
		Output: output,
	}

	responseBody, err := json.Marshal(result)
	if err != nil {
		WriteServerError(w, "Could not marshal output")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

func RunTests(mach machine.Machine, testsString string, verbose bool) string {
	var outputBuilder strings.Builder
	var err error

	// split tests
	tests := strings.Split(testsString, "\n")
	// remove extra test
	tests = tests[:len(tests)-1]

	totalAccept := 0
	totalReject := 0
	totalError := 0

	// Run all tests
	for _, test := range tests {
		outputBuilder.WriteString(fmt.Sprintf("Simulating with \"%s\".\n", test))

		var stepCount int = 0

		// Step through a single test
		conf := mach.Start(test)
		for {
			if verbose {
				outputBuilder.WriteString(conf.Print())
			}

			// check if accept or reject and break
			if mach.IsAccept(conf) {
				totalAccept += 1
				outputBuilder.WriteString("Accepted.\n\n")
				break
			} else if mach.IsReject(conf) {
				totalReject += 1
				outputBuilder.WriteString("Rejected.\n\n")
				break
			}

			if stepCount > 500 {
				outputBuilder.WriteString("Error: Program took too long or encountered an infinite loop.\n")
				outputBuilder.WriteString("Skipping this test.\n\n")
				totalError += 1
				break
			}

			// step
			conf, err = mach.Step(conf)
			stepCount += 1
			if err != nil {
				outputBuilder.WriteString(fmt.Sprintf("Error: %s\n", err.Error()))
				outputBuilder.WriteString("Skipping this test.\n\n")
				break
			}

		}
	}
	outputBuilder.WriteString(fmt.Sprintf("%d accepted.\n", totalAccept))
	outputBuilder.WriteString(fmt.Sprintf("%d rejected.\n", totalReject))
	outputBuilder.WriteString(fmt.Sprintf("%d errors.\n", totalError))

	return outputBuilder.String()
}
