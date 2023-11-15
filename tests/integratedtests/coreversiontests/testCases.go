package coreversiontests

import (
	"reflect"

	"gitlab.com/auk-go/core/corecomparator"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coreversion"
	"gitlab.com/auk-go/core/issetter"
)

var (
	arrangeTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]interface{}{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	arrangeStringTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]string{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	defaultVersionTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]*coreversion.Version{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	versionCreationTestCases = []testWrapper{
		{
			Title: "Create versions with different args and methods.",
			ArrangeInput: []coreversion.Version{
				coreversion.New.Invalid(),
				coreversion.New.Default("1.2.3.4"),
				coreversion.New.Default("5.3.6"),
				coreversion.New.Default("5.3"),
				coreversion.New.Default("9"),
				coreversion.New.Default("v1.2.3.4"),
				coreversion.New.Default("v5.3.6"),
				coreversion.New.Default("v5.3"),
				coreversion.New.Default("v9"),
				coreversion.New.Default(""),
			},
			ExpectedInput: []string{
				"0 : invalid - ",
				"1 : v1.2.3.4 (compact: 1.2.3.4, display: v1.2.3.4)",
				"2 : v5.3.6 (compact: 5.3.6, display: v5.3.6)",
				"3 : v5.3 (compact: 5.3, display: v5.3)",
				"4 : v9 (compact: 9, display: v9)",
				"5 : v1.2.3.4 (compact: 1.2.3.4, display: v1.2.3.4)",
				"6 : v5.3.6 (compact: 5.3.6, display: v5.3.6)",
				"7 : v5.3 (compact: 5.3, display: v5.3)",
				"8 : v9 (compact: 9, display: v9)",
				"9 : invalid - ",
			},
			VerifyTypeOf: defaultVersionTypeVerification,
			IsEnable:     issetter.True,
		},
	}

	versionCreationUsingStringTestCases = []testWrapper{
		{
			Title: "Create versions using string.",
			ArrangeInput: []string{
				"-1",
				"1.2.3.4",
				"5.3.6",
				"5.3",
				"9",
				"v1.2.3.4",
				"v5.3.6",
				"v5.3",
				"v9",
				"5.*.1",
				"7.*.*",
				"1.*.*.10",
				"-1.555.*.11",
				"v-1.576.*.12",
				"8v-1.581.*.13",
				"8v-1.x565.*.u14",
				"8v-1.*.*.u15",
				"v5.-5",
				"v10.-6",
				"v11...7",
				"v12...8",
				"xv12...9",
				"12..5",
				"13..6",
				"14..7",
				"",
			},
			ExpectedInput: []string{
				"0 : invalid - v-1",
				"1 : v1.2.3.4 (compact: 1.2.3.4, display: v1.2.3.4)",
				"2 : v5.3.6 (compact: 5.3.6, display: v5.3.6)",
				"3 : v5.3 (compact: 5.3, display: v5.3)",
				"4 : v9 (compact: 9, display: v9)",
				"5 : v1.2.3.4 (compact: 1.2.3.4, display: v1.2.3.4)",
				"6 : v5.3.6 (compact: 5.3.6, display: v5.3.6)",
				"7 : v5.3 (compact: 5.3, display: v5.3)",
				"8 : v9 (compact: 9, display: v9)",
				"9 : v5.0.1 (compact: 5.*.1, display: v5.*.1)",
				"10 : v7 (compact: 7.*.*, display: v7.*.*)",
				"11 : v1.0.0.10 (compact: 1.*.*.10, display: v1.*.*.10)",
				"12 : invalid - v-1.555.0.11",
				"13 : invalid - v-1.576.0.12",
				"14 : invalid - v-1.581.0.13",
				"15 : invalid - v8v-1.x565.*.u14",
				"16 : invalid - v8v-1.*.*.u15",
				"17 : invalid - v5",
				"18 : invalid - v10",
				"19 : v11.0.0.7 (compact: 11...7, display: v11...7)",
				"20 : v12.0.0.8 (compact: 12...8, display: v12...8)",
				"21 : invalid - v-1.0.0.9",
				"22 : v12.0.5 (compact: 12..5, display: v12..5)",
				"23 : v13.0.6 (compact: 13..6, display: v13..6)",
				"24 : v14.0.7 (compact: 14..7, display: v14..7)",
				"25 : invalid - ",
			},
			VerifyTypeOf: arrangeStringTypeVerification,
			IsEnable:     issetter.True,
		},
	}

	comparisonStringTestCases = []testWrapper{
		{
			Title: "Versions comparisons",
			ArrangeInput: []coretests.LeftRightExpect{
				{
					Left:   "1.2.5",
					Right:  "1.2.4",
					Expect: corecomparator.LeftGreater,
				},
				{
					Left:   "1.5.5",
					Right:  "1.*.8",
					Expect: corecomparator.LeftGreater,
				},
				{
					Left:   "1.2",
					Right:  "1.2.1",
					Expect: corecomparator.LeftLess,
				},
				{
					Left:   "1.2",
					Right:  "1.2.1",
					Expect: corecomparator.LeftLess,
				},
			},
			ExpectedInput: []string{
				"0 : Expect : LeftGreater - Left [v1.2.5] - [v1.2.4] Right , match: true",
				"1 : Expect : LeftGreater - Left [v1.5.5] - [v1.0.8] Right , match: true",
			},
			VerifyTypeOf: arrangeStringTypeVerification,
			IsEnable:     issetter.True,
		},
	}
)
