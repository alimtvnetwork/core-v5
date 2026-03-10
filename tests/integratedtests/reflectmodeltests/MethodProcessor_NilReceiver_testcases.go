package reflectmodeltests

import (
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/coretests/results"
	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
)

// =============================================================================
// MethodProcessor nil receiver test cases (migrated from inline t.Error tests)
// =============================================================================

var methodProcessorNilReceiverTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "HasValidFunc on nil returns false",
		Func:  (*reflectmodel.MethodProcessor).HasValidFunc,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "IsInvalid on nil returns true",
		Func:  (*reflectmodel.MethodProcessor).IsInvalid,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "Func on nil returns nil",
		Func:  (*reflectmodel.MethodProcessor).Func,
		Expected: results.ResultAny{
			Value:    "<nil>",
			Panicked: false,
		},
	},
	{
		Title: "ReturnLength on nil returns -1",
		Func:  (*reflectmodel.MethodProcessor).ReturnLength,
		Expected: results.ResultAny{
			Value:    "-1",
			Panicked: false,
		},
	},
	{
		Title: "IsPublicMethod on nil returns false",
		Func:  (*reflectmodel.MethodProcessor).IsPublicMethod,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "GetType on nil returns nil",
		Func:  (*reflectmodel.MethodProcessor).GetType,
		Expected: results.ResultAny{
			Value:    "<nil>",
			Panicked: false,
		},
	},
	{
		Title: "GetInArgsTypes on nil returns empty",
		Func:  (*reflectmodel.MethodProcessor).GetInArgsTypes,
		Expected: results.ResultAny{
			Panicked: false,
		},
	},
	{
		Title: "GetOutArgsTypes on nil returns empty",
		Func:  (*reflectmodel.MethodProcessor).GetOutArgsTypes,
		Expected: results.ResultAny{
			Panicked: false,
		},
	},
	{
		Title: "GetInArgsTypesNames on nil returns empty",
		Func:  (*reflectmodel.MethodProcessor).GetInArgsTypesNames,
		Expected: results.ResultAny{
			Panicked: false,
		},
	},
	{
		Title: "Invoke on nil returns error",
		Func:  (*reflectmodel.MethodProcessor).Invoke,
		Expected: results.ResultAny{
			Panicked: false,
			Error:    results.ExpectAnyError,
		},
	},
}
