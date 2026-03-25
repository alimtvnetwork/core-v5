package coreoncetests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_BoolOnce_Core(t *testing.T) {
	for caseIndex, tc := range boolOnceCoreTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewBoolOncePtr(func() bool { return initVal })

		// Act
		actual := args.Map{
			"value":  once.Value(),
			"string": once.String(),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"r1":        r1,
			"r2":        r2,
			"r3":        r3,
			"callCount": callCount,
		}
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_BoolOnce_Json(t *testing.T) {
	for caseIndex, tc := range boolOnceJsonTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewBoolOncePtr(func() bool { return initVal })

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

// Ensure fmt is used
var _ = fmt.Sprintf
