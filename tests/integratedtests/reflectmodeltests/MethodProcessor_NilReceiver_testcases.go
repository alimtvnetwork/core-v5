package reflectmodeltests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/reflectcore/reflectmodel"
)

// =============================================================================
// MethodProcessor nil receiver test cases (migrated from inline t.Error tests)
// =============================================================================

var methodProcessorNilReceiverTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "HasValidFunc on nil returns false",
		Func:  (*reflectmodel.MethodProcessor).HasValidFunc,
		Expected: args.Map{
			"value":    "false",
			"panicked": false,
		},
	},
	{
		Title: "IsInvalid on nil returns true",
		Func:  (*reflectmodel.MethodProcessor).IsInvalid,
		Expected: args.Map{
			"value":    "true",
			"panicked": false,
		},
	},
	{
		Title: "Func on nil returns nil",
		Func:  (*reflectmodel.MethodProcessor).Func,
		Expected: args.Map{
			"value":    "<nil>",
			"panicked": false,
		},
	},
	{
		Title: "ReturnLength on nil returns -1",
		Func:  (*reflectmodel.MethodProcessor).ReturnLength,
		Expected: args.Map{
			"value":    "-1",
			"panicked": false,
		},
	},
	{
		Title: "IsPublicMethod on nil returns false",
		Func:  (*reflectmodel.MethodProcessor).IsPublicMethod,
		Expected: args.Map{
			"value":    "false",
			"panicked": false,
		},
	},
	{
		Title: "GetType on nil returns nil",
		Func:  (*reflectmodel.MethodProcessor).GetType,
		Expected: args.Map{
			"value":    "<nil>",
			"panicked": false,
		},
	},
	{
		Title: "GetInArgsTypes on nil returns empty",
		Func:  (*reflectmodel.MethodProcessor).GetInArgsTypes,
		Expected: args.Map{
			"panicked": false,
		},
	},
	{
		Title: "GetOutArgsTypes on nil returns empty",
		Func:  (*reflectmodel.MethodProcessor).GetOutArgsTypes,
		Expected: args.Map{
			"panicked": false,
		},
	},
	{
		Title: "GetInArgsTypesNames on nil returns empty",
		Func:  (*reflectmodel.MethodProcessor).GetInArgsTypesNames,
		Expected: args.Map{
			"panicked": false,
		},
	},
	{
		Title: "Invoke on nil returns error",
		Func:  (*reflectmodel.MethodProcessor).Invoke,
		Expected: args.Map{
			"panicked": false,
			"hasError": true,
		},
	},
}
