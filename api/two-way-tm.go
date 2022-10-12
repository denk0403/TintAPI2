package handler

import (
	"net/http"

	. "github.com/denk0403/TintAPI2/utils"
	"gopkg.in/yaml.v2"

	"github.com/cjcodell1/tint/machine"
	tm "github.com/cjcodell1/tint/machine/turing/ways/two"
)

// Schema for describing the behavior of a Two-way Turing Machine.
type twoTMSpec struct {
	Start       string
	Accept      string
	Reject      string
	Transitions [][]string
}

// Handles "/api/two-way-tm" endpoint. Runs a Two-way Turing Machine program on the given tests.
func HandleTwoWayTM(w http.ResponseWriter, r *http.Request) {
	var err error
	var mach machine.Machine

	submission, shouldContinue := TryParseTintSubmission(w, r)
	if !shouldContinue {
		return
	}

	var spec twoTMSpec

	err = yaml.Unmarshal([]byte(submission.Program), &spec)
	if err != nil {
		WriteClientError(w, err.Error())
		return
	}

	mach, err = tm.MakeTuringMachine(spec.Transitions, spec.Start, spec.Accept, spec.Reject)
	if err != nil {
		WriteClientError(w, err.Error())
		return
	}

	output := RunTests(mach, submission.Tests, submission.Verbose)

	SendOutputResponse(w, output)
}
