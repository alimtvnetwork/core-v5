package corerangestests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/tests/testwrappers/coredynamictestwrappers"
)

func Test_IntRanges_ValidCases(t *testing.T) {
	for _, testCase := range coredynamictestwrappers.ReflectSetFromToInvalidTestCases {
		// Act
		err := coredynamic.ReflectSetFromTo(
			testCase.From,
			testCase.To)

		typeStatus := coredynamic.TypeSameStatus(
			testCase.To,
			testCase.ExpectedValue)
		testCase.SetActual(testCase.To)

		// Assert
		Convey(testCase.CaseTitle(), t, func() {
			So(err, ShouldBeNil)
			typeStatus.MustBeSame()
			switch convertedFrom := testCase.From.(type) {
			case *coretests.DraftType:
				toField := testCase.ToFieldToDraftType()
				expectedField := testCase.ExpectedFieldToDraftType()
				toFieldEqualErr := toField.
					VerifyNotEqualExcludingInnerFieldsErr(
						expectedField)
				fromFieldEqualErr := convertedFrom.
					VerifyNotEqualExcludingInnerFieldsErr(expectedField)

				So(toFieldEqualErr, ShouldBeNil)
				So(fromFieldEqualErr, ShouldBeNil)

			case coretests.DraftType:
				toField := testCase.ToFieldToDraftType()
				expectedField := testCase.ExpectedFieldToDraftType()
				toFieldEqualErr := toField.
					VerifyNotEqualExcludingInnerFieldsErr(
						expectedField)
				fromFieldEqualErr := convertedFrom.
					VerifyNotEqualExcludingInnerFieldsErr(expectedField)

				So(toFieldEqualErr, ShouldBeNil)
				So(fromFieldEqualErr, ShouldBeNil)

			case []byte, *[]byte:
				// expecting unmarshalling to type
				// From, To: ([]byte or *[]byte, otherType) -- try unmarshal, reflect
				// To, Expected should be same
				toField := testCase.ToFieldToDraftType()
				toFieldEqualErr := toField.
					VerifyNotEqualExcludingInnerFieldsErr(
						testCase.ExpectedFieldToDraftType())
				So(toFieldEqualErr, ShouldBeNil)
			}
		})
	}
}
