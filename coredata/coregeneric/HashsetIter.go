package coregeneric

import "iter"

// All returns an iter.Seq2[T, bool] over all set members.
// The bool value is always true (matching the internal map representation).
//
//	for item, _ := range hashset.All() { ... }
func (it *Hashset[T]) All() iter.Seq2[T, bool] {
	return func(yield func(T, bool) bool) {
		for item, v := range it.items {
			if !yield(item, v) {
				return
			}
		}
	}
}

// Values returns an iter.Seq[T] over all set members.
//
//	for item := range hashset.Values() { ... }
func (it *Hashset[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for item := range it.items {
			if !yield(item) {
				return
			}
		}
	}
}
