package coregeneric

import "strings"

// =============================================================================
// Generic same-type constructors
// =============================================================================

// NewTripleOf creates a valid Triple[T, T, T] where all three share the same type.
//
// Usage:
//
//	triple := coregeneric.NewTripleOf(1, 2, 3)       // Triple[int, int, int]
//	triple := coregeneric.NewTripleOf("a", "b", "c") // Triple[string, string, string]
func NewTripleOf[T any](left, middle, right T) *Triple[T, T, T] {
	return NewTriple(left, middle, right)
}

// InvalidTripleOf creates an invalid Triple[T, T, T] with a message.
func InvalidTripleOf[T any](message string) *Triple[T, T, T] {
	return InvalidTriple[T, T, T](message)
}

// =============================================================================
// String split → Triple[string, string, string]
// =============================================================================

// TripleFromSplit splits a string by separator and creates a Triple.
//
// If the split produces:
//   - 0 parts → invalid triple
//   - 1 part  → Left only (invalid)
//   - 2 parts → Left + Right, Middle = "" (invalid)
//   - 3 parts → Left, Middle, Right (valid)
//   - 4+ parts → Left = first, Middle = second, Right = last (valid)
//
// Usage:
//
//	triple := coregeneric.TripleFromSplit("a.b.c", ".")
//	triple.Left   // "a"
//	triple.Middle // "b"
//	triple.Right  // "c"
func TripleFromSplit(input, sep string) *Triple[string, string, string] {
	parts := strings.Split(input, sep)

	return tripleFromParts(parts)
}

// TripleFromSplitTrimmed splits a string and trims whitespace from all parts.
func TripleFromSplitTrimmed(input, sep string) *Triple[string, string, string] {
	parts := strings.Split(input, sep)

	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	return tripleFromParts(parts)
}

// TripleFromSplitN splits into at most 3 parts.
// If there are more separators, the third part contains the rest.
//
// Usage:
//
//	triple := coregeneric.TripleFromSplitN("a:b:c:d:e", ":")
//	triple.Left   // "a"
//	triple.Middle // "b"
//	triple.Right  // "c:d:e"
func TripleFromSplitN(input, sep string) *Triple[string, string, string] {
	parts := strings.SplitN(input, sep, 3)

	return tripleFromParts(parts)
}

// TripleFromSplitNTrimmed splits into at most 3 parts with trimming.
func TripleFromSplitNTrimmed(input, sep string) *Triple[string, string, string] {
	parts := strings.SplitN(input, sep, 3)

	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	return tripleFromParts(parts)
}

// TripleFromSlice creates a Triple from a string slice.
func TripleFromSlice(parts []string) *Triple[string, string, string] {
	return tripleFromParts(parts)
}

func tripleFromParts(parts []string) *Triple[string, string, string] {
	length := len(parts)

	switch {
	case length == 0:
		return InvalidTriple[string, string, string]("empty input")
	case length == 1:
		return NewTripleWithMessage(parts[0], "", "", false, "only one part found")
	case length == 2:
		return NewTripleWithMessage(parts[0], "", parts[1], false, "only two parts found")
	default:
		return NewTriple(parts[0], parts[1], parts[length-1])
	}
}

// =============================================================================
// Number division → Triple[T, T, T]
// =============================================================================

// TripleDivide splits a numeric value into three parts.
//
// For integers: Left = value/3, Middle = value/3, Right = value - 2*(value/3).
// Handles remainders by putting extra in the Right.
//
// Usage:
//
//	triple := coregeneric.TripleDivide(10)  // {3, 3, 4}
//	triple := coregeneric.TripleDivide(9)   // {3, 3, 3}
func TripleDivide[T Numeric](value T) *Triple[T, T, T] {
	third := value / 3
	remainder := value - third*2

	return NewTriple(third, third, remainder)
}

// TripleDivideWeighted splits a numeric value by two ratios.
// leftRatio + middleRatio should be <= 1.0. Remainder goes to Right.
//
// Usage:
//
//	triple := coregeneric.TripleDivideWeighted(100, 0.2, 0.3)  // {20, 30, 50}
func TripleDivideWeighted[T Numeric](value T, leftRatio, middleRatio float64) *Triple[T, T, T] {
	left := T(float64(value) * leftRatio)
	middle := T(float64(value) * middleRatio)
	right := value - left - middle

	return NewTriple(left, middle, right)
}
