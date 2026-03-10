package corepayload

import (
	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corestr"
)

type emptyCreator struct{}

func (it *emptyCreator) Attributes() *Attributes {
	return &Attributes{}
}

func (it *emptyCreator) AttributesDefaults() *Attributes {
	return &Attributes{
		KeyValuePairs:    corestr.Empty.Hashmap(),
		AnyKeyValuePairs: coredynamic.EmptyMapAnyItems(),
		DynamicPayloads:  []byte{},
	}
}

func (it *emptyCreator) PayloadWrapper() *PayloadWrapper {
	return &PayloadWrapper{}
}

func (it *emptyCreator) PayloadsCollection() *PayloadsCollection {
	return &PayloadsCollection{
		Items: []*PayloadWrapper{},
	}
}
