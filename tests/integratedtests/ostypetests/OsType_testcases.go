package ostypetests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/ostype"
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
		ExpectedInput: "unknown",
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
// Expected order: groupName, isUnix, isWindows
var variationGroupTestCases = []coretestcases.CaseV1{
	{
		Title:        "Windows variation has WindowsGroup",
		ArrangeInput: ostype.Windows,
		ExpectedInput: args.Three[string, string, string]{
			First:  "Windows", // groupName
			Second: "false",   // isUnix
			Third:  "true",    // isWindows
		},
	},
	{
		Title:        "Linux variation has UnixGroup",
		ArrangeInput: ostype.Linux,
		ExpectedInput: args.Three[string, string, string]{
			First:  "Unix", // groupName
			Second: "true", // isUnix
			Third:  "false", // isWindows
		},
	},
	{
		Title:        "DarwinOrMacOs variation has UnixGroup",
		ArrangeInput: ostype.DarwinOrMacOs,
		ExpectedInput: args.Three[string, string, string]{
			First:  "Unix", // groupName
			Second: "true", // isUnix
			Third:  "false", // isWindows
		},
	},
	{
		Title:        "Android variation has AndroidGroup",
		ArrangeInput: ostype.Android,
		ExpectedInput: args.Three[string, string, string]{
			First:  "Android", // groupName
			Second: "false",   // isUnix
			Third:  "false",   // isWindows
		},
	},
}

// variationIdentityTestCases
// Expected order: isWindows, isLinux, isDarwin, isValid
var variationIdentityTestCases = []coretestcases.CaseV1{
	{
		Title:        "Windows identity checks",
		ArrangeInput: ostype.Windows,
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "true",  // isWindows
			Second: "false", // isLinux
			Third:  "false", // isDarwin
			Fourth: "true",  // isValid
		},
	},
	{
		Title:        "Linux identity checks",
		ArrangeInput: ostype.Linux,
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "false", // isWindows
			Second: "true",  // isLinux
			Third:  "false", // isDarwin
			Fourth: "true",  // isValid
		},
	},
	{
		Title:        "DarwinOrMacOs identity checks",
		ArrangeInput: ostype.DarwinOrMacOs,
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "false", // isWindows
			Second: "false", // isLinux
			Third:  "true",  // isDarwin
			Fourth: "true",  // isValid
		},
	},
	{
		Title:        "Any (default) is invalid",
		ArrangeInput: ostype.Any,
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "false", // isWindows
			Second: "false", // isLinux
			Third:  "false", // isDarwin
			Fourth: "false", // isValid
		},
	},
}
