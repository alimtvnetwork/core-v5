package namevaluetests

import (
	"fmt"
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
// Test: StringMapAny.String
// ==========================================================================

func Test_StringMapAny_Populated(t *testing.T) {
	tc := stringMapAnyTestCases[0]

	instance := namevalue.StringMapAny{
		Name:  "config",
		Value: map[string]any{"host": "localhost", "port": 8080},
	}
	result := instance.String()

	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%v", result != ""),
		fmt.Sprintf("%v", strings.Contains(result, "config")),
	)
}

func Test_StringMapAny_Empty(t *testing.T) {
	tc := stringMapAnyTestCases[1]

	instance := namevalue.StringMapAny{
		Name:  "empty",
		Value: map[string]any{},
	}
	result := instance.String()

	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%v", result != ""),
		fmt.Sprintf("%v", strings.Contains(result, "empty")),
	)
}

func Test_StringMapAny_Nil(t *testing.T) {
	tc := stringMapAnyTestCases[2]

	instance := namevalue.StringMapAny{
		Name:  "nothing",
		Value: nil,
	}
	result := instance.String()

	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%v", result != ""),
		fmt.Sprintf("%v", strings.Contains(result, "nothing")),
	)
}

// ==========================================================================
// Test: StringMapString.String
// ==========================================================================

func Test_StringMapString_Populated(t *testing.T) {
	tc := stringMapStringTestCases[0]

	instance := namevalue.StringMapString{
		Name:  "headers",
		Value: map[string]string{"Content-Type": "application/json"},
	}
	result := instance.String()

	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%v", result != ""),
		fmt.Sprintf("%v", strings.Contains(result, "headers")),
	)
}

func Test_StringMapString_Nil(t *testing.T) {
	tc := stringMapStringTestCases[1]

	instance := namevalue.StringMapString{
		Name:  "nothing",
		Value: nil,
	}
	result := instance.String()

	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%v", result != ""),
		fmt.Sprintf("%v", strings.Contains(result, "nothing")),
	)
}

// ==========================================================================
// Test: Dispose — StringAny
// ==========================================================================

func Test_Dispose_StringAny(t *testing.T) {
	tc := genericDisposeTestCases[0]

	inst := &namevalue.StringAny{Name: "key", Value: "val"}
	inst.Dispose()

	tc.ShouldBeEqual(t, 0,
		inst.Name,
		fmt.Sprintf("%v", inst.Value == nil),
	)
}

// ==========================================================================
// Test: Dispose — StringString
// ==========================================================================

func Test_Dispose_StringString(t *testing.T) {
	tc := genericDisposeTestCases[1]

	inst := &namevalue.StringString{Name: "key", Value: "val"}
	inst.Dispose()

	tc.ShouldBeEqual(t, 0,
		inst.Name,
		inst.Value,
	)
}

// ==========================================================================
// Test: Dispose — StringInt
// ==========================================================================

func Test_Dispose_StringInt(t *testing.T) {
	tc := genericDisposeTestCases[2]

	inst := &namevalue.StringInt{Name: "count", Value: 42}
	inst.Dispose()

	tc.ShouldBeEqual(t, 0,
		inst.Name,
		fmt.Sprintf("%d", inst.Value),
	)
}

// ==========================================================================
// Test: JsonString — StringAny
// ==========================================================================

func Test_JsonString_StringAny(t *testing.T) {
	tc := genericJsonStringTestCases[0]

	inst := namevalue.StringAny{Name: "server", Value: "api.example.com"}
	jsonStr := inst.JsonString()

	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%v", jsonStr != ""),
		fmt.Sprintf("%v", strings.Contains(jsonStr, "server")),
	)
}

// ==========================================================================
// Test: JsonString — StringInt
// ==========================================================================

func Test_JsonString_StringInt(t *testing.T) {
	tc := genericJsonStringTestCases[1]

	inst := namevalue.StringInt{Name: "port", Value: 443}
	jsonStr := inst.JsonString()

	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%v", jsonStr != ""),
		fmt.Sprintf("%v", strings.Contains(jsonStr, "port")),
	)
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
				Name:  fmt.Sprintf("key%d", i),
				Value: i,
			})
		}
		length := fmt.Sprintf("%d", collection.Length())
		isEmpty := fmt.Sprintf("%v", collection.IsEmpty())

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, length, isEmpty)
	}
}

// ==========================================================================
// Test: Chmod integration
// ==========================================================================

func Test_Chmod_Integration_Verification(t *testing.T) {
	for caseIndex, testCase := range chmodIntegrationTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		name, hasName := input.GetAsString("name")

		// Act
		if !hasName {
			result := namevalue.VarNameValues()
			testCase.ShouldBeEqual(t, caseIndex, result)
			continue
		}

		path, _ := input.GetAsString("path")
		message, hasMessage := input.GetAsString("message")

		nv := namevalue.StringAny{
			Name:  name,
			Value: path,
		}

		if hasMessage && message != "" {
			result := namevalue.MessageNameValues(message, nv)
			containsMessage := fmt.Sprintf("%v", strings.Contains(result, message))
			containsPath := fmt.Sprintf("%v", strings.Contains(result, path))
			testCase.ShouldBeEqual(t, caseIndex, containsMessage, containsPath)
		} else {
			result := namevalue.VarNameValues(nv)
			isNotEmpty := fmt.Sprintf("%v", result != "")
			containsPath := fmt.Sprintf("%v", strings.Contains(result, path))
			testCase.ShouldBeEqual(t, caseIndex, isNotEmpty, containsPath)
		}
	}
}
