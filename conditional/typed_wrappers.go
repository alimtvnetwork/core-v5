package conditional

// ============================================================================
// Additional typed convenience wrappers not covered by the per-type files.
//
// These provide ergonomic, non-generic entry points for composite types
// and cross-type wrappers used in typed_wrappers.go historically.
//
// All primary typed wrappers are now in their respective typed_*.go files.
// ============================================================================

// --- IfTrueFunc composite wrappers ---

func IfTrueFuncStrings(isTrue bool, trueValueFunc func() []string) []string {
	return IfTrueFunc[[]string](isTrue, trueValueFunc)
}

func IfTrueFuncBytes(isTrue bool, trueValueFunc func() []byte) []byte {
	return IfTrueFunc[[]byte](isTrue, trueValueFunc)
}

// --- IfSlice for any ---

func IfSliceAny(isTrue bool, trueValue, falseValue []any) []any {
	return IfSlice[any](isTrue, trueValue, falseValue)
}

// --- IfAny ---

func IfAny(isTrue bool, trueValue, falseValue any) any {
	return If[any](isTrue, trueValue, falseValue)
}

// --- IfFunc for any ---

func IfFuncAny(isTrue bool, trueValueFunc, falseValueFunc func() any) any {
	return IfFunc[any](isTrue, trueValueFunc, falseValueFunc)
}
