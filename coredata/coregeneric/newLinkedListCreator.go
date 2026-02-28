package coregeneric

// newLinkedListCreator aggregates per-type linked list creators.
//
// Usage:
//
//	coregeneric.New.LinkedList.String.Empty()
//	coregeneric.New.LinkedList.Int.From([]int{1, 2, 3})
//	coregeneric.New.LinkedList.Float64.Items(1.0, 2.5, 3.7)
type newLinkedListCreator struct {
	String  typedLinkedListCreator[string]
	Int     typedLinkedListCreator[int]
	Int8    typedLinkedListCreator[int8]
	Int16   typedLinkedListCreator[int16]
	Int32   typedLinkedListCreator[int32]
	Int64   typedLinkedListCreator[int64]
	Uint    typedLinkedListCreator[uint]
	Float32 typedLinkedListCreator[float32]
	Float64 typedLinkedListCreator[float64]
	Byte    typedLinkedListCreator[byte]
	Bool    typedLinkedListCreator[bool]
	Any     typedLinkedListCreator[any]
}
