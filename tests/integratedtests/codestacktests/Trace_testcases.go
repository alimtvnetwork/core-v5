package codestacktests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var traceTestCases = []coretestcases.CaseV1{
	{
		Title: "Trace basic properties -- valid trace returns correct fields",
		ArrangeInput: args.Map{
			"when":        "valid trace created",
			"packageName": "mypackage",
			"methodName":  "MyMethod",
			"pkgMethod":   "mypackage.MyMethod",
			"filePath":    "/src/mypackage/file.go",
			"line":        42,
		},
		ExpectedInput: args.Map{
			"packageName": "mypackage",
			"methodName":  "MyMethod",
			"pkgMethod":   "mypackage.MyMethod",
			"filePath":    "/src/mypackage/file.go",
			"lineNumber":  42,
			"isNil":       false,
			"isNotNil":    true,
			"hasIssues":   false,
		},
	},
}

var traceNilTestCases = []coretestcases.CaseV1{
	{
		Title: "Trace nil -- nil trace returns correct defaults",
		ArrangeInput: args.Map{
			"when": "nil trace",
		},
		ExpectedInput: args.Map{
			"isNil":    true,
			"isNotNil": false,
			"string":   "",
		},
	},
}

var traceDisposeTestCases = []coretestcases.CaseV1{
	{
		Title: "Trace Dispose -- clears all fields",
		ArrangeInput: args.Map{
			"when":        "dispose called on valid trace",
			"packageName": "pkg",
			"methodName":  "Method",
			"pkgMethod":   "pkg.Method",
			"filePath":    "/file.go",
			"line":        10,
		},
		ExpectedInput: args.Map{
			"packageName": "",
			"methodName":  "",
			"pkgMethod":   "",
			"filePath":    "",
			"lineNumber":  0,
			"isOkay":      false,
		},
	},
}

var traceCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "Trace Clone -- creates independent copy",
		ArrangeInput: args.Map{
			"when":        "clone valid trace",
			"packageName": "clonepkg",
			"methodName":  "CloneMethod",
			"pkgMethod":   "clonepkg.CloneMethod",
			"filePath":    "/clone/file.go",
			"line":        99,
		},
		ExpectedInput: args.Map{
			"packageName": "clonepkg",
			"methodName":  "CloneMethod",
			"filePath":    "/clone/file.go",
			"lineNumber":  99,
		},
	},
}

var fileWithLineStringMethodTestCases = []coretestcases.CaseV1{
	{
		Title: "FileWithLine String method -- returns formatted path:line",
		ArrangeInput: args.Map{
			"when": "file with line created",
			"file": "/src/app.go",
			"line": 100,
		},
		ExpectedInput: args.Map{
			"isNil":    false,
			"isNotNil": true,
			"hasLine":  true,
		},
	},
}
