package anycmptests

import (
	"fmt"
	"reflect"

	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/issetter"
)

var (
	defaultTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]string{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}
	arrangeInterfaceArrayTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]interface{}{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	arrangeFmtStringerTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]fmt.Stringer{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	anyItemsToCsvStringSingleQuoteTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given strings will be displayed as csv. " +
					"On all true options, it will look like format: '%s', eg. '%s', '%s', '%s'...",
				ArrangeInput: []interface{}{
					1,
					2,
					"alim",
					"created",
					"{curly}",
					"which wraps",
					"",
					"any string to",
					"curly",
					"even empty ones",
					"and",
					"{curly ones}",
					"{left curly exists",
					"right curly exists}",
				},
				ExpectedInput: []string{
					"'1', " +
						"'2', " +
						"'alim', " +
						"'created', " +
						"'{curly}', " +
						"'which wraps', " +
						"'', " +
						"'any string to', " +
						"'curly', " +
						"'even empty ones', " +
						"'and', '{curly ones}', " +
						"'{left curly exists', " +
						"'right curly exists}'",
				},
				VerifyTypeOf: arrangeInterfaceArrayTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}
)
