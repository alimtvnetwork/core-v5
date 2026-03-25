package corecmp

import "github.com/alimtvnetwork/core/corecomparator"

// AnyItem compares two any values for equality using Go's built-in == operator.
//
// Comparison logic (evaluated in order):
//  1. Both nil           → Equal
//  2. One nil, other not → NotEqual
//  3. left == right      → Equal (works for comparable built-in types)
//  4. Otherwise          → Inconclusive (types may not be comparable via ==)
//
// Returns:
//   - corecomparator.Equal        — both values are nil or structurally equal
//   - corecomparator.NotEqual     — exactly one value is nil
//   - corecomparator.Inconclusive — values differ or are not comparable via ==
//
// Note: This function does NOT perform deep comparison. Slices, maps, and
// structs without a comparable implementation will always return Inconclusive
// when they differ. Use reflect.DeepEqual or typed comparators for those cases.
func AnyItem(left, right any) corecomparator.Compare {
	if left == nil && right == nil {
		return corecomparator.Equal
	}

	if left == nil || right == nil {
		return corecomparator.NotEqual
	}

	if left == right {
		return corecomparator.Equal
	}

	return corecomparator.Inconclusive
}
