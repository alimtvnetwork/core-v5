package trydotests

import (
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/coretests/results"
	"github.com/alimtvnetwork/core/internal/trydo"
)

// =============================================================================
// WrappedErr nil receiver test cases
// (migrated from first element of CaseV1 slices in WrappedErr_testcases.go)
// =============================================================================

var wrappedErrNilSafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "IsDefined on nil returns false",
		Func:  (*trydo.WrappedErr).IsDefined,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "IsInvalid on nil returns true",
		Func:  (*trydo.WrappedErr).IsInvalid,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "ErrorString on nil returns empty",
		Func:  (*trydo.WrappedErr).ErrorString,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "ExceptionString on nil returns empty",
		Func:  (*trydo.WrappedErr).ExceptionString,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "String on nil returns empty",
		Func:  (*trydo.WrappedErr).String,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "ExceptionType on nil returns nil",
		Func: func(w *trydo.WrappedErr) bool {
			return w.ExceptionType() == nil
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "IsInvalidException on nil returns true",
		Func:  (*trydo.WrappedErr).IsInvalidException,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "HasErrorOrException on nil returns false",
		Func:  (*trydo.WrappedErr).HasErrorOrException,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "IsBothPresent on nil returns false",
		Func:  (*trydo.WrappedErr).IsBothPresent,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "HasException on nil returns false",
		Func:  (*trydo.WrappedErr).HasException,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
}
