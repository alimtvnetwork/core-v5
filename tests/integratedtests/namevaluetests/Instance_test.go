package namevaluetests

import (
	"fmt"
	"strings"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/namevalue"
)

// region StringAny (backward-compat) tests

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

// endregion

// region StringString tests

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

// endregion

// region StringInt tests

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

// endregion

// region StringMapAny tests

func Test_StringMapAny_String_Verification(t *testing.T) {
	for caseIndex, testCase := range stringMapAnyTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		name, _ := input.GetAsString("name")

		var mapVal map[string]any
		switch name {
		case "config":
			mapVal = map[string]any{"host": "localhost", "port": 8080}
		case "empty":
			mapVal = map[string]any{}
		default:
			mapVal = nil
		}

		// Act
		instance := namevalue.StringMapAny{
			Name:  name,
			Value: mapVal,
		}
		result := instance.String()
		isNotEmpty := fmt.Sprintf("%v", result != "")
		hasName := fmt.Sprintf("%v", strings.Contains(result, name))

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, isNotEmpty, hasName)
	}
}

// endregion

// region StringMapString tests

func Test_StringMapString_String_Verification(t *testing.T) {
	for caseIndex, testCase := range stringMapStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		name, _ := input.GetAsString("name")

		var mapVal map[string]string
		switch name {
		case "headers":
			mapVal = map[string]string{"Content-Type": "application/json"}
		default:
			mapVal = nil
		}

		// Act
		instance := namevalue.StringMapString{
			Name:  name,
			Value: mapVal,
		}
		result := instance.String()
		isNotEmpty := fmt.Sprintf("%v", result != "")
		hasName := fmt.Sprintf("%v", strings.Contains(result, name))

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, isNotEmpty, hasName)
	}
}

// endregion

// region Dispose tests

func Test_Generic_Dispose_Verification(t *testing.T) {
	for caseIndex, testCase := range genericDisposeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		typeStr, _ := input.GetAsString("type")
		name, _ := input.GetAsString("name")

		var resultName string
		var resultValue string

		// Act
		switch typeStr {
		case "stringany":
			value, _ := input.Get("value")
			inst := &namevalue.StringAny{Name: name, Value: value}
			inst.Dispose()
			resultName = inst.Name
			resultValue = fmt.Sprintf("%v", inst.Value == nil)
		case "stringstring":
			value, _ := input.GetAsString("value")
			inst := &namevalue.StringString{Name: name, Value: value}
			inst.Dispose()
			resultName = inst.Name
			resultValue = inst.Value
		case "stringint":
			inst := &namevalue.StringInt{Name: name, Value: 42}
			inst.Dispose()
			resultName = inst.Name
			resultValue = fmt.Sprintf("%d", inst.Value)
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, resultName, resultValue)
	}
}

// endregion

// region JsonString tests

func Test_Generic_JsonString_Verification(t *testing.T) {
	for caseIndex, testCase := range genericJsonStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		name, _ := input.GetAsString("name")
		value, _ := input.Get("value")

		// Act
		var jsonStr string
		switch value.(type) {
		case int:
			inst := namevalue.StringInt{Name: name, Value: value.(int)}
			jsonStr = inst.JsonString()
		default:
			valStr, _ := value.(string)
			inst := namevalue.StringAny{Name: name, Value: valStr}
			jsonStr = inst.JsonString()
		}
		isNotEmpty := fmt.Sprintf("%v", jsonStr != "")
		containsName := fmt.Sprintf("%v", strings.Contains(jsonStr, name))

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, isNotEmpty, containsName)
	}
}

// endregion

// region Collection tests

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

// endregion

// region Chmod integration tests

func Test_Chmod_Integration_Verification(t *testing.T) {
	for caseIndex, testCase := range chmodIntegrationTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		name, hasName := input.GetAsString("name")

		// Act
		if !hasName {
			// Negative: empty slice test
			result := errcore.VarNameValues()

			// Assert
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
			// MessageNameValues test
			result := errcore.MessageNameValues(message, nv)
			containsMessage := fmt.Sprintf("%v", strings.Contains(result, message))
			containsPath := fmt.Sprintf("%v", strings.Contains(result, path))

			// Assert
			testCase.ShouldBeEqual(t, caseIndex, containsMessage, containsPath)
		} else {
			// VarNameValues test
			result := errcore.VarNameValues(nv)
			isNotEmpty := fmt.Sprintf("%v", result != "")
			containsPath := fmt.Sprintf("%v", strings.Contains(result, path))

			// Assert
			testCase.ShouldBeEqual(t, caseIndex, isNotEmpty, containsPath)
		}
	}
}

// endregion
