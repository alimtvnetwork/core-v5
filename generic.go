package core

// EmptySlicePtr returns a pointer to an empty slice of type T.
// It replaces EmptyAnysPtr, EmptyFloat32Ptr, EmptyFloat64Ptr,
// EmptyBoolsPtr, EmptyIntsPtr, EmptyBytePtr, EmptyStringsPtr,
// EmptyPointerStringsPtr and similar per-type functions.
//
// Usage:
//
//	ints := core.EmptySlicePtr[int]()       // returns *[]int
//	strs := core.EmptySlicePtr[string]()    // returns *[]string
func EmptySlicePtr[T any]() *[]T {
	s := make([]T, 0)
	return &s
}

// SlicePtrByLength returns a pointer to a zero-valued slice of type T with the given length.
// It replaces StringsPtrByLength and similar per-type functions.
func SlicePtrByLength[T any](length int) *[]T {
	s := make([]T, length)
	return &s
}

// SlicePtrByCapacity returns a pointer to a slice of type T with the given length and capacity.
// It replaces StringsPtrByCapacity, PointerStringsPtrByCapacity and similar per-type functions.
func SlicePtrByCapacity[T any](length, cap int) *[]T {
	s := make([]T, length, cap)
	return &s
}

// EmptyMapPtr returns a pointer to an empty map of type map[K]V.
// It replaces EmptyStringsMapPtr, EmptyStringToIntMapPtr and similar per-type functions.
func EmptyMapPtr[K comparable, V any]() *map[K]V {
	m := make(map[K]V)
	return &m
}
