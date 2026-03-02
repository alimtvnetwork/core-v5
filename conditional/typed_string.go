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

// Deprecated: Use IfSliceString instead.
func IfSlicePtrString(
	isTrue bool,
	trueValue, falseValue []string,
) []string {
	return IfSlice[string](isTrue, trueValue, falseValue)
}

// Deprecated: Use IfSlice[string] with func wrappers instead.
func IfSlicePtrFuncString(
	isTrue bool,
	trueValueFunc, falseValueFunc func() []string,
) []string {
	if isTrue {
		return trueValueFunc()
	}
	return falseValueFunc()
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

// NilDerefString is a typed convenience wrapper for NilDeref[string].
func NilDerefString(valuePointer *string) string {
	return NilDeref[string](valuePointer)
}

// NilDerefPtrString is a typed convenience wrapper for NilDerefPtr[string].
func NilDerefPtrString(valuePointer *string) *string {
	return NilDerefPtr[string](valuePointer)
}

// NilValString is a typed convenience wrapper for NilVal[string].
func NilValString(valuePointer *string, onNil, onNonNil string) string {
	return NilVal[string](valuePointer, onNil, onNonNil)
}

// NilValPtrString is a typed convenience wrapper for NilValPtr[string].
func NilValPtrString(valuePointer *string, onNil, onNonNil string) *string {
	return NilValPtr[string](valuePointer, onNil, onNonNil)
}
