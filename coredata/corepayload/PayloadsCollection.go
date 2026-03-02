package corepayload

// PayloadsCollection is a collection of PayloadWrapper pointers.
//
// Getters and query methods are in PayloadsCollectionGetters.go.
// Filter and search methods are in PayloadsCollectionFilter.go.
// Paging methods are in PayloadsCollectionPaging.go.
// JSON serialization and string methods are in PayloadsCollectionJson.go.
// Mutation methods (Add, Remove, Clone, etc.) are below.
type PayloadsCollection struct {
	Items []*PayloadWrapper
}

// =============================================================================
// Mutation — Add
// =============================================================================

func (it *PayloadsCollection) Add(
	payloadWrapper PayloadWrapper,
) *PayloadsCollection {
	it.Items = append(
		it.Items,
		&payloadWrapper)

	return it
}

func (it *PayloadsCollection) Adds(
	payloadWrappers ...PayloadWrapper,
) *PayloadsCollection {
	if len(payloadWrappers) == 0 {
		return it
	}

	for i := 0; i < len(payloadWrappers); i++ {
		it.Items = append(
			it.Items,
			&payloadWrappers[i])
	}

	return it
}

func (it *PayloadsCollection) AddsPtr(
	payloadWrappers ...*PayloadWrapper,
) *PayloadsCollection {
	if len(payloadWrappers) == 0 {
		return it
	}

	it.Items = append(
		it.Items,
		payloadWrappers...)

	return it
}

func (it *PayloadsCollection) AddsPtrOptions(
	isSkipHasIssuedPayloads bool,
	payloadWrappers ...*PayloadWrapper,
) *PayloadsCollection {
	if len(payloadWrappers) == 0 {
		return it
	}

	for i := 0; i < len(payloadWrappers); i++ {
		item := payloadWrappers[i]

		if isSkipHasIssuedPayloads && item.HasIssuesOrEmpty() {
			continue
		}

		it.Items = append(
			it.Items,
			item)
	}

	return it
}

func (it *PayloadsCollection) AddsOptions(
	isSkipHasIssuedPayloads bool,
	payloadWrappers ...PayloadWrapper,
) *PayloadsCollection {
	if len(payloadWrappers) == 0 {
		return it
	}

	for i := 0; i < len(payloadWrappers); i++ {
		item := payloadWrappers[i]

		if isSkipHasIssuedPayloads && item.HasIssuesOrEmpty() {
			continue
		}

		it.Items = append(
			it.Items,
			&item)
	}

	return it
}

func (it *PayloadsCollection) AddsIf(
	isAdd bool,
	payloadWrappers ...PayloadWrapper,
) *PayloadsCollection {
	isSkip := !isAdd

	if isSkip {
		return it
	}

	return it.Adds(payloadWrappers...)
}

func (it *PayloadsCollection) InsertAt(
	index int,
	item PayloadWrapper,
) *PayloadsCollection {
	it.Items = append(it.Items[:index+1], it.Items[index:]...)
	it.Items[index] = &item

	return it
}

// =============================================================================
// Mutation — Concat, Reverse, Clone
// =============================================================================

func (it *PayloadsCollection) ConcatNew(
	additionalItems ...PayloadWrapper,
) *PayloadsCollection {
	cloned := it.Clone()

	return cloned.Adds(additionalItems...)
}

func (it *PayloadsCollection) ConcatNewPtr(
	additionalItemsPtr ...*PayloadWrapper,
) *PayloadsCollection {
	cloned := it.Clone()

	return cloned.AddsPtr(
		additionalItemsPtr...)
}

func (it *PayloadsCollection) Reverse() *PayloadsCollection {
	length := it.Length()

	if length <= 1 {
		return it
	}

	if length == 2 {
		it.Items[0], it.Items[1] = it.Items[1], it.Items[0]

		return it
	}

	mid := length / 2
	lastIndex := length - 1

	for i := 0; i < mid; i++ {
		it.Items[i], it.Items[lastIndex-i] =
			it.Items[lastIndex-i], it.Items[i]
	}

	return it
}

func (it PayloadsCollection) Clone() PayloadsCollection {
	list := New.PayloadsCollection.UsingCap(it.Length())

	return *list.AddsPtr(it.Items...)
}

func (it *PayloadsCollection) ClonePtr() *PayloadsCollection {
	if it == nil {
		return nil
	}

	list := New.PayloadsCollection.UsingCap(it.Length())

	return list.AddsPtr(it.Items...)
}

// =============================================================================
// Mutation — Clear, Dispose
// =============================================================================

func (it *PayloadsCollection) Clear() *PayloadsCollection {
	if it == nil {
		return it
	}

	tempItems := it.Items
	clearFunc := func() {
		for _, item := range tempItems {
			item.Dispose()
		}
	}

	go clearFunc()

	it.Items = []*PayloadWrapper{}

	return it
}

func (it *PayloadsCollection) Dispose() {
	if it == nil {
		return
	}

	it.Clear()
	it.Items = nil
}
