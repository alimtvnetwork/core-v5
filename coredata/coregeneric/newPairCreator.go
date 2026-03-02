package coregeneric

// newPairCreator provides factory methods for Pair[L, R] types.
//
// Usage:
//
//	coregeneric.New.Pair.StringString("key", "value")
//	coregeneric.New.Pair.StringInt("name", 42)
//	coregeneric.New.Pair.Any("left", "right")
//	coregeneric.New.Pair.Split("key=value", "=")
//	coregeneric.New.Pair.DivideInt(100)
type newPairCreator struct{}

// --- Type-specific constructors ---

// StringString creates a valid Pair[string, string].
func (it newPairCreator) StringString(left, right string) *Pair[string, string] {
	return NewPair(left, right)
}

// StringInt creates a valid Pair[string, int].
func (it newPairCreator) StringInt(left string, right int) *Pair[string, int] {
	return NewPair(left, right)
}

// StringInt64 creates a valid Pair[string, int64].
func (it newPairCreator) StringInt64(left string, right int64) *Pair[string, int64] {
	return NewPair(left, right)
}

// StringFloat64 creates a valid Pair[string, float64].
func (it newPairCreator) StringFloat64(left string, right float64) *Pair[string, float64] {
	return NewPair(left, right)
}

// StringBool creates a valid Pair[string, bool].
func (it newPairCreator) StringBool(left string, right bool) *Pair[string, bool] {
	return NewPair(left, right)
}

// StringAny creates a valid Pair[string, any].
func (it newPairCreator) StringAny(left string, right any) *Pair[string, any] {
	return NewPair(left, right)
}

// IntInt creates a valid Pair[int, int].
func (it newPairCreator) IntInt(left, right int) *Pair[int, int] {
	return NewPair(left, right)
}

// IntString creates a valid Pair[int, string].
func (it newPairCreator) IntString(left int, right string) *Pair[int, string] {
	return NewPair(left, right)
}

// Any creates a valid Pair[any, any].
func (it newPairCreator) Any(left, right any) *Pair[any, any] {
	return NewPair(left, right)
}

// --- Invalid constructors ---

// InvalidStringString creates an invalid Pair[string, string] with a message.
func (it newPairCreator) InvalidStringString(message string) *Pair[string, string] {
	return InvalidPair[string, string](message)
}

// InvalidAny creates an invalid Pair[any, any] with a message.
func (it newPairCreator) InvalidAny(message string) *Pair[any, any] {
	return InvalidPair[any, any](message)
}

// --- String split constructors ---

// Split splits a string by separator into a Pair[string, string].
//
//	pair := coregeneric.New.Pair.Split("key=value", "=")
//	pair.Left  // "key"
//	pair.Right // "value"
func (it newPairCreator) Split(input, sep string) *Pair[string, string] {
	return PairFromSplit(input, sep)
}

// SplitTrimmed splits and trims whitespace from both parts.
func (it newPairCreator) SplitTrimmed(input, sep string) *Pair[string, string] {
	return PairFromSplitTrimmed(input, sep)
}

// SplitFull splits at the first separator only; Right gets everything after.
//
//	pair := coregeneric.New.Pair.SplitFull("a:b:c:d", ":")
//	pair.Left  // "a"
//	pair.Right // "b:c:d"
func (it newPairCreator) SplitFull(input, sep string) *Pair[string, string] {
	return PairFromSplitFull(input, sep)
}

// SplitFullTrimmed is SplitFull with trimmed output.
func (it newPairCreator) SplitFullTrimmed(input, sep string) *Pair[string, string] {
	return PairFromSplitFullTrimmed(input, sep)
}

// FromSlice creates a Pair from a string slice (first → Left, last → Right).
func (it newPairCreator) FromSlice(parts []string) *Pair[string, string] {
	return PairFromSlice(parts)
}

// --- Number division constructors ---

// DivideInt splits an int into two halves.
//
//	pair := coregeneric.New.Pair.DivideInt(11) // {5, 6}
func (it newPairCreator) DivideInt(value int) *Pair[int, int] {
	return PairDivide(value)
}

// DivideInt64 splits an int64 into two halves.
func (it newPairCreator) DivideInt64(value int64) *Pair[int64, int64] {
	return PairDivide(value)
}

// DivideFloat64 splits a float64 into two equal halves.
func (it newPairCreator) DivideFloat64(value float64) *Pair[float64, float64] {
	return PairDivide(value)
}

// DivideIntWeighted splits an int by ratio (0.0–1.0).
//
//	pair := coregeneric.New.Pair.DivideIntWeighted(100, 0.3)  // {30, 70}
func (it newPairCreator) DivideIntWeighted(value int, leftRatio float64) *Pair[int, int] {
	return PairDivideWeighted(value, leftRatio)
}

// DivideFloat64Weighted splits a float64 by ratio (0.0–1.0).
func (it newPairCreator) DivideFloat64Weighted(value float64, leftRatio float64) *Pair[float64, float64] {
	return PairDivideWeighted(value, leftRatio)
}
