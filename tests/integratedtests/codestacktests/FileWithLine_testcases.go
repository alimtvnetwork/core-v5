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
		ExpectedInput: args.Map{
			"filePath":   "/src/main.go",
			"lineNumber": "42",
			"isValid":    "true",
		},
	},
	{
		Title: "FileWithLine with empty path",
		ArrangeInput: args.Map{
			"when": "given empty file path",
			"file": "",
			"line": 0,
		},
		ExpectedInput: args.Map{
			"filePath":   "",
			"lineNumber": "0",
			"isValid":    "true",
		},
	},
}
