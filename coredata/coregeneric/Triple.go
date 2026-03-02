package coregeneric

import "fmt"

// Triple is a generic three-value container.
// It generalizes corestr.LeftMiddleRight from string-only to any types.
//
// Usage:
//
//	triple := coregeneric.NewTriple("left", 42, true)
//	triple.Left   // "left"
//	triple.Middle // 42
//	triple.Right  // true
type Triple[A any, B any, C any] struct {
	Left    A      `json:"Left,omitempty"`
	Middle  B      `json:"Middle,omitempty"`
	Right   C      `json:"Right,omitempty"`
	IsValid bool   `json:"IsValid,omitempty"`
	Message string `json:"Message,omitempty"`
}

// NewTriple creates a valid Triple.
func NewTriple[A any, B any, C any](left A, middle B, right C) *Triple[A, B, C] {
	return &Triple[A, B, C]{
		Left:    left,
		Middle:  middle,
		Right:   right,
		IsValid: true,
	}
}

// NewTripleWithMessage creates a Triple with validity and message.
func NewTripleWithMessage[A any, B any, C any](
	left A,
	middle B,
	right C,
	isValid bool,
	message string,
) *Triple[A, B, C] {
	return &Triple[A, B, C]{
		Left:    left,
		Middle:  middle,
		Right:   right,
		IsValid: isValid,
		Message: message,
	}
}

// InvalidTriple creates an invalid Triple with zero values and a message.
func InvalidTriple[A any, B any, C any](message string) *Triple[A, B, C] {
	return &Triple[A, B, C]{
		IsValid: false,
		Message: message,
	}
}

// InvalidTripleNoMessage creates an invalid Triple with zero values.
func InvalidTripleNoMessage[A any, B any, C any]() *Triple[A, B, C] {
	return &Triple[A, B, C]{
		IsValid: false,
	}
}

// HasMessage returns true if Message is non-empty.
func (it *Triple[A, B, C]) HasMessage() bool {
	return it != nil && it.Message != ""
}

// IsInvalid returns true if the triple is not valid.
func (it *Triple[A, B, C]) IsInvalid() bool {
	return it == nil || !it.IsValid
}

// Values returns all three values.
func (it *Triple[A, B, C]) Values() (A, B, C) {
	return it.Left, it.Middle, it.Right
}

// Clone creates a shallow copy of the Triple.
func (it *Triple[A, B, C]) Clone() *Triple[A, B, C] {
	if it == nil {
		return nil
	}

	return &Triple[A, B, C]{
		Left:    it.Left,
		Middle:  it.Middle,
		Right:   it.Right,
		IsValid: it.IsValid,
		Message: it.Message,
	}
}

// IsEqual compares two Triples using fmt.Sprintf for value comparison.
func (it *Triple[A, B, C]) IsEqual(another *Triple[A, B, C]) bool {
	if it == nil && another == nil {
		return true
	}

	if it == nil || another == nil {
		return false
	}

	return it.IsValid == another.IsValid &&
		fmt.Sprintf("%v", it.Left) == fmt.Sprintf("%v", another.Left) &&
		fmt.Sprintf("%v", it.Middle) == fmt.Sprintf("%v", another.Middle) &&
		fmt.Sprintf("%v", it.Right) == fmt.Sprintf("%v", another.Right)
}

// String returns a formatted string representation.
func (it *Triple[A, B, C]) String() string {
	if it == nil {
		return ""
	}

	return fmt.Sprintf(
		"{Left: %v, Middle: %v, Right: %v, IsValid: %v}",
		it.Left, it.Middle, it.Right, it.IsValid,
	)
}

// Clear resets the Triple to zero values.
func (it *Triple[A, B, C]) Clear() {
	if it == nil {
		return
	}

	var zeroA A
	var zeroB B
	var zeroC C

	it.Left = zeroA
	it.Middle = zeroB
	it.Right = zeroC
	it.IsValid = false
	it.Message = ""
}

// Dispose is an alias for Clear.
func (it *Triple[A, B, C]) Dispose() {
	it.Clear()
}
