package reflectinternal

import (
	"testing"
)

func TestStackTrace_New(t *testing.T) {
	st := CodeStack.New(0)
	if !st.IsOkay {
		t.Fatal("should be okay")
	}
}

func TestStackTrace_NewDefault(t *testing.T) {
	st := CodeStack.NewDefault()
	if !st.IsOkay {
		t.Fatal("should be okay")
	}
}

func TestStackTrace_Message(t *testing.T) {
	st := CodeStack.New(0)
	if st.Message() == "" {
		t.Fatal("expected non-empty")
	}
}

func TestStackTrace_ShortString(t *testing.T) {
	st := CodeStack.New(0)
	if st.ShortString() == "" {
		t.Fatal("expected non-empty")
	}
}

func TestStackTrace_IsNil(t *testing.T) {
	st := &StackTrace{}
	if st.IsNil() {
		t.Fatal("should not be nil")
	}
	var nilST *StackTrace
	if !nilST.IsNil() {
		t.Fatal("should be nil")
	}
}

func TestStackTrace_HasIssues(t *testing.T) {
	st := CodeStack.New(0)
	if st.HasIssues() {
		t.Fatal("should not have issues")
	}
	bad := &StackTrace{}
	if !bad.HasIssues() {
		t.Fatal("should have issues")
	}
	var nilST *StackTrace
	if !nilST.HasIssues() {
		t.Fatal("nil should have issues")
	}
}

func TestStackTrace_IsNotNil(t *testing.T) {
	st := &StackTrace{}
	if !st.IsNotNil() {
		t.Fatal("should be not nil")
	}
}

func TestStackTrace_String(t *testing.T) {
	st := CodeStack.New(0)
	if st.String() == "" {
		t.Fatal("expected non-empty")
	}
	var nilST *StackTrace
	if nilST.String() != "" {
		t.Fatal("expected empty for nil")
	}
}

func TestStackTrace_StringUsingFmt(t *testing.T) {
	st := CodeStack.New(0)
	s := st.StringUsingFmt(func(trace StackTrace) string {
		return trace.PackageName
	})
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestStackTrace_FileWithLine(t *testing.T) {
	st := CodeStack.New(0)
	fwl := st.FileWithLine()
	if fwl.FilePath == "" {
		t.Fatal("expected file path")
	}
}

func TestStackTrace_FullFilePath(t *testing.T) {
	st := CodeStack.New(0)
	if st.FullFilePath() == "" {
		t.Fatal("expected non-empty")
	}
}

func TestStackTrace_FileName(t *testing.T) {
	st := CodeStack.New(0)
	if st.FileName() == "" {
		t.Fatal("expected non-empty")
	}
}

func TestStackTrace_LineNumber(t *testing.T) {
	st := CodeStack.New(0)
	if st.LineNumber() == 0 {
		t.Fatal("expected non-zero")
	}
}

func TestStackTrace_FileWithLineString(t *testing.T) {
	st := CodeStack.New(0)
	if st.FileWithLineString() == "" {
		t.Fatal("expected non-empty")
	}
}

func TestStackTrace_JsonModel(t *testing.T) {
	st := CodeStack.New(0)
	_ = st.JsonModel()
	_ = st.JsonModelAny()
}

func TestStackTrace_Dispose(t *testing.T) {
	st := CodeStack.New(0)
	st.Dispose()
	if st.PackageName != "" {
		t.Fatal("expected empty after dispose")
	}
	var nilST *StackTrace
	nilST.Dispose()
}

func TestStackTrace_JsonString(t *testing.T) {
	st := CodeStack.New(0)
	if st.JsonString() == "" {
		t.Fatal("expected non-empty")
	}
}

func TestStackTrace_Clone(t *testing.T) {
	st := CodeStack.New(0)
	c := st.Clone()
	if c.PackageName != st.PackageName {
		t.Fatal("clone mismatch")
	}
}

func TestStackTrace_ClonePtr(t *testing.T) {
	st := CodeStack.New(0)
	cp := st.ClonePtr()
	if cp == nil {
		t.Fatal("expected non-nil")
	}
	var nilST *StackTrace
	if nilST.ClonePtr() != nil {
		t.Fatal("expected nil")
	}
}

func TestCodeStack_MethodName(t *testing.T) {
	m := CodeStack.MethodName(0)
	if m == "" {
		t.Fatal("expected non-empty")
	}
}

func TestCodeStack_MethodNameWithLine(t *testing.T) {
	m := CodeStack.MethodNameWithLine(0)
	if m == "" {
		t.Fatal("expected non-empty")
	}
}

func TestCodeStack_FileWithLine(t *testing.T) {
	s := CodeStack.FileWithLine(0)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestCodeStack_LastFileWithLines(t *testing.T) {
	lines := CodeStack.LastFileWithLines(0, 3)
	if len(lines) != 3 {
		t.Fatal("expected 3")
	}
}

func TestCodeStack_LastFileWithLine(t *testing.T) {
	s := CodeStack.LastFileWithLine(0, 3)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestCodeStack_FilePath(t *testing.T) {
	p := CodeStack.FilePath(0)
	if p == "" {
		t.Fatal("expected non-empty")
	}
}

func TestCodeStack_NewFileWithLines(t *testing.T) {
	lines := CodeStack.NewFileWithLines(0, 3)
	if len(lines) == 0 {
		t.Fatal("expected non-empty")
	}
}

func TestCodeStack_NewFileWithLine(t *testing.T) {
	fwl := CodeStack.NewFileWithLine(0)
	if fwl.FilePath == "" {
		t.Fatal("expected non-empty")
	}
}

func TestCodeStack_NewStacks(t *testing.T) {
	stacks := CodeStack.NewStacks(0, 3)
	if len(stacks) != 3 {
		t.Fatal("expected 3")
	}
}

func TestCodeStack_StacksStrings(t *testing.T) {
	s := CodeStack.StacksStrings(0)
	_ = s
}

func TestCodeStack_StacksStringsCount(t *testing.T) {
	s := CodeStack.StacksStringsCount(0, 3)
	_ = s
}

func TestCodeStack_StacksStringsFiltered(t *testing.T) {
	s := CodeStack.StacksStringsFiltered(0, 5)
	_ = s
}

func TestCodeStack_StacksString(t *testing.T) {
	s := CodeStack.StacksString(0)
	_ = s
}

func TestCodeStack_StacksStringDefault(t *testing.T) {
	s := CodeStack.StacksStringDefault(0)
	_ = s
}

func TestCodeStack_StacksStringCount(t *testing.T) {
	s := CodeStack.StacksStringCount(0, 3)
	_ = s
}

func TestCodeStack_SingleStack(t *testing.T) {
	s := CodeStack.SingleStack(0)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestIsSystemLibraryPath(t *testing.T) {
	if !isSystemLibraryPath("/usr/local/go", "/usr/local/go/src/fmt/print.go") {
		t.Fatal("should be system path")
	}
	if isSystemLibraryPath("/usr/local/go", "/home/user/project/main.go") {
		t.Fatal("should not be system path")
	}
	if isSystemLibraryPath("", "/anything") {
		t.Fatal("empty goroot should not match")
	}
}
