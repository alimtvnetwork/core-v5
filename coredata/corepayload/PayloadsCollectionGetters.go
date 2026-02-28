package corepayload

import "gitlab.com/auk-go/core/defaultcapacity"

// PayloadsCollectionGetters.go — Read-only accessor and query methods extracted from PayloadsCollection.go

// =============================================================================
// Length and state
// =============================================================================

func (it *PayloadsCollection) Length() int {
	if it == nil {
		return 0
	}

	return len(it.Items)
}

func (it *PayloadsCollection) Count() int {
	return it.Length()
}

func (it *PayloadsCollection) IsEmpty() bool {
	return it.Length() == 0
}

func (it *PayloadsCollection) HasAnyItem() bool {
	return !it.IsEmpty()
}

func (it *PayloadsCollection) LastIndex() int {
	return it.Length() - 1
}

func (it *PayloadsCollection) HasIndex(index int) bool {
	return it.LastIndex() >= index
}

// =============================================================================
// Element access
// =============================================================================

func (it *PayloadsCollection) FirstDynamic() any {
	return it.Items[0]
}

func (it *PayloadsCollection) First() *PayloadWrapper {
	return it.Items[0]
}

func (it *PayloadsCollection) LastDynamic() any {
	return it.Items[it.LastIndex()]
}

func (it *PayloadsCollection) Last() *PayloadWrapper {
	return it.Items[it.LastIndex()]
}

func (it *PayloadsCollection) FirstOrDefaultDynamic() any {
	return it.FirstOrDefault()
}

func (it *PayloadsCollection) FirstOrDefault() *PayloadWrapper {
	if it.IsEmpty() {
		return nil
	}

	return it.First()
}

func (it *PayloadsCollection) LastOrDefaultDynamic() any {
	return it.LastOrDefault()
}

func (it *PayloadsCollection) LastOrDefault() *PayloadWrapper {
	if it.IsEmpty() {
		return nil
	}

	return it.Last()
}

// =============================================================================
// Slice operations
// =============================================================================

func (it *PayloadsCollection) SkipDynamic(skippingItemsCount int) any {
	return it.Items[skippingItemsCount:]
}

func (it *PayloadsCollection) Skip(skippingItemsCount int) []*PayloadWrapper {
	return it.Items[skippingItemsCount:]
}

func (it *PayloadsCollection) SkipCollection(skippingItemsCount int) *PayloadsCollection {
	return &PayloadsCollection{
		Items: it.Items[skippingItemsCount:],
	}
}

func (it *PayloadsCollection) TakeDynamic(takeDynamicItems int) any {
	return it.Items[:takeDynamicItems]
}

func (it *PayloadsCollection) Take(takeDynamicItems int) []*PayloadWrapper {
	return it.Items[:takeDynamicItems]
}

func (it *PayloadsCollection) TakeCollection(takeDynamicItems int) *PayloadsCollection {
	return &PayloadsCollection{
		Items: it.Items[:takeDynamicItems],
	}
}

func (it *PayloadsCollection) LimitCollection(limit int) *PayloadsCollection {
	return &PayloadsCollection{
		Items: it.Items[:limit],
	}
}

func (it *PayloadsCollection) SafeLimitCollection(limit int) *PayloadsCollection {
	limit = defaultcapacity.
		MaxLimit(it.Length(), limit)

	return &PayloadsCollection{
		Items: it.Items[:limit],
	}
}

func (it *PayloadsCollection) LimitDynamic(limit int) any {
	return it.Take(limit)
}

func (it *PayloadsCollection) Limit(limit int) []*PayloadWrapper {
	return it.Take(limit)
}

// =============================================================================
// Strings
// =============================================================================

func (it *PayloadsCollection) Strings() []string {
	list := make([]string, it.Length())

	for i, item := range it.Items {
		list[i] = item.String()
	}

	return list
}

// =============================================================================
// Equality
// =============================================================================

func (it *PayloadsCollection) IsEqual(another *PayloadsCollection) bool {
	if it == nil && another == nil {
		return true
	}

	if it == nil || another == nil {
		return false
	}

	if it.Length() != another.Length() {
		return false
	}

	return it.IsEqualItems(another.Items...)
}

func (it *PayloadsCollection) IsEqualItems(lines ...*PayloadWrapper) bool {
	if it == nil && lines == nil {
		return true
	}

	if it == nil || lines == nil {
		return false
	}

	if it.Length() != len(lines) {
		return false
	}

	for i, item := range it.Items {
		anotherItem := lines[i]

		if !item.IsEqual(anotherItem) {
			return false
		}
	}

	return true
}
