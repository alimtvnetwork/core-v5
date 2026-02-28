package coredynamic

// newAnyCollectionCreator embeds the generic creator for Collection[any].
//
// This is a concrete Collection[any] creator — not a parameterized generic.
// For the true generic factory, see newGenericCollectionCreator[T].
type newAnyCollectionCreator struct {
	newGenericCollectionCreator[any]
}
