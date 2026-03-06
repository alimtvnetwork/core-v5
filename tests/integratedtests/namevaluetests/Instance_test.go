package namevaluetests

import (
	"strings"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/namevalue"
)

// ==========================================================================
// Test: StringAny.String
// ==========================================================================

func Test_StringAny_String_Verification(t *testing.T) {
	for caseIndex, testCase := range stringAnyStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		name, _ := input.GetAsString("name")
		value, _ := input.Get("value")

		// Act
		instance := namevalue.StringAny{
			Name:  name,
			Value: value,
		}
		result := instance.String()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

// ==========================================================================
// Test: StringString.String
// ==========================================================================

func Test_StringString_String_Verification(t *testing.T) {
	for caseIndex, testCase := range stringStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		name, _ := input.GetAsString("name")
		value, _ := input.GetAsString("value")

		// Act
		instance := namevalue.StringString{
			Name:  name,
			Value: value,
		}
		result := instance.String()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

// ==========================================================================
// Test: StringInt.String
// ==========================================================================

func Test_StringInt_String_Verification(t *testing.T) {
	for caseIndex, testCase := range stringIntTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		name, _ := input.GetAsString("name")
		value, _ := input.Get("value")
		intVal := value.(int)

		// Act
		instance := namevalue.StringInt{
			Name:  name,
			Value: intVal,
		}
		result := instance.String()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

// ==========================================================================
// Test: StringMapAny.String — Populated
// ==========================================================================

func Test_StringMapAny_Populated(t *testing.T) {
	tc := stringMapAnyPopulatedTestCase

	instance := namevalue.StringMapAny{
		Name:  "config",
		Value: map[string]any{"host": "localhost", "port": 8080},
	}
	result := instance.String()

	actual := args.Map{
		"isValidJson":  result != "",
		"containsName": strings.Contains(result, "config"),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: StringMapAny.String — Empty
// ==========================================================================

func Test_StringMapAny_Empty(t *testing.T) {
	tc := stringMapAnyEmptyTestCase

	instance := namevalue.StringMapAny{
		Name:  "empty",
		Value: map[string]any{},
	}
	result := instance.String()

	actual := args.Map{
		"isValidJson":  result != "",
		"containsName": strings.Contains(result, "empty"),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: StringMapAny.String — Nil
// ==========================================================================

func Test_StringMapAny_Nil(t *testing.T) {
	tc := stringMapAnyNilTestCase

	instance := namevalue.StringMapAny{
		Name:  "nothing",
		Value: nil,
	}
	result := instance.String()

	actual := args.Map{
		"isValidJson":  result != "",
		"containsName": strings.Contains(result, "nothing"),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: StringMapString.String — Populated
// ==========================================================================

func Test_StringMapString_Populated(t *testing.T) {
	tc := stringMapStringPopulatedTestCase

	instance := namevalue.StringMapString{
		Name:  "headers",
		Value: map[string]string{"Content-Type": "application/json"},
	}
	result := instance.String()

	actual := args.Map{
		"isValidJson":  result != "",
		"containsName": strings.Contains(result, "headers"),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: StringMapString.String — Nil
// ==========================================================================

func Test_StringMapString_Nil(t *testing.T) {
	tc := stringMapStringNilTestCase

	instance := namevalue.StringMapString{
		Name:  "nothing",
		Value: nil,
	}
	result := instance.String()

	actual := args.Map{
		"isValidJson":  result != "",
		"containsName": strings.Contains(result, "nothing"),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Dispose — StringAny
// ==========================================================================

func Test_Dispose_StringAny(t *testing.T) {
	tc := disposeStringAnyTestCase

	inst := &namevalue.StringAny{Name: "key", Value: "val"}
	inst.Dispose()

	actual := args.Map{
		"disposedName": inst.Name,
		"isNilValue":   inst.Value == nil,
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Dispose — StringString
// ==========================================================================

func Test_Dispose_StringString(t *testing.T) {
	tc := disposeStringStringTestCase

	inst := &namevalue.StringString{Name: "key", Value: "val"}
	inst.Dispose()

	actual := args.Map{
		"disposedName":  inst.Name,
		"disposedValue": inst.Value,
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Dispose — StringInt
// ==========================================================================

func Test_Dispose_StringInt(t *testing.T) {
	tc := disposeStringIntTestCase

	inst := &namevalue.StringInt{Name: "count", Value: 42}
	inst.Dispose()

	actual := args.Map{
		"disposedName":  inst.Name,
		"disposedValue": inst.Value,
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: JsonString — StringAny
// ==========================================================================

func Test_JsonString_StringAny(t *testing.T) {
	tc := jsonStringStringAnyTestCase

	inst := namevalue.StringAny{Name: "server", Value: "api.example.com"}
	jsonStr := inst.JsonString()

	actual := args.Map{
		"isValidJson": jsonStr != "",
		"containsKey": strings.Contains(jsonStr, "server"),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: JsonString — StringInt
// ==========================================================================

func Test_JsonString_StringInt(t *testing.T) {
	tc := jsonStringStringIntTestCase

	inst := namevalue.StringInt{Name: "port", Value: 443}
	jsonStr := inst.JsonString()

	actual := args.Map{
		"isValidJson":    jsonStr != "",
		"containsNumber": strings.Contains(jsonStr, "port"),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection
// ==========================================================================

func Test_Collection_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count, _ := input.Get("count")
		countInt := count.(int)

		// Act
		collection := namevalue.NewCollection()
		for i := 0; i < countInt; i++ {
			collection.Add(namevalue.StringAny{
				Name:  "key",
				Value: i,
			})
		}

		actual := args.Map{
			"length":  collection.Length(),
			"isEmpty": collection.IsEmpty(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// Test: Chmod — VarNameValues with single item
// ==========================================================================

func Test_Chmod_VarNameValues_Single(t *testing.T) {
	tc := chmodVarNameValuesSingleTestCase

	nv := namevalue.StringAny{
		Name:  "Location",
		Value: "/tmp/test",
	}
	result := namevalue.VarNameValues(nv)

	actual := args.Map{
		"containsName":  result != "",
		"containsValue": strings.Contains(result, "/tmp/test"),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Chmod — MessageNameValues
// ==========================================================================

func Test_Chmod_MessageNameValues(t *testing.T) {
	tc := chmodMessageNameValuesTestCase

	nv := namevalue.StringAny{
		Name:  "Path",
		Value: "/usr/local/bin",
	}
	result := namevalue.MessageNameValues("chmod verification failed", nv)

	actual := args.Map{
		"containsMessage":   strings.Contains(result, "chmod verification failed"),
		"containsNameValue": strings.Contains(result, "/usr/local/bin"),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Chmod — VarNameValues empty
// ==========================================================================

func Test_Chmod_VarNameValues_Empty(t *testing.T) {
	tc := chmodVarNameValuesEmptyTestCase
	result := namevalue.VarNameValues()
	tc.ShouldBeEqual(t, 0, result)
}
