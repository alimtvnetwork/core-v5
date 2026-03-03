package coreoncetests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coreonce"
)

func Test_BoolOnce_Core(t *testing.T) {
	for caseIndex, tc := range boolOnceCoreTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewBoolOncePtr(func() bool { return initVal })

		// Act & Assert
		tc.Case.ShouldBeEqual(t, caseIndex,
			fmt.Sprintf("%v", once.Value()),
			once.String(),
		)
	}
}

func Test_BoolOnce_Caching(t *testing.T) {
	for caseIndex, tc := range boolOnceCachingTestCases {
		// Arrange
		callCount := 0
		initVal := tc.InitValue
		once := coreonce.NewBoolOncePtr(func() bool {
			callCount++

			return initVal
		})

		// Act
		r1 := once.Value()
		r2 := once.Value()
		r3 := once.Value()

		// Assert
		tc.Case.ShouldBeEqual(t, caseIndex,
			fmt.Sprintf("%v", r1),
			fmt.Sprintf("%v", r2),
			fmt.Sprintf("%v", r3),
			fmt.Sprintf("%d", callCount),
		)
	}
}

func Test_BoolOnce_Json(t *testing.T) {
	for caseIndex, tc := range boolOnceJsonTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewBoolOncePtr(func() bool { return initVal })

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
