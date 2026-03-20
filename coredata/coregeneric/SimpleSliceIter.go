package coregeneric

import "iter"

// All returns an iter.Seq2[int, T] over all items with index.
//
//	for i, item := range simpleSlice.All() { ... }
func (it *SimpleSlice[T]) All() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, item := range *it {
			if !yield(i, item) {
				return
			}
		}
	}
}

// Values returns an iter.Seq[T] over all items.
//
//	for item := range simpleSlice.Values() { ... }
func (it *SimpleSlice[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, item := range *it {
			if !yield(item) {
				return
			}
		}
	}
}
