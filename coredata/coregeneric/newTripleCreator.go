package coregeneric

// newTripleCreator provides factory methods for Triple[A, B, C] types.
//
// Usage:
//
//	coregeneric.New.Triple.StringStringString("left", "mid", "right")
//	coregeneric.New.Triple.Any("a", 42, true)
type newTripleCreator struct{}

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

// InvalidStringStringString creates an invalid Triple[string, string, string] with a message.
func (it newTripleCreator) InvalidStringStringString(message string) *Triple[string, string, string] {
	return InvalidTriple[string, string, string](message)
}

// InvalidAny creates an invalid Triple[any, any, any] with a message.
func (it newTripleCreator) InvalidAny(message string) *Triple[any, any, any] {
	return InvalidTriple[any, any, any](message)
}
