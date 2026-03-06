package coreoncetests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coreonce"
	"gitlab.com/auk-go/core/coretests/args"
)

func Test_StringOnce_Core(t *testing.T) {
	for caseIndex, tc := range stringOnceCoreTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewStringOncePtr(func() string { return initVal })

		// Act
		actual := args.Map{
			"value":               once.Value(),
			"string":              once.String(),
			"isEmpty":             once.IsEmpty(),
			"isEmptyOrWhitespace": once.IsEmptyOrWhitespace(),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_StringOnce_Caching(t *testing.T) {
	for caseIndex, tc := range stringOnceCachingTestCases {
		// Arrange
		callCount := 0
		initVal := tc.InitValue
		once := coreonce.NewStringOncePtr(func() string {
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

func Test_StringOnce_Match(t *testing.T) {
	for caseIndex, tc := range stringOnceMatchTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewStringOncePtr(func() string { return initVal })

		// Act
		var actual args.Map

		hasPrefix := tc.MatchArg == "prefix"
		hasSuffix := tc.MatchArg == "suffix"

		if hasPrefix {
			actual = args.Map{
				"matchResult":   once.HasPrefix(tc.MatchArg),
				"noMatchResult": once.HasPrefix("data"),
			}
		} else if hasSuffix {
			actual = args.Map{
				"matchResult":   once.HasSuffix(tc.MatchArg),
				"noMatchResult": once.HasSuffix("data"),
			}
		} else if tc.MatchArg == tc.InitValue {
			actual = args.Map{
				"matchResult":   once.IsEqual(tc.MatchArg),
				"noMatchResult": once.IsEqual("xyz"),
			}
		} else {
			actual = args.Map{
				"matchResult":   once.IsContains(tc.MatchArg),
				"noMatchResult": once.IsContains("xyz"),
			}
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_StringOnce_Split(t *testing.T) {
	for caseIndex, tc := range stringOnceSplitTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewStringOncePtr(func() string { return initVal })

		// Act
		var actual args.Map

		switch tc.Method {
		case "splitBy":
			parts := once.SplitBy(tc.Splitter)
			actual = args.Map{
				"partsLength": len(parts),
				"firstPart":   parts[0],
				"lastPart":    parts[len(parts)-1],
			}
		case "splitLeftRightTrim":
			left, right := once.SplitLeftRightTrim(tc.Splitter)
			actual = args.Map{
				"left":  left,
				"right": right,
			}
		default: // splitLeftRight
			left, right := once.SplitLeftRight(tc.Splitter)
			actual = args.Map{
				"left":  left,
				"right": right,
			}
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_StringOnce_Json(t *testing.T) {
	for caseIndex, tc := range stringOnceJsonTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewStringOncePtr(func() string { return initVal })

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
