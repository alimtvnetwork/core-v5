package coreoncetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_AnyOnce_Core(t *testing.T) {
	for caseIndex, tc := range anyOnceCoreTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewAnyOncePtr(func() any { return initVal })

		// Act — trigger initialization
		_ = once.Value()

		actual := args.Map{
			"isNull":                    once.IsNull(),
			"isStringEmpty":             once.IsStringEmpty(),
			"isStringEmptyOrWhitespace": once.IsStringEmptyOrWhitespace(),
			"isInitialized":             once.IsInitialized(),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_AnyOnce_CastString(t *testing.T) {
	tc := anyOnceCastStringTestCase

	// Arrange
	once := coreonce.NewAnyOncePtr(func() any { return tc.InitValue })

	// Act
	val, ok := once.CastValueString()

	actual := args.Map{
		"castValue":   val,
		"castSuccess": ok,
	}

	// Assert
	tc.Case.ShouldBeEqualMapFirst(t, actual)
}

func Test_AnyOnce_CastStrings(t *testing.T) {
	tc := anyOnceCastStringsTestCase

	// Arrange
	once := coreonce.NewAnyOncePtr(func() any { return tc.InitValue })

	// Act
	val, ok := once.CastValueStrings()

	actual := args.Map{
		"castLen":     len(val),
		"castSuccess": ok,
	}

	// Assert
	tc.Case.ShouldBeEqualMapFirst(t, actual)
}

func Test_AnyOnce_CastBytes(t *testing.T) {
	tc := anyOnceCastBytesTestCase

	// Arrange
	once := coreonce.NewAnyOncePtr(func() any { return tc.InitValue })

	// Act
	val, ok := once.CastValueBytes()

	actual := args.Map{
		"castLen":     len(val),
		"castSuccess": ok,
	}

	// Assert
	tc.Case.ShouldBeEqualMapFirst(t, actual)
}

func Test_AnyOnce_CastMap(t *testing.T) {
	tc := anyOnceCastMapTestCase

	// Arrange
	once := coreonce.NewAnyOncePtr(func() any { return tc.InitValue })

	// Act
	val, ok := once.CastValueHashmapMap()

	actual := args.Map{
		"castLen":     len(val),
		"castSuccess": ok,
	}

	// Assert
	tc.Case.ShouldBeEqualMapFirst(t, actual)
}

func Test_AnyOnce_Caching(t *testing.T) {
	for caseIndex, tc := range anyOnceCachingTestCases {
		// Arrange
		callCount := 0
		initVal := tc.InitValue
		once := coreonce.NewAnyOncePtr(func() any {
			callCount++

			return initVal
		})

		// Act
		_ = once.Value()
		_ = once.Value()
		_ = once.Execute()

		actual := args.Map{
			"callCount": callCount,
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_AnyOnce_Json(t *testing.T) {
	for caseIndex, tc := range anyOnceJsonTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewAnyOncePtr(func() any { return initVal })

		// Act
		data, err := once.Serialize()

		actual := args.Map{
			"noError":             err == nil,
			"dataLengthAboveZero": len(data) > 0,
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_AnyOnce_Constructor(t *testing.T) {
	for caseIndex, tc := range anyOnceConstructorTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewAnyOnce(func() any { return initVal })

		// Act
		actual := args.Map{
			"isNull": once.IsNull(),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
