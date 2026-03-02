package coregeneric

// Numeric is a constraint for all built-in numeric types.
// Used by PairDivide, TripleDivide, and other arithmetic operations.
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// IntegerNumeric is a constraint for integer-only numeric types.
type IntegerNumeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// FloatNumeric is a constraint for floating-point numeric types.
type FloatNumeric interface {
	~float32 | ~float64
}
