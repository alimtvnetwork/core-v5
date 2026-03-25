package chmodhelpertests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var attributeTestCases = []coretestcases.CaseV1{
	{
		Title: "Attribute rwx all true -- ToByte returns 7, ToRwxString returns rwx",
		ArrangeInput: args.Map{
			"when":    "all permissions enabled",
			"read":    true,
			"write":   true,
			"execute": true,
		},
		ExpectedInput: args.Map{
			"toByte":      7,
			"toRwxString": "rwx",
			"isEmpty":     false,
			"isDefined":   true,
			"isZero":      false,
			"isInvalid":   false,
		},
	},
	{
		Title: "Attribute all false -- ToByte returns 0, ToRwxString returns ---",
		ArrangeInput: args.Map{
			"when":    "no permissions",
			"read":    false,
			"write":   false,
			"execute": false,
		},
		ExpectedInput: args.Map{
			"toByte":      0,
			"toRwxString": "---",
			"isEmpty":     true,
			"isDefined":   false,
			"isZero":      true,
			"isInvalid":   true,
		},
	},
	{
		Title: "Attribute read only -- ToByte returns 4, ToRwxString returns r--",
		ArrangeInput: args.Map{
			"when":    "read only",
			"read":    true,
			"write":   false,
			"execute": false,
		},
		ExpectedInput: args.Map{
			"toByte":      4,
			"toRwxString": "r--",
			"isEmpty":     false,
			"isDefined":   true,
			"isZero":      false,
			"isInvalid":   false,
		},
	},
	{
		Title: "Attribute write+execute -- ToByte returns 3, ToRwxString returns -wx",
		ArrangeInput: args.Map{
			"when":    "write and execute",
			"read":    false,
			"write":   true,
			"execute": true,
		},
		ExpectedInput: args.Map{
			"toByte":      3,
			"toRwxString": "-wx",
			"isEmpty":     false,
			"isDefined":   true,
			"isZero":      false,
			"isInvalid":   false,
		},
	},
	{
		Title: "Attribute read+write -- ToByte returns 6, ToRwxString returns rw-",
		ArrangeInput: args.Map{
			"when":    "read and write",
			"read":    true,
			"write":   true,
			"execute": false,
		},
		ExpectedInput: args.Map{
			"toByte":      6,
			"toRwxString": "rw-",
			"isEmpty":     false,
			"isDefined":   true,
			"isZero":      false,
			"isInvalid":   false,
		},
	},
	{
		Title: "Attribute read+execute -- ToByte returns 5, ToRwxString returns r-x",
		ArrangeInput: args.Map{
			"when":    "read and execute",
			"read":    true,
			"write":   false,
			"execute": true,
		},
		ExpectedInput: args.Map{
			"toByte":      5,
			"toRwxString": "r-x",
			"isEmpty":     false,
			"isDefined":   true,
			"isZero":      false,
			"isInvalid":   false,
		},
	},
}

var attributeEqualTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEqual same attributes -- returns true",
		ArrangeInput: args.Map{
			"when":     "same rwx attributes",
			"leftRwx":  "rwx",
			"rightRwx": "rwx",
		},
		ExpectedInput: args.Map{
			"isEqual": true,
		},
	},
	{
		Title: "IsEqual different attributes -- returns false",
		ArrangeInput: args.Map{
			"when":     "different attributes",
			"leftRwx":  "rwx",
			"rightRwx": "r--",
		},
		ExpectedInput: args.Map{
			"isEqual": false,
		},
	},
}

var attributeCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "Clone creates independent copy -- values match",
		ArrangeInput: args.Map{
			"when":    "clone rwx attribute",
			"read":    true,
			"write":   false,
			"execute": true,
		},
		ExpectedInput: args.Map{
			"cloneRead":    true,
			"cloneWrite":   false,
			"cloneExecute": true,
			"isEqual":      true,
		},
	},
}

var usingByteTestCases = []coretestcases.CaseV1{
	{
		Title: "UsingByte 0 -- all false",
		ArrangeInput: args.Map{
			"when":  "byte value 0",
			"input": 0,
		},
		ExpectedInput: args.Map{
			"read":    false,
			"write":   false,
			"execute": false,
			"toByte":  0,
		},
	},
	{
		Title: "UsingByte 1 -- execute only",
		ArrangeInput: args.Map{
			"when":  "byte value 1",
			"input": 1,
		},
		ExpectedInput: args.Map{
			"read":    false,
			"write":   false,
			"execute": true,
			"toByte":  1,
		},
	},
	{
		Title: "UsingByte 2 -- write only",
		ArrangeInput: args.Map{
			"when":  "byte value 2",
			"input": 2,
		},
		ExpectedInput: args.Map{
			"read":    false,
			"write":   true,
			"execute": false,
			"toByte":  2,
		},
	},
	{
		Title: "UsingByte 3 -- write+execute",
		ArrangeInput: args.Map{
			"when":  "byte value 3",
			"input": 3,
		},
		ExpectedInput: args.Map{
			"read":    false,
			"write":   true,
			"execute": true,
			"toByte":  3,
		},
	},
	{
		Title: "UsingByte 4 -- read only",
		ArrangeInput: args.Map{
			"when":  "byte value 4",
			"input": 4,
		},
		ExpectedInput: args.Map{
			"read":    true,
			"write":   false,
			"execute": false,
			"toByte":  4,
		},
	},
	{
		Title: "UsingByte 5 -- read+execute",
		ArrangeInput: args.Map{
			"when":  "byte value 5",
			"input": 5,
		},
		ExpectedInput: args.Map{
			"read":    true,
			"write":   false,
			"execute": true,
			"toByte":  5,
		},
	},
	{
		Title: "UsingByte 6 -- read+write",
		ArrangeInput: args.Map{
			"when":  "byte value 6",
			"input": 6,
		},
		ExpectedInput: args.Map{
			"read":    true,
			"write":   true,
			"execute": false,
			"toByte":  6,
		},
	},
	{
		Title: "UsingByte 7 -- all true",
		ArrangeInput: args.Map{
			"when":  "byte value 7",
			"input": 7,
		},
		ExpectedInput: args.Map{
			"read":    true,
			"write":   true,
			"execute": true,
			"toByte":  7,
		},
	},
}

var usingRwxStringTestCases = []coretestcases.CaseV1{
	{
		Title: "UsingRwxString 'rwx' -- all true",
		ArrangeInput: args.Map{
			"when":  "rwx string",
			"input": "rwx",
		},
		ExpectedInput: args.Map{
			"read":    true,
			"write":   true,
			"execute": true,
		},
	},
	{
		Title: "UsingRwxString '---' -- all false",
		ArrangeInput: args.Map{
			"when":  "no permissions string",
			"input": "---",
		},
		ExpectedInput: args.Map{
			"read":    false,
			"write":   false,
			"execute": false,
		},
	},
	{
		Title: "UsingRwxString 'r-x' -- read+execute",
		ArrangeInput: args.Map{
			"when":  "read execute string",
			"input": "r-x",
		},
		ExpectedInput: args.Map{
			"read":    true,
			"write":   false,
			"execute": true,
		},
	},
	{
		Title: "UsingRwxString 'rw-' -- read+write",
		ArrangeInput: args.Map{
			"when":  "read write string",
			"input": "rw-",
		},
		ExpectedInput: args.Map{
			"read":    true,
			"write":   true,
			"execute": false,
		},
	},
}

var expandCharRwxTestCases = []coretestcases.CaseV1{
	{
		Title: "ExpandCharRwx '755' -- returns '7','5','5'",
		ArrangeInput: args.Map{
			"when":  "755 variant",
			"input": "755",
		},
		ExpectedInput: args.Map{
			"r": "55",
			"w": "53",
			"x": "53",
		},
	},
	{
		Title: "ExpandCharRwx '644' -- returns '6','4','4'",
		ArrangeInput: args.Map{
			"when":  "644 variant",
			"input": "644",
		},
		ExpectedInput: args.Map{
			"r": "54",
			"w": "52",
			"x": "52",
		},
	},
}

var isPathExistsTestCases = []coretestcases.CaseV1{
	{
		Title: "IsPathExists existing path -- returns true",
		ArrangeInput: args.Map{
			"when":  "current directory",
			"input": ".",
		},
		ExpectedInput: args.Map{
			"exists":  true,
			"invalid": false,
		},
	},
	{
		Title: "IsPathExists non-existing path -- returns false",
		ArrangeInput: args.Map{
			"when":  "non-existing path",
			"input": "/absolutely/non/existing/path/xyz123",
		},
		ExpectedInput: args.Map{
			"exists":  false,
			"invalid": true,
		},
	},
}

var isDirectoryTestCases = []coretestcases.CaseV1{
	{
		Title: "IsDirectory on current dir -- returns true",
		ArrangeInput: args.Map{
			"when":  "current directory",
			"input": ".",
		},
		ExpectedInput: args.Map{
			"isDir": true,
		},
	},
	{
		Title: "IsDirectory on non-existing -- returns false",
		ArrangeInput: args.Map{
			"when":  "non-existing path",
			"input": "/non/existing/dir/xyz123",
		},
		ExpectedInput: args.Map{
			"isDir": false,
		},
	},
}

var getRwxLengthErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "GetRwxLengthError valid length 3 -- returns nil",
		ArrangeInput: args.Map{
			"when":  "valid rwx string",
			"input": "rwx",
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
	{
		Title: "GetRwxLengthError invalid length -- returns error",
		ArrangeInput: args.Map{
			"when":  "invalid length string",
			"input": "rw",
		},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
	{
		Title: "GetRwxLengthError empty string -- returns error",
		ArrangeInput: args.Map{
			"when":  "empty string",
			"input": "",
		},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
}
