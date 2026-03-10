package coredynamictests

import (
	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/coretests/results"
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
	{
		Title: "ValueMarshal on nil returns nil bytes and error",
		Func: func(d *coredynamic.Dynamic) bool {
			bytes, err := d.ValueMarshal()
			return bytes == nil && err != nil
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "ReflectSetTo on nil returns error",
		Func: func(d *coredynamic.Dynamic) bool {
			var target string
			err := d.ReflectSetTo(&target)
			return err != nil
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "Deserialize on nil returns invalid and error",
		Func: func(d *coredynamic.Dynamic) bool {
			result, err := d.Deserialize([]byte(`{}`))
			return result != nil && err != nil
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "UnmarshalJSON on nil returns error",
		Func: func(d *coredynamic.Dynamic) bool {
			err := d.UnmarshalJSON([]byte(`{}`))
			return err != nil
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
}
