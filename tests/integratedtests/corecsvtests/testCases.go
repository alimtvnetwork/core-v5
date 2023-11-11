package corecsvtests

import (
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

	anyItemsToCsvStringDoubleQuoteTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given strings will be displayed as csv. " +
					"It will look like format: \"%s\", eg. \"%s\", \"%s\", \"%s\"...",
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
					"\"1\", " +
						"\"2\", " +
						"\"alim\", " +
						"\"created\", " +
						"\"{curly}\", " +
						"\"which wraps\", " +
						"\"\", " +
						"\"any string to\", " +
						"\"curly\", " +
						"\"even empty ones\", " +
						"\"and\", " +
						"\"{curly ones}\", " +
						"\"{left curly exists\", " +
						"\"right curly exists}\"",
				},
				VerifyTypeOf: arrangeInterfaceArrayTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	anyItemsToCsvStringNoQuoteTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given strings will be displayed as csv. " +
					"It will look like format: %s, eg. %s, %s, %s...",
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
					"1, 2, " +
						"alim, created, {curly}, which wraps, , " +
						"any string to, curly, even empty ones, " +
						"and, {curly ones}, " +
						"{left curly exists, right curly exists}",
				},
				VerifyTypeOf: arrangeInterfaceArrayTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	stringersToCsvStringSingleQuoteTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given strings will be displayed as csv. " +
					"On all true options, it will look like format: '%s', eg. '%s', '%s', '%s'...",
				ArrangeInput: []string{
					"1",
					"2",
					"alim",
					"created",
					"{curly}",
					"",
					"any string to",
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
						"'', " +
						"'any string to', " +
						"'and', " +
						"'{curly ones}', " +
						"'{left curly exists', " +
						"'right curly exists}'",
				},
				VerifyTypeOf: defaultTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	stringersToCsvStringDoubleQuoteTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given strings will be displayed as csv. " +
					"Double quote format: \"%s\", eg. \"%s\", \"%s\", \"%s\"...",
				ArrangeInput: []string{
					"1",
					"2",
					"alim",
					"created",
					"{curly}",
					"",
					"any string to",
					"and",
					"{curly ones}",
					"{left curly exists",
					"right curly exists}",
				},
				ExpectedInput: []string{
					"\"1\", \"2\", \"alim\", \"created\", \"{curly}\", \"\", \"any string to\", \"and\", \"{curly ones}\", \"{left curly exists\", \"right curly exists}\"",
				},
				VerifyTypeOf: defaultTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	stringersToCsvStringNoQuoteTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given strings will be displayed as csv. " +
					"No quote format: %s, eg. %s, %s, %s...",
				ArrangeInput: []string{
					"1",
					"2",
					"alim",
					"created",
					"{curly}",
					"",
					"any string to",
					"and",
					"{curly ones}",
					"{left curly exists",
					"right curly exists}",
				},
				ExpectedInput: []string{
					"1, 2, alim, created, {curly}, , any string to, and, {curly ones}, {left curly exists, right curly exists}",
				},
				VerifyTypeOf: defaultTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}
)
