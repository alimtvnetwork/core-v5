package reflectmodeltests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/reflectcore/reflectmodel"
)

// =============================================================================
// ReflectValueKind nil receiver test cases (migrated from inline t.Error tests)
// =============================================================================

var reflectValueKindNilReceiverTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "IsInvalid on nil returns true",
		Func:  (*reflectmodel.ReflectValueKind).IsInvalid,
		Expected: args.Map{
			"value":    "true",
			"panicked": false,
		},
	},
	{
		Title: "HasError on nil returns false",
		Func:  (*reflectmodel.ReflectValueKind).HasError,
		Expected: args.Map{
			"value":    "false",
			"panicked": false,
		},
	},
	{
		Title: "IsEmptyError on nil returns true",
		Func:  (*reflectmodel.ReflectValueKind).IsEmptyError,
		Expected: args.Map{
			"value":    "true",
			"panicked": false,
		},
	},
	{
		Title: "ActualInstance on nil returns nil",
		Func:  (*reflectmodel.ReflectValueKind).ActualInstance,
		Expected: args.Map{
			"value":    "<nil>",
			"panicked": false,
		},
	},
	{
		Title: "PkgPath on nil returns empty",
		Func:  (*reflectmodel.ReflectValueKind).PkgPath,
		Expected: args.Map{
			"value":    "",
			"panicked": false,
		},
	},
	{
		Title: "TypeName on nil returns empty",
		Func:  (*reflectmodel.ReflectValueKind).TypeName,
		Expected: args.Map{
			"value":    "",
			"panicked": false,
		},
	},
	{
		Title: "PointerRv on nil returns nil",
		Func:  (*reflectmodel.ReflectValueKind).PointerRv,
		Expected: args.Map{
			"panicked": false,
		},
	},
	{
		Title: "PointerInterface on nil returns nil",
		Func:  (*reflectmodel.ReflectValueKind).PointerInterface,
		Expected: args.Map{
			"value":    "<nil>",
			"panicked": false,
		},
	},
}
