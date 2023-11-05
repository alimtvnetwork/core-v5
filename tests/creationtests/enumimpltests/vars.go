package enumimpltests

import (
	"reflect"

	"gitlab.com/auk-go/core/coreimpl/enumimpl"
	"gitlab.com/auk-go/core/coretests"
)

var (
	checker1                                     = enumimpl.LeftRightDiffCheckerImpl
	typeVerifyOfForDynamicMapSimpleDiffTestCases = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf(LeftRightDynamicMapWithDefaultChecker{}),
		ActualInput:   reflect.TypeOf(""),
		ExpectedInput: reflect.TypeOf(""),
	}
)
