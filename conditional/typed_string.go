package conditional

// IfString is a typed convenience wrapper for If[string].
func IfString(
	isTrue bool,
	trueValue, falseValue string,
) string {
	return If[string](isTrue, trueValue, falseValue)
}

// IfFuncString is a typed convenience wrapper for IfFunc[string].
func IfFuncString(
	isTrue bool,
	trueValueFunc, falseValueFunc func() string,
) string {
	return IfFunc[string](isTrue, trueValueFunc, falseValueFunc)
}

// IfTrueFuncString is a typed convenience wrapper for IfTrueFunc[string].
func IfTrueFuncString(
	isTrue bool,
	trueValueFunc func() string,
) string {
	return IfTrueFunc[string](isTrue, trueValueFunc)
}

// IfSliceString is a typed convenience wrapper for IfSlice[string].
func IfSliceString(
	isTrue bool,
	trueValue, falseValue []string,
) []string {
	return IfSlice[string](isTrue, trueValue, falseValue)
}

// IfSlicePtrString is a typed convenience wrapper for IfSlicePtr[string].
func IfSlicePtrString(
	isTrue bool,
	trueValue, falseValue *[]string,
) *[]string {
	return IfSlicePtr[string](isTrue, trueValue, falseValue)
}

// IfSlicePtrFuncString is a typed convenience wrapper for IfSlicePtrFunc[string].
func IfSlicePtrFuncString(
	isTrue bool,
	trueValueFunc, falseValueFunc func() *[]string,
) *[]string {
	return IfSlicePtrFunc[string](isTrue, trueValueFunc, falseValueFunc)
}

// IfPtrString is a typed convenience wrapper for IfPtr[string].
func IfPtrString(
	isTrue bool,
	trueValue, falseValue *string,
) *string {
	return IfPtr[string](isTrue, trueValue, falseValue)
}

// NilDefString is a typed convenience wrapper for NilDef[string].
// Note: This differs from the deprecated NilDefStr(ptr) which uses EmptyString as default.
// This version requires an explicit default value.
func NilDefString(
	valuePointer *string,
	defVal string,
) string {
	return NilDef[string](valuePointer, defVal)
}

// NilDefPtrString is a typed convenience wrapper for NilDefPtr[string].
func NilDefPtrString(
	valuePointer *string,
	defVal string,
) *string {
	return NilDefPtr[string](valuePointer, defVal)
}
