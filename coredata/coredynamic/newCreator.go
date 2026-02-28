package coredynamic

// newCreator is the root aggregator for the New Creator pattern.
// Usage: coredynamic.New.Collection.String.Cap(10)
type newCreator struct {
	Collection newCollectionCreator
}
