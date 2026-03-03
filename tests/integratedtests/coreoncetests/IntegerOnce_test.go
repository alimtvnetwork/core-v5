package coreoncetests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coreonce"
)

func Test_IntegerOnce_Core(t *testing.T) {
	for caseIndex, tc := range integerOnceCoreTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewIntegerOncePtr(func() int { return initVal })

		// Act & Assert
		tc.Case.ShouldBeEqual(t, caseIndex,
			fmt.Sprintf("%d", once.Value()),
			once.String(),
			fmt.Sprintf("%v", once.IsZero()),
			fmt.Sprintf("%v", once.IsEmpty()),
			fmt.Sprintf("%v", once.IsAboveZero()),
			fmt.Sprintf("%v", once.IsPositive()),
			fmt.Sprintf("%v", once.IsLessThanZero()),
			fmt.Sprintf("%v", once.IsNegative()),
		)
	}
}

func Test_IntegerOnce_Caching(t *testing.T) {
	for caseIndex, tc := range integerOnceCachingTestCases {
		// Arrange
		callCount := 0
		initVal := tc.InitValue
		once := coreonce.NewIntegerOncePtr(func() int {
			callCount++

			return initVal
		})

		// Act
		r1 := once.Value()
		r2 := once.Value()

		// Assert
		tc.Case.ShouldBeEqual(t, caseIndex,
			fmt.Sprintf("%d", r1),
			fmt.Sprintf("%d", r2),
			fmt.Sprintf("%d", callCount),
		)
	}
}

func Test_IntegerOnce_Compare(t *testing.T) {
	for caseIndex, tc := range integerOnceCompareTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewIntegerOncePtr(func() int { return initVal })
		cmpVal := tc.CompareValue

		// Act
		var actLines []string

		isLessThanCase := tc.InitValue < tc.CompareValue
		if isLessThanCase {
			actLines = []string{
				fmt.Sprintf("%v", once.IsLessThan(cmpVal)),
				fmt.Sprintf("%v", once.IsLessThan(initVal)),
				fmt.Sprintf("%v", once.IsLessThanEqual(initVal)),
			}
		} else {
			actLines = []string{
				fmt.Sprintf("%v", once.IsAbove(cmpVal)),
				fmt.Sprintf("%v", once.IsAbove(initVal)),
				fmt.Sprintf("%v", once.IsAboveEqual(initVal)),
			}
		}

		// Assert
		tc.Case.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

func Test_IntegerOnce_Json(t *testing.T) {
	for caseIndex, tc := range integerOnceJsonTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewIntegerOncePtr(func() int { return initVal })

		// Act
		data, err := once.MarshalJSON()
		noError := err == nil

		// Assert
		tc.Case.ShouldBeEqual(t, caseIndex,
			fmt.Sprintf("%v", noError),
			string(data),
		)
	}
}
