package corepayloadtests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coredata/corepayload"
	"gitlab.com/auk-go/core/errcore"
)

func Test_PayloadWrapper_Create_Verification(t *testing.T) {
	for caseIndex, testCase := range payloadWrapperCreateTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		name, _ := input.GetAsString("name")
		id, _ := input.GetAsString("id")
		line := []byte("some payload data")

		// Act
		payload, err := corepayload.New.PayloadWrapper.Create(
			name, id, "task-type", "category", line,
		)
		errcore.HandleErr(err)

		jsonResult := payload.JsonPtr()
		hasJson := !jsonResult.IsEmpty()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			payload.Name,
			payload.Identifier,
			fmt.Sprintf("%v", hasJson),
		)
	}
}

func Test_PayloadWrapper_DeserializeRoundtrip_Verification(t *testing.T) {
	for caseIndex, testCase := range payloadWrapperDeserializeRoundtripTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		name, _ := input.GetAsString("name")
		id, _ := input.GetAsString("id")
		line := []byte("roundtrip payload bytes")

		// Act
		payload, err := corepayload.New.PayloadWrapper.Create(
			name, id, "task-type", "category", line,
		)
		errcore.HandleErr(err)

		jsonResult := payload.JsonPtr()
		restored, restoreErr := corepayload.New.PayloadWrapper.DeserializeUsingJsonResult(jsonResult)
		errcore.HandleErr(restoreErr)

		restoredJson := restored.JsonPtr()
		isEqual := jsonResult.IsEqual(*restoredJson)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			restored.Name,
			restored.Identifier,
			fmt.Sprintf("%v", isEqual),
		)
	}
}

func Test_PayloadWrapper_Clone_Verification(t *testing.T) {
	for caseIndex, testCase := range payloadWrapperCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		name, _ := input.GetAsString("name")
		id, _ := input.GetAsString("id")
		newName, _ := input.GetAsString("new_name")
		line := []byte("clone payload")

		// Act
		original, err := corepayload.New.PayloadWrapper.Create(
			name, id, "task-type", "category", line,
		)
		errcore.HandleErr(err)

		cloned, cloneErr := original.ClonePtr(true)
		errcore.HandleErr(cloneErr)

		cloned.Name = newName
		originalUnchanged := original.Name != cloned.Name

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			original.Name,
			cloned.Name,
			fmt.Sprintf("%v", originalUnchanged),
		)
	}
}

func Test_PayloadWrapper_DeserializeToMany_Verification(t *testing.T) {
	for caseIndex, testCase := range payloadWrapperDeserializeToManyTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count := input.GetAsIntDefault("count", 3)

		// Act
		wrappers := make([]*corepayload.PayloadWrapper, 0, count)
		for i := 0; i < count; i++ {
			payload, createErr := corepayload.New.PayloadWrapper.Create(
				fmt.Sprintf("item-%d", i),
				fmt.Sprintf("id-%d", i),
				"task", "cat",
				[]byte(fmt.Sprintf("data-%d", i)),
			)
			errcore.HandleErr(createErr)
			wrappers = append(wrappers, payload)
		}

		jsonSlice := corejson.Serialize.Apply(wrappers)
		jsonSlice.HandleError()

		deserialized, deserializeErr := corepayload.New.PayloadWrapper.DeserializeToMany(jsonSlice.Bytes)
		errcore.HandleErr(deserializeErr)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			fmt.Sprintf("%d", len(deserialized)),
		)
	}
}
