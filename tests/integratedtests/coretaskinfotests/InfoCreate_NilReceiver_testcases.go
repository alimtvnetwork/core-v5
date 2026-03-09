package coretaskinfotests

import (
	"gitlab.com/auk-go/core/coretaskinfo"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/coretests/results"
)

// =============================================================================
// Info nil receiver test cases
// (migrated from CaseV1 variables in InfoCreate_testcases.go)
// =============================================================================

var infoNilSafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "SafeName on nil returns empty",
		Func:  (*coretaskinfo.Info).SafeName,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "SafeDescription on nil returns empty",
		Func:  (*coretaskinfo.Info).SafeDescription,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "SafeUrl on nil returns empty",
		Func:  (*coretaskinfo.Info).SafeUrl,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "SafeHintUrl on nil returns empty",
		Func:  (*coretaskinfo.Info).SafeHintUrl,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "SafeErrorUrl on nil returns empty",
		Func:  (*coretaskinfo.Info).SafeErrorUrl,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "SafeExampleUrl on nil returns empty",
		Func:  (*coretaskinfo.Info).SafeExampleUrl,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "IsNull on nil returns true",
		Func:  (*coretaskinfo.Info).IsNull,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "IsDefined on nil returns false",
		Func:  (*coretaskinfo.Info).IsDefined,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "IsEmpty on nil returns true",
		Func:  (*coretaskinfo.Info).IsEmpty,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "HasAnyItem on nil returns false",
		Func:  (*coretaskinfo.Info).HasAnyItem,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "ClonePtr on nil returns nil",
		Func: func(info *coretaskinfo.Info) bool {
			return info.ClonePtr() == nil
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "PrettyJsonString on nil returns empty",
		Func:  (*coretaskinfo.Info).PrettyJsonString,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
}
