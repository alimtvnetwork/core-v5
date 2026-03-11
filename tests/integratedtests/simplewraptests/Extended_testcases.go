package simplewraptests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ============================================================================
// WithDoubleQuote
// ============================================================================

var withDoubleQuoteTestCases = []coretestcases.CaseV1{
	{
		Title: "WithDoubleQuote wraps string",
		ArrangeInput: args.Map{
			"input": "hello",
		},
		ExpectedInput: "\"hello\"",
	},
	{
		Title: "WithDoubleQuote wraps empty string",
		ArrangeInput: args.Map{
			"input": "",
		},
		ExpectedInput: "\"\"",
	},
}

// ============================================================================
// WithSingleQuote
// ============================================================================

var withSingleQuoteTestCases = []coretestcases.CaseV1{
	{
		Title: "WithSingleQuote wraps string",
		ArrangeInput: args.Map{
			"input": "hello",
		},
		ExpectedInput: "'hello'",
	},
}

// ============================================================================
// WithCurly
// ============================================================================

var withCurlyTestCases = []coretestcases.CaseV1{
	{
		Title: "WithCurly wraps string",
		ArrangeInput: args.Map{
			"input": "hello",
		},
		ExpectedInput: "{hello}",
	},
	{
		Title: "WithCurly wraps integer",
		ArrangeInput: args.Map{
			"input": 42,
		},
		ExpectedInput: "{42}",
	},
}

// ============================================================================
// WithParenthesis
// ============================================================================

var withParenthesisTestCases = []coretestcases.CaseV1{
	{
		Title: "WithParenthesis wraps string",
		ArrangeInput: args.Map{
			"input": "hello",
		},
		ExpectedInput: "(hello)",
	},
}

// ============================================================================
// WithBrackets (single string)
// ============================================================================

var withBracketsStrTestCases = []coretestcases.CaseV1{
	{
		Title: "WithBrackets wraps string in square brackets",
		ArrangeInput: args.Map{
			"input": "hello",
		},
		ExpectedInput: "[hello]",
	},
}

// ============================================================================
// TitleSquare
// ============================================================================

var titleSquareTestCases = []coretestcases.CaseV1{
	{
		Title: "TitleSquare formats title: [value]",
		ArrangeInput: args.Map{
			"title": "my title",
			"value": "my value",
		},
		ExpectedInput: "my title: [my value]",
	},
}

// ============================================================================
// TitleSquareMeta
// ============================================================================

var titleSquareMetaTestCases = []coretestcases.CaseV1{
	{
		Title: "TitleSquareMeta formats title: [value] (meta)",
		ArrangeInput: args.Map{
			"title": "my title",
			"value": "my value",
			"meta":  "meta info",
		},
		ExpectedInput: "my title: [my value] (meta info)",
	},
}

// ============================================================================
// With
// ============================================================================

var withTestCases = []coretestcases.CaseV1{
	{
		Title: "With concatenates start source end",
		ArrangeInput: args.Map{
			"start":  "<<",
			"source": "hello",
			"end":    ">>",
		},
		ExpectedInput: "<<hello>>",
	},
	{
		Title: "With handles empty strings",
		ArrangeInput: args.Map{
			"start":  "",
			"source": "hello",
			"end":    "",
		},
		ExpectedInput: "hello",
	},
}

// ============================================================================
// WithPtr
// ============================================================================

var withPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "WithPtr concatenates with non-nil ptrs",
		ArrangeInput: args.Map{
			"start":  "<<",
			"source": "hello",
			"end":    ">>",
		},
		ExpectedInput: "<<hello>>",
	},
	{
		Title: "WithPtr handles nil start",
		ArrangeInput: args.Map{
			"startNil": true,
			"source":   "hello",
			"end":      ">>",
		},
		ExpectedInput: "hello>>",
	},
}

// ============================================================================
// ToJsonName
// ============================================================================

var toJsonNameTestCases = []coretestcases.CaseV1{
	{
		Title: "ToJsonName wraps in double quotes",
		ArrangeInput: args.Map{
			"input": "myField",
		},
		ExpectedInput: "\"myField\"",
	},
}

// ============================================================================
// ConditionalWrapWith
// ============================================================================

var conditionalWrapWithTestCases = []coretestcases.CaseV1{
	{
		Title: "ConditionalWrapWith adds both when missing",
		ArrangeInput: args.Map{
			"input": "hello",
			"start": byte('{'),
			"end":   byte('}'),
		},
		ExpectedInput: "{hello}",
	},
	{
		Title: "ConditionalWrapWith preserves existing wrap",
		ArrangeInput: args.Map{
			"input": "{hello}",
			"start": byte('{'),
			"end":   byte('}'),
		},
		ExpectedInput: "{hello}",
	},
	{
		Title: "ConditionalWrapWith adds end when only start present",
		ArrangeInput: args.Map{
			"input": "{hello",
			"start": byte('{'),
			"end":   byte('}'),
		},
		ExpectedInput: "{hello}",
	},
	{
		Title: "ConditionalWrapWith adds start when only end present",
		ArrangeInput: args.Map{
			"input": "hello}",
			"start": byte('{'),
			"end":   byte('}'),
		},
		ExpectedInput: "{hello}",
	},
	{
		Title: "ConditionalWrapWith handles empty string",
		ArrangeInput: args.Map{
			"input": "",
			"start": byte('{'),
			"end":   byte('}'),
		},
		ExpectedInput: "{}",
	},
	{
		Title: "ConditionalWrapWith handles single char matching start",
		ArrangeInput: args.Map{
			"input": "{",
			"start": byte('{'),
			"end":   byte('}'),
		},
		ExpectedInput: "{}", 
	},
}

// ============================================================================
// WithDoubleQuoteAny
// ============================================================================

var withDoubleQuoteAnyTestCases = []coretestcases.CaseV1{
	{
		Title: "WithDoubleQuoteAny wraps any value",
		ArrangeInput: args.Map{
			"input": 42,
		},
		ExpectedInput: "\"42\"",
	},
}
