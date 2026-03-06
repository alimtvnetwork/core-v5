package coreoncetests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/coreonce"
	"gitlab.com/auk-go/core/coretests/args"
)

func Test_AnyErrorOnce_Core(t *testing.T) {
	for caseIndex, tc := range anyErrorOnceCoreTestCases {
		// Arrange
		initVal := tc.InitValue
		initErr := tc.InitErr
		once := coreonce.NewAnyErrorOncePtr(func() (any, error) {
			return initVal, initErr
		})

		// Act
		actual := args.Map{
			"hasError":   once.HasError(),
			"isValid":    once.IsValid(),
			"isSuccess":  once.IsSuccess(),
			"isInvalid":  once.IsInvalid(),
			"isFailed":   once.IsFailed(),
			"isNull":     once.IsNull(),
			"isEmpty":    once.IsEmpty(),
			"hasAnyItem": once.HasAnyItem(),
			"isDefined":  once.IsDefined(),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_AnyErrorOnce_Caching(t *testing.T) {
	for caseIndex, tc := range anyErrorOnceCachingTestCases {
		// Arrange
		callCount := 0
		initVal := tc.InitValue
		once := coreonce.NewAnyErrorOncePtr(func() (any, error) {
			callCount++

			return initVal, nil
		})

		// Act
		v1, _ := once.Value()
		v2, _ := once.Execute()

		actual := args.Map{
			"callCount":      callCount,
			"executeEqValue": v1 == v2,
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_AnyErrorOnce_ValueMust_Success(t *testing.T) {
	tc := anyErrorOnceMustSuccessTestCase

	// Arrange
	once := coreonce.NewAnyErrorOncePtr(func() (any, error) {
		return tc.InitValue, nil
	})

	// Act
	panicked := callPanics(func() { once.ValueMust() })

	actual := args.Map{
		"didPanic": panicked,
	}

	// Assert
	tc.Case.ShouldBeEqualMapFirst(t, actual)
}

func Test_AnyErrorOnce_ValueMust_Panic(t *testing.T) {
	tc := anyErrorOnceMustPanicTestCase

	// Arrange
	once := coreonce.NewAnyErrorOncePtr(func() (any, error) {
		return nil, tc.InitErr
	})

	// Act
	panicked := callPanics(func() { once.ValueMust() })

	actual := args.Map{
		"didPanic": panicked,
	}

	// Assert
	tc.Case.ShouldBeEqualMapFirst(t, actual)
}

func Test_AnyErrorOnce_CastString(t *testing.T) {
	tc := anyErrorOnceCastStringTestCase

	// Arrange
	once := coreonce.NewAnyErrorOncePtr(func() (any, error) {
		return tc.InitValue, nil
	})

	// Act
	val, err, ok := once.CastValueString()

	actual := args.Map{
		"castValue":   val,
		"castSuccess": ok,
		"noError":     err == nil,
	}

	// Assert
	tc.Case.ShouldBeEqualMapFirst(t, actual)
}

func Test_AnyErrorOnce_Json(t *testing.T) {
	for caseIndex, tc := range anyErrorOnceJsonTestCases {
		// Arrange
		initVal := tc.InitValue
		initErr := tc.InitErr
		once := coreonce.NewAnyErrorOncePtr(func() (any, error) {
			return initVal, initErr
		})

		// Act
		var actual args.Map

		data, err := once.Serialize()
		if tc.InitErr != nil {
			actual = args.Map{
				"hasError": err != nil,
			}
		} else {
			actual = args.Map{
				"noError":             err == nil,
				"dataLengthAboveZero": len(data) > 0,
			}
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_AnyErrorOnce_Constructor(t *testing.T) {
	for caseIndex, tc := range anyErrorOnceConstructorTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewAnyErrorOnce(func() (any, error) {
			return initVal, nil
		})

		// Act
		_, err := once.Value()

		actual := args.Map{
			"isNull":  once.IsNull(),
			"noError": err == nil,
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
