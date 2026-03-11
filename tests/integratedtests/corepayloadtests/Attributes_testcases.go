package corepayloadtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================
// Attributes.IsEqual — Regression for logic inversion bug
// ==========================================

var attributesIsEqualTestCases = []coretestcases.CaseV1{
	{
		Title: "Attributes.IsEqual returns true -- both nil",
		ArrangeInput: args.Map{
			"when":      "both attributes are nil",
			"left_nil":  true,
			"right_nil": true,
		},
		ExpectedInput: args.Map{
			"isEqual": true,
		},
	},
	{
		Title: "Attributes.IsEqual returns false -- left nil right non-nil",
		ArrangeInput: args.Map{
			"when":      "left is nil right is not",
			"left_nil":  true,
			"right_nil": false,
		},
		ExpectedInput: args.Map{
			"isEqual": false,
		},
	},
	{
		Title: "Attributes.IsEqual returns false -- right nil left non-nil",
		ArrangeInput: args.Map{
			"when":      "right is nil left is not",
			"left_nil":  false,
			"right_nil": true,
		},
		ExpectedInput: args.Map{
			"isEqual": false,
		},
	},
	{
		Title: "Attributes.IsEqual returns true -- same pointer identity",
		ArrangeInput: args.Map{
			"when":         "same pointer identity",
			"same_pointer": true,
		},
		ExpectedInput: args.Map{
			"isEqual": true,
		},
	},
	{
		Title: "Attributes.IsEqual returns true -- same dynamic payloads",
		ArrangeInput: args.Map{
			"when":    "same dynamic payloads on both",
			"payload": "test-data",
		},
		ExpectedInput: args.Map{
			"isEqual": true,
		},
	},
	{
		Title: "Different dynamic payloads are not equal",
		ArrangeInput: args.Map{
			"when":          "different dynamic payloads",
			"left_payload":  "data-a",
			"right_payload": "data-b",
		},
		ExpectedInput: args.Map{
			"isEqual": false,
		},
	},
}

// ==========================================
// Attributes.Clone — Regression for deep clone independence
// ==========================================

var attributesCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "Nil attributes shallow clone returns nil",
		ArrangeInput: args.Map{
			"when":     "attributes is nil",
			"nil_attr": true,
			"deep":     false,
		},
		ExpectedInput: args.Map{
			"isNil":    true,
			"hasError": false,
		},
	},
	{
		Title: "Shallow clone preserves dynamic payloads",
		ArrangeInput: args.Map{
			"when":    "shallow clone with dynamic payloads",
			"payload": "clone-payload",
			"deep":    false,
		},
		ExpectedInput: args.Map{
			"clonedPayload": "clone-payload",
			"isEqual":       true,
		},
	},
	{
		Title: "Deep clone returns error",
		ArrangeInput: args.Map{
			"when":    "deep clone then mutate original",
			"payload": "deep-clone-data",
			"deep":    true,
		},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
}

// ==========================================
// Attributes.IsSafeValid — Regression for negation bug
// ==========================================

var attributesIsSafeValidTestCases = []coretestcases.CaseV1{
	{
		Title: "Nil attributes IsSafeValid returns false",
		ArrangeInput: args.Map{
			"when":     "attributes is nil",
			"nil_attr": true,
		},
		ExpectedInput: args.Map{
			"isSafeValid": false,
		},
	},
	{
		Title: "Empty attributes IsSafeValid returns false",
		ArrangeInput: args.Map{
			"when":  "attributes has no data",
			"empty": true,
		},
		ExpectedInput: args.Map{
			"isSafeValid": false,
		},
	},
	{
		Title: "Attributes with payload IsSafeValid returns true",
		ArrangeInput: args.Map{
			"when":    "attributes has dynamic payload",
			"payload": "valid-data",
		},
		ExpectedInput: args.Map{
			"isSafeValid": true,
		},
	},
}

// ==========================================
// AuthInfo.Clone — Regression for missing Identifier field
// ==========================================

var authInfoCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "Nil AuthInfo ClonePtr returns nil",
		ArrangeInput: args.Map{
			"when":     "auth info is nil",
			"nil_auth": true,
		},
		ExpectedInput: args.Map{
			"isNil": true,
		},
	},
	{
		Title: "Clone preserves Identifier field",
		ArrangeInput: args.Map{
			"when":          "auth info has identifier",
			"identifier":    "user-42",
			"action_type":   "login",
			"resource_name": "/api/data",
		},
		ExpectedInput: args.Map{
			"identifier":   "user-42",
			"actionType":   "login",
			"resourceName": "/api/data",
		},
	},
	{
		Title: "Clone is independent — mutating clone does not affect original",
		ArrangeInput: args.Map{
			"when":            "clone mutated after creation",
			"identifier":      "original-id",
			"action_type":     "read",
			"new_action_type": "write",
		},
		ExpectedInput: args.Map{
			"originalAction": "read",
			"clonedAction":   "write",
		},
	},
	{
		Title: "Clone preserves all fields with empty Identifier",
		ArrangeInput: args.Map{
			"when":          "identifier is empty string",
			"identifier":    "",
			"action_type":   "delete",
			"resource_name": "/api/remove",
		},
		ExpectedInput: args.Map{
			"identifier":   "",
			"actionType":   "delete",
			"resourceName": "/api/remove",
		},
	},
}
