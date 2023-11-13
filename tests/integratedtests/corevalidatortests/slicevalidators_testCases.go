package corevalidatortests

import (
	"reflect"

	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/issetter"
	"gitlab.com/auk-go/core/tests/testwrappers/corevalidatortestwrappers"
)

var (
	arrangeArgsTwoTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]coretests.ArgTwo{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	sliceValidatorsTestCases = []corevalidatortestwrappers.SliceValidatorsWrapper{
		{
			Case: coretestcases.CaseV1{
				Title: "Diff check against invalid comparisons",
				ArrangeInput: []coretests.ArgTwo{
					{
						First:  1,
						Second: byte(2),
					},
					{
						First:  1,
						Second: float64(5),
					},
					{
						First:  "1",
						Second: 1,
					},
				},
				ExpectedInput: []string{
					"Wrong expectation",
					"Wrong expectation",
					"Wrong expectation",
				},
				VerifyTypeOf: arrangeArgsTwoTypeVerification,
				IsEnable:     issetter.True,
			},
			Validator: corevalidator.SliceValidator{
				ValidatorCoreCondition: corevalidator.DefaultTrimCoreCondition,
				ExpectedLines: []string{
					"Wrong expectation",
				},
			},
		},
	}
)
