package coreinstruction

import (
	"gitlab.com/auk-go/core/constants"
)

type Identifiers struct {
	Ids []BaseIdentifier `json:"Ids,omitempty"`
}

func EmptyIdentifiers() *Identifiers {
	return &Identifiers{
		Ids: []BaseIdentifier{},
	}
}

func NewIdentifiersUsingCap(
	capacity int,
) *Identifiers {
	slice := make(
		[]BaseIdentifier,
		0,
		capacity)

	return &Identifiers{Ids: slice}
}

func NewIdentifiers(
	ids ...string,
) *Identifiers {
	slice := make(
		[]BaseIdentifier,
		len(ids))

	if len(ids) == 0 {
		return &Identifiers{
			Ids: []BaseIdentifier{},
		}
	}

	for i, id := range ids {
		slice[i] = BaseIdentifier{
			Id: id,
		}
	}

	return &Identifiers{
		Ids: slice,
	}
}

func (it *Identifiers) Length() int {
	if it == nil {
		return 0
	}

	return len(it.Ids)
}

func (it *Identifiers) IsEmpty() bool {
	return it.Length() == 0
}

func (it *Identifiers) IndexOf(id string) int {
	if id == constants.EmptyString || it.IsEmpty() {
		return constants.InvalidNotFoundCase
	}

	for index, baseIdentifier := range it.Ids {
		if baseIdentifier.Id == id {
			return index
		}
	}

	return constants.InvalidNotFoundCase
}

func (it *Identifiers) GetById(id string) *BaseIdentifier {
	if id == constants.EmptyString || it.IsEmpty() {
		return nil
	}

	for i := range it.Ids {
		if it.Ids[i].Id == id {
			return &it.Ids[i]
		}
	}

	return nil
}

func (it *Identifiers) Add(
	id string,
) *Identifiers {
	if id == constants.EmptyString {
		return it
	}

	it.Ids = append(
		it.Ids,
		BaseIdentifier{Id: id})

	return it
}

func (it *Identifiers) Adds(
	ids ...string,
) *Identifiers {
	if len(ids) == 0 {
		return it
	}

	for _, id := range ids {
		it.Ids = append(
			it.Ids,
			BaseIdentifier{Id: id})
	}

	return it
}

func (it *Identifiers) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *Identifiers) Clone() *Identifiers {
	length := it.Length()

	slice := make(
		[]BaseIdentifier,
		length)

	if length == 0 {
		return &Identifiers{
			Ids: slice,
		}
	}

	for i, baseIdentifier := range it.Ids {
		slice[i] = *baseIdentifier.Clone()
	}

	return &Identifiers{
		Ids: slice,
	}
}
