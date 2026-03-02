package coregeneric

// Numeric is a constraint for all built-in numeric types (integers + floats).
// Used by PairDivide, TripleDivide, and other arithmetic operations.
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// IntegerNumeric is a constraint for integer-only numeric types (signed + unsigned).
type IntegerNumeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// FloatNumeric is a constraint for floating-point numeric types.
type FloatNumeric interface {
	~float32 | ~float64
}

// UnsignedNumeric is a constraint for unsigned integer types only.
// These types cannot hold negative values.
type UnsignedNumeric interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// SignedNumeric is a constraint for all signed numeric types (signed integers + floats).
// These types can represent negative values.
type SignedNumeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}

// SignedInteger is a constraint for signed integer types only.
type SignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
