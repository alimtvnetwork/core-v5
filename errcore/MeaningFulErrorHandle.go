package errcore

func MeaningfulErrorHandle(
	rawErrType RawErrorType,
	funcName string,
	err error,
) {
	if err == nil {
		return
	}

	meaningfulErr := MeaningfulError(rawErrType, funcName, err)

	panic(meaningfulErr)
}
