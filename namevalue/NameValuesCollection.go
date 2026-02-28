package namevalue

import (
	"gitlab.com/auk-go/core/constants"
)

// NameValuesCollection is the backward-compatible alias for Collection[string, any].
type NameValuesCollection = Collection[string, any]

func NewNameValuesCollection(capacity int) *NameValuesCollection {
	return NewGenericCollection[string, any](capacity)
}

func NewCollection() *NameValuesCollection {
	return NewGenericCollection[string, any](constants.Capacity5)
}

func NewNewNameValuesCollectionUsing(
	isClone bool,
	items ...StringAny,
) *NameValuesCollection {
	return NewGenericCollectionUsing[string, any](isClone, items...)
}

func EmptyNameValuesCollection() *NameValuesCollection {
	return NewGenericCollection[string, any](0)
}
