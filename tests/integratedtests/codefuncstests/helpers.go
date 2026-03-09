package codefuncstests

import "errors"

// errTest is a reusable sentinel error for test cases.
var errTest = errors.New("test-error")

// sampleNamedAction tracks calls for NamedActionFuncWrapper tests.
type namedActionTracker struct {
	CalledWith string
}

func (it *namedActionTracker) Action(name string) {
	it.CalledWith = name
}
