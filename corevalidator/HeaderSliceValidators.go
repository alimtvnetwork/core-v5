package corevalidator

import (
	"log/slog"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/errcore"
)

type HeaderSliceValidators []HeaderSliceValidator

func (it HeaderSliceValidators) Length() int {
	if it == nil {
		return 0
	}

	return len(it)
}

func (it HeaderSliceValidators) IsEmpty() bool {
	return it == nil || len(it) == 0
}

func (it HeaderSliceValidators) SetActualOnAll(actualLines ...string) {
	if it.IsEmpty() {
		return
	}

	for _, sliceValidator := range it {
		sliceValidator.SetActual(actualLines)
	}
}

func (it HeaderSliceValidators) IsValid(
	isCaseSensitive bool,
) bool {
	return it.IsMatch(isCaseSensitive)
}

func (it HeaderSliceValidators) IsMatch(
	isCaseSensitive bool,
) bool {
	if it.IsEmpty() {
		return true
	}

	for _, sliceValidator := range it {
		if !sliceValidator.IsValid(isCaseSensitive) {
			return false
		}
	}

	return true
}

func (it HeaderSliceValidators) VerifyAll(
	header string,
	params *Parameter,
	isPrintError bool,
) error {
	if it.IsEmpty() {
		return nil
	}

	errs := corestr.New.SimpleSlice.Cap(constants.Capacity2)

	for _, sliceValidator := range it {
		err := sliceValidator.AllVerifyError(params)

		if err != nil {
			diffMsg := errcore.LineDiffToString(
				params.CaseIndex,
				params.Header,
				sliceValidator.ActualLines,
				sliceValidator.ExpectedLines,
			)

			errs.AddError(err)

			if len(diffMsg) > 0 {
				errs.Add(diffMsg)
			}
		}
	}

	if errs.IsEmpty() {
		return nil
	}

	errs.InsertAt(0, header)
	err := errs.AsDefaultError()

	if isPrintError {
		slog.Error("verification failed", "error", err)
	}

	return err
}

func (it HeaderSliceValidators) AssertVerifyAll(
	t *testing.T,
	params *Parameter,
) {
	if it.IsEmpty() {
		return
	}

	finalError := it.VerifyAllError(params)

	convey.Convey(params.Header, t, func() {
		convey.So(
			finalError,
			should.BeNil)
	})
}

func (it HeaderSliceValidators) VerifyAllError(
	params *Parameter,
) error {
	if it.IsEmpty() {
		return nil
	}

	errs := corestr.New.SimpleSlice.Cap(constants.Capacity2)

	for _, sliceValidator := range it {
		err := sliceValidator.AllVerifyError(params)

		if err != nil {
			diffMsg := errcore.LineDiffToString(
				params.CaseIndex,
				params.Header,
				sliceValidator.ActualLines,
				sliceValidator.ExpectedLines,
			)

			errs.AddError(err)

			if len(diffMsg) > 0 {
				errs.Add(diffMsg)
			}
		}
	}

	header := params.Header

	errs.InsertAt(0, header)

	return errs.AsDefaultError()
}

func (it HeaderSliceValidators) AssertVerifyAllUsingActual(
	t *testing.T,
	params *Parameter,
	actualLines ...string,
) {
	if it.IsEmpty() {
		return
	}

	finalError := it.VerifyAllErrorUsingActual(
		params,
		actualLines...)

	convey.Convey(params.Header, t, func() {
		convey.So(
			finalError,
			should.BeNil)
	})
}

func (it HeaderSliceValidators) VerifyAllErrorUsingActual(
	params *Parameter,
	actualLines ...string,
) error {
	if it.IsEmpty() {
		return nil
	}

	errs := corestr.New.SimpleSlice.Cap(constants.Capacity2)

	for _, sliceValidator := range it {
		sliceValidator.SetActual(actualLines)
		err := sliceValidator.AllVerifyError(params)

		if err != nil {
			diffMsg := errcore.LineDiffToString(
				params.CaseIndex,
				params.Header,
				sliceValidator.ActualLines,
				sliceValidator.ExpectedLines,
			)

			errs.AddError(err)

			if len(diffMsg) > 0 {
				errs.Add(diffMsg)
			}
		}
	}

	header := params.Header

	errs.InsertAt(0, header)

	return errs.AsDefaultError()
}

// VerifyFirst
//
// Only collect using the SliceValidator.VerifyFirstError
func (it HeaderSliceValidators) VerifyFirst(
	params *Parameter,
	isPrintError bool,
) error {
	if it.IsEmpty() {
		return nil
	}

	errs := corestr.New.SimpleSlice.Cap(constants.Capacity2)

	for _, sliceValidator := range it {
		err := sliceValidator.VerifyFirstError(params)

		if err != nil {
			diffMsg := errcore.LineDiffToString(
				params.CaseIndex,
				params.Header,
				sliceValidator.ActualLines,
				sliceValidator.ExpectedLines,
			)

			errs.AddError(err)

			if len(diffMsg) > 0 {
				errs.Add(diffMsg)
			}
		}
	}

	if errs.IsEmpty() {
		return nil
	}

	header := params.Header

	errs.InsertAt(0, header)
	err := errs.AsDefaultError()

	if isPrintError {
		slog.Error("verification failed", "error", err)
	}

	return err
}

func (it HeaderSliceValidators) VerifyUpto(
	isPrintErr,
	isFirstOnly bool,
	length int,
	params *Parameter,
) error {
	if it.IsEmpty() {
		return nil
	}

	errs := corestr.New.SimpleSlice.Cap(constants.Capacity2)

	for _, sliceValidator := range it {
		err := sliceValidator.AllVerifyErrorUptoLength(
			isFirstOnly,
			params,
			length)

		if err != nil {
			diffMsg := errcore.LineDiffToString(
				params.CaseIndex,
				params.Header,
				sliceValidator.ActualLines,
				sliceValidator.ExpectedLines,
			)

			errs.AddError(err)

			if len(diffMsg) > 0 {
				errs.Add(diffMsg)
			}
		}
	}

	if errs.IsEmpty() {
		return nil
	}

	errs.InsertAt(0, params.Header)
	err := errs.AsDefaultError()

	if isPrintErr {
		slog.Error("verification failed", "error", err)
	}

	return err
}
