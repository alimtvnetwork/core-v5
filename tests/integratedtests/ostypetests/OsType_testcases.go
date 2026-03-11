package ostypetests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/ostype"
)

var getVariantTestCases = []coretestcases.CaseV1{
	{
		Title:         "windows maps to Windows",
		ArrangeInput:  "windows",
		ExpectedInput: "windows",
	},
	{
		Title:         "linux maps to Linux",
		ArrangeInput:  "linux",
		ExpectedInput: "linux",
	},
	{
		Title:         "darwin maps to DarwinOrMacOs",
		ArrangeInput:  "darwin",
		ExpectedInput: "darwin",
	},
	{
		Title:         "freebsd maps to FreeBsd",
		ArrangeInput:  "freebsd",
		ExpectedInput: "freebsd",
	},
	{
		Title:         "android maps to Android",
		ArrangeInput:  "android",
		ExpectedInput: "android",
	},
	{
		Title:         "unknown string maps to Unknown",
		ArrangeInput:  "totally_unknown_os",
		ExpectedInput: "Unknown",
	},
}

var getGroupTestCases = []coretestcases.CaseV1{
	{
		Title:         "windows is WindowsGroup",
		ArrangeInput:  "windows",
		ExpectedInput: "Windows",
	},
	{
		Title:         "linux is UnixGroup",
		ArrangeInput:  "linux",
		ExpectedInput: "Unix",
	},
	{
		Title:         "darwin is UnixGroup",
		ArrangeInput:  "darwin",
		ExpectedInput: "Unix",
	},
	{
		Title:         "android is AndroidGroup",
		ArrangeInput:  "android",
		ExpectedInput: "Android",
	},
	{
		Title:         "unknown returns InvalidGroup",
		ArrangeInput:  "totally_unknown_os",
		ExpectedInput: "Invalid",
	},
}

// variationGroupTestCases
// Expected: groupName, isUnix, isWindows
var variationGroupTestCases = []coretestcases.CaseV1{
	{
		Title:        "Windows variation has WindowsGroup",
		ArrangeInput: ostype.Windows,
		ExpectedInput: args.Map{
			"groupName": "Windows",
			"isUnix":    "false",
			"isWindows": "true",
		},
	},
	{
		Title:        "Linux variation has UnixGroup",
		ArrangeInput: ostype.Linux,
		ExpectedInput: args.Map{
			"groupName": "Unix",
			"isUnix":    "true",
			"isWindows": "false",
		},
	},
	{
		Title:        "DarwinOrMacOs variation has UnixGroup",
		ArrangeInput: ostype.DarwinOrMacOs,
		ExpectedInput: args.Map{
			"groupName": "Unix",
			"isUnix":    "true",
			"isWindows": "false",
		},
	},
	{
		Title:        "Android variation has AndroidGroup",
		ArrangeInput: ostype.Android,
		ExpectedInput: args.Map{
			"groupName": "Android",
			"isUnix":    "false",
			"isWindows": "false",
		},
	},
}

// variationIdentityTestCases
// Expected: isWindows, isLinux, isDarwin, isValid
var variationIdentityTestCases = []coretestcases.CaseV1{
	{
		Title:        "Windows identity checks",
		ArrangeInput: ostype.Windows,
		ExpectedInput: args.Map{
			"isWindows": "true",
			"isLinux":   "false",
			"isDarwin":  "false",
			"isValid":   "true",
		},
	},
	{
		Title:        "Linux identity checks",
		ArrangeInput: ostype.Linux,
		ExpectedInput: args.Map{
			"isWindows": "false",
			"isLinux":   "true",
			"isDarwin":  "false",
			"isValid":   "true",
		},
	},
	{
		Title:        "DarwinOrMacOs identity checks",
		ArrangeInput: ostype.DarwinOrMacOs,
		ExpectedInput: args.Map{
			"isWindows": "false",
			"isLinux":   "false",
			"isDarwin":  "true",
			"isValid":   "true",
		},
	},
	{
		Title:        "Any (default) is invalid",
		ArrangeInput: ostype.Any,
		ExpectedInput: args.Map{
			"isWindows": "false",
			"isLinux":   "false",
			"isDarwin":  "false",
			"isValid":   "false",
		},
	},
}
