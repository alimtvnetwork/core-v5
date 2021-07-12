package coreinstruction

import (
	"gitlab.com/evatix-go/core/reqtype"
)

type BaseModifyAs struct {
	ModifyAs reqtype.Request `json:"ModifyAs"`
}

func NewModifyAs(modifyAs reqtype.Request) BaseModifyAs {
	return BaseModifyAs{
		ModifyAs: modifyAs,
	}
}

func (b *BaseModifyAs) SetModifyAs(modifyAs reqtype.Request) {
	b.ModifyAs = modifyAs
}
