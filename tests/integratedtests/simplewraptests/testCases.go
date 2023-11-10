package simplewraptests

import (
	"reflect"

	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/issetter"
)

var (
	stringsSliceTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]string{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	curlyWrapIfEnabledValidTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given strings will be curly wrapped without checking. " +
					"Curly wrapped guaranteed, duplicate curly wrap is possible.",
				ArrangeInput: []string{
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
					"{alim}",
					"{created}",
					"{{curly}}",
					"{which wraps}",
					"{}",
					"{any string to}",
					"{curly}",
					"{even empty ones}",
					"{and}",
					"{{curly ones}}",
					"{{left curly exists}",
					"{right curly exists}}",
				},
				VerifyTypeOf: stringsSliceTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}
	curlyWrapIfDisabledValidTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given strings will be curly wrapped without checking. " +
					"Curly wrapped guaranteed, duplicate curly wrap is possible.",
				ArrangeInput: []string{
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
				VerifyTypeOf: stringsSliceTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	curlyWrapOptionsValidTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given strings will be curly wrapped with checking. " +
					"Curly wrapped guaranteed, no duplicate curly wrap possible.",
				ArrangeInput: []string{
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
					"{alim}",
					"{created}",
					"{curly}",
					"{which wraps}",
					"{}",
					"{any string to}",
					"{curly}",
					"{even empty ones}",
					"{and}",
					"{curly ones}",
					"{left curly exists}",
					"{right curly exists}",
				},
				VerifyTypeOf: stringsSliceTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}
	parenthesisValidTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given strings will be parenthesis ( wrapped ) with no checking. " +
					"Parenthesis wrapped guaranteed, duplicate wrap possible.",
				ArrangeInput: []string{
					"alim",
					"created",
					"(parenthesis)",
					"which wraps",
					"",
					"any string to",
					"parenthesis",
					"even empty ones",
					"and",
					"(parenthesis ones)",
					"(left parenthesis exists",
					"right parenthesis exists)",
				},
				ExpectedInput: []string{
					"(alim)",
					"(created)",
					"((parenthesis))",
					"(which wraps)",
					"()",
					"(any string to)",
					"(parenthesis)",
					"(even empty ones)",
					"(and)",
					"((parenthesis ones))",
					"((left parenthesis exists)",
					"(right parenthesis exists))",
				},
				VerifyTypeOf: stringsSliceTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	parenthesisDisabledRemainsAsItIsTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given strings will be parenthesis ( wrapped ) with no checking. " +
					"Parenthesis wrapped guaranteed, duplicate wrap possible.",
				ArrangeInput: []string{
					"alim",
					"created",
					"(parenthesis)",
					"which wraps",
					"",
					"any string to",
					"parenthesis",
					"even empty ones",
					"and",
					"(parenthesis ones)",
					"(left parenthesis exists",
					"right parenthesis exists)",
				},
				ExpectedInput: []string{
					"alim",
					"created",
					"(parenthesis)",
					"which wraps",
					"",
					"any string to",
					"parenthesis",
					"even empty ones",
					"and",
					"(parenthesis ones)",
					"(left parenthesis exists",
					"right parenthesis exists)",
				},
				VerifyTypeOf: stringsSliceTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	squareBracketWrapTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given strings will be square bracket [ wrapped ] with no checking. " +
					"Square bracket wrapped guaranteed, duplicate wrap possible.",
				ArrangeInput: []string{
					"alim",
					"created",
					"[sq bracket]",
					"which wraps",
					"",
					"any string to",
					"parenthesis",
					"even empty ones",
					"and",
					"[square]",
					"[left sq exists",
					"right sq exists]",
				},
				ExpectedInput: []string{
					"[alim]",
					"[created]",
					"[[sq bracket]]",
					"[which wraps]",
					"[]",
					"[any string to]",
					"[parenthesis]",
					"[even empty ones]",
					"[and]",
					"[[square]]",
					"[[left sq exists]",
					"[right sq exists]]",
				},
				VerifyTypeOf: stringsSliceTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	squareBracketWrapDisabledTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given strings will be NOT square bracket [ wrapped ]. " +
					"Square bracket wrapped is NOT guaranteed.",
				ArrangeInput: []string{
					"alim",
					"created",
					"[sq bracket]",
					"which wraps",
					"",
					"any string to",
					"parenthesis",
					"even empty ones",
					"and",
					"[square]",
					"[left sq exists",
					"right sq exists]",
				},
				ExpectedInput: []string{
					"alim",
					"created",
					"[sq bracket]",
					"which wraps",
					"",
					"any string to",
					"parenthesis",
					"even empty ones",
					"and",
					"[square]",
					"[left sq exists",
					"right sq exists]",
				},
				VerifyTypeOf: stringsSliceTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	titleCurlyMetaTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Title curly meta should look like : " +
					"Square bracket wrapped is NOT guaranteed.",
				ArrangeInput: []string{
					"my title",       // title
					"some \"value\"", // value
					corejson.Serialize.ToString(coretests.DraftType{
						// meta
						SampleString1: "Some meta information",
					}),
				},
				ExpectedInput: []string{
					"alim",
					"created",
					"[sq bracket]",
					"which wraps",
					"",
					"any string to",
					"parenthesis",
					"even empty ones",
					"and",
					"[square]",
					"[left sq exists",
					"right sq exists]",
				},
				VerifyTypeOf: stringsSliceTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}
)
