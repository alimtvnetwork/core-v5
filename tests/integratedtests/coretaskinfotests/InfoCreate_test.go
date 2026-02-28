package coretaskinfotests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretaskinfo"
	"gitlab.com/auk-go/core/coretests/args"
)

func Test_Info_Create_Verification(t *testing.T) {
	for caseIndex, testCase := range infoCreateTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal := input.GetDirectLower("isNil")

		var name, desc, url string
		var isNull, isDefined string

		if isNilVal == true {
			// Act — nil info
			var info *coretaskinfo.Info
			name = info.SafeName()
			desc = info.SafeDescription()
			url = info.SafeUrl()
			isNull = fmt.Sprintf("%v", info.IsNull())
			isDefined = fmt.Sprintf("%v", info.IsDefined())
		} else {
			// Act — real info
			nameStr, _ := input.GetAsString("name")
			descStr, _ := input.GetAsString("desc")
			urlStr, _ := input.GetAsString("url")

			info := coretaskinfo.New.Info.Default(nameStr, descStr, urlStr)
			name = info.SafeName()
			desc = info.SafeDescription()
			url = info.SafeUrl()
			isNull = fmt.Sprintf("%v", info.IsNull())
			isDefined = fmt.Sprintf("%v", info.IsDefined())
		}

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			name,
			desc,
			url,
			isNull,
			isDefined,
		)
	}
}

func Test_Info_Serialize_Verification(t *testing.T) {
	for caseIndex, testCase := range infoSerializeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nameStr, _ := input.GetAsString("name")
		descStr, _ := input.GetAsString("desc")
		urlStr, _ := input.GetAsString("url")

		// Act
		info := coretaskinfo.New.Info.Default(nameStr, descStr, urlStr)
		jsonBytes := info.JsonPtr().Bytes
		deserialized, err := coretaskinfo.New.Info.Deserialized(jsonBytes)

		noErr := fmt.Sprintf("%v", err == nil)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			deserialized.SafeName(),
			deserialized.SafeDescription(),
			deserialized.SafeUrl(),
			noErr,
		)
	}
}
