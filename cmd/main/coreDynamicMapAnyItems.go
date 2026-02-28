package main

import (
	"fmt"

	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/namevalue"
)

func coreDynamicMapAnyItems() {
	fmt.Println(errcore.VarTwoNoType("varName", "ss", "varCount", 2))
	fmt.Println(errcore.MessageVarTwo("current message", "varName", "ss", "varCount", 2))
	fmt.Println(errcore.MessageVarMap(
		"current message",
		map[string]any{
			"key1": 1,
			"key2": 1,
			"key3": "",
		}))

	fmt.Println(errcore.MessageNameValues(
		"current message",
		namevalue.StringAny{
			Name:  "name1",
			Value: "nil",
		},
		namevalue.StringAny{
			Name:  "name2",
			Value: 2,
		}))

	fmt.Println("MapAnyItems")
	mapAnyItems := coredynamic.NewMapAnyItems(200)
	collection := corestr.New.Collection.Cap(100)
	collection.Adds("alim-1", "alim-2", "alim-3", "alim-4")
	mapAnyItems.Add("alim-something", collection)
	mapAnyItems.Add("alim-something2", collection)
	mapAnyItems.Add("alim-something3", collection.ConcatNew(1, "alim 5"))
	mapAnyItems.Add("alim-something4", collection)
	mapAnyItems.Add("alim-something5", collection)
	mapAnyItems.Add("alim-something6", collection)
	mapAnyItems.Add("alim-something7", collection)
	mapAnyItems.Add("alim-something8", collection)
	mapAnyItems.Add("alim-something9", collection)

	pagedItems := mapAnyItems.GetPagedCollection(2)

	for _, pagedItem := range pagedItems {
		fmt.Println(pagedItem.AllKeys())
	}

	jsonResult := mapAnyItems.JsonPtr()
	targetCollection := corestr.Empty.Collection()
	mapAnyItems.GetItemRef("alim-something3", targetCollection)
	fmt.Println("4", targetCollection)

	emptyMapAnyItems := coredynamic.EmptyMapAnyItems()
	deserializedCollection := corestr.Empty.Collection()
	keyAnyRequest := corejson.KeyAny{
		Key:    "alim-something3",
		AnyInf: deserializedCollection,
	}

	parseErr := emptyMapAnyItems.JsonParseSelfInject(jsonResult)
	newJsonResult := emptyMapAnyItems.Json()
	fmt.Println(parseErr)
	collectionJsonResult := emptyMapAnyItems.JsonResultOfKey("alim-something")

	manyItemsErr := emptyMapAnyItems.GetManyItemsRefs(keyAnyRequest)
	fmt.Println("alim-something3, err:", manyItemsErr)
	fmt.Println("\"alim-something3\":", keyAnyRequest.AnyInf)
	fmt.Println("\"alim-something3\":", deserializedCollection)

	unmarshalErr := emptyMapAnyItems.GetUsingUnmarshallManyAt(keyAnyRequest)
	fmt.Println("alim-something3, err:", unmarshalErr)
	fmt.Println("\"alim-something3\":", keyAnyRequest.AnyInf)
	fmt.Println("\"alim-something3\":", deserializedCollection)

	fmt.Println(jsonResult.JsonString())
	fmt.Println(newJsonResult.JsonString())
	fmt.Println("jsonResult == newJsonResult :", jsonResult.IsEqual(newJsonResult))
	fmt.Println(collectionJsonResult.JsonString())
	newLinkedList := corestr.Empty.LinkedList()

	newLinkedList.JsonParseSelfInject(collectionJsonResult)
	fmt.Println(newLinkedList)
	fmt.Println(mapAnyItems)

	anyCollection := coredynamic.NewAnyCollection(10)
	anyCollection.AddAnySliceFromSingleItem(pagedItems[0].AllKeys())
	fmt.Println(anyCollection)
}
