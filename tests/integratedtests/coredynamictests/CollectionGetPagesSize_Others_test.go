package coredynamictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/errcore"
)

func Test_AnyCollection_GetPagesSize_Verification(t *testing.T) {
	for caseIndex, testCase := range anyCollectionGetPagesSizeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count, _ := input.GetAsInt("count")
		eachPageSize, _ := input.GetAsInt("eachPageSize")

		collection := coredynamic.NewAnyCollection(count)
		for i := 0; i < count; i++ {
			collection.Add(i)
		}

		// Act
		result := collection.GetPagesSize(eachPageSize)

		// Assert
		errcore.AssertDiffOnMismatch(
			t, caseIndex, testCase.Title,
			[]string{fmt.Sprintf("%v", result)},
			testCase.ExpectedInput.([]string),
		)
	}
}

func Test_DynamicCollection_GetPagesSize_Verification(t *testing.T) {
	for caseIndex, testCase := range dynamicCollectionGetPagesSizeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count, _ := input.GetAsInt("count")
		eachPageSize, _ := input.GetAsInt("eachPageSize")

		collection := coredynamic.NewDynamicCollection(count)
		for i := 0; i < count; i++ {
			collection.Add(coredynamic.Dynamic{Value: i})
		}

		// Act
		result := collection.GetPagesSize(eachPageSize)

		// Assert
		errcore.AssertDiffOnMismatch(
			t, caseIndex, testCase.Title,
			[]string{fmt.Sprintf("%v", result)},
			testCase.ExpectedInput.([]string),
		)
	}
}

func Test_KeyValCollection_GetPagesSize_Verification(t *testing.T) {
	for caseIndex, testCase := range keyValCollectionGetPagesSizeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count, _ := input.GetAsInt("count")
		eachPageSize, _ := input.GetAsInt("eachPageSize")

		collection := coredynamic.NewKeyValCollection(count)
		for i := 0; i < count; i++ {
			collection.Add(coredynamic.KeyVal{
				Key:   fmt.Sprintf("key-%d", i),
				Value: i,
			})
		}

		// Act
		result := collection.GetPagesSize(eachPageSize)

		// Assert
		errcore.AssertDiffOnMismatch(
			t, caseIndex, testCase.Title,
			[]string{fmt.Sprintf("%v", result)},
			testCase.ExpectedInput.([]string),
		)
	}
}
