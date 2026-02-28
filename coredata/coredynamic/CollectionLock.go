package coredynamic

import (
	"fmt"
	"sync"
)

// --- Thread-Safe Accessors ---

// LengthLock returns the number of items with mutex protection.
func (it *Collection[T]) LengthLock() int {
	it.Lock()
	defer it.Unlock()

	if it == nil {
		return 0
	}
	return len(it.items)
}

// IsEmptyLock returns true if the collection has no items, with mutex protection.
func (it *Collection[T]) IsEmptyLock() bool {
	it.Lock()
	defer it.Unlock()

	return it == nil || len(it.items) == 0
}

// --- Thread-Safe Mutators ---

// AddLock appends a single item with mutex protection.
func (it *Collection[T]) AddLock(item T) *Collection[T] {
	it.Lock()
	defer it.Unlock()

	it.items = append(it.items, item)
	return it
}

// AddsLock appends multiple items with mutex protection.
func (it *Collection[T]) AddsLock(items ...T) *Collection[T] {
	it.Lock()
	defer it.Unlock()

	it.items = append(it.items, items...)
	return it
}

// AddManyLock is an alias for AddsLock.
func (it *Collection[T]) AddManyLock(items ...T) *Collection[T] {
	return it.AddsLock(items...)
}

// AddCollectionLock appends all items from another collection with mutex protection.
func (it *Collection[T]) AddCollectionLock(other *Collection[T]) *Collection[T] {
	if other == nil || other.IsEmpty() {
		return it
	}

	it.Lock()
	defer it.Unlock()

	it.items = append(it.items, other.items...)
	return it
}

// AddCollectionsLock appends items from multiple collections with mutex protection.
func (it *Collection[T]) AddCollectionsLock(others ...*Collection[T]) *Collection[T] {
	it.Lock()
	defer it.Unlock()

	for _, other := range others {
		if other == nil || other.IsEmpty() {
			continue
		}
		it.items = append(it.items, other.items...)
	}
	return it
}

// AddIfLock appends an item only when isAdd is true, with mutex protection.
func (it *Collection[T]) AddIfLock(isAdd bool, item T) *Collection[T] {
	if !isAdd {
		return it
	}

	it.Lock()
	defer it.Unlock()

	it.items = append(it.items, item)
	return it
}

// RemoveAtLock removes the item at the given index with mutex protection.
func (it *Collection[T]) RemoveAtLock(index int) bool {
	it.Lock()
	defer it.Unlock()

	if !it.HasIndex(index) {
		return false
	}
	it.items = append(it.items[:index], it.items[index+1:]...)
	return true
}

// ClearLock removes all items with mutex protection, keeping allocated capacity.
func (it *Collection[T]) ClearLock() {
	it.Lock()
	defer it.Unlock()

	it.items = it.items[:0]
}

// --- Thread-Safe Readers ---

// ItemsLock returns a copy of the items slice with mutex protection.
func (it *Collection[T]) ItemsLock() []T {
	it.Lock()
	defer it.Unlock()

	if it.items == nil {
		return []T{}
	}

	copied := make([]T, len(it.items))
	copy(copied, it.items)
	return copied
}

// FirstLock returns the first item with mutex protection. Panics if empty.
func (it *Collection[T]) FirstLock() T {
	it.Lock()
	defer it.Unlock()

	return it.items[0]
}

// LastLock returns the last item with mutex protection. Panics if empty.
func (it *Collection[T]) LastLock() T {
	it.Lock()
	defer it.Unlock()

	return it.items[len(it.items)-1]
}

// --- Thread-Safe Async ---

// AddWithWgLock appends an item and calls wg.Done(), with mutex protection.
func (it *Collection[T]) AddWithWgLock(
	wg *sync.WaitGroup,
	item T,
) *Collection[T] {
	it.Lock()
	defer it.Unlock()

	it.items = append(it.items, item)
	wg.Done()

	return it
}

// LoopLock iterates over a copy of items with mutex protection.
// Return true from the callback to break.
func (it *Collection[T]) LoopLock(
	fn func(index int, item T) (isBreak bool),
) {
	it.Lock()
	copied := make([]T, len(it.items))
	copy(copied, it.items)
	it.Unlock()

	for i, item := range copied {
		if fn(i, item) {
			return
		}
	}
}

// FilterLock returns a new Collection containing only items that match
// the predicate, with mutex protection during read.
func (it *Collection[T]) FilterLock(
	predicate func(T) bool,
) *Collection[T] {
	it.Lock()
	copied := make([]T, len(it.items))
	copy(copied, it.items)
	it.Unlock()

	result := NewCollection[T](len(copied) / 2)

	for _, item := range copied {
		if predicate(item) {
			result.items = append(result.items, item)
		}
	}
	return result
}

// StringsLock returns string representations of items with mutex protection.
func (it *Collection[T]) StringsLock() []string {
	it.Lock()
	defer it.Unlock()

	slice := make([]string, len(it.items))

	for i, item := range it.items {
		slice[i] = fmt.Sprintf("%v", item)
	}
	return slice
}
