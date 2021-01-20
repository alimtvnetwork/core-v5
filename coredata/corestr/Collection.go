package corestr

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"gitlab.com/evatix-go/core/constants"
)

type Collection struct {
	elements *[]string
	sync.Mutex
}

func (collection *Collection) Capacity() int {
	if collection.elements == nil {
		return 0
	}

	return cap(*collection.elements)
}

func (collection *Collection) Length() int {
	if collection.elements == nil {
		return 0
	}

	return len(*collection.elements)
}

func (collection *Collection) LengthLock() int {
	collection.Lock()
	defer collection.Unlock()

	if collection.elements == nil {
		return 0
	}

	return len(*collection.elements)
}

func (collection *Collection) IsEquals(
	anotherCollection Collection,
) bool {
	return collection.IsEqualsWithSensitivePtr(
		&anotherCollection,
		true)
}

func (collection *Collection) IsEqualsPtr(
	anotherCollection *Collection,
) bool {
	return collection.IsEqualsWithSensitivePtr(
		anotherCollection,
		true)
}

func (collection *Collection) IsEqualsWithSensitivePtr(
	anotherCollection *Collection,
	isCaseSensitive bool,
) bool {
	if anotherCollection == nil {
		return false
	}

	if collection == anotherCollection {
		return true
	}

	if collection.IsEmpty() && anotherCollection.IsEmpty() {
		return true
	}

	if collection.IsEmpty() || anotherCollection.IsEmpty() {
		return false
	}

	if collection.Length() != anotherCollection.Length() {
		return false
	}

	leftItems := collection.elements
	rightItems := anotherCollection.elements

	if isCaseSensitive {
		for i, leftVal := range *leftItems {
			if leftVal != (*rightItems)[i] {
				return false
			}
		}

		return true
	}

	for i, leftVal := range *leftItems {
		if !strings.EqualFold(leftVal, (*rightItems)[i]) {
			return false
		}
	}

	return true
}

func (collection *Collection) IsEmptyLock() bool {
	collection.Lock()
	defer collection.Unlock()

	return collection.elements == nil ||
		*collection.elements == nil ||
		len(*collection.elements) == 0
}

func (collection *Collection) IsEmpty() bool {
	return collection.elements == nil ||
		*collection.elements == nil ||
		len(*collection.elements) == 0
}

func (collection *Collection) AddLock(str string) *Collection {
	collection.Lock()
	defer collection.Unlock()

	*collection.elements = append(
		*collection.elements,
		str)

	return collection
}

func (collection *Collection) Add(str string) *Collection {
	*collection.elements = append(
		*collection.elements,
		str)

	return collection
}

func (collection *Collection) AddsLock(items ...string) *Collection {
	collection.Lock()
	defer collection.Unlock()

	*collection.elements = append(
		*collection.elements,
		items...)

	return collection
}

func (collection *Collection) Adds(items ...string) *Collection {
	*collection.elements = append(
		*collection.elements,
		items...)

	return collection
}

func (collection *Collection) AddPtr(str *string) *Collection {
	*collection.elements = append(
		*collection.elements,
		*str)

	return collection
}

func (collection *Collection) AddPtrLock(str *string) *Collection {
	collection.Lock()
	defer collection.Unlock()

	*collection.elements = append(
		*collection.elements,
		*str)

	return collection
}

func (collection *Collection) AddWithWgLock(
	str string,
	group *sync.WaitGroup,
) *Collection {
	collection.Lock()
	defer collection.Unlock()

	*collection.elements = append(
		*collection.elements,
		str)

	group.Done()

	return collection
}

func (collection *Collection) AddsPtrLock(itemsPtr ...*string) *Collection {
	collection.Lock()
	defer collection.Unlock()

	for _, str := range itemsPtr {
		*collection.elements = append(
			*collection.elements,
			*str)
	}

	return collection
}

func (collection *Collection) AddStringsPtrWgLock(
	str *[]string,
	group *sync.WaitGroup,
) *Collection {
	collection.Lock()
	defer collection.Unlock()

	*collection.elements = append(
		*collection.elements,
		*str...)

	group.Done()

	return collection
}

func (collection *Collection) AddStringsPtr(str *[]string) *Collection {
	*collection.elements = append(
		*collection.elements,
		*str...)

	return collection
}

func (collection *Collection) AppendCollection(
	anotherCollection Collection,
) *Collection {
	*collection.elements = append(
		*collection.elements,
		*anotherCollection.elements...)

	return collection
}

func (collection *Collection) AppendCollectionPtr(
	anotherCollection *Collection,
) *Collection {
	*collection.elements = append(
		*collection.elements,
		*anotherCollection.elements...)

	return collection
}

func (collection *Collection) AppendCollectionsPtr(
	anotherCollectionsPtr ...*Collection,
) *Collection {
	capacitiesIncrease := 0
	for _, currentCollection := range anotherCollectionsPtr {
		if currentCollection == nil || currentCollection.IsEmpty() {
			continue
		}

		capacitiesIncrease += currentCollection.Length()
	}

	collection.AddCapacity(capacitiesIncrease)

	for _, currentCollection := range anotherCollectionsPtr {
		if currentCollection == nil || currentCollection.IsEmpty() {
			continue
		}

		*collection.elements = append(
			*collection.elements,
			*currentCollection.elements...)
	}

	return collection
}

// Continue on nil
func (collection *Collection) AppendAnysLock(anys ...interface{}) *Collection {
	for _, any := range anys {
		if any == nil {
			continue
		}

		anyStr := fmt.Sprintf(constants.SprintValueFormat, any)

		collection.Lock()
		*collection.elements = append(
			*collection.elements,
			anyStr)
		collection.Unlock()
	}

	return collection
}

// Continue on nil
func (collection *Collection) AppendAnys(anys ...interface{}) *Collection {
	for _, any := range anys {
		if any == nil {
			continue
		}

		anyStr := fmt.Sprintf(
			constants.SprintValueFormat,
			any,
		)

		*collection.elements = append(
			*collection.elements,
			anyStr)
	}

	return collection
}

// Skip on nil
func (collection *Collection) AppendAnysUsingFilter(
	filter IsStringFilter,
	anys ...interface{},
) *Collection {
	for _, any := range anys {
		if any == nil {
			continue
		}

		anyStr := fmt.Sprintf(
			constants.SprintValueFormat,
			any)

		result, isKeep := filter(anyStr)

		if !isKeep {
			continue
		}

		*collection.elements = append(
			*collection.elements,
			result)
	}

	return collection
}

// Skip on nil
func (collection *Collection) AppendAnysUsingFilterLock(
	filter IsStringFilter,
	anys ...interface{},
) *Collection {
	for _, any := range anys {
		if any == nil {
			continue
		}

		anyStr := fmt.Sprintf(constants.SprintValueFormat, any)
		result, isKeep := filter(anyStr)

		if !isKeep {
			continue
		}

		collection.Lock()
		*collection.elements = append(
			*collection.elements,
			result)
		collection.Unlock()
	}

	return collection
}

// Continue on nil
func (collection *Collection) AppendNonEmptyAnys(anys ...interface{}) *Collection {
	for _, any := range anys {
		if any == nil {
			continue
		}

		anyStr := fmt.Sprintf(constants.SprintValueFormat, any)
		if anyStr == "" {
			continue
		}

		*collection.elements = append(
			*collection.elements,
			anyStr)
	}

	return collection
}

// Skip on nil
func (collection *Collection) AddsPtr(itemsPtr ...*string) *Collection {
	for _, str := range itemsPtr {
		if str == nil {
			continue
		}

		*collection.elements = append(
			*collection.elements,
			*str)
	}

	return collection
}

func (collection *Collection) AddsNonEmptyPtr(itemsPtr ...*string) *Collection {
	for _, str := range itemsPtr {
		if str == nil || *str == "" {
			continue
		}

		*collection.elements = append(
			*collection.elements,
			*str)
	}

	return collection
}

func (collection *Collection) AddsNonEmptyPtrLock(itemsPtr ...*string) *Collection {
	for _, str := range itemsPtr {
		if str == nil || *str == "" {
			continue
		}

		collection.Lock()
		*collection.elements = append(
			*collection.elements,
			*str)
		collection.Unlock()
	}

	return collection
}

func (collection *Collection) UniqueBoolMapLock() *map[string]bool {
	collection.Lock()
	defer collection.Unlock()

	return collection.UniqueBoolMap()
}

func (collection *Collection) UniqueBoolMap() *map[string]bool {
	respectiveMap := make(
		map[string]bool,
		collection.Length())

	for _, item := range *collection.elements {
		respectiveMap[item] = true
	}

	return &respectiveMap
}

func (collection *Collection) UniqueListPtr() *[]string {
	boolMap := collection.UniqueBoolMap()
	list := make([]string, len(*boolMap))

	i := 0
	for str := range *boolMap {
		list[i] = str
		i++
	}

	return &list
}

func (collection *Collection) UniqueListPtrLock() *[]string {
	collection.Lock()
	defer collection.Unlock()

	return collection.UniqueListPtr()
}

func (collection *Collection) UniqueListLock() []string {
	collection.Lock()
	defer collection.Unlock()

	return collection.UniqueList()
}

func (collection *Collection) UniqueList() []string {
	return *collection.UniqueListPtr()
}

func (collection *Collection) List() []string {
	return *collection.elements
}

// must return a slice
func (collection *Collection) Filter(filter IsStringFilter) *[]string {
	if collection.IsEmpty() {
		return &([]string{})
	}

	list := make([]string, 0, collection.Length())

	for _, element := range *collection.elements {
		result, isKeep := filter(element)

		if isKeep {
			list = append(list, result)
		}
	}

	return &list
}

// must return a slice
func (collection *Collection) FilterLock(filter IsStringFilter) *[]string {
	elements := collection.ListPtrLock()
	length := len(*elements)

	if length == 0 {
		return elements
	}

	list := make([]string, 0, length)

	for _, element := range *elements {
		result, isKeep := filter(element)

		if isKeep {
			list = append(list, result)
		}
	}

	return &list
}

// must return a collection
func (collection *Collection) FilteredCollection(filter IsStringFilter) *Collection {
	return NewCollectionUsingStrings(collection.Filter(filter))
}

// must return a collection
func (collection *Collection) FilteredCollectionLock(filter IsStringFilter) *Collection {
	return NewCollectionUsingStrings(collection.FilterLock(filter))
}

// must return a slice
func (collection *Collection) FilterPtrLock(filterPtr IsStringPointerFilter) *[]*string {

	elements := collection.ListPtrLock()
	length := len(*elements)

	if length == 0 {
		return &([]*string{})
	}

	list := make([]*string, 0, length)

	for _, element := range *elements {
		result, isKeep := filterPtr(&element)

		if isKeep {
			list = append(list, result)
		}
	}

	return &list
}

// must return a slice
func (collection *Collection) FilterPtr(filterPtr IsStringPointerFilter) *[]*string {
	if collection.IsEmpty() {
		return &([]*string{})
	}

	list := make([]*string, 0, collection.Length())

	for _, element := range *collection.elements {
		result, isKeep := filterPtr(&element)

		if isKeep {
			list = append(list, result)
		}
	}

	return &list
}

// must return a slice
func (collection *Collection) NonEmptyListPtr() *[]string {
	if collection.IsEmpty() {
		return &([]string{})
	}

	list := make([]string, 0, collection.Length())

	for _, element := range *collection.elements {
		if element == "" {
			continue
		}

		list = append(list, element)
	}

	return &list
}

func (collection *Collection) Hashset() *Hashset {
	return NewHashsetUsingArray(collection.elements)
}

func (collection *Collection) HashsetLock() *Hashset {
	return NewHashsetUsingArray(collection.ListPtrLock())
}

// direct return pointer
func (collection *Collection) ListPtr() *[]string {
	return collection.elements
}

// returns a copy of the elements
//
// must return a slice
func (collection *Collection) ListPtrLock() *[]string {
	collection.Lock()
	defer collection.Unlock()

	if collection.elements == nil ||
		*collection.elements == nil {
		return &([]string{})
	}

	return &(*collection.elements)
}

func (collection *Collection) HasLock(str string) bool {
	collection.Lock()
	defer collection.Unlock()

	return collection.Has(str)
}

func (collection *Collection) Has(str string) bool {
	if collection.IsEmpty() {
		return false
	}

	for _, element := range *collection.elements {
		if element == str {
			return true
		}
	}

	return false
}

func (collection *Collection) HasAll(items ...string) bool {
	if collection.IsEmpty() {
		return false
	}

	for _, element := range *collection.elements {
		if !collection.IsContainsPtr(&element) {
			return false
		}
	}

	return true
}

func (collection *Collection) Sorted() *Collection {
	if collection.IsEmpty() {
		return collection
	}

	sort.Strings(*collection.elements)

	return collection
}

func (collection *Collection) SortedLock() *Collection {
	if collection.IsEmptyLock() {
		return collection
	}

	collection.Lock()
	defer collection.Unlock()

	sort.Strings(*collection.elements)

	return collection
}

func (collection *Collection) HasUsingSensitivity(str string, isCaseSensitive bool) bool {
	if isCaseSensitive {
		return collection.Has(str)
	}

	for _, element := range *collection.elements {
		if strings.EqualFold(element, str) {
			return true
		}
	}

	return false
}

func (collection *Collection) IsContainsPtr(item *string) bool {
	if item == nil || collection.IsEmpty() {
		return false
	}

	for _, element := range *collection.elements {
		if element == *item {
			return true
		}
	}

	return false
}

// nil will return false.
func (collection *Collection) GetHashsetPlusHasAll(items *[]string) (*Hashset, bool) {
	hashset := collection.Hashset()

	if items == nil || collection.IsEmpty() {
		return hashset, false
	}

	return hashset, hashset.HasAllStringsPtr(items)
}

// nil will return false.
func (collection *Collection) IsContainsAllPtr(items *[]string) bool {
	if items == nil {
		return false
	}

	if collection.IsEmpty() {
		return false
	}

	for _, item := range *items {
		if !collection.IsContainsPtr(&item) {
			return false
		}
	}

	return true
}

// nil will return false.
func (collection *Collection) IsContainsAll(items ...string) bool {
	if items == nil {
		return false
	}

	return collection.IsContainsAllPtr(&items)
}

// nil will return false.
func (collection *Collection) IsContainsAllLock(items ...string) bool {
	collection.Lock()
	defer collection.Unlock()

	if items == nil {
		return false
	}

	return collection.IsContainsAllPtr(&items)
}

func (collection *Collection) CharCollectionMap() *CharCollectionMap {
	length := collection.Length()
	lengthByFourBestGuess := length / 4
	runeMap := NewCharCollectionMap(
		length,
		lengthByFourBestGuess)

	for _, item := range *collection.elements {
		runeMap.AddStringPtr(&item)
	}

	return runeMap
}

func (collection *Collection) String() string {
	if collection.IsEmpty() {
		return commonJoiner + NoElements
	}

	return commonJoiner +
		strings.Join(
			*collection.elements,
			commonJoiner)
}

func (collection *Collection) StringLock() string {
	if collection.IsEmptyLock() {
		return commonJoiner + NoElements
	}

	collection.Lock()
	defer collection.Unlock()

	return commonJoiner +
		strings.Join(
			*collection.elements,
			commonJoiner)
}

func (collection *Collection) Join(
	separator string,
) string {
	return strings.Join(*collection.elements, separator)
}

func (collection *Collection) AddCapacity(
	capacities ...int,
) *Collection {
	if capacities == nil || len(capacities) == 0 {
		return collection
	}

	currentCapacity := collection.Capacity()

	for _, capacity := range capacities {
		currentCapacity += capacity
	}

	return collection.Resize(currentCapacity)
}

// Only resize if capacity is bigger than the current one
func (collection *Collection) Resize(
	newCapacity int,
) *Collection {
	capacity := collection.Capacity()
	if capacity >= newCapacity {
		return collection
	}

	newItems := make([]string, collection.Length(), newCapacity)
	copy(newItems, *collection.elements)

	collection.elements = &newItems

	return collection
}

func (collection *Collection) Joins(
	separator string,
	items ...string,
) string {
	if items == nil || len(items) == 0 {
		return strings.Join(items, separator)
	}

	newItems := make([]string, 0, collection.Length()+len(items))
	copy(newItems, *collection.elements)

	newItems = append(newItems, items...)

	return strings.Join(newItems, separator)
}
