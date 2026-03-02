package coregeneric

// newPairCreator provides factory methods for Pair[L, R] types.
//
// Usage:
//
//	coregeneric.New.Pair.StringString("key", "value")
//	coregeneric.New.Pair.StringInt("name", 42)
//	coregeneric.New.Pair.Any("left", "right")
type newPairCreator struct{}

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

// InvalidStringString creates an invalid Pair[string, string] with a message.
func (it newPairCreator) InvalidStringString(message string) *Pair[string, string] {
	return InvalidPair[string, string](message)
}

// InvalidAny creates an invalid Pair[any, any] with a message.
func (it newPairCreator) InvalidAny(message string) *Pair[any, any] {
	return InvalidPair[any, any](message)
}
