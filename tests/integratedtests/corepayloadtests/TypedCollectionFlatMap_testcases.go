package corepayloadtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
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
		Title: "FlatMapTypedPayloads returns 6 tags -- 3 users with 2 tags each",
		ArrangeInput: args.Map{
			"when": "given 3 users with 2 tags each",
		},
		ExpectedInput: args.Map{
			"count": 6,
			"tag0":  "go", "tag1": "rust",
			"tag2": "python", "tag3": "java",
			"tag4": "ts", "tag5": "js",
		},
	},
}

// =============================================================================
// FlatMapTypedPayloadData — data-level
// =============================================================================

var flatMapTypedPayloadDataTestCases = []coretestcases.CaseV1{
	{
		Title: "FlatMapTypedPayloadData returns 6 tags -- 3 users data accessor",
		ArrangeInput: args.Map{
			"when": "given 3 users with 2 tags each via data accessor",
		},
		ExpectedInput: args.Map{
			"count": 6,
			"tag0":  "go", "tag1": "rust",
			"tag2": "python", "tag3": "java",
			"tag4": "ts", "tag5": "js",
		},
	},
}

// =============================================================================
// FlatMap on empty collection
// =============================================================================

var flatMapEmptyCollectionTestCases = []coretestcases.CaseV1{
	{
		Title: "FlatMapTypedPayloads returns count 0 -- empty collection",
		ArrangeInput: args.Map{
			"when": "given empty collection",
		},
		ExpectedInput: args.Map{
			"count": 0,
		},
	},
}

// =============================================================================
// FlatMap producing empty slices per item
// =============================================================================

var flatMapNoOutputTestCases = []coretestcases.CaseV1{
	{
		Title: "FlatMapTypedPayloadData returns count 0 -- mapper returns nil slices",
		ArrangeInput: args.Map{
			"when": "mapper returns nil for each item",
		},
		ExpectedInput: args.Map{
			"count": 0,
		},
	},
}

// =============================================================================
// Edge: nil wrappers in collection
// =============================================================================

var nilWrapperEdgeCaseTestCases = []coretestcases.CaseV1{
	{
		Title: "TypedPayloadCollection returns isValid false -- nil wrapper in collection",
		ArrangeInput: args.Map{
			"when": "collection contains a nil wrapper",
		},
		ExpectedInput: args.Map{
			"isValid": false,
			"length":  4,
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
		ExpectedInput: args.Map{
			"count": 2,
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
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
}

// Note: Nil receiver test cases migrated to TypedCollection_NilReceiver_testcases.go
// using CaseNilSafe pattern with function literal wrappers for generic types.
