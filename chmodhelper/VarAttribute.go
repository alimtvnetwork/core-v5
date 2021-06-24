package chmodhelper

import (
	"gitlab.com/evatix-go/core/issetter"
)

type VarAttribute struct {
	rawInput    string
	isFixedType bool
	isRead      issetter.Value
	isWrite     issetter.Value
	isExecute   issetter.Value
}

func (varAttribute *VarAttribute) IsFixedType() bool {
	return varAttribute.isFixedType
}

func (varAttribute *VarAttribute) HasWildcard() bool {
	return !varAttribute.isFixedType
}

// ToCompileFixAttr must check IsFixedType, before calling.
func (varAttribute *VarAttribute) ToCompileFixAttr() *Attribute {
	if varAttribute.isFixedType {
		return &Attribute{
			IsRead:    varAttribute.isRead.IsTrue(),
			IsWrite:   varAttribute.isWrite.IsTrue(),
			IsExecute: varAttribute.isExecute.IsTrue(),
		}
	}

	return nil
}

// ToCompileAttr if fixed type then fixed param can be nil
func (varAttribute *VarAttribute) ToCompileAttr(fixed *Attribute) Attribute {
	if varAttribute.isFixedType {
		return Attribute{
			IsRead:    varAttribute.isRead.IsTrue(),
			IsWrite:   varAttribute.isWrite.IsTrue(),
			IsExecute: varAttribute.isExecute.IsTrue(),
		}
	}

	return Attribute{
		IsRead:    varAttribute.isRead.WildcardApply(fixed.IsRead),
		IsWrite:   varAttribute.isWrite.WildcardApply(fixed.IsWrite),
		IsExecute: varAttribute.isExecute.WildcardApply(fixed.IsExecute),
	}
}

func (varAttribute *VarAttribute) Clone() *VarAttribute {
	if varAttribute == nil {
		return nil
	}

	return &VarAttribute{
		rawInput:    varAttribute.rawInput,
		isFixedType: varAttribute.IsFixedType(),
		isRead:      varAttribute.isRead,
		isWrite:     varAttribute.isWrite,
		isExecute:   varAttribute.isExecute,
	}
}

func (varAttribute *VarAttribute) IsEqualPtr(next *VarAttribute) bool {
	if varAttribute == nil && next == nil {
		return true
	}

	if varAttribute == nil || next == nil {
		return false
	}

	isRead := next.isRead == varAttribute.isRead
	isWrite := next.isWrite == varAttribute.isWrite
	isExecute := next.isExecute == varAttribute.isExecute

	return isRead &&
		isWrite &&
		isExecute
}

func (varAttribute *VarAttribute) String() string {
	return varAttribute.rawInput
}
