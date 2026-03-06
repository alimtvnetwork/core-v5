package enumimpltests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coreimpl/enumimpl"
	"gitlab.com/auk-go/core/coretests/args"
)

func Test_EnumByte_MinMax(t *testing.T) {
	for caseIndex, tc := range enumByteTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		enumMap := input["enum-map"].(enumimpl.DynamicMap)
		enumImpl := enumMap.BasicByte("unknown type")

		// Act
		actual := args.Map{
			"min": fmt.Sprintf("%v", enumImpl.Min()),
			"max": fmt.Sprintf("%v", enumImpl.Max()),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_EnumInt8_MinMax(t *testing.T) {
	for caseIndex, tc := range enumInt8TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		enumMap := input["enum-map"].(enumimpl.DynamicMap)
		enumImpl := enumMap.BasicInt8("unknown type")

		// Act
		actual := args.Map{
			"min": fmt.Sprintf("%v", enumImpl.Min()),
			"max": fmt.Sprintf("%v", enumImpl.Max()),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_EnumInt16_MinMax(t *testing.T) {
	for caseIndex, tc := range enumInt16TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		enumMap := input["enum-map"].(enumimpl.DynamicMap)
		enumImpl := enumMap.BasicInt16("unknown type")

		// Act
		actual := args.Map{
			"min": fmt.Sprintf("%v", enumImpl.Min()),
			"max": fmt.Sprintf("%v", enumImpl.Max()),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_EnumInt32_MinMax(t *testing.T) {
	for caseIndex, tc := range enumInt32TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		enumMap := input["enum-map"].(enumimpl.DynamicMap)
		enumImpl := enumMap.BasicInt32("unknown type")

		// Act
		actual := args.Map{
			"min": fmt.Sprintf("%v", enumImpl.Min()),
			"max": fmt.Sprintf("%v", enumImpl.Max()),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_EnumUInt16_MinMax(t *testing.T) {
	for caseIndex, tc := range enumUInt16TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		enumMap := input["enum-map"].(enumimpl.DynamicMap)
		enumImpl := enumMap.BasicUInt16("unknown type")

		// Act
		actual := args.Map{
			"min": fmt.Sprintf("%v", enumImpl.Min()),
			"max": fmt.Sprintf("%v", enumImpl.Max()),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_EnumString_MinMax(t *testing.T) {
	for caseIndex, tc := range enumStringTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		enumMap := input["enum-map"].(enumimpl.DynamicMap)
		enumImpl := enumMap.BasicString("unknown type")

		// Act
		actual := args.Map{
			"min": fmt.Sprintf("%v", enumImpl.Min()),
			"max": fmt.Sprintf("%v", enumImpl.Max()),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
