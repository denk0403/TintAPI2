package handler

import (
	"net/http"

	. "github.com/denk0403/TintAPI2/utils"
	"gopkg.in/yaml.v2"

	"github.com/cjcodell1/tint/machine"
	"github.com/cjcodell1/tint/machine/finite/dfa"
)

type dfaSpec struct {
	// These must be export, yaml parser requires it.
	Start       string
	Accepts     []string `yaml:"accept-states"` // renamed to accept-states
	Transitions [][]string
}

func HandleDFA(w http.ResponseWriter, r *http.Request) {
	var err error
	var mach machine.Machine

	input, shouldContinue := DoDumbHTTPStuff(w, r)
	if !shouldContinue {
		return
	}

	var spec dfaSpec

	err = yaml.Unmarshal([]byte(input.Program), &spec)
	if err != nil {
		WriteClientError(w, err.Error())
		return
	}

	mach, err = dfa.MakeDFA(spec.Transitions, spec.Start, spec.Accepts)
	if err != nil {
		WriteClientError(w, err.Error())
		return
	}

	output := RunTests(mach, input.Tests, input.Verbose)

	SendOutputResponse(w, output)
}
