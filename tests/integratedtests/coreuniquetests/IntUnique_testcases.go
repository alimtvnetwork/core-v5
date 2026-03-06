package coreuniquetests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var intUniqueGetRemovesDuplicatesTestCase = coretestcases.CaseV1{
	Title: "Get removes duplicates",
	ArrangeInput: args.Map{
		"when":  "given slice with duplicates",
		"input": []int{1, 2, 2, 3, 3, 3},
	},
	ExpectedInput: "3",
}

var intUniqueGetAlreadyUniqueTestCase = coretestcases.CaseV1{
	Title: "Get returns same for already unique",
	ArrangeInput: args.Map{
		"when":  "given slice without duplicates",
		"input": []int{1, 2, 3},
	},
	ExpectedInput: "3",
}

var intUniqueGetNilTestCase = coretestcases.CaseV1{
	Title: "Get handles nil",
	ArrangeInput: args.Map{
		"when": "given nil slice",
	},
	ExpectedInput: "true",
}
