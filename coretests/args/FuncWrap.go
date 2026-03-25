package args

import (
	"reflect"

	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

// FuncWrap wraps a Go function value with reflection metadata for
// dynamic invocation, argument validation, and introspection.
//
// Type parameter T represents the type of the wrapped function.
// Use FuncWrapAny (= FuncWrap[any]) for untyped usage with reflection-based creation.
//
// Example (typed):
//
//	fw := args.NewTypedFuncWrap(myFunc)
//	results, err := fw.Invoke("arg1", 42)
//
// Example (untyped, via creator):
//
//	fw := args.NewFuncWrap.Default(myFunc)
//	results, err := fw.Invoke("arg1", 42)
type FuncWrap[T any] struct {
	Name                 string         `json:",omitempty"`
	FullName             string         `json:",omitempty"`
	Func                 T              `json:"-"`
	isInvalid            bool           `json:"IsInvalid,omitempty"`
	rvType               reflect.Type   `json:"-"`
	rv                   reflect.Value  `json:"-"`
	inArgsTypesNames     []string       `json:"-"`
	inArgsTypes          []reflect.Type `json:"-"`
	outArgsTypes         []reflect.Type `json:"-"`
	pkgNameOnly          string
	funcDirectInvokeName string
	pkgPath              string
	inArgsMap            Map
	outArgsMap           Map
	inArgsNames          []string
	outArgsTypesNames    []string
	outArgsNames         []string
}

// NewTypedFuncWrap creates a type-safe FuncWrap[T] from a typed function.
// This is a package-level function because Go does not support generic methods
// on non-generic receiver types.
//
// Example:
//
//	fw := args.NewTypedFuncWrap(func(s string) int { return len(s) })
//	fmt.Println(fw.ArgsCount()) // 1
func NewTypedFuncWrap[T any](fn T) *FuncWrap[T] {
	anyFn := any(fn)

	if reflectinternal.Is.Null(anyFn) {
		return &FuncWrap[T]{
			Func:      fn,
			isInvalid: true,
		}
	}

	typeOf := reflect.TypeOf(anyFn)
	kind := typeOf.Kind()

	if kind != reflect.Func {
		return &FuncWrap[T]{
			Func:      fn,
			isInvalid: true,
			rvType:    typeOf,
		}
	}

	fullName, nameOnly := reflectinternal.GetFunc.FullNameWithName(anyFn)

	return &FuncWrap[T]{
		Name:      nameOnly,
		FullName:  fullName,
		Func:      fn,
		isInvalid: false,
		rvType:    typeOf,
		rv:        reflect.ValueOf(anyFn),
	}
}

// GetFuncName returns the short name of the wrapped function.
func (it *FuncWrap[T]) GetFuncName() string {
	if it == nil {
		return ""
	}

	return it.Name
}

// GetPascalCaseFuncName returns the function name in PascalCase format.
func (it *FuncWrap[T]) GetPascalCaseFuncName() string {
	if it == nil {
		return ""
	}

	return pascalCaseFunc(it.Name)
}

// HasValidFunc checks whether the FuncWrap holds a valid, callable function.
func (it *FuncWrap[T]) HasValidFunc() bool {
	return it != nil &&
		!it.isInvalid &&
		it.rv.IsValid() &&
		reflectinternal.Is.Func(it.Func)
}

// IsInvalid returns true if the FuncWrap is nil, marked invalid,
// or does not hold a valid function reference.
func (it *FuncWrap[T]) IsInvalid() bool {
	return it == nil ||
		it.isInvalid ||
		!it.rv.IsValid() ||
		!it.HasValidFunc()
}

// IsValid returns true if the FuncWrap holds a valid, callable function.
func (it *FuncWrap[T]) IsValid() bool {
	return !it.IsInvalid()
}

// PkgPath returns the full package path of the wrapped function.
func (it *FuncWrap[T]) PkgPath() string {
	if it.IsInvalid() {
		return ""
	}

	if len(it.pkgPath) > 0 {
		return it.pkgPath
	}

	it.pkgPath = reflectinternal.GetFunc.GetPkgPathFullName(it.FullName)

	return it.pkgPath
}

// PkgNameOnly returns only the package name (without path) of the wrapped function.
func (it *FuncWrap[T]) PkgNameOnly() string {
	if it.IsInvalid() {
		return ""
	}

	if len(it.pkgNameOnly) > 0 {
		return it.pkgNameOnly
	}

	it.pkgNameOnly = reflectinternal.Utils.PkgNameOnly(it.Func)

	return it.pkgNameOnly
}

// FuncDirectInvokeName returns the direct invocation name of the function,
// suitable for code generation.
func (it *FuncWrap[T]) FuncDirectInvokeName() string {
	if it.IsInvalid() {
		return ""
	}

	if len(it.funcDirectInvokeName) > 0 {
		return it.funcDirectInvokeName
	}

	it.funcDirectInvokeName = reflectinternal.GetFunc.FuncDirectInvokeNameUsingFullName(it.FullName)

	return it.funcDirectInvokeName
}

// GetType returns the reflect.Type of the wrapped function.
func (it *FuncWrap[T]) GetType() reflect.Type {
	if it.IsInvalid() {
		return nil
	}

	return it.rvType
}

// IsPublicMethod returns true if the wrapped function is an exported method.
func (it *FuncWrap[T]) IsPublicMethod() bool {
	return it != nil && it.rvType.PkgPath() == ""
}

// IsPrivateMethod returns true if the wrapped function is an unexported method.
func (it *FuncWrap[T]) IsPrivateMethod() bool {
	return it != nil && it.rvType.PkgPath() != ""
}

// IsNotEqual returns true if the two FuncWraps are not equal.
func (it *FuncWrap[T]) IsNotEqual(another *FuncWrap[T]) bool {
	return !it.IsEqual(another)
}

// IsEqualValue compares equality using a value (non-pointer) FuncWrap.
func (it *FuncWrap[T]) IsEqualValue(another FuncWrap[T]) bool {
	return it.IsEqual(&another)
}

// IsEqual performs a deep equality check between two FuncWraps,
// comparing validity, name, visibility, argument counts, and argument types.
func (it *FuncWrap[T]) IsEqual(another *FuncWrap[T]) bool {
	if it == nil && another == nil {
		return true
	}

	if it == nil || another == nil {
		return false
	}

	if it == another {
		return true
	}

	if it.IsInvalid() != another.IsInvalid() {
		return false
	}

	if it.Name != another.Name {
		return false
	}

	if it.IsPublicMethod() != another.IsPublicMethod() {
		return false
	}

	if it.ArgsCount() != another.ArgsCount() {
		return false
	}

	if it.ReturnLength() != another.ReturnLength() {
		return false
	}

	isInArgsOkay, _ := it.InArgsVerifyRv(another.GetInArgsTypes())

	if !isInArgsOkay {
		return false
	}

	isOutArgsOkay, _ := it.OutArgsVerifyRv(another.GetOutArgsTypes())

	if !isOutArgsOkay {
		return false
	}

	return true
}
