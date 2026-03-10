package coregeneric

import "github.com/alimtvnetwork/core/corecomparator"

// =============================================================================
// Comparison — returns corecomparator.Compare enum
// =============================================================================

// CompareNumeric compares two numeric values and returns a corecomparator.Compare result.
//
//	Equal       if left == right
//	LeftGreater if left > right
//	LeftLess    if left < right
func CompareNumeric[T Numeric](left, right T) corecomparator.Compare {
	if left == right {
		return corecomparator.Equal
	}

	if left > right {
		return corecomparator.LeftGreater
	}

	return corecomparator.LeftLess
}

// =============================================================================
// Relational predicates
// =============================================================================

// IsLess returns true if left < right.
func IsLess[T Numeric](left, right T) bool {
	return left < right
}

// IsLessOrEqual returns true if left <= right.
func IsLessOrEqual[T Numeric](left, right T) bool {
	return left <= right
}

// IsGreater returns true if left > right.
func IsGreater[T Numeric](left, right T) bool {
	return left > right
}

// IsGreaterOrEqual returns true if left >= right.
func IsGreaterOrEqual[T Numeric](left, right T) bool {
	return left >= right
}

// IsNumericEqual returns true if left == right.
func IsNumericEqual[T Numeric](left, right T) bool {
	return left == right
}

// IsNotEqual returns true if left != right.
func IsNotEqual[T Numeric](left, right T) bool {
	return left != right
}

// =============================================================================
// Clamping and bounding
// =============================================================================

// Clamp constrains value to the inclusive range [minVal, maxVal].
func Clamp[T Numeric](value, minVal, maxVal T) T {
	if value < minVal {
		return minVal
	}

	if value > maxVal {
		return maxVal
	}

	return value
}

// ClampMin returns the greater of value and minVal.
func ClampMin[T Numeric](value, minVal T) T {
	if value < minVal {
		return minVal
	}

	return value
}

// ClampMax returns the lesser of value and maxVal.
func ClampMax[T Numeric](value, maxVal T) T {
	if value > maxVal {
		return maxVal
	}

	return value
}

// InRange returns true if value is within [minVal, maxVal] inclusive.
func InRange[T Numeric](value, minVal, maxVal T) bool {
	return value >= minVal && value <= maxVal
}

// InRangeExclusive returns true if value is within (minVal, maxVal) exclusive.
func InRangeExclusive[T Numeric](value, minVal, maxVal T) bool {
	return value > minVal && value < maxVal
}

// =============================================================================
// Arithmetic utilities
// =============================================================================

// Abs returns the absolute value. Only meaningful for signed types.
func Abs[T SignedNumeric](value T) T {
	if value < 0 {
		return -value
	}

	return value
}

// AbsDiff returns the absolute difference between two numeric values.
func AbsDiff[T SignedNumeric](left, right T) T {
	diff := left - right

	if diff < 0 {
		return -diff
	}

	return diff
}

// Sum returns the sum of all provided values.
func Sum[T Numeric](values ...T) T {
	var total T

	for _, v := range values {
		total += v
	}

	return total
}

// MinOf returns the minimum of two values.
func MinOf[T Numeric](a, b T) T {
	if a < b {
		return a
	}

	return b
}

// MaxOf returns the maximum of two values.
func MaxOf[T Numeric](a, b T) T {
	if a > b {
		return a
	}

	return b
}

// MinOfSlice returns the minimum value in a slice.
// Panics on empty slice.
func MinOfSlice[T Numeric](values []T) T {
	result := values[0]

	for _, v := range values[1:] {
		if v < result {
			result = v
		}
	}

	return result
}

// MaxOfSlice returns the maximum value in a slice.
// Panics on empty slice.
func MaxOfSlice[T Numeric](values []T) T {
	result := values[0]

	for _, v := range values[1:] {
		if v > result {
			result = v
		}
	}

	return result
}

// =============================================================================
// Sign and zero checks
// =============================================================================

// IsZero returns true if value is the zero value for its type.
func IsZero[T Numeric](value T) bool {
	return value == 0
}

// IsPositive returns true if value > 0.
func IsPositive[T Numeric](value T) bool {
	return value > 0
}

// IsNegative returns true if value < 0.
// For unsigned types this always returns false.
func IsNegative[T SignedNumeric](value T) bool {
	return value < 0
}

// IsNonNegative returns true if value >= 0.
// For unsigned types this always returns true.
func IsNonNegative[T Numeric](value T) bool {
	return value >= 0
}

// Sign returns -1, 0, or 1 indicating the sign of the value.
func Sign[T SignedNumeric](value T) int {
	if value < 0 {
		return -1
	}

	if value > 0 {
		return 1
	}

	return 0
}

// =============================================================================
// Safe division
// =============================================================================

// SafeDiv divides numerator by denominator, returning zero on division by zero.
func SafeDiv[T Numeric](numerator, denominator T) T {
	if denominator == 0 {
		return 0
	}

	return numerator / denominator
}

// SafeDivOrDefault divides numerator by denominator,
// returning defaultVal on division by zero.
func SafeDivOrDefault[T Numeric](numerator, denominator, defaultVal T) T {
	if denominator == 0 {
		return defaultVal
	}

	return numerator / denominator
}
