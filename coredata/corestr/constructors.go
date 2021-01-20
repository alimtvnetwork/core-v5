package corestr

import (
	"sync"

	"gitlab.com/evatix-go/core/converters"
)

// --------- Hashset starts ----------

func EmptyHashset() *Hashset {
	return NewHashset(0)
}

func NewHashset(length int) *Hashset {
	hashset := make(map[string]bool, length)

	return &Hashset{
		elementsMap:   &hashset,
		hasMapUpdated: false,
		cachedList:    nil,
		length:        length,
		isEmptySet:    true,
		Mutex:         sync.Mutex{},
	}
}

func NewHashsetWithValues(items ...string) *Hashset {
	if items == nil {
		return EmptyHashset()
	}

	return NewHashsetUsingArray(&items)
}

func NewUsingStringPointersArray(inputArray *[]*string) *Hashset {
	if inputArray == nil || *inputArray == nil {
		return NewHashset(defaultHashsetItems)
	}

	maps := converters.StringsPointersToStringBoolMap(inputArray)

	return NewHashsetUsingMap(maps)
}

func NewHashsetUsingCollection(collection *Collection) *Hashset {
	if collection == nil || collection.IsEmpty() {
		return EmptyHashset()
	}

	return NewHashsetUsingArray(collection.elements)
}

func NewHashsetUsingArray(inputArray *[]string) *Hashset {
	if inputArray == nil || *inputArray == nil {
		return EmptyHashset()
	}

	maps := converters.StringsToMap(inputArray)

	return NewHashsetUsingMap(maps)
}

func NewHashsetUsingMap(mapString *map[string]bool) *Hashset {
	if mapString == nil || *mapString == nil {
		return NewHashset(defaultHashsetItems)
	}

	length := len(*mapString)

	return &Hashset{
		elementsMap:   mapString,
		hasMapUpdated: false,
		cachedList:    nil,
		length:        length,
		isEmptySet:    length == 0,
		Mutex:         sync.Mutex{},
	}
}

// --------- Collection starts ----------

func NewCollection(capacity int) *Collection {
	collection := make([]string, 0, capacity)

	return &Collection{
		elements: &collection,
	}
}

func EmptyCollection() *Collection {
	collection := make([]string, 0, 0)

	return &Collection{
		elements: &collection,
	}
}

func NewCollectionUsingStrings(stringItems *[]string) *Collection {
	return &Collection{
		elements: stringItems,
	}
}

func NewCollectionUsingLength(len, capacity int) *Collection {
	collection := make([]string, len, capacity)

	return &Collection{
		elements: &collection,
	}
}
