package corefuncs

import "gitlab.com/auk-go/core/errcore"

// InOutErrFuncWrapperOf is the generic version of InOutErrFuncWrapper.
//
// It wraps a typed InOutErrFuncOf[TIn, TOut] with a name for error reporting
// and provides conversion to ActionFunc/ActionReturnsErrorFunc.
//
// Usage:
//
//	wrapper := corefuncs.InOutErrFuncWrapperOf[string, int]{
//	    Name:   "strlen",
//	    Action: func(s string) (int, error) { return len(s), nil },
//	}
//	result, err := wrapper.Exec("hello") // result is int(5)
type InOutErrFuncWrapperOf[TIn any, TOut any] struct {
	Name   string
	Action InOutErrFuncOf[TIn, TOut]
}

// Exec runs the wrapped function with the given typed input.
func (it InOutErrFuncWrapperOf[TIn, TOut]) Exec(
	input TIn,
) (output TOut, err error) {
	return it.Action(input)
}

// AsActionFunc returns an ActionFunc that executes with the given input.
// Panics on error via errcore.MustBeEmpty.
func (it InOutErrFuncWrapperOf[TIn, TOut]) AsActionFunc(input TIn) ActionFunc {
	return func() {
		errcore.MustBeEmpty(
			it.AsActionReturnsErrorFunc(input)())
	}
}

// AsActionReturnsErrorFunc returns an ActionReturnsErrorFunc that captures the input.
func (it InOutErrFuncWrapperOf[TIn, TOut]) AsActionReturnsErrorFunc(
	input TIn,
) ActionReturnsErrorFunc {
	return func() error {
		_, err := it.Action(input)

		if err != nil {
			return errcore.
				FailedToExecuteType.
				Error(err.Error()+", function name:", it.Name)
		}

		return err
	}
}

// ToLegacy converts to the non-generic InOutErrFuncWrapper for backward compatibility.
func (it InOutErrFuncWrapperOf[TIn, TOut]) ToLegacy() InOutErrFuncWrapper {
	return InOutErrFuncWrapper{
		Name: it.Name,
		Action: func(input any) (any, error) {
			return it.Action(input.(TIn))
		},
	}
}
