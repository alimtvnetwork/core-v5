package coregeneric

// New is the root aggregator for the New Creator pattern in coregeneric.
//
// Usage:
//
//	coregeneric.New.Collection.String.Cap(10)
//	coregeneric.New.Hashset.Int.Empty()
//	coregeneric.New.Hashmap.StringString.Cap(20)
//	coregeneric.New.SimpleSlice.Float64.Items(1.0, 2.5, 3.7)
//	coregeneric.New.LinkedList.String.Empty()
var New = &newCreator{}
