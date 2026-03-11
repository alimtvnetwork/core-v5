package ostypetests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/ostype"
)

var getVariantTestCases = []coretestcases.CaseV1{
	{
		Title:         "GetVariant returns 'windows' -- 'windows' input",
		ArrangeInput:  "windows",
		ExpectedInput: "windows",
	},
	{
		Title:         "GetVariant returns 'linux' -- 'linux' input",
		ArrangeInput:  "linux",
		ExpectedInput: "linux",
	},
	{
		Title:         "GetVariant returns 'darwin' -- 'darwin' input",
		ArrangeInput:  "darwin",
		ExpectedInput: "darwin",
	},
	{
		Title:         "GetVariant returns 'freebsd' -- 'freebsd' input",
		ArrangeInput:  "freebsd",
		ExpectedInput: "freebsd",
	},
	{
		Title:         "GetVariant returns 'android' -- 'android' input",
		ArrangeInput:  "android",
		ExpectedInput: "android",
	},
	{
		Title:         "GetVariant returns 'Unknown' -- unrecognized OS string",
		ArrangeInput:  "totally_unknown_os",
		ExpectedInput: "Unknown",
	},
}

var getGroupTestCases = []coretestcases.CaseV1{
	{
		Title:         "GetGroup returns WindowsGroup -- 'windows' input",
		ArrangeInput:  "windows",
		ExpectedInput: "WindowsGroup",
	},
	{
		Title:         "GetGroup returns UnixGroup -- 'linux' input",
		ArrangeInput:  "linux",
		ExpectedInput: "UnixGroup",
	},
	{
		Title:         "GetGroup returns UnixGroup -- 'darwin' input",
		ArrangeInput:  "darwin",
		ExpectedInput: "UnixGroup",
	},
	{
		Title:         "GetGroup returns AndroidGroup -- 'android' input",
		ArrangeInput:  "android",
		ExpectedInput: "AndroidGroup",
	},
	{
		Title:         "GetGroup returns JavaScriptGroup -- unrecognized OS string",
		ArrangeInput:  "totally_unknown_os",
		ExpectedInput: "JavaScriptGroup",
	},
}

// variationGroupTestCases
// Expected: groupName, isUnix, isWindows
var variationGroupTestCases = []coretestcases.CaseV1{
	{
		Title:        "Variation.Group returns WindowsGroup and isWindows true -- Windows variant",
		ArrangeInput: ostype.Windows,
		ExpectedInput: args.Map{
			"groupName": "WindowsGroup",
			"isUnix":    "false",
			"isWindows": "true",
		},
	},
	{
		Title:        "Variation.Group returns UnixGroup and isUnix true -- Linux variant",
		ArrangeInput: ostype.Linux,
		ExpectedInput: args.Map{
			"groupName": "UnixGroup",
			"isUnix":    "true",
			"isWindows": "false",
		},
	},
	{
		Title:        "Variation.Group returns UnixGroup and isUnix true -- DarwinOrMacOs variant",
		ArrangeInput: ostype.DarwinOrMacOs,
		ExpectedInput: args.Map{
			"groupName": "UnixGroup",
			"isUnix":    "true",
			"isWindows": "false",
		},
	},
	{
		Title:        "Variation.Group returns AndroidGroup -- Android variant",
		ArrangeInput: ostype.Android,
		ExpectedInput: args.Map{
			"groupName": "AndroidGroup",
			"isUnix":    "false",
			"isWindows": "false",
		},
	},
}

// variationIdentityTestCases
// Expected: isWindows, isLinux, isDarwin, isValid
var variationIdentityTestCases = []coretestcases.CaseV1{
	{
		Title:        "Variation returns isWindows true and isValid true -- Windows variant",
		ArrangeInput: ostype.Windows,
		ExpectedInput: args.Map{
			"isWindows": "true",
			"isLinux":   "false",
			"isDarwin":  "false",
			"isValid":   "true",
		},
	},
	{
		Title:        "Variation returns isLinux true and isValid true -- Linux variant",
		ArrangeInput: ostype.Linux,
		ExpectedInput: args.Map{
			"isWindows": "false",
			"isLinux":   "true",
			"isDarwin":  "false",
			"isValid":   "true",
		},
	},
	{
		Title:        "Variation returns isDarwin true and isValid true -- DarwinOrMacOs variant",
		ArrangeInput: ostype.DarwinOrMacOs,
		ExpectedInput: args.Map{
			"isWindows": "false",
			"isLinux":   "false",
			"isDarwin":  "true",
			"isValid":   "true",
		},
	},
	{
		Title:        "Variation returns isValid false -- Any (default) variant",
		ArrangeInput: ostype.Any,
		ExpectedInput: args.Map{
			"isWindows": "false",
			"isLinux":   "false",
			"isDarwin":  "false",
			"isValid":   "false",
		},
	},
}