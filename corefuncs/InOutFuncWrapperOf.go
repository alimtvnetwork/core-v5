package corefuncs

import "github.com/alimtvnetwork/core/errcore"

// InOutFuncWrapperOf is the generic version of an input-output function wrapper.
//
// It wraps a typed InOutFuncOf[TIn, TOut] with a name for identification
// and provides conversion to ActionFunc/ActionReturnsErrorFunc.
//
// Usage:
//
//	wrapper := corefuncs.InOutFuncWrapperOf[string, int]{
//	    Name:   "strlen",
//	    Action: func(s string) int { return len(s) },
//	}
//	result := wrapper.Exec("hello") // result is int(5)
type InOutFuncWrapperOf[TIn any, TOut any] struct {
	Name   string
	Action InOutFuncOf[TIn, TOut]
}

// Exec runs the wrapped function with the given typed input.
func (it InOutFuncWrapperOf[TIn, TOut]) Exec(input TIn) TOut {
	return it.Action(input)
}

// AsActionFunc returns an ActionFunc that executes with the given input (discarding output).
func (it InOutFuncWrapperOf[TIn, TOut]) AsActionFunc(input TIn) ActionFunc {
	return func() {
		it.Action(input)
	}
}

// AsActionReturnsErrorFunc returns an ActionReturnsErrorFunc (never errors).
func (it InOutFuncWrapperOf[TIn, TOut]) AsActionReturnsErrorFunc(
	input TIn,
) ActionReturnsErrorFunc {
	return func() error {
		it.Action(input)

		return nil
	}
}

// ToLegacy converts to the non-generic InOutErrFuncWrapper for backward compatibility.
func (it InOutFuncWrapperOf[TIn, TOut]) ToLegacy() InOutErrFuncWrapper {
	return InOutErrFuncWrapper{
		Name: it.Name,
		Action: func(input any) (any, error) {
			return it.Action(input.(TIn)), nil
		},
	}
}

// SerializeOutputFuncWrapperOf wraps a typed serialization function.
//
// Usage:
//
//	wrapper := corefuncs.SerializeOutputFuncWrapperOf[MyStruct]{
//	    Name:   "json-marshal",
//	    Action: func(m MyStruct) ([]byte, error) { return json.Marshal(m) },
//	}
//	bytes, err := wrapper.Exec(myStruct)
type SerializeOutputFuncWrapperOf[TIn any] struct {
	Name   string
	Action SerializeOutputFuncOf[TIn]
}

// Exec runs the serialization function with typed input.
func (it SerializeOutputFuncWrapperOf[TIn]) Exec(input TIn) ([]byte, error) {
	return it.Action(input)
}

// AsActionReturnsErrorFunc returns an ActionReturnsErrorFunc that captures the input.
func (it SerializeOutputFuncWrapperOf[TIn]) AsActionReturnsErrorFunc(
	input TIn,
) ActionReturnsErrorFunc {
	return func() error {
		_, err := it.Action(input)

		if err != nil {
			return errcore.
				FailedToExecuteType.
				Error(err.Error()+", function name:", it.Name)
		}

		return nil
	}
}
