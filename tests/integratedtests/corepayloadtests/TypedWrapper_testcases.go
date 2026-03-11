package corepayloadtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// testProduct is a sample struct for TypedPayloadWrapper deserialization tests.
type testProduct struct {
	SKU   string  `json:"sku"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

var typedWrapperDeserializationTestCases = []coretestcases.CaseV1{
	{
		Title: "TypedPayloadWrapper.Deserialize returns all fields -- name 'product-create'",
		ArrangeInput: args.Map{
			"name": "product-create", "id": "prod-1",
			"sku": "SKU-100", "title": "Widget", "price": 29.99,
		},
		ExpectedInput: args.Map{
			"name": "product-create", "id": "prod-1",
			"sku": "SKU-100", "title": "Widget", "price": "29.99",
		},
	},
	{
		Title: "TypedPayloadWrapper.Deserialize returns empty fields -- name 'empty-product'",
		ArrangeInput: args.Map{
			"name": "empty-product", "id": "prod-2",
			"sku": "", "title": "", "price": 0.0,
		},
		ExpectedInput: args.Map{
			"name": "empty-product", "id": "prod-2",
			"sku": "", "title": "", "price": "0.00",
		},
	},
}

var typedWrapperRoundTripTestCases = []coretestcases.CaseV1{
	{
		Title: "TypedPayloadWrapper returns preserved data -- round-trip 'round-trip'",
		ArrangeInput: args.Map{
			"name": "round-trip", "id": "rt-1",
			"sku": "RT-SKU", "title": "Round Trip Product", "price": 55.50,
		},
		ExpectedInput: args.Map{
			"name": "round-trip", "id": "rt-1",
			"sku": "RT-SKU", "title": "Round Trip Product", "price": "55.50",
		},
	},
	{
		Title: "TypedPayloadWrapper returns preserved special chars -- round-trip 'special-chars'",
		ArrangeInput: args.Map{
			"name": "special-chars", "id": "sc-1",
			"sku": "SC-001", "title": `Quote "test" & <html>`, "price": 0.01,
		},
		ExpectedInput: args.Map{
			"name": "special-chars", "id": "sc-1",
			"sku": "SC-001", "title": `Quote "test" & <html>`, "price": "0.01",
		},
	},
}

var typedWrapperCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "Deep clone produces independent copy",
		ArrangeInput: args.Map{
			"name": "clone-source", "id": "cl-1",
			"sku": "CL-SKU", "title": "Original", "price": 100.0,
		},
		ExpectedInput: args.Map{
			"originalName": "clone-source", "originalId": "cl-1",
			"originalSku": "CL-SKU", "originalTitle": "Original",
			"originalPrice": "100.00", "clonedTitle": "Modified",
		},
	},
}

var typedWrapperSetDataTestCases = []coretestcases.CaseV1{
	{
		Title: "SetTypedData updates both typed data and raw payloads",
		ArrangeInput: args.Map{
			"name": "set-data", "id": "sd-1",
			"sku": "SD-SKU", "title": "Before", "price": 10.0,
			"new_title": "After", "new_price": 20.0,
		},
		ExpectedInput: args.Map{
			"directTitle": "After", "directPrice": "20.00",
			"reparsedTitle": "After", "reparsedPrice": "20.00",
		},
	},
}

var typedWrapperNilTestCase = coretestcases.CaseV1{
	Title: "Nil wrapper returns error on creation",
	ArrangeInput: args.Map{
		"when": "passing nil wrapper",
	},
	ExpectedInput: args.Map{
		"hasError": true,
	},
}

var typedWrapperInvalidJsonTestCase = coretestcases.CaseV1{
	Title: "Invalid JSON bytes return error on deserialization",
	ArrangeInput: args.Map{
		"when":  "passing invalid json",
		"bytes": "not-valid-json{{{",
	},
	ExpectedInput: args.Map{
		"hasError": true,
	},
}

var typedWrapperDeserializeToManyTestCases = []coretestcases.CaseV1{
	{
		Title: "DeserializeToMany parses array of typed wrappers",
		ArrangeInput: args.Map{
			"count": 3,
		},
		ExpectedInput: args.Map{
			"count": 3, "title0": "item-0", "title1": "item-1", "title2": "item-2",
		},
	},
}

// ==========================================================================
// TypedPayloadWrapper — Metadata accessors
// ==========================================================================

var typedWrapperMetadataAccessorsTestCase = coretestcases.CaseV1{
	Title: "MetadataAccessors return correct values",
	ExpectedInput: args.Map{
		"name":             "meta-name",
		"identifier":       "meta-id",
		"categoryName":     "meta-category",
		"isParsed":         true,
		"isEmpty":          false,
		"hasError":         false,
		"hasSingleRecord":  true,
		"payloadsNonEmpty": true,
	},
}

// ==========================================================================
// TypedPayloadWrapper — TypedDataJson
// ==========================================================================

var typedWrapperTypedDataJsonTestCase = coretestcases.CaseV1{
	Title: "TypedDataJson returns non-empty results",
	ExpectedInput: args.Map{
		"dataJsonNonEmpty":  true,
		"dataJsonPtrNonNil": true,
		"jsonBytesNonEmpty": true,
		"jsonBytesNoError":  true,
	},
}
