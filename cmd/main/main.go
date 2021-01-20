package main

import (
	"fmt"

	"gitlab.com/evatix-go/core/coredata/corestr"
)

func main() {
	// fmt.Println(osconsts.IsWindows)
	// fmt.Println(osconsts.IsUnixGroup)

	items := []string{" Alim NewHashset first,", "alim 1,", "alim next, ", "0 alim ,"}
	collection := corestr.NewCollectionUsingStrings(&items)

	collection.Resize(100)
	fmt.Println("Capacity :", collection.Capacity())

	length := 10
	newItems := make([]string, 0, length+len(items))
	newItems = append(newItems, items...)

	for i := 0; i < length; i++ {
		str := fmt.Sprintf("%d . %s", i, "alim ")
		newItems = append(newItems, str)
	}

	// fmt.Println(newItems)
	// fmt.Println("Sorted : ", collection.Add("0. Next Alim").Sorted().String())

	cmap := corestr.NewCharCollectionMap(
		10,
		5)

	var onComplete corestr.OnComplete = func(stringsMap *corestr.CharCollectionMap) {
		stringsMap.PrintLock(true)

		hashset := cmap.HashsetByStringFirstChar("a")

		json := hashset.Json()

		fmt.Println("json:\n", json.JsonString())

		h, e := hashset.NewUsingJson(json)

		fmt.Println("Data from JSON:\n", h, e)
	}

	cmap.AddStringsPtrAsyncLock(collection.ListPtr(), nil)
	cmap.AddStringsPtrAsyncLock(collection.ListPtr(), nil)
	cmap.AddStringsPtrAsyncLock(collection.ListPtr(), nil)
	cmap.AddStringsPtrAsyncLock(collection.ListPtr(), nil)
	cmap.AddStringsPtrAsyncLock(collection.ListPtr(), nil)
	cmap.AddStringsPtrAsyncLock(collection.ListPtr(), nil)
	cmap.AddStringsPtrAsyncLock(collection.ListPtr(), nil)
	cmap.AddStringsPtrAsyncLock(collection.ListPtr(), nil)
	cmap.AddStringsPtrAsyncLock(collection.ListPtr(), nil)
	cmap.AddStringsPtrAsyncLock(collection.ListPtr(), onComplete)
}
