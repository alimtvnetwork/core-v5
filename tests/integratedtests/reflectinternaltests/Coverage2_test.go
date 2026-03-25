package reflectinternaltests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

// ── getFunc ──

func Test_Cov2_GetFunc_FullName(t *testing.T) {
	fn := func() {}
	fullName := reflectinternal.GetFunc.FullName(fn)
	nilName := reflectinternal.GetFunc.FullName(nil)

	actual := args.Map{
		"hasName": fullName != "",
		"nilName": nilName,
	}
	expected := args.Map{
		"hasName": true,
		"nilName": "",
	}
	expected.ShouldBeEqual(t, 0, "GetFunc_FullName returns correct value -- with args", actual)
}

func Test_Cov2_GetFunc_FullNameWithName(t *testing.T) {
	fn := func() {}
	fullName, name := reflectinternal.GetFunc.FullNameWithName(fn)
	nilFull, nilName := reflectinternal.GetFunc.FullNameWithName(nil)

	actual := args.Map{
		"hasFullName": fullName != "",
		"hasName":     name != "",
		"nilFull":     nilFull,
		"nilName":     nilName,
	}
	expected := args.Map{
		"hasFullName": true,
		"hasName":     true,
		"nilFull":     "",
		"nilName":     "",
	}
	expected.ShouldBeEqual(t, 0, "GetFunc_FullNameWithName returns non-empty -- with args", actual)
}

func Test_Cov2_GetFunc_NameOnly(t *testing.T) {
	fn := func() {}
	name := reflectinternal.GetFunc.NameOnly(fn)
	nilName := reflectinternal.GetFunc.NameOnly(nil)

	actual := args.Map{
		"hasName": name != "",
		"nilName": nilName,
	}
	expected := args.Map{
		"hasName": true,
		"nilName": "",
	}
	expected.ShouldBeEqual(t, 0, "GetFunc_NameOnly returns correct value -- with args", actual)
}

func Test_Cov2_GetFunc_All(t *testing.T) {
	full, pkg, method := reflectinternal.GetFunc.All("mypackage.MyFunc")
	emFull, emPkg, emMethod := reflectinternal.GetFunc.All("")

	actual := args.Map{
		"full":     full,
		"pkg":      pkg,
		"method":   method,
		"emFull":   emFull,
		"emPkg":    emPkg,
		"emMethod": emMethod,
	}
	expected := args.Map{
		"full":     "mypackage.MyFunc",
		"pkg":      "mypackage",
		"method":   "MyFunc",
		"emFull":   "",
		"emPkg":    "",
		"emMethod": "",
	}
	expected.ShouldBeEqual(t, 0, "GetFunc_All returns correct value -- with args", actual)
}

func Test_Cov2_GetFunc_FuncDirectInvokeName(t *testing.T) {
	fn := func() {}
	name := reflectinternal.GetFunc.FuncDirectInvokeName(fn)
	emptyName := reflectinternal.GetFunc.FuncDirectInvokeNameUsingFullName("")

	actual := args.Map{
		"hasName":  name != "",
		"emptyRes": emptyName,
	}
	expected := args.Map{
		"hasName":  true,
		"emptyRes": "",
	}
	expected.ShouldBeEqual(t, 0, "GetFunc_FuncDirectInvokeName returns correct value -- with args", actual)
}

func Test_Cov2_GetFunc_PascalFuncName(t *testing.T) {
	actual := args.Map{
		"simple":  reflectinternal.GetFunc.PascalFuncName("hello"),
		"single":  reflectinternal.GetFunc.PascalFuncName("h"),
		"empty":   reflectinternal.GetFunc.PascalFuncName(""),
		"already": reflectinternal.GetFunc.PascalFuncName("Hello"),
	}
	expected := args.Map{
		"simple":  "Hello",
		"single":  "H",
		"empty":   "",
		"already": "Hello",
	}
	expected.ShouldBeEqual(t, 0, "GetFunc_PascalFuncName returns correct value -- with args", actual)
}

func Test_Cov2_GetFunc_GetPkgPathFullName(t *testing.T) {
	result := reflectinternal.GetFunc.GetPkgPathFullName("github.com/org/repo/pkg.Func")
	empty := reflectinternal.GetFunc.GetPkgPathFullName("")
	noslash := reflectinternal.GetFunc.GetPkgPathFullName("pkg.Func")

	actual := args.Map{
		"result":  result != "",
		"empty":   empty,
		"noslash": noslash,
	}
	expected := args.Map{
		"result":  true,
		"empty":   "",
		"noslash": "pkg.Func",
	}
	expected.ShouldBeEqual(t, 0, "GetFunc_GetPkgPathFullName returns correct value -- with args", actual)
}

func Test_Cov2_GetFunc_RunTime(t *testing.T) {
	fn := func() {}
	rt := reflectinternal.GetFunc.RunTime(fn)
	nilRt := reflectinternal.GetFunc.RunTime(nil)
	intRt := reflectinternal.GetFunc.RunTime(42)

	actual := args.Map{
		"rtNotNil": rt != nil,
		"nilRt":    nilRt == nil,
		"intRt":    intRt == nil,
	}
	expected := args.Map{
		"rtNotNil": true,
		"nilRt":    true,
		"intRt":    true,
	}
	expected.ShouldBeEqual(t, 0, "GetFunc_RunTime returns correct value -- with args", actual)
}

func Test_Cov2_GetFunc_GetMethod(t *testing.T) {
	type myStruct struct{}
	s := myStruct{}

	nilMethod := reflectinternal.GetFunc.GetMethod("", s)
	missingMethod := reflectinternal.GetFunc.GetMethod("NotExist", s)
	nilItem := reflectinternal.GetFunc.GetMethod("Name", nil)

	actual := args.Map{
		"nilMethod":     nilMethod == nil,
		"missingMethod": missingMethod == nil,
		"nilItem":       nilItem == nil,
	}
	expected := args.Map{
		"nilMethod":     true,
		"missingMethod": true,
		"nilItem":       true,
	}
	expected.ShouldBeEqual(t, 0, "GetFunc_GetMethod returns correct value -- with args", actual)
}

func Test_Cov2_GetFunc_GetMethods(t *testing.T) {
	type myStruct struct{}
	s := myStruct{}
	methods := reflectinternal.GetFunc.GetMethods(s)
	nilMethods := reflectinternal.GetFunc.GetMethods(nil)

	actual := args.Map{
		"methodsNotNil":    methods != nil,
		"nilMethodsNotNil": nilMethods != nil,
	}
	expected := args.Map{
		"methodsNotNil":    true,
		"nilMethodsNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "GetFunc_GetMethods returns correct value -- with args", actual)
}

func Test_Cov2_GetFunc_GetMethodsNames(t *testing.T) {
	type myStruct struct{}
	s := myStruct{}
	names := reflectinternal.GetFunc.GetMethodsNames(s)
	nilNames := reflectinternal.GetFunc.GetMethodsNames(nil)

	actual := args.Map{
		"namesNotNil":    names != nil,
		"nilNamesNotNil": nilNames != nil,
	}
	expected := args.Map{
		"namesNotNil":    names != nil,
		"nilNamesNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "GetFunc_GetMethodsNames returns correct value -- with args", actual)
}

func Test_Cov2_GetFunc_GetMethodsMap(t *testing.T) {
	type myStruct struct{}
	s := myStruct{}
	m := reflectinternal.GetFunc.GetMethodsMap(s)
	nilMap := reflectinternal.GetFunc.GetMethodsMap(nil)

	actual := args.Map{
		"mapNotNil":    m != nil,
		"nilMapNotNil": nilMap != nil,
	}
	expected := args.Map{
		"mapNotNil":    true,
		"nilMapNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "GetFunc_GetMethodsMap returns correct value -- with args", actual)
}

// ── reflectUtils ──

func Test_Cov2_Utils_MaxLimit(t *testing.T) {
	actual := args.Map{
		"noMax":     reflectinternal.Utils.MaxLimit(10, -1),
		"underMax":  reflectinternal.Utils.MaxLimit(5, 10),
		"overMax":   reflectinternal.Utils.MaxLimit(15, 10),
		"equalMax":  reflectinternal.Utils.MaxLimit(10, 10),
	}
	expected := args.Map{
		"noMax":     10,
		"underMax":  5,
		"overMax":   10,
		"equalMax":  10,
	}
	expected.ShouldBeEqual(t, 0, "Utils_MaxLimit returns correct value -- with args", actual)
}

func Test_Cov2_Utils_AppendArgs(t *testing.T) {
	emptyResult := reflectinternal.Utils.AppendArgs("item", []any{})
	withResult := reflectinternal.Utils.AppendArgs("item", []any{"a", "b"})

	actual := args.Map{
		"emptyLen": len(emptyResult),
		"withLen":  len(withResult),
		"first":    emptyResult[0],
	}
	expected := args.Map{
		"emptyLen": 1,
		"withLen":  3,
		"first":    "item",
	}
	expected.ShouldBeEqual(t, 0, "Utils_AppendArgs returns correct value -- with args", actual)
}

func Test_Cov2_Utils_PkgNameOnly(t *testing.T) {
	fn := func() {}
	name := reflectinternal.Utils.PkgNameOnly(fn)

	actual := args.Map{
		"hasName": name != "",
	}
	expected := args.Map{
		"hasName": true,
	}
	expected.ShouldBeEqual(t, 0, "Utils_PkgNameOnly returns correct value -- with args", actual)
}

func Test_Cov2_Utils_FullNameToPkgName(t *testing.T) {
	name := reflectinternal.Utils.FullNameToPkgName("mypackage.MyFunc")

	actual := args.Map{
		"name": name,
	}
	expected := args.Map{
		"name": "mypackage",
	}
	expected.ShouldBeEqual(t, 0, "Utils_FullNameToPkgName returns correct value -- with args", actual)
}

func Test_Cov2_Utils_IsReflectTypeMatch(t *testing.T) {
	intType := reflect.TypeOf(0)
	strType := reflect.TypeOf("")
	anyType := reflect.TypeOf((*any)(nil)).Elem()

	ok1, err1 := reflectinternal.Utils.IsReflectTypeMatch(intType, intType)
	ok2, err2 := reflectinternal.Utils.IsReflectTypeMatch(intType, strType)
	ok3, err3 := reflectinternal.Utils.IsReflectTypeMatch(anyType, strType)

	actual := args.Map{
		"sameOk":   ok1,
		"sameErr":  err1 == nil,
		"diffOk":   ok2,
		"diffErr":  err2 != nil,
		"anyOk":    ok3,
		"anyNoErr": err3 == nil,
	}
	expected := args.Map{
		"sameOk":   true,
		"sameErr":  true,
		"diffOk":   false,
		"diffErr":  true,
		"anyOk":    true,
		"anyNoErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Utils_IsReflectTypeMatch returns correct value -- with args", actual)
}

func Test_Cov2_Utils_IsReflectTypeMatchAny(t *testing.T) {
	ok1, err1 := reflectinternal.Utils.IsReflectTypeMatchAny(42, 99)
	ok2, err2 := reflectinternal.Utils.IsReflectTypeMatchAny(42, "hello")

	actual := args.Map{
		"sameOk":  ok1,
		"sameErr": err1 == nil,
		"diffOk":  ok2,
		"diffErr": err2 != nil,
	}
	expected := args.Map{
		"sameOk":  true,
		"sameErr": true,
		"diffOk":  false,
		"diffErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Utils_IsReflectTypeMatchAny returns correct value -- with args", actual)
}

func Test_Cov2_Utils_VerifyReflectTypesAny(t *testing.T) {
	ok1, err1 := reflectinternal.Utils.VerifyReflectTypesAny(
		[]any{42, "hello"},
		[]any{99, "world"},
	)
	ok2, err2 := reflectinternal.Utils.VerifyReflectTypesAny(
		[]any{42},
		[]any{42, "extra"},
	)
	ok3, err3 := reflectinternal.Utils.VerifyReflectTypesAny(
		[]any{42, "hello"},
		[]any{99, 100},
	)

	actual := args.Map{
		"matchOk":     ok1,
		"matchErr":    err1 == nil,
		"lenMismatch": ok2,
		"lenErr":      err2 != nil,
		"typeMismatch": ok3,
		"typeErr":     err3 != nil,
	}
	expected := args.Map{
		"matchOk":     true,
		"matchErr":    true,
		"lenMismatch": false,
		"lenErr":      true,
		"typeMismatch": false,
		"typeErr":     true,
	}
	expected.ShouldBeEqual(t, 0, "Utils_VerifyReflectTypesAny returns correct value -- with args", actual)
}

func Test_Cov2_Utils_VerifyReflectTypes(t *testing.T) {
	intType := reflect.TypeOf(0)
	strType := reflect.TypeOf("")

	ok1, err1 := reflectinternal.Utils.VerifyReflectTypes(
		"test",
		[]reflect.Type{intType, strType},
		[]reflect.Type{intType, strType},
	)
	ok2, err2 := reflectinternal.Utils.VerifyReflectTypes(
		"test",
		[]reflect.Type{intType},
		[]reflect.Type{intType, strType},
	)

	actual := args.Map{
		"matchOk":  ok1,
		"matchErr": err1 == nil,
		"diffOk":   ok2,
		"diffErr":  err2 != nil,
	}
	expected := args.Map{
		"matchOk":  true,
		"matchErr": true,
		"diffOk":   false,
		"diffErr":  true,
	}
	expected.ShouldBeEqual(t, 0, "Utils_VerifyReflectTypes returns correct value -- with args", actual)
}

// ── reflectGetter additional ──

func Test_Cov2_ReflectGetter_FieldNameWithValuesMap(t *testing.T) {
	type testStruct struct {
		Name string
		age  int
	}
	s := testStruct{Name: "Alice", age: 30}

	result, err := reflectinternal.ReflectGetter.FieldNameWithValuesMap(s)
	nilResult, nilErr := reflectinternal.ReflectGetter.FieldNameWithValuesMap(nil)

	actual := args.Map{
		"resultLen": len(result),
		"noError":   err == nil,
		"nilLen":    len(nilResult),
		"nilHasErr": nilErr != nil,
	}
	expected := args.Map{
		"resultLen": 2,
		"noError":   true,
		"nilLen":    0,
		"nilHasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ReflectGetter_FieldNameWithValuesMap returns non-empty -- with args", actual)
}

func Test_Cov2_ReflectGetter_NullFieldsMap(t *testing.T) {
	type testStruct struct {
		Name *string
		Age  int
	}
	s := testStruct{Name: nil, Age: 0}

	result := reflectinternal.ReflectGetter.NullFieldsMap(s)
	nilResult := reflectinternal.ReflectGetter.NullFieldsMap(nil)

	actual := args.Map{
		"hasNullName": result["Name"],
		"nilLen":      len(nilResult),
	}
	expected := args.Map{
		"hasNullName": true,
		"nilLen":      0,
	}
	expected.ShouldBeEqual(t, 0, "ReflectGetter_NullFieldsMap returns correct value -- with args", actual)
}

func Test_Cov2_ReflectGetter_NullOrZeroFieldsMap(t *testing.T) {
	type testStruct struct {
		Name string
		Age  int
	}
	s := testStruct{Name: "", Age: 0}

	result := reflectinternal.ReflectGetter.NullOrZeroFieldsMap(s)
	nilResult := reflectinternal.ReflectGetter.NullOrZeroFieldsMap(nil)

	actual := args.Map{
		"hasZeroName": result["Name"],
		"hasZeroAge":  result["Age"],
		"nilLen":      len(nilResult),
	}
	expected := args.Map{
		"hasZeroName": true,
		"hasZeroAge":  true,
		"nilLen":      0,
	}
	expected.ShouldBeEqual(t, 0, "ReflectGetter_NullOrZeroFieldsMap returns correct value -- with args", actual)
}

// ── reflectGetUsingReflectValue additional ──

func Test_Cov2_ReflectGetRv_FieldNameWithTypeMap(t *testing.T) {
	type testStruct struct {
		Name string
		Age  int
	}
	s := testStruct{Name: "Alice", Age: 30}
	rv := reflect.ValueOf(s)
	result := reflectinternal.ReflectGetterUsingReflectValue.FieldNameWithTypeMap(rv)

	nilResult := reflectinternal.ReflectGetterUsingReflectValue.FieldNameWithTypeMap(
		reflect.ValueOf(42),
	)

	actual := args.Map{
		"resultLen": len(result),
		"nilResult": nilResult == nil,
	}
	expected := args.Map{
		"resultLen": 2,
		"nilResult": true,
	}
	expected.ShouldBeEqual(t, 0, "ReflectGetRv_FieldNameWithTypeMap returns non-empty -- with args", actual)
}

func Test_Cov2_ReflectGetRv_FieldNameWithValuesMap(t *testing.T) {
	type testStruct struct {
		Name string
		age  int
	}
	s := testStruct{Name: "Alice", age: 30}
	rv := reflect.ValueOf(s)

	result, err := reflectinternal.ReflectGetterUsingReflectValue.FieldNameWithValuesMap(rv)

	actual := args.Map{
		"resultLen": len(result),
		"noError":   err == nil,
		"hasName":   result["Name"] == "Alice",
	}
	expected := args.Map{
		"resultLen": 2,
		"noError":   true,
		"hasName":   true,
	}
	expected.ShouldBeEqual(t, 0, "ReflectGetRv_FieldNameWithValuesMap returns non-empty -- with args", actual)
}

// ── MapConverter additional ──

func Test_Cov2_MapConverter_ToKeysStrings(t *testing.T) {
	result, err := reflectinternal.MapConverter.ToKeysStrings(map[string]int{"a": 1})

	actual := args.Map{
		"len":   len(result),
		"noErr": err == nil,
	}
	expected := args.Map{
		"len":   1,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "MapConverter_ToKeysStrings returns correct value -- with args", actual)
}

func Test_Cov2_MapConverter_ToValuesAny(t *testing.T) {
	result, err := reflectinternal.MapConverter.ToValuesAny(map[string]int{"a": 1})
	nilResult, nilErr := reflectinternal.MapConverter.ToValuesAny(nil)

	actual := args.Map{
		"len":    len(result),
		"noErr":  err == nil,
		"nilLen": len(nilResult),
		"nilErr": nilErr == nil,
	}
	expected := args.Map{
		"len":    1,
		"noErr":  true,
		"nilLen": 0,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "MapConverter_ToValuesAny returns non-empty -- with args", actual)
}

func Test_Cov2_MapConverter_ToKeysAny(t *testing.T) {
	result, err := reflectinternal.MapConverter.ToKeysAny(map[string]int{"a": 1})
	nilResult, nilErr := reflectinternal.MapConverter.ToKeysAny(nil)

	actual := args.Map{
		"len":    len(result),
		"noErr":  err == nil,
		"nilLen": len(nilResult),
		"nilErr": nilErr == nil,
	}
	expected := args.Map{
		"len":    1,
		"noErr":  true,
		"nilLen": 0,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "MapConverter_ToKeysAny returns correct value -- with args", actual)
}

func Test_Cov2_MapConverter_ToStringsMust(t *testing.T) {
	result := reflectinternal.MapConverter.ToStringsMust(map[string]int{"a": 1})

	actual := args.Map{
		"len": len(result),
	}
	expected := args.Map{
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "MapConverter_ToStringsMust returns correct value -- with args", actual)
}

func Test_Cov2_MapConverter_ToSortedStringsMust(t *testing.T) {
	result := reflectinternal.MapConverter.ToSortedStringsMust(map[string]int{"b": 2, "a": 1})
	nilResult := reflectinternal.MapConverter.ToSortedStringsMust(nil)

	actual := args.Map{
		"len":    len(result),
		"first":  result[0],
		"nilLen": len(nilResult),
	}
	expected := args.Map{
		"len":    2,
		"first":  "a",
		"nilLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "MapConverter_ToSortedStringsMust returns correct value -- with args", actual)
}

func Test_Cov2_MapConverter_ToMapStringAnyRv_NonStringKey(t *testing.T) {
	m := map[int]string{1: "one", 2: "two"}
	rv := reflect.ValueOf(m)
	result, err := reflectinternal.MapConverter.ToMapStringAnyRv(rv)

	actual := args.Map{
		"len":   len(result),
		"noErr": err == nil,
	}
	expected := args.Map{
		"len":   2,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "MapConverter_ToMapStringAnyRv_NonStringKey returns correct value -- with args", actual)
}

func Test_Cov2_MapConverter_ToMapStringAnyRv_NotMap(t *testing.T) {
	rv := reflect.ValueOf(42)
	_, err := reflectinternal.MapConverter.ToMapStringAnyRv(rv)

	actual := args.Map{
		"hasErr": err != nil,
	}
	expected := args.Map{
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "MapConverter_ToMapStringAnyRv_NotMap returns correct value -- with args", actual)
}

func Test_Cov2_MapConverter_ToStringsRv_NotMap(t *testing.T) {
	rv := reflect.ValueOf(42)
	_, err := reflectinternal.MapConverter.ToStringsRv(rv)

	actual := args.Map{
		"hasErr": err != nil,
	}
	expected := args.Map{
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "MapConverter_ToStringsRv_NotMap returns correct value -- with args", actual)
}

// ── reflectPath ──

func Test_Cov2_Path_RepoDir(t *testing.T) {
	result := reflectinternal.Path.RepoDir()

	actual := args.Map{
		"notEmpty": result != "",
	}
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Path_RepoDir returns correct value -- with args", actual)
}

func Test_Cov2_Path_CurDir(t *testing.T) {
	result := reflectinternal.Path.CurDir()

	actual := args.Map{
		"notEmpty": result != "",
	}
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Path_CurDir returns correct value -- with args", actual)
}

// ── TypeNamesReferenceString ──

func Test_Cov2_TypeNamesReferenceString(t *testing.T) {
	result := reflectinternal.TypeNamesReferenceString(true, 42, "hello")

	actual := args.Map{
		"notEmpty": result != "",
	}
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "TypeNamesReferenceString returns correct value -- with args", actual)
}

// ── TypeNamesString ──

func Test_Cov2_TypeNamesString(t *testing.T) {
	result := reflectinternal.TypeNamesString(true, 42, "hello")

	actual := args.Map{
		"notEmpty": result != "",
	}
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "TypeNamesString returns correct value -- with args", actual)
}

// ── ReflectType converter ──

func Test_Cov2_ReflectType_SafeName(t *testing.T) {
	result := reflectinternal.ReflectType.SafeName(42)
	nilResult := reflectinternal.ReflectType.SafeName(nil)

	actual := args.Map{
		"result":    result,
		"nilResult": nilResult,
	}
	expected := args.Map{
		"result":    "int",
		"nilResult": "",
	}
	expected.ShouldBeEqual(t, 0, "ReflectType_SafeName returns correct value -- with args", actual)
}
