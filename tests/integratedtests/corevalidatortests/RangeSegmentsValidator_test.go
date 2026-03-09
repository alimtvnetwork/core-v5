package corevalidatortests

import (
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/corevalidator"
)

// ==========================================
// LengthOfVerifierSegments
// ==========================================

func Test_RangeSegmentsValidator_LengthOfVerifierSegments(t *testing.T) {
	for caseIndex, tc := range rangeSegmentsValidatorLengthTestCases {
		// Arrange
		v := tc.ArrangeInput.(*corevalidator.RangeSegmentsValidator)

		// Act
		actual := args.Map{
			"length": v.LengthOfVerifierSegments(),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Validators
// ==========================================

func Test_RangeSegmentsValidator_Validators(t *testing.T) {
	for caseIndex, tc := range rangeSegmentsValidatorValidatorsTestCases {
		// Arrange
		v := tc.ArrangeInput.(*corevalidator.RangeSegmentsValidator)
		v.SetActual(rangeSegActualLines)

		// Act
		validators := v.Validators()
		actual := args.Map{
			"hasValidators": len(validators) > 0,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// VerifyAll
// ==========================================

func Test_RangeSegmentsValidator_VerifyAll(t *testing.T) {
	for caseIndex, tc := range rangeSegmentsValidatorVerifyAllTestCases {
		// Arrange
		v := tc.ArrangeInput.(*corevalidator.RangeSegmentsValidator)
		params := &corevalidator.Parameter{
			CaseIndex:       0,
			Header:          "test",
			IsCaseSensitive: true,
		}

		// Act
		err := v.VerifyAll("header", rangeSegActualLines, params, false)
		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// VerifySimple
// ==========================================

func Test_RangeSegmentsValidator_VerifySimple(t *testing.T) {
	for caseIndex, tc := range rangeSegmentsValidatorVerifySimpleTestCases {
		// Arrange
		v := tc.ArrangeInput.(*corevalidator.RangeSegmentsValidator)
		params := &corevalidator.Parameter{
			CaseIndex:       0,
			Header:          "test",
			IsCaseSensitive: true,
		}

		// Act
		err := v.VerifySimple(rangeSegActualLines, params, false)
		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// VerifyFirst
// ==========================================

func Test_RangeSegmentsValidator_VerifyFirst(t *testing.T) {
	for caseIndex, tc := range rangeSegmentsValidatorVerifyFirstTestCases {
		// Arrange
		v := tc.ArrangeInput.(*corevalidator.RangeSegmentsValidator)
		params := &corevalidator.Parameter{
			CaseIndex:       0,
			Header:          "test",
			IsCaseSensitive: true,
		}

		// Act
		err := v.VerifyFirst("header", rangeSegActualLines, params, false)
		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// VerifyUpto
// ==========================================

func Test_RangeSegmentsValidator_VerifyUpto(t *testing.T) {
	for caseIndex, tc := range rangeSegmentsValidatorVerifyUptoTestCases {
		// Arrange
		v := tc.ArrangeInput.(*corevalidator.RangeSegmentsValidator)
		params := &corevalidator.Parameter{
			CaseIndex:       0,
			Header:          "test",
			IsCaseSensitive: true,
		}

		// Act
		err := v.VerifyUpto("header", rangeSegActualLines, params, 3, false)
		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// VerifyFirstDefault
// ==========================================

func Test_RangeSegmentsValidator_VerifyFirstDefault(t *testing.T) {
	for caseIndex, tc := range rangeSegmentsValidatorVerifyFirstDefaultTestCases {
		// Arrange
		v := tc.ArrangeInput.(*corevalidator.RangeSegmentsValidator)
		params := &corevalidator.Parameter{
			CaseIndex:       0,
			Header:          "test",
			IsCaseSensitive: true,
		}

		// Act
		err := v.VerifyFirstDefault(rangeSegActualLines, params, false)
		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// VerifyUptoDefault
// ==========================================

func Test_RangeSegmentsValidator_VerifyUptoDefault(t *testing.T) {
	for caseIndex, tc := range rangeSegmentsValidatorVerifyUptoDefaultTestCases {
		// Arrange
		v := tc.ArrangeInput.(*corevalidator.RangeSegmentsValidator)
		params := &corevalidator.Parameter{
			CaseIndex:       0,
			Header:          "test",
			IsCaseSensitive: true,
		}

		// Act
		err := v.VerifyUptoDefault(rangeSegActualLines, params, 3, false)
		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
