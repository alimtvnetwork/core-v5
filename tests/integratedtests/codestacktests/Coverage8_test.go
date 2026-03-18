package codestacktests

import (
	"testing"

	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── FileWithLine ──

func Test_Cov8_FileWithLine_FullFilePath(t *testing.T) {
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	actual := args.Map{"path": fwl.FullFilePath(), "line": fwl.LineNumber()}
	expected := args.Map{"path": "/tmp/test.go", "line": 10}
	expected.ShouldBeEqual(t, 0, "FileWithLine FullFilePath and LineNumber", actual)
}

func Test_Cov8_FileWithLine_IsNil(t *testing.T) {
	var fwl *codestack.FileWithLine
	actual := args.Map{"isNil": fwl.IsNil(), "isNotNil": fwl.IsNotNil()}
	expected := args.Map{"isNil": true, "isNotNil": false}
	expected.ShouldBeEqual(t, 0, "FileWithLine IsNil", actual)
}

func Test_Cov8_FileWithLine_IsNotNil(t *testing.T) {
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 1}
	actual := args.Map{"isNil": fwl.IsNil(), "isNotNil": fwl.IsNotNil()}
	expected := args.Map{"isNil": false, "isNotNil": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine IsNotNil", actual)
}

func Test_Cov8_FileWithLine_String(t *testing.T) {
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	actual := args.Map{"notEmpty": fwl.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine String", actual)
}

func Test_Cov8_FileWithLine_String_Nil(t *testing.T) {
	var fwl *codestack.FileWithLine
	actual := args.Map{"empty": fwl.String() == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine String nil", actual)
}

func Test_Cov8_FileWithLine_StringUsingFmt(t *testing.T) {
	fwl := codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	result := fwl.StringUsingFmt(func(f codestack.FileWithLine) string { return "custom" })
	actual := args.Map{"val": result}
	expected := args.Map{"val": "custom"}
	expected.ShouldBeEqual(t, 0, "FileWithLine StringUsingFmt", actual)
}

func Test_Cov8_FileWithLine_FileWithLine(t *testing.T) {
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	result := fwl.FileWithLine()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine FileWithLine method", actual)
}

func Test_Cov8_FileWithLine_JsonModel(t *testing.T) {
	fwl := codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	model := fwl.JsonModel()
	actual := args.Map{"path": model.FilePath}
	expected := args.Map{"path": "/tmp/test.go"}
	expected.ShouldBeEqual(t, 0, "FileWithLine JsonModel", actual)
}

func Test_Cov8_FileWithLine_JsonModelAny(t *testing.T) {
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	result := fwl.JsonModelAny()
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine JsonModelAny", actual)
}

func Test_Cov8_FileWithLine_JsonString(t *testing.T) {
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	result := fwl.JsonString()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine JsonString", actual)
}

func Test_Cov8_FileWithLine_Json(t *testing.T) {
	fwl := codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	result := fwl.Json()
	actual := args.Map{"noErr": result.MeaningfulError() == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine Json", actual)
}

func Test_Cov8_FileWithLine_JsonPtr(t *testing.T) {
	fwl := codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	result := fwl.JsonPtr()
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine JsonPtr", actual)
}

func Test_Cov8_FileWithLine_ParseInjectUsingJson(t *testing.T) {
	original := codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	jsonResult := original.JsonPtr()
	var target codestack.FileWithLine
	result, err := target.ParseInjectUsingJson(jsonResult)
	actual := args.Map{"noErr": err == nil, "path": result.FilePath}
	expected := args.Map{"noErr": true, "path": "/tmp/test.go"}
	expected.ShouldBeEqual(t, 0, "FileWithLine ParseInjectUsingJson", actual)
}

func Test_Cov8_FileWithLine_ParseInjectUsingJson_Error(t *testing.T) {
	badJson := corejson.NewPtr("not a FileWithLine struct")
	var target codestack.FileWithLine
	_, err := target.ParseInjectUsingJson(badJson)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine ParseInjectUsingJson error", actual)
}

func Test_Cov8_FileWithLine_ParseInjectUsingJsonMust(t *testing.T) {
	original := codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	jsonResult := original.JsonPtr()
	var target codestack.FileWithLine
	result := target.ParseInjectUsingJsonMust(jsonResult)
	actual := args.Map{"path": result.FilePath}
	expected := args.Map{"path": "/tmp/test.go"}
	expected.ShouldBeEqual(t, 0, "FileWithLine ParseInjectUsingJsonMust", actual)
}

func Test_Cov8_FileWithLine_JsonParseSelfInject(t *testing.T) {
	original := codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	jsonResult := original.JsonPtr()
	var target codestack.FileWithLine
	err := target.JsonParseSelfInject(jsonResult)
	actual := args.Map{"noErr": err == nil, "path": target.FilePath}
	expected := args.Map{"noErr": true, "path": "/tmp/test.go"}
	expected.ShouldBeEqual(t, 0, "FileWithLine JsonParseSelfInject", actual)
}

func Test_Cov8_FileWithLine_AsFileLiner(t *testing.T) {
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	liner := fwl.AsFileLiner()
	actual := args.Map{"path": liner.FullFilePath(), "line": liner.LineNumber()}
	expected := args.Map{"path": "/tmp/test.go", "line": 10}
	expected.ShouldBeEqual(t, 0, "FileWithLine AsFileLiner", actual)
}

// ── Trace ──

func Test_Cov8_Trace_New_Create(t *testing.T) {
	trace := codestack.New.Create(0)
	actual := args.Map{"isOkay": trace.IsOkay, "notNil": trace.IsNotNil()}
	expected := args.Map{"isOkay": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "Trace New.Create", actual)
}

func Test_Cov8_Trace_New_Default(t *testing.T) {
	trace := codestack.New.Default()
	actual := args.Map{"isOkay": trace.IsOkay}
	expected := args.Map{"isOkay": true}
	expected.ShouldBeEqual(t, 0, "Trace New.Default", actual)
}

func Test_Cov8_Trace_New_SkipOne(t *testing.T) {
	trace := codestack.New.SkipOne()
	actual := args.Map{"isOkay": trace.IsOkay}
	expected := args.Map{"isOkay": true}
	expected.ShouldBeEqual(t, 0, "Trace New.SkipOne", actual)
}

func Test_Cov8_Trace_New_Ptr(t *testing.T) {
	trace := codestack.New.Ptr(0)
	actual := args.Map{"notNil": trace != nil, "isOkay": trace.IsOkay}
	expected := args.Map{"notNil": true, "isOkay": true}
	expected.ShouldBeEqual(t, 0, "Trace New.Ptr", actual)
}

func Test_Cov8_Trace_Message(t *testing.T) {
	trace := codestack.New.Create(0)
	msg := trace.Message()
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Trace Message", actual)
}

func Test_Cov8_Trace_Message_Cached(t *testing.T) {
	trace := codestack.New.Create(0)
	msg1 := trace.Message()
	msg2 := trace.Message()
	actual := args.Map{"same": msg1 == msg2}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "Trace Message cached", actual)
}

func Test_Cov8_Trace_ShortString(t *testing.T) {
	trace := codestack.New.Create(0)
	ss := trace.ShortString()
	actual := args.Map{"notEmpty": ss != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Trace ShortString", actual)
}

func Test_Cov8_Trace_ShortString_Cached(t *testing.T) {
	trace := codestack.New.Create(0)
	ss1 := trace.ShortString()
	ss2 := trace.ShortString()
	actual := args.Map{"same": ss1 == ss2}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "Trace ShortString cached", actual)
}

func Test_Cov8_Trace_IsNil(t *testing.T) {
	var trace *codestack.Trace
	actual := args.Map{"isNil": trace.IsNil(), "isNotNil": trace.IsNotNil()}
	expected := args.Map{"isNil": true, "isNotNil": false}
	expected.ShouldBeEqual(t, 0, "Trace IsNil", actual)
}

func Test_Cov8_Trace_HasIssues_Nil(t *testing.T) {
	var trace *codestack.Trace
	actual := args.Map{"hasIssues": trace.HasIssues()}
	expected := args.Map{"hasIssues": true}
	expected.ShouldBeEqual(t, 0, "Trace HasIssues nil", actual)
}

func Test_Cov8_Trace_HasIssues_NotOkay(t *testing.T) {
	trace := &codestack.Trace{IsOkay: false}
	actual := args.Map{"hasIssues": trace.HasIssues()}
	expected := args.Map{"hasIssues": true}
	expected.ShouldBeEqual(t, 0, "Trace HasIssues not okay", actual)
}

func Test_Cov8_Trace_HasIssues_MissingPackage(t *testing.T) {
	trace := &codestack.Trace{IsOkay: true, PackageMethodName: "test", PackageName: ""}
	actual := args.Map{"hasIssues": trace.HasIssues()}
	expected := args.Map{"hasIssues": true}
	expected.ShouldBeEqual(t, 0, "Trace HasIssues missing package", actual)
}

func Test_Cov8_Trace_String(t *testing.T) {
	trace := codestack.New.Create(0)
	actual := args.Map{"notEmpty": trace.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Trace String", actual)
}

func Test_Cov8_Trace_String_Nil(t *testing.T) {
	var trace *codestack.Trace
	actual := args.Map{"empty": trace.String() == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Trace String nil", actual)
}

func Test_Cov8_Trace_StringUsingFmt(t *testing.T) {
	trace := codestack.New.Create(0)
	result := trace.StringUsingFmt(func(tr codestack.Trace) string { return "custom" })
	actual := args.Map{"val": result}
	expected := args.Map{"val": "custom"}
	expected.ShouldBeEqual(t, 0, "Trace StringUsingFmt", actual)
}

func Test_Cov8_Trace_FileWithLine(t *testing.T) {
	trace := codestack.New.Create(0)
	fwl := trace.FileWithLine()
	actual := args.Map{"pathNotEmpty": fwl.FilePath != "", "lineGt0": fwl.Line > 0}
	expected := args.Map{"pathNotEmpty": true, "lineGt0": true}
	expected.ShouldBeEqual(t, 0, "Trace FileWithLine", actual)
}

func Test_Cov8_Trace_FullFilePath(t *testing.T) {
	trace := codestack.New.Create(0)
	actual := args.Map{"notEmpty": trace.FullFilePath() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Trace FullFilePath", actual)
}

func Test_Cov8_Trace_FileName(t *testing.T) {
	trace := codestack.New.Create(0)
	actual := args.Map{"notEmpty": trace.FileName() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Trace FileName", actual)
}

func Test_Cov8_Trace_LineNumber(t *testing.T) {
	trace := codestack.New.Create(0)
	actual := args.Map{"gt0": trace.LineNumber() > 0}
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "Trace LineNumber", actual)
}

func Test_Cov8_Trace_FileWithLineString(t *testing.T) {
	trace := codestack.New.Create(0)
	result := trace.FileWithLineString()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Trace FileWithLineString", actual)
}

func Test_Cov8_Trace_JsonModel(t *testing.T) {
	trace := codestack.New.Create(0)
	model := trace.JsonModel()
	actual := args.Map{"isOkay": model.IsOkay}
	expected := args.Map{"isOkay": true}
	expected.ShouldBeEqual(t, 0, "Trace JsonModel", actual)
}

func Test_Cov8_Trace_JsonModelAny(t *testing.T) {
	trace := codestack.New.Create(0)
	result := trace.JsonModelAny()
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Trace JsonModelAny", actual)
}

func Test_Cov8_Trace_JsonString(t *testing.T) {
	trace := codestack.New.Create(0)
	result := trace.JsonString()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Trace JsonString", actual)
}

func Test_Cov8_Trace_Json(t *testing.T) {
	trace := codestack.New.Create(0)
	result := trace.Json()
	actual := args.Map{"noErr": result.MeaningfulError() == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Trace Json", actual)
}

func Test_Cov8_Trace_JsonPtr(t *testing.T) {
	trace := codestack.New.Create(0)
	result := trace.JsonPtr()
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Trace JsonPtr", actual)
}

func Test_Cov8_Trace_Clone(t *testing.T) {
	trace := codestack.New.Create(0)
	cloned := trace.Clone()
	actual := args.Map{"samePath": cloned.FilePath == trace.FilePath}
	expected := args.Map{"samePath": true}
	expected.ShouldBeEqual(t, 0, "Trace Clone", actual)
}

func Test_Cov8_Trace_ClonePtr(t *testing.T) {
	trace := codestack.New.Create(0)
	cloned := trace.ClonePtr()
	actual := args.Map{"notNil": cloned != nil, "samePath": cloned.FilePath == trace.FilePath}
	expected := args.Map{"notNil": true, "samePath": true}
	expected.ShouldBeEqual(t, 0, "Trace ClonePtr", actual)
}

func Test_Cov8_Trace_ClonePtr_Nil(t *testing.T) {
	var trace *codestack.Trace
	cloned := trace.ClonePtr()
	isNil := cloned == nil
	actual := args.Map{"isNil": isNil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Trace ClonePtr nil", actual)
}

func Test_Cov8_Trace_ParseInjectUsingJson(t *testing.T) {
	trace := codestack.New.Create(0)
	jsonResult := trace.JsonPtr()
	var target codestack.Trace
	result, err := target.ParseInjectUsingJson(jsonResult)
	actual := args.Map{"noErr": err == nil, "notNil": result != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "Trace ParseInjectUsingJson", actual)
}

func Test_Cov8_Trace_ParseInjectUsingJsonMust(t *testing.T) {
	trace := codestack.New.Create(0)
	jsonResult := trace.JsonPtr()
	var target codestack.Trace
	result := target.ParseInjectUsingJsonMust(jsonResult)
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Trace ParseInjectUsingJsonMust", actual)
}

func Test_Cov8_Trace_JsonParseSelfInject(t *testing.T) {
	trace := codestack.New.Create(0)
	jsonResult := trace.JsonPtr()
	var target codestack.Trace
	err := target.JsonParseSelfInject(jsonResult)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Trace JsonParseSelfInject", actual)
}

func Test_Cov8_Trace_Dispose(t *testing.T) {
	trace := codestack.New.Create(0)
	trace.Dispose()
	actual := args.Map{"pathEmpty": trace.FilePath == "", "line": trace.Line}
	expected := args.Map{"pathEmpty": true, "line": 0}
	expected.ShouldBeEqual(t, 0, "Trace Dispose", actual)
}

func Test_Cov8_Trace_Dispose_Nil(t *testing.T) {
	var trace *codestack.Trace
	trace.Dispose() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Trace Dispose nil", actual)
}

func Test_Cov8_Trace_AsFileLiner(t *testing.T) {
	trace := codestack.New.Create(0)
	liner := trace.AsFileLiner()
	actual := args.Map{"notEmpty": liner.FullFilePath() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Trace AsFileLiner", actual)
}

// ── TraceCollection ──

func Test_Cov8_TraceCollection_Basic(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	actual := args.Map{"hasAny": tc.HasAnyItem(), "lenGt0": tc.Length() > 0, "count": tc.Count()}
	expected := args.Map{"hasAny": true, "lenGt0": true, "count": tc.Length()}
	expected.ShouldBeEqual(t, 0, "TraceCollection basic", actual)
}

func Test_Cov8_TraceCollection_IsEmpty(t *testing.T) {
	tc := &codestack.TraceCollection{}
	actual := args.Map{"isEmpty": tc.IsEmpty(), "lastIndex": tc.LastIndex()}
	expected := args.Map{"isEmpty": true, "lastIndex": -1}
	expected.ShouldBeEqual(t, 0, "TraceCollection IsEmpty", actual)
}

func Test_Cov8_TraceCollection_Length_Nil(t *testing.T) {
	var tc *codestack.TraceCollection
	actual := args.Map{"len": tc.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TraceCollection Length nil", actual)
}

func Test_Cov8_TraceCollection_FirstAndLast(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	if tc.Length() > 0 {
		first := tc.First()
		last := tc.Last()
		actual := args.Map{"firstOk": first.IsOkay, "lastOk": last.IsOkay}
		expected := args.Map{"firstOk": true, "lastOk": true}
		expected.ShouldBeEqual(t, 0, "TraceCollection First and Last", actual)
	}
}

func Test_Cov8_TraceCollection_FirstOrDefault_Empty(t *testing.T) {
	tc := &codestack.TraceCollection{}
	result := tc.FirstOrDefault()
	isNil := result == nil
	actual := args.Map{"isNil": isNil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection FirstOrDefault empty", actual)
}

func Test_Cov8_TraceCollection_LastOrDefault_Empty(t *testing.T) {
	tc := &codestack.TraceCollection{}
	result := tc.LastOrDefault()
	isNil := result == nil
	actual := args.Map{"isNil": isNil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection LastOrDefault empty", actual)
}

func Test_Cov8_TraceCollection_FirstOrDefault_NonEmpty(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.FirstOrDefault()
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection FirstOrDefault non-empty", actual)
}

func Test_Cov8_TraceCollection_FirstDynamic(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.FirstDynamic()
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection FirstDynamic", actual)
}

func Test_Cov8_TraceCollection_LastDynamic(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.LastDynamic()
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection LastDynamic", actual)
}

func Test_Cov8_TraceCollection_FirstOrDefaultDynamic(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.FirstOrDefaultDynamic()
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection FirstOrDefaultDynamic", actual)
}

func Test_Cov8_TraceCollection_LastOrDefaultDynamic(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.LastOrDefaultDynamic()
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection LastOrDefaultDynamic", actual)
}

func Test_Cov8_TraceCollection_HasIndex(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	actual := args.Map{"has0": tc.HasIndex(0), "hasBig": tc.HasIndex(99999)}
	expected := args.Map{"has0": true, "hasBig": false}
	expected.ShouldBeEqual(t, 0, "TraceCollection HasIndex", actual)
}

func Test_Cov8_TraceCollection_Strings(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	strs := tc.Strings()
	actual := args.Map{"lenGt0": len(strs) > 0}
	expected := args.Map{"lenGt0": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection Strings", actual)
}

func Test_Cov8_TraceCollection_ShortStrings(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	strs := tc.ShortStrings()
	actual := args.Map{"lenGt0": len(strs) > 0}
	expected := args.Map{"lenGt0": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection ShortStrings", actual)
}

func Test_Cov8_TraceCollection_FileWithLines(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	fwls := tc.FileWithLines()
	actual := args.Map{"lenGt0": len(fwls) > 0}
	expected := args.Map{"lenGt0": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection FileWithLines", actual)
}

func Test_Cov8_TraceCollection_FileWithLinesStrings(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	strs := tc.FileWithLinesStrings()
	actual := args.Map{"lenGt0": len(strs) > 0}
	expected := args.Map{"lenGt0": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection FileWithLinesStrings", actual)
}

func Test_Cov8_TraceCollection_CodeStacksString(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.CodeStacksString()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection CodeStacksString", actual)
}

func Test_Cov8_TraceCollection_CodeStacksString_Empty(t *testing.T) {
	tc := codestack.TraceCollection{}
	result := tc.CodeStacksString()
	actual := args.Map{"empty": result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection CodeStacksString empty", actual)
}

func Test_Cov8_TraceCollection_StackTraces(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.StackTraces()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection StackTraces", actual)
}

func Test_Cov8_TraceCollection_StackTracesJsonResult(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.StackTracesJsonResult()
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection StackTracesJsonResult", actual)
}

func Test_Cov8_TraceCollection_StackTracesBytes(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.StackTracesBytes()
	actual := args.Map{"lenGt0": len(result) > 0}
	expected := args.Map{"lenGt0": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection StackTracesBytes", actual)
}

func Test_Cov8_TraceCollection_StackTracesBytes_Empty(t *testing.T) {
	tc := codestack.TraceCollection{}
	result := tc.StackTracesBytes()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TraceCollection StackTracesBytes empty", actual)
}

func Test_Cov8_TraceCollection_Clone(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	cloned := tc.Clone()
	actual := args.Map{"sameLen": cloned.Length() == tc.Length()}
	expected := args.Map{"sameLen": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection Clone", actual)
}

func Test_Cov8_TraceCollection_ClonePtr(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	cloned := tc.ClonePtr()
	actual := args.Map{"notNil": cloned != nil, "sameLen": cloned.Length() == tc.Length()}
	expected := args.Map{"notNil": true, "sameLen": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection ClonePtr", actual)
}

func Test_Cov8_TraceCollection_ClonePtr_Nil(t *testing.T) {
	var tc *codestack.TraceCollection
	cloned := tc.ClonePtr()
	isNil := cloned == nil
	actual := args.Map{"isNil": isNil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection ClonePtr nil", actual)
}

func Test_Cov8_TraceCollection_Reverse_Empty(t *testing.T) {
	tc := &codestack.TraceCollection{}
	result := tc.Reverse()
	actual := args.Map{"empty": result.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection Reverse empty", actual)
}

func Test_Cov8_TraceCollection_Reverse_Single(t *testing.T) {
	trace := codestack.New.Create(0)
	tc := &codestack.TraceCollection{Items: []codestack.Trace{trace}}
	result := tc.Reverse()
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection Reverse single", actual)
}

func Test_Cov8_TraceCollection_Reverse_Two(t *testing.T) {
	t1 := codestack.Trace{PackageName: "a"}
	t2 := codestack.Trace{PackageName: "b"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2}}
	result := tc.Reverse()
	actual := args.Map{"first": result.Items[0].PackageName, "second": result.Items[1].PackageName}
	expected := args.Map{"first": "b", "second": "a"}
	expected.ShouldBeEqual(t, 0, "TraceCollection Reverse two", actual)
}

func Test_Cov8_TraceCollection_Reverse_Three(t *testing.T) {
	t1 := codestack.Trace{PackageName: "a"}
	t2 := codestack.Trace{PackageName: "b"}
	t3 := codestack.Trace{PackageName: "c"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2, t3}}
	result := tc.Reverse()
	actual := args.Map{"first": result.Items[0].PackageName, "last": result.Items[2].PackageName}
	expected := args.Map{"first": "c", "last": "a"}
	expected.ShouldBeEqual(t, 0, "TraceCollection Reverse three", actual)
}

func Test_Cov8_TraceCollection_Skip(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	if tc.Length() > 1 {
		skipped := tc.Skip(1)
		actual := args.Map{"lenLess": len(skipped) == tc.Length()-1}
		expected := args.Map{"lenLess": true}
		expected.ShouldBeEqual(t, 0, "TraceCollection Skip", actual)
	}
}

func Test_Cov8_TraceCollection_SkipDynamic_BeyondLength(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.SkipDynamic(99999)
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection SkipDynamic beyond", actual)
}

func Test_Cov8_TraceCollection_SkipCollection(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	if tc.Length() > 1 {
		result := tc.SkipCollection(1)
		actual := args.Map{"lenLess": result.Length() == tc.Length()-1}
		expected := args.Map{"lenLess": true}
		expected.ShouldBeEqual(t, 0, "TraceCollection SkipCollection", actual)
	}
}

func Test_Cov8_TraceCollection_Take(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	if tc.Length() > 1 {
		taken := tc.Take(1)
		actual := args.Map{"len": len(taken)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "TraceCollection Take", actual)
	}
}

func Test_Cov8_TraceCollection_TakeDynamic(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	if tc.Length() > 0 {
		result := tc.TakeDynamic(1)
		actual := args.Map{"notNil": result != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "TraceCollection TakeDynamic", actual)
	}
}

func Test_Cov8_TraceCollection_TakeCollection(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	if tc.Length() > 0 {
		result := tc.TakeCollection(1)
		actual := args.Map{"len": result.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "TraceCollection TakeCollection", actual)
	}
}

func Test_Cov8_TraceCollection_LimitCollection(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	if tc.Length() > 0 {
		result := tc.LimitCollection(1)
		actual := args.Map{"len": result.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "TraceCollection LimitCollection", actual)
	}
}

func Test_Cov8_TraceCollection_SafeLimitCollection(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.SafeLimitCollection(999)
	actual := args.Map{"sameLen": result.Length() == tc.Length()}
	expected := args.Map{"sameLen": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection SafeLimitCollection", actual)
}

func Test_Cov8_TraceCollection_LimitDynamic(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	if tc.Length() > 0 {
		result := tc.LimitDynamic(1)
		actual := args.Map{"notNil": result != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "TraceCollection LimitDynamic", actual)
	}
}

func Test_Cov8_TraceCollection_Limit(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	if tc.Length() > 0 {
		result := tc.Limit(1)
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "TraceCollection Limit", actual)
	}
}

func Test_Cov8_TraceCollection_GetPagesSize(t *testing.T) {
	t1 := codestack.Trace{PackageName: "a"}
	t2 := codestack.Trace{PackageName: "b"}
	t3 := codestack.Trace{PackageName: "c"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2, t3}}
	actual := args.Map{"pages": tc.GetPagesSize(2), "zero": tc.GetPagesSize(0)}
	expected := args.Map{"pages": 2, "zero": 0}
	expected.ShouldBeEqual(t, 0, "TraceCollection GetPagesSize", actual)
}

func Test_Cov8_TraceCollection_GetPagedCollection(t *testing.T) {
	t1 := codestack.Trace{PackageName: "a"}
	t2 := codestack.Trace{PackageName: "b"}
	t3 := codestack.Trace{PackageName: "c"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2, t3}}
	pages := tc.GetPagedCollection(2)
	actual := args.Map{"pageCount": len(pages)}
	expected := args.Map{"pageCount": 2}
	expected.ShouldBeEqual(t, 0, "TraceCollection GetPagedCollection", actual)
}

func Test_Cov8_TraceCollection_GetPagedCollection_SmallPage(t *testing.T) {
	t1 := codestack.Trace{PackageName: "a"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1}}
	pages := tc.GetPagedCollection(5)
	actual := args.Map{"pageCount": len(pages)}
	expected := args.Map{"pageCount": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection GetPagedCollection small", actual)
}

func Test_Cov8_TraceCollection_GetSinglePageCollection(t *testing.T) {
	t1 := codestack.Trace{PackageName: "a"}
	t2 := codestack.Trace{PackageName: "b"}
	t3 := codestack.Trace{PackageName: "c"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2, t3}}
	page := tc.GetSinglePageCollection(2, 1)
	actual := args.Map{"len": page.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "TraceCollection GetSinglePageCollection", actual)
}

func Test_Cov8_TraceCollection_Filter(t *testing.T) {
	t1 := codestack.Trace{PackageName: "a", IsOkay: true, PackageMethodName: "a.test"}
	t2 := codestack.Trace{PackageName: "b", IsOkay: true, PackageMethodName: "b.test"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2}}
	result := tc.Filter(func(trace *codestack.Trace) (bool, bool) {
		return trace.PackageName == "a", false
	})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection Filter", actual)
}

func Test_Cov8_TraceCollection_Filter_WithBreak(t *testing.T) {
	t1 := codestack.Trace{PackageName: "a"}
	t2 := codestack.Trace{PackageName: "b"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2}}
	result := tc.Filter(func(trace *codestack.Trace) (bool, bool) {
		return true, trace.PackageName == "a"
	})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection Filter with break", actual)
}

func Test_Cov8_TraceCollection_FilterTraceCollection(t *testing.T) {
	t1 := codestack.Trace{PackageName: "a"}
	t2 := codestack.Trace{PackageName: "b"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2}}
	result := tc.FilterTraceCollection(func(trace *codestack.Trace) (bool, bool) {
		return trace.PackageName == "a", false
	})
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection FilterTraceCollection", actual)
}

func Test_Cov8_TraceCollection_FilterPackageNameTraceCollection(t *testing.T) {
	t1 := codestack.Trace{PackageName: "a"}
	t2 := codestack.Trace{PackageName: "b"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2}}
	result := tc.FilterPackageNameTraceCollection("a")
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection FilterPackageName", actual)
}

func Test_Cov8_TraceCollection_SkipFilterPackageNameTraceCollection(t *testing.T) {
	t1 := codestack.Trace{PackageName: "a"}
	t2 := codestack.Trace{PackageName: "b"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2}}
	result := tc.SkipFilterPackageNameTraceCollection("a")
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection SkipFilterPackageName", actual)
}

func Test_Cov8_TraceCollection_FilterMethodNameTraceCollection(t *testing.T) {
	t1 := codestack.Trace{MethodName: "test"}
	t2 := codestack.Trace{MethodName: "other"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2}}
	result := tc.FilterMethodNameTraceCollection("test")
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection FilterMethodName", actual)
}

func Test_Cov8_TraceCollection_SkipFilterMethodNameTraceCollection(t *testing.T) {
	t1 := codestack.Trace{MethodName: "test"}
	t2 := codestack.Trace{MethodName: "other"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2}}
	result := tc.SkipFilterMethodNameTraceCollection("test")
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection SkipFilterMethodName", actual)
}

func Test_Cov8_TraceCollection_FilterFullMethodNameTraceCollection(t *testing.T) {
	t1 := codestack.Trace{PackageMethodName: "pkg.Method"}
	t2 := codestack.Trace{PackageMethodName: "pkg.Other"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2}}
	result := tc.FilterFullMethodNameTraceCollection("pkg.Method")
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection FilterFullMethodName", actual)
}

func Test_Cov8_TraceCollection_SkipFilterFullMethodNameTraceCollection(t *testing.T) {
	t1 := codestack.Trace{PackageMethodName: "pkg.Method"}
	t2 := codestack.Trace{PackageMethodName: "pkg.Other"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2}}
	result := tc.SkipFilterFullMethodNameTraceCollection("pkg.Method")
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection SkipFilterFullMethodName", actual)
}

func Test_Cov8_TraceCollection_SkipFilterFilenameTraceCollection(t *testing.T) {
	t1 := codestack.Trace{FilePath: "/a.go"}
	t2 := codestack.Trace{FilePath: "/b.go"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2}}
	result := tc.SkipFilterFilenameTraceCollection("/a.go")
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection SkipFilterFilename", actual)
}

func Test_Cov8_TraceCollection_FilterWithLimit(t *testing.T) {
	t1 := codestack.Trace{PackageName: "a"}
	t2 := codestack.Trace{PackageName: "b"}
	t3 := codestack.Trace{PackageName: "c"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2, t3}}
	result := tc.FilterWithLimit(2, func(trace *codestack.Trace) (bool, bool) {
		return true, false
	})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "TraceCollection FilterWithLimit", actual)
}

func Test_Cov8_TraceCollection_Json(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.Json()
	actual := args.Map{"noErr": result.MeaningfulError() == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection Json", actual)
}

func Test_Cov8_TraceCollection_JsonPtr(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.JsonPtr()
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection JsonPtr", actual)
}

func Test_Cov8_TraceCollection_JsonString(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.JsonString()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection JsonString", actual)
}

func Test_Cov8_TraceCollection_JsonString_Empty(t *testing.T) {
	tc := codestack.TraceCollection{}
	result := tc.JsonString()
	actual := args.Map{"empty": result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection JsonString empty", actual)
}

func Test_Cov8_TraceCollection_String(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.String()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection String", actual)
}

func Test_Cov8_TraceCollection_String_Empty(t *testing.T) {
	tc := codestack.TraceCollection{}
	result := tc.String()
	actual := args.Map{"empty": result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection String empty", actual)
}

func Test_Cov8_TraceCollection_JsonModel(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	model := tc.JsonModel()
	actual := args.Map{"lenGt0": len(model) > 0}
	expected := args.Map{"lenGt0": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection JsonModel", actual)
}

func Test_Cov8_TraceCollection_JsonModelAny(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.JsonModelAny()
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection JsonModelAny", actual)
}

func Test_Cov8_TraceCollection_Serializer(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	bytes, err := tc.Serializer()
	actual := args.Map{"noErr": err == nil, "lenGt0": len(bytes) > 0}
	expected := args.Map{"noErr": true, "lenGt0": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection Serializer", actual)
}

func Test_Cov8_TraceCollection_Join(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.Join(", ")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection Join", actual)
}

func Test_Cov8_TraceCollection_JoinLines(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.JoinLines()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection JoinLines", actual)
}

func Test_Cov8_TraceCollection_JoinShortStrings(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.JoinShortStrings(", ")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection JoinShortStrings", actual)
}

func Test_Cov8_TraceCollection_JoinFileWithLinesStrings(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.JoinFileWithLinesStrings(", ")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection JoinFileWithLinesStrings", actual)
}

func Test_Cov8_TraceCollection_JoinJsonStrings(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.JoinJsonStrings(", ")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection JoinJsonStrings", actual)
}

func Test_Cov8_TraceCollection_JsonStrings(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.JsonStrings()
	actual := args.Map{"lenGt0": len(result) > 0}
	expected := args.Map{"lenGt0": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection JsonStrings", actual)
}

func Test_Cov8_TraceCollection_CsvStrings(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.CsvStrings()
	actual := args.Map{"lenGt0": len(result) > 0}
	expected := args.Map{"lenGt0": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection CsvStrings", actual)
}

func Test_Cov8_TraceCollection_CsvStrings_Empty(t *testing.T) {
	tc := codestack.TraceCollection{}
	result := tc.CsvStrings()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TraceCollection CsvStrings empty", actual)
}

func Test_Cov8_TraceCollection_JoinCsv(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.JoinCsv()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection JoinCsv", actual)
}

func Test_Cov8_TraceCollection_JoinCsvLine(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.JoinCsvLine()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection JoinCsvLine", actual)
}

func Test_Cov8_TraceCollection_JoinUsingFmt(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.JoinUsingFmt(func(trace *codestack.Trace) string { return "x" }, ", ")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection JoinUsingFmt", actual)
}

func Test_Cov8_TraceCollection_StringsUsingFmt(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.StringsUsingFmt(func(trace *codestack.Trace) string { return "x" })
	actual := args.Map{"lenGt0": len(result) > 0, "first": result[0]}
	expected := args.Map{"lenGt0": true, "first": "x"}
	expected.ShouldBeEqual(t, 0, "TraceCollection StringsUsingFmt", actual)
}

func Test_Cov8_TraceCollection_CodeStacksStringLimit(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.CodeStacksStringLimit(1)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection CodeStacksStringLimit", actual)
}

func Test_Cov8_TraceCollection_CodeStacksStringLimit_Empty(t *testing.T) {
	tc := codestack.TraceCollection{}
	result := tc.CodeStacksStringLimit(1)
	actual := args.Map{"empty": result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection CodeStacksStringLimit empty", actual)
}

func Test_Cov8_TraceCollection_FileWithLinesString(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.FileWithLinesString()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection FileWithLinesString", actual)
}

func Test_Cov8_TraceCollection_FileWithLinesString_Empty(t *testing.T) {
	tc := &codestack.TraceCollection{}
	result := tc.FileWithLinesString()
	actual := args.Map{"empty": result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection FileWithLinesString empty", actual)
}

func Test_Cov8_TraceCollection_Add(t *testing.T) {
	tc := &codestack.TraceCollection{}
	trace := codestack.Trace{PackageName: "test"}
	tc.Add(trace)
	actual := args.Map{"len": tc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection Add", actual)
}

func Test_Cov8_TraceCollection_Adds(t *testing.T) {
	tc := &codestack.TraceCollection{}
	tc.Adds(codestack.Trace{PackageName: "a"}, codestack.Trace{PackageName: "b"})
	actual := args.Map{"len": tc.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "TraceCollection Adds", actual)
}

func Test_Cov8_TraceCollection_Adds_Empty(t *testing.T) {
	tc := &codestack.TraceCollection{}
	tc.Adds()
	actual := args.Map{"len": tc.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TraceCollection Adds empty", actual)
}

func Test_Cov8_TraceCollection_AddsIf_True(t *testing.T) {
	tc := &codestack.TraceCollection{}
	tc.AddsIf(true, codestack.Trace{PackageName: "a"})
	actual := args.Map{"len": tc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection AddsIf true", actual)
}

func Test_Cov8_TraceCollection_AddsIf_False(t *testing.T) {
	tc := &codestack.TraceCollection{}
	tc.AddsIf(false, codestack.Trace{PackageName: "a"})
	actual := args.Map{"len": tc.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TraceCollection AddsIf false", actual)
}

func Test_Cov8_TraceCollection_AddsPtr(t *testing.T) {
	tc := &codestack.TraceCollection{}
	trace := &codestack.Trace{PackageName: "a", IsOkay: true, PackageMethodName: "a.test"}
	tc.AddsPtr(false, trace)
	actual := args.Map{"len": tc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection AddsPtr", actual)
}

func Test_Cov8_TraceCollection_AddsPtr_SkipIssues(t *testing.T) {
	tc := &codestack.TraceCollection{}
	trace := &codestack.Trace{IsOkay: false}
	tc.AddsPtr(true, trace)
	actual := args.Map{"len": tc.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TraceCollection AddsPtr skip issues", actual)
}

func Test_Cov8_TraceCollection_AddsPtr_Nil(t *testing.T) {
	tc := &codestack.TraceCollection{}
	tc.AddsPtr(false, nil)
	actual := args.Map{"len": tc.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TraceCollection AddsPtr nil", actual)
}

func Test_Cov8_TraceCollection_AddsPtr_Empty(t *testing.T) {
	tc := &codestack.TraceCollection{}
	tc.AddsPtr(false)
	actual := args.Map{"len": tc.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TraceCollection AddsPtr empty", actual)
}

func Test_Cov8_TraceCollection_ConcatNew(t *testing.T) {
	tc := &codestack.TraceCollection{Items: []codestack.Trace{{PackageName: "a"}}}
	result := tc.ConcatNew(codestack.Trace{PackageName: "b"})
	actual := args.Map{"len": result.Length(), "origLen": tc.Length()}
	expected := args.Map{"len": 2, "origLen": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection ConcatNew", actual)
}

func Test_Cov8_TraceCollection_ConcatNewPtr(t *testing.T) {
	tc := &codestack.TraceCollection{Items: []codestack.Trace{{PackageName: "a"}}}
	trace := &codestack.Trace{PackageName: "b", IsOkay: true, PackageMethodName: "b.test"}
	result := tc.ConcatNewPtr(trace)
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "TraceCollection ConcatNewPtr", actual)
}

func Test_Cov8_TraceCollection_IsEqual(t *testing.T) {
	t1 := codestack.Trace{PackageName: "a"}
	tc1 := &codestack.TraceCollection{Items: []codestack.Trace{t1}}
	tc2 := &codestack.TraceCollection{Items: []codestack.Trace{t1}}
	actual := args.Map{"equal": tc1.IsEqual(tc2)}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection IsEqual", actual)
}

func Test_Cov8_TraceCollection_IsEqual_BothNil(t *testing.T) {
	var tc1 *codestack.TraceCollection
	var tc2 *codestack.TraceCollection
	actual := args.Map{"equal": tc1.IsEqual(tc2)}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection IsEqual both nil", actual)
}

func Test_Cov8_TraceCollection_IsEqual_OneNil(t *testing.T) {
	tc1 := &codestack.TraceCollection{}
	var tc2 *codestack.TraceCollection
	actual := args.Map{"equal": tc1.IsEqual(tc2)}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "TraceCollection IsEqual one nil", actual)
}

func Test_Cov8_TraceCollection_IsEqual_DiffLen(t *testing.T) {
	tc1 := &codestack.TraceCollection{Items: []codestack.Trace{{PackageName: "a"}}}
	tc2 := &codestack.TraceCollection{Items: []codestack.Trace{{PackageName: "a"}, {PackageName: "b"}}}
	actual := args.Map{"equal": tc1.IsEqual(tc2)}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "TraceCollection IsEqual diff len", actual)
}

func Test_Cov8_TraceCollection_IsEqualItems(t *testing.T) {
	t1 := codestack.Trace{PackageName: "a"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1}}
	actual := args.Map{"equal": tc.IsEqualItems(t1)}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection IsEqualItems", actual)
}

func Test_Cov8_TraceCollection_IsEqualItems_DiffLen(t *testing.T) {
	tc := &codestack.TraceCollection{Items: []codestack.Trace{{PackageName: "a"}}}
	actual := args.Map{"equal": tc.IsEqualItems(codestack.Trace{PackageName: "a"}, codestack.Trace{PackageName: "b"})}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "TraceCollection IsEqualItems diff len", actual)
}

func Test_Cov8_TraceCollection_IsEqualItems_NotEqual(t *testing.T) {
	tc := &codestack.TraceCollection{Items: []codestack.Trace{{PackageName: "a"}}}
	actual := args.Map{"equal": tc.IsEqualItems(codestack.Trace{PackageName: "b"})}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "TraceCollection IsEqualItems not equal", actual)
}

func Test_Cov8_TraceCollection_ParseInjectUsingJson(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	jsonResult := tc.JsonPtr()
	var target codestack.TraceCollection
	result, err := target.ParseInjectUsingJson(jsonResult)
	actual := args.Map{"noErr": err == nil, "lenGt0": result.Length() > 0}
	expected := args.Map{"noErr": true, "lenGt0": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection ParseInjectUsingJson", actual)
}

func Test_Cov8_TraceCollection_ParseInjectUsingJsonMust(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	jsonResult := tc.JsonPtr()
	var target codestack.TraceCollection
	result := target.ParseInjectUsingJsonMust(jsonResult)
	actual := args.Map{"lenGt0": result.Length() > 0}
	expected := args.Map{"lenGt0": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection ParseInjectUsingJsonMust", actual)
}

func Test_Cov8_TraceCollection_JsonParseSelfInject(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	jsonResult := tc.JsonPtr()
	var target codestack.TraceCollection
	err := target.JsonParseSelfInject(jsonResult)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection JsonParseSelfInject", actual)
}

func Test_Cov8_TraceCollection_AsJsonContractsBinder(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.AsJsonContractsBinder()
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection AsJsonContractsBinder", actual)
}

func Test_Cov8_TraceCollection_AsJsoner(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.AsJsoner()
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection AsJsoner", actual)
}

func Test_Cov8_TraceCollection_AsJsonParseSelfInjector(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.AsJsonParseSelfInjector()
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection AsJsonParseSelfInjector", actual)
}

func Test_Cov8_TraceCollection_Clear(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	tc.Clear()
	actual := args.Map{"empty": tc.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection Clear", actual)
}

func Test_Cov8_TraceCollection_Clear_Nil(t *testing.T) {
	var tc *codestack.TraceCollection
	result := tc.Clear()
	isNil := result == nil
	actual := args.Map{"isNil": isNil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection Clear nil", actual)
}

func Test_Cov8_TraceCollection_Dispose(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	tc.Dispose()
	isNilItems := tc.Items == nil
	actual := args.Map{"nilItems": isNilItems}
	expected := args.Map{"nilItems": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection Dispose", actual)
}

func Test_Cov8_TraceCollection_Dispose_Nil(t *testing.T) {
	var tc *codestack.TraceCollection
	tc.Dispose() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection Dispose nil", actual)
}

func Test_Cov8_TraceCollection_InsertAt(t *testing.T) {
	t1 := codestack.Trace{PackageName: "a"}
	t2 := codestack.Trace{PackageName: "b"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2}}
	inserted := codestack.Trace{PackageName: "x"}
	tc.InsertAt(0, inserted)
	actual := args.Map{"first": tc.Items[0].PackageName, "len": tc.Length()}
	expected := args.Map{"first": "x", "len": 3}
	expected.ShouldBeEqual(t, 0, "TraceCollection InsertAt", actual)
}

// ── NameOf ──

func Test_Cov8_NameOf_All_Empty(t *testing.T) {
	full, pkg, method := codestack.NameOf.All("")
	actual := args.Map{"full": full, "pkg": pkg, "method": method}
	expected := args.Map{"full": "", "pkg": "", "method": ""}
	expected.ShouldBeEqual(t, 0, "NameOf.All empty", actual)
}

func Test_Cov8_NameOf_Method(t *testing.T) {
	method := codestack.NameOf.Method()
	actual := args.Map{"notEmpty": method != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameOf.Method", actual)
}

func Test_Cov8_NameOf_Package(t *testing.T) {
	pkg := codestack.NameOf.Package()
	actual := args.Map{"notEmpty": pkg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameOf.Package", actual)
}

// ── StacksTo ──

func Test_Cov8_StacksTo_Bytes(t *testing.T) {
	result := codestack.StacksTo.Bytes(0)
	actual := args.Map{"lenGt0": len(result) > 0}
	expected := args.Map{"lenGt0": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.Bytes", actual)
}

func Test_Cov8_StacksTo_BytesDefault(t *testing.T) {
	result := codestack.StacksTo.BytesDefault()
	actual := args.Map{"lenGt0": len(result) > 0}
	expected := args.Map{"lenGt0": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.BytesDefault", actual)
}

func Test_Cov8_StacksTo_String(t *testing.T) {
	result := codestack.StacksTo.String(0, 5)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.String", actual)
}

func Test_Cov8_StacksTo_StringDefault(t *testing.T) {
	result := codestack.StacksTo.StringDefault()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.StringDefault", actual)
}

func Test_Cov8_StacksTo_StringNoCount(t *testing.T) {
	result := codestack.StacksTo.StringNoCount(0)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.StringNoCount", actual)
}

func Test_Cov8_StacksTo_JsonString(t *testing.T) {
	result := codestack.StacksTo.JsonString(0)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.JsonString", actual)
}

func Test_Cov8_StacksTo_JsonStringDefault(t *testing.T) {
	result := codestack.StacksTo.JsonStringDefault()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.JsonStringDefault", actual)
}

func Test_Cov8_StacksTo_StringUsingFmt(t *testing.T) {
	result := codestack.StacksTo.StringUsingFmt(
		func(trace *codestack.Trace) string { return "x" },
		0, 5,
	)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.StringUsingFmt", actual)
}

// ── File / Dir ──

func Test_Cov8_File_Name(t *testing.T) {
	result := codestack.File.Name(0)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "File.Name", actual)
}

func Test_Cov8_File_Path(t *testing.T) {
	result := codestack.File.Path(0)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "File.Path", actual)
}

func Test_Cov8_File_CurrentFilePath(t *testing.T) {
	result := codestack.File.CurrentFilePath()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "File.CurrentFilePath", actual)
}

func Test_Cov8_File_PathLineSep(t *testing.T) {
	path, line := codestack.File.PathLineSep(0)
	actual := args.Map{"pathNotEmpty": path != "", "lineGt0": line > 0}
	expected := args.Map{"pathNotEmpty": true, "lineGt0": true}
	expected.ShouldBeEqual(t, 0, "File.PathLineSep", actual)
}

func Test_Cov8_File_PathLineSepDefault(t *testing.T) {
	path, line := codestack.File.PathLineSepDefault()
	actual := args.Map{"pathNotEmpty": path != "", "lineGt0": line > 0}
	expected := args.Map{"pathNotEmpty": true, "lineGt0": true}
	expected.ShouldBeEqual(t, 0, "File.PathLineSepDefault", actual)
}

func Test_Cov8_File_FilePathWithLineString(t *testing.T) {
	result := codestack.File.FilePathWithLineString(0)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "File.FilePathWithLineString", actual)
}

func Test_Cov8_File_PathLineStringDefault(t *testing.T) {
	result := codestack.File.PathLineStringDefault()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "File.PathLineStringDefault", actual)
}

func Test_Cov8_Dir_CurDir(t *testing.T) {
	result := codestack.Dir.CurDir()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir.CurDir", actual)
}

func Test_Cov8_Dir_Get(t *testing.T) {
	result := codestack.Dir.Get(0)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir.Get", actual)
}

func Test_Cov8_Dir_RepoDir(t *testing.T) {
	result := codestack.Dir.RepoDir()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir.RepoDir", actual)
}

func Test_Cov8_Dir_CurDirJoin(t *testing.T) {
	result := codestack.Dir.CurDirJoin("sub", "dir")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir.CurDirJoin", actual)
}

func Test_Cov8_Dir_RepoDirJoin(t *testing.T) {
	result := codestack.Dir.RepoDirJoin("sub")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir.RepoDirJoin", actual)
}

// ── NameOf extended ──

func Test_Cov8_NameOf_MethodByFullName(t *testing.T) {
	result := codestack.NameOf.MethodByFullName("github.com/pkg.Type.Method")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameOf.MethodByFullName", actual)
}

func Test_Cov8_NameOf_PackageByFullName(t *testing.T) {
	result := codestack.NameOf.PackageByFullName("github.com/pkg.Type.Method")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameOf.PackageByFullName", actual)
}

func Test_Cov8_NameOf_CurrentFuncFullPath(t *testing.T) {
	result := codestack.NameOf.CurrentFuncFullPath("github.com/pkg.Type.Method")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameOf.CurrentFuncFullPath", actual)
}

func Test_Cov8_NameOf_JoinPackageNameWithRelative(t *testing.T) {
	result := codestack.NameOf.JoinPackageNameWithRelative("github.com/pkg.Type.Method", "SubMethod")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameOf.JoinPackageNameWithRelative", actual)
}

func Test_Cov8_NameOf_MethodStackSkip(t *testing.T) {
	result := codestack.NameOf.MethodStackSkip(0)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameOf.MethodStackSkip", actual)
}

func Test_Cov8_NameOf_PackageStackSkip(t *testing.T) {
	result := codestack.NameOf.PackageStackSkip(0)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameOf.PackageStackSkip", actual)
}

// ── newStacksCreator ──

func Test_Cov8_StackTrace_SkipOne(t *testing.T) {
	tc := codestack.New.StackTrace.SkipOne()
	actual := args.Map{"hasAny": tc.HasAnyItem()}
	expected := args.Map{"hasAny": true}
	expected.ShouldBeEqual(t, 0, "StackTrace.SkipOne", actual)
}

func Test_Cov8_StackTrace_Default(t *testing.T) {
	tc := codestack.New.StackTrace.Default(0, 5)
	actual := args.Map{"hasAny": tc.HasAnyItem()}
	expected := args.Map{"hasAny": true}
	expected.ShouldBeEqual(t, 0, "StackTrace.Default", actual)
}

func Test_Cov8_StackTrace_DefaultCount(t *testing.T) {
	tc := codestack.New.StackTrace.DefaultCount(0)
	actual := args.Map{"hasAny": tc.HasAnyItem()}
	expected := args.Map{"hasAny": true}
	expected.ShouldBeEqual(t, 0, "StackTrace.DefaultCount", actual)
}

// ── NewStackTraces / NewDefaultStackTraces ──

func Test_Cov8_TraceCollection_NewStackTraces(t *testing.T) {
	tc := codestack.TraceCollection{}
	result := tc.NewStackTraces(0)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection NewStackTraces", actual)
}

func Test_Cov8_TraceCollection_NewDefaultStackTraces(t *testing.T) {
	tc := codestack.TraceCollection{}
	result := tc.NewDefaultStackTraces()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection NewDefaultStackTraces", actual)
}

func Test_Cov8_TraceCollection_NewStackTracesJsonResult(t *testing.T) {
	tc := codestack.TraceCollection{}
	result := tc.NewStackTracesJsonResult(0)
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection NewStackTracesJsonResult", actual)
}

func Test_Cov8_TraceCollection_NewDefaultStackTracesJsonResult(t *testing.T) {
	tc := codestack.TraceCollection{}
	result := tc.NewDefaultStackTracesJsonResult()
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection NewDefaultStackTracesJsonResult", actual)
}

// ── newTraceCollection ──

func Test_Cov8_NewTraceCollection_Empty(t *testing.T) {
	// Test via TraceCollection Clear which creates empty
	tc := codestack.New.StackTrace.SkipNone()
	tc.Clear()
	actual := args.Map{"empty": tc.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "newTraceCollection Empty via Clear", actual)
}

// ── isSkippablePackage ──

func Test_Cov8_Trace_Create_HighSkipIndex(t *testing.T) {
	trace := codestack.New.Create(9999)
	actual := args.Map{"isOkay": trace.IsOkay}
	expected := args.Map{"isOkay": false}
	expected.ShouldBeEqual(t, 0, "Trace Create high skip", actual)
}
