package regexnewtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// New.Default — creates regex without lock (var-level usage)
// =============================================================================

var newDefaultTestCases = []coretestcases.MapGherkins{
	{
		Title: "New.Default returns compiled regex -- valid pattern '\\d+'",
		When:  "given a valid digit pattern",
		Input: args.Map{
			params.pattern: "\\d+",
		},
		Expected: args.Map{
			params.regexNotNil: true,
			params.hasError:    false,
		},
	},
	{
		Title: "New.Default returns error -- invalid pattern '[bad'",
		When:  "given an invalid pattern",
		Input: args.Map{
			params.pattern: "[bad",
		},
		Expected: args.Map{
			params.regexNotNil: false,
			params.hasError:    true,
		},
	},
}

// =============================================================================
// New.DefaultLockIf — creates regex with conditional lock
// =============================================================================

var newDefaultLockIfTestCases = []coretestcases.MapGherkins{
	{
		Title: "New.DefaultLockIf returns compiled regex -- lock true valid pattern",
		When:  "given valid pattern with lock true",
		Input: args.Map{
			params.pattern: "\\w+",
			params.isLock:  true,
		},
		Expected: args.Map{
			params.regexNotNil: true,
			params.hasError:    false,
		},
	},
	{
		Title: "New.DefaultLockIf returns compiled regex -- lock false valid pattern",
		When:  "given valid pattern with lock false",
		Input: args.Map{
			params.pattern: "[a-z]+",
			params.isLock:  false,
		},
		Expected: args.Map{
			params.regexNotNil: true,
			params.hasError:    false,
		},
	},
	{
		Title: "New.DefaultLockIf returns error -- lock true invalid pattern",
		When:  "given invalid pattern with lock true",
		Input: args.Map{
			params.pattern: "[bad",
			params.isLock:  true,
		},
		Expected: args.Map{
			params.regexNotNil: false,
			params.hasError:    true,
		},
	},
	{
		Title: "New.DefaultLockIf returns error -- lock false invalid pattern",
		When:  "given invalid pattern with lock false",
		Input: args.Map{
			params.pattern: "(unclosed",
			params.isLock:  false,
		},
		Expected: args.Map{
			params.regexNotNil: false,
			params.hasError:    true,
		},
	},
}

// =============================================================================
// New.DefaultApplicableLock — creates regex with applicability flag
// =============================================================================

var newDefaultApplicableLockTestCases = []coretestcases.MapGherkins{
	{
		Title: "New.DefaultApplicableLock returns applicable regex -- valid pattern '\\d+'",
		When:  "given a valid digit pattern",
		Input: args.Map{
			params.pattern: "\\d+",
		},
		Expected: args.Map{
			params.regexNotNil:  true,
			params.hasError:     false,
			params.isApplicable: true,
		},
	},
	{
		Title: "New.DefaultApplicableLock returns not applicable -- invalid pattern '[bad'",
		When:  "given an invalid pattern",
		Input: args.Map{
			params.pattern: "[bad",
		},
		Expected: args.Map{
			params.regexNotNil:  false,
			params.hasError:     true,
			params.isApplicable: false,
		},
	},
	{
		Title: "New.DefaultApplicableLock returns applicable -- empty pattern",
		When:  "given an empty pattern",
		Input: args.Map{
			params.pattern: "",
		},
		Expected: args.Map{
			params.regexNotNil:  true,
			params.hasError:     false,
			params.isApplicable: true,
		},
	},
}

// =============================================================================
// New.LazyRegex.NewLockIf — conditional lock lazy creation
// =============================================================================

var newLazyRegexNewLockIfTestCases = []coretestcases.MapGherkins{
	{
		Title: "NewLockIf returns applicable LazyRegex -- lock true valid pattern",
		When:  "given valid pattern with lock true",
		Input: args.Map{
			params.pattern:      "\\d+",
			params.compareInput: "123",
			params.isLock:       true,
		},
		Expected: args.Map{
			params.isDefined:    true,
			params.isApplicable: true,
			params.isMatch:      true,
		},
	},
	{
		Title: "NewLockIf returns applicable LazyRegex -- lock false valid pattern",
		When:  "given valid pattern with lock false",
		Input: args.Map{
			params.pattern:      "[a-z]+",
			params.compareInput: "hello",
			params.isLock:       false,
		},
		Expected: args.Map{
			params.isDefined:    true,
			params.isApplicable: true,
			params.isMatch:      true,
		},
	},
	{
		Title: "NewLockIf returns non-applicable LazyRegex -- lock true invalid pattern",
		When:  "given invalid pattern with lock true",
		Input: args.Map{
			params.pattern:      "[bad",
			params.compareInput: "test",
			params.isLock:       true,
		},
		Expected: args.Map{
			params.isDefined:    true,
			params.isApplicable: false,
			params.isMatch:      false,
		},
	},
}

// =============================================================================
// New.LazyRegex.AllPatternsMap — returns all cached patterns
// =============================================================================

var allPatternsMapTestCases = []coretestcases.MapGherkins{
	{
		Title: "AllPatternsMap returns non-empty map -- after creating patterns",
		When:  "given patterns have been created",
		Input: args.Map{
			params.pattern: "\\d+",
		},
		Expected: args.Map{
			params.isNotEmpty: true,
		},
	},
}
