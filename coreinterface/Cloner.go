package coreinterface

// Cloner defines a generic interface for types that can produce
// an independent copy of themselves.
//
// Any type implementing Cloner[T] guarantees that mutations to the
// clone will not affect the original, and vice versa.
//
// Usage:
//
//	type RequestAttribute struct { ... }
//
//	func (it *RequestAttribute) Clone() *RequestAttribute {
//	    return &RequestAttribute{ ... }
//	}
//
//	// RequestAttribute now satisfies Cloner[*RequestAttribute]
type Cloner[T any] interface {
	Clone() T
}
