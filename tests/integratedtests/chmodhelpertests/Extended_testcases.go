package chmodhelpertests

import (
	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/coretests/results"
)

type pathExistStatType = chmodhelper.PathExistStat
type nilSafeResult = results.ResultAny

// ── PathExistStat nil-safety tests ──

var extPathExistStatNilSafeCases = []coretestcases.CaseNilSafe{
	{
		Title: "HasError on nil receiver returns false",
		Func:  (*pathExistStatType).HasError,
		Expected: nilSafeResult{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "IsEmptyError on nil receiver returns true",
		Func:  (*pathExistStatType).IsEmptyError,
		Expected: nilSafeResult{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "HasFileInfo on nil receiver returns false",
		Func:  (*pathExistStatType).HasFileInfo,
		Expected: nilSafeResult{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "IsInvalidFileInfo on nil receiver returns true",
		Func:  (*pathExistStatType).IsInvalidFileInfo,
		Expected: nilSafeResult{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "IsFile on nil receiver returns false",
		Func:  (*pathExistStatType).IsFile,
		Expected: nilSafeResult{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "IsDir on nil receiver returns false",
		Func:  (*pathExistStatType).IsDir,
		Expected: nilSafeResult{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "IsInvalid on nil receiver returns true",
		Func:  (*pathExistStatType).IsInvalid,
		Expected: nilSafeResult{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "HasAnyIssues on nil receiver returns true",
		Func:  (*pathExistStatType).HasAnyIssues,
		Expected: nilSafeResult{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "Dispose on nil receiver does not panic",
		Func:  (*pathExistStatType).Dispose,
		Expected: nilSafeResult{
			Panicked: false,
		},
	},
	{
		Title: "String on nil receiver returns empty",
		Func:  (*pathExistStatType).String,
		Expected: nilSafeResult{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "NotExistError on nil receiver returns nil",
		Func:  (*pathExistStatType).NotExistError,
		Expected: nilSafeResult{
			Panicked: false,
		},
	},
	{
		Title: "NotAFileError on nil receiver returns nil",
		Func:  (*pathExistStatType).NotAFileError,
		Expected: nilSafeResult{
			Panicked: false,
		},
	},
	{
		Title: "NotADirError on nil receiver returns nil",
		Func:  (*pathExistStatType).NotADirError,
		Expected: nilSafeResult{
			Panicked: false,
		},
	},
	{
		Title: "MeaningFullError on nil receiver returns nil",
		Func:  (*pathExistStatType).MeaningFullError,
		Expected: nilSafeResult{
			Panicked: false,
		},
	},
	{
		Title: "LastModifiedDate on nil receiver returns nil",
		Func:  (*pathExistStatType).LastModifiedDate,
		Expected: nilSafeResult{
			Panicked: false,
		},
	},
	{
		Title: "FileMode on nil receiver returns nil",
		Func:  (*pathExistStatType).FileMode,
		Expected: nilSafeResult{
			Panicked: false,
		},
	},
	{
		Title: "Size on nil receiver returns nil",
		Func:  (*pathExistStatType).Size,
		Expected: nilSafeResult{
			Panicked: false,
		},
	},
}

// ── PathExistStat method tests ──

var extPathExistStatTestCases = []coretestcases.CaseV1{
	{
		Title: "GetPathExistStat on invalid path -- returns not exist",
		ArrangeInput: args.Map{
			"when": "invalid path",
			"path": "/nonexistent/path/that/does/not/exist/at/all",
		},
		ExpectedInput: args.Map{
			"isExist":  false,
			"isInvalid": true,
		},
	},
}

// ── chmodVerifier test cases ──

var extChmodVerifierTestCases = []coretestcases.CaseV1{
	{
		Title: "GetRwx9 returns 9 char string for standard filemode",
		ArrangeInput: args.Map{
			"when": "standard 0755 filemode",
		},
		ExpectedInput: args.Map{
			"rwx9Length": 9,
		},
	},
	{
		Title: "GetRwxFull returns 10 char string for standard filemode",
		ArrangeInput: args.Map{
			"when": "standard 0755 filemode",
		},
		ExpectedInput: args.Map{
			"rwxFullLength": 10,
		},
	},
}
