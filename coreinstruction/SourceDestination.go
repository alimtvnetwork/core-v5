package coreinstruction

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

type SourceDestination struct {
	Source      string `json:"Source,omitempty"`
	Destination string `json:"Destination,omitempty"`
}

func (it *SourceDestination) IsNull() bool {
	return it == nil
}

func (it *SourceDestination) IsSourceEmpty() bool {
	return it == nil || it.Source == ""
}

func (it *SourceDestination) IsDestinationEmpty() bool {
	return it == nil || it.Destination == ""
}

func (it SourceDestination) String() string {
	return fmt.Sprintf(
		constants.SourceDestinationFormat,
		it.Source,
		it.Destination)
}

func (it SourceDestination) FromName() string {
	return it.Source
}

func (it SourceDestination) ToName() string {
	return it.Destination
}

func (it *SourceDestination) SetFromName(from string) {
	if it == nil {
		return
	}

	it.Source = from
}

func (it *SourceDestination) SetToName(to string) {
	if it == nil {
		return
	}

	it.Destination = to
}

func (it *SourceDestination) FromTo() *FromTo {
	if it == nil {
		return nil
	}

	return &FromTo{
		From: it.Source,
		To:   it.Destination,
	}
}

func (it *SourceDestination) Rename() *Rename {
	if it == nil {
		return nil
	}

	return &Rename{
		Existing: it.Source,
		New:      it.Destination,
	}
}

func (it *SourceDestination) Clone() *SourceDestination {
	if it == nil {
		return nil
	}

	return &SourceDestination{
		Source:      it.Source,
		Destination: it.Destination,
	}
}
