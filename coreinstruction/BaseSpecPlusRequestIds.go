package coreinstruction

type BaseSpecPlusRequestIds struct {
	Specification *Specification          `json:"Specification,omitempty"`
	RequestIds    *IdentifierWithIsGlobal `json:"RequestIds,omitempty"`
}

func NewBaseSpecPlusRequestIds(
	spec *Specification,
	reqIds *IdentifierWithIsGlobal,
) *BaseSpecPlusRequestIds {
	return &BaseSpecPlusRequestIds{
		Specification: spec,
		RequestIds:    reqIds,
	}
}

func (b *BaseSpecPlusRequestIds) HasSpec() bool {
	return b != nil && b.Specification != nil
}

func (b *BaseSpecPlusRequestIds) HasRequestIds() bool {
	return b != nil && b.RequestIds != nil
}

func (b *BaseSpecPlusRequestIds) Clone() *BaseSpecPlusRequestIds {
	return &BaseSpecPlusRequestIds{
		Specification: b.Specification.Clone(),
		RequestIds:    b.RequestIds.Clone(),
	}
}
