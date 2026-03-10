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
		Title:         "Nil returns length 0",
		ArrangeInput:  nil,
		ExpectedInput: args.Map{"length": 0},
	},
	{
		Title:         "Empty slice returns length 0",
		ArrangeInput:  corevalidator.HeaderSliceValidators{},
		ExpectedInput: args.Map{"length": 0},
	},
	{
		Title: "Single item returns length 1",
		ArrangeInput: corevalidator.HeaderSliceValidators{
			newMatchingHeaderSliceValidator(),
		},
		ExpectedInput: args.Map{"length": 1},
	},
	{
		Title: "Two items returns length 2",
		ArrangeInput: corevalidator.HeaderSliceValidators{
			newMatchingHeaderSliceValidator(),
			newMismatchHeaderSliceValidator(),
		},
		ExpectedInput: args.Map{"length": 2},
	},
}

var headerSliceValidatorsIsEmptyTestCases = []coretestcases.CaseV1{
	{
		Title:         "Nil is empty",
		ArrangeInput:  nil,
		ExpectedInput: args.Map{"isEmpty": true},
	},
	{
		Title:         "Empty slice is empty",
		ArrangeInput:  corevalidator.HeaderSliceValidators{},
		ExpectedInput: args.Map{"isEmpty": true},
	},
	{
		Title: "Non-empty slice is not empty",
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
		Title:         "Nil matches",
		ArrangeInput:  nil,
		ExpectedInput: args.Map{"isMatch": true},
	},
	{
		Title:         "Empty matches",
		ArrangeInput:  corevalidator.HeaderSliceValidators{},
		ExpectedInput: args.Map{"isMatch": true},
	},
	{
		Title: "All matching returns true",
		ArrangeInput: corevalidator.HeaderSliceValidators{
			newMatchingHeaderSliceValidator(),
		},
		ExpectedInput: args.Map{"isMatch": true},
	},
	{
		Title: "One mismatch returns false",
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
		Title:         "Empty returns nil error",
		ArrangeInput:  corevalidator.HeaderSliceValidators{},
		ExpectedInput: args.Map{"hasError": false},
	},
	{
		Title: "All matching returns nil error",
		ArrangeInput: corevalidator.HeaderSliceValidators{
			newMatchingHeaderSliceValidator(),
		},
		ExpectedInput: args.Map{"hasError": false},
	},
	{
		Title: "Mismatch returns error",
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
		Title:         "Empty returns nil error",
		ArrangeInput:  corevalidator.HeaderSliceValidators{},
		ExpectedInput: args.Map{"hasError": false},
	},
	{
		Title: "Matching returns nil error",
		ArrangeInput: corevalidator.HeaderSliceValidators{
			newMatchingHeaderSliceValidator(),
		},
		ExpectedInput: args.Map{"hasError": false},
	},
	{
		Title: "Mismatch returns error",
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
		Title:         "Empty returns nil error",
		ArrangeInput:  corevalidator.HeaderSliceValidators{},
		ExpectedInput: args.Map{"hasError": false},
	},
	{
		Title: "Matching within length returns nil error",
		ArrangeInput: corevalidator.HeaderSliceValidators{
			newMatchingHeaderSliceValidator(),
		},
		ExpectedInput: args.Map{"hasError": false},
	},
	{
		Title: "Mismatch returns error",
		ArrangeInput: corevalidator.HeaderSliceValidators{
			newMismatchHeaderSliceValidator(),
		},
		ExpectedInput: args.Map{"hasError": true},
	},
}
