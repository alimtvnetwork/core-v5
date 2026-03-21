package argstests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Map: GetFuncName returns empty when no FuncWrap ──
// Covers Map.go L107

func Test_Cov9_Map_GetFuncName_Empty(t *testing.T) {
	m := args.Map{"key": "value"}

	result := m.GetFuncName()

	actual := args.Map{"empty": result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "GetFuncName returns empty -- no FuncWrap in map", actual)
}

// ── Map: InvokeMust panics on invalid func ──
// Covers Map.go L422-423

func Test_Cov9_Map_InvokeMust_Panic(t *testing.T) {
	m := args.Map{"workFunc": "not-a-func"}

	didPanic := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		m.InvokeMust()
	}()

	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "InvokeMust panics -- invalid work func", actual)
}

// ── Map: InvokeWithValidArgs ──
// Covers Map.go L434-439

func Test_Cov9_Map_InvokeWithValidArgs(t *testing.T) {
	fn := func(s string) string { return s + "!" }
	m := args.Map{
		"workFunc": fn,
		"first":    "hello",
	}

	results, err := m.InvokeWithValidArgs()

	actual := args.Map{
		"noErr":     err == nil,
		"hasResult": len(results) > 0,
	}
	expected := args.Map{
		"noErr":     true,
		"hasResult": true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeWithValidArgs invokes function -- valid args", actual)
}

// ── FuncWrap: IsEqual various comparison paths ──
// Covers FuncWrap.go L215-245

func Test_Cov9_FuncWrap_IsEqual_Different(t *testing.T) {
	fn1 := func(s string) string { return s }
	fn2 := func(s string, i int) string { return s }

	fw1 := args.NewFuncWrap.Default(fn1)
	fw2 := args.NewFuncWrap.Default(fn2)

	result := fw1.IsEqual(fw2)

	actual := args.Map{"equal": result}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns false -- different arg counts", actual)
}

func Test_Cov9_FuncWrap_IsEqual_Nil(t *testing.T) {
	fn1 := func() {}
	fw1 := args.NewFuncWrap.Default(fn1)

	result := fw1.IsEqual(nil)

	actual := args.Map{"equal": result}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns false -- nil other", actual)
}

// ── FuncWrap: InvokeFirstAndError ──
// Covers FuncWrapInvoke.go L121-137

func Test_Cov9_FuncWrap_InvokeFirstAndError(t *testing.T) {
	fn := func(s string) (string, error) { return s + "!", nil }
	fw := args.NewFuncWrap.Default(fn)

	first, funcErr, procErr := fw.InvokeFirstAndError("hello")

	actual := args.Map{
		"first":     first,
		"noFuncErr": funcErr == nil,
		"noProcErr": procErr == nil,
	}
	expected := args.Map{
		"first":     "hello!",
		"noFuncErr": true,
		"noProcErr": true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeFirstAndError returns first and nil errors", actual)
}

// ── FuncWrap: InvokeError ──
// Covers FuncWrapInvoke.go L107-114

func Test_Cov9_FuncWrap_InvokeError(t *testing.T) {
	fn := func() error { return nil }
	fw := args.NewFuncWrap.Default(fn)

	funcErr, procErr := fw.InvokeError()

	actual := args.Map{
		"noFuncErr": funcErr == nil,
		"noProcErr": procErr == nil,
	}
	expected := args.Map{
		"noFuncErr": true,
		"noProcErr": true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeError returns nil errors -- no-error func", actual)
}

// ── FuncWrap: GetResponseOfInvoke with index ──
// Covers FuncWrapInvoke.go L85-101

func Test_Cov9_FuncWrap_GetResponseOfInvoke(t *testing.T) {
	fn := func() (string, int) { return "a", 42 }
	fw := args.NewFuncWrap.Default(fn)

	result, err := fw.GetResponseOfInvoke(1)

	actual := args.Map{"result": result, "noErr": err == nil}
	expected := args.Map{"result": 42, "noErr": true}
	expected.ShouldBeEqual(t, 0, "GetResponseOfInvoke returns indexed result", actual)
}

// ── FuncWrap: InvokeSkip with panic ──
// Covers FuncWrapInvoke.go L61-71

func Test_Cov9_FuncWrap_InvokeSkip_Panic(t *testing.T) {
	fn := func() { panic("test panic") }
	fw := args.NewFuncWrap.Default(fn)

	_, err := fw.Invoke()

	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Invoke returns error -- function panics", actual)
}

// ── funcDetector: GetFuncWrap with FuncWrapAny directly ──
// Covers funcDetector.go L16-17

func Test_Cov9_FuncDetector_DirectFuncWrap(t *testing.T) {
	fn := func() {}
	fw := args.NewFuncWrap.Default(fn)

	result := args.FuncDetector.GetFuncWrap(fw)

	actual := args.Map{"notNil": result != nil, "valid": result.IsValid()}
	expected := args.Map{"notNil": true, "valid": true}
	expected.ShouldBeEqual(t, 0, "GetFuncWrap returns same instance -- direct FuncWrap", actual)
}

// ── funcDetector: GetFuncWrap with raw function ──
// Covers funcDetector.go L18-22 (default)

func Test_Cov9_FuncDetector_RawFunc(t *testing.T) {
	fn := func(s string) string { return s }

	result := args.FuncDetector.GetFuncWrap(fn)

	actual := args.Map{"notNil": result != nil, "valid": result.IsValid()}
	expected := args.Map{"notNil": true, "valid": true}
	expected.ShouldBeEqual(t, 0, "GetFuncWrap creates new FuncWrap -- raw function", actual)
}

// ── newFuncWrapCreator: MethodToFunc nil method ──
// Covers newFuncWrapCreator.go L101-106

func Test_Cov9_NewFuncWrap_MethodToFunc_Nil(t *testing.T) {
	fw, err := args.NewFuncWrap.MethodToFunc(nil)

	actual := args.Map{"invalid": fw.IsInvalid(), "hasErr": err != nil}
	expected := args.Map{"invalid": true, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "MethodToFunc returns invalid -- nil method", actual)
}

// ── newFuncWrapCreator: StructToMap ──
// Covers newFuncWrapCreator.go L122-141

type cov9TestStruct struct{}

func (s *cov9TestStruct) Hello() string { return "hello" }
func (s *cov9TestStruct) World() string { return "world" }

func Test_Cov9_NewFuncWrap_StructToMap(t *testing.T) {
	s := &cov9TestStruct{}
	fm, err := args.NewFuncWrap.StructToMap(s)

	actual := args.Map{"noErr": err == nil, "hasEntries": len(fm) > 0}
	expected := args.Map{"noErr": true, "hasEntries": true}
	expected.ShouldBeEqual(t, 0, "StructToMap creates FuncMap -- valid struct", actual)
}

// ── OneFunc: InvokeMust / InvokeWithValidArgs / InvokeArgs ──
// Covers OneFunc.go L84-105

func Test_Cov9_OneFunc_InvokeMust(t *testing.T) {
	of := &args.OneFunc[string]{
		First:    "hi",
		WorkFunc: func(s string) string { return s + "!" },
	}

	results := of.InvokeMust("test")

	actual := args.Map{"hasResult": len(results) > 0}
	expected := args.Map{"hasResult": true}
	expected.ShouldBeEqual(t, 0, "OneFunc InvokeMust returns result -- valid call", actual)
}

func Test_Cov9_OneFunc_InvokeWithValidArgs(t *testing.T) {
	of := &args.OneFunc[string]{
		First:    "hi",
		WorkFunc: func(s string) string { return s + "!" },
	}

	results, err := of.InvokeWithValidArgs()

	actual := args.Map{"noErr": err == nil, "hasResult": len(results) > 0}
	expected := args.Map{"noErr": true, "hasResult": true}
	expected.ShouldBeEqual(t, 0, "OneFunc InvokeWithValidArgs returns result", actual)
}

func Test_Cov9_OneFunc_InvokeArgs(t *testing.T) {
	of := &args.OneFunc[string]{
		First:    "hi",
		WorkFunc: func(s string) string { return s + "!" },
	}

	results, err := of.InvokeArgs(1)

	actual := args.Map{"noErr": err == nil, "hasResult": len(results) > 0}
	expected := args.Map{"noErr": true, "hasResult": true}
	expected.ShouldBeEqual(t, 0, "OneFunc InvokeArgs returns result", actual)
}

// ── Holder: InvokeMust / InvokeWithValidArgs / InvokeArgs ──
// Covers Holder.go L188-204

func Test_Cov9_Holder_InvokeMust(t *testing.T) {
	fn := func() string { return "test" }
	h := &args.Holder[string]{}
	

	results := h.InvokeMust()

	actual := args.Map{"hasResult": len(results) > 0}
	expected := args.Map{"hasResult": true}
	expected.ShouldBeEqual(t, 0, "Holder InvokeMust returns result", actual)
}

func Test_Cov9_Holder_InvokeWithValidArgs(t *testing.T) {
	fn := func() string { return "test" }
	h := &args.Holder[string]{}
	

	results, err := h.InvokeWithValidArgs()

	actual := args.Map{"noErr": err == nil, "hasResult": len(results) > 0}
	expected := args.Map{"noErr": true, "hasResult": true}
	expected.ShouldBeEqual(t, 0, "Holder InvokeWithValidArgs returns result", actual)
}

func Test_Cov9_Holder_InvokeArgs(t *testing.T) {
	fn := func() string { return "test" }
	h := &args.Holder[string]{}
	

	results, err := h.InvokeArgs(0)

	actual := args.Map{"noErr": err == nil, "hasResult": len(results) > 0}
	expected := args.Map{"noErr": true, "hasResult": true}
	expected.ShouldBeEqual(t, 0, "Holder InvokeArgs returns result", actual)
}

// ── Dynamic: Invoke / InvokeMust / InvokeWithValidArgs ──
// Covers Dynamic.go L51,76-94

func Test_Cov9_Dynamic_Invoke(t *testing.T) {
	fn := func() string { return "dynamic" }
	d := &args.Dynamic[string]{
		Params: args.Map{"workFunc": fn},
	}

	results, err := d.Invoke()

	actual := args.Map{"noErr": err == nil, "hasResult": len(results) > 0}
	expected := args.Map{"noErr": true, "hasResult": true}
	expected.ShouldBeEqual(t, 0, "Dynamic Invoke returns result", actual)
}

func Test_Cov9_Dynamic_InvokeMust(t *testing.T) {
	fn := func() string { return "dynamic" }
	d := &args.Dynamic[string]{
		Params: args.Map{"workFunc": fn},
	}

	results := d.InvokeMust()

	actual := args.Map{"hasResult": len(results) > 0}
	expected := args.Map{"hasResult": true}
	expected.ShouldBeEqual(t, 0, "Dynamic InvokeMust returns result", actual)
}

func Test_Cov9_Dynamic_InvokeWithValidArgs(t *testing.T) {
	fn := func() string { return "dynamic" }
	d := &args.Dynamic[string]{
		Params: args.Map{"workFunc": fn},
	}

	results, err := d.InvokeWithValidArgs()

	actual := args.Map{"noErr": err == nil, "hasResult": len(results) > 0}
	expected := args.Map{"noErr": true, "hasResult": true}
	expected.ShouldBeEqual(t, 0, "Dynamic InvokeWithValidArgs returns result", actual)
}
