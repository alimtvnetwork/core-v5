package codestacktests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var fileWithLineTestCases = []coretestcases.CaseV1{
	{
		Title: "FileWithLine returns correct path and lineNumber -- '/src/main.go' line 42",
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
