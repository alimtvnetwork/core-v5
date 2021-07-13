package corestr

import (
	"fmt"
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

type SimpleSlice struct {
	Items []string `json:"Items,omitempty"`
}

func NewSimpleSlice(capacity int) *SimpleSlice {
	slice := make([]string, 0, capacity)

	return &SimpleSlice{
		slice,
	}
}

func EmptySimpleSlice() *SimpleSlice {
	return NewSimpleSlice(0)
}

func (it *SimpleSlice) Add(item string) *SimpleSlice {
	it.Items = append(it.Items, item)

	return it
}

func (it *SimpleSlice) AddIf(isAdd bool, item string) *SimpleSlice {
	if !isAdd {
		return it
	}

	it.Items = append(it.Items, item)

	return it
}

func (it *SimpleSlice) Adds(items ...string) *SimpleSlice {
	for _, item := range items {
		it.Items = append(
			it.Items,
			item)
	}

	return it
}

func (it *SimpleSlice) AddsIf(
	isAdd bool,
	items ...string,
) *SimpleSlice {
	if !isAdd {
		return it
	}

	return it.Adds(items...)
}

func (it *SimpleSlice) AddErr(err error) *SimpleSlice {
	if err != nil {
		return it.Add(err.Error())
	}

	return it
}

func (it *SimpleSlice) FirstDynamic() interface{} {
	return it.Items[0]
}

func (it *SimpleSlice) First() string {
	return it.Items[0]
}

func (it *SimpleSlice) LastDynamic() interface{} {
	return it.Items[it.LastIndex()]
}

func (it *SimpleSlice) Last() string {
	return it.Items[it.LastIndex()]
}

func (it *SimpleSlice) FirstOrDefaultDynamic() interface{} {
	return it.FirstOrDefault()
}

func (it *SimpleSlice) FirstOrDefault() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return it.First()
}

func (it *SimpleSlice) LastOrDefaultDynamic() interface{} {
	return it.LastOrDefault()
}

func (it *SimpleSlice) LastOrDefault() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return it.Last()
}

func (it *SimpleSlice) SkipDynamic(skippingItemsCount int) interface{} {
	return it.Items[skippingItemsCount:]
}

func (it *SimpleSlice) Skip(skippingItemsCount int) []string {
	return it.Items[skippingItemsCount:]
}

func (it *SimpleSlice) TakeDynamic(takeDynamicItems int) interface{} {
	return it.Items[:takeDynamicItems]
}

func (it *SimpleSlice) Take(takeDynamicItems int) []string {
	return it.Items[:takeDynamicItems]
}

func (it *SimpleSlice) LimitDynamic(limit int) interface{} {
	return it.Take(limit)
}

func (it *SimpleSlice) Limit(limit int) []string {
	return it.Take(limit)
}

func (it *SimpleSlice) Length() int {
	if it == nil {
		return 0
	}

	return len(it.Items)
}

func (it *SimpleSlice) Count() int {
	return it.Length()
}

func (it *SimpleSlice) IsEmpty() bool {
	return it == nil || it.Length() == 0
}

func (it *SimpleSlice) HasAnyItem() bool {
	return !it.IsEmpty()
}

func (it *SimpleSlice) LastIndex() int {
	return it.Length() - 1
}

func (it *SimpleSlice) HasIndex(index int) bool {
	return it.LastIndex() >= index
}

func (it *SimpleSlice) Strings() []string {
	return it.Items
}

func (it *SimpleSlice) String() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return strings.Join(
		it.Items,
		constants.NewLineUnix)
}

func (it *SimpleSlice) ConcatNewStrings(items ...string) []string {
	if it == nil {
		return CloneSlice(items)
	}

	slice := make(
		[]string,
		0,
		it.Length()+len(items))

	for _, item := range it.Items {
		slice = append(slice, item)
	}

	for _, item := range items {
		slice = append(slice, item)
	}

	return slice
}

func (it *SimpleSlice) ConcatNew(items ...string) *SimpleSlice {
	concatNew := it.ConcatNewStrings(items...)

	return &SimpleSlice{
		concatNew,
	}
}

func (it *SimpleSlice) ToCollection(isClone bool) *Collection {
	return NewCollectionUsingStrings(&it.Items, isClone)
}

func (it *SimpleSlice) CsvStrings() []string {
	if it.IsEmpty() {
		return []string{}
	}

	newSlice := make([]string, it.Length())

	for i, item := range it.Items {
		newSlice[i] = fmt.Sprintf(
			constants.SprintSingleQuoteFormat,
			item)
	}

	return newSlice
}
