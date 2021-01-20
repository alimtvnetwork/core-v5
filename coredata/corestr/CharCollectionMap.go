package corestr

import (
	"fmt"
	"strings"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreindexes"
)

type CharCollectionMap struct {
	elementsMap            *map[byte]*Collection
	selfCollectionCapacity int
	sync.Mutex
}

// CharCollectionMap.selfCollectionCapacity, capacity minimum 10 will be set if lower than 10 is given.
//
// For lower than 5 use the EmptyCharCollectionMap collection definition.
func NewCharCollectionMap(
	capacity, selfCollectionCapacity int,
) *CharCollectionMap {
	if capacity < constants.ArbitraryCapacity10 {
		capacity = constants.ArbitraryCapacity10
	}

	mapElements := make(map[byte]*Collection, capacity)

	if selfCollectionCapacity < constants.ArbitraryCapacity10 {
		selfCollectionCapacity = constants.ArbitraryCapacity10
	}

	return &CharCollectionMap{
		elementsMap:            &mapElements,
		selfCollectionCapacity: selfCollectionCapacity,
	}
}

// selfCollectionCapacity = 0
func EmptyCharCollectionMap() *CharCollectionMap {
	mapElements := make(map[byte]*Collection, 0)

	return &CharCollectionMap{
		elementsMap:            &mapElements,
		selfCollectionCapacity: 0,
	}
}

func (charCollectionMap *CharCollectionMap) GetChar(
	str string,
) byte {
	if str != "" {
		return str[coreindexes.First]
	}

	return emptyChar
}

func (charCollectionMap *CharCollectionMap) GetCharOfPtr(
	str *string,
) byte {
	if str == nil || *str == "" {
		return emptyChar
	}

	return (*str)[coreindexes.First]
}

func (charCollectionMap *CharCollectionMap) GetCharsPtrGroups(
	items *[]string,
) *CharCollectionMap {
	if items == nil || *items == nil {
		return StaticEmptyCharCollectionMapPtr
	}

	length := len(*items)

	if length == 0 {
		return nil
	}

	collectionMap := NewCharCollectionMap(
		length,
		length/3)

	return collectionMap.AddStringsPtr(items)
}

func (charCollectionMap *CharCollectionMap) GetMap() *map[byte]*Collection {
	return charCollectionMap.elementsMap
}

// Sends a copy of items
func (charCollectionMap *CharCollectionMap) GetMapLock() *map[byte]*Collection {
	charCollectionMap.Lock()
	defer charCollectionMap.Unlock()

	if charCollectionMap.IsEmpty() {
		return &(map[byte]*Collection{})
	}

	return &(*charCollectionMap.elementsMap)
}

func (charCollectionMap *CharCollectionMap) SummaryStringLock() string {
	length := charCollectionMap.LengthLock()
	collectionOfCollection := make(
		[]string,
		length+1)

	collectionOfCollection[coreindexes.First] = fmt.Sprintf(
		summaryOfCharCollectionMapLengthFormat,
		charCollectionMap,
		length)

	i := 1
	for key, collection := range *charCollectionMap.GetMapLock() {
		collectionOfCollection[i] = fmt.Sprintf(
			charCollectionMapSingleItemFormat,
			string(key),
			collection.LengthLock())

		i++
	}

	return strings.Join(
		collectionOfCollection,
		constants.EmptyString)
}

func (charCollectionMap *CharCollectionMap) SummaryString() string {
	collectionOfCollection := make(
		[]string,
		charCollectionMap.Length()+1)

	collectionOfCollection[coreindexes.First] = fmt.Sprintf(
		summaryOfCharCollectionMapLengthFormat,
		charCollectionMap,
		charCollectionMap.Length())

	i := 1
	for key, collection := range *charCollectionMap.elementsMap {
		collectionOfCollection[i] = fmt.Sprintf(
			charCollectionMapSingleItemFormat,
			string(key),
			collection.Length())

		i++
	}

	return strings.Join(
		collectionOfCollection,
		constants.EmptyString)
}

func (charCollectionMap *CharCollectionMap) String() string {
	collectionOfCollection := make(
		[]string,
		charCollectionMap.Length()*2+1)

	collectionOfCollection[coreindexes.First] =
		charCollectionMap.SummaryString()

	i := 1
	for key, collection := range *charCollectionMap.elementsMap {
		collectionOfCollection[i] = fmt.Sprintf(
			charCollectionMapLengthFormat,
			string(key))

		i++
		collectionOfCollection[i] = collection.String()
		i++
	}

	return strings.Join(
		collectionOfCollection,
		constants.EmptyString)
}

func (charCollectionMap *CharCollectionMap) StringLock() string {
	collectionOfCollection := make(
		[]string,
		charCollectionMap.LengthLock()*2+1)

	collectionOfCollection[coreindexes.First] =
		charCollectionMap.SummaryStringLock()

	i := 1
	for key, collection := range *charCollectionMap.GetMapLock() {
		collectionOfCollection[i] = fmt.Sprintf(
			charCollectionMapLengthFormat,
			string(key))

		i++
		collectionOfCollection[i] =
			collection.StringLock()
		i++
	}

	return strings.Join(
		collectionOfCollection,
		constants.EmptyString)
}

func (charCollectionMap *CharCollectionMap) Print(isPrint bool) {
	if !isPrint {
		return
	}

	fmt.Println(
		charCollectionMap.String(),
	)
}

func (charCollectionMap *CharCollectionMap) PrintLock(isPrint bool) {
	if !isPrint {
		return
	}

	fmt.Println(
		charCollectionMap.StringLock(),
	)
}

func (charCollectionMap *CharCollectionMap) IsEmpty() bool {
	return charCollectionMap.elementsMap == nil ||
		*charCollectionMap.elementsMap == nil ||
		len(*charCollectionMap.elementsMap) == 0
}

func (charCollectionMap *CharCollectionMap) IsEmptyLock() bool {
	charCollectionMap.Lock()
	defer charCollectionMap.Unlock()

	return charCollectionMap.
		elementsMap == nil ||
		*charCollectionMap.elementsMap == nil ||
		len(*charCollectionMap.elementsMap) == 0
}

// Get the char of the string given and get the length of how much is there.
func (charCollectionMap *CharCollectionMap) LengthOfCollectionFromFirstChar(
	str string,
) int {
	char := charCollectionMap.GetChar(str)

	collection, has := (*charCollectionMap.elementsMap)[char]

	if has {
		return collection.Length()
	}

	return 0
}

func (charCollectionMap *CharCollectionMap) Has(
	str string,
) bool {
	if charCollectionMap.IsEmpty() {
		return false
	}

	char := charCollectionMap.
		GetChar(str)

	collection, has := (*charCollectionMap.elementsMap)[char]

	if has {
		return collection.Has(str)
	}

	return false
}

func (charCollectionMap *CharCollectionMap) HasWithCollection(
	str string,
) (bool, *Collection) {
	if charCollectionMap.IsEmpty() {
		return false, StaticEmptyCollectionPtr
	}

	char := charCollectionMap.
		GetChar(str)

	collection, has := (*charCollectionMap.elementsMap)[char]

	if has {
		return collection.Has(str), collection
	}

	return false, StaticEmptyCollectionPtr
}

func (charCollectionMap *CharCollectionMap) HasWithCollectionLock(
	str string,
) (bool, *Collection) {
	charCollectionMap.Lock()
	defer charCollectionMap.Unlock()

	if charCollectionMap.IsEmpty() {
		return false, StaticEmptyCollectionPtr
	}

	char := charCollectionMap.
		GetChar(str)

	collection, has := (*charCollectionMap.elementsMap)[char]

	if has {
		return collection.HasLock(str), collection
	}

	return false, StaticEmptyCollectionPtr
}

func (charCollectionMap *CharCollectionMap) LengthOf(char byte) int {
	if charCollectionMap.IsEmpty() {
		return 0
	}

	collection, has := (*charCollectionMap.elementsMap)[char]

	if has {
		return collection.Length()
	}

	return 0
}

func (charCollectionMap *CharCollectionMap) LengthOfLock(char byte) int {
	charCollectionMap.Lock()
	defer charCollectionMap.Unlock()

	if charCollectionMap.IsEmpty() {
		return 0
	}

	collection, has := (*charCollectionMap.elementsMap)[char]

	if has {
		return collection.Length()
	}

	return 0
}

func (charCollectionMap *CharCollectionMap) Length() int {
	if charCollectionMap.
		elementsMap == nil ||
		*charCollectionMap.elementsMap == nil {
		return 0
	}

	return len(*charCollectionMap.elementsMap)
}

func (charCollectionMap *CharCollectionMap) LengthLock() int {
	charCollectionMap.Lock()
	defer charCollectionMap.Unlock()

	if charCollectionMap.
		elementsMap == nil ||
		*charCollectionMap.elementsMap == nil {
		return 0
	}

	return len(*charCollectionMap.elementsMap)
}

func (charCollectionMap *CharCollectionMap) IsEqualsPtrLock(
	another *CharCollectionMap,
) bool {
	charCollectionMap.Lock()
	defer charCollectionMap.Unlock()

	return charCollectionMap.IsEqualsWithCaseSensitivityPtr(
		another,
		true)
}

func (charCollectionMap *CharCollectionMap) IsEqualsPtr(
	another *CharCollectionMap,
) bool {
	return charCollectionMap.IsEqualsWithCaseSensitivityPtr(
		another,
		true)
}

func (charCollectionMap *CharCollectionMap) IsEqualsWithCaseSensitivityPtrLock(
	another *CharCollectionMap,
	isCaseSensitive bool,
) bool {
	charCollectionMap.Lock()
	defer charCollectionMap.Unlock()

	return charCollectionMap.IsEqualsWithCaseSensitivityPtr(
		another,
		isCaseSensitive)
}

func (charCollectionMap *CharCollectionMap) IsEqualsWithCaseSensitivityPtr(
	another *CharCollectionMap,
	isCaseSensitive bool,
) bool {
	if another == nil {
		return false
	}

	if another == charCollectionMap {
		return true
	}

	if another.IsEmpty() && charCollectionMap.IsEmpty() {
		return true
	}

	if another.IsEmpty() || charCollectionMap.IsEmpty() {
		return false
	}

	if another.Length() != charCollectionMap.Length() {
		return false
	}

	leftMap := charCollectionMap.elementsMap
	rightMap := another.elementsMap

	for key, collection := range *leftMap {
		rCollection, has := (*rightMap)[key]

		if !has {
			return false
		}

		if !rCollection.IsEqualsWithSensitivePtr(
			collection,
			isCaseSensitive) {
			return false
		}
	}

	return true
}

func (charCollectionMap *CharCollectionMap) AddLock(
	str string,
) *CharCollectionMap {
	char := charCollectionMap.GetChar(str)

	charCollectionMap.Lock()
	collection, has := (*charCollectionMap.
		elementsMap)[char]
	charCollectionMap.Unlock()

	if has {
		collection.AddLock(str)

		return charCollectionMap
	}

	newCollection := NewCollection(charCollectionMap.selfCollectionCapacity)
	newCollection.Add(str)

	charCollectionMap.Lock()
	(*charCollectionMap.elementsMap)[char] = newCollection
	charCollectionMap.Unlock()

	return charCollectionMap
}

func (charCollectionMap *CharCollectionMap) Add(
	str string,
) *CharCollectionMap {
	char := charCollectionMap.GetChar(str)

	collection, has := (*charCollectionMap.
		elementsMap)[char]

	if has {
		collection.Add(str)
	}

	newCollection := NewCollection(charCollectionMap.selfCollectionCapacity)
	newCollection.Add(str)
	(*charCollectionMap.
		elementsMap)[char] = newCollection

	return charCollectionMap
}

func (charCollectionMap *CharCollectionMap) AddStringPtr(
	str *string,
) *CharCollectionMap {
	char := charCollectionMap.GetCharOfPtr(str)

	collection, has := (*charCollectionMap.
		elementsMap)[char]

	if has {
		collection.AddPtr(str)
	}

	newCollection := NewCollection(charCollectionMap.selfCollectionCapacity)
	newCollection.AddPtr(str)
	(*charCollectionMap.
		elementsMap)[char] = newCollection

	return charCollectionMap
}

func (charCollectionMap *CharCollectionMap) AddStringPtrLock(
	str *string,
) *CharCollectionMap {
	defer charCollectionMap.Unlock()
	char := charCollectionMap.GetCharOfPtr(str)

	charCollectionMap.Lock()
	collection, has := (*charCollectionMap.
		elementsMap)[char]
	charCollectionMap.Unlock()

	if has {
		collection.AddPtrLock(str)

		return charCollectionMap
	}

	newCollection := NewCollection(charCollectionMap.selfCollectionCapacity)
	newCollection.AddPtr(str)

	charCollectionMap.Lock()
	(*charCollectionMap.
		elementsMap)[char] = newCollection
	charCollectionMap.Unlock()

	return charCollectionMap
}

// Assuming all items starts with same chars
func (charCollectionMap *CharCollectionMap) AddSameStartingCharItems(
	char byte,
	allItemsWithSameChar *[]string,
) *CharCollectionMap {
	if allItemsWithSameChar == nil ||
		*allItemsWithSameChar == nil ||
		len(*allItemsWithSameChar) == 0 {
		return charCollectionMap
	}

	values, has := (*charCollectionMap.
		elementsMap)[char]

	if has {
		values.AddStringsPtr(allItemsWithSameChar)

		return charCollectionMap
	}

	(*charCollectionMap.
		elementsMap)[char] =
		NewCollectionUsingStrings(allItemsWithSameChar)

	return charCollectionMap
}

func (charCollectionMap *CharCollectionMap) AddPtrStringsLock(
	simpleStrings *[]*string,
) *CharCollectionMap {
	if simpleStrings == nil ||
		*simpleStrings == nil ||
		len(*simpleStrings) == 0 {
		return charCollectionMap
	}

	for _, item := range *simpleStrings {
		foundCollection := charCollectionMap.GetCollectionLock(
			*item, true)

		foundCollection.AddPtrLock(item)
	}

	return charCollectionMap
}

func (charCollectionMap *CharCollectionMap) AddStringsPtrAsyncLock(
	largeStringsCollection *[]string,
	onComplete OnComplete,
) *CharCollectionMap {
	if largeStringsCollection == nil ||
		*largeStringsCollection == nil {
		return charCollectionMap
	}

	length := len(*largeStringsCollection)

	if length == 0 {
		return charCollectionMap
	}

	isListIsTooLargeAndHasExistingData := length > RegularCollectionEfficiencyLimit &&
		charCollectionMap.Length() > DoubleLimit

	if isListIsTooLargeAndHasExistingData {
		return charCollectionMap.
			efficientAddOfLargeItems(
				largeStringsCollection,
				onComplete)
	}

	wg := &sync.WaitGroup{}
	wg.Add(length)

	for _, item := range *largeStringsCollection {
		foundCollection := charCollectionMap.GetCollectionLock(
			item,
			true)

		go foundCollection.AddWithWgLock(
			item,
			wg,
		)
	}

	wg.Wait()

	if onComplete != nil {
		onComplete(charCollectionMap)
	}

	return charCollectionMap
}

func (charCollectionMap *CharCollectionMap) efficientAddOfLargeItems(
	largeStringsCollection *[]string, onComplete OnComplete,
) *CharCollectionMap {
	allCharsMap := charCollectionMap.
		GetCharsPtrGroups(largeStringsCollection)

	wg := &sync.WaitGroup{}
	wg.Add(allCharsMap.Length())

	for key, collection := range *allCharsMap.elementsMap {
		foundCollection := charCollectionMap.GetCollectionLock(
			string(key),
			true)

		go foundCollection.AddStringsPtrWgLock(
			collection.elements,
			wg,
		)
	}

	wg.Wait()

	if onComplete != nil {
		onComplete(charCollectionMap)
	}

	return charCollectionMap
}

func (charCollectionMap *CharCollectionMap) AddStringsPtr(
	items *[]string,
) *CharCollectionMap {
	if items == nil ||
		*items == nil ||
		len(*items) == 0 {
		return charCollectionMap
	}

	for _, item := range *items {
		charCollectionMap.AddStringPtr(&item)
	}

	return charCollectionMap
}

func (charCollectionMap *CharCollectionMap) AddStrings(
	items ...string,
) *CharCollectionMap {
	if items == nil ||
		len(items) == 0 {
		return charCollectionMap
	}

	for _, item := range items {
		charCollectionMap.AddStringPtr(&item)
	}

	return charCollectionMap
}

func (charCollectionMap *CharCollectionMap) GetCollection(
	strFirstChar string,
	isAddNewOnEmpty bool,
) *Collection {
	char := charCollectionMap.GetChar(strFirstChar)

	collection, has := (*charCollectionMap.elementsMap)[char]

	if has {
		return collection
	}

	if isAddNewOnEmpty {
		newCollection := NewCollection(charCollectionMap.selfCollectionCapacity)
		(*charCollectionMap.elementsMap)[char] = newCollection

		return newCollection
	}

	return nil
}

func (charCollectionMap *CharCollectionMap) GetCollectionLock(
	strFirstChar string,
	isAddNewOnEmpty bool,
) *Collection {
	charCollectionMap.Lock()
	defer charCollectionMap.Unlock()

	return charCollectionMap.GetCollection(
		strFirstChar,
		isAddNewOnEmpty)
}

func (charCollectionMap *CharCollectionMap) AddCollection(
	str string,
	stringsWithSameStartChar *Collection,
) *Collection {
	isNilOrEmptyCollectionGiven := stringsWithSameStartChar == nil ||
		stringsWithSameStartChar.IsEmpty()

	foundCollection := charCollectionMap.GetCollection(
		str,
		false)

	has := foundCollection != nil
	isAddToCollection := has && !isNilOrEmptyCollectionGiven
	hasCollectionHoweverNothingToAdd := has && isNilOrEmptyCollectionGiven

	if isAddToCollection {
		foundCollection.AddStringsPtr(stringsWithSameStartChar.elements)

		return foundCollection
	} else if hasCollectionHoweverNothingToAdd {
		return foundCollection
	}

	char := charCollectionMap.GetChar(str)

	if isNilOrEmptyCollectionGiven {
		// create new
		newCollection := NewCollection(
			charCollectionMap.selfCollectionCapacity)
		(*charCollectionMap.elementsMap)[char] = newCollection

		return newCollection
	}

	// elements exist or stringsWithSameStartChar exists
	(*charCollectionMap.elementsMap)[char] =
		stringsWithSameStartChar

	return stringsWithSameStartChar
}

func (charCollectionMap *CharCollectionMap) AddCollectionItems(
	collectionWithDiffStarts *Collection,
) *CharCollectionMap {
	if collectionWithDiffStarts == nil ||
		collectionWithDiffStarts.IsEmpty() {
		return charCollectionMap
	}

	charCollectionMap.AddStringsPtr(collectionWithDiffStarts.elements)

	return charCollectionMap
}

func (charCollectionMap *CharCollectionMap) AddCollectionItemsAsyncLock(
	collectionWithDiffStarts *Collection,
	onComplete OnComplete,
) *CharCollectionMap {
	if collectionWithDiffStarts == nil ||
		collectionWithDiffStarts.IsEmpty() {
		return charCollectionMap
	}

	go charCollectionMap.AddStringsPtrAsyncLock(
		collectionWithDiffStarts.elements,
		onComplete)

	return charCollectionMap
}

func (charCollectionMap *CharCollectionMap) AddCollectionLock(
	str string,
	stringsWithSameStartChar *Collection,
) *Collection {
	isNilOrEmptyCollectionGiven := stringsWithSameStartChar == nil ||
		stringsWithSameStartChar.IsEmpty()

	foundCollection := charCollectionMap.GetCollectionLock(
		str,
		false)
	has := foundCollection != nil
	isAddToCollection := has && !isNilOrEmptyCollectionGiven
	hasCollectionHoweverNothingToAdd := has && isNilOrEmptyCollectionGiven

	if isAddToCollection {
		foundCollection.AddStringsPtr(stringsWithSameStartChar.elements)

		return foundCollection
	} else if hasCollectionHoweverNothingToAdd {
		return foundCollection
	}

	char := charCollectionMap.GetChar(str)

	if isNilOrEmptyCollectionGiven {
		// create new
		newCollection := NewCollection(
			charCollectionMap.selfCollectionCapacity)
		charCollectionMap.Lock()
		(*charCollectionMap.elementsMap)[char] = newCollection
		charCollectionMap.Unlock()

		return newCollection
	}

	// elements exist or stringsWithSameStartChar exists
	charCollectionMap.Lock()
	(*charCollectionMap.elementsMap)[char] =
		stringsWithSameStartChar
	charCollectionMap.Unlock()

	return stringsWithSameStartChar
}

func (charCollectionMap *CharCollectionMap) GetCollectionByChar(
	char byte,
) *Collection {
	return (*charCollectionMap.elementsMap)[char]
}

func (charCollectionMap *CharCollectionMap) HashsetByChar(
	char byte,
) *Hashset {
	collection := (*charCollectionMap.elementsMap)[char]

	return NewHashsetUsingCollection(collection)
}

func (charCollectionMap *CharCollectionMap) HashsetByCharLock(
	char byte,
) *Hashset {
	charCollectionMap.Lock()
	collection := (*charCollectionMap.elementsMap)[char]
	charCollectionMap.Unlock()

	if collection == nil {
		return EmptyHashset()
	}

	items := collection.ListPtrLock()

	return NewHashsetUsingArray(items)
}

func (charCollectionMap *CharCollectionMap) HashsetByStringFirstChar(
	str string,
) *Hashset {
	char := charCollectionMap.GetChar(str)

	return charCollectionMap.HashsetByChar(char)
}

func (charCollectionMap *CharCollectionMap) HashsetByStringFirstCharLock(
	str string,
) *Hashset {
	char := charCollectionMap.GetChar(str)

	return charCollectionMap.HashsetByCharLock(char)
}
