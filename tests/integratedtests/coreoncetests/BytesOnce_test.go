package coreoncetests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_BytesOnce_Core(t *testing.T) {
	for caseIndex, tc := range bytesOnceCoreTestCases {
		// Arrange
		var once *coreonce.BytesOnce

		if tc.UseNilInit {
			once = &coreonce.BytesOnce{}
		} else {
			initBytes := tc.InitBytes
			once = coreonce.NewBytesOncePtr(func() []byte { return initBytes })
		}

		// Act
		val := once.Value()
		actual := args.Map{
			"stringOfValue": string(val),
			"stringMethod":  once.String(),
			"isEmpty":       once.IsEmpty(),
			"length":        once.Length(),
			"isNil":         val == nil,
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_BytesOnce_Caching(t *testing.T) {
	for caseIndex, tc := range bytesOnceCachingTestCases {
		// Arrange
		callCount := 0
		initBytes := tc.InitBytes
		once := coreonce.NewBytesOncePtr(func() []byte {
			callCount++

			return initBytes
		})

		// Act
		r1 := once.Value()
		r2 := once.Value()
		r3 := once.Value()

		actual := args.Map{
			"r1":             string(r1),
			"r2":             string(r2),
			"r3":             string(r3),
			"callCount":      callCount,
			"executeEqValue": string(once.Execute()) == string(once.Value()),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_BytesOnce_JSON(t *testing.T) {
	for caseIndex, tc := range bytesOnceJsonTestCases {
		// Arrange
		initBytes := tc.InitBytes
		once := coreonce.NewBytesOncePtr(func() []byte { return initBytes })

		// Act
		var actual args.Map

		if tc.ReplaceBytes != nil {
			input, _ := json.Marshal(tc.ReplaceBytes)
			err := once.UnmarshalJSON(input)

			actual = args.Map{
				"noError":  err == nil,
				"newValue": string(once.Value()),
			}
		} else {
			data, err := once.MarshalJSON()

			actual = args.Map{
				"noError":             err == nil,
				"dataLengthAboveZero": len(data) > 0,
			}
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_BytesOnce_Constructor(t *testing.T) {
	for caseIndex, tc := range bytesOnceConstructorTestCases {
		// Arrange
		initBytes := tc.InitBytes
		once := coreonce.NewBytesOnce(func() []byte { return initBytes })

		// Act
		actual := args.Map{
			"constructedValue": string(once.Value()),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// Ensure fmt is used
var _ = fmt.Sprintf
