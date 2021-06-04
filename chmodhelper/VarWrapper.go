package chmodhelper

type VarWrapper struct {
	rawInput            string
	isFixedType         bool
	Owner, Group, Other VarAttribute
}

func (varWrapper *VarWrapper) IsFixedType() bool {
	return varWrapper.isFixedType
}

func (varWrapper *VarWrapper) ToCompileFixedPtr() *Wrapper {
	if varWrapper.IsFixedType() {
		return varWrapper.ToCompileWrapperPtr(nil)
	}

	return nil
}

// ToCompileWrapper if Fixed type then fixed input can be nil.
func (varWrapper *VarWrapper) ToCompileWrapper(fixed *Wrapper) Wrapper {
	return *varWrapper.ToCompileWrapperPtr(fixed)
}

// ToCompileWrapperPtr if Fixed type then fixed input can be nil.
func (varWrapper *VarWrapper) ToCompileWrapperPtr(fixed *Wrapper) *Wrapper {
	if varWrapper.IsFixedType() {
		return &Wrapper{
			Owner: *varWrapper.Owner.ToCompileFixAttr(),
			Group: *varWrapper.Group.ToCompileFixAttr(),
			Other: *varWrapper.Other.ToCompileFixAttr(),
		}
	}

	return &Wrapper{
		Owner: varWrapper.Owner.ToCompileAttr(&fixed.Owner),
		Group: varWrapper.Group.ToCompileAttr(&fixed.Group),
		Other: varWrapper.Other.ToCompileAttr(&fixed.Other),
	}
}
