package coregeneric

import "fmt"

// Pair is a generic two-value container.
// It generalizes corestr.LeftRight from string-only to any types L, R.
//
// Usage:
//
//	pair := coregeneric.NewPair("key", 42)
//	pair.Left  // "key"
//	pair.Right // 42
type Pair[L any, R any] struct {
	Left    L      `json:"Left,omitempty"`
	Right   R      `json:"Right,omitempty"`
	IsValid bool   `json:"IsValid,omitempty"`
	Message string `json:"Message,omitempty"`
}

// NewPair creates a valid Pair[L, R].
func NewPair[L any, R any](left L, right R) *Pair[L, R] {
	return &Pair[L, R]{
		Left:    left,
		Right:   right,
		IsValid: true,
	}
}

// NewPairWithMessage creates a Pair[L, R] with validity and message.
func NewPairWithMessage[L any, R any](
	left L,
	right R,
	isValid bool,
	message string,
) *Pair[L, R] {
	return &Pair[L, R]{
		Left:    left,
		Right:   right,
		IsValid: isValid,
		Message: message,
	}
}

// InvalidPair creates an invalid Pair with zero values and a message.
func InvalidPair[L any, R any](message string) *Pair[L, R] {
	return &Pair[L, R]{
		IsValid: false,
		Message: message,
	}
}

// InvalidPairNoMessage creates an invalid Pair with zero values and no message.
func InvalidPairNoMessage[L any, R any]() *Pair[L, R] {
	return &Pair[L, R]{
		IsValid: false,
	}
}

// HasMessage returns true if Message is non-empty.
func (it *Pair[L, R]) HasMessage() bool {
	return it != nil && it.Message != ""
}

// IsInvalid returns true if the pair is not valid.
func (it *Pair[L, R]) IsInvalid() bool {
	return it == nil || !it.IsValid
}

// Values returns both Left and Right.
func (it *Pair[L, R]) Values() (L, R) {
	return it.Left, it.Right
}

// Clone creates a shallow copy of the Pair.
func (it *Pair[L, R]) Clone() *Pair[L, R] {
	if it == nil {
		return nil
	}

	return &Pair[L, R]{
		Left:    it.Left,
		Right:   it.Right,
		IsValid: it.IsValid,
		Message: it.Message,
	}
}

// IsEqual compares two Pairs using fmt.Sprintf for value comparison.
// For comparable types, prefer direct field comparison.
func (it *Pair[L, R]) IsEqual(another *Pair[L, R]) bool {
	if it == nil && another == nil {
		return true
	}

	if it == nil || another == nil {
		return false
	}

	return it.IsValid == another.IsValid &&
		fmt.Sprintf("%v", it.Left) == fmt.Sprintf("%v", another.Left) &&
		fmt.Sprintf("%v", it.Right) == fmt.Sprintf("%v", another.Right)
}

// String returns a formatted string representation.
func (it *Pair[L, R]) String() string {
	if it == nil {
		return ""
	}

	return fmt.Sprintf("{Left: %v, Right: %v, IsValid: %v}", it.Left, it.Right, it.IsValid)
}

// Clear resets the Pair to zero values.
func (it *Pair[L, R]) Clear() {
	if it == nil {
		return
	}

	var zeroL L
	var zeroR R

	it.Left = zeroL
	it.Right = zeroR
	it.IsValid = false
	it.Message = ""
}

// Dispose is an alias for Clear.
func (it *Pair[L, R]) Dispose() {
	it.Clear()
}
