package versionindexestests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/enums/versionindexes"
	"github.com/alimtvnetwork/core/errcore"
)

var indexByName = map[string]versionindexes.Index{
	"Major": versionindexes.Major,
	"Minor": versionindexes.Minor,
	"Patch": versionindexes.Patch,
	"Build": versionindexes.Build,
}

func Test_Index_JsonRoundtrip_Verification(t *testing.T) {
	for caseIndex, testCase := range jsonRoundtripTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		indexName, _ := input.GetAsString("index")
		idx := indexByName[indexName]

		// Act
		jsonResult := idx.Json()
		var restored versionindexes.Index
		err := restored.JsonParseSelfInject(&jsonResult)
		errcore.HandleErr(err)

		actual := args.Map{
			"indexName":  restored.Name(),
			"indexValue": restored.ToNumberString(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Index_NameAndNameValue_Verification(t *testing.T) {
	for caseIndex, testCase := range nameAndNameValueTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		indexName, _ := input.GetAsString("index")
		idx := indexByName[indexName]

		// Act
		actual := args.Map{
			"name":      idx.Name(),
			"nameValue": idx.NameValue(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Index_JsonParseSelfInject_Verification(t *testing.T) {
	for caseIndex, testCase := range jsonParseSelfInjectTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		sourceName, _ := input.GetAsString("source")
		targetName, _ := input.GetAsString("target")
		source := indexByName[sourceName]
		target := indexByName[targetName]

		// Act
		sourceJson := source.Json()
		err := target.JsonParseSelfInject(&sourceJson)
		errcore.HandleErr(err)

		actual := args.Map{
			"resultName":      target.Name(),
			"resultNameValue": target.NameValue(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
