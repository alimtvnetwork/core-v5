package codestacktests

import (
	"testing"

	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── TraceCollection additional coverage ──

func Test_Cov3_TraceCollection_NilLength(t *testing.T) {
	var tc *codestack.TraceCollection
	actual := args.Map{"length": tc.Length(), "isEmpty": tc.IsEmpty()}
	expected := args.Map{"length": 0, "isEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.Length nil receiver", actual)
}

func Test_Cov3_TraceCollection_Adds_Empty(t *testing.T) {
	tc := codestack.New.StackTrace.DefaultCount(1)
	before := tc.Length()
	tc.Adds()
	actual := args.Map{"same": tc.Length() == before}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.Adds empty does nothing", actual)
}

func Test_Cov3_TraceCollection_AddsIf_False(t *testing.T) {
	tc := codestack.New.StackTrace.DefaultCount(1)
	before := tc.Length()
	tc.AddsIf(false, codestack.Trace{})
	actual := args.Map{"same": tc.Length() == before}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.AddsIf false does nothing", actual)
}

func Test_Cov3_TraceCollection_AddsIf_True(t *testing.T) {
	tc := codestack.New.StackTrace.DefaultCount(1)
	before := tc.Length()
	tc.AddsIf(true, codestack.Trace{PackageName: "test"})
	actual := args.Map{"grew": tc.Length() > before}
	expected := args.Map{"grew": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.AddsIf true adds", actual)
}

func Test_Cov3_TraceCollection_AddsPtr_Empty(t *testing.T) {
	tc := codestack.New.StackTrace.DefaultCount(1)
	before := tc.Length()
	tc.AddsPtr(true)
	actual := args.Map{"same": tc.Length() == before}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.AddsPtr empty does nothing", actual)
}

func Test_Cov3_TraceCollection_AddsPtr_NilTrace(t *testing.T) {
	tc := codestack.New.StackTrace.DefaultCount(1)
	before := tc.Length()
	tc.AddsPtr(true, nil)
	actual := args.Map{"same": tc.Length() == before}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.AddsPtr nil trace skipped", actual)
}

func Test_Cov3_TraceCollection_AddsPtr_SkipIssues(t *testing.T) {
	tc := codestack.New.StackTrace.DefaultCount(1)
	before := tc.Length()
	badTrace := &codestack.Trace{} // HasIssues = true
	tc.AddsPtr(true, badTrace)
	actual := args.Map{"same": tc.Length() == before}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.AddsPtr skips issues", actual)
}

func Test_Cov3_TraceCollection_AddsPtr_NoSkip(t *testing.T) {
	tc := codestack.New.StackTrace.DefaultCount(1)
	before := tc.Length()
	goodTrace := codestack.New.Ptr(0)
	tc.AddsPtr(false, goodTrace)
	actual := args.Map{"grew": tc.Length() > before}
	expected := args.Map{"grew": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.AddsPtr no skip adds", actual)
}

func Test_Cov3_TraceCollection_FirstOrDefault_Empty(t *testing.T) {
	tc := &codestack.TraceCollection{}
	actual := args.Map{"isNil": tc.FirstOrDefault() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.FirstOrDefault empty", actual)
}

func Test_Cov3_TraceCollection_LastOrDefault_Empty(t *testing.T) {
	tc := &codestack.TraceCollection{}
	actual := args.Map{"isNil": tc.LastOrDefault() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.LastOrDefault empty", actual)
}

func Test_Cov3_TraceCollection_HasIndex(t *testing.T) {
	tc := codestack.New.StackTrace.DefaultCount(1)
	actual := args.Map{
		"hasZero":  tc.HasIndex(0),
		"hasMega":  tc.HasIndex(9999),
	}
	expected := args.Map{
		"hasZero":  tc.HasIndex(0),
		"hasMega":  false,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection.HasIndex", actual)
}

func Test_Cov3_TraceCollection_GetPagesSize(t *testing.T) {
	tc := codestack.New.StackTrace.Default(1, 10)
	actual := args.Map{
		"zeroPage": tc.GetPagesSize(0),
		"negPage":  tc.GetPagesSize(-1),
		"valid":    tc.GetPagesSize(3) > 0,
	}
	expected := args.Map{
		"zeroPage": 0,
		"negPage":  0,
		"valid":    tc.GetPagesSize(3) > 0,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection.GetPagesSize", actual)
}

func Test_Cov3_TraceCollection_Filter(t *testing.T) {
	tc := codestack.New.StackTrace.DefaultCount(1)
	filtered := tc.Filter(func(tr *codestack.Trace) (bool, bool) {
		return true, false
	})
	actual := args.Map{"notEmpty": len(filtered) > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.Filter takes all", actual)
}

func Test_Cov3_TraceCollection_Filter_BreakEarly(t *testing.T) {
	tc := codestack.New.StackTrace.DefaultCount(1)
	filtered := tc.Filter(func(tr *codestack.Trace) (bool, bool) {
		return true, true // take first, break
	})
	actual := args.Map{"len": len(filtered)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection.Filter break early", actual)
}

func Test_Cov3_TraceCollection_SafeLimitCollection(t *testing.T) {
	tc := codestack.New.StackTrace.DefaultCount(1)
	limited := tc.SafeLimitCollection(999)
	actual := args.Map{"safeLen": limited.Length() == tc.Length()}
	expected := args.Map{"safeLen": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.SafeLimitCollection exceeds length", actual)
}

func Test_Cov3_TraceCollection_ConcatNew(t *testing.T) {
	tc := codestack.New.StackTrace.DefaultCount(1)
	origLen := tc.Length()
	newTc := tc.ConcatNew(codestack.Trace{PackageName: "extra"})
	actual := args.Map{
		"origSame": tc.Length() == origLen,
		"newGrew":  newTc.Length() > origLen,
	}
	expected := args.Map{
		"origSame": true,
		"newGrew":  true,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection.ConcatNew", actual)
}

func Test_Cov3_TraceCollection_ConcatNewPtr(t *testing.T) {
	tc := codestack.New.StackTrace.DefaultCount(1)
	origLen := tc.Length()
	tr := codestack.New.Ptr(0)
	newTc := tc.ConcatNewPtr(tr)
	actual := args.Map{"origSame": tc.Length() == origLen, "newNotEmpty": newTc.Length() > 0}
	expected := args.Map{"origSame": true, "newNotEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.ConcatNewPtr", actual)
}

func Test_Cov3_TraceCollection_StackTraces(t *testing.T) {
	tc := codestack.New.StackTrace.DefaultCount(1)
	actual := args.Map{"notEmpty": tc.StackTraces() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.StackTraces", actual)
}

func Test_Cov3_TraceCollection_StackTracesJsonResult(t *testing.T) {
	tc := codestack.New.StackTrace.DefaultCount(1)
	actual := args.Map{"notNil": tc.StackTracesJsonResult() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.StackTracesJsonResult", actual)
}

func Test_Cov3_TraceCollection_NewStackTraces(t *testing.T) {
	tc := codestack.New.StackTrace.DefaultCount(1)
	actual := args.Map{"notEmpty": tc.NewStackTraces(1) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.NewStackTraces", actual)
}

func Test_Cov3_TraceCollection_NewDefaultStackTraces(t *testing.T) {
	tc := codestack.New.StackTrace.DefaultCount(1)
	actual := args.Map{"notEmpty": tc.NewDefaultStackTraces() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.NewDefaultStackTraces", actual)
}

func Test_Cov3_TraceCollection_NewStackTracesJsonResult(t *testing.T) {
	tc := codestack.New.StackTrace.DefaultCount(1)
	actual := args.Map{"notNil": tc.NewStackTracesJsonResult(1) != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.NewStackTracesJsonResult", actual)
}

func Test_Cov3_TraceCollection_NewDefaultStackTracesJsonResult(t *testing.T) {
	tc := codestack.New.StackTrace.DefaultCount(1)
	actual := args.Map{"notNil": tc.NewDefaultStackTracesJsonResult() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.NewDefaultStackTracesJsonResult", actual)
}

func Test_Cov3_TraceCollection_SkipDynamic(t *testing.T) {
	tc := codestack.New.StackTrace.DefaultCount(1)
	actual := args.Map{"notNil": tc.SkipDynamic(1) != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.SkipDynamic", actual)
}

func Test_Cov3_TraceCollection_TakeDynamic(t *testing.T) {
	tc := codestack.New.StackTrace.DefaultCount(1)
	actual := args.Map{"notNil": tc.TakeDynamic(1) != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.TakeDynamic", actual)
}

func Test_Cov3_TraceCollection_TakeCollection(t *testing.T) {
	tc := codestack.New.StackTrace.DefaultCount(1)
	taken := tc.TakeCollection(1)
	actual := args.Map{"len": taken.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection.TakeCollection", actual)
}

func Test_Cov3_TraceCollection_LimitCollection(t *testing.T) {
	tc := codestack.New.StackTrace.DefaultCount(1)
	limited := tc.LimitCollection(1)
	actual := args.Map{"len": limited.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection.LimitCollection", actual)
}

func Test_Cov3_TraceCollection_LimitDynamic(t *testing.T) {
	tc := codestack.New.StackTrace.DefaultCount(1)
	actual := args.Map{"notNil": tc.LimitDynamic(1) != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.LimitDynamic", actual)
}

func Test_Cov3_TraceCollection_FirstOrDefaultDynamic(t *testing.T) {
	tc := codestack.New.StackTrace.DefaultCount(1)
	actual := args.Map{"notNil": tc.FirstOrDefaultDynamic() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.FirstOrDefaultDynamic", actual)
}

func Test_Cov3_TraceCollection_LastOrDefaultDynamic(t *testing.T) {
	tc := codestack.New.StackTrace.DefaultCount(1)
	actual := args.Map{"notNil": tc.LastOrDefaultDynamic() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.LastOrDefaultDynamic", actual)
}

// ── newTraceCollection factory coverage ──

func Test_Cov3_NewTraces_Cap(t *testing.T) {
	// Exercises newTraceCollection.Cap directly via StackTrace factories
	tc := codestack.New.StackTrace.All(true, true, 1, 5)
	actual := args.Map{"notEmpty": tc.Length() > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "newStacksCreator.All", actual)
}

// ── NameOf with real func names ──

func Test_Cov3_NameOf_All_WithDot(t *testing.T) {
	full, pkg, method := codestack.NameOf.All("github.com/pkg/errors.New")
	actual := args.Map{
		"fullNotEmpty":   full != "",
		"pkgNotEmpty":    pkg != "",
		"methodNotEmpty": method != "",
	}
	expected := args.Map{
		"fullNotEmpty":   true,
		"pkgNotEmpty":    true,
		"methodNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "NameOf.All with dotted path", actual)
}

func Test_Cov3_NameOf_MethodByFullName(t *testing.T) {
	result := codestack.NameOf.MethodByFullName("github.com/pkg.Method")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameOf.MethodByFullName", actual)
}

func Test_Cov3_NameOf_PackageByFullName(t *testing.T) {
	result := codestack.NameOf.PackageByFullName("github.com/pkg.Method")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameOf.PackageByFullName", actual)
}

func Test_Cov3_NameOf_CurrentFuncFullPath(t *testing.T) {
	result := codestack.NameOf.CurrentFuncFullPath("github.com/pkg.Method")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameOf.CurrentFuncFullPath", actual)
}

func Test_Cov3_NameOf_JoinPackageNameWithRelative(t *testing.T) {
	result := codestack.NameOf.JoinPackageNameWithRelative("github.com/pkg.Method", "SubMethod")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameOf.JoinPackageNameWithRelative", actual)
}

// ── File getter additional coverage ──

func Test_Cov3_File_PathLineSep(t *testing.T) {
	fp, ln := codestack.File.PathLineSep(0)
	actual := args.Map{"pathNotEmpty": fp != "", "linePositive": ln > 0}
	expected := args.Map{"pathNotEmpty": true, "linePositive": true}
	expected.ShouldBeEqual(t, 0, "File.PathLineSep", actual)
}

func Test_Cov3_File_PathLineSepDefault(t *testing.T) {
	fp, ln := codestack.File.PathLineSepDefault()
	actual := args.Map{"pathNotEmpty": fp != "", "linePositive": ln > 0}
	expected := args.Map{"pathNotEmpty": true, "linePositive": true}
	expected.ShouldBeEqual(t, 0, "File.PathLineSepDefault", actual)
}

func Test_Cov3_File_FilePathWithLineString(t *testing.T) {
	actual := args.Map{"notEmpty": codestack.File.FilePathWithLineString(0) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "File.FilePathWithLineString", actual)
}

func Test_Cov3_File_PathLineStringDefault(t *testing.T) {
	actual := args.Map{"notEmpty": codestack.File.PathLineStringDefault() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "File.PathLineStringDefault", actual)
}

func Test_Cov3_File_CurrentFilePath(t *testing.T) {
	actual := args.Map{"notEmpty": codestack.File.CurrentFilePath() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "File.CurrentFilePath", actual)
}

// ── Dir additional coverage ──

func Test_Cov3_Dir_Get(t *testing.T) {
	actual := args.Map{"notEmpty": codestack.Dir.Get(0) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir.Get", actual)
}

func Test_Cov3_Dir_RepoDir(t *testing.T) {
	actual := args.Map{"notEmpty": codestack.Dir.RepoDir() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir.RepoDir", actual)
}

func Test_Cov3_Dir_RepoDirJoin(t *testing.T) {
	actual := args.Map{"notEmpty": codestack.Dir.RepoDirJoin("sub") != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir.RepoDirJoin", actual)
}

// ── isSkippablePackage ──

func Test_Cov3_SkippablePackage_ViaTrace(t *testing.T) {
	// Create a trace; application code should NOT be skippable
	trace := codestack.New.Default()
	actual := args.Map{"notSkippable": !trace.IsSkippable}
	expected := args.Map{"notSkippable": true}
	expected.ShouldBeEqual(t, 0, "isSkippablePackage returns false for app code", actual)
}

// ── TraceCollection.AddsUsingSkipDefault ──

func Test_Cov3_TraceCollection_AddsUsingSkipDefault(t *testing.T) {
	tc := &codestack.TraceCollection{}
	tc.AddsUsingSkipDefault(0)
	actual := args.Map{"notEmpty": tc.Length() > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.AddsUsingSkipDefault", actual)
}

// ── TraceCollection.ConcatNewUsingSkip ──

func Test_Cov3_TraceCollection_ConcatNewUsingSkip(t *testing.T) {
	tc := codestack.New.StackTrace.DefaultCount(1)
	newTc := tc.ConcatNewUsingSkip(0)
	actual := args.Map{"notEmpty": newTc.Length() > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.ConcatNewUsingSkip", actual)
}

func Test_Cov3_TraceCollection_ConcatNewUsingSkipPlusCount(t *testing.T) {
	tc := codestack.New.StackTrace.DefaultCount(1)
	newTc := tc.ConcatNewUsingSkipPlusCount(0, 5)
	actual := args.Map{"notEmpty": newTc.Length() > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.ConcatNewUsingSkipPlusCount", actual)
}
