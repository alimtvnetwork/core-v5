package coredynamictests

import (
	"fmt"
	"reflect"
	"runtime"

	"gitlab.com/auk-go/core/coredata/coredynamic"
)

// =============================================================================
// DynamicBoolMethodRef
//
// Wraps a pointer-receiver bool method on *Dynamic.
// Storing the actual method reference ensures a build error if the method
// is renamed on the Dynamic struct.
// =============================================================================

type DynamicBoolMethodRef struct {
	fn func(*coredynamic.Dynamic) bool
}

func NewBoolRef(fn func(*coredynamic.Dynamic) bool) DynamicBoolMethodRef {
	return DynamicBoolMethodRef{fn: fn}
}

func (r DynamicBoolMethodRef) Call(d *coredynamic.Dynamic) bool {
	return r.fn(d)
}

// Name returns the short method name extracted via runtime reflection.
func (r DynamicBoolMethodRef) Name() string {
	fullName := runtime.FuncForPC(
		reflect.ValueOf(r.fn).Pointer(),
	).Name()

	return extractShortName(fullName)
}

func (r DynamicBoolMethodRef) String() string {
	return fmt.Sprintf("BoolMethodRef(%s)", r.Name())
}

// =============================================================================
// DynamicStringMethodRef
//
// Wraps a pointer-receiver string method on *Dynamic.
// =============================================================================

type DynamicStringMethodRef struct {
	fn func(*coredynamic.Dynamic) string
}

func NewStringRef(fn func(*coredynamic.Dynamic) string) DynamicStringMethodRef {
	return DynamicStringMethodRef{fn: fn}
}

func (r DynamicStringMethodRef) Call(d *coredynamic.Dynamic) string {
	return r.fn(d)
}

func (r DynamicStringMethodRef) Name() string {
	fullName := runtime.FuncForPC(
		reflect.ValueOf(r.fn).Pointer(),
	).Name()

	return extractShortName(fullName)
}

// =============================================================================
// Bool Method References
//
// Each variable below holds a compile-time reference to a *Dynamic method.
// If any method is renamed, these lines produce a build error.
// =============================================================================

var (
	refIsNull              = NewBoolRef((*coredynamic.Dynamic).IsNull)
	refIsValid             = NewBoolRef((*coredynamic.Dynamic).IsValid)
	refIsInvalid           = NewBoolRef((*coredynamic.Dynamic).IsInvalid)
	refIsStringType        = NewBoolRef((*coredynamic.Dynamic).IsStringType)
	refIsNumber            = NewBoolRef((*coredynamic.Dynamic).IsNumber)
	refIsPrimitive         = NewBoolRef((*coredynamic.Dynamic).IsPrimitive)
	refIsFunc              = NewBoolRef((*coredynamic.Dynamic).IsFunc)
	refIsSliceOrArray      = NewBoolRef((*coredynamic.Dynamic).IsSliceOrArray)
	refIsSliceOrArrayOrMap = NewBoolRef((*coredynamic.Dynamic).IsSliceOrArrayOrMap)
	refIsMap               = NewBoolRef((*coredynamic.Dynamic).IsMap)
	refIsPointer           = NewBoolRef((*coredynamic.Dynamic).IsPointer)
	refIsValueType         = NewBoolRef((*coredynamic.Dynamic).IsValueType)
	refIsStruct            = NewBoolRef((*coredynamic.Dynamic).IsStruct)
	refIsStructStringNullOrEmpty = NewBoolRef(
		(*coredynamic.Dynamic).IsStructStringNullOrEmpty,
	)
)

// =============================================================================
// String Method References
// =============================================================================

var (
	refString         = NewStringRef((*coredynamic.Dynamic).String)
	refValueString    = NewStringRef((*coredynamic.Dynamic).ValueString)
	refReflectTypeName = NewStringRef((*coredynamic.Dynamic).ReflectTypeName)
)

// =============================================================================
// Constructor References
//
// Package-level function references — build error if renamed.
// =============================================================================

var (
	refNewDynamicValid  = coredynamic.NewDynamicValid
	refNewDynamic       = coredynamic.NewDynamic
	refInvalidDynamic   = coredynamic.InvalidDynamic
	refInvalidDynamicPtr = coredynamic.InvalidDynamicPtr
	refNewDynamicPtr    = coredynamic.NewDynamicPtr
)

// =============================================================================
// Helpers
// =============================================================================

// extractShortName extracts the method name from a fully qualified
// runtime function name like "gitlab.com/auk-go/core/coredata/coredynamic.(*Dynamic).IsNull-fm".
func extractShortName(fullName string) string {
	// Find last dot for the method name
	for i := len(fullName) - 1; i >= 0; i-- {
		if fullName[i] == '.' {
			name := fullName[i+1:]

			// Strip "-fm" suffix added by Go runtime for method values
			if len(name) > 3 && name[len(name)-3:] == "-fm" {
				name = name[:len(name)-3]
			}

			return name
		}
	}

	return fullName
}

// init validates all method references are resolvable at test startup.
// This provides an early, clear failure message if any method is missing.
func init() {
	dynamicPtrType := reflect.TypeOf((*coredynamic.Dynamic)(nil))

	requiredMethods := []string{
		refIsNull.Name(),
		refIsValid.Name(),
		refIsInvalid.Name(),
		refIsStringType.Name(),
		refIsNumber.Name(),
		refIsPrimitive.Name(),
		refIsFunc.Name(),
		refIsSliceOrArray.Name(),
		refIsSliceOrArrayOrMap.Name(),
		refIsMap.Name(),
		refIsPointer.Name(),
		refIsValueType.Name(),
		refIsStruct.Name(),
		refIsStructStringNullOrEmpty.Name(),
		refString.Name(),
		refValueString.Name(),
		refReflectTypeName.Name(),
	}

	for _, name := range requiredMethods {
		if _, ok := dynamicPtrType.MethodByName(name); !ok {
			panic(fmt.Sprintf(
				"Dynamic method reference validation failed: method %q not found on *Dynamic",
				name,
			))
		}
	}
}
