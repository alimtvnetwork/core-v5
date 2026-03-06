package coreoncetests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/coreonce"
	"gitlab.com/auk-go/core/coretests/args"
)

func Test_ByteOnce_Core(t *testing.T) {
	for caseIndex, tc := range byteOnceCoreTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewByteOncePtr(func() byte { return initVal })

		// Act
		actual := args.Map{
			"value":      int(once.Value()),
			"int":        once.Int(),
			"string":     once.String(),
			"isEmpty":    once.IsEmpty(),
			"isZero":     once.IsZero(),
			"isNegative": once.IsNegative(),
			"isPositive": once.IsPositive(),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_ByteOnce_Caching(t *testing.T) {
	for caseIndex, tc := range byteOnceCachingTestCases {
		// Arrange
		callCount := 0
		initVal := tc.InitValue
		once := coreonce.NewByteOncePtr(func() byte {
			callCount++

			return initVal
		})

		// Act
		r1 := once.Value()
		r2 := once.Value()

		// Assert
		actual := args.Map{
			"r1":        int(r1),
			"r2":        int(r2),
			"callCount": callCount,
		}
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_ByteOnce_Json(t *testing.T) {
	for caseIndex, tc := range byteOnceJsonTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewByteOncePtr(func() byte { return initVal })

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

func Test_ByteOnce_Serialize(t *testing.T) {
	for caseIndex, tc := range byteOnceSerializeTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewByteOncePtr(func() byte { return initVal })

		// Act
		data, err := once.Serialize()

		// Assert
		actual := args.Map{
			"noError":         err == nil,
			"serializedValue": string(data),
		}
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_ByteOnce_Constructor(t *testing.T) {
	for caseIndex, tc := range byteOnceConstructorTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewByteOnce(func() byte { return initVal })

		// Act
		actual := args.Map{
			"constructedValue": int(once.Value()),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
