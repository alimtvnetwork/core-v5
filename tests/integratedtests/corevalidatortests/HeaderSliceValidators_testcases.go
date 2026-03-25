package corevalidatortests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// ==========================================
// Shared helpers
// ==========================================

func newMatchingHeaderSliceValidator() corevalidator.HeaderSliceValidator {
	return corevalidator.HeaderSliceValidator{
		Header: "test-header",
		SliceValidator: corevalidator.SliceValidator{
			Condition:     corevalidator.DefaultDisabledCoreCondition,
			CompareAs:     stringcompareas.Equal,
			ActualLines:   []string{"a", "b"},
			ExpectedLines: []string{"a", "b"},
		},
	}
}

func newMismatchHeaderSliceValidator() corevalidator.HeaderSliceValidator {
	return corevalidator.HeaderSliceValidator{
		Header: "mismatch-header",
		SliceValidator: corevalidator.SliceValidator{
			Condition:     corevalidator.DefaultDisabledCoreCondition,
			CompareAs:     stringcompareas.Equal,
			ActualLines:   []string{"x"},
			ExpectedLines: []string{"y"},
		},
	}
}

// ==========================================
// Length / IsEmpty
// ==========================================

var headerSliceValidatorsLengthTestCases = []coretestcases.CaseV1{
	{
		Title:         "HeaderSliceValidators.Length returns 0 -- nil input",
		ArrangeInput:  nil,
		ExpectedInput: args.Map{"length": 0},
	},
	{
		Title:         "HeaderSliceValidators.Length returns 0 -- empty slice",
		ArrangeInput:  corevalidator.HeaderSliceValidators{},
		ExpectedInput: args.Map{"length": 0},
	},
	{
		Title: "HeaderSliceValidators.Length returns 1 -- single item",
		ArrangeInput: corevalidator.HeaderSliceValidators{
			newMatchingHeaderSliceValidator(),
		},
		ExpectedInput: args.Map{"length": 1},
	},
	{
		Title: "HeaderSliceValidators.Length returns 2 -- two items",
		ArrangeInput: corevalidator.HeaderSliceValidators{
			newMatchingHeaderSliceValidator(),
			newMismatchHeaderSliceValidator(),
		},
		ExpectedInput: args.Map{"length": 2},
	},
}

var headerSliceValidatorsIsEmptyTestCases = []coretestcases.CaseV1{
	{
		Title:         "HeaderSliceValidators.IsEmpty returns true -- nil input",
		ArrangeInput:  nil,
		ExpectedInput: args.Map{"isEmpty": true},
	},
	{
		Title:         "HeaderSliceValidators.IsEmpty returns true -- empty slice",
		ArrangeInput:  corevalidator.HeaderSliceValidators{},
		ExpectedInput: args.Map{"isEmpty": true},
	},
	{
		Title: "HeaderSliceValidators.IsEmpty returns false -- non-empty slice",
		ArrangeInput: corevalidator.HeaderSliceValidators{
			newMatchingHeaderSliceValidator(),
		},
		ExpectedInput: args.Map{"isEmpty": false},
	},
}

// ==========================================
// IsMatch / IsValid
// ==========================================

var headerSliceValidatorsIsMatchTestCases = []coretestcases.CaseV1{
	{
		Title:         "HeaderSliceValidators.IsMatch returns true -- nil input",
		ArrangeInput:  nil,
		ExpectedInput: args.Map{"isMatch": true},
	},
	{
		Title:         "HeaderSliceValidators.IsMatch returns true -- empty slice",
		ArrangeInput:  corevalidator.HeaderSliceValidators{},
		ExpectedInput: args.Map{"isMatch": true},
	},
	{
		Title: "HeaderSliceValidators.IsMatch returns true -- all matching",
		ArrangeInput: corevalidator.HeaderSliceValidators{
			newMatchingHeaderSliceValidator(),
		},
		ExpectedInput: args.Map{"isMatch": true},
	},
	{
		Title: "HeaderSliceValidators.IsMatch returns false -- one mismatch",
		ArrangeInput: corevalidator.HeaderSliceValidators{
			newMatchingHeaderSliceValidator(),
			newMismatchHeaderSliceValidator(),
		},
		ExpectedInput: args.Map{"isMatch": false},
	},
}

// ==========================================
// VerifyAll
// ==========================================

var headerSliceValidatorsVerifyAllTestCases = []coretestcases.CaseV1{
	{
		Title:         "HeaderSliceValidators.VerifyAll returns nil -- empty slice",
		ArrangeInput:  corevalidator.HeaderSliceValidators{},
		ExpectedInput: args.Map{"hasError": false},
	},
	{
		Title: "HeaderSliceValidators.VerifyAll returns nil -- all matching",
		ArrangeInput: corevalidator.HeaderSliceValidators{
			newMatchingHeaderSliceValidator(),
		},
		ExpectedInput: args.Map{"hasError": false},
	},
	{
		Title: "HeaderSliceValidators.VerifyAll returns error -- mismatch found",
		ArrangeInput: corevalidator.HeaderSliceValidators{
			newMismatchHeaderSliceValidator(),
		},
		ExpectedInput: args.Map{"hasError": true},
	},
}

// ==========================================
// VerifyFirst
// ==========================================

var headerSliceValidatorsVerifyFirstTestCases = []coretestcases.CaseV1{
	{
		Title:         "HeaderSliceValidators.VerifyFirst returns nil -- empty slice",
		ArrangeInput:  corevalidator.HeaderSliceValidators{},
		ExpectedInput: args.Map{"hasError": false},
	},
	{
		Title: "HeaderSliceValidators.VerifyFirst returns nil -- matching item",
		ArrangeInput: corevalidator.HeaderSliceValidators{
			newMatchingHeaderSliceValidator(),
		},
		ExpectedInput: args.Map{"hasError": false},
	},
	{
		Title: "HeaderSliceValidators.VerifyFirst returns error -- mismatch found",
		ArrangeInput: corevalidator.HeaderSliceValidators{
			newMismatchHeaderSliceValidator(),
		},
		ExpectedInput: args.Map{"hasError": true},
	},
}

// ==========================================
// VerifyUpto
// ==========================================

var headerSliceValidatorsVerifyUptoTestCases = []coretestcases.CaseV1{
	{
		Title:         "HeaderSliceValidators.VerifyUpto returns nil -- empty slice",
		ArrangeInput:  corevalidator.HeaderSliceValidators{},
		ExpectedInput: args.Map{"hasError": false},
	},
	{
		Title: "HeaderSliceValidators.VerifyUpto returns nil -- matching within length",
		ArrangeInput: corevalidator.HeaderSliceValidators{
			newMatchingHeaderSliceValidator(),
		},
		ExpectedInput: args.Map{"hasError": false},
	},
	{
		Title: "HeaderSliceValidators.VerifyUpto returns error -- mismatch found",
		ArrangeInput: corevalidator.HeaderSliceValidators{
			newMismatchHeaderSliceValidator(),
		},
		ExpectedInput: args.Map{"hasError": true},
	},
}
