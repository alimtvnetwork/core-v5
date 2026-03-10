package coredynamictests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================
// MapAnyItems — Add and AllKeys
// ==========================================

var mapAnyItemsAddAndKeysTestCase = coretestcases.CaseV1{
	Title: "MapAnyItems Add stores items and AllKeys returns keys",
	ArrangeInput: args.Map{
		"when":     "given 3 items added",
		"capacity": 10,
		"keys":     []string{"key1", "key2", "key3"},
	},
	ExpectedInput: args.Map{
		"keyCount": 3,
		"hasAll":   true,
	},
}

// ==========================================
// MapAnyItems — GetPagedCollection
// ==========================================

var mapAnyItemsPagedTestCase = coretestcases.CaseV1{
	Title: "GetPagedCollection splits items into pages of given size",
	ArrangeInput: args.Map{
		"when":      "given 9 items paged by 2",
		"itemCount": 9,
		"pageSize":  2,
	},
	ExpectedInput: args.Map{
		"pageCount": 5,
	},
}

// ==========================================
// MapAnyItems — JSON roundtrip
// ==========================================

var mapAnyItemsJsonRoundtripTestCase = coretestcases.CaseV1{
	Title: "MapAnyItems JSON serialize then deserialize is equal",
	ArrangeInput: args.Map{
		"when":      "given map serialized and deserialized",
		"itemCount": 4,
	},
	ExpectedInput: args.Map{
		"isEqual": true,
	},
}

// ==========================================
// MapAnyItems — GetItemRef
// ==========================================

var mapAnyItemsGetItemRefTestCase = coretestcases.CaseV1{
	Title: "GetItemRef populates target with stored value",
	ArrangeInput: args.Map{
		"when": "given key exists in map",
		"key":  "target-key",
	},
	ExpectedInput: args.Map{
		"hasItems": true,
	},
}
