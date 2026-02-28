package corefuncs

import "gitlab.com/auk-go/core/errcore"

// InActionReturnsErrFuncWrapperOf is the generic version of an input-action-error wrapper.
//
// It wraps a typed InActionReturnsErrFuncOf[TIn] with a name for error reporting
// and provides conversion to ActionFunc/ActionReturnsErrorFunc.
//
// Usage:
//
//	wrapper := corefuncs.InActionReturnsErrFuncWrapperOf[string]{
//	    Name:   "validate-email",
//	    Action: func(email string) error { return validateEmail(email) },
//	}
//	err := wrapper.Exec("test@example.com")
type InActionReturnsErrFuncWrapperOf[TIn any] struct {
	Name   string
	Action InActionReturnsErrFuncOf[TIn]
}

// Exec runs the wrapped function with the given typed input.
func (it InActionReturnsErrFuncWrapperOf[TIn]) Exec(input TIn) error {
	return it.Action(input)
}

// AsActionFunc returns an ActionFunc that executes with the given input.
// Panics on error via errcore.MustBeEmpty.
func (it InActionReturnsErrFuncWrapperOf[TIn]) AsActionFunc(input TIn) ActionFunc {
	return func() {
		errcore.MustBeEmpty(
			it.AsActionReturnsErrorFunc(input)())
	}
}

// AsActionReturnsErrorFunc returns an ActionReturnsErrorFunc that captures the input.
func (it InActionReturnsErrFuncWrapperOf[TIn]) AsActionReturnsErrorFunc(
	input TIn,
) ActionReturnsErrorFunc {
	return func() error {
		err := it.Action(input)

		if err != nil {
			return errcore.
				FailedToExecuteType.
				Error(err.Error()+", function name:", it.Name)
		}

		return nil
	}
}

// ToLegacy converts to a legacy any-based wrapper for backward compatibility.
func (it InActionReturnsErrFuncWrapperOf[TIn]) ToLegacy() InOutErrFuncWrapper {
	return InOutErrFuncWrapper{
		Name: it.Name,
		Action: func(input any) (any, error) {
			return nil, it.Action(input.(TIn))
		},
	}
}
