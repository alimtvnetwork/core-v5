package reflectmodeltests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/coretests/results"
	"gitlab.com/auk-go/core/reflectcore/reflectmodel"
)

// =============================================================================
// ReflectValueKind nil receiver test cases (migrated from inline t.Error tests)
// =============================================================================

var reflectValueKindNilReceiverTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "IsInvalid on nil returns true",
		Func:  (*reflectmodel.ReflectValueKind).IsInvalid,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "HasError on nil returns false",
		Func:  (*reflectmodel.ReflectValueKind).HasError,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "IsEmptyError on nil returns true",
		Func:  (*reflectmodel.ReflectValueKind).IsEmptyError,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "ActualInstance on nil returns nil",
		Func:  (*reflectmodel.ReflectValueKind).ActualInstance,
		Expected: results.ResultAny{
			Value:    "<nil>",
			Panicked: false,
		},
	},
	{
		Title: "PkgPath on nil returns empty",
		Func:  (*reflectmodel.ReflectValueKind).PkgPath,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "TypeName on nil returns empty",
		Func:  (*reflectmodel.ReflectValueKind).TypeName,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "PointerRv on nil returns nil",
		Func:  (*reflectmodel.ReflectValueKind).PointerRv,
		Expected: results.ResultAny{
			Panicked: false,
		},
	},
	{
		Title: "PointerInterface on nil returns nil",
		Func:  (*reflectmodel.ReflectValueKind).PointerInterface,
		Expected: results.ResultAny{
			Value:    "<nil>",
			Panicked: false,
		},
	},
}
