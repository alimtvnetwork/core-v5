package corevalidatortests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var simpleSliceValidatorSetActualTestCase = coretestcases.CaseV1{
	Title: "SimpleSliceValidator SetActual returns same instance",
	ExpectedInput: args.Map{
		"sameInstance": true,
	},
}

var simpleSliceValidatorSliceValidatorTestCase = coretestcases.CaseV1{
	Title: "SimpleSliceValidator SliceValidator returns non-nil",
	ExpectedInput: args.Map{
		"isNotNil": true,
	},
}

var simpleSliceValidatorVerifyAllTestCases = []coretestcases.CaseV1{
	{
		Title: "SimpleSliceValidator VerifyAll matching returns nil",
		ArrangeInput: args.Map{
			"expected": []string{"a", "b"},
			"actual":   []string{"a", "b"},
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
	{
		Title: "SimpleSliceValidator VerifyAll mismatch returns error",
		ArrangeInput: args.Map{
			"expected": []string{"a", "b"},
			"actual":   []string{"x", "y"},
		},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
}

var simpleSliceValidatorVerifyFirstTestCase = coretestcases.CaseV1{
	Title: "SimpleSliceValidator VerifyFirst matching returns nil",
	ArrangeInput: args.Map{
		"expected": []string{"a"},
		"actual":   []string{"a"},
	},
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

var simpleSliceValidatorVerifyUptoTestCase = coretestcases.CaseV1{
	Title: "SimpleSliceValidator VerifyUpto matching returns nil",
	ArrangeInput: args.Map{
		"expected": []string{"a", "b", "c"},
		"actual":   []string{"a", "b", "c"},
		"length":   2,
	},
	ExpectedInput: args.Map{
		"hasError": false,
	},
}
