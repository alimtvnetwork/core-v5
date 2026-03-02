package corepayloadtests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================
// Attributes.IsEqual — Regression for logic inversion bug
// ==========================================

var attributesIsEqualTestCases = []coretestcases.CaseV1{
	{
		Title: "Both nil Attributes are equal",
		ArrangeInput: args.Map{
			"when":     "both attributes are nil",
			"left_nil": true,
			"right_nil": true,
		},
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "Left nil, right non-nil are not equal",
		ArrangeInput: args.Map{
			"when":     "left is nil right is not",
			"left_nil": true,
			"right_nil": false,
		},
		ExpectedInput: []string{
			"false",
		},
	},
	{
		Title: "Right nil, left non-nil are not equal",
		ArrangeInput: args.Map{
			"when":     "right is nil left is not",
			"left_nil": false,
			"right_nil": true,
		},
		ExpectedInput: []string{
			"false",
		},
	},
	{
		Title: "Same pointer returns equal",
		ArrangeInput: args.Map{
			"when":         "same pointer identity",
			"same_pointer": true,
		},
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "Equal dynamic payloads are equal",
		ArrangeInput: args.Map{
			"when":    "same dynamic payloads on both",
			"payload": "test-data",
		},
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "Different dynamic payloads are not equal",
		ArrangeInput: args.Map{
			"when":          "different dynamic payloads",
			"left_payload":  "data-a",
			"right_payload": "data-b",
		},
		ExpectedInput: []string{
			"false",
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
		ExpectedInput: []string{
			"true",  // clonedPtr is nil
			"false", // no error
		},
	},
	{
		Title: "Shallow clone preserves dynamic payloads",
		ArrangeInput: args.Map{
			"when":    "shallow clone with dynamic payloads",
			"payload": "clone-payload",
			"deep":    false,
		},
		ExpectedInput: []string{
			"clone-payload", // cloned payload string
			"true",          // is equal to original
		},
	},
	{
		Title: "Deep clone creates independent copy",
		ArrangeInput: args.Map{
			"when":    "deep clone then mutate original",
			"payload": "deep-clone-data",
			"deep":    true,
		},
		ExpectedInput: []string{
			"deep-clone-data", // cloned payload preserved
			"true",            // is equal before mutation
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
		ExpectedInput: []string{
			"false",
		},
	},
	{
		Title: "Empty attributes IsSafeValid returns false",
		ArrangeInput: args.Map{
			"when":  "attributes has no data",
			"empty": true,
		},
		ExpectedInput: []string{
			"false",
		},
	},
	{
		Title: "Attributes with payload IsSafeValid returns true",
		ArrangeInput: args.Map{
			"when":    "attributes has dynamic payload",
			"payload": "valid-data",
		},
		ExpectedInput: []string{
			"true",
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
		ExpectedInput: []string{
			"true", // result is nil
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
		ExpectedInput: []string{
			"user-42",
			"login",
			"/api/data",
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
		ExpectedInput: []string{
			"read",  // original unchanged
			"write", // clone mutated
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
		ExpectedInput: []string{
			"",
			"delete",
			"/api/remove",
		},
	},
}
