package reflectinternal

import "reflect"

type isChecker struct{}

func IsConclusive(left, right interface{}) (isEqual, isConclusive bool) {
	if left == right {
		return true, true
	}

	if left == nil && right == nil {
		return true, true
	}

	if left == nil || right == nil {
		return false, true
	}

	leftRv := reflect.ValueOf(left)
	rightRv := reflect.ValueOf(right)
	isLeftNull := IsNullUsingReflectValue(leftRv)
	isRightNull := IsNullUsingReflectValue(rightRv)
	isBothEqual := isLeftNull == isRightNull

	if isLeftNull && isBothEqual {
		// both null
		return true, true
	} else if !isBothEqual && isLeftNull || isRightNull {
		// any null but the other is not
		return false, true
	}

	if leftRv.Type() != rightRv.Type() {
		return false, true
	}

	return false, false
}

func IsAnyEqual(left, right interface{}) bool {
	isEqual, isConclusive := IsConclusive(left, right)

	if isConclusive {
		return isEqual
	}

	return reflect.DeepEqual(left, right)
}

func IsFunc(item interface{}) bool {
	if item == nil {
		return true
	}

	typeOf := reflect.TypeOf(item)

	return IsFuncTypeOf(typeOf)
}

func IsFuncTypeOf(typeOf reflect.Type) bool {
	kind := typeOf.Kind()

	switch kind {
	case reflect.Func:
		return true
	}

	return false
}

func IsNotNull(item interface{}) bool {
	return !IsNull(item)
}
func IsNull(item interface{}) bool {
	if item == nil {
		return true
	}

	rv := reflect.ValueOf(item)

	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Slice:
		return rv.IsNil()
	default:
		return false
	}
}

func IsNullUsingReflectValue(rv reflect.Value) bool {
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Slice:
		return rv.IsNil()
	default:
		return false
	}
}

// IsNumber function returns true if the kind passed to it is one of the
// primitive types (reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
//
//	reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
//	reflect.Float32, reflect.Float64)
func IsNumber(kind reflect.Kind) bool {
	switch kind {
	case
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

// IsPrimitive function returns true if the kind passed to it is one of the
// primitive types (boolean, int, uint, float, string)
func IsPrimitive(kind reflect.Kind) bool {
	switch kind {
	case
		reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.String:
		return true
	default:
		return false
	}
}

// IsZero
//
//	returns true if the current value is null
//	or reflect value is zero
//
// Reference:
//   - Stackoverflow Example : https://stackoverflow.com/a/23555352
func IsZero(anyItem interface{}) bool {
	if IsNull(anyItem) {
		return true
	}

	return IsZeroReflectValue(reflect.ValueOf(anyItem))
}

// IsZeroReflectValue
//
//	returns true if the current value is null
//	or reflect value is zero
//
// Reference:
//   - Stackoverflow Example : https://stackoverflow.com/a/23555352
func IsZeroReflectValue(rv reflect.Value) bool {
	switch rv.Kind() {
	case reflect.Func, reflect.Map, reflect.Slice, reflect.Ptr:
		return rv.IsNil()
	case reflect.Array:
		isAllZero := true
		for i := 0; i < rv.Len(); i++ {
			isAllZero = isAllZero && IsZeroReflectValue(rv.Index(i))
		}

		return isAllZero
	case reflect.Struct:
		isAllZero := true
		for i := 0; i < rv.NumField(); i++ {
			isAllZero = isAllZero && IsZeroReflectValue(rv.Field(i))
		}

		return isAllZero
	}

	// Compare other types directly:
	z := reflect.Zero(rv.Type())

	return rv.Interface() == z.Interface()
}
