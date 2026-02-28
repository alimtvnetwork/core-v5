package coregeneric

// newSimpleSliceCreator aggregates per-type simple slice creators.
//
// Usage:
//
//	coregeneric.New.SimpleSlice.String.Cap(10)
//	coregeneric.New.SimpleSlice.Int.Items(1, 2, 3)
//	coregeneric.New.SimpleSlice.Float64.From(existingSlice)
type newSimpleSliceCreator struct {
	String  typedSimpleSliceCreator[string]
	Int     typedSimpleSliceCreator[int]
	Int8    typedSimpleSliceCreator[int8]
	Int16   typedSimpleSliceCreator[int16]
	Int32   typedSimpleSliceCreator[int32]
	Int64   typedSimpleSliceCreator[int64]
	Uint    typedSimpleSliceCreator[uint]
	Uint8   typedSimpleSliceCreator[uint8]
	Uint16  typedSimpleSliceCreator[uint16]
	Uint32  typedSimpleSliceCreator[uint32]
	Uint64  typedSimpleSliceCreator[uint64]
	Float32 typedSimpleSliceCreator[float32]
	Float64 typedSimpleSliceCreator[float64]
	Byte    typedSimpleSliceCreator[byte]
	Bool    typedSimpleSliceCreator[bool]
	Any     typedSimpleSliceCreator[any]
}

// newSimpleSliceCreatorNote:
// SimpleSlice uses [T any] constraint because it wraps a raw []T.
// For ordered operations (Sort, Min, Max), use the package-level
// SortSimpleSlice[T cmp.Ordered]() functions.
