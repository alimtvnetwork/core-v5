package coregeneric

import "iter"

// All returns an iter.Seq2[int, T] that yields each item with its index.
// Supports Go 1.23+ range-over-func:
//
//	for i, item := range collection.All() { ... }
func (it *Collection[T]) All() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, item := range it.items {
			if !yield(i, item) {
				return
			}
		}
	}
}

// Values returns an iter.Seq[T] that yields each item without index.
//
//	for item := range collection.Values() { ... }
func (it *Collection[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, item := range it.items {
			if !yield(item) {
				return
			}
		}
	}
}

// Backward returns an iter.Seq2[int, T] that yields items in reverse order.
//
//	for i, item := range collection.Backward() { ... }
func (it *Collection[T]) Backward() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i := len(it.items) - 1; i >= 0; i-- {
			if !yield(i, it.items[i]) {
				return
			}
		}
	}
}
