package isanytests

import (
	"testing"

	"gitlab.com/auk-go/core/corecsv"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/isany"
)

// Test_Extended_Defined_TypedNil verifies isany.Defined with typed-nil error and *int.
// Migrated from cmd/main/nullTesting01.go.
func Test_Extended_Defined_TypedNil(t *testing.T) {
	for caseIndex, testCase := range extendedDefinedTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]any)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		for i, input := range inputs {
			actualSlice.AppendFmt(
				booleanPrintFormatWithType,
				i,
				isany.Defined(input),
				input,
				input,
			)
		}

		finalActLines := actualSlice.Strings()
		finalTestCase := coretestcases.
			CaseV1(testCase.BaseTestCase)

		// Assert
		finalTestCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}

// Test_Extended_Null_TypedNil verifies isany.Null with typed-nil error and *int.
// Migrated from cmd/main/nullTesting01.go.
func Test_Extended_Null_TypedNil(t *testing.T) {
	for caseIndex, testCase := range extendedNullTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]any)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		for i, input := range inputs {
			actualSlice.AppendFmt(
				booleanPrintFormatWithType,
				i,
				isany.Null(input),
				input,
				input,
			)
		}

		finalActLines := actualSlice.Strings()
		finalTestCase := coretestcases.
			CaseV1(testCase.BaseTestCase)

		// Assert
		finalTestCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}

// Test_Extended_DefinedBoth_TypedNil verifies isany.DefinedBoth with error and *int typed nils.
// Migrated from cmd/main/nullTesting02.go.
func Test_Extended_DefinedBoth_TypedNil(t *testing.T) {
	for caseIndex, testCase := range extendedDefinedBothTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]args.Two)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		for i, parameter := range inputs {
			f := parameter.First
			s := parameter.Second

			actualSlice.AppendFmt(
				defaultCaseIndexBoolStringFmt,
				i,
				isany.DefinedBoth(f, s),
				corecsv.AnyToTypesCsvDefault(f, s),
			)
		}

		finalActLines := actualSlice.Strings()
		finalTestCase := coretestcases.
			CaseV1(testCase.BaseTestCase)

		// Assert
		finalTestCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}

// Test_Extended_NullBoth_TypedNil verifies isany.NullBoth with error and *int typed nils.
// Migrated from cmd/main/nullTesting02.go.
func Test_Extended_NullBoth_TypedNil(t *testing.T) {
	for caseIndex, testCase := range extendedNullBothTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]args.Two)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		for i, parameter := range inputs {
			f := parameter.First
			s := parameter.Second

			actualSlice.AppendFmt(
				defaultCaseIndexBoolStringFmt,
				i,
				isany.NullBoth(f, s),
				corecsv.AnyToTypesCsvDefault(f, s),
			)
		}

		finalActLines := actualSlice.Strings()
		finalTestCase := coretestcases.
			CaseV1(testCase.BaseTestCase)

		// Assert
		finalTestCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
