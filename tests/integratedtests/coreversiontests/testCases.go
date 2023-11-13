package coreversiontests

import (
	"reflect"

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

	arrangeArgsTwoTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]coretests.ArgTwo{}),
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
			ArrangeInput: []*coreversion.Version{
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
				"0 : invalid (empty)",
				"1 : v1.2.3.4 (compact: 1.2.3.4, display: v1.2.3.4)",
				"2 : v5.3.6 (compact: 5.3.6, display: v5.3.6)",
				"3 : v5.3 (compact: 5.3, display: v5.3)",
				"4 : v9 (compact: 9, display: v9)",
				"5 : v1.2.3.4 (compact: 1.2.3.4, display: v1.2.3.4)",
				"6 : v5.3.6 (compact: 5.3.6, display: v5.3.6)",
				"7 : v5.3 (compact: 5.3, display: v5.3)",
				"8 : v9 (compact: 9, display: v9)",
				"9 : invalid (empty)",
			},
			VerifyTypeOf: defaultVersionTypeVerification,
			IsEnable:     issetter.True,
		},
	}
)
