package coregeneric

// newHashmapCreator aggregates per-type hashmap creators.
//
// Usage:
//
//	coregeneric.New.Hashmap.StringString.Cap(10)
//	coregeneric.New.Hashmap.StringInt.Empty()
//	coregeneric.New.Hashmap.StringAny.From(existingMap)
type newHashmapCreator struct {
	StringString  typedHashmapCreator[string, string]
	StringInt     typedHashmapCreator[string, int]
	StringInt64   typedHashmapCreator[string, int64]
	StringFloat64 typedHashmapCreator[string, float64]
	StringBool    typedHashmapCreator[string, bool]
	StringAny     typedHashmapCreator[string, any]
	IntString     typedHashmapCreator[int, string]
	IntInt        typedHashmapCreator[int, int]
	IntAny        typedHashmapCreator[int, any]
}
