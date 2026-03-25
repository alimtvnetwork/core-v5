package reflectinternal

import (
	"testing"
)

func TestGetFunc_All(t *testing.T) {
	full, pkg, method := GetFunc.All("github.com/alimtvnetwork/core/codestack.TestFunc")
	if full == "" || pkg == "" || method == "" {
		t.Fatal("expected non-empty")
	}

	// empty
	f2, p2, m2 := GetFunc.All("")
	if f2 != "" || p2 != "" || m2 != "" {
		t.Fatal("expected all empty")
	}

	// simple name without forward slash
	f3, p3, m3 := GetFunc.All("pkg.Method")
	if f3 == "" || p3 == "" || m3 == "" {
		t.Fatal("expected non-empty for simple")
	}
}

func TestGetFunc_RunTime(t *testing.T) {
	fn := func() {}
	rt := GetFunc.RunTime(fn)
	if rt == nil {
		t.Fatal("expected non-nil")
	}

	rt2 := GetFunc.RunTime(nil)
	if rt2 != nil {
		t.Fatal("expected nil")
	}

	rt3 := GetFunc.RunTime(42)
	if rt3 != nil {
		t.Fatal("expected nil for non-func")
	}
}

func TestGetFunc_FullName(t *testing.T) {
	fn := func() {}
	name := GetFunc.FullName(fn)
	if name == "" {
		t.Fatal("expected non-empty")
	}

	name2 := GetFunc.FullName(nil)
	if name2 != "" {
		t.Fatal("expected empty for nil")
	}
}

func TestGetFunc_FullNameWithName(t *testing.T) {
	fn := func() {}
	full, name := GetFunc.FullNameWithName(fn)
	if full == "" || name == "" {
		t.Fatal("expected non-empty")
	}

	full2, name2 := GetFunc.FullNameWithName(nil)
	if full2 != "" || name2 != "" {
		t.Fatal("expected empty for nil")
	}
}

func TestGetFunc_NameOnly(t *testing.T) {
	fn := func() {}
	name := GetFunc.NameOnly(fn)
	if name == "" {
		t.Fatal("expected non-empty")
	}

	name2 := GetFunc.NameOnly(nil)
	if name2 != "" {
		t.Fatal("expected empty for nil")
	}
}

func TestGetFunc_NameOnlyByStack(t *testing.T) {
	name := GetFunc.NameOnlyByStack(0)
	if name == "" {
		t.Fatal("expected non-empty")
	}
}

func TestGetFunc_FuncDirectInvokeName(t *testing.T) {
	fn := func() {}
	name := GetFunc.FuncDirectInvokeName(fn)
	if name == "" {
		t.Fatal("expected non-empty")
	}
}

func TestGetFunc_FuncDirectInvokeNameUsingFullName(t *testing.T) {
	name := GetFunc.FuncDirectInvokeNameUsingFullName(
		"github.com/alimtvnetwork/core/codestack.TestFunc",
	)
	if name == "" {
		t.Fatal("expected non-empty")
	}

	name2 := GetFunc.FuncDirectInvokeNameUsingFullName("")
	if name2 != "" {
		t.Fatal("expected empty")
	}

	// no forward slash
	name3 := GetFunc.FuncDirectInvokeNameUsingFullName("pkg.Method")
	if name3 != "pkg.Method" {
		t.Fatal("expected same", name3)
	}
}

func TestGetFunc_PascalFuncName(t *testing.T) {
	s := GetFunc.PascalFuncName("hello")
	if s != "Hello" {
		t.Fatal("expected Hello, got:", s)
	}

	s2 := GetFunc.PascalFuncName("")
	if s2 != "" {
		t.Fatal("expected empty")
	}

	s3 := GetFunc.PascalFuncName("a")
	if s3 != "A" {
		t.Fatal("expected A")
	}
}

func TestGetFunc_GetMethod(t *testing.T) {
	type testStruct struct{}
	m := GetFunc.GetMethod("", testStruct{})
	if m != nil {
		t.Fatal("expected nil")
	}

	m2 := GetFunc.GetMethod("NonExistent", testStruct{})
	if m2 != nil {
		t.Fatal("expected nil")
	}

	m3 := GetFunc.GetMethod("test", nil)
	if m3 != nil {
		t.Fatal("expected nil for nil")
	}
}

func TestGetFunc_GetMethods(t *testing.T) {
	type testStruct struct{ Name string }
	methods := GetFunc.GetMethods(testStruct{})
	_ = methods

	methods2 := GetFunc.GetMethods(nil)
	if len(methods2) != 0 {
		t.Fatal("expected empty")
	}
}

func TestGetFunc_GetMethodsNames(t *testing.T) {
	type testStruct struct{ Name string }
	names := GetFunc.GetMethodsNames(testStruct{})
	_ = names

	names2 := GetFunc.GetMethodsNames(nil)
	if len(names2) != 0 {
		t.Fatal("expected empty")
	}
}

func TestGetFunc_GetMethodsMap(t *testing.T) {
	type testStruct struct{ Name string }
	m := GetFunc.GetMethodsMap(testStruct{})
	_ = m

	m2 := GetFunc.GetMethodsMap(nil)
	if len(m2) != 0 {
		t.Fatal("expected empty")
	}
}

func TestGetFunc_GetPkgPath(t *testing.T) {
	fn := func() {}
	p := GetFunc.GetPkgPath(fn)
	_ = p
}

func TestGetFunc_GetPkgPathFullName(t *testing.T) {
	p := GetFunc.GetPkgPathFullName("github.com/alimtvnetwork/core/codestack.TestFunc")
	if p == "" {
		t.Fatal("expected non-empty")
	}

	p2 := GetFunc.GetPkgPathFullName("")
	if p2 != "" {
		t.Fatal("expected empty")
	}

	p3 := GetFunc.GetPkgPathFullName("simple")
	if p3 != "simple" {
		t.Fatal("expected same")
	}
}

func TestGetFunc_fixFinalFuncName(t *testing.T) {
	s := GetFunc.fixFinalFuncName("TestFunc-fm")
	if s != "TestFunc" {
		t.Fatal("expected TestFunc, got:", s)
	}

	s2 := GetFunc.fixFinalFuncName("TestFunc")
	if s2 != "TestFunc" {
		t.Fatal("expected TestFunc")
	}
}

func TestGetFunc_firstLastDefault(t *testing.T) {
	f, l := GetFunc.firstLastDefault([]string{})
	if f != "" || l != "" {
		t.Fatal("expected empty")
	}

	f2, l2 := GetFunc.firstLastDefault([]string{"a"})
	if f2 != "a" || l2 != "" {
		t.Fatal("unexpected")
	}

	f3, l3 := GetFunc.firstLastDefault([]string{"a", "b"})
	if f3 != "a" || l3 != "b" {
		t.Fatal("unexpected")
	}
}
