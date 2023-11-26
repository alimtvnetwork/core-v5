package codestack

func (it newCreator) CollectionByCap(capacity int) *TraceCollection {
	slice := make([]Trace, 0, capacity)

	return &TraceCollection{
		slice,
	}
}

func (it newCreator) Collection() *TraceCollection {
	return it.CollectionByCap(DefaultStackCount + 5)
}

func (it newCreator) CollectionUsing(
	isClone bool,
	traces ...Trace,
) *TraceCollection {
	if traces == nil {
		return it.EmptyTraces()
	}

	if !isClone {
		return &TraceCollection{
			traces,
		}
	}

	slice := it.CollectionByCap(len(traces))

	return slice.Adds(traces...)
}

func (it newCreator) EmptyTraces() *TraceCollection {
	return it.CollectionByCap(0)
}
