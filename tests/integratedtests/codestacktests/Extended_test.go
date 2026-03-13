package codestacktests

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// stacksTo: Reflection-based method discovery
// =============================================================================

func Test_StacksTo_MethodDiscovery_Ext(t *testing.T) {
	for caseIndex, testCase := range extStacksToMethodDiscoveryTestCases {
		// Arrange
		rv := reflect.ValueOf(codestack.StacksTo)
		rt := rv.Type()

		methodMap := map[string]bool{}
		for i := 0; i < rt.NumMethod(); i++ {
			methodMap[rt.Method(i).Name] = true
		}

		// Act
		actual := args.Map{
			"hasBytesMethod":             methodMap["Bytes"],
			"hasBytesDefaultMethod":      methodMap["BytesDefault"],
			"hasStringMethod":            methodMap["String"],
			"hasStringUsingFmtMethod":    methodMap["StringUsingFmt"],
			"hasJsonStringMethod":        methodMap["JsonString"],
			"hasJsonStringDefaultMethod": methodMap["JsonStringDefault"],
			"hasStringNoCountMethod":     methodMap["StringNoCount"],
			"hasStringDefaultMethod":     methodMap["StringDefault"],
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// stacksTo: Bytes via reflection
// =============================================================================

func Test_StacksTo_Bytes_Reflection_Ext(t *testing.T) {
	for caseIndex, testCase := range extStacksToBytesTestCases {
		// Arrange
		rv := reflect.ValueOf(codestack.StacksTo)
		method := rv.MethodByName("Bytes")
		input := testCase.ArrangeInput.(args.Map)
		skipIndex, _ := input.GetAsInt("skipIndex")

		// Act
		results := method.Call([]reflect.Value{
			reflect.ValueOf(skipIndex),
		})
		bytes := results[0].Bytes()

		actual := args.Map{
			"notEmpty": len(bytes) > 0,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// stacksTo: BytesDefault via reflection
// =============================================================================

func Test_StacksTo_BytesDefault_Reflection_Ext(t *testing.T) {
	for caseIndex, testCase := range extStacksToBytesDefaultTestCases {
		// Arrange
		rv := reflect.ValueOf(codestack.StacksTo)
		method := rv.MethodByName("BytesDefault")

		// Act
		results := method.Call(nil)
		bytes := results[0].Bytes()

		actual := args.Map{
			"notEmpty": len(bytes) > 0,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// stacksTo: String via reflection
// =============================================================================

func Test_StacksTo_String_Reflection_Ext(t *testing.T) {
	for caseIndex, testCase := range extStacksToStringTestCases {
		// Arrange
		rv := reflect.ValueOf(codestack.StacksTo)
		method := rv.MethodByName("String")
		input := testCase.ArrangeInput.(args.Map)
		skipIndex, _ := input.GetAsInt("skipIndex")
		count, _ := input.GetAsInt("count")

		// Act
		results := method.Call([]reflect.Value{
			reflect.ValueOf(skipIndex),
			reflect.ValueOf(count),
		})
		str := results[0].String()

		actual := args.Map{
			"notEmpty": str != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// stacksTo: StringUsingFmt via reflection
// =============================================================================

func Test_StacksTo_StringUsingFmt_Reflection_Ext(t *testing.T) {
	for caseIndex, testCase := range extStacksToStringUsingFmtTestCases {
		// Arrange
		rv := reflect.ValueOf(codestack.StacksTo)
		method := rv.MethodByName("StringUsingFmt")

		formatter := codestack.Formatter(func(trace *codestack.Trace) string {
			return fmt.Sprintf(
				"%s:%d",
				trace.PackageMethodName,
				trace.Line,
			)
		})

		// Act
		results := method.Call([]reflect.Value{
			reflect.ValueOf(formatter),
			reflect.ValueOf(0),
			reflect.ValueOf(5),
		})
		str := results[0].String()

		actual := args.Map{
			"notEmpty": str != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// stacksTo: JsonString via reflection
// =============================================================================

func Test_StacksTo_JsonString_Reflection_Ext(t *testing.T) {
	for caseIndex, testCase := range extStacksToJsonStringTestCases {
		// Arrange
		rv := reflect.ValueOf(codestack.StacksTo)
		method := rv.MethodByName("JsonString")
		input := testCase.ArrangeInput.(args.Map)
		skipIndex, _ := input.GetAsInt("skipIndex")

		// Act
		results := method.Call([]reflect.Value{
			reflect.ValueOf(skipIndex),
		})
		str := results[0].String()

		actual := args.Map{
			"notEmpty": str != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// stacksTo: JsonStringDefault via reflection
// =============================================================================

func Test_StacksTo_JsonStringDefault_Reflection_Ext(t *testing.T) {
	for caseIndex, testCase := range extStacksToJsonStringDefaultTestCases {
		// Arrange
		rv := reflect.ValueOf(codestack.StacksTo)
		method := rv.MethodByName("JsonStringDefault")

		// Act
		results := method.Call(nil)
		str := results[0].String()

		actual := args.Map{
			"notEmpty": str != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// stacksTo: StringNoCount via reflection
// =============================================================================

func Test_StacksTo_StringNoCount_Reflection_Ext(t *testing.T) {
	for caseIndex, testCase := range extStacksToStringNoCountTestCases {
		// Arrange
		rv := reflect.ValueOf(codestack.StacksTo)
		method := rv.MethodByName("StringNoCount")
		input := testCase.ArrangeInput.(args.Map)
		skipIndex, _ := input.GetAsInt("skipIndex")

		// Act
		results := method.Call([]reflect.Value{
			reflect.ValueOf(skipIndex),
		})
		str := results[0].String()

		actual := args.Map{
			"notEmpty": str != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// stacksTo: StringDefault via reflection
// =============================================================================

func Test_StacksTo_StringDefault_Reflection_Ext(t *testing.T) {
	for caseIndex, testCase := range extStacksToStringDefaultTestCases {
		// Arrange
		rv := reflect.ValueOf(codestack.StacksTo)
		method := rv.MethodByName("StringDefault")

		// Act
		results := method.Call(nil)
		str := results[0].String()

		actual := args.Map{
			"notEmpty": str != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// currentNameOf
// =============================================================================

func Test_NameOf_All_Ext(t *testing.T) {
	for caseIndex, testCase := range extNameOfAllTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		fullName, _ := input.GetAsString("fullName")

		// Act
		fullMethod, pkgName, methodName := codestack.NameOf.All(fullName)

		actual := args.Map{
			"hasFullMethod":  fullMethod != "",
			"hasPackageName": pkgName != "",
			"hasMethodName":  methodName != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NameOf_Method_Ext(t *testing.T) {
	for caseIndex, testCase := range extNameOfMethodTestCases {
		// Act
		methodName := codestack.NameOf.Method()

		actual := args.Map{
			"hasMethodName": methodName != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NameOf_Package_Ext(t *testing.T) {
	for caseIndex, testCase := range extNameOfPackageTestCases {
		// Act
		pkgName := codestack.NameOf.Package()

		actual := args.Map{
			"hasPackageName": pkgName != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NameOf_MethodByFullName_Ext(t *testing.T) {
	// Arrange
	fullName := "github.com/mypackage.(*MyStruct).DoWork"

	// Act
	methodName := codestack.NameOf.MethodByFullName(fullName)

	// Assert
	if methodName == "" {
		t.Error("MethodByFullName should return non-empty")
	}
}

func Test_NameOf_PackageByFullName_Ext(t *testing.T) {
	// Arrange
	fullName := "github.com/mypackage.(*MyStruct).DoWork"

	// Act
	pkgName := codestack.NameOf.PackageByFullName(fullName)

	// Assert
	if pkgName == "" {
		t.Error("PackageByFullName should return non-empty")
	}
}

func Test_NameOf_JoinPackageNameWithRelative_Ext(t *testing.T) {
	// Arrange
	fullName := "github.com/mypackage.(*MyStruct).DoWork"

	// Act
	result := codestack.NameOf.JoinPackageNameWithRelative(
		fullName,
		"SubStruct.Method",
	)

	// Assert
	if result == "" {
		t.Error("JoinPackageNameWithRelative should return non-empty")
	}
}

func Test_NameOf_CurrentFuncFullPath_Ext(t *testing.T) {
	// Arrange
	fullName := "github.com/mypackage.(*MyStruct).DoWork"

	// Act
	result := codestack.NameOf.CurrentFuncFullPath(fullName)

	// Assert
	if result == "" {
		t.Error("CurrentFuncFullPath should return non-empty")
	}
}

func Test_NameOf_MethodStackSkip_Ext(t *testing.T) {
	// Act
	result := codestack.NameOf.MethodStackSkip(0)

	// Assert
	if result == "" {
		t.Error("MethodStackSkip should return non-empty")
	}
}

func Test_NameOf_PackageStackSkip_Ext(t *testing.T) {
	// Act
	result := codestack.NameOf.PackageStackSkip(0)

	// Assert
	if result == "" {
		t.Error("PackageStackSkip should return non-empty")
	}
}

// =============================================================================
// newCreator
// =============================================================================

func Test_NewCreator_Default_Ext(t *testing.T) {
	for caseIndex, testCase := range extNewCreatorDefaultTestCases {
		// Act
		trace := codestack.New.Default()

		actual := args.Map{
			"isOkay":        trace.IsOkay,
			"hasFilePath":   trace.FilePath != "",
			"hasMethodName": trace.MethodName != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NewCreator_SkipOne_Ext(t *testing.T) {
	for caseIndex, testCase := range extNewCreatorSkipOneTestCases {
		// Act
		trace := codestack.New.SkipOne()

		actual := args.Map{
			"isOkay":      trace.IsOkay,
			"hasFilePath": trace.FilePath != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NewCreator_Ptr_Ext(t *testing.T) {
	for caseIndex, testCase := range extNewCreatorPtrTestCases {
		// Act
		trace := codestack.New.Ptr(0)

		actual := args.Map{
			"isNil":  trace == nil,
			"isOkay": trace.IsOkay,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// newStacksCreator via reflection
// =============================================================================

func Test_NewStacksCreator_Default_Reflection_Ext(t *testing.T) {
	for caseIndex, testCase := range extNewStacksCreatorDefaultTestCases {
		// Arrange
		rv := reflect.ValueOf(codestack.New.StackTrace)
		method := rv.MethodByName("Default")

		// Act
		results := method.Call([]reflect.Value{
			reflect.ValueOf(0),
			reflect.ValueOf(3),
		})
		collection := results[0].Interface().(codestack.TraceCollection)

		actual := args.Map{
			"hasItems": collection.HasAnyItem(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NewStacksCreator_SkipOne_Reflection_Ext(t *testing.T) {
	for caseIndex, testCase := range extNewStacksCreatorSkipOneTestCases {
		// Arrange
		rv := reflect.ValueOf(codestack.New.StackTrace)
		method := rv.MethodByName("SkipOne")

		// Act
		results := method.Call(nil)
		collection := results[0].Interface().(codestack.TraceCollection)

		actual := args.Map{
			"hasItems": collection.HasAnyItem(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NewStacksCreator_SkipNone_Reflection_Ext(t *testing.T) {
	for caseIndex, testCase := range extNewStacksCreatorSkipNoneTestCases {
		// Arrange
		rv := reflect.ValueOf(codestack.New.StackTrace)
		method := rv.MethodByName("SkipNone")

		// Act
		results := method.Call(nil)
		collection := results[0].Interface().(codestack.TraceCollection)

		actual := args.Map{
			"hasItems": collection.HasAnyItem(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NewStacksCreator_DefaultCount_Reflection_Ext(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(codestack.New.StackTrace)
	method := rv.MethodByName("DefaultCount")

	// Act
	results := method.Call([]reflect.Value{
		reflect.ValueOf(0),
	})
	collection := results[0].Interface().(codestack.TraceCollection)

	// Assert
	if !collection.HasAnyItem() {
		t.Error("DefaultCount should return items")
	}
}

func Test_NewStacksCreator_All_Reflection_Ext(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(codestack.New.StackTrace)
	method := rv.MethodByName("All")

	// Act
	results := method.Call([]reflect.Value{
		reflect.ValueOf(true),
		reflect.ValueOf(true),
		reflect.ValueOf(0),
		reflect.ValueOf(5),
	})
	collection := results[0].Interface().(codestack.TraceCollection)

	// Assert
	if !collection.HasAnyItem() {
		t.Error("All should return items")
	}
}

// =============================================================================
// dirGetter
// =============================================================================

func Test_Dir_CurDir_Ext(t *testing.T) {
	for caseIndex, testCase := range extDirGetterCurDirTestCases {
		// Act
		result := codestack.Dir.CurDir()

		actual := args.Map{
			"notEmpty": result != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Dir_RepoDir_Ext(t *testing.T) {
	for caseIndex, testCase := range extDirGetterRepoDirTestCases {
		// Act
		result := codestack.Dir.RepoDir()

		actual := args.Map{
			"notEmpty": result != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Dir_Get_Ext(t *testing.T) {
	// Act
	result := codestack.Dir.Get(0)

	// Assert
	if result == "" {
		t.Error("Dir.Get should return non-empty")
	}
}

func Test_Dir_CurDirJoin_Ext(t *testing.T) {
	// Act
	result := codestack.Dir.CurDirJoin("sub", "path")

	// Assert
	if result == "" {
		t.Error("Dir.CurDirJoin should return non-empty")
	}
}

func Test_Dir_RepoDirJoin_Ext(t *testing.T) {
	// Act
	result := codestack.Dir.RepoDirJoin("sub")

	// Assert
	if result == "" {
		t.Error("Dir.RepoDirJoin should return non-empty")
	}
}

// =============================================================================
// fileGetter
// =============================================================================

func Test_File_Path_Ext(t *testing.T) {
	for caseIndex, testCase := range extFileGetterPathTestCases {
		// Act
		result := codestack.File.Path(0)

		actual := args.Map{
			"notEmpty": result != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_File_Name_Ext(t *testing.T) {
	for caseIndex, testCase := range extFileGetterNameTestCases {
		// Act
		result := codestack.File.Name(0)

		actual := args.Map{
			"notEmpty": result != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_File_CurrentFilePath_Ext(t *testing.T) {
	for caseIndex, testCase := range extFileGetterCurrentFilePathTestCases {
		// Act
		result := codestack.File.CurrentFilePath()

		actual := args.Map{
			"notEmpty": result != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_File_PathLineSep_Ext(t *testing.T) {
	// Act
	filePath, lineNumber := codestack.File.PathLineSep(0)

	// Assert
	if filePath == "" {
		t.Error("PathLineSep should return non-empty file path")
	}
	if lineNumber <= 0 {
		t.Error("PathLineSep should return positive line number")
	}
}

func Test_File_PathLineSepDefault_Ext(t *testing.T) {
	// Act
	filePath, lineNumber := codestack.File.PathLineSepDefault()

	// Assert
	if filePath == "" {
		t.Error("PathLineSepDefault should return non-empty file path")
	}
	if lineNumber <= 0 {
		t.Error("PathLineSepDefault should return positive line number")
	}
}

func Test_File_FilePathWithLineString_Ext(t *testing.T) {
	// Act
	result := codestack.File.FilePathWithLineString(0)

	// Assert
	if result == "" {
		t.Error("FilePathWithLineString should return non-empty")
	}
}

func Test_File_PathLineStringDefault_Ext(t *testing.T) {
	// Act
	result := codestack.File.PathLineStringDefault()

	// Assert
	if result == "" {
		t.Error("PathLineStringDefault should return non-empty")
	}
}

// =============================================================================
// TraceCollection: additional uncovered methods via reflection
// =============================================================================

func Test_TraceCollection_StackTracesJsonResult_Reflection_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()
	rv := reflect.ValueOf(collection).Elem()
	method := rv.MethodByName("StackTracesJsonResult")

	// Act
	results := method.Call(nil)

	// Assert
	if results[0].IsNil() {
		t.Error("StackTracesJsonResult should not be nil")
	}
}

func Test_TraceCollection_NewStackTraces_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.NewStackTraces(0)

	// Assert
	if result == "" {
		t.Error("NewStackTraces should return non-empty")
	}
}

func Test_TraceCollection_NewDefaultStackTraces_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.NewDefaultStackTraces()

	// Assert
	if result == "" {
		t.Error("NewDefaultStackTraces should return non-empty")
	}
}

func Test_TraceCollection_NewStackTracesJsonResult_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.NewStackTracesJsonResult(0)

	// Assert
	if result == nil {
		t.Error("NewStackTracesJsonResult should not be nil")
	}
}

func Test_TraceCollection_NewDefaultStackTracesJsonResult_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.NewDefaultStackTracesJsonResult()

	// Assert
	if result == nil {
		t.Error("NewDefaultStackTracesJsonResult should not be nil")
	}
}

func Test_TraceCollection_GetPagedCollection_Ext(t *testing.T) {
	// Arrange
	collection := &codestack.TraceCollection{
		Items: []codestack.Trace{
			newTestTrace("p1", 1),
			newTestTrace("p2", 2),
			newTestTrace("p3", 3),
			newTestTrace("p4", 4),
			newTestTrace("p5", 5),
		},
	}

	// Act
	pages := collection.GetPagedCollection(2)

	// Assert
	if len(pages) != 3 {
		t.Errorf("expected 3 pages, got %d", len(pages))
	}
}

func Test_TraceCollection_GetSinglePageCollection_Ext(t *testing.T) {
	// Arrange
	collection := &codestack.TraceCollection{
		Items: []codestack.Trace{
			newTestTrace("p1", 1),
			newTestTrace("p2", 2),
			newTestTrace("p3", 3),
			newTestTrace("p4", 4),
			newTestTrace("p5", 5),
		},
	}

	// Act
	page := collection.GetSinglePageCollection(2, 2)

	// Assert
	if page.Length() != 2 {
		t.Errorf("expected 2 items, got %d", page.Length())
	}
}

func Test_TraceCollection_GetSinglePageCollection_SmallList_Ext(t *testing.T) {
	// Arrange
	collection := &codestack.TraceCollection{
		Items: []codestack.Trace{
			newTestTrace("p1", 1),
		},
	}

	// Act
	page := collection.GetSinglePageCollection(5, 1)

	// Assert
	if page.Length() != 1 {
		t.Errorf("expected 1 item, got %d", page.Length())
	}
}

func Test_TraceCollection_FilterWithLimit_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.FilterWithLimit(
		2,
		func(trace *codestack.Trace) (bool, bool) {
			return true, false
		},
	)

	// Assert
	if len(result) != 2 {
		t.Errorf("expected 2 items, got %d", len(result))
	}
}

func Test_TraceCollection_FilterTraceCollection_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.FilterTraceCollection(
		func(trace *codestack.Trace) (bool, bool) {
			return trace.PackageName == "pkg1", false
		},
	)

	// Assert
	if result.Length() != 1 {
		t.Errorf("expected 1 item, got %d", result.Length())
	}
}

func Test_TraceCollection_SkipFilterMethodName_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.SkipFilterMethodNameTraceCollection("TestMethod")

	// Assert
	if result.Length() != 0 {
		t.Errorf("expected 0 items, got %d", result.Length())
	}
}

func Test_TraceCollection_FilterFullMethodName_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.FilterFullMethodNameTraceCollection("pkg1.TestMethod")

	// Assert
	if result.Length() != 1 {
		t.Errorf("expected 1 item, got %d", result.Length())
	}
}

func Test_TraceCollection_SkipFilterFullMethodName_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.SkipFilterFullMethodNameTraceCollection("pkg1.TestMethod")

	// Assert
	if result.Length() != 2 {
		t.Errorf("expected 2 items, got %d", result.Length())
	}
}

func Test_TraceCollection_SkipFilterFilename_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.SkipFilterFilenameTraceCollection("/src/pkg1/file.go")

	// Assert
	if result.Length() != 2 {
		t.Errorf("expected 2 items, got %d", result.Length())
	}
}

func Test_TraceCollection_FileWithLinesStrings_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	strs := collection.FileWithLinesStrings()

	// Assert
	if len(strs) != 3 {
		t.Errorf("expected 3, got %d", len(strs))
	}
}

func Test_TraceCollection_FileWithLinesString_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.FileWithLinesString()

	// Assert
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_TraceCollection_FileWithLinesString_Empty_Ext(t *testing.T) {
	// Arrange
	collection := emptyCollection()

	// Act
	result := collection.FileWithLinesString()

	// Assert
	if result != "" {
		t.Error("empty should return empty")
	}
}

func Test_TraceCollection_JoinShortStrings_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.JoinShortStrings(",")

	// Assert
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_TraceCollection_JoinFileWithLinesStrings_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.JoinFileWithLinesStrings(",")

	// Assert
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_TraceCollection_JoinJsonStrings_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.JoinJsonStrings(",")

	// Assert
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_TraceCollection_JsonStrings_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.JsonStrings()

	// Assert
	if len(result) != 3 {
		t.Errorf("expected 3, got %d", len(result))
	}
}

func Test_TraceCollection_Join_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.Join(",")

	// Assert
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_TraceCollection_JoinLines_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.JoinLines()

	// Assert
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_TraceCollection_JoinCsv_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.JoinCsv()

	// Assert
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_TraceCollection_JoinCsvLine_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.JoinCsvLine()

	// Assert
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_TraceCollection_CodeStacksStringLimit_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.CodeStacksStringLimit(2)

	// Assert
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_TraceCollection_CodeStacksStringLimit_Empty_Ext(t *testing.T) {
	// Arrange
	collection := emptyCollection()

	// Act
	result := collection.CodeStacksStringLimit(2)

	// Assert
	if result != "" {
		t.Error("empty should return empty")
	}
}

func Test_TraceCollection_CsvStrings_Empty_Ext(t *testing.T) {
	// Arrange
	collection := emptyCollection()

	// Act
	result := collection.CsvStrings()

	// Assert
	if len(result) != 0 {
		t.Errorf("expected 0, got %d", len(result))
	}
}

func Test_TraceCollection_FirstDynamic_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.FirstDynamic()

	// Assert
	if result == nil {
		t.Error("should return non-nil")
	}
}

func Test_TraceCollection_LastDynamic_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.LastDynamic()

	// Assert
	if result == nil {
		t.Error("should return non-nil")
	}
}

func Test_TraceCollection_FirstOrDefaultDynamic_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.FirstOrDefaultDynamic()

	// Assert
	if result == nil {
		t.Error("should return non-nil")
	}
}

func Test_TraceCollection_LastOrDefaultDynamic_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.LastOrDefaultDynamic()

	// Assert
	if result == nil {
		t.Error("should return non-nil")
	}
}

func Test_TraceCollection_SkipDynamic_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.SkipDynamic(1)

	// Assert
	if result == nil {
		t.Error("should return non-nil")
	}
}

func Test_TraceCollection_Skip_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.Skip(1)

	// Assert
	if len(result) != 2 {
		t.Errorf("expected 2, got %d", len(result))
	}
}

func Test_TraceCollection_TakeDynamic_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.TakeDynamic(2)

	// Assert
	if result == nil {
		t.Error("should return non-nil")
	}
}

func Test_TraceCollection_Take_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.Take(2)

	// Assert
	if len(result) != 2 {
		t.Errorf("expected 2, got %d", len(result))
	}
}

func Test_TraceCollection_LimitDynamic_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.LimitDynamic(2)

	// Assert
	if result == nil {
		t.Error("should return non-nil")
	}
}

func Test_TraceCollection_Limit_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.Limit(2)

	// Assert
	if len(result) != 2 {
		t.Errorf("expected 2, got %d", len(result))
	}
}

func Test_TraceCollection_Count_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	count := collection.Count()

	// Assert
	if count != 3 {
		t.Errorf("expected 3, got %d", count)
	}
}

func Test_TraceCollection_JsonModel_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	model := collection.JsonModel()

	// Assert
	if len(model) != 3 {
		t.Errorf("expected 3, got %d", len(model))
	}
}

func Test_TraceCollection_JsonModelAny_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.JsonModelAny()

	// Assert
	if result == nil {
		t.Error("should return non-nil")
	}
}

func Test_TraceCollection_Json_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.Json()

	// Assert
	if result.HasError() {
		t.Errorf("should not have error: %v", result.MeaningfulError())
	}
}

func Test_TraceCollection_AsJsonContractsBinder_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	binder := collection.AsJsonContractsBinder()

	// Assert
	if binder == nil {
		t.Error("should not be nil")
	}
}

func Test_TraceCollection_AsJsoner_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	jsoner := collection.AsJsoner()

	// Assert
	if jsoner == nil {
		t.Error("should not be nil")
	}
}

func Test_TraceCollection_AsJsonParseSelfInjector_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	injector := collection.AsJsonParseSelfInjector()

	// Assert
	if injector == nil {
		t.Error("should not be nil")
	}
}

func Test_TraceCollection_ParseInjectUsingJson_Ext(t *testing.T) {
	// Arrange
	original := newTestCollection()
	jsonResult := original.JsonPtr()

	// Act
	var target codestack.TraceCollection
	result, err := target.ParseInjectUsingJson(jsonResult)

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if result.Length() != 3 {
		t.Errorf("expected 3, got %d", result.Length())
	}
}

func Test_TraceCollection_ParseInjectUsingJsonMust_Ext(t *testing.T) {
	// Arrange
	original := newTestCollection()
	jsonResult := original.JsonPtr()

	// Act
	var target codestack.TraceCollection
	result := target.ParseInjectUsingJsonMust(jsonResult)

	// Assert
	if result.Length() != 3 {
		t.Errorf("expected 3, got %d", result.Length())
	}
}

func Test_TraceCollection_JsonParseSelfInject_Ext(t *testing.T) {
	// Arrange
	original := newTestCollection()
	jsonResult := original.JsonPtr()

	// Act
	var target codestack.TraceCollection
	err := target.JsonParseSelfInject(jsonResult)

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func Test_TraceCollection_IsEqualItems_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.IsEqualItems(collection.Items...)

	// Assert
	if !result {
		t.Error("should be equal")
	}
}

func Test_TraceCollection_IsEqualItems_DiffLength_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.IsEqualItems(newTestTrace("pkg1", 10))

	// Assert
	if result {
		t.Error("should not be equal")
	}
}

func Test_TraceCollection_ConcatNewPtr_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()
	trace := newTestTrace("extra", 99)

	// Act
	result := collection.ConcatNewPtr(&trace)

	// Assert
	if result.Length() != 4 {
		t.Errorf("expected 4, got %d", result.Length())
	}
}

func Test_TraceCollection_ConcatNewUsingSkipPlusCount_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.ConcatNewUsingSkipPlusCount(0, 3)

	// Assert
	if result.Length() < 3 {
		t.Errorf("expected at least 3, got %d", result.Length())
	}
}

func Test_TraceCollection_ConcatNewUsingSkip_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.ConcatNewUsingSkip(0)

	// Assert
	if result.Length() < 3 {
		t.Errorf("expected at least 3, got %d", result.Length())
	}
}

func Test_TraceCollection_InsertAt_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	inserted := newTestTrace("inserted", 99)
	collection.InsertAt(1, inserted)

	// Assert
	if collection.Items[1].PackageName != "inserted" {
		t.Error("item should be inserted at index 1")
	}
}

func Test_TraceCollection_AddsPtr_Valid_Ext(t *testing.T) {
	// Arrange
	collection := emptyCollection()
	trace := newTestTrace("valid", 1)

	// Act
	collection.AddsPtr(false, &trace)

	// Assert
	if collection.Length() != 1 {
		t.Errorf("expected 1, got %d", collection.Length())
	}
}

func Test_TraceCollection_Adds_Empty_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	collection.Adds()

	// Assert
	if collection.Length() != 3 {
		t.Errorf("expected 3, got %d", collection.Length())
	}
}

func Test_TraceCollection_AddsPtr_Empty_Ext(t *testing.T) {
	// Arrange
	collection := emptyCollection()

	// Act
	collection.AddsPtr(false)

	// Assert
	if collection.Length() != 0 {
		t.Errorf("expected 0, got %d", collection.Length())
	}
}

func Test_TraceCollection_StringsUsingFmt_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.StringsUsingFmt(func(trace *codestack.Trace) string {
		return trace.PackageName
	})

	// Assert
	if len(result) != 3 {
		t.Errorf("expected 3, got %d", len(result))
	}
}

func Test_TraceCollection_AddsUsingSkipDefault_Ext(t *testing.T) {
	// Arrange
	collection := emptyCollection()

	// Act
	collection.AddsUsingSkipDefault(0)

	// Assert - should have items from current stack
	if collection.Length() == 0 {
		t.Error("should have items from stack")
	}
}

func Test_TraceCollection_AddsUsingSkipUsingFilter_Ext(t *testing.T) {
	// Arrange
	collection := emptyCollection()

	// Act
	collection.AddsUsingSkipUsingFilter(
		true,
		true,
		0,
		10,
		func(trace *codestack.Trace) (bool, bool) {
			return true, false
		},
	)

	// Assert
	if collection.Length() == 0 {
		t.Error("should have items from filtered stack")
	}
}

// =============================================================================
// Trace: additional Json methods via reflection
// =============================================================================

func Test_Trace_Json_Ext(t *testing.T) {
	// Arrange
	trace := codestack.Trace{
		PackageName:       "pkg",
		MethodName:        "Method",
		PackageMethodName: "pkg.Method",
		FilePath:          "/file.go",
		Line:              10,
		IsOkay:            true,
	}

	// Act
	result := trace.Json()

	// Assert
	if result.HasError() {
		t.Errorf("should not have error: %v", result.MeaningfulError())
	}
}

func Test_Trace_JsonPtr_Ext(t *testing.T) {
	// Arrange
	trace := codestack.Trace{
		PackageName: "pkg",
		FilePath:    "/file.go",
		Line:        10,
	}

	// Act
	result := trace.JsonPtr()

	// Assert
	if result == nil {
		t.Error("should not be nil")
	}
}

func Test_Trace_JsonModel_Ext(t *testing.T) {
	// Arrange
	trace := codestack.Trace{PackageName: "pkg"}

	// Act
	model := trace.JsonModel()

	// Assert
	if model.PackageName != "pkg" {
		t.Error("should return same trace as model")
	}
}

func Test_Trace_JsonModelAny_Ext(t *testing.T) {
	// Arrange
	trace := &codestack.Trace{PackageName: "pkg"}

	// Act
	result := trace.JsonModelAny()

	// Assert
	if result == nil {
		t.Error("should not be nil")
	}
}

func Test_Trace_JsonString_Ext(t *testing.T) {
	// Arrange
	trace := &codestack.Trace{
		PackageName: "pkg",
		FilePath:    "/file.go",
		Line:        10,
	}

	// Act
	result := trace.JsonString()

	// Assert
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_Trace_ParseInjectUsingJson_Ext(t *testing.T) {
	// Arrange
	original := codestack.Trace{
		PackageName:       "pkg",
		PackageMethodName: "pkg.M",
		FilePath:          "/f.go",
		Line:              10,
	}
	jsonResult := original.JsonPtr()

	// Act
	var target codestack.Trace
	result, err := target.ParseInjectUsingJson(jsonResult)

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if result.PackageName != "pkg" {
		t.Error("should preserve PackageName")
	}
}

func Test_Trace_ParseInjectUsingJsonMust_Ext(t *testing.T) {
	// Arrange
	original := codestack.Trace{PackageName: "pkg", FilePath: "/f.go"}
	jsonResult := original.JsonPtr()

	// Act
	var target codestack.Trace
	result := target.ParseInjectUsingJsonMust(jsonResult)

	// Assert
	if result == nil {
		t.Error("should not be nil")
	}
}

func Test_Trace_JsonParseSelfInject_Ext(t *testing.T) {
	// Arrange
	original := codestack.Trace{PackageName: "pkg", FilePath: "/f.go"}
	jsonResult := original.JsonPtr()

	// Act
	var target codestack.Trace
	err := target.JsonParseSelfInject(jsonResult)

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func Test_Trace_NilDispose_Ext(t *testing.T) {
	// Arrange
	var trace *codestack.Trace

	// Act - should not panic
	trace.Dispose()
}

// =============================================================================
// FileWithLine: additional methods
// =============================================================================

func Test_FileWithLine_JsonModel_Ext(t *testing.T) {
	// Arrange
	fwl := codestack.FileWithLine{FilePath: "/f.go", Line: 5}

	// Act
	model := fwl.JsonModel()

	// Assert
	if model.FilePath != "/f.go" {
		t.Error("should return same model")
	}
}

func Test_FileWithLine_JsonModelAny_Ext(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{FilePath: "/f.go", Line: 5}

	// Act
	result := fwl.JsonModelAny()

	// Assert
	if result == nil {
		t.Error("should not be nil")
	}
}

func Test_FileWithLine_Json_Ext(t *testing.T) {
	// Arrange
	fwl := codestack.FileWithLine{FilePath: "/f.go", Line: 5}

	// Act
	result := fwl.Json()

	// Assert
	if result.HasError() {
		t.Errorf("error: %v", result.Err())
	}
}

func Test_FileWithLine_JsonPtr_Ext(t *testing.T) {
	// Arrange
	fwl := codestack.FileWithLine{FilePath: "/f.go", Line: 5}

	// Act
	result := fwl.JsonPtr()

	// Assert
	if result == nil {
		t.Error("should not be nil")
	}
}

func Test_FileWithLine_ParseInjectUsingJson_Ext(t *testing.T) {
	// Arrange
	original := codestack.FileWithLine{FilePath: "/f.go", Line: 5}
	jsonResult := original.JsonPtr()

	// Act
	var target codestack.FileWithLine
	result, err := target.ParseInjectUsingJson(jsonResult)

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if result.FilePath != "/f.go" {
		t.Error("should preserve FilePath")
	}
}

func Test_FileWithLine_ParseInjectUsingJsonMust_Ext(t *testing.T) {
	// Arrange
	original := codestack.FileWithLine{FilePath: "/f.go", Line: 5}
	jsonResult := original.JsonPtr()

	// Act
	var target codestack.FileWithLine
	result := target.ParseInjectUsingJsonMust(jsonResult)

	// Assert
	if result == nil {
		t.Error("should not be nil")
	}
}

func Test_FileWithLine_JsonParseSelfInject_Ext(t *testing.T) {
	// Arrange
	original := codestack.FileWithLine{FilePath: "/f.go", Line: 5}
	jsonResult := original.JsonPtr()

	// Act
	var target codestack.FileWithLine
	err := target.JsonParseSelfInject(jsonResult)

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func Test_FileWithLine_FileWithLine_Method_Ext(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{FilePath: "/f.go", Line: 42}

	// Act
	result := fwl.FileWithLine()

	// Assert
	if result != "/f.go:42" {
		t.Errorf("expected '/f.go:42', got '%s'", result)
	}
}

// =============================================================================
// Trace: Message caching via reflection (tests lazy once)
// =============================================================================

func Test_Trace_Message_CachesOnSecondCall_Ext(t *testing.T) {
	// Arrange
	trace := &codestack.Trace{
		PackageName:       "pkg",
		MethodName:        "Method",
		PackageMethodName: "pkg.Method",
		FilePath:          "/file.go",
		Line:              10,
		IsOkay:            true,
	}

	// Act
	msg1 := trace.Message()
	msg2 := trace.Message() // cached path

	// Assert
	if msg1 != msg2 {
		t.Error("Message should return same value on second call")
	}
}

func Test_Trace_ShortString_CachesOnSecondCall_Ext(t *testing.T) {
	// Arrange
	trace := &codestack.Trace{
		PackageName:       "pkg",
		MethodName:        "Method",
		PackageMethodName: "pkg.Method",
		FilePath:          "/file.go",
		Line:              10,
		IsOkay:            true,
	}

	// Act
	s1 := trace.ShortString()
	s2 := trace.ShortString() // cached path

	// Assert
	if s1 != s2 {
		t.Error("ShortString should return same value on second call")
	}
}
