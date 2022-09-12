package handler

import (
	"net/http"

	. "github.com/denk0403/TintAPI2/utils"
	"gopkg.in/yaml.v2"

	"github.com/cjcodell1/tint/machine"
	tm "github.com/cjcodell1/tint/machine/turing/ways/one"
)

type twoTMSpec struct {
	// These must be exported, yaml parser requires it.
	Start       string
	Accept      string
	Reject      string
	Transitions [][]string
}

func HandleTwoWayTM(w http.ResponseWriter, r *http.Request) {
	var err error
	var mach machine.Machine

	input, shouldContinue := DoDumbHTTPStuff(w, r)
	if !shouldContinue {
		return
	}

	var spec twoTMSpec

	err = yaml.Unmarshal([]byte(input.Program), &spec)
	if err != nil {
		WriteClientError(w, err.Error())
		return
	}

	mach, err = tm.MakeTuringMachine(spec.Transitions, spec.Start, spec.Accept, spec.Reject)
	if err != nil {
		WriteClientError(w, err.Error())
		return
	}

	output := RunTests(mach, input.Tests, input.Verbose)

	SendOutputResponse(w, output)
}
