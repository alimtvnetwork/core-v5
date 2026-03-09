package coredynamictests

import (
	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/coretests/results"
)

// =============================================================================
// LeftRight nil receiver test cases
// (migrated from first element of CaseV1 slices in LeftRight_testcases.go)
// =============================================================================

var leftRightNilSafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "IsEmpty true on nil receiver",
		Func:  (*coredynamic.LeftRight).IsEmpty,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "HasLeft false on nil receiver",
		Func:  (*coredynamic.LeftRight).HasLeft,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "HasRight false on nil receiver",
		Func:  (*coredynamic.LeftRight).HasRight,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "IsLeftEmpty true on nil receiver",
		Func:  (*coredynamic.LeftRight).IsLeftEmpty,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "IsRightEmpty true on nil receiver",
		Func:  (*coredynamic.LeftRight).IsRightEmpty,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "HasAnyItem false on nil receiver",
		Func:  (*coredynamic.LeftRight).HasAnyItem,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "DeserializeLeft on nil returns nil",
		Func: func(lr *coredynamic.LeftRight) bool {
			return lr.DeserializeLeft() == nil
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "DeserializeRight on nil returns nil",
		Func: func(lr *coredynamic.LeftRight) bool {
			return lr.DeserializeRight() == nil
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "LeftToDynamic on nil returns nil",
		Func: func(lr *coredynamic.LeftRight) bool {
			return lr.LeftToDynamic() == nil
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "RightToDynamic on nil returns nil",
		Func: func(lr *coredynamic.LeftRight) bool {
			return lr.RightToDynamic() == nil
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
}
