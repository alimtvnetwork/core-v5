package codestacktests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── newCreator ──

func Test_Cov4_NewCreator_Default(t *testing.T) {
	trace := codestack.New.Default()
	actual := args.Map{"notNil": trace != nil, "hasFunc": trace.FuncName != ""}
	expected := args.Map{"notNil": true, "hasFunc": true}
	expected.ShouldBeEqual(t, 0, "New.Default returns valid Trace -- from test func", actual)
}

func Test_Cov4_NewCreator_SkipOne(t *testing.T) {
	trace := codestack.New.SkipOne()
	actual := args.Map{"notNil": trace != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "New.SkipOne returns valid Trace -- skip one frame", actual)
}

func Test_Cov4_NewCreator_Ptr(t *testing.T) {
	trace := codestack.New.Ptr()
	actual := args.Map{"notNil": trace != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "New.Ptr returns valid Trace ptr -- default skip", actual)
}

func Test_Cov4_NewCreator_Create(t *testing.T) {
	trace := codestack.New.Create(1)
	actual := args.Map{"notNil": trace != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "New.Create returns valid Trace -- skip 1", actual)
}

// ── Trace methods ──

func Test_Cov4_Trace_Message(t *testing.T) {
	trace := codestack.New.Default()
	msg := trace.Message()
	actual := args.Map{"hasContent": len(msg) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Trace.Message returns non-empty -- default", actual)
}

func Test_Cov4_Trace_ShortString(t *testing.T) {
	trace := codestack.New.Default()
	s := trace.ShortString()
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Trace.ShortString returns non-empty -- default", actual)
}

func Test_Cov4_Trace_IsNilNotNil(t *testing.T) {
	trace := codestack.New.Default()
	var nilTrace *codestack.Trace
	actual := args.Map{
		"isNil":    trace.IsNil(),
		"isNotNil": trace.IsNotNil(),
		"nilIsNil": nilTrace.IsNil(),
	}
	expected := args.Map{
		"isNil":    false,
		"isNotNil": true,
		"nilIsNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Trace IsNil/IsNotNil returns correct -- valid and nil", actual)
}

func Test_Cov4_Trace_HasIssues(t *testing.T) {
	trace := codestack.New.Default()
	actual := args.Map{"hasIssues": trace.HasIssues()}
	expected := args.Map{"hasIssues": false}
	expected.ShouldBeEqual(t, 0, "Trace.HasIssues returns false -- valid trace", actual)
}

func Test_Cov4_Trace_String(t *testing.T) {
	trace := codestack.New.Default()
	s := trace.String()
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Trace.String returns non-empty -- default", actual)
}

func Test_Cov4_Trace_StringUsingFmt(t *testing.T) {
	trace := codestack.New.Default()
	s := trace.StringUsingFmt(" | ")
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Trace.StringUsingFmt returns non-empty -- pipe sep", actual)
}

func Test_Cov4_Trace_FileWithLine(t *testing.T) {
	trace := codestack.New.Default()
	fwl := trace.FileWithLine()
	actual := args.Map{"notNil": fwl != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Trace.FileWithLine returns non-nil -- default", actual)
}

func Test_Cov4_Trace_FullFilePath(t *testing.T) {
	trace := codestack.New.Default()
	fp := trace.FullFilePath()
	actual := args.Map{"hasContent": len(fp) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Trace.FullFilePath returns non-empty -- default", actual)
}

func Test_Cov4_Trace_FileName(t *testing.T) {
	trace := codestack.New.Default()
	fn := trace.FileName()
	actual := args.Map{"hasContent": len(fn) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Trace.FileName returns non-empty -- default", actual)
}

func Test_Cov4_Trace_LineNumber(t *testing.T) {
	trace := codestack.New.Default()
	ln := trace.LineNumber()
	actual := args.Map{"positive": ln > 0}
	expected := args.Map{"positive": true}
	expected.ShouldBeEqual(t, 0, "Trace.LineNumber returns positive -- default", actual)
}

func Test_Cov4_Trace_FileWithLineString(t *testing.T) {
	trace := codestack.New.Default()
	s := trace.FileWithLineString()
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Trace.FileWithLineString returns non-empty -- default", actual)
}

func Test_Cov4_Trace_Clone(t *testing.T) {
	trace := codestack.New.Default()
	cloned := trace.Clone()
	actual := args.Map{
		"notNil":   cloned.FuncName != "",
		"sameFunc": cloned.FuncName == trace.FuncName,
	}
	expected := args.Map{
		"notNil":   true,
		"sameFunc": true,
	}
	expected.ShouldBeEqual(t, 0, "Trace.Clone returns same data -- default", actual)
}

func Test_Cov4_Trace_ClonePtr(t *testing.T) {
	trace := codestack.New.Ptr()
	cloned := trace.ClonePtr()
	actual := args.Map{
		"notNil":     cloned != nil,
		"notSamePtr": cloned != trace,
		"sameFunc":   cloned.FuncName == trace.FuncName,
	}
	expected := args.Map{
		"notNil":     true,
		"notSamePtr": true,
		"sameFunc":   true,
	}
	expected.ShouldBeEqual(t, 0, "Trace.ClonePtr returns different ptr same data -- default", actual)
}

func Test_Cov4_Trace_ClonePtr_Nil(t *testing.T) {
	var trace *codestack.Trace
	cloned := trace.ClonePtr()
	actual := args.Map{"isNil": cloned == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Trace.ClonePtr returns nil -- nil receiver", actual)
}

func Test_Cov4_Trace_AsFileLiner(t *testing.T) {
	trace := codestack.New.Default()
	liner := trace.AsFileLiner()
	actual := args.Map{"notNil": liner != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Trace.AsFileLiner returns non-nil -- default", actual)
}

func Test_Cov4_Trace_JsonModel(t *testing.T) {
	trace := codestack.New.Default()
	model := trace.JsonModel()
	actual := args.Map{"notEmpty": len(model) > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Trace.JsonModel returns non-empty -- default", actual)
}

func Test_Cov4_Trace_JsonModelAny(t *testing.T) {
	trace := codestack.New.Default()
	model := trace.JsonModelAny()
	actual := args.Map{"notNil": model != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Trace.JsonModelAny returns non-nil -- default", actual)
}

func Test_Cov4_Trace_JsonString(t *testing.T) {
	trace := codestack.New.Default()
	js := trace.JsonString()
	actual := args.Map{"hasContent": len(js) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Trace.JsonString returns non-empty -- default", actual)
}

func Test_Cov4_Trace_Json(t *testing.T) {
	trace := codestack.New.Default()
	result := trace.Json()
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Trace.Json returns non-nil result -- default", actual)
}

func Test_Cov4_Trace_JsonPtr(t *testing.T) {
	trace := codestack.New.Default()
	result := trace.JsonPtr()
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Trace.JsonPtr returns non-nil result -- default", actual)
}

func Test_Cov4_Trace_Dispose(t *testing.T) {
	trace := codestack.New.Default()
	trace.Dispose()
	actual := args.Map{"funcEmpty": trace.FuncName == ""}
	expected := args.Map{"funcEmpty": true}
	expected.ShouldBeEqual(t, 0, "Trace.Dispose clears fields -- after dispose", actual)
}

// ── newStacksCreator ──

func Test_Cov4_NewStacks_Default(t *testing.T) {
	traces := codestack.NewStacks.Default()
	actual := args.Map{"hasItems": traces != nil && traces.Length() > 0}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "NewStacks.Default returns non-empty -- from test", actual)
}

func Test_Cov4_NewStacks_DefaultCount(t *testing.T) {
	traces := codestack.NewStacks.DefaultCount(3)
	actual := args.Map{"hasItems": traces != nil}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "NewStacks.DefaultCount returns traces -- count 3", actual)
}

func Test_Cov4_NewStacks_SkipOne(t *testing.T) {
	traces := codestack.NewStacks.SkipOne()
	actual := args.Map{"notNil": traces != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewStacks.SkipOne returns traces -- skip 1", actual)
}

func Test_Cov4_NewStacks_SkipNone(t *testing.T) {
	traces := codestack.NewStacks.SkipNone()
	actual := args.Map{"notNil": traces != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewStacks.SkipNone returns traces -- no skip", actual)
}

func Test_Cov4_NewStacks_All(t *testing.T) {
	traces := codestack.NewStacks.All(1, 5)
	actual := args.Map{"notNil": traces != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewStacks.All returns traces -- skip 1 count 5", actual)
}

// ── TraceCollection basic methods ──

func Test_Cov4_TraceCollection_NewAndBasic(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	actual := args.Map{
		"notNil":  tc != nil,
		"isEmpty": tc.IsEmpty(),
		"length":  tc.Length(),
		"count":   tc.Count(),
	}
	expected := args.Map{
		"notNil":  true,
		"isEmpty": true,
		"length":  0,
		"count":   0,
	}
	expected.ShouldBeEqual(t, 0, "NewTraceCollection.Default returns empty -- new collection", actual)
}

func Test_Cov4_TraceCollection_Add(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	trace := codestack.New.Default()
	tc.Add(trace)
	actual := args.Map{
		"length":     tc.Length(),
		"hasAny":     tc.HasAnyItem(),
		"lastIndex":  tc.LastIndex(),
		"hasIndex0":  tc.HasIndex(0),
		"hasIndex99": tc.HasIndex(99),
	}
	expected := args.Map{
		"length":     1,
		"hasAny":     true,
		"lastIndex":  0,
		"hasIndex0":  true,
		"hasIndex99": false,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection.Add works -- single item", actual)
}

func Test_Cov4_TraceCollection_FirstLast(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	t1 := codestack.New.Default()
	tc.Add(t1)
	actual := args.Map{
		"first":           tc.First().FuncName == t1.FuncName,
		"last":            tc.Last().FuncName == t1.FuncName,
		"firstOrDefault":  tc.FirstOrDefault().FuncName == t1.FuncName,
		"lastOrDefault":   tc.LastOrDefault().FuncName == t1.FuncName,
	}
	expected := args.Map{
		"first":           true,
		"last":            true,
		"firstOrDefault":  true,
		"lastOrDefault":   true,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection First/Last returns correct -- single item", actual)
}

func Test_Cov4_TraceCollection_Strings(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.Add(codestack.New.Default())
	strs := tc.Strings()
	actual := args.Map{"len": len(strs)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection.Strings returns correct len -- single item", actual)
}

func Test_Cov4_TraceCollection_Clone(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.Add(codestack.New.Default())
	cloned := tc.Clone()
	actual := args.Map{
		"sameLen": cloned.Length() == tc.Length(),
	}
	expected := args.Map{
		"sameLen": true,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection.Clone returns same length -- single item", actual)
}

func Test_Cov4_TraceCollection_ClonePtr(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.Add(codestack.New.Default())
	cloned := tc.ClonePtr()
	actual := args.Map{
		"notNil":     cloned != nil,
		"notSamePtr": cloned != tc,
	}
	expected := args.Map{
		"notNil":     true,
		"notSamePtr": true,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection.ClonePtr returns different ptr -- single item", actual)
}

func Test_Cov4_TraceCollection_Join(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.Add(codestack.New.Default())
	joined := tc.Join(", ")
	actual := args.Map{"hasContent": len(joined) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.Join returns non-empty -- single item", actual)
}

func Test_Cov4_TraceCollection_JoinLines(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.Add(codestack.New.Default())
	joined := tc.JoinLines()
	actual := args.Map{"hasContent": len(joined) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.JoinLines returns non-empty -- single item", actual)
}

func Test_Cov4_TraceCollection_String(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.Add(codestack.New.Default())
	s := tc.String()
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.String returns non-empty -- single item", actual)
}

func Test_Cov4_TraceCollection_JsonString(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.Add(codestack.New.Default())
	js := tc.JsonString()
	actual := args.Map{"hasContent": len(js) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.JsonString returns non-empty -- single item", actual)
}

func Test_Cov4_TraceCollection_Clear(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.Add(codestack.New.Default())
	tc.Clear()
	actual := args.Map{"isEmpty": tc.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.Clear empties collection -- after clear", actual)
}

func Test_Cov4_TraceCollection_Dispose(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.Add(codestack.New.Default())
	tc.Dispose()
	actual := args.Map{"isEmpty": tc.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.Dispose empties collection -- after dispose", actual)
}

func Test_Cov4_TraceCollection_Reverse(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.Add(codestack.New.Default())
	tc.Add(codestack.New.SkipOne())
	reversed := tc.Reverse()
	actual := args.Map{"length": reversed.Length()}
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "TraceCollection.Reverse returns same length -- two items", actual)
}

func Test_Cov4_TraceCollection_Skip(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.Add(codestack.New.Default())
	tc.Add(codestack.New.Default())
	skipped := tc.Skip(1)
	actual := args.Map{"len": len(skipped)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection.Skip returns correct len -- skip 1 of 2", actual)
}

func Test_Cov4_TraceCollection_Take(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.Add(codestack.New.Default())
	tc.Add(codestack.New.Default())
	taken := tc.Take(1)
	actual := args.Map{"len": len(taken)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection.Take returns correct len -- take 1 of 2", actual)
}

func Test_Cov4_TraceCollection_Limit(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.Add(codestack.New.Default())
	tc.Add(codestack.New.Default())
	limited := tc.Limit(1)
	actual := args.Map{"len": len(limited)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection.Limit returns correct len -- limit 1 of 2", actual)
}

// ── currentNameOf ──

func Test_Cov4_CurrentNameOf_Method(t *testing.T) {
	name := codestack.CurrentNameOf.Method()
	actual := args.Map{"hasContent": len(name) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "CurrentNameOf.Method returns non-empty -- from test", actual)
}

func Test_Cov4_CurrentNameOf_Package(t *testing.T) {
	name := codestack.CurrentNameOf.Package()
	actual := args.Map{"hasContent": len(name) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "CurrentNameOf.Package returns non-empty -- from test", actual)
}

func Test_Cov4_CurrentNameOf_All(t *testing.T) {
	name := codestack.CurrentNameOf.All()
	actual := args.Map{"hasContent": len(name) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "CurrentNameOf.All returns non-empty -- from test", actual)
}

func Test_Cov4_CurrentNameOf_CurrentFuncFullPath(t *testing.T) {
	name := codestack.CurrentNameOf.CurrentFuncFullPath()
	actual := args.Map{"hasContent": len(name) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "CurrentNameOf.CurrentFuncFullPath returns non-empty -- from test", actual)
}

// ── dirGetter ──

func Test_Cov4_Dir_CurDir(t *testing.T) {
	dir := codestack.Dir.CurDir()
	actual := args.Map{"hasContent": len(dir) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Dir.CurDir returns non-empty -- from test", actual)
}

func Test_Cov4_Dir_CurDirJoin(t *testing.T) {
	dir := codestack.Dir.CurDirJoin("sub")
	actual := args.Map{"hasSub": strings.Contains(dir, "sub")}
	expected := args.Map{"hasSub": true}
	expected.ShouldBeEqual(t, 0, "Dir.CurDirJoin returns path with sub -- from test", actual)
}

func Test_Cov4_Dir_Get(t *testing.T) {
	dir := codestack.Dir.Get(1)
	actual := args.Map{"hasContent": len(dir) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Dir.Get returns non-empty -- skip 1", actual)
}

// ── fileGetter ──

func Test_Cov4_File_Name(t *testing.T) {
	name := codestack.File.Name()
	actual := args.Map{"hasContent": len(name) > 0, "hasGo": strings.HasSuffix(name, ".go")}
	expected := args.Map{"hasContent": true, "hasGo": true}
	expected.ShouldBeEqual(t, 0, "File.Name returns .go file -- from test", actual)
}

func Test_Cov4_File_Path(t *testing.T) {
	path := codestack.File.Path(1)
	actual := args.Map{"hasContent": len(path) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "File.Path returns non-empty -- skip 1", actual)
}

func Test_Cov4_File_CurrentFilePath(t *testing.T) {
	fp := codestack.File.CurrentFilePath()
	actual := args.Map{"hasContent": len(fp) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "File.CurrentFilePath returns non-empty -- from test", actual)
}

func Test_Cov4_File_PathLineSep(t *testing.T) {
	s := codestack.File.PathLineSep(1, ":")
	actual := args.Map{"hasColon": strings.Contains(s, ":")}
	expected := args.Map{"hasColon": true}
	expected.ShouldBeEqual(t, 0, "File.PathLineSep returns path:line -- colon sep", actual)
}

func Test_Cov4_File_PathLineSepDefault(t *testing.T) {
	s := codestack.File.PathLineSepDefault()
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "File.PathLineSepDefault returns non-empty -- default", actual)
}

func Test_Cov4_File_PathLineStringDefault(t *testing.T) {
	s := codestack.File.PathLineStringDefault()
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "File.PathLineStringDefault returns non-empty -- default", actual)
}

func Test_Cov4_File_FilePathWithLineString(t *testing.T) {
	s := codestack.File.FilePathWithLineString(1, ":")
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "File.FilePathWithLineString returns non-empty -- skip 1", actual)
}

// ── stacksTo ──

func Test_Cov4_StacksTo_String(t *testing.T) {
	s := codestack.StacksTo.String(1, 3)
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.String returns non-empty -- skip 1 count 3", actual)
}

func Test_Cov4_StacksTo_StringDefault(t *testing.T) {
	s := codestack.StacksTo.StringDefault()
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.StringDefault returns non-empty -- default", actual)
}

func Test_Cov4_StacksTo_StringNoCount(t *testing.T) {
	s := codestack.StacksTo.StringNoCount(1, 3)
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.StringNoCount returns non-empty -- skip 1 count 3", actual)
}

func Test_Cov4_StacksTo_Bytes(t *testing.T) {
	b := codestack.StacksTo.Bytes(1, 3)
	actual := args.Map{"hasContent": len(b) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.Bytes returns non-empty -- skip 1 count 3", actual)
}

func Test_Cov4_StacksTo_BytesDefault(t *testing.T) {
	b := codestack.StacksTo.BytesDefault()
	actual := args.Map{"hasContent": len(b) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.BytesDefault returns non-empty -- default", actual)
}

func Test_Cov4_StacksTo_JsonString(t *testing.T) {
	s := codestack.StacksTo.JsonString(1, 3)
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.JsonString returns non-empty -- skip 1 count 3", actual)
}

func Test_Cov4_StacksTo_JsonStringDefault(t *testing.T) {
	s := codestack.StacksTo.JsonStringDefault()
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.JsonStringDefault returns non-empty -- default", actual)
}

func Test_Cov4_StacksTo_StringUsingFmt(t *testing.T) {
	s := codestack.StacksTo.StringUsingFmt(1, 3, " | ")
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.StringUsingFmt returns non-empty -- pipe sep", actual)
}

// ── FileWithLine ──

func Test_Cov4_FileWithLine_Basic(t *testing.T) {
	fwl := codestack.NewFileWithLine("/some/path.go", 42)
	actual := args.Map{
		"path":   fwl.FullFilePath(),
		"line":   fwl.LineNumber(),
		"notNil": fwl.IsNotNil(),
		"isNil":  fwl.IsNil(),
		"str":    len(fwl.String()) > 0,
	}
	expected := args.Map{
		"path":   "/some/path.go",
		"line":   42,
		"notNil": true,
		"isNil":  false,
		"str":    true,
	}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns correct values -- basic", actual)
}

func Test_Cov4_FileWithLine_NilSafety(t *testing.T) {
	var fwl *codestack.FileWithLine
	actual := args.Map{"isNil": fwl.IsNil()}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine.IsNil returns true -- nil receiver", actual)
}

func Test_Cov4_FileWithLine_Json(t *testing.T) {
	fwl := codestack.NewFileWithLine("/path.go", 10)
	js := fwl.JsonString()
	actual := args.Map{"hasContent": len(js) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine.JsonString returns non-empty -- valid", actual)
}

func Test_Cov4_FileWithLine_AsFileLiner(t *testing.T) {
	fwl := codestack.NewFileWithLine("/path.go", 10)
	liner := fwl.AsFileLiner()
	actual := args.Map{"notNil": liner != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine.AsFileLiner returns non-nil -- valid", actual)
}

// ── TraceCollection advanced ──

func Test_Cov4_TraceCollection_SkipTakeCollections(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.Add(codestack.New.Default())
	tc.Add(codestack.New.Default())
	tc.Add(codestack.New.Default())

	skipCol := tc.SkipCollection(1)
	takeCol := tc.TakeCollection(2)
	limitCol := tc.LimitCollection(1)
	safeLimitCol := tc.SafeLimitCollection(100)

	actual := args.Map{
		"skipLen":      skipCol.Length(),
		"takeLen":      takeCol.Length(),
		"limitLen":     limitCol.Length(),
		"safeLimitLen": safeLimitCol.Length(),
	}
	expected := args.Map{
		"skipLen":      2,
		"takeLen":      2,
		"limitLen":     1,
		"safeLimitLen": 3,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection Skip/Take/Limit collections -- 3 items", actual)
}

func Test_Cov4_TraceCollection_ConcatNew(t *testing.T) {
	tc1 := codestack.NewTraceCollection.Default()
	tc1.Add(codestack.New.Default())
	tc2 := codestack.NewTraceCollection.Default()
	tc2.Add(codestack.New.Default())
	concat := tc1.ConcatNew(tc2)
	actual := args.Map{"len": concat.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "TraceCollection.ConcatNew merges -- two collections", actual)
}

func Test_Cov4_TraceCollection_ConcatNewPtr(t *testing.T) {
	tc1 := codestack.NewTraceCollection.Default()
	tc1.Add(codestack.New.Default())
	tc2 := codestack.NewTraceCollection.Default()
	tc2.Add(codestack.New.Default())
	concat := tc1.ConcatNewPtr(tc2)
	actual := args.Map{"notNil": concat != nil, "len": concat.Length()}
	expected := args.Map{"notNil": true, "len": 2}
	expected.ShouldBeEqual(t, 0, "TraceCollection.ConcatNewPtr merges -- two collections", actual)
}

func Test_Cov4_TraceCollection_InsertAt(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.Add(codestack.New.Default())
	tc.Add(codestack.New.Default())
	inserted := codestack.New.SkipOne()
	tc.InsertAt(1, inserted)
	actual := args.Map{"len": tc.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "TraceCollection.InsertAt adds at index -- 3 items", actual)
}

func Test_Cov4_TraceCollection_FileWithLines(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.Add(codestack.New.Default())
	fwls := tc.FileWithLines()
	actual := args.Map{"len": len(fwls)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection.FileWithLines returns correct -- single item", actual)
}

func Test_Cov4_TraceCollection_FileWithLinesStrings(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.Add(codestack.New.Default())
	strs := tc.FileWithLinesStrings()
	actual := args.Map{"len": len(strs)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection.FileWithLinesStrings returns 1 -- single item", actual)
}

func Test_Cov4_TraceCollection_ShortStrings(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.Add(codestack.New.Default())
	strs := tc.ShortStrings()
	actual := args.Map{"len": len(strs)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection.ShortStrings returns 1 -- single item", actual)
}

func Test_Cov4_TraceCollection_JoinMethods(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.Add(codestack.New.Default())
	actual := args.Map{
		"joinCsv":       len(tc.JoinCsv()) > 0,
		"joinCsvLine":   len(tc.JoinCsvLine()) > 0,
		"joinShort":     len(tc.JoinShortStrings()) > 0,
		"joinFwl":       len(tc.JoinFileWithLinesStrings()) > 0,
		"joinJson":      len(tc.JoinJsonStrings()) > 0,
		"joinUsingFmt":  len(tc.JoinUsingFmt(" | ")) > 0,
	}
	expected := args.Map{
		"joinCsv":       true,
		"joinCsvLine":   true,
		"joinShort":     true,
		"joinFwl":       true,
		"joinJson":      true,
		"joinUsingFmt":  true,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection join methods return non-empty -- single item", actual)
}

func Test_Cov4_TraceCollection_CodeStacksString(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.Add(codestack.New.Default())
	s := tc.CodeStacksString()
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.CodeStacksString returns non-empty -- single item", actual)
}

func Test_Cov4_TraceCollection_FileWithLinesString(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.Add(codestack.New.Default())
	s := tc.FileWithLinesString()
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.FileWithLinesString returns non-empty -- single item", actual)
}

func Test_Cov4_TraceCollection_CodeStacksStringLimit(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.Add(codestack.New.Default())
	tc.Add(codestack.New.Default())
	s := tc.CodeStacksStringLimit(1)
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.CodeStacksStringLimit returns non-empty -- limit 1", actual)
}

func Test_Cov4_TraceCollection_CsvStrings(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.Add(codestack.New.Default())
	strs := tc.CsvStrings()
	actual := args.Map{"len": len(strs)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection.CsvStrings returns 1 -- single item", actual)
}

func Test_Cov4_TraceCollection_JsonModel(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.Add(codestack.New.Default())
	model := tc.JsonModel()
	actual := args.Map{"notEmpty": len(model) > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.JsonModel returns non-empty -- single item", actual)
}

func Test_Cov4_TraceCollection_Json(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.Add(codestack.New.Default())
	result := tc.Json()
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.Json returns non-nil -- single item", actual)
}

func Test_Cov4_TraceCollection_JsonPtr(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.Add(codestack.New.Default())
	result := tc.JsonPtr()
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.JsonPtr returns non-nil -- single item", actual)
}

func Test_Cov4_TraceCollection_JsonStrings(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.Add(codestack.New.Default())
	strs := tc.JsonStrings()
	actual := args.Map{"len": len(strs)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection.JsonStrings returns 1 -- single item", actual)
}

func Test_Cov4_TraceCollection_StringsUsingFmt(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.Add(codestack.New.Default())
	strs := tc.StringsUsingFmt(" | ")
	actual := args.Map{"len": len(strs)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection.StringsUsingFmt returns 1 -- single item", actual)
}

func Test_Cov4_TraceCollection_IsEqual(t *testing.T) {
	tc1 := codestack.NewTraceCollection.Default()
	tc2 := codestack.NewTraceCollection.Default()
	actual := args.Map{"equal": tc1.IsEqual(tc2)}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.IsEqual returns true -- both empty", actual)
}

func Test_Cov4_TraceCollection_StackTracesBytes(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.Add(codestack.New.Default())
	b := tc.StackTracesBytes()
	actual := args.Map{"hasContent": len(b) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.StackTracesBytes returns bytes -- single item", actual)
}

// ── newTraceCollection ──

func Test_Cov4_NewTraceCollection_Cap(t *testing.T) {
	tc := codestack.NewTraceCollection.Cap(10)
	actual := args.Map{"notNil": tc != nil, "isEmpty": tc.IsEmpty()}
	expected := args.Map{"notNil": true, "isEmpty": true}
	expected.ShouldBeEqual(t, 0, "NewTraceCollection.Cap returns empty collection -- cap 10", actual)
}

func Test_Cov4_NewTraceCollection_Empty(t *testing.T) {
	tc := codestack.NewTraceCollection.Empty()
	actual := args.Map{"notNil": tc != nil, "isEmpty": tc.IsEmpty()}
	expected := args.Map{"notNil": true, "isEmpty": true}
	expected.ShouldBeEqual(t, 0, "NewTraceCollection.Empty returns empty -- default", actual)
}

func Test_Cov4_NewTraceCollection_Using(t *testing.T) {
	traces := []codestack.Trace{codestack.New.Default()}
	tc := codestack.NewTraceCollection.Using(traces)
	actual := args.Map{"notNil": tc != nil, "len": tc.Length()}
	expected := args.Map{"notNil": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "NewTraceCollection.Using returns collection -- 1 item", actual)
}

// ── TraceCollection StackTraces/NewStackTraces ──

func Test_Cov4_TraceCollection_StackTraces(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.Add(codestack.New.Default())
	stJson := tc.StackTraces()
	actual := args.Map{"hasContent": len(stJson) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.StackTraces returns non-empty -- single item", actual)
}

func Test_Cov4_TraceCollection_NewStackTraces(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	newTc := tc.NewStackTraces(1, 5)
	actual := args.Map{"notNil": newTc != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.NewStackTraces returns non-nil -- skip 1", actual)
}

func Test_Cov4_TraceCollection_NewDefaultStackTraces(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	newTc := tc.NewDefaultStackTraces()
	actual := args.Map{"notNil": newTc != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.NewDefaultStackTraces returns non-nil -- default", actual)
}

func Test_Cov4_TraceCollection_AddsUsingSkipDefault(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.AddsUsingSkipDefault()
	actual := args.Map{"hasItems": tc.HasAnyItem()}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.AddsUsingSkipDefault populates -- default", actual)
}

func Test_Cov4_TraceCollection_ConcatNewUsingSkip(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	result := tc.ConcatNewUsingSkip(1, 3)
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.ConcatNewUsingSkip returns non-nil -- skip 1", actual)
}

func Test_Cov4_TraceCollection_GetPagesSize(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	for i := 0; i < 5; i++ {
		tc.Add(codestack.New.Default())
	}
	pages := tc.GetPagesSize(2)
	actual := args.Map{"pages": pages}
	expected := args.Map{"pages": 3}
	expected.ShouldBeEqual(t, 0, "TraceCollection.GetPagesSize returns 3 -- 5 items page 2", actual)
}

func Test_Cov4_TraceCollection_Serializer(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	tc.Add(codestack.New.Default())
	serializer := tc.Serializer()
	actual := args.Map{"notNil": serializer != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.Serializer returns non-nil -- single item", actual)
}

func Test_Cov4_TraceCollection_AsJsoner(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	jsoner := tc.AsJsoner()
	actual := args.Map{"notNil": jsoner != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.AsJsoner returns non-nil -- default", actual)
}

func Test_Cov4_TraceCollection_AsJsonContractsBinder(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	binder := tc.AsJsonContractsBinder()
	actual := args.Map{"notNil": binder != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.AsJsonContractsBinder returns non-nil -- default", actual)
}

func Test_Cov4_TraceCollection_AsJsonParseSelfInjector(t *testing.T) {
	tc := codestack.NewTraceCollection.Default()
	injector := tc.AsJsonParseSelfInjector()
	actual := args.Map{"notNil": injector != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.AsJsonParseSelfInjector returns non-nil -- default", actual)
}
