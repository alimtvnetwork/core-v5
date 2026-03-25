package coreuniquetests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var intUniqueGetRemovesDuplicatesTestCase = coretestcases.CaseV1{
	Title: "Get returns unique count -- slice with duplicates [1,2,2,3,3,3]",
	ArrangeInput: args.Map{
		"when":  "given slice with duplicates",
		"input": []int{1, 2, 2, 3, 3, 3},
	},
	ExpectedInput: "3",
}

var intUniqueGetAlreadyUniqueTestCase = coretestcases.CaseV1{
	Title: "Get returns same count -- already unique slice [1,2,3]",
	ArrangeInput: args.Map{
		"when":  "given slice without duplicates",
		"input": []int{1, 2, 3},
	},
	ExpectedInput: "3",
}

var intUniqueGetNilTestCase = coretestcases.CaseV1{
	Title: "Get returns nil-safe result -- nil slice input",
	ArrangeInput: args.Map{
		"when": "given nil slice",
	},
	ExpectedInput: "true",
}
