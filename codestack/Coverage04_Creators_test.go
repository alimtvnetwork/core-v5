package codestack

import (
	"testing"
)

func TestNameOf_Method(t *testing.T) {
	m := NameOf.Method()
	if m == "" {
		t.Fatal("expected method name")
	}
}

func TestNameOf_Package(t *testing.T) {
	p := NameOf.Package()
	if p == "" {
		t.Fatal("expected package name")
	}
}

func TestNameOf_All(t *testing.T) {
	full, pkg, method := NameOf.All("github.com/alimtvnetwork/core/codestack.TestNameOf_All")
	if full == "" || pkg == "" || method == "" {
		t.Fatal("expected non-empty")
	}

	// empty
	f2, p2, m2 := NameOf.All("")
	if f2 != "" || p2 != "" || m2 != "" {
		t.Fatal("expected all empty")
	}
}

func TestNameOf_AllStackSkip(t *testing.T) {
	full, pkg, method := NameOf.AllStackSkip(0)
	if full == "" || pkg == "" || method == "" {
		t.Fatal("expected non-empty")
	}
}

func TestNameOf_MethodStackSkip(t *testing.T) {
	m := NameOf.MethodStackSkip(0)
	if m == "" {
		t.Fatal("expected method name")
	}
}

func TestNameOf_PackageStackSkip(t *testing.T) {
	p := NameOf.PackageStackSkip(0)
	if p == "" {
		t.Fatal("expected package name")
	}
}

func TestNameOf_MethodByFullName(t *testing.T) {
	m := NameOf.MethodByFullName("github.com/alimtvnetwork/core/codestack.TestMethod")
	_ = m
}

func TestNameOf_PackageByFullName(t *testing.T) {
	p := NameOf.PackageByFullName("github.com/alimtvnetwork/core/codestack.TestMethod")
	_ = p
}

func TestNameOf_CurrentFuncFullPath(t *testing.T) {
	s := NameOf.CurrentFuncFullPath("github.com/alimtvnetwork/core/codestack.TestMethod")
	_ = s
}

func TestNameOf_JoinPackageNameWithRelative(t *testing.T) {
	s := NameOf.JoinPackageNameWithRelative(
		"github.com/alimtvnetwork/core/codestack.TestMethod",
		"NewFunc",
	)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestStacksCreator_All(t *testing.T) {
	tc := New.StackTrace.All(true, true, 0, 5)
	_ = tc
}

func TestStacksCreator_Default(t *testing.T) {
	tc := New.StackTrace.Default(0, 5)
	_ = tc
}

func TestStacksCreator_DefaultCount(t *testing.T) {
	tc := New.StackTrace.DefaultCount(0)
	_ = tc
}

func TestStacksCreator_SkipOne(t *testing.T) {
	tc := New.StackTrace.SkipOne()
	_ = tc
}

func TestStacksCreator_SkipNone(t *testing.T) {
	tc := New.StackTrace.SkipNone()
	_ = tc
}

func TestStacksTo_String(t *testing.T) {
	s := StacksTo.String(0, 3)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestStacksTo_StringDefault(t *testing.T) {
	s := StacksTo.StringDefault()
	_ = s
}

func TestStacksTo_StringNoCount(t *testing.T) {
	s := StacksTo.StringNoCount(0)
	_ = s
}

func TestStacksTo_Bytes(t *testing.T) {
	b := StacksTo.Bytes(0)
	_ = b
}

func TestStacksTo_BytesDefault(t *testing.T) {
	b := StacksTo.BytesDefault()
	_ = b
}

func TestStacksTo_JsonString(t *testing.T) {
	s := StacksTo.JsonString(0)
	_ = s
}

func TestStacksTo_JsonStringDefault(t *testing.T) {
	s := StacksTo.JsonStringDefault()
	_ = s
}

func TestStacksTo_StringUsingFmt(t *testing.T) {
	s := StacksTo.StringUsingFmt(func(tr *Trace) string {
		return tr.PackageName
	}, 0, 3)
	_ = s
}

func TestFile_Name(t *testing.T) {
	n := File.Name(0)
	if n == "" {
		t.Fatal("expected non-empty")
	}
}

func TestFile_Path(t *testing.T) {
	p := File.Path(0)
	if p == "" {
		t.Fatal("expected non-empty")
	}
}

func TestFile_PathLineSep(t *testing.T) {
	fp, ln := File.PathLineSep(0)
	if fp == "" || ln == 0 {
		t.Fatal("expected non-empty")
	}
}

func TestFile_PathLineSepDefault(t *testing.T) {
	fp, ln := File.PathLineSepDefault()
	if fp == "" || ln == 0 {
		t.Fatal("expected non-empty")
	}
}

func TestFile_FilePathWithLineString(t *testing.T) {
	s := File.FilePathWithLineString(0)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestFile_PathLineStringDefault(t *testing.T) {
	s := File.PathLineStringDefault()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestFile_CurrentFilePath(t *testing.T) {
	s := File.CurrentFilePath()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestDir_CurDir(t *testing.T) {
	d := Dir.CurDir()
	if d == "" {
		t.Fatal("expected non-empty")
	}
}

func TestDir_CurDirJoin(t *testing.T) {
	d := Dir.CurDirJoin("subdir")
	if d == "" {
		t.Fatal("expected non-empty")
	}
}

func TestDir_RepoDir(t *testing.T) {
	d := Dir.RepoDir()
	_ = d
}

func TestDir_RepoDirJoin(t *testing.T) {
	d := Dir.RepoDirJoin("subdir")
	_ = d
}

func TestDir_Get(t *testing.T) {
	d := Dir.Get(0)
	if d == "" {
		t.Fatal("expected non-empty")
	}
}

func TestNewTraceCollection_Default(t *testing.T) {
	tc := New.traces.Default()
	if tc == nil {
		t.Fatal("expected non-nil")
	}
}

func TestNewTraceCollection_Using(t *testing.T) {
	tc := New.traces.Using(false, New.Create(0))
	if tc.Length() != 1 {
		t.Fatal("expected 1")
	}

	tc2 := New.traces.Using(true, New.Create(0))
	if tc2.Length() != 1 {
		t.Fatal("expected 1")
	}

	tc3 := New.traces.Using(false)
	if tc3.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func TestSkippablePrefixes(t *testing.T) {
	if !isSkippablePackage("runtime") {
		t.Fatal("runtime should be skippable")
	}
	if isSkippablePackage("mypackage") {
		t.Fatal("mypackage should not be skippable")
	}
}

func TestTraceCollection_AddsUsingSkipDefault(t *testing.T) {
	tc := New.traces.Empty()
	tc.AddsUsingSkipDefault(0)
	_ = tc
}

func TestTraceCollection_AddsUsingSkipUsingFilter(t *testing.T) {
	tc := New.traces.Empty()
	tc.AddsUsingSkipUsingFilter(true, true, 0, 5, func(tr *Trace) (bool, bool) {
		return true, false
	})
	_ = tc
}

func TestTraceCollection_ConcatNewUsingSkipPlusCount(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	tc2 := tc.ConcatNewUsingSkipPlusCount(0, 3)
	_ = tc2
}

func TestTraceCollection_ConcatNewUsingSkip(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	tc2 := tc.ConcatNewUsingSkip(0)
	_ = tc2
}
