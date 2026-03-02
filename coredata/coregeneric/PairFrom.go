package coregeneric

import "strings"

// =============================================================================
// Generic same-type constructors
// =============================================================================

// NewPairOf creates a valid Pair[T, T] where both Left and Right share the same type.
//
// Usage:
//
//	pair := coregeneric.NewPairOf(10, 20)       // Pair[int, int]
//	pair := coregeneric.NewPairOf("a", "b")     // Pair[string, string]
func NewPairOf[T any](left, right T) *Pair[T, T] {
	return NewPair(left, right)
}

// InvalidPairOf creates an invalid Pair[T, T] with a message.
func InvalidPairOf[T any](message string) *Pair[T, T] {
	return InvalidPair[T, T](message)
}

// =============================================================================
// String split → Pair[string, string]
// =============================================================================

// PairFromSplit splits a string by separator and creates a Pair[string, string].
//
// If the split produces:
//   - 0 parts → invalid pair
//   - 1 part  → Left = parts[0], Right = "" (marked invalid)
//   - 2 parts → Left = parts[0], Right = parts[1] (valid)
//   - 3+ parts → Left = parts[0], Right = last part (valid, extras dropped)
//
// Usage:
//
//	pair := coregeneric.PairFromSplit("key=value", "=")
//	pair.Left  // "key"
//	pair.Right // "value"
func PairFromSplit(input, sep string) *Pair[string, string] {
	parts := strings.SplitN(input, sep, 2)

	return pairFromParts(parts)
}

// PairFromSplitTrimmed splits a string and trims whitespace from both parts.
func PairFromSplitTrimmed(input, sep string) *Pair[string, string] {
	parts := strings.SplitN(input, sep, 2)

	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	return pairFromParts(parts)
}

// PairFromSplitFull splits without limiting parts count.
// Left = first part, Right = everything after first separator (re-joined).
//
// Usage:
//
//	pair := coregeneric.PairFromSplitFull("a:b:c:d", ":")
//	pair.Left  // "a"
//	pair.Right // "b:c:d"
func PairFromSplitFull(input, sep string) *Pair[string, string] {
	idx := strings.Index(input, sep)
	if idx < 0 {
		return NewPairWithMessage(input, "", false, "separator not found")
	}

	return NewPair(input[:idx], input[idx+len(sep):])
}

// PairFromSplitFullTrimmed is PairFromSplitFull with trimmed output.
func PairFromSplitFullTrimmed(input, sep string) *Pair[string, string] {
	idx := strings.Index(input, sep)
	if idx < 0 {
		return NewPairWithMessage(
			strings.TrimSpace(input), "", false, "separator not found",
		)
	}

	return NewPair(
		strings.TrimSpace(input[:idx]),
		strings.TrimSpace(input[idx+len(sep):]),
	)
}

// PairFromSlice creates a Pair from a string slice.
//
//   - 0 items → invalid
//   - 1 item  → Left only, invalid
//   - 2+ items → Left = first, Right = last, valid
func PairFromSlice(parts []string) *Pair[string, string] {
	return pairFromParts(parts)
}

func pairFromParts(parts []string) *Pair[string, string] {
	length := len(parts)

	switch {
	case length == 0:
		return InvalidPair[string, string]("empty input")
	case length == 1:
		return NewPairWithMessage(parts[0], "", false, "only one part found")
	default:
		return NewPair(parts[0], parts[length-1])
	}
}

// =============================================================================
// Number division → Pair[T, T]
// =============================================================================

// PairDivide splits a numeric value into two parts.
//
// For integers: Left = value/2, Right = value - value/2 (handles odd numbers).
// For floats: Left = value/2.0, Right = value/2.0 (exact halves).
//
// Usage:
//
//	pair := coregeneric.PairDivide(10)  // {5, 5}
//	pair := coregeneric.PairDivide(11)  // {5, 6}
func PairDivide[T Numeric](value T) *Pair[T, T] {
	half := value / 2
	remainder := value - half

	return NewPair(half, remainder)
}

// PairDivideWeighted splits a numeric value by a ratio (0.0 to 1.0).
// leftRatio determines how much goes to Left.
//
// Usage:
//
//	pair := coregeneric.PairDivideWeighted(100, 0.3)  // {30, 70}
func PairDivideWeighted[T Numeric](value T, leftRatio float64) *Pair[T, T] {
	left := T(float64(value) * leftRatio)
	right := value - left

	return NewPair(left, right)
}
