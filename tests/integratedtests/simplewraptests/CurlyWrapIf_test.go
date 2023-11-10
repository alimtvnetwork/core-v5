package simplewraptests

import (
	"testing"
	
	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/simplewrap"
)

func Test_CurlyWrapIf_Enabled_Wraps_All_Dont_CheckConditionally(t *testing.T) {
	// Arrange
	sliceValidator := corevalidator.SliceValidator{
		ValidatorCoreCondition: corevalidator.DefaultTrimCoreCondition,
	}
	
	for caseIndex, testCase := range curlyWrapIfEnabledValidTestCases {
		inputs := testCase.Arrange()
		actualSlice := corestr.New.SimpleSlice.Cap(len(inputs))
		
		for _, input := range inputs {
			actualSlice.Add(simplewrap.CurlyWrapIf(true, input))
		}
		
		finalActual := actualSlice.Strings()
		testCase.SetActual(finalActual)
		sliceValidator.SetActual(finalActual)
		sliceValidator.ExpectedLines = testCase.ExpectedInput.([]string)
		
		nextBaseParam := corevalidator.ValidatorParamsBase{
			CaseIndex:          caseIndex,
			Header:             testCase.Title,
			IsAttachUserInputs: true,
			IsCaseSensitive:    true,
		}
		
		// Act
		validationFinalError := sliceValidator.AllVerifyError(
			&nextBaseParam)
		
		// Assert
		convey.Convey(testCase.Title, t, func() {
			errcore.PrintErrorWithTestIndex(
				caseIndex,
				testCase.Title,
				validationFinalError)
			
			convey.So(
				validationFinalError,
				should.BeNil)
		})
		
		convey.Convey(testCase.Title+" - type verify", t, func() {
			convey.So(
				testCase.TypeValidationError(),
				should.BeNil)
		})
	}
}

func Test_CurlyWrapIf_Disabled_Wraps_All_Dont_CheckConditionally(t *testing.T) {
	// Arrange
	sliceValidator := corevalidator.SliceValidator{
		ValidatorCoreCondition: corevalidator.DefaultTrimCoreCondition,
	}
	
	for caseIndex, testCase := range curlyWrapIfDisabledValidTestCases {
		inputs := testCase.Arrange()
		actualSlice := corestr.New.SimpleSlice.Cap(len(inputs))
		
		for _, input := range inputs {
			actualSlice.Add(simplewrap.CurlyWrapIf(false, input))
		}
		
		finalActual := actualSlice.Strings()
		testCase.SetActual(finalActual)
		sliceValidator.SetActual(finalActual)
		sliceValidator.ExpectedLines = testCase.ExpectedInput.([]string)
		
		nextBaseParam := corevalidator.ValidatorParamsBase{
			CaseIndex:          caseIndex,
			Header:             testCase.Title,
			IsAttachUserInputs: true,
			IsCaseSensitive:    true,
		}
		
		// Act
		validationFinalError := sliceValidator.AllVerifyError(
			&nextBaseParam)
		
		// Assert
		convey.Convey(testCase.Title, t, func() {
			errcore.PrintErrorWithTestIndex(
				caseIndex,
				testCase.Title,
				validationFinalError)
			
			convey.So(
				validationFinalError,
				should.BeNil)
		})
		
		convey.Convey(testCase.Title+" - type verify", t, func() {
			convey.So(
				testCase.TypeValidationError(),
				should.BeNil)
		})
	}
}

// func Test_When_DoubleQuoteWrapElements_SkipQuoteOnPresent_Disabled_Should_Have_DuplicateDoubleQuotations(t *testing.T) {
// 	testCases := []string{
// 		"some-elem",
// 		"alim-elem",
// 		"\"has-quote\"",
// 		"",
// 		"\"",
// 		"\"first",
// 		"last\"",
// 		"'",
// 		"simple",
// 	}
// 	expectation := []string{
// 		"\"some-elem\"",
// 		"\"alim-elem\"",
// 		"\"\"has-quote\"\"",
// 		"\"\"",
// 		"\"\"\"",
// 		"\"\"first\"",
// 		"\"last\"\"",
// 		"\"'\"",
// 		"\"simple\"",
// 	}
// 	
// 	// Act
// 	actual := simplewrap.
// 		DoubleQuoteWrapElements(
// 			false,
// 			testCases...)
// 	
// 	// Assert
// 	convey.Convey("Wrap strings with double quote, if exists already then skip adding", t, func() {
// 		convey.So(actual, should.Equal, expectation)
// 	})
// }
