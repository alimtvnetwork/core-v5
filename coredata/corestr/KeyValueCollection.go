package corestr

import (
	"gitlab.com/evatix-go/core/internal/strutilinternal"
)

type KeyValueCollection struct {
	Items []*KeyValuePair
}

func (it *KeyValueCollection) Count() int {
	return it.Length()
}

func (it *KeyValueCollection) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *KeyValueCollection) LastIndex() int {
	return it.Length() - 1
}

func (it *KeyValueCollection) HasIndex(
	index int,
) bool {
	return it.LastIndex() >= index
}

func (it *KeyValueCollection) Strings() []string {
	if it.IsEmpty() {
		return []string{}
	}

	slice := make([]string, it.Length())

	for i, keyVal := range it.Items {
		slice[i] = keyVal.String()
	}

	return slice
}

func (it *KeyValueCollection) StringsUsingFormat(
	format string,
) []string {
	if it.IsEmpty() {
		return []string{}
	}

	slice := make([]string, it.Length())

	for i, keyVal := range it.Items {
		slice[i] = keyVal.FormatString(format)
	}

	return slice
}

func (it *KeyValueCollection) String() string {
	return strutilinternal.AnyToString(it.Strings())
}

func (it *KeyValueCollection) Length() int {
	if it == nil {
		return 0
	}

	return len(it.Items)
}

func (it *KeyValueCollection) IsEmpty() bool {
	return it.Length() == 0
}

func (it *KeyValueCollection) Add(key, val string) *KeyValueCollection {
	it.Items = append(it.Items, &KeyValuePair{
		Key:   key,
		Value: val,
	})

	return it
}

func (it *KeyValueCollection) Adds(keyValues ...KeyValuePair) *KeyValueCollection {
	if len(keyValues) == 0 {
		return it
	}

	for _, keyVal := range keyValues {
		it.Items = append(it.Items, &KeyValuePair{
			Key:   keyVal.Key,
			Value: keyVal.Value,
		})
	}

	return it
}

func (it *KeyValueCollection) AddMap(
	inputMap map[string]string,
) *KeyValueCollection {
	if inputMap == nil || len(inputMap) == 0 {
		return it
	}

	for key, val := range inputMap {
		it.Items = append(it.Items, &KeyValuePair{
			Key:   key,
			Value: val,
		})
	}

	return it
}

func (it *KeyValueCollection) AddHashsetMap(
	inputMap map[string]bool,
) *KeyValueCollection {
	if inputMap == nil || len(inputMap) == 0 {
		return it
	}

	for key := range inputMap {
		it.Items = append(it.Items, &KeyValuePair{
			Key:   key,
			Value: key,
		})
	}

	return it
}

func (it *KeyValueCollection) AddHashset(
	inputHashset *Hashset,
) *KeyValueCollection {
	if inputHashset == nil || inputHashset.IsEmpty() {
		return it
	}

	for key := range *inputHashset.items {
		it.Items = append(it.Items, &KeyValuePair{
			Key:   key,
			Value: key,
		})
	}

	return it
}

func (it *KeyValueCollection) AddsHashmap(
	hashmap *Hashmap,
) *KeyValueCollection {
	if hashmap == nil || hashmap.IsEmpty() {
		return it
	}

	for key, val := range *hashmap.items {
		it.Items = append(it.Items, &KeyValuePair{
			Key:   key,
			Value: val,
		})
	}

	return it
}

func (it *KeyValueCollection) Hashmap() *Hashmap {
	length := it.Length()
	hashmap := NewHashmap(length)

	if length == 0 {
		return hashmap
	}

	for _, keyVal := range it.Items {
		hashmap.AddOrUpdate(keyVal.Key, keyVal.Value)
	}

	return hashmap
}

func (it *KeyValueCollection) Map() *map[string]string {
	hashmap := it.Hashmap()

	return hashmap.items
}

func (it *KeyValueCollection) AddsHashmaps(
	hashmaps ...*Hashmap,
) *KeyValueCollection {
	if hashmaps == nil || len(hashmaps) == 0 {
		return it
	}

	for _, hashmap := range hashmaps {
		it.AddsHashmap(hashmap)
	}

	return it
}
