package corefuncs

// New is the root aggregator for the New Creator pattern in corefuncs.
//
// Usage:
//
//	// Create wrappers via the New creator
//	wrapper := corefuncs.New.InOutErr("transform", myFunc)
//	wrapper := corefuncs.New.ResultDelegating("unmarshal", myFunc)
//	wrapper := corefuncs.New.InActionErr("validate", myFunc)
//	wrapper := corefuncs.New.InOut("convert", myFunc)
//	wrapper := corefuncs.New.ActionErr("cleanup", myFunc)
//	wrapper := corefuncs.New.IsSuccess("check", myFunc)
//	wrapper := corefuncs.New.NamedAction("log", myFunc)
//	wrapper := corefuncs.New.Serialize("marshal", myFunc)
var New = &newFuncCreator{}

// newFuncCreator is the root aggregator for function wrapper creation.
type newFuncCreator struct{}

// =============================================================================
// Generic Typed Wrapper Creators
// =============================================================================

// InOutErr creates a generic InOutErrFuncWrapperOf[TIn, TOut].
func (it *newFuncCreator) InOutErr() inOutErrCreator {
	return inOutErrCreator{}
}

// ResultDelegating creates a generic ResultDelegatingFuncWrapperOf[T].
func (it *newFuncCreator) ResultDelegating() resultDelegatingCreator {
	return resultDelegatingCreator{}
}

// InActionErr creates a generic InActionReturnsErrFuncWrapperOf[TIn].
func (it *newFuncCreator) InActionErr() inActionErrCreator {
	return inActionErrCreator{}
}

// InOut creates a generic InOutFuncWrapperOf[TIn, TOut].
func (it *newFuncCreator) InOut() inOutCreator {
	return inOutCreator{}
}

// Serialize creates a generic SerializeOutputFuncWrapperOf[TIn].
func (it *newFuncCreator) Serialize() serializeCreator {
	return serializeCreator{}
}

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
// Sub-creators for generic wrappers (needed because Go can't infer type params
// on struct constructors — these enable: New.InOutErr().Of("name", fn))
// =============================================================================

type inOutErrCreator struct{}

// Of creates a InOutErrFuncWrapperOf[TIn, TOut].
// Usage: corefuncs.New.InOutErr().Of("name", fn) — note: requires explicit type params at call site.
// Prefer NewInOutErrWrapper[TIn, TOut]("name", fn) for cleaner syntax.
func NewInOutErrWrapper[TIn any, TOut any](
	name string,
	action InOutErrFuncOf[TIn, TOut],
) InOutErrFuncWrapperOf[TIn, TOut] {
	return InOutErrFuncWrapperOf[TIn, TOut]{
		Name:   name,
		Action: action,
	}
}

type resultDelegatingCreator struct{}

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

type inActionErrCreator struct{}

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

type inOutCreator struct{}

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

type serializeCreator struct{}

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
