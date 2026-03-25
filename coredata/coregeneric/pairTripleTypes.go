package coregeneric

// =============================================================================
// Pair[L, R] — type aliases for common combinations
// =============================================================================

type StringStringPair = Pair[string, string]
type StringIntPair = Pair[string, int]
type StringInt64Pair = Pair[string, int64]
type StringFloat64Pair = Pair[string, float64]
type StringBoolPair = Pair[string, bool]
type StringAnyPair = Pair[string, any]
type IntIntPair = Pair[int, int]
type IntStringPair = Pair[int, string]
type IntAnyPair = Pair[int, any]
type AnyAnyPair = Pair[any, any]

// =============================================================================
// Triple[A, B, C] — type aliases for common combinations
// =============================================================================

type StringStringStringTriple = Triple[string, string, string]
type StringIntStringTriple = Triple[string, int, string]
type StringAnyAnyTriple = Triple[string, any, any]
type AnyAnyAnyTriple = Triple[any, any, any]
