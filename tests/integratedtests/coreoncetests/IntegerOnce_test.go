package coreoncetests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_IntegerOnce_Core(t *testing.T) {
	for caseIndex, tc := range integerOnceCoreTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewIntegerOncePtr(func() int { return initVal })

		// Act
		actual := args.Map{
			"value":          once.Value(),
			"string":         once.String(),
			"isZero":         once.IsZero(),
			"isEmpty":        once.IsEmpty(),
			"isAboveZero":    once.IsAboveZero(),
			"isPositive":     once.IsPositive(),
			"isLessThanZero": once.IsLessThanZero(),
			"isNegative":     once.IsNegative(),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"r1":        r1,
			"r2":        r2,
			"callCount": callCount,
		}
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegerOnce_Compare(t *testing.T) {
	for caseIndex, tc := range integerOnceCompareTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewIntegerOncePtr(func() int { return initVal })
		cmpVal := tc.CompareValue

		// Act
		var actual args.Map

		isLessThanCase := tc.InitValue < tc.CompareValue
		if isLessThanCase {
			actual = args.Map{
				"isLessThanCompare":   once.IsLessThan(cmpVal),
				"isLessThanSelf":      once.IsLessThan(initVal),
				"isLessThanEqualSelf": once.IsLessThanEqual(initVal),
			}
		} else {
			actual = args.Map{
				"isAboveCompare":   once.IsAbove(cmpVal),
				"isAboveSelf":      once.IsAbove(initVal),
				"isAboveEqualSelf": once.IsAboveEqual(initVal),
			}
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegerOnce_Json(t *testing.T) {
	for caseIndex, tc := range integerOnceJsonTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewIntegerOncePtr(func() int { return initVal })

		// Act
		data, err := once.MarshalJSON()

		// Assert
		actual := args.Map{
			"noError":        err == nil,
			"marshaledValue": string(data),
		}
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// Ensure fmt is used (for potential future use)
var _ = fmt.Sprintf
