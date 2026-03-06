package conditional

// ============================================================================
// Typed convenience wrappers for generic conditional functions.
//
// These provide ergonomic, non-generic entry points for the most common types.
// They replace the deprecated per-type functions (Bool, Int, String, etc.)
// with consistently named alternatives.
//
// Naming convention:
//   If<Type>          — ternary value select
//   IfFunc<Type>      — ternary with lazy evaluation (both branches)
//   IfTrueFunc<Type>  — lazy evaluation of true branch only, zero on false
//   IfSlice<Type>     — ternary for slice types
//   IfPtr<Type>       — ternary for pointer types
// ============================================================================

// --- If<Type> wrappers (replace Bool, Int, String, Byte, Interface) ---

func IfBool(isTrue bool, trueValue, falseValue bool) bool {
	return If[bool](isTrue, trueValue, falseValue)
}

func IfInt(isTrue bool, trueValue, falseValue int) int {
	return If[int](isTrue, trueValue, falseValue)
}

func IfString(isTrue bool, trueValue, falseValue string) string {
	return If[string](isTrue, trueValue, falseValue)
}

func IfByte(isTrue bool, trueValue, falseValue byte) byte {
	return If[byte](isTrue, trueValue, falseValue)
}

func IfAny(isTrue bool, trueValue, falseValue any) any {
	return If[any](isTrue, trueValue, falseValue)
}

func IfFloat64(isTrue bool, trueValue, falseValue float64) float64 {
	return If[float64](isTrue, trueValue, falseValue)
}

// --- IfFunc<Type> wrappers (replace BoolFunc, StringFunc, InterfaceFunc) ---

func IfFuncBool(isTrue bool, trueValueFunc, falseValueFunc func() bool) bool {
	return IfFunc[bool](isTrue, trueValueFunc, falseValueFunc)
}

func IfFuncString(isTrue bool, trueValueFunc, falseValueFunc func() string) string {
	return IfFunc[string](isTrue, trueValueFunc, falseValueFunc)
}

func IfFuncAny(isTrue bool, trueValueFunc, falseValueFunc func() any) any {
	return IfFunc[any](isTrue, trueValueFunc, falseValueFunc)
}

func IfFuncInt(isTrue bool, trueValueFunc, falseValueFunc func() int) int {
	return IfFunc[int](isTrue, trueValueFunc, falseValueFunc)
}

// --- IfTrueFunc<Type> wrappers (replace BooleanTrueFunc, StringTrueFunc, BytesTrueFunc) ---

func IfTrueFuncBool(isTrue bool, trueValueFunc func() bool) bool {
	return IfTrueFunc[bool](isTrue, trueValueFunc)
}

func IfTrueFuncString(isTrue bool, trueValueFunc func() string) string {
	return IfTrueFunc[string](isTrue, trueValueFunc)
}

func IfTrueFuncStrings(isTrue bool, trueValueFunc func() []string) []string {
	return IfTrueFunc[[]string](isTrue, trueValueFunc)
}

func IfTrueFuncBytes(isTrue bool, trueValueFunc func() []byte) []byte {
	return IfTrueFunc[[]byte](isTrue, trueValueFunc)
}

// --- IfSlice<Type> wrappers (replace Booleans, Integers, Strings, Bytes, Interfaces) ---

func IfSliceBool(isTrue bool, trueValue, falseValue []bool) []bool {
	return IfSlice[bool](isTrue, trueValue, falseValue)
}

func IfSliceInt(isTrue bool, trueValue, falseValue []int) []int {
	return IfSlice[int](isTrue, trueValue, falseValue)
}

func IfSliceString(isTrue bool, trueValue, falseValue []string) []string {
	return IfSlice[string](isTrue, trueValue, falseValue)
}

func IfSliceByte(isTrue bool, trueValue, falseValue []byte) []byte {
	return IfSlice[byte](isTrue, trueValue, falseValue)
}

func IfSliceAny(isTrue bool, trueValue, falseValue []any) []any {
	return IfSlice[any](isTrue, trueValue, falseValue)
}

// --- IfPtr<Type> wrappers (replace StringPtr) ---

func IfPtrString(isTrue bool, trueValue, falseValue *string) *string {
	return IfPtr[string](isTrue, trueValue, falseValue)
}

func IfPtrInt(isTrue bool, trueValue, falseValue *int) *int {
	return IfPtr[int](isTrue, trueValue, falseValue)
}

func IfPtrBool(isTrue bool, trueValue, falseValue *bool) *bool {
	return IfPtr[bool](isTrue, trueValue, falseValue)
}
