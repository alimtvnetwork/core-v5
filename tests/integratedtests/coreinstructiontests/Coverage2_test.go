package coreinstructiontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coreinstruction"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// ── FlatSpecification ──

func Test_Cov2_FlatSpec_InvalidFlatSpecification(t *testing.T) {
	flat := coreinstruction.InvalidFlatSpecification()
	actual := args.Map{"valid": flat.IsValid, "id": flat.Id}
	expected := args.Map{"valid": false, "id": ""}
	expected.ShouldBeEqual(t, 0, "InvalidFlatSpecification", actual)
}

func Test_Cov2_FlatSpec_BaseAccessors(t *testing.T) {
	spec := coreinstruction.NewSpecification("id1", "disp", "tp", []string{"t1"}, true)
	flat := spec.FlatSpecification()
	actual := args.Map{
		"baseId":       flat.BaseIdentifier().Id,
		"baseDisplay":  flat.BaseDisplay().Display,
		"baseType":     flat.BaseType().Type,
		"baseIsGlobal": flat.BaseIsGlobal().IsGlobal,
		"tagsLen":      len(flat.BaseTags().Tags),
	}
	expected := args.Map{
		"baseId":       "id1",
		"baseDisplay":  "disp",
		"baseType":     "tp",
		"baseIsGlobal": true,
		"tagsLen":      1,
	}
	expected.ShouldBeEqual(t, 0, "FlatSpec_BaseAccessors", actual)
}

func Test_Cov2_FlatSpec_SpecCaching(t *testing.T) {
	flat := &coreinstruction.FlatSpecification{Id: "x", Display: "d", Type: "t"}
	s1 := flat.Spec()
	s2 := flat.Spec()
	actual := args.Map{"same": s1 == s2, "id": s1.Id}
	expected := args.Map{"same": true, "id": "x"}
	expected.ShouldBeEqual(t, 0, "FlatSpec_SpecCaching", actual)
}

func Test_Cov2_FlatSpec_NilSpec(t *testing.T) {
	var flat *coreinstruction.FlatSpecification
	actual := args.Map{"isNil": flat.Spec() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "FlatSpec_NilSpec", actual)
}

func Test_Cov2_FlatSpec_Clone(t *testing.T) {
	flat := &coreinstruction.FlatSpecification{Id: "x", Tags: []string{"a"}}
	cloned := flat.Clone()
	var nilFlat *coreinstruction.FlatSpecification
	actual := args.Map{
		"id":       cloned.Id,
		"tagsLen":  len(cloned.Tags),
		"nilClone": nilFlat.Clone() == nil,
	}
	expected := args.Map{
		"id":       "x",
		"tagsLen":  1,
		"nilClone": true,
	}
	expected.ShouldBeEqual(t, 0, "FlatSpec_Clone", actual)
}

// ── StringCompare ──

func Test_Cov2_StringCompare_Constructors(t *testing.T) {
	eq := coreinstruction.NewStringCompareEqual("abc", "abc")
	contains := coreinstruction.NewStringCompareContains(false, "bc", "abc")
	starts := coreinstruction.NewStringCompareStartsWith(false, "ab", "abc")
	ends := coreinstruction.NewStringCompareEndsWith(false, "bc", "abc")
	regex := coreinstruction.NewStringCompareRegex(`\w+`, "abc")

	actual := args.Map{
		"eqMatch":       eq.IsMatch(),
		"containsMatch": contains.IsMatch(),
		"startsMatch":   starts.IsMatch(),
		"endsMatch":     ends.IsMatch(),
		"regexMatch":    regex.IsMatch(),
	}
	expected := args.Map{
		"eqMatch":       true,
		"containsMatch": true,
		"startsMatch":   true,
		"endsMatch":     true,
		"regexMatch":    true,
	}
	expected.ShouldBeEqual(t, 0, "StringCompare_Constructors", actual)
}

func Test_Cov2_StringCompare_NilMethods(t *testing.T) {
	var sc *coreinstruction.StringCompare
	actual := args.Map{
		"isInvalid":  sc.IsInvalid(),
		"isDefined":  sc.IsDefined(),
		"isMatch":    sc.IsMatch(),
		"matchFail":  sc.IsMatchFailed(),
		"verifyErr":  sc.VerifyError() == nil,
	}
	expected := args.Map{
		"isInvalid":  true,
		"isDefined":  false,
		"isMatch":    true,
		"matchFail":  false,
		"verifyErr":  true,
	}
	expected.ShouldBeEqual(t, 0, "StringCompare_Nil", actual)
}

func Test_Cov2_StringCompare_VerifyError_Regex(t *testing.T) {
	sc := coreinstruction.NewStringCompareRegex(`\w+`, "abc")
	actual := args.Map{"noErr": sc.VerifyError() == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyError_Regex", actual)
}

func Test_Cov2_StringCompare_VerifyError_NonRegex(t *testing.T) {
	sc := coreinstruction.NewStringCompare(stringcompareas.Equal, false, "abc", "abc")
	actual := args.Map{"noErr": sc.VerifyError() == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyError_NonRegex", actual)
}

// ── StringSearch ──

func Test_Cov2_StringSearch_Methods(t *testing.T) {
	var nilSS *coreinstruction.StringSearch
	actual := args.Map{
		"nilIsEmpty": nilSS.IsEmpty(),
		"nilIsExist": nilSS.IsExist(),
		"nilHas":     nilSS.Has(),
		"nilMatch":   nilSS.IsMatch("anything"),
		"nilVerify":  nilSS.VerifyError("anything") == nil,
	}
	expected := args.Map{
		"nilIsEmpty": true,
		"nilIsExist": false,
		"nilHas":     false,
		"nilMatch":   true,
		"nilVerify":  true,
	}
	expected.ShouldBeEqual(t, 0, "StringSearch_Nil", actual)
}

func Test_Cov2_StringSearch_IsAllMatch(t *testing.T) {
	var nilSS *coreinstruction.StringSearch
	actual := args.Map{
		"emptyMatch":   nilSS.IsAllMatch(),
		"anyMatchFail": nilSS.IsAnyMatchFailed("a"),
	}
	expected := args.Map{
		"emptyMatch":   true,
		"anyMatchFail": false,
	}
	expected.ShouldBeEqual(t, 0, "StringSearch_AllMatch", actual)
}

// ── NameList / NameListCollection ──

func Test_Cov2_NameList_DeepClone(t *testing.T) {
	nl := &coreinstruction.NameList{Name: "test", List: corestr.NewSimpleSlicePtr("a", "b")}
	cloned := nl.DeepClone()
	actual := args.Map{"name": cloned.Name, "notNil": cloned.List != nil}
	expected := args.Map{"name": "test", "notNil": true}
	expected.ShouldBeEqual(t, 0, "NameList_DeepClone", actual)
}

func Test_Cov2_NameList_String(t *testing.T) {
	nl := coreinstruction.NameList{Name: "test"}
	actual := args.Map{"notEmpty": nl.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameList_String", actual)
}

func Test_Cov2_NameListCollection(t *testing.T) {
	var nilNlc *coreinstruction.NameListCollection
	nlc := &coreinstruction.NameListCollection{NameLists: []coreinstruction.NameList{{Name: "a"}}}

	actual := args.Map{
		"nilIsNull":  nilNlc.IsNull(),
		"nilIsAny":   nilNlc.IsAnyNull(),
		"nilIsEmpty": nilNlc.IsEmpty(),
		"nilLength":  nilNlc.Length(),
		"length":     nlc.Length(),
		"hasAny":     nlc.HasAnyItem(),
		"notEmpty":   nlc.String() != "",
	}
	expected := args.Map{
		"nilIsNull":  true,
		"nilIsAny":   true,
		"nilIsEmpty": true,
		"nilLength":  0,
		"length":     1,
		"hasAny":     true,
		"notEmpty":   true,
	}
	expected.ShouldBeEqual(t, 0, "NameListCollection", actual)
}

// ── Identifiers ──

func Test_Cov2_Identifiers_AddsAndLookup(t *testing.T) {
	ids := coreinstruction.EmptyIdentifiers()
	ids.Add("a").Adds("b", "c")
	actual := args.Map{
		"length":   ids.Length(),
		"hasAny":   ids.HasAnyItem(),
		"indexOf":  ids.IndexOf("b"),
		"notFound": ids.IndexOf("z"),
		"getById":  ids.GetById("a") != nil,
		"getNil":   ids.GetById("z") == nil,
	}
	expected := args.Map{
		"length":   3,
		"hasAny":   true,
		"indexOf":  1,
		"notFound": -1,
		"getById":  true,
		"getNil":   true,
	}
	expected.ShouldBeEqual(t, 0, "Identifiers", actual)
}

func Test_Cov2_Identifiers_EmptyOps(t *testing.T) {
	ids := coreinstruction.EmptyIdentifiers()
	ids.Add("") // should skip
	actual := args.Map{
		"isEmpty":  ids.IsEmpty(),
		"indexOf":  ids.IndexOf(""),
		"getEmpty": ids.GetById("") == nil,
	}
	expected := args.Map{
		"isEmpty":  true,
		"indexOf":  -1,
		"getEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Identifiers_Empty", actual)
}

func Test_Cov2_Identifiers_Clone(t *testing.T) {
	ids := coreinstruction.NewIdentifiers("a", "b")
	cloned := ids.Clone()
	actual := args.Map{"len": cloned.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Identifiers_Clone", actual)
}

func Test_Cov2_Identifiers_NewUsingCap(t *testing.T) {
	ids := coreinstruction.NewIdentifiersUsingCap(5)
	actual := args.Map{"len": ids.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NewIdentifiersUsingCap", actual)
}

// ── LineIdentifier ──

func Test_Cov2_LineIdentifier_Methods(t *testing.T) {
	var nilLi *coreinstruction.LineIdentifier
	actual := args.Map{
		"nilInvalid":   nilLi.IsInvalidLineNumber(),
		"nilHasLine":   nilLi.HasLineNumber(),
		"nilNewLine":   nilLi.IsNewLineRequest(),
		"nilDelete":    nilLi.IsDeleteLineRequest(),
		"nilModify":    nilLi.IsModifyLineRequest(),
		"nilAddOrMod":  nilLi.IsAddNewOrModifyLineRequest(),
		"nilToBase":    nilLi.ToBaseLineIdentifier() == nil,
		"nilClone":     nilLi.Clone() == nil,
		"nilInvalidLn": nilLi.IsInvalidLineNumberUsingLastLineNumber(10),
	}
	expected := args.Map{
		"nilInvalid":   true,
		"nilHasLine":   false,
		"nilNewLine":   false,
		"nilDelete":    false,
		"nilModify":    false,
		"nilAddOrMod":  false,
		"nilToBase":    true,
		"nilClone":     true,
		"nilInvalidLn": true,
	}
	expected.ShouldBeEqual(t, 0, "LineIdentifier_Nil", actual)
}

// ── BaseTypeDotFilter ──

func Test_Cov2_BaseTypeDotFilter(t *testing.T) {
	f := &coreinstruction.BaseTypeDotFilter{}
	// Uses zero value — just trigger the function
	splits := f.GetDotSplitTypes()
	splits2 := f.GetDotSplitTypes() // cached
	actual := args.Map{"same": len(splits) == len(splits2)}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "BaseTypeDotFilter", actual)
}

// ── RequestSpecification ──

func Test_Cov2_RequestSpecification_Clone(t *testing.T) {
	rs := coreinstruction.RequestSpecification{}
	rs.Id = "x"
	rs.Tags = []string{"t1"}
	cloned := rs.Clone()
	actual := args.Map{"id": cloned.Id, "tagsLen": len(cloned.Tags)}
	expected := args.Map{"id": "x", "tagsLen": 1}
	expected.ShouldBeEqual(t, 0, "RequestSpecification_Clone", actual)
}
