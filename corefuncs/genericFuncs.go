package corefuncs

// Generic function types — type-safe versions of the any-based func types.
// These allow callers to work with concrete types instead of any,
// eliminating the need for type assertions at call sites.
//
// Each generic type mirrors a corresponding any-based type in funcs.go:
//
//	InOutFunc           → InOutFuncOf[TIn, TOut]
//	InOutErrFunc        → InOutErrFuncOf[TIn, TOut]
//	SerializeOutputFunc → SerializeOutputFuncOf[TIn]
//	InActionReturnsErrFunc → InActionReturnsErrFuncOf[TIn]
//	ResultDelegatingFunc   → ResultDelegatingFuncOf[T]
type (
	// InOutFuncOf is a generic version of InOutFunc.
	//
	//	any-based:  func(input any) (output any)
	//	generic:    func(input TIn) (output TOut)
	InOutFuncOf[TIn any, TOut any] func(input TIn) (output TOut)

	// InOutErrFuncOf is a generic version of InOutErrFunc.
	//
	//	any-based:  func(input any) (output any, err error)
	//	generic:    func(input TIn) (output TOut, err error)
	InOutErrFuncOf[TIn any, TOut any] func(input TIn) (output TOut, err error)

	// SerializeOutputFuncOf is a generic version of SerializeOutputFunc.
	//
	//	any-based:  func(input any) (serializedBytes []byte, err error)
	//	generic:    func(input TIn) (serializedBytes []byte, err error)
	SerializeOutputFuncOf[TIn any] func(input TIn) (serializedBytes []byte, err error)

	// InActionReturnsErrFuncOf is a generic version of InActionReturnsErrFunc.
	//
	//	any-based:  func(input any) (err error)
	//	generic:    func(input TIn) (err error)
	InActionReturnsErrFuncOf[TIn any] func(input TIn) (err error)

	// ResultDelegatingFuncOf is a generic version of ResultDelegatingFunc.
	//
	//	any-based:  func(resultDelegatedTo any) error
	//	generic:    func(resultDelegatedTo T) error
	//
	// resultDelegatedTo can be unmarshal or marshal or reflect set target.
	ResultDelegatingFuncOf[T any] func(resultDelegatedTo T) error
)
