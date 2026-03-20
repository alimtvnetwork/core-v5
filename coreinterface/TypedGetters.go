package coreinterface

// TypedValueGetter is a generic interface for types that expose a typed Value().
//
// This replaces the need for type-specific getter interfaces like
// ValueIntegerGetter, ValueStringGetter, ValueFloat64Getter, etc.
//
// Existing type-specific getters are retained for backward compatibility
// but new code should prefer TypedValueGetter[T].
//
// Usage:
//
//	type MyConfig struct { val string }
//	func (c MyConfig) Value() string { return c.val }
//	// MyConfig satisfies TypedValueGetter[string]
type TypedValueGetter[T any] interface {
	Value() T
}

// TypedValuesGetter is a generic interface for types that expose
// a typed slice via Values().
//
// Usage:
//
//	type StringList struct { items []string }
//	func (s StringList) Values() []string { return s.items }
//	// StringList satisfies TypedValuesGetter[string]
type TypedValuesGetter[T any] interface {
	Values() []T
}

// TypedKeyValueGetter is a generic interface for key-value accessors.
//
// Usage:
//
//	type Entry struct { k string; v int }
//	func (e Entry) Key() string { return e.k }
//	func (e Entry) Value() int  { return e.v }
//	// Entry satisfies TypedKeyValueGetter[string, int]
type TypedKeyValueGetter[K comparable, V any] interface {
	Key() K
	TypedValueGetter[V]
}
