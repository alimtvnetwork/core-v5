package isanytests

import (
	"reflect"

	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/issetter"
)

var (
	arrangeInterfaceArrayTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]interface{}{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	anyItemsToCsvStringSingleQuoteTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "left and right is null and both are equal.",
				ArrangeInput: []interface{}{
					coretests.DataHolder{
						First:  nil,
						Second: nil,
					},
					coretests.DataHolder{
						First:  1,
						Second: nil,
					},
					coretests.DataHolder{
						First:  1,
						Second: 2,
					},
					coretests.DataHolder{
						First:  &coretests.DraftType{},
						Second: nil,
					},
					coretests.DataHolder{
						First:  nil,
						Second: &coretests.DraftType{},
					},
					coretests.DataHolder{
						First:  &coretests.DraftType{},
						Second: &coretests.DraftType{},
					},
					coretests.DataHolder{
						First:  arrangeInterfaceArrayTypeVerification,
						Second: arrangeInterfaceArrayTypeVerification,
					},
				},
				ExpectedInput: []string{
					"0 : Equal (<nil>, <nil>)",
					"1 : NotEqual (int, <nil>)",
					"2 : Inconclusive (int, int)",
					"3 : NotEqual (*coretests.DraftType, <nil>)",
					"4 : NotEqual (<nil>, *coretests.DraftType)",
					"5 : Inconclusive (*coretests.DraftType, *coretests.DraftType)",
					"6 : Equal (*coretests.VerifyTypeOf, *coretests.VerifyTypeOf)",
				},
				VerifyTypeOf: arrangeInterfaceArrayTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}
)
