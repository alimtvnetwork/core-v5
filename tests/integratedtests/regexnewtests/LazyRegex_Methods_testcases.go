package regexnewtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var lazyRegexCompileTestCases = []coretestcases.MapGherkins{
	{
		Title: "Compile valid pattern returns no error",
		When:  "given a valid pattern to Compile",
		Input: args.Map{
			params.pattern: "\\d+",
		},
		Expected: args.Map{
			params.regexNotNil: true,
			params.hasError:    false,
			params.isCompiled:  true,
		},
	},
	{
		Title: "Compile invalid pattern returns error",
		When:  "given an invalid pattern to Compile",
		Input: args.Map{
			params.pattern: "[bad",
		},
		Expected: args.Map{
			params.regexNotNil: false,
			params.hasError:    true,
			params.isCompiled:  true,
		},
	},
	{
		Title: "Compile empty pattern on undefined lazy returns error",
		When:  "given empty pattern to Compile",
		Input: args.Map{
			params.pattern: "",
		},
		Expected: args.Map{
			params.regexNotNil: false,
			params.hasError:    true,
			params.isCompiled:  false,
		},
	},
}

var lazyRegexHasErrorTestCases = []coretestcases.MapGherkins{
	{
		Title: "HasError returns false for valid pattern",
		When:  "given valid pattern for HasError",
		Input: args.Map{
			params.pattern: "hello",
		},
		Expected: args.Map{
			params.hasError:  false,
			params.isInvalid: false,
		},
	},
	{
		Title: "HasError returns true for invalid pattern",
		When:  "given invalid pattern for HasError",
		Input: args.Map{
			params.pattern: "[broken",
		},
		Expected: args.Map{
			params.hasError:  true,
			params.isInvalid: true,
		},
	},
}

var lazyRegexMatchBytesTestCases = []coretestcases.MapGherkins{
	{
		Title: "IsMatchBytes returns true for matching bytes",
		When:  "given matching byte input",
		Input: args.Map{
			params.pattern:      "\\d+",
			params.compareInput: "abc123",
		},
		Expected: args.Map{
			params.isMatchBytes:       true,
			params.isFailedMatchBytes: false,
		},
	},
	{
		Title: "IsMatchBytes returns false for non-matching bytes",
		When:  "given non-matching byte input",
		Input: args.Map{
			params.pattern:      "^\\d+$",
			params.compareInput: "abc",
		},
		Expected: args.Map{
			params.isMatchBytes:       false,
			params.isFailedMatchBytes: true,
		},
	},
}

var lazyRegexMatchErrorTestCases = []coretestcases.MapGherkins{
	{
		Title: "LazyRegex.MatchError returns nil on match",
		When:  "given matching input to LazyRegex.MatchError",
		Input: args.Map{
			params.pattern:      "^hello$",
			params.compareInput: "hello",
		},
		Expected: args.Map{
			params.isNoError: true,
		},
	},
	{
		Title: "LazyRegex.MatchError returns error on mismatch",
		When:  "given non-matching input to LazyRegex.MatchError",
		Input: args.Map{
			params.pattern:      "^\\d+$",
			params.compareInput: "abc",
		},
		Expected: args.Map{
			params.isNoError: false,
		},
	},
	{
		Title: "LazyRegex.MatchError returns error for invalid regex",
		When:  "given invalid regex to LazyRegex.MatchError",
		Input: args.Map{
			params.pattern:      "[bad",
			params.compareInput: "test",
		},
		Expected: args.Map{
			params.isNoError: false,
		},
	},
}

var lazyRegexStringTestCases = []coretestcases.MapGherkins{
	{
		Title: "String returns pattern for valid LazyRegex",
		When:  "given valid pattern for String",
		Input: args.Map{
			params.pattern: "\\d+",
		},
		Expected: args.Map{
			params.stringResult: "\\d+",
		},
	},
	{
		Title: "Pattern returns the original pattern",
		When:  "given email pattern for Pattern",
		Input: args.Map{
			params.pattern: `[a-z]+@[a-z]+\.[a-z]+`,
		},
		Expected: args.Map{
			params.stringResult: `[a-z]+@[a-z]+\.[a-z]+`,
		},
	},
}
