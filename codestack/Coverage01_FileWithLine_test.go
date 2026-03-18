package codestack

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

func TestFileWithLine_FullFilePath(t *testing.T) {
	f := FileWithLine{FilePath: "/tmp/test.go", Line: 42}
	if f.FullFilePath() != "/tmp/test.go" {
		t.Fatal("unexpected")
	}
}

func TestFileWithLine_LineNumber(t *testing.T) {
	f := FileWithLine{FilePath: "/tmp/test.go", Line: 42}
	if f.LineNumber() != 42 {
		t.Fatal("unexpected")
	}
}

func TestFileWithLine_IsNil_IsNotNil(t *testing.T) {
	f := &FileWithLine{}
	if f.IsNil() {
		t.Fatal("should not be nil")
	}
	if !f.IsNotNil() {
		t.Fatal("should be not nil")
	}

	var nilF *FileWithLine
	if !nilF.IsNil() {
		t.Fatal("should be nil")
	}
	if nilF.IsNotNil() {
		t.Fatal("should not be not nil")
	}
}

func TestFileWithLine_String(t *testing.T) {
	f := &FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	s := f.String()
	if s == "" {
		t.Fatal("expected non-empty")
	}

	var nilF *FileWithLine
	if nilF.String() != "" {
		t.Fatal("expected empty for nil")
	}
}

func TestFileWithLine_StringUsingFmt(t *testing.T) {
	f := FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	s := f.StringUsingFmt(func(fwl FileWithLine) string {
		return fwl.FilePath
	})
	if s != "/tmp/test.go" {
		t.Fatal("unexpected")
	}
}

func TestFileWithLine_FileWithLineMethod(t *testing.T) {
	f := &FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	s := f.FileWithLine()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestFileWithLine_JsonModel(t *testing.T) {
	f := FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	m := f.JsonModel()
	if m.FilePath != "/tmp/test.go" {
		t.Fatal("unexpected")
	}
}

func TestFileWithLine_JsonModelAny(t *testing.T) {
	f := &FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	a := f.JsonModelAny()
	if a == nil {
		t.Fatal("expected non-nil")
	}
}

func TestFileWithLine_JsonString(t *testing.T) {
	f := &FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	s := f.JsonString()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestFileWithLine_Json_JsonPtr(t *testing.T) {
	f := FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	j := f.Json()
	if j.HasError() {
		t.Fatal(j.Error)
	}
	jp := f.JsonPtr()
	if jp == nil {
		t.Fatal("expected non-nil")
	}
}

func TestFileWithLine_ParseInjectUsingJson(t *testing.T) {
	f := FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	jr := corejson.NewPtr(f)
	target := &FileWithLine{}
	result, err := target.ParseInjectUsingJson(jr)
	if err != nil {
		t.Fatal(err)
	}
	if result.FilePath != "/tmp/test.go" {
		t.Fatal("unexpected")
	}

	// error case
	badJr := corejson.NewResult.UsingBytes([]byte("invalid"))
	_, err2 := target.ParseInjectUsingJson(badJr.Ptr())
	if err2 == nil {
		t.Fatal("expected error")
	}
}

func TestFileWithLine_ParseInjectUsingJsonMust(t *testing.T) {
	f := FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	jr := corejson.NewPtr(f)
	target := &FileWithLine{}
	result := target.ParseInjectUsingJsonMust(jr)
	if result.FilePath != "/tmp/test.go" {
		t.Fatal("unexpected")
	}
}

func TestFileWithLine_JsonParseSelfInject(t *testing.T) {
	f := FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	jr := corejson.NewPtr(f)
	target := &FileWithLine{}
	err := target.JsonParseSelfInject(jr)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFileWithLine_AsFileLiner(t *testing.T) {
	f := &FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	liner := f.AsFileLiner()
	if liner == nil {
		t.Fatal("expected non-nil")
	}
	if liner.FullFilePath() != "/tmp/test.go" {
		t.Fatal("unexpected")
	}
}
