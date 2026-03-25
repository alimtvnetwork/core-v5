package coregeneric

// newTripleCreator provides factory methods for Triple[A, B, C] types.
//
// Usage:
//
//	coregeneric.New.Triple.StringStringString("left", "mid", "right")
//	coregeneric.New.Triple.Any("a", 42, true)
//	coregeneric.New.Triple.Split("a.b.c", ".")
//	coregeneric.New.Triple.DivideInt(100)
type newTripleCreator struct{}

// --- Type-specific constructors ---

// StringStringString creates a valid Triple[string, string, string].
func (it newTripleCreator) StringStringString(left, middle, right string) *Triple[string, string, string] {
	return NewTriple(left, middle, right)
}

// StringIntString creates a valid Triple[string, int, string].
func (it newTripleCreator) StringIntString(left string, middle int, right string) *Triple[string, int, string] {
	return NewTriple(left, middle, right)
}

// StringAnyAny creates a valid Triple[string, any, any].
func (it newTripleCreator) StringAnyAny(left string, middle, right any) *Triple[string, any, any] {
	return NewTriple(left, middle, right)
}

// Any creates a valid Triple[any, any, any].
func (it newTripleCreator) Any(left, middle, right any) *Triple[any, any, any] {
	return NewTriple(left, middle, right)
}

// --- Invalid constructors ---

// InvalidStringStringString creates an invalid Triple[string, string, string] with a message.
func (it newTripleCreator) InvalidStringStringString(message string) *Triple[string, string, string] {
	return InvalidTriple[string, string, string](message)
}

// InvalidAny creates an invalid Triple[any, any, any] with a message.
func (it newTripleCreator) InvalidAny(message string) *Triple[any, any, any] {
	return InvalidTriple[any, any, any](message)
}

// --- String split constructors ---

// Split splits a string by separator into a Triple[string, string, string].
//
//	triple := coregeneric.New.Triple.Split("a.b.c", ".")
//	triple.Left   // "a"
//	triple.Middle // "b"
//	triple.Right  // "c"
func (it newTripleCreator) Split(input, sep string) *Triple[string, string, string] {
	return TripleFromSplit(input, sep)
}

// SplitTrimmed splits and trims whitespace from all parts.
func (it newTripleCreator) SplitTrimmed(input, sep string) *Triple[string, string, string] {
	return TripleFromSplitTrimmed(input, sep)
}

// SplitN splits into at most 3 parts. Third part contains the remainder.
//
//	triple := coregeneric.New.Triple.SplitN("a:b:c:d:e", ":")
//	triple.Left   // "a"
//	triple.Middle // "b"
//	triple.Right  // "c:d:e"
func (it newTripleCreator) SplitN(input, sep string) *Triple[string, string, string] {
	return TripleFromSplitN(input, sep)
}

// SplitNTrimmed splits into at most 3 parts with trimming.
func (it newTripleCreator) SplitNTrimmed(input, sep string) *Triple[string, string, string] {
	return TripleFromSplitNTrimmed(input, sep)
}

// FromSlice creates a Triple from a string slice.
func (it newTripleCreator) FromSlice(parts []string) *Triple[string, string, string] {
	return TripleFromSlice(parts)
}

// --- Number division constructors ---

// DivideInt splits an int into three parts.
//
//	triple := coregeneric.New.Triple.DivideInt(10) // {3, 3, 4}
func (it newTripleCreator) DivideInt(value int) *Triple[int, int, int] {
	return TripleDivide(value)
}

// DivideInt64 splits an int64 into three parts.
func (it newTripleCreator) DivideInt64(value int64) *Triple[int64, int64, int64] {
	return TripleDivide(value)
}

// DivideFloat64 splits a float64 into three equal parts.
func (it newTripleCreator) DivideFloat64(value float64) *Triple[float64, float64, float64] {
	return TripleDivide(value)
}

// DivideIntWeighted splits an int by two ratios. Remainder goes to Right.
//
//	triple := coregeneric.New.Triple.DivideIntWeighted(100, 0.2, 0.3) // {20, 30, 50}
func (it newTripleCreator) DivideIntWeighted(
	value int,
	leftRatio, middleRatio float64,
) *Triple[int, int, int] {
	return TripleDivideWeighted(value, leftRatio, middleRatio)
}

// DivideFloat64Weighted splits a float64 by two ratios.
func (it newTripleCreator) DivideFloat64Weighted(
	value float64,
	leftRatio, middleRatio float64,
) *Triple[float64, float64, float64] {
	return TripleDivideWeighted(value, leftRatio, middleRatio)
}
