package coredynamictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/errcore"
)

func Test_MapAnyItems_AddAndKeys_Verification(t *testing.T) {
	for caseIndex, testCase := range mapAnyItemsAddAndKeysTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		capacity := input.GetAsIntDefault("capacity", 10)
		keys := input["keys"].([]string)

		// Act
		mapItems := coredynamic.NewMapAnyItems(capacity)
		collection := corestr.New.Collection.Cap(10)
		collection.Adds("a", "b", "c")

		for _, key := range keys {
			mapItems.Add(key, collection)
		}

		allKeys := mapItems.AllKeys()
		hasAll := true
		for _, key := range keys {
			found := false
			for _, k := range allKeys {
				if k == key {
					found = true
					break
				}
			}
			if !found {
				hasAll = false
				break
			}
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			fmt.Sprintf("%d", len(allKeys)),
			fmt.Sprintf("%v", hasAll),
		)
	}
}

func Test_MapAnyItems_Paged_Verification(t *testing.T) {
	for caseIndex, testCase := range mapAnyItemsPagedTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		itemCount := input.GetAsIntDefault("itemCount", 9)
		pageSize := input.GetAsIntDefault("pageSize", 2)

		// Act
		mapItems := coredynamic.NewMapAnyItems(itemCount + 5)
		collection := corestr.New.Collection.Cap(5)
		collection.Adds("a", "b")

		for i := 0; i < itemCount; i++ {
			mapItems.Add(fmt.Sprintf("key-%d", i), collection)
		}

		pagedItems := mapItems.GetPagedCollection(pageSize)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			fmt.Sprintf("%d", len(pagedItems)),
		)
	}
}

func Test_MapAnyItems_JsonRoundtrip_Verification(t *testing.T) {
	for caseIndex, testCase := range mapAnyItemsJsonRoundtripTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		itemCount := input.GetAsIntDefault("itemCount", 4)

		// Act
		mapItems := coredynamic.NewMapAnyItems(itemCount + 5)
		collection := corestr.New.Collection.Cap(5)
		collection.Adds("val1", "val2")

		for i := 0; i < itemCount; i++ {
			mapItems.Add(fmt.Sprintf("item-%d", i), collection)
		}

		jsonResult := mapItems.JsonPtr()
		restored := coredynamic.EmptyMapAnyItems()
		parseErr := restored.JsonParseSelfInject(jsonResult)
		errcore.HandleErr(parseErr)

		newJsonResult := restored.Json()
		isEqual := jsonResult.IsEqual(newJsonResult)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			fmt.Sprintf("%v", isEqual),
		)
	}
}

func Test_MapAnyItems_GetItemRef_Verification(t *testing.T) {
	for caseIndex, testCase := range mapAnyItemsGetItemRefTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		key, _ := input.GetAsString("key")

		// Act
		mapItems := coredynamic.NewMapAnyItems(10)
		collection := corestr.New.Collection.Cap(5)
		collection.Adds("x", "y", "z")
		mapItems.Add(key, collection)

		target := corestr.Empty.Collection()
		mapItems.GetItemRef(key, target)
		hasItems := target.HasAnyItem()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			fmt.Sprintf("%v", hasItems),
		)
	}
}
