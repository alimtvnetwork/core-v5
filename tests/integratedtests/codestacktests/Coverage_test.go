package codestacktests

import (
	"testing"

	"github.com/alimtvnetwork/core/codestack"
)

// ── FileWithLine nil-safety ──

func Test_Cov_FileWithLine_NilSafe(t *testing.T) {
	for caseIndex, tc := range coverageFileWithLineNilSafeCases {
		// Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// ── Trace nil-safety ──

func Test_Cov_Trace_NilSafe(t *testing.T) {
	for caseIndex, tc := range coverageTraceNilSafeCases {
		// Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// ── FileWithLine value tests ──

func Test_Cov_FileWithLine_Value(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{
		FilePath: "/tmp/test.go",
		Line:     42,
	}

	// Act & Assert
	if fwl.FullFilePath() != "/tmp/test.go" {
		t.Error("FullFilePath mismatch")
	}

	if fwl.LineNumber() != 42 {
		t.Error("LineNumber mismatch")
	}

	if fwl.IsNil() {
		t.Error("should not be nil")
	}

	if !fwl.IsNotNil() {
		t.Error("should be not nil")
	}

	if fwl.String() == "" {
		t.Error("String should not be empty")
	}

	if fwl.FileWithLine() == "" {
		t.Error("FileWithLine should not be empty")
	}

	// JsonModel
	model := fwl.JsonModel()
	if model.FilePath != "/tmp/test.go" {
		t.Error("JsonModel FilePath mismatch")
	}

	// JsonModelAny
	modelAny := fwl.JsonModelAny()
	if modelAny == nil {
		t.Error("JsonModelAny should not be nil")
	}

	// Json
	jsonResult := fwl.Json()
	if jsonResult.JsonString() == "" {
		t.Error("Json string should not be empty")
	}

	// JsonPtr
	jsonPtr := fwl.JsonPtr()
	if jsonPtr == nil {
		t.Error("JsonPtr should not be nil")
	}

	// JsonString
	js := fwl.JsonString()
	if js == "" {
		t.Error("JsonString should not be empty")
	}

	// StringUsingFmt
	fmtStr := fwl.StringUsingFmt(func(f codestack.FileWithLine) string {
		return f.FilePath
	})
	if fmtStr != "/tmp/test.go" {
		t.Error("StringUsingFmt mismatch")
	}

	// AsFileLiner
	liner := fwl.AsFileLiner()
	if liner == nil {
		t.Error("AsFileLiner should not be nil")
	}
}

func Test_Cov_FileWithLine_ParseJson(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{
		FilePath: "/tmp/test.go",
		Line:     42,
	}
	jsonResult := fwl.Json()
	jsonPtr := &jsonResult

	// Act
	target := &codestack.FileWithLine{}
	result, err := target.ParseInjectUsingJson(jsonPtr)

	// Assert
	if err != nil {
		t.Errorf("ParseInjectUsingJson error: %v", err)
	}

	if result.FilePath != "/tmp/test.go" {
		t.Error("parsed FilePath mismatch")
	}
}

func Test_Cov_FileWithLine_ParseJsonMust(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{
		FilePath: "/tmp/test.go",
		Line:     42,
	}
	jsonResult := fwl.Json()
	jsonPtr := &jsonResult

	// Act
	target := &codestack.FileWithLine{}
	result := target.ParseInjectUsingJsonMust(jsonPtr)

	// Assert
	if result.FilePath != "/tmp/test.go" {
		t.Error("ParseInjectUsingJsonMust FilePath mismatch")
	}
}

func Test_Cov_FileWithLine_JsonParseSelfInject(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{
		FilePath: "/tmp/test.go",
		Line:     42,
	}
	jsonResult := fwl.Json()
	jsonPtr := &jsonResult

	// Act
	target := &codestack.FileWithLine{}
	err := target.JsonParseSelfInject(jsonPtr)

	// Assert
	if err != nil {
		t.Errorf("JsonParseSelfInject error: %v", err)
	}
}

// ── Trace value tests ──

func Test_Cov_Trace_Value(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act & Assert
	if trace.IsNil() {
		t.Error("should not be nil pointer — it's a value")
	}

	if !trace.IsOkay {
		t.Error("Default trace should be okay")
	}

	if trace.PackageMethodName == "" {
		t.Error("PackageMethodName should not be empty")
	}

	if trace.Message() == "" {
		t.Error("Message should not be empty")
	}

	if trace.ShortString() == "" {
		t.Error("ShortString should not be empty")
	}

	if trace.FullFilePath() == "" {
		t.Error("FullFilePath should not be empty")
	}

	if trace.FileName() == "" {
		t.Error("FileName should not be empty")
	}

	if trace.LineNumber() == 0 {
		t.Error("LineNumber should not be 0")
	}

	if trace.FileWithLineString() == "" {
		t.Error("FileWithLineString should not be empty")
	}

	fwl := trace.FileWithLine()
	if fwl.FilePath == "" {
		t.Error("FileWithLine FilePath should not be empty")
	}

	if trace.String() == "" {
		t.Error("String should not be empty")
	}
}

func Test_Cov_Trace_StringUsingFmt(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act
	result := trace.StringUsingFmt(func(tr codestack.Trace) string {
		return tr.PackageName
	})

	// Assert
	if result == "" {
		t.Error("StringUsingFmt should not be empty")
	}
}

func Test_Cov_Trace_Clone(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act
	cloned := trace.Clone()

	// Assert
	if cloned.PackageMethodName != trace.PackageMethodName {
		t.Error("Clone PackageMethodName mismatch")
	}

	clonedPtr := trace.ClonePtr()
	if clonedPtr == nil {
		t.Error("ClonePtr should not be nil")
	}
}

func Test_Cov_Trace_Json(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act & Assert
	model := trace.JsonModel()
	if model.PackageName == "" {
		t.Error("JsonModel PackageName should not be empty")
	}

	modelAny := trace.JsonModelAny()
	if modelAny == nil {
		t.Error("JsonModelAny should not be nil")
	}

	js := trace.JsonString()
	if js == "" {
		t.Error("JsonString should not be empty")
	}

	jsonResult := trace.Json()
	if jsonResult.JsonString() == "" {
		t.Error("Json string should not be empty")
	}

	jsonPtr := trace.JsonPtr()
	if jsonPtr == nil {
		t.Error("JsonPtr should not be nil")
	}

	liner := trace.AsFileLiner()
	if liner == nil {
		t.Error("AsFileLiner should not be nil")
	}
}

func Test_Cov_Trace_ParseJson(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	jsonResult := trace.Json()
	jsonPtr := &jsonResult

	// Act
	target := &codestack.Trace{}
	result, err := target.ParseInjectUsingJson(jsonPtr)

	// Assert
	if err != nil {
		t.Errorf("ParseInjectUsingJson error: %v", err)
	}

	if result.PackageName == "" {
		t.Error("parsed PackageName should not be empty")
	}
}

func Test_Cov_Trace_ParseJsonMust(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	jsonResult := trace.Json()
	jsonPtr := &jsonResult

	// Act
	target := &codestack.Trace{}
	result := target.ParseInjectUsingJsonMust(jsonPtr)

	// Assert
	if result.PackageName == "" {
		t.Error("ParseInjectUsingJsonMust PackageName mismatch")
	}
}

func Test_Cov_Trace_JsonParseSelfInject(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	jsonResult := trace.Json()
	jsonPtr := &jsonResult

	// Act
	target := &codestack.Trace{}
	err := target.JsonParseSelfInject(jsonPtr)

	// Assert
	if err != nil {
		t.Errorf("JsonParseSelfInject error: %v", err)
	}
}

func Test_Cov_Trace_Dispose(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act
	trace.Dispose()

	// Assert
	if trace.PackageName != "" {
		t.Error("PackageName should be empty after Dispose")
	}

	if trace.IsOkay {
		t.Error("IsOkay should be false after Dispose")
	}
}

func Test_Cov_Trace_HasIssues(t *testing.T) {
	// Arrange
	trace := codestack.Trace{}

	// Act
	hasIssues := trace.HasIssues()

	// Assert
	if !hasIssues {
		t.Error("empty Trace should have issues")
	}
}

// ── TraceCollection tests (unique coverage methods) ──

func Test_Cov_TraceCollection_NewAndBasic(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Default()

	// Act & Assert
	if tc.IsEmpty() {
		t.Error("Default TraceCollection should not be empty")
	}

	if !tc.HasAnyItem() {
		t.Error("should have items")
	}

	if tc.Length() == 0 {
		t.Error("Length should not be 0")
	}

	if tc.Count() == 0 {
		t.Error("Count should not be 0")
	}

	if tc.LastIndex() < 0 {
		t.Error("LastIndex should be >= 0")
	}
}

func Test_Cov_TraceCollection_FirstLast(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Default()

	// Act & Assert
	first := tc.First()
	if first.PackageName == "" {
		t.Error("First should have PackageName")
	}

	last := tc.Last()
	if last.PackageName == "" {
		t.Error("Last should have PackageName")
	}

	firstDyn := tc.FirstDynamic()
	if firstDyn == nil {
		t.Error("FirstDynamic should not be nil")
	}

	lastDyn := tc.LastDynamic()
	if lastDyn == nil {
		t.Error("LastDynamic should not be nil")
	}

	firstOrDefault := tc.FirstOrDefault()
	if firstOrDefault.PackageName == "" {
		t.Error("FirstOrDefault should have PackageName")
	}

	lastOrDefault := tc.LastOrDefault()
	if lastOrDefault.PackageName == "" {
		t.Error("LastOrDefault should have PackageName")
	}
}

func Test_Cov_TraceCollection_Strings(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Default()

	// Act & Assert
	strs := tc.Strings()
	if len(strs) == 0 {
		t.Error("Strings should not be empty")
	}

	shortStrs := tc.ShortStrings()
	if len(shortStrs) == 0 {
		t.Error("ShortStrings should not be empty")
	}

	joinStr := tc.Join(", ")
	if joinStr == "" {
		t.Error("Join should not be empty")
	}

	joinLines := tc.JoinLines()
	if joinLines == "" {
		t.Error("JoinLines should not be empty")
	}

	csvStr := tc.JoinCsv()
	if csvStr == "" {
		t.Error("JoinCsv should not be empty")
	}

	jsonStr := tc.JsonString()
	if jsonStr == "" {
		t.Error("JsonString should not be empty")
	}

	str := tc.String()
	if str == "" {
		t.Error("String should not be empty")
	}
}

func Test_Cov_TraceCollection_SkipTake(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Default()
	length := tc.Length()

	// Act & Assert
	skipped := tc.Skip(1)
	if len(skipped) >= length {
		t.Error("Skip should reduce length")
	}

	taken := tc.Take(1)
	if len(taken) != 1 {
		t.Error("Take 1 should return 1 item")
	}

	limited := tc.Limit(1)
	if len(limited) != 1 {
		t.Error("Limit 1 should return 1 item")
	}

	skipCol := tc.SkipCollection(1)
	if skipCol.Length() >= length {
		t.Error("SkipCollection should reduce length")
	}

	takeCol := tc.TakeCollection(1)
	if takeCol.Length() != 1 {
		t.Error("TakeCollection should return 1")
	}

	limitCol := tc.LimitCollection(1)
	if limitCol.Length() != 1 {
		t.Error("LimitCollection should return 1")
	}

	safeLimit := tc.SafeLimitCollection(1)
	if safeLimit.Length() != 1 {
		t.Error("SafeLimitCollection should return 1")
	}
}

func Test_Cov_TraceCollection_FileWithLines(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Default()

	// Act & Assert
	fwls := tc.FileWithLines()
	if len(fwls) == 0 {
		t.Error("FileWithLines should not be empty")
	}

	fwlStrs := tc.FileWithLinesStrings()
	if len(fwlStrs) == 0 {
		t.Error("FileWithLinesStrings should not be empty")
	}

	fwlStr := tc.FileWithLinesString(", ")
	if fwlStr == "" {
		t.Error("FileWithLinesString should not be empty")
	}

	joinFwlStr := tc.JoinFileWithLinesStrings(", ")
	if joinFwlStr == "" {
		t.Error("JoinFileWithLinesStrings should not be empty")
	}
}

func Test_Cov_TraceCollection_Json(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Default()

	// Act & Assert
	jsonStrs := tc.JsonStrings()
	if len(jsonStrs) == 0 {
		t.Error("JsonStrings should not be empty")
	}

	joinJsonStr := tc.JoinJsonStrings(", ")
	if joinJsonStr == "" {
		t.Error("JoinJsonStrings should not be empty")
	}

	jsonModel := tc.JsonModel()
	if jsonModel == nil {
		t.Error("JsonModel should not be nil")
	}

	jsonModelAny := tc.JsonModelAny()
	if jsonModelAny == nil {
		t.Error("JsonModelAny should not be nil")
	}

	jsonResult := tc.Json()
	if jsonResult.JsonString() == "" {
		t.Error("Json should not be empty")
	}

	jsonPtr := tc.JsonPtr()
	if jsonPtr == nil {
		t.Error("JsonPtr should not be nil")
	}

	csvStrs := tc.CsvStrings()
	if len(csvStrs) == 0 {
		t.Error("CsvStrings should not be empty")
	}
}

func Test_Cov_TraceCollection_Reverse(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Default()

	// Act
	reversed := tc.Reverse()

	// Assert
	if reversed.Length() != tc.Length() {
		t.Error("Reverse should preserve length")
	}
}

func Test_Cov_TraceCollection_IsEqual(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Default()

	// Act & Assert
	if !tc.IsEqual(tc) {
		t.Error("collection should be equal to itself")
	}
}

func Test_Cov_TraceCollection_Clone(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Default()

	// Act
	cloned := tc.Clone()

	// Assert
	if cloned.Length() != tc.Length() {
		t.Error("Clone should preserve length")
	}

	clonedPtr := tc.ClonePtr()
	if clonedPtr == nil {
		t.Error("ClonePtr should not be nil")
	}
}

func Test_Cov_TraceCollection_ClearDispose(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Default()

	// Act
	tc.Clear()

	// Assert
	if !tc.IsEmpty() {
		t.Error("should be empty after Clear")
	}
}

func Test_Cov_TraceCollection_Add(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	trace := codestack.New.Default()

	// Act
	tc.Add(trace)

	// Assert
	if tc.IsEmpty() {
		t.Error("should not be empty after Add")
	}
}

func Test_TraceCollection_Paging_Cov(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Default()

	// Act
	pages := tc.GetPagesSize(2)

	// Assert
	if pages < 1 {
		t.Error("GetPagesSize should return at least 1")
	}
}

func Test_TraceCollection_CodeStacksString_Cov(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Default()

	// Act & Assert
	csStr := tc.CodeStacksString(", ")
	if csStr == "" {
		t.Error("CodeStacksString should not be empty")
	}

	csStrLimit := tc.CodeStacksStringLimit(", ", 1)
	if csStrLimit == "" {
		t.Error("CodeStacksStringLimit should not be empty")
	}
}

func Test_TraceCollection_StringsUsingFmt_Cov(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Default()

	// Act
	strs := tc.StringsUsingFmt(func(tr codestack.Trace) string {
		return tr.PackageName
	})

	// Assert
	if len(strs) == 0 {
		t.Error("StringsUsingFmt should not be empty")
	}

	joinStr := tc.JoinUsingFmt(", ", func(tr codestack.Trace) string {
		return tr.PackageName
	})
	if joinStr == "" {
		t.Error("JoinUsingFmt should not be empty")
	}
}

func Test_TraceCollection_JoinShortStrings_Cov(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Default()

	// Act
	joinShort := tc.JoinShortStrings(", ")

	// Assert
	if joinShort == "" {
		t.Error("JoinShortStrings should not be empty")
	}
}

func Test_TraceCollection_JoinCsvLine_Cov(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Default()

	// Act
	csvLine := tc.JoinCsvLine()

	// Assert
	if csvLine == "" {
		t.Error("JoinCsvLine should not be empty")
	}
}

func Test_TraceCollection_HasIndex_Cov(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Default()

	// Act & Assert
	if !tc.HasIndex(0) {
		t.Error("HasIndex 0 should be true")
	}

	if tc.HasIndex(9999) {
		t.Error("HasIndex 9999 should be false")
	}
}

func Test_TraceCollection_Serializer_Cov(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Default()

	// Act
	s := tc.Serializer()

	// Assert
	if s == nil {
		t.Error("Serializer should not be nil")
	}
}

func Test_TraceCollection_StackTracesBytes_Cov(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Default()

	// Act
	bytes := tc.StackTracesBytes()

	// Assert
	if len(bytes) == 0 {
		t.Error("StackTracesBytes should not be empty")
	}
}

func Test_TraceCollection_ParseJson_Cov(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Default()
	jsonResult := tc.Json()
	jsonPtr := &jsonResult

	// Act
	target := &codestack.TraceCollection{}
	err := target.JsonParseSelfInject(jsonPtr)

	// Assert
	if err != nil {
		t.Errorf("JsonParseSelfInject error: %v", err)
	}
}

func Test_TraceCollection_Dispose_Cov(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Default()

	// Act
	tc.Dispose()

	// Assert
	if !tc.IsEmpty() {
		t.Error("should be empty after Dispose")
	}
}

// ── NameOf tests ──

func Test_NameOf_Method_Cov(t *testing.T) {
	// Act
	name := codestack.NameOf.Method("github.com/alimtvnetwork/core/codestack.Test_NameOf_Method_Cov")

	// Assert
	if name == "" {
		t.Error("Method should not be empty")
	}
}

func Test_NameOf_Package_Cov(t *testing.T) {
	// Act
	name := codestack.NameOf.Package("github.com/alimtvnetwork/core/codestack.Test_NameOf_Package_Cov")

	// Assert
	if name == "" {
		t.Error("Package should not be empty")
	}
}

func Test_NameOf_All_Cov(t *testing.T) {
	// Act
	full, pkg, method := codestack.NameOf.All("github.com/alimtvnetwork/core/codestack.Test_NameOf_All_Cov")

	// Assert
	if full == "" {
		t.Error("full should not be empty")
	}

	if pkg == "" {
		t.Error("pkg should not be empty")
	}

	if method == "" {
		t.Error("method should not be empty")
	}
}

// ── newCreator tests ──

func Test_NewCreator_SkipOne_Cov(t *testing.T) {
	// Act
	trace := codestack.New.SkipOne()

	// Assert
	if trace.PackageName == "" {
		t.Error("SkipOne PackageName should not be empty")
	}
}

func Test_NewCreator_Ptr_Cov(t *testing.T) {
	// Act
	trace := codestack.New.Ptr(0)

	// Assert
	if trace == nil {
		t.Error("Ptr should not be nil")
	}
}

// ── StackTrace tests ──

func Test_StackTrace_DefaultCount_Cov(t *testing.T) {
	// Act
	tc := codestack.New.StackTrace.DefaultCount(3)

	// Assert
	if tc.IsEmpty() {
		t.Error("DefaultCount should not be empty")
	}
}

func Test_StackTrace_SkipOne_Cov(t *testing.T) {
	// Act
	tc := codestack.New.StackTrace.SkipOne()

	// Assert
	if tc.IsEmpty() {
		t.Error("SkipOne should not be empty")
	}
}

func Test_StackTrace_SkipNone_Cov(t *testing.T) {
	// Act
	tc := codestack.New.StackTrace.SkipNone()

	// Assert
	if tc.IsEmpty() {
		t.Error("SkipNone should not be empty")
	}
}

// ── StacksTo tests ──

func Test_StacksTo_String_Cov(t *testing.T) {
	// Act
	result := codestack.StacksTo.String(2, 5)

	// Assert
	if result == "" {
		t.Error("StacksTo.String should not be empty")
	}
}

func Test_StacksTo_StringDefault_Cov(t *testing.T) {
	// Act
	result := codestack.StacksTo.StringDefault()

	// Assert
	if result == "" {
		t.Error("StacksTo.StringDefault should not be empty")
	}
}

func Test_StacksTo_Bytes_Cov(t *testing.T) {
	// Act
	result := codestack.StacksTo.Bytes(2, 5)

	// Assert
	if len(result) == 0 {
		t.Error("StacksTo.Bytes should not be empty")
	}
}

func Test_StacksTo_BytesDefault_Cov(t *testing.T) {
	// Act
	result := codestack.StacksTo.BytesDefault()

	// Assert
	if len(result) == 0 {
		t.Error("StacksTo.BytesDefault should not be empty")
	}
}

func Test_StacksTo_JsonString_Cov(t *testing.T) {
	// Act
	result := codestack.StacksTo.JsonString(2, 5)

	// Assert
	if result == "" {
		t.Error("StacksTo.JsonString should not be empty")
	}
}

func Test_StacksTo_JsonStringDefault_Cov(t *testing.T) {
	// Act
	result := codestack.StacksTo.JsonStringDefault()

	// Assert
	if result == "" {
		t.Error("StacksTo.JsonStringDefault should not be empty")
	}
}

func Test_StacksTo_StringNoCount_Cov(t *testing.T) {
	// Act
	result := codestack.StacksTo.StringNoCount(5)

	// Assert
	if result == "" {
		t.Error("StacksTo.StringNoCount should not be empty")
	}
}

// ── File getter tests ──

func Test_File_Name_Cov(t *testing.T) {
	// Act
	name := codestack.File.Name(0)

	// Assert
	if name == "" {
		t.Error("File.Name should not be empty")
	}
}

func Test_File_Path_Cov(t *testing.T) {
	// Act
	path := codestack.File.Path(0)

	// Assert
	if path == "" {
		t.Error("File.Path should not be empty")
	}
}

// ── Dir getter tests ──

func Test_Dir_CurDir_Cov(t *testing.T) {
	// Act
	dir := codestack.Dir.CurDir()

	// Assert
	if dir == "" {
		t.Error("Dir.CurDir should not be empty")
	}
}

func Test_Dir_CurDirJoin_Cov(t *testing.T) {
	// Act
	dir := codestack.Dir.CurDirJoin("subdir")

	// Assert
	if dir == "" {
		t.Error("Dir.CurDirJoin should not be empty")
	}
}

func Test_TraceCollection_Concat_Cov(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Default()

	// Act
	concatted := tc.ConcatNew(codestack.New.Default())

	// Assert
	if concatted.Length() <= tc.Length() {
		t.Error("ConcatNew should increase length")
	}

	concatPtr := tc.ConcatNewPtr(codestack.New.Default())
	if concatPtr == nil {
		t.Error("ConcatNewPtr should not be nil")
	}
}

func Test_TraceCollection_Filters_Cov(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Default()

	// Act & Assert
	filtered := tc.Filter(func(index int, trace codestack.Trace) bool {
		return true
	})
	if filtered.Length() == 0 {
		t.Error("Filter should not be empty")
	}

	filteredLimit := tc.FilterWithLimit(1, func(index int, trace codestack.Trace) bool {
		return true
	})
	if filteredLimit.Length() == 0 {
		t.Error("FilterWithLimit should not be empty")
	}
}

func Test_TraceCollection_AsBindings_Cov(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Default()

	// Act & Assert
	binder := tc.AsJsonContractsBinder()
	if binder == nil {
		t.Error("AsJsonContractsBinder should not be nil")
	}

	jsoner := tc.AsJsoner()
	if jsoner == nil {
		t.Error("AsJsoner should not be nil")
	}

	injector := tc.AsJsonParseSelfInjector()
	if injector == nil {
		t.Error("AsJsonParseSelfInjector should not be nil")
	}
}
