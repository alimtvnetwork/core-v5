package stringcompareastests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var globMatchTestCases = []coretestcases.CaseV1{
	// === Positive: basic wildcard ===
	{
		Title: "Glob - star matches any filename",
		ArrangeInput: args.Map{
			"when":         "given glob with star wildcard",
			"pattern":      "*.go",
			"content":      "main.go",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Map{
			"isMatch":   "true",
			"isInverse": "false",
		},
	},
	{
		Title: "Glob - star matches path segment",
		ArrangeInput: args.Map{
			"when":         "given glob matching dynamic directory",
			"pattern":      "build-*/result.json",
			"content":      "build-20260303/result.json",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Map{
			"isMatch":   "true",
			"isInverse": "false",
		},
	},
	{
		Title: "Glob - question mark matches single char",
		ArrangeInput: args.Map{
			"when":         "given glob with question mark",
			"pattern":      "file?.txt",
			"content":      "fileA.txt",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Map{
			"isMatch":   "true",
			"isInverse": "false",
		},
	},
	{
		Title: "Glob - character class matches range",
		ArrangeInput: args.Map{
			"when":         "given glob with character class",
			"pattern":      "log[0-9].txt",
			"content":      "log5.txt",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Map{
			"isMatch":   "true",
			"isInverse": "false",
		},
	},
	{
		Title: "Glob - exact match without wildcards",
		ArrangeInput: args.Map{
			"when":         "given glob with no wildcards",
			"pattern":      "exact.txt",
			"content":      "exact.txt",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Map{
			"isMatch":   "true",
			"isInverse": "false",
		},
	},

	// === Case insensitive ===
	{
		Title: "Glob - ignore case matches different casing",
		ArrangeInput: args.Map{
			"when":         "given glob with ignore case",
			"pattern":      "*.GO",
			"content":      "main.go",
			"isIgnoreCase": true,
		},
		ExpectedInput: args.Map{
			"isMatch":   "true",
			"isInverse": "false",
		},
	},
	{
		Title: "Glob - case sensitive rejects different casing",
		ArrangeInput: args.Map{
			"when":         "given glob without ignore case",
			"pattern":      "*.GO",
			"content":      "main.go",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Map{
			"isMatch":   "false",
			"isInverse": "true",
		},
	},

	// === Negative: no match ===
	{
		Title: "Glob - no match returns false",
		ArrangeInput: args.Map{
			"when":         "given non-matching content",
			"pattern":      "*.go",
			"content":      "main.rs",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Map{
			"isMatch":   "false",
			"isInverse": "true",
		},
	},
	{
		Title: "Glob - question mark rejects multi-char",
		ArrangeInput: args.Map{
			"when":         "given content longer than question mark allows",
			"pattern":      "file?.txt",
			"content":      "fileAB.txt",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Map{
			"isMatch":   "false",
			"isInverse": "true",
		},
	},

	// === Edge: invalid pattern ===
	{
		Title: "Glob - invalid pattern returns false gracefully",
		ArrangeInput: args.Map{
			"when":         "given malformed glob pattern",
			"pattern":      "[invalid",
			"content":      "anything",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Map{
			"isMatch":   "false",
			"isInverse": "true",
		},
	},

	// === Edge: empty ===
	{
		Title: "Glob - empty pattern matches empty content",
		ArrangeInput: args.Map{
			"when":         "given empty pattern and content",
			"pattern":      "",
			"content":      "",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Map{
			"isMatch":   "true",
			"isInverse": "false",
		},
	},
	{
		Title: "Glob - star matches empty content",
		ArrangeInput: args.Map{
			"when":         "given star pattern with empty content",
			"pattern":      "*",
			"content":      "",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Map{
			"isMatch":   "true",
			"isInverse": "false",
		},
	},
	{
		Title: "Glob - star matches any content",
		ArrangeInput: args.Map{
			"when":         "given star pattern with any content",
			"pattern":      "*",
			"content":      "anything-at-all",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Map{
			"isMatch":   "true",
			"isInverse": "false",
		},
	},
}
