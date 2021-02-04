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
	// fmt.Println("Capacity :", collection.Capacity())

	length := 10
	newItems := make([]string, 0, length+len(items))
	newItems = append(newItems, items...)

	for i := 0; i < length; i++ {
		str := fmt.Sprintf("%d . %s", i, "alim ")
		newItems = append(newItems, str)
	}

	// fmt.Println(newItems)
	// fmt.Println("SortedAsc : ", collection.Add("0. Next Alim").SortedAsc().String())

	cmap := corestr.NewCharCollectionMap(
		10,
		5)

	var onComplete corestr.OnCompleteCharCollectionMap = func(stringsMap *corestr.CharCollectionMap) {
		linkedList := corestr.NewLinkedList().
			AddStringsPtr(collection.ListPtr())

		fmt.Println("before:\n", linkedList)

		node := linkedList.SafeIndexAt(2)

		fmt.Println("node :\n", node)

		finalNode := node.AddNext(linkedList, "alim +++").AddNext(linkedList, "alim ++ 2")

		fmt.Println("after:\n", linkedList)

		finalNode.AddCollectionToNode(linkedList, true, collection)
		linkedList.RemoveAll()
		fmt.Println("after items:\n", linkedList, "\nlen:\n", linkedList.Length())
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
