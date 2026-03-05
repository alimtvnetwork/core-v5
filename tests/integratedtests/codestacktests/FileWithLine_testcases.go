package codestacktests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var fileWithLineTestCases = []coretestcases.CaseV1{
	{
		Title: "FileWithLine stores path and line",
		ArrangeInput: args.Map{
			"when": "given file path and line number",
			"file": "/src/main.go",
			"line": 42,
		},
		ExpectedInput: args.Three[string, string, string]{
			First:  "/src/main.go", // filePath
			Second: "42",           // lineNumber
			Third:  "true",         // isValid
		},
	},
	{
		Title: "FileWithLine with empty path",
		ArrangeInput: args.Map{
			"when": "given empty file path",
			"file": "",
			"line": 0,
		},
		ExpectedInput: args.Three[string, string, string]{
			First:  "",     // filePath
			Second: "0",    // lineNumber
			Third:  "true", // isValid
		},
	},
}
