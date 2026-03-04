package args

import (
	"errors"
	"fmt"
)

func (it *FuncWrap) MustBeValid() {
	if it == nil {
		panic("cannot execute on nil func-wrap")
	}

	if it.IsInvalid() {
		panic("func-wrap invalid - " + it.Name)
	}
}

func (it *FuncWrap) ValidationError() error {
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

func (it *FuncWrap) ValidateMethodArgs(args []any) error {
	expectedCount := it.ArgsCount()
	given := len(args)

	if given != expectedCount {
		return errors.New(
			it.argsCountMismatchErrorMessage(expectedCount, given, args),
		)
	}

	_, err := it.VerifyInArgs(args)

	return err
}

func (it *FuncWrap) InvalidError() error {
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
