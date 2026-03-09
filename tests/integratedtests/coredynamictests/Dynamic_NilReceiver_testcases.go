package coredynamictests

import (
	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/coretests/results"
)

// =============================================================================
// Dynamic nil receiver test cases
// (migrated from standalone CaseV1 variables in Dynamic_testcases.go)
// =============================================================================

var dynamicNilSafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "ClonePtr returns nil on nil receiver",
		Func: func(d *coredynamic.Dynamic) bool {
			return d.ClonePtr() == nil
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "Bytes returns nil,false on nil receiver",
		Func: func(d *coredynamic.Dynamic) bool {
			raw, ok := d.Bytes()
			return raw == nil && !ok
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "ValueNullErr returns error on nil receiver",
		Func: func(d *coredynamic.Dynamic) error {
			return d.ValueNullErr()
		},
		Expected: results.ResultAny{
			Error:    results.ExpectAnyError,
			Panicked: false,
		},
	},
}
