package trydo

func WrapPanic(voidFunc func()) Exception {
	var exception Exception

	Block{
		Try: func() {
			voidFunc()
		},
		Catch: func(e Exception) {
			exception = e
		},
		Finally: nil,
	}.Do()

	return exception
}

func ErrorFuncWrapPanic(errFunc func() error) WrappedErr {
	var exception Exception
	var err error
	var hasThrown bool

	catchFunc := func(e Exception) {
		exception = e
		hasThrown = true
	}

	defer func() {
		r := recover()

		catchFunc(r)
	}()

	err = errFunc()

	return WrappedErr{
		Error:     err,
		Exception: exception,
		HasThrown: hasThrown,
		HasError:  err != nil,
	}
}
