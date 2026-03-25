package args

import (
	"errors"
	"fmt"
)

// MustBeValid panics if the FuncWrap is nil or invalid.
func (it *FuncWrap[T]) MustBeValid() {
	if it == nil {
		panic("cannot execute on nil func-wrap")
	}

	if it.IsInvalid() {
		panic("func-wrap invalid - " + it.Name)
	}
}

// ValidationError returns an error if the FuncWrap is nil or invalid,
// or nil if it is valid.
func (it *FuncWrap[T]) ValidationError() error {
	if it == nil {
		return errors.New("cannot execute on nil func-wrap")
	}

	if it.IsInvalid() {
		return fmt.Errorf(
			"func-wrap is invalid:\n    given type: %T\n    name: %s",
			it.Func,
			it.Name,
		)
	}

	return nil
}

// ValidateMethodArgs validates that the given arguments match the
// expected count and types of the wrapped function's parameters.
func (it *FuncWrap[T]) ValidateMethodArgs(args []any) error {
	expectedCount := it.ArgsCount()
	given := len(args)

	if given != expectedCount {
		return errors.New(
			it.argsCountMismatchErrorMessage(
				expectedCount,
				given,
				args,
			),
		)
	}

	_, err := it.VerifyInArgs(args)

	return err
}

// InvalidError returns a descriptive error explaining why the FuncWrap is invalid,
// or nil if it is valid.
func (it *FuncWrap[T]) InvalidError() error {
	if it == nil {
		return errors.New("func-wrap is nil")
	}

	if !it.rv.IsValid() {
		return errors.New("reflect value is invalid")
	}

	if !it.HasValidFunc() {
		return errors.New("func-wrap request doesn't hold a valid func reflect")
	}

	return nil
}
