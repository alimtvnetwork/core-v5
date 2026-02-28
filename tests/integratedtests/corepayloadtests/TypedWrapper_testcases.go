package corepayloadtests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// testProduct is a sample struct for TypedPayloadWrapper deserialization tests.
type testProduct struct {
	SKU   string  `json:"sku"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

var typedWrapperDeserializationTestCases = []coretestcases.CaseV1{
	{
		Title: "Deserialize typed wrapper preserves all fields",
		ArrangeInput: args.Map{
			"name":  "product-create",
			"id":    "prod-1",
			"sku":   "SKU-100",
			"title": "Widget",
			"price": 29.99,
		},
		ExpectedInput: []string{
			"product-create",
			"prod-1",
			"SKU-100",
			"Widget",
			"29.99",
		},
	},
	{
		Title: "Deserialize typed wrapper with empty payload fields",
		ArrangeInput: args.Map{
			"name":  "empty-product",
			"id":    "prod-2",
			"sku":   "",
			"title": "",
			"price": 0.0,
		},
		ExpectedInput: []string{
			"empty-product",
			"prod-2",
			"",
			"",
			"0.00",
		},
	},
}

var typedWrapperRoundTripTestCases = []coretestcases.CaseV1{
	{
		Title: "Round-trip serialize then deserialize preserves data",
		ArrangeInput: args.Map{
			"name":  "round-trip",
			"id":    "rt-1",
			"sku":   "RT-SKU",
			"title": "Round Trip Product",
			"price": 55.50,
		},
		ExpectedInput: []string{
			"round-trip",
			"rt-1",
			"RT-SKU",
			"Round Trip Product",
			"55.50",
		},
	},
	{
		Title: "Round-trip with special characters in title",
		ArrangeInput: args.Map{
			"name":  "special-chars",
			"id":    "sc-1",
			"sku":   "SC-001",
			"title": `Quote "test" & <html>`,
			"price": 0.01,
		},
		ExpectedInput: []string{
			"special-chars",
			"sc-1",
			"SC-001",
			`Quote "test" & <html>`,
			"0.01",
		},
	},
}

var typedWrapperCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "Deep clone produces independent copy",
		ArrangeInput: args.Map{
			"name":  "clone-source",
			"id":    "cl-1",
			"sku":   "CL-SKU",
			"title": "Original",
			"price": 100.0,
		},
		ExpectedInput: []string{
			"clone-source",
			"cl-1",
			"CL-SKU",
			"Original",
			"100.00",
			"Modified",
		},
	},
}

var typedWrapperSetDataTestCases = []coretestcases.CaseV1{
	{
		Title: "SetTypedData updates both typed data and raw payloads",
		ArrangeInput: args.Map{
			"name":      "set-data",
			"id":        "sd-1",
			"sku":       "SD-SKU",
			"title":     "Before",
			"price":     10.0,
			"new_title": "After",
			"new_price": 20.0,
		},
		ExpectedInput: []string{
			"After",
			"20.00",
			"After",
			"20.00",
		},
	},
}

var typedWrapperNilAndInvalidTestCases = []coretestcases.CaseV1{
	{
		Title: "Nil wrapper returns error on creation",
		ArrangeInput: args.Map{
			"when": "passing nil wrapper",
		},
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "Invalid JSON bytes return error on deserialization",
		ArrangeInput: args.Map{
			"when":  "passing invalid json",
			"bytes": "not-valid-json{{{",
		},
		ExpectedInput: []string{
			"true",
		},
	},
}

var typedWrapperDeserializeToManyTestCases = []coretestcases.CaseV1{
	{
		Title: "DeserializeToMany parses array of typed wrappers",
		ArrangeInput: args.Map{
			"count": 3,
		},
		ExpectedInput: []string{
			"3",
			"item-0",
			"item-1",
			"item-2",
		},
	},
}
