package corefuncs

// New is the root aggregator for the New Creator pattern in corefuncs.
//
// Usage:
//
//	// Legacy wrappers via New creator
//	wrapper := corefuncs.New.ActionErr("cleanup", myFunc)
//	wrapper := corefuncs.New.IsSuccess("check", myFunc)
//	wrapper := corefuncs.New.NamedAction("log", myFunc)
//	wrapper := corefuncs.New.LegacyInOutErr("transform", myFunc)
//	wrapper := corefuncs.New.LegacyResultDelegating("unmarshal", myFunc)
//
//	// Generic wrappers via package-level constructors
//	wrapper := corefuncs.NewInOutErrWrapper[string, int]("parse", myFunc)
//	wrapper := corefuncs.NewInOutWrapper[string, int]("convert", myFunc)
//	wrapper := corefuncs.NewInActionErrWrapper[string]("validate", myFunc)
//	wrapper := corefuncs.NewResultDelegatingWrapper[*MyStruct]("unmarshal", myFunc)
//	wrapper := corefuncs.NewSerializeWrapper[MyStruct]("marshal", myFunc)
var New = &newFuncCreator{}

// newFuncCreator is the root aggregator for function wrapper creation.
type newFuncCreator struct{}

// =============================================================================
// Legacy (any-based) Wrapper Creators
// =============================================================================

// ActionErr creates a legacy ActionReturnsErrorFuncWrapper.
func (it *newFuncCreator) ActionErr(
	name string,
	action ActionReturnsErrorFunc,
) ActionReturnsErrorFuncWrapper {
	return ActionReturnsErrorFuncWrapper{
		Name:   name,
		Action: action,
	}
}

// IsSuccess creates a legacy IsSuccessFuncWrapper.
func (it *newFuncCreator) IsSuccess(
	name string,
	action IsSuccessFunc,
) IsSuccessFuncWrapper {
	return IsSuccessFuncWrapper{
		Name:   name,
		Action: action,
	}
}

// NamedAction creates a legacy NamedActionFuncWrapper.
func (it *newFuncCreator) NamedAction(
	name string,
	action NamedActionFunc,
) NamedActionFuncWrapper {
	return NamedActionFuncWrapper{
		Name:   name,
		Action: action,
	}
}

// LegacyInOutErr creates a legacy InOutErrFuncWrapper.
func (it *newFuncCreator) LegacyInOutErr(
	name string,
	action InOutErrFunc,
) InOutErrFuncWrapper {
	return InOutErrFuncWrapper{
		Name:   name,
		Action: action,
	}
}

// LegacyResultDelegating creates a legacy ResultDelegatingFuncWrapper.
func (it *newFuncCreator) LegacyResultDelegating(
	name string,
	action ResultDelegatingFunc,
) ResultDelegatingFuncWrapper {
	return ResultDelegatingFuncWrapper{
		Name:   name,
		Action: action,
	}
}

// =============================================================================
// Generic Wrapper Constructors (package-level, type-parameterized)
// =============================================================================

// NewInOutErrWrapper creates a InOutErrFuncWrapperOf[TIn, TOut].
func NewInOutErrWrapper[TIn any, TOut any](
	name string,
	action InOutErrFuncOf[TIn, TOut],
) InOutErrFuncWrapperOf[TIn, TOut] {
	return InOutErrFuncWrapperOf[TIn, TOut]{
		Name:   name,
		Action: action,
	}
}

// NewResultDelegatingWrapper creates a ResultDelegatingFuncWrapperOf[T].
func NewResultDelegatingWrapper[T any](
	name string,
	action ResultDelegatingFuncOf[T],
) ResultDelegatingFuncWrapperOf[T] {
	return ResultDelegatingFuncWrapperOf[T]{
		Name:   name,
		Action: action,
	}
}

// NewInActionErrWrapper creates an InActionReturnsErrFuncWrapperOf[TIn].
func NewInActionErrWrapper[TIn any](
	name string,
	action InActionReturnsErrFuncOf[TIn],
) InActionReturnsErrFuncWrapperOf[TIn] {
	return InActionReturnsErrFuncWrapperOf[TIn]{
		Name:   name,
		Action: action,
	}
}

// NewInOutWrapper creates an InOutFuncWrapperOf[TIn, TOut].
func NewInOutWrapper[TIn any, TOut any](
	name string,
	action InOutFuncOf[TIn, TOut],
) InOutFuncWrapperOf[TIn, TOut] {
	return InOutFuncWrapperOf[TIn, TOut]{
		Name:   name,
		Action: action,
	}
}

// NewSerializeWrapper creates a SerializeOutputFuncWrapperOf[TIn].
func NewSerializeWrapper[TIn any](
	name string,
	action SerializeOutputFuncOf[TIn],
) SerializeOutputFuncWrapperOf[TIn] {
	return SerializeOutputFuncWrapperOf[TIn]{
		Name:   name,
		Action: action,
	}
}
