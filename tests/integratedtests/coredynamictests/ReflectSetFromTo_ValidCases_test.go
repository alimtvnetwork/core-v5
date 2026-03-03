package coredynamictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/tests/testwrappers/coredynamictestwrappers"
)

// Test_ReflectSetFromTo_ValidCases
//
// Valid Inputs:
//   - From, To: (null, null)                          -- do nothing
//   - From, To: (sameTypePointer, sameTypePointer)    -- try reflection
//   - From, To: (sameTypeNonPointer, sameTypePointer) -- try reflection
//   - From, To: ([]byte or *[]byte, otherType)        -- try unmarshal, reflect
//   - From, To: (otherType, *[]byte)                  -- try marshal, reflect
func Test_ReflectSetFromTo_ValidCases(t *testing.T) {
	for caseIndex, testCase := range coredynamictestwrappers.ReflectSetFromToValidTestCases {
		// Act
		err := coredynamic.ReflectSetFromTo(
			testCase.From,
			testCase.To,
		)

		typeStatus := coredynamic.TypeSameStatus(
			testCase.To,
			testCase.ExpectedValue,
		)
		testCase.SetActual(testCase.To)

		// Assert
		actLines := []string{
			fmt.Sprintf("%v", err == nil),
			fmt.Sprintf("%v", typeStatus.IsSame()),
		}
		expected := []string{"true", "true"}

		switch convertedFrom := testCase.From.(type) {
		case *coretests.DraftType:
			toField := testCase.ToFieldToDraftType()
			expectedField := testCase.ExpectedFieldToDraftType()
			toFieldEqualErr := toField.
				VerifyNotEqualExcludingInnerFieldsErr(expectedField)
			fromFieldEqualErr := convertedFrom.
				VerifyNotEqualExcludingInnerFieldsErr(expectedField)

			actLines = append(actLines,
				fmt.Sprintf("%v", toFieldEqualErr == nil),
				fmt.Sprintf("%v", fromFieldEqualErr == nil),
			)
			expected = append(expected, "true", "true")

		case coretests.DraftType:
			toField := testCase.ToFieldToDraftType()
			expectedField := testCase.ExpectedFieldToDraftType()
			toFieldEqualErr := toField.
				VerifyNotEqualExcludingInnerFieldsErr(expectedField)
			fromFieldEqualErr := convertedFrom.
				VerifyNotEqualExcludingInnerFieldsErr(expectedField)

			actLines = append(actLines,
				fmt.Sprintf("%v", toFieldEqualErr == nil),
				fmt.Sprintf("%v", fromFieldEqualErr == nil),
			)
			expected = append(expected, "true", "true")

		case []byte, *[]byte:
			toField := testCase.ToFieldToDraftType()
			toFieldEqualErr := toField.
				VerifyNotEqualExcludingInnerFieldsErr(
					testCase.ExpectedFieldToDraftType(),
				)
			actLines = append(actLines,
				fmt.Sprintf("%v", toFieldEqualErr == nil),
			)
			expected = append(expected, "true")
		}

		errcore.PrintLineDiff(caseIndex, testCase.Header, actLines, expected)

		for i, act := range actLines {
			if act != expected[i] {
				t.Errorf("[case %d] %s: line %d got %q, want %q",
					caseIndex, testCase.Header, i, act, expected[i])
			}
		}
	}
}
