package chmodhelpertests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var rwxWrapperCreateTestCases = []coretestcases.CaseV1{
	{
		Title: "RwxWrapper Create 755 -- returns correct owner/group/other",
		ArrangeInput: args.Map{
			"when":  "standard 755",
			"input": "755",
		},
		ExpectedInput: args.Map{
			"ownerRwx":    "rwx",
			"groupRwx":    "r-x",
			"otherRwx":    "r-x",
			"fullRwx":     "-rwxr-xr-x",
			"rwx9":        "rwxr-xr-x",
			"fileMode":    "0755",
			"rwxCompiled": "755",
			"isEmpty":     false,
			"isDefined":   true,
		},
	},
	{
		Title: "RwxWrapper Create 644 -- returns correct owner/group/other",
		ArrangeInput: args.Map{
			"when":  "standard 644",
			"input": "644",
		},
		ExpectedInput: args.Map{
			"ownerRwx":    "rw-",
			"groupRwx":    "r--",
			"otherRwx":    "r--",
			"fullRwx":     "-rw-r--r--",
			"rwx9":        "rw-r--r--",
			"fileMode":    "0644",
			"rwxCompiled": "644",
			"isEmpty":     false,
			"isDefined":   true,
		},
	},
	{
		Title: "RwxWrapper Create 777 -- all permissions",
		ArrangeInput: args.Map{
			"when":  "all permissions 777",
			"input": "777",
		},
		ExpectedInput: args.Map{
			"ownerRwx":    "rwx",
			"groupRwx":    "rwx",
			"otherRwx":    "rwx",
			"fullRwx":     "-rwxrwxrwx",
			"rwx9":        "rwxrwxrwx",
			"fileMode":    "0777",
			"rwxCompiled": "777",
			"isEmpty":     false,
			"isDefined":   true,
		},
	},
	{
		Title: "RwxWrapper Create 000 -- no permissions",
		ArrangeInput: args.Map{
			"when":  "no permissions 000",
			"input": "000",
		},
		ExpectedInput: args.Map{
			"ownerRwx":    "---",
			"groupRwx":    "---",
			"otherRwx":    "---",
			"fullRwx":     "----------",
			"rwx9":        "---------",
			"fileMode":    "0000",
			"rwxCompiled": "000",
			"isEmpty":     true,
			"isDefined":   false,
		},
	},
	{
		Title: "RwxWrapper Create 400 -- owner read only",
		ArrangeInput: args.Map{
			"when":  "owner read only 400",
			"input": "400",
		},
		ExpectedInput: args.Map{
			"ownerRwx":    "r--",
			"groupRwx":    "---",
			"otherRwx":    "---",
			"fullRwx":     "-r--------",
			"rwx9":        "r--------",
			"fileMode":    "0400",
			"rwxCompiled": "400",
			"isEmpty":     false,
			"isDefined":   true,
		},
	},
}

var rwxWrapperRwxFullStringTestCases = []coretestcases.CaseV1{
	{
		Title: "RwxFullString valid 10 char -- parses correctly",
		ArrangeInput: args.Map{
			"when":  "valid hyphenated rwx",
			"input": "-rwxr-xr-x",
		},
		ExpectedInput: args.Map{
			"ownerRwx":  "rwx",
			"groupRwx":  "r-x",
			"otherRwx":  "r-x",
			"hasError":  false,
			"isDefined": true,
		},
	},
	{
		Title: "RwxFullString invalid length -- returns error",
		ArrangeInput: args.Map{
			"when":  "invalid short string",
			"input": "rwxr-x",
		},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
}

var rwxWrapperBytesTestCases = []coretestcases.CaseV1{
	{
		Title: "RwxWrapper Bytes 755 -- returns [7,5,5]",
		ArrangeInput: args.Map{
			"when":  "755 wrapper",
			"input": "755",
		},
		ExpectedInput: args.Map{
			"byte0": 7,
			"byte1": 5,
			"byte2": 5,
		},
	},
	{
		Title: "RwxWrapper Bytes 644 -- returns [6,4,4]",
		ArrangeInput: args.Map{
			"when":  "644 wrapper",
			"input": "644",
		},
		ExpectedInput: args.Map{
			"byte0": 6,
			"byte1": 4,
			"byte2": 4,
		},
	},
}

var rwxWrapperCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "RwxWrapper Clone -- creates equal independent copy",
		ArrangeInput: args.Map{
			"when":  "clone 755",
			"input": "755",
		},
		ExpectedInput: args.Map{
			"isEqual":    true,
			"ownerRwx":   "rwx",
			"groupRwx":   "r-x",
			"otherRwx":   "r-x",
			"clonedNull": false,
		},
	},
}

var rwxWrapperIsEqualTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEqualPtr same wrappers -- returns true",
		ArrangeInput: args.Map{
			"when":  "same 755",
			"left":  "755",
			"right": "755",
		},
		ExpectedInput: args.Map{
			"isEqual": true,
		},
	},
	{
		Title: "IsEqualPtr different wrappers -- returns false",
		ArrangeInput: args.Map{
			"when":  "755 vs 644",
			"left":  "755",
			"right": "644",
		},
		ExpectedInput: args.Map{
			"isEqual": false,
		},
	},
}

var rwxWrapperVariantTestCases = []coretestcases.CaseV1{
	{
		Title: "RwxWrapper from Variant 755 -- correct wrapper",
		ArrangeInput: args.Map{
			"when":  "variant 755",
			"input": "755",
		},
		ExpectedInput: args.Map{
			"fullRwx":  "-rwxr-xr-x",
			"hasError": false,
		},
	},
	{
		Title: "RwxWrapper from Variant 644 -- correct wrapper",
		ArrangeInput: args.Map{
			"when":  "variant 644",
			"input": "644",
		},
		ExpectedInput: args.Map{
			"fullRwx":  "-rw-r--r--",
			"hasError": false,
		},
	},
}

var attrVariantTestCases = []coretestcases.CaseV1{
	{
		Title: "AttrVariant Execute -- value is 1",
		ArrangeInput: args.Map{
			"when":  "Execute variant",
			"input": 1,
		},
		ExpectedInput: args.Map{
			"value":       1,
			"attrRead":    false,
			"attrWrite":   false,
			"attrExecute": true,
		},
	},
	{
		Title: "AttrVariant ReadWriteExecute -- value is 7",
		ArrangeInput: args.Map{
			"when":  "ReadWriteExecute variant",
			"input": 7,
		},
		ExpectedInput: args.Map{
			"value":       7,
			"attrRead":    true,
			"attrWrite":   true,
			"attrExecute": true,
		},
	},
	{
		Title: "AttrVariant Read -- value is 4",
		ArrangeInput: args.Map{
			"when":  "Read variant",
			"input": 4,
		},
		ExpectedInput: args.Map{
			"value":       4,
			"attrRead":    true,
			"attrWrite":   false,
			"attrExecute": false,
		},
	},
	{
		Title: "AttrVariant ReadWrite -- value is 6",
		ArrangeInput: args.Map{
			"when":  "ReadWrite variant",
			"input": 6,
		},
		ExpectedInput: args.Map{
			"value":       6,
			"attrRead":    true,
			"attrWrite":   true,
			"attrExecute": false,
		},
	},
}

var rwxWrapper9StringTestCases = []coretestcases.CaseV1{
	{
		Title: "Rwx9 valid 9 char -- parses correctly",
		ArrangeInput: args.Map{
			"when":  "valid 9 char rwx",
			"input": "rwxr-xr-x",
		},
		ExpectedInput: args.Map{
			"ownerRwx":  "rwx",
			"groupRwx":  "r-x",
			"otherRwx":  "r-x",
			"hasError":  false,
			"isDefined": true,
		},
	},
	{
		Title: "Rwx9 invalid length -- returns error",
		ArrangeInput: args.Map{
			"when":  "too short",
			"input": "rwx",
		},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
}

var rwxWrapperOctalTestCases = []coretestcases.CaseV1{
	{
		Title: "ToCompiledOctalBytes4Digits 755 -- returns [0,7,5,5] as string chars",
		ArrangeInput: args.Map{
			"when":  "755 octal",
			"input": "755",
		},
		ExpectedInput: args.Map{
			"octal4": "0755",
		},
	},
	{
		Title: "ToCompiledOctalBytes4Digits 644 -- returns [0,6,4,4] as string chars",
		ArrangeInput: args.Map{
			"when":  "644 octal",
			"input": "644",
		},
		ExpectedInput: args.Map{
			"octal4": "0644",
		},
	},
}
