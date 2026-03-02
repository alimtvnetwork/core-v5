package coregeneric

// newCreator is the root aggregator for the New Creator pattern.
//
// Usage:
//
//	coregeneric.New.Collection.String.Cap(10)
//	coregeneric.New.Hashset.Int.Empty()
//	coregeneric.New.Hashmap.StringString.Cap(20)
//	coregeneric.New.SimpleSlice.Float64.Items(1.0, 2.5)
//	coregeneric.New.LinkedList.String.Empty()
//	coregeneric.New.Pair.StringString("key", "value")
//	coregeneric.New.Triple.StringStringString("a", "b", "c")
type newCreator struct {
	Collection  newCollectionCreator
	Hashset     newHashsetCreator
	Hashmap     newHashmapCreator
	SimpleSlice newSimpleSliceCreator
	LinkedList  newLinkedListCreator
	Pair        newPairCreator
	Triple      newTripleCreator
}
