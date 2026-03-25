package errcore

import (
	"github.com/alimtvnetwork/core/internal/strutilinternal"
)

func MeaningfulErrorWithData(
	rawErrType RawErrorType,
	funcName string,
	err error,
	data any,
) error {
	if err == nil {
		return nil
	}

	return rawErrType.Error(
		funcName,
		err.Error()+strutilinternal.AnyToString(data))
}
