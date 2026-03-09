package regexnewtests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// =============================================================================
// CreateApplicableLock
// =============================================================================

var createApplicableLockTestCases = []coretestcases.CaseV1{
	{
		Title: "CreateApplicableLock valid pattern returns applicable",
		ArrangeInput: args.Map{
			"pattern": "\\d+",
		},
		ExpectedInput: args.Map{
			"regexNotNil":  true,
			"hasError":     false,
			"isApplicable": true,
		},
	},
	{
		Title: "CreateApplicableLock invalid pattern returns not applicable",
		ArrangeInput: args.Map{
			"pattern": "[bad",
		},
		ExpectedInput: args.Map{
			"regexNotNil":  false,
			"hasError":     true,
			"isApplicable": false,
		},
	},
}

// =============================================================================
// CreateLockIf
// =============================================================================

var createLockIfTestCases = []coretestcases.CaseV1{
	{
		Title: "CreateLockIf isLock=true valid pattern compiles",
		ArrangeInput: args.Map{
			"isLock":  true,
			"pattern": "^hello$",
		},
		ExpectedInput: args.Map{
			"regexNotNil": true,
			"hasError":    false,
		},
	},
	{
		Title: "CreateLockIf isLock=false valid pattern compiles",
		ArrangeInput: args.Map{
			"isLock":  false,
			"pattern": "^hello$",
		},
		ExpectedInput: args.Map{
			"regexNotNil": true,
			"hasError":    false,
		},
	},
	{
		Title: "CreateLockIf isLock=true invalid pattern returns error",
		ArrangeInput: args.Map{
			"isLock":  true,
			"pattern": "[broken",
		},
		ExpectedInput: args.Map{
			"regexNotNil": false,
			"hasError":    true,
		},
	},
}

// =============================================================================
// CreateMustLockIf
// =============================================================================

var createMustLockIfTestCases = []coretestcases.CaseV1{
	{
		Title: "CreateMustLockIf isLock=true valid pattern returns regex",
		ArrangeInput: args.Map{
			"isLock":  true,
			"pattern": "\\w+",
		},
		ExpectedInput: args.Map{
			"regexNotNil": true,
			"panicked":    false,
		},
	},
	{
		Title: "CreateMustLockIf isLock=false valid pattern returns regex",
		ArrangeInput: args.Map{
			"isLock":  false,
			"pattern": "\\w+",
		},
		ExpectedInput: args.Map{
			"regexNotNil": true,
			"panicked":    false,
		},
	},
	{
		Title: "CreateMustLockIf invalid pattern panics",
		ArrangeInput: args.Map{
			"isLock":  true,
			"pattern": "[bad",
		},
		ExpectedInput: args.Map{
			"regexNotNil": false,
			"panicked":    true,
		},
	},
}

// =============================================================================
// MatchUsingFuncErrorLock
// =============================================================================

var matchUsingFuncErrorLockTestCases = []coretestcases.CaseV1{
	{
		Title: "MatchUsingFuncErrorLock match returns nil error",
		ArrangeInput: args.Map{
			"pattern":  "\\d+",
			"comparing": "abc123",
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
	{
		Title: "MatchUsingFuncErrorLock mismatch returns error",
		ArrangeInput: args.Map{
			"pattern":  "^\\d+$",
			"comparing": "abc",
		},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
	{
		Title: "MatchUsingFuncErrorLock invalid regex returns error",
		ArrangeInput: args.Map{
			"pattern":  "[broken",
			"comparing": "test",
		},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
}

// =============================================================================
// MatchUsingCustomizeErrorFuncLock
// =============================================================================

var matchUsingCustomizeErrorFuncLockTestCases = []coretestcases.CaseV1{
	{
		Title: "CustomizeErrorFuncLock match returns nil",
		ArrangeInput: args.Map{
			"pattern":       "^hello$",
			"comparing":     "hello",
			"useCustomizer": false,
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
	{
		Title: "CustomizeErrorFuncLock mismatch nil customizer uses default error",
		ArrangeInput: args.Map{
			"pattern":       "^\\d+$",
			"comparing":     "abc",
			"useCustomizer": false,
		},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
	{
		Title: "CustomizeErrorFuncLock mismatch with custom error func",
		ArrangeInput: args.Map{
			"pattern":       "^\\d+$",
			"comparing":     "abc",
			"useCustomizer": true,
		},
		ExpectedInput: args.Map{
			"hasError":       true,
			"isCustomError": true,
		},
	},
}
