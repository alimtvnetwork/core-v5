package coredynamictests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

func Test_MapAnyItems_AddAndKeys(t *testing.T) {
	tc := mapAnyItemsAddAndKeysTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
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

	actual := args.Map{
		"keyCount": len(allKeys),
		"hasAll":   hasAll,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapAnyItems_Paged(t *testing.T) {
	tc := mapAnyItemsPagedTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
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

	actual := args.Map{
		"pageCount": len(pagedItems),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapAnyItems_JsonRoundtrip(t *testing.T) {
	tc := mapAnyItemsJsonRoundtripTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
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

	actual := args.Map{
		"isEqual": jsonResult.IsEqual(newJsonResult),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapAnyItems_GetItemRef(t *testing.T) {
	tc := mapAnyItemsGetItemRefTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	key, _ := input.GetAsString("key")

	// Act
	mapItems := coredynamic.NewMapAnyItems(10)
	collection := corestr.New.Collection.Cap(5)
	collection.Adds("x", "y", "z")
	mapItems.Add(key, collection)

	target := corestr.Empty.Collection()
	mapItems.GetItemRef(key, target)

	actual := args.Map{
		"hasItems": target.HasAnyItem(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
