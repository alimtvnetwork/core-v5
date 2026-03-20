package coregeneric

import "iter"

// All returns an iter.Seq2[K, V] over all key-value pairs.
//
//	for key, value := range hashmap.All() { ... }
func (it *Hashmap[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range it.items {
			if !yield(k, v) {
				return
			}
		}
	}
}

// Keys returns an iter.Seq[K] over all keys.
//
//	for key := range hashmap.Keys() { ... }
func (it *Hashmap[K, V]) Keys() iter.Seq[K] {
	return func(yield func(K) bool) {
		for k := range it.items {
			if !yield(k) {
				return
			}
		}
	}
}

// Values returns an iter.Seq[V] over all values.
//
//	for value := range hashmap.Values() { ... }
func (it *Hashmap[K, V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range it.items {
			if !yield(v) {
				return
			}
		}
	}
}
