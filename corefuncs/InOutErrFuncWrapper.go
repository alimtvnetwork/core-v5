package corefuncs

import "github.com/alimtvnetwork/core/errcore"

type InOutErrFuncWrapper struct {
	Name   string
	Action InOutErrFunc
}

func (it InOutErrFuncWrapper) Exec(
	input any,
) (output any, err error) {
	return it.Action(input)
}

func (it InOutErrFuncWrapper) AsActionFunc(input any) ActionFunc {
	return func() {
		errcore.MustBeEmpty(
			it.AsActionReturnsErrorFunc(input)())
	}
}

func (it InOutErrFuncWrapper) AsActionReturnsErrorFunc(
	input any,
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
