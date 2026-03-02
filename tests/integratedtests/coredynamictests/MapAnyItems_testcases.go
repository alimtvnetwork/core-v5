package coredynamictests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================
// MapAnyItems — Add and AllKeys
// ==========================================

var mapAnyItemsAddAndKeysTestCases = []coretestcases.CaseV1{
	{
		Title: "MapAnyItems Add stores items and AllKeys returns keys",
		ArrangeInput: args.Map{
			"when":     "given 3 items added",
			"capacity": 10,
			"keys":     []string{"key1", "key2", "key3"},
		},
		ExpectedInput: []string{
			"3",
			"true",
		},
	},
}

// ==========================================
// MapAnyItems — GetPagedCollection
// ==========================================

var mapAnyItemsPagedTestCases = []coretestcases.CaseV1{
	{
		Title: "GetPagedCollection splits items into pages of given size",
		ArrangeInput: args.Map{
			"when":      "given 9 items paged by 2",
			"itemCount": 9,
			"pageSize":  2,
		},
		ExpectedInput: []string{
			"5",
		},
	},
}

// ==========================================
// MapAnyItems — JSON roundtrip
// ==========================================

var mapAnyItemsJsonRoundtripTestCases = []coretestcases.CaseV1{
	{
		Title: "MapAnyItems JSON serialize then deserialize is equal",
		ArrangeInput: args.Map{
			"when":      "given map serialized and deserialized",
			"itemCount": 4,
		},
		ExpectedInput: []string{
			"true",
		},
	},
}

// ==========================================
// MapAnyItems — GetItemRef
// ==========================================

var mapAnyItemsGetItemRefTestCases = []coretestcases.CaseV1{
	{
		Title: "GetItemRef populates target with stored value",
		ArrangeInput: args.Map{
			"when": "given key exists in map",
			"key":  "target-key",
		},
		ExpectedInput: []string{
			"true",
		},
	},
}
