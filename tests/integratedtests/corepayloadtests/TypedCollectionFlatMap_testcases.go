package corepayloadtests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// testUserWithTags extends testUser with tags for FlatMap testing.
type testUserWithTags struct {
	Name string   `json:"name"`
	Age  int      `json:"age"`
	Tags []string `json:"tags"`
}

// =============================================================================
// FlatMapTypedPayloads — wrapper-level
// =============================================================================

var flatMapTypedPayloadsTestCases = []coretestcases.CaseV1{
	{
		Title: "FlatMapTypedPayloads flattens tags from all users",
		ArrangeInput: args.Map{
			"when": "given 3 users with 2 tags each",
		},
		ExpectedInput: []string{
			"6",
			"go",
			"rust",
			"python",
			"java",
			"ts",
			"js",
		},
	},
}

// =============================================================================
// FlatMapTypedPayloadData — data-level
// =============================================================================

var flatMapTypedPayloadDataTestCases = []coretestcases.CaseV1{
	{
		Title: "FlatMapTypedPayloadData flattens tags via data accessor",
		ArrangeInput: args.Map{
			"when": "given 3 users with 2 tags each via data accessor",
		},
		ExpectedInput: []string{
			"6",
			"go",
			"rust",
			"python",
			"java",
			"ts",
			"js",
		},
	},
}

// =============================================================================
// FlatMap on empty collection
// =============================================================================

var flatMapEmptyCollectionTestCases = []coretestcases.CaseV1{
	{
		Title: "FlatMapTypedPayloads on empty collection returns empty",
		ArrangeInput: args.Map{
			"when": "given empty collection",
		},
		ExpectedInput: []string{
			"0",
		},
	},
}

// =============================================================================
// FlatMap producing empty slices per item
// =============================================================================

var flatMapNoOutputTestCases = []coretestcases.CaseV1{
	{
		Title: "FlatMapTypedPayloadData returns empty when mapper returns nil slices",
		ArrangeInput: args.Map{
			"when": "mapper returns nil for each item",
		},
		ExpectedInput: []string{
			"0",
		},
	},
}

// =============================================================================
// Edge: nil wrappers in collection
// =============================================================================

var nilWrapperEdgeCaseTestCases = []coretestcases.CaseV1{
	{
		Title: "Collection with nil wrapper — IsValid returns false",
		ArrangeInput: args.Map{
			"when": "collection contains a nil wrapper",
		},
		ExpectedInput: []string{
			"false",
			"3",
		},
	},
}

// =============================================================================
// Edge: deserialization failure via TypedPayloadCollectionFromPayloads
// =============================================================================

var deserializationFailureTestCases = []coretestcases.CaseV1{
	{
		Title: "TypedPayloadCollectionFromPayloads skips items with invalid payloads",
		ArrangeInput: args.Map{
			"when":          "2 valid + 1 invalid payload wrappers",
			"valid_count":   2,
			"invalid_count": 1,
		},
		ExpectedInput: []string{
			"2",
		},
	},
}

// =============================================================================
// Edge: deserialization failure via TypedPayloadCollectionDeserialize
// =============================================================================

var collectionDeserializeInvalidTestCases = []coretestcases.CaseV1{
	{
		Title: "TypedPayloadCollectionDeserialize with invalid bytes returns error",
		ArrangeInput: args.Map{
			"when":  "passing invalid json bytes",
			"bytes": "{{not-json-at-all",
		},
		ExpectedInput: []string{
			"true",
		},
	},
}

// =============================================================================
// Edge: nil receiver safety
// =============================================================================

var nilReceiverTestCases = []coretestcases.CaseV1{
	{
		Title: "Nil collection — Length returns 0 and IsEmpty returns true",
		ArrangeInput: args.Map{
			"when": "collection pointer is nil",
		},
		ExpectedInput: []string{
			"0",
			"true",
			"true",
		},
	},
}
