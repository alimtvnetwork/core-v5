package coredynamictests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================
// MapAnyItems — Add and AllKeys
// ==========================================

var mapAnyItemsAddAndKeysTestCase = coretestcases.CaseV1{
	Title: "Add returns stored items and AllKeys returns keys -- 3 items added",
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
	Title: "GetPagedCollection returns 5 pages -- 9 items paged by 2",
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
	Title: "JSON serialize then deserialize returns equal map -- 4 items",
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
	Title: "GetItemRef returns stored value -- existing key 'target-key'",
	ArrangeInput: args.Map{
		"when": "given key exists in map",
		"key":  "target-key",
	},
	ExpectedInput: args.Map{
		"hasItems": true,
	},
}
