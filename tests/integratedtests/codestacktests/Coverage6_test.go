package codestacktests

import (
	"testing"

	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── currentNameOf ──

func Test_Cov6_CurrentNameOf(t *testing.T) {
	name := codestack.NameOf.MethodStackSkip(0)
	actual := args.Map{"notEmpty": name != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CurrentNameOf", actual)
}

// ── skippablePrefixes ──

func Test_Cov6_SkippablePrefixes(t *testing.T) {
	prefixes := codestack.SkippablePrefixes()
	actual := args.Map{"gt0": len(prefixes) > 0}
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "SkippablePrefixes", actual)
}

// ── New.Default / New.Ptr / New.Skip ──

func Test_Cov6_New_Default(t *testing.T) {
	trace := codestack.New.Default()
	actual := args.Map{"isOkay": trace.IsOkay, "pkgNotEmpty": trace.PackageName != ""}
	expected := args.Map{"isOkay": true, "pkgNotEmpty": true}
	expected.ShouldBeEqual(t, 0, "New.Default", actual)
}

func Test_Cov6_New_Ptr(t *testing.T) {
	trace := codestack.New.Ptr(0)
	actual := args.Map{"notNil": trace != nil, "isOkay": trace.IsOkay}
	expected := args.Map{"notNil": true, "isOkay": true}
	expected.ShouldBeEqual(t, 0, "New.Ptr", actual)
}

func Test_Cov6_New_Skip(t *testing.T) {
	trace := codestack.New.Skip(0)
	actual := args.Map{"isOkay": trace.IsOkay}
	expected := args.Map{"isOkay": true}
	expected.ShouldBeEqual(t, 0, "New.Skip", actual)
}

func Test_Cov6_New_SkipPtr(t *testing.T) {
	trace := codestack.New.SkipPtr(0)
	actual := args.Map{"notNil": trace != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "New.SkipPtr", actual)
}

// ── Trace — all getters ──

func Test_Cov6_Trace_Getters(t *testing.T) {
	trace := codestack.New.Default()
	actual := args.Map{
		"isOkay":     trace.IsOkay,
		"isNotOkay":  trace.IsNotOkay(),
		"isNil":      trace.IsNil(),
		"isNotNil":   trace.IsNotNil(),
		"hasIssues":  trace.HasIssues(),
		"string":     trace.String() != "",
		"shortStr":   trace.ShortString() != "",
		"message":    trace.Message() != "",
		"filePath":   trace.FullFilePath() != "",
		"fileName":   trace.FileName() != "",
		"lineNum":    trace.LineNumber() > 0,
		"fwlStr":     trace.FileWithLineString() != "",
		"jsonStr":    trace.JsonString() != "",
		"pkgMethod":  trace.PackageMethodName != "",
	}
	expected := args.Map{
		"isOkay": true, "isNotOkay": false, "isNil": false, "isNotNil": true,
		"hasIssues": false, "string": true, "shortStr": true, "message": true,
		"filePath": true, "fileName": true, "lineNum": true, "fwlStr": true,
		"jsonStr": true, "pkgMethod": true,
	}
	expected.ShouldBeEqual(t, 0, "Trace getters", actual)
}

func Test_Cov6_Trace_FileWithLine(t *testing.T) {
	trace := codestack.New.Default()
	fwl := trace.FileWithLine()
	actual := args.Map{"notEmpty": fwl.FilePath != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Trace FileWithLine", actual)
}

func Test_Cov6_Trace_Clone(t *testing.T) {
	trace := codestack.New.Default()
	cloned := trace.Clone()
	clonedPtr := trace.ClonePtr()
	actual := args.Map{"pkg": cloned.PackageName, "ptrPkg": clonedPtr.PackageName != ""}
	expected := args.Map{"pkg": trace.PackageName, "ptrPkg": true}
	expected.ShouldBeEqual(t, 0, "Trace Clone", actual)
}

func Test_Cov6_Trace_ClonePtr_Nil(t *testing.T) {
	var trace *codestack.Trace
	actual := args.Map{"isNil": trace.ClonePtr() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Trace ClonePtr nil", actual)
}

func Test_Cov6_Trace_Dispose(t *testing.T) {
	trace := codestack.New.Default()
	trace.Dispose()
	actual := args.Map{"pkg": trace.PackageName}
	expected := args.Map{"pkg": ""}
	expected.ShouldBeEqual(t, 0, "Trace Dispose", actual)
}

func Test_Cov6_Trace_JsonModel(t *testing.T) {
	trace := codestack.New.Default()
	model := trace.JsonModel()
	modelAny := trace.JsonModelAny()
	actual := args.Map{"pkg": model.PackageName != "", "anyNotNil": modelAny != nil}
	expected := args.Map{"pkg": true, "anyNotNil": true}
	expected.ShouldBeEqual(t, 0, "Trace JsonModel", actual)
}

func Test_Cov6_Trace_Json(t *testing.T) {
	trace := codestack.New.Default()
	r := trace.Json()
	rp := trace.JsonPtr()
	actual := args.Map{"hasBytes": r.HasBytes(), "ptrNotNil": rp != nil}
	expected := args.Map{"hasBytes": true, "ptrNotNil": true}
	expected.ShouldBeEqual(t, 0, "Trace Json", actual)
}

// ── FileWithLine — all methods ──

func Test_Cov6_FileWithLine_Basic(t *testing.T) {
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 42}
	actual := args.Map{
		"path":    fwl.FullFilePath(),
		"line":    fwl.LineNumber(),
		"isNil":   fwl.IsNil(),
		"notNil":  fwl.IsNotNil(),
		"string":  fwl.String() != "",
		"fwlStr":  fwl.FileWithLine() != "",
		"jsonStr": fwl.JsonString() != "",
	}
	expected := args.Map{
		"path": "/tmp/test.go", "line": 42, "isNil": false, "notNil": true,
		"string": true, "fwlStr": true, "jsonStr": true,
	}
	expected.ShouldBeEqual(t, 0, "FileWithLine basic", actual)
}

func Test_Cov6_FileWithLine_Nil(t *testing.T) {
	var fwl *codestack.FileWithLine
	actual := args.Map{"isNil": fwl.IsNil(), "str": fwl.String()}
	expected := args.Map{"isNil": true, "str": ""}
	expected.ShouldBeEqual(t, 0, "FileWithLine nil", actual)
}

func Test_Cov6_FileWithLine_Clone(t *testing.T) {
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 42}
	cloned := fwl.Clone()
	clonedPtr := fwl.ClonePtr()
	var nilFwl *codestack.FileWithLine
	actual := args.Map{"path": cloned.FilePath, "ptrPath": clonedPtr.FilePath, "nilClone": nilFwl.ClonePtr() == nil}
	expected := args.Map{"path": "/tmp/test.go", "ptrPath": "/tmp/test.go", "nilClone": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine Clone", actual)
}

func Test_Cov6_FileWithLine_JsonModel(t *testing.T) {
	fwl := codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 42}
	model := fwl.JsonModel()
	modelAny := fwl.JsonModelAny()
	actual := args.Map{"path": model.FilePath, "anyNotNil": modelAny != nil}
	expected := args.Map{"path": "/tmp/test.go", "anyNotNil": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine JsonModel", actual)
}

func Test_Cov6_FileWithLine_StringUsingFmt(t *testing.T) {
	fwl := codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 42}
	result := fwl.StringUsingFmt(func(f codestack.FileWithLine) string { return f.FilePath })
	actual := args.Map{"result": result}
	expected := args.Map{"result": "/tmp/test.go"}
	expected.ShouldBeEqual(t, 0, "FileWithLine StringUsingFmt", actual)
}

func Test_Cov6_FileWithLine_Json(t *testing.T) {
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 42}
	r := fwl.Json()
	rp := fwl.JsonPtr()
	actual := args.Map{"hasBytes": r.HasBytes(), "ptrNotNil": rp != nil}
	expected := args.Map{"hasBytes": true, "ptrNotNil": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine Json", actual)
}

func Test_Cov6_FileWithLine_Dispose(t *testing.T) {
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 42}
	fwl.Dispose()
	actual := args.Map{"path": fwl.FilePath}
	expected := args.Map{"path": ""}
	expected.ShouldBeEqual(t, 0, "FileWithLine Dispose", actual)
}

func Test_Cov6_FileWithLine_Dispose_Nil(t *testing.T) {
	var fwl *codestack.FileWithLine
	fwl.Dispose()
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine Dispose nil", actual)
}

// ── StacksTo ──

func Test_Cov6_StacksTo_String(t *testing.T) {
	tc := codestack.New.StackTrace.Default(1, 3)
	result := codestack.StacksTo.String(tc)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo String", actual)
}

// ── Dir ──

func Test_Cov6_Dir_CurrentDir(t *testing.T) {
	dir := codestack.Dir.CurrentDir()
	actual := args.Map{"notEmpty": len(dir) > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir CurrentDir", actual)
}

func Test_Cov6_Dir_TestDir(t *testing.T) {
	dir := codestack.Dir.TestDir()
	actual := args.Map{"notEmpty": len(dir) > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir TestDir", actual)
}

// ── File ──

func Test_Cov6_File_CurrentFileName(t *testing.T) {
	file := codestack.File.CurrentFileName()
	actual := args.Map{"notEmpty": len(file) > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "File CurrentFileName", actual)
}
