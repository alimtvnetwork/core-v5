package coreoncetests

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"gitlab.com/auk-go/core/coredata/coreonce"
)

func newErrorOnce(initError string) *coreonce.ErrorOnce {
	if initError == "" {
		return coreonce.NewErrorOncePtr(func() error { return nil })
	}

	if initError == "empty-marker" {
		return coreonce.NewErrorOncePtr(func() error { return errors.New("") })
	}

	return coreonce.NewErrorOncePtr(func() error { return errors.New(initError) })
}

func Test_ErrorOnce_Core(t *testing.T) {
	for caseIndex, tc := range errorOnceCoreTestCases {
		// Arrange
		once := newErrorOnce(tc.InitError)

		// Act & Assert
		tc.Case.ShouldBeEqual(t, caseIndex,
			fmt.Sprintf("%v", once.HasError()),
			fmt.Sprintf("%v", once.IsValid()),
			fmt.Sprintf("%v", once.IsSuccess()),
			fmt.Sprintf("%v", once.IsEmpty()),
			fmt.Sprintf("%v", once.IsInvalid()),
			fmt.Sprintf("%v", once.IsFailed()),
			fmt.Sprintf("%v", once.HasAnyItem()),
			fmt.Sprintf("%v", once.IsDefined()),
			once.Message(),
		)
	}
}

func Test_ErrorOnce_Caching(t *testing.T) {
	for caseIndex, tc := range errorOnceCachingTestCases {
		// Arrange
		callCount := 0
		initErr := tc.InitError
		once := coreonce.NewErrorOncePtr(func() error {
			callCount++

			return errors.New(initErr)
		})

		// Act
		r1 := once.Value()
		r2 := once.Value()
		r3 := once.Value()

		// Assert
		tc.Case.ShouldBeEqual(t, caseIndex,
			r1.Error(),
			r2.Error(),
			r3.Error(),
			fmt.Sprintf("%d", callCount),
		)
	}
}

func Test_ErrorOnce_NullOrEmpty(t *testing.T) {
	for caseIndex, tc := range errorOnceNullOrEmptyTestCases {
		// Arrange
		once := newErrorOnce(tc.InitError)

		// Act & Assert
		tc.Case.ShouldBeEqual(t, caseIndex,
			fmt.Sprintf("%v", once.IsNullOrEmpty()),
		)
	}
}

func Test_ErrorOnce_MessageEqual(t *testing.T) {
	for caseIndex, tc := range errorOnceMessageEqualTestCases {
		// Arrange
		once := newErrorOnce(tc.InitError)

		// Act & Assert
		tc.Case.ShouldBeEqual(t, caseIndex,
			fmt.Sprintf("%v", once.IsMessageEqual(tc.MatchMsg)),
			fmt.Sprintf("%v", once.IsMessageEqual("other")),
		)
	}
}

func Test_ErrorOnce_ConcatNew(t *testing.T) {
	for caseIndex, tc := range errorOnceConcatTestCases {
		// Arrange
		once := newErrorOnce(tc.InitError)

		// Act
		result := once.ConcatNewString(tc.ExtraMsg)

		var actLines []string

		isNilError := tc.InitError == ""
		if isNilError {
			actLines = []string{result}
		} else {
			actLines = []string{
				fmt.Sprintf("%v", strings.Contains(result, tc.InitError)),
				fmt.Sprintf("%v", strings.Contains(result, tc.ExtraMsg)),
			}
		}

		// Assert
		tc.Case.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

func Test_ErrorOnce_Json(t *testing.T) {
	for caseIndex, tc := range errorOnceJsonTestCases {
		// Arrange
		once := newErrorOnce(tc.InitError)

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
