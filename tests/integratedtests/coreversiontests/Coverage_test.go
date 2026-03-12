package coreversiontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreversion"
)

// ==========================================
// Version field accessors
// ==========================================

func Test_Version_MajorString(t *testing.T) {
	v := coreversion.New.Create("1.2.3")
	if v.MajorString() != "1" {
		t.Errorf("expected '1', got '%s'", v.MajorString())
	}
}

func Test_Version_MinorString(t *testing.T) {
	v := coreversion.New.Create("1.2.3")
	if v.MinorString() != "2" {
		t.Errorf("expected '2', got '%s'", v.MinorString())
	}
}

func Test_Version_PatchString(t *testing.T) {
	v := coreversion.New.Create("1.2.3")
	if v.PatchString() != "3" {
		t.Errorf("expected '3', got '%s'", v.PatchString())
	}
}

func Test_Version_BuildString(t *testing.T) {
	v := coreversion.New.Create("1.2.3.4")
	if v.BuildString() != "4" {
		t.Errorf("expected '4', got '%s'", v.BuildString())
	}
}

// ==========================================
// Version nil receivers
// ==========================================

func Test_Version_Nil_MajorString(t *testing.T) {
	var v *coreversion.Version
	if v.MajorString() != "" {
		t.Error("nil should return empty")
	}
}

func Test_Version_Nil_MinorString(t *testing.T) {
	var v *coreversion.Version
	if v.MinorString() != "" {
		t.Error("nil should return empty")
	}
}

func Test_Version_Nil_PatchString(t *testing.T) {
	var v *coreversion.Version
	if v.PatchString() != "" {
		t.Error("nil should return empty")
	}
}

func Test_Version_Nil_BuildString(t *testing.T) {
	var v *coreversion.Version
	if v.BuildString() != "" {
		t.Error("nil should return empty")
	}
}

// ==========================================
// Has / IsInvalid checks
// ==========================================

func Test_Version_HasMajor(t *testing.T) {
	v := coreversion.New.Create("1.0.0")
	if !v.HasMajor() {
		t.Error("should have major")
	}
}

func Test_Version_HasMinor(t *testing.T) {
	v := coreversion.New.Create("1.2.0")
	if !v.HasMinor() {
		t.Error("should have minor")
	}
}

func Test_Version_HasPatch(t *testing.T) {
	v := coreversion.New.Create("1.2.3")
	if !v.HasPatch() {
		t.Error("should have patch")
	}
}

func Test_Version_HasBuild(t *testing.T) {
	v := coreversion.New.Create("1.2.3.4")
	if !v.HasBuild() {
		t.Error("should have build")
	}
}

func Test_Version_IsMajorInvalid_Nil(t *testing.T) {
	var v *coreversion.Version
	if !v.IsMajorInvalid() {
		t.Error("nil should be invalid")
	}
}

func Test_Version_IsMinorInvalid_Nil(t *testing.T) {
	var v *coreversion.Version
	if !v.IsMinorInvalid() {
		t.Error("nil should be invalid")
	}
}

func Test_Version_IsPatchInvalid_Nil(t *testing.T) {
	var v *coreversion.Version
	if !v.IsPatchInvalid() {
		t.Error("nil should be invalid")
	}
}

func Test_Version_IsBuildInvalid_Nil(t *testing.T) {
	var v *coreversion.Version
	if !v.IsBuildInvalid() {
		t.Error("nil should be invalid")
	}
}

// ==========================================
// InvalidOrZero checks
// ==========================================

func Test_Version_IsMajorInvalidOrZero(t *testing.T) {
	v := coreversion.New.Create("0.1.0")
	if !v.IsMajorInvalidOrZero() {
		t.Error("major 0 should be invalid or zero")
	}
}

func Test_Version_IsMinorInvalidOrZero(t *testing.T) {
	v := coreversion.New.Create("1.0.0")
	if !v.IsMinorInvalidOrZero() {
		t.Error("minor 0 should be invalid or zero")
	}
}

func Test_Version_IsPatchInvalidOrZero(t *testing.T) {
	v := coreversion.New.Create("1.1.0")
	if !v.IsPatchInvalidOrZero() {
		t.Error("patch 0 should be invalid or zero")
	}
}

func Test_Version_IsBuildInvalidOrZero(t *testing.T) {
	v := coreversion.New.Create("1.1.1")
	if !v.IsBuildInvalidOrZero() {
		t.Error("no build should be invalid or zero")
	}
}

// ==========================================
// Display methods
// ==========================================

func Test_Version_VersionDisplay(t *testing.T) {
	v := coreversion.New.Create("1.2.3")
	d := v.VersionDisplay()
	if d == "" {
		t.Error("should return non-empty display")
	}
}

func Test_Version_VersionDisplay_Nil(t *testing.T) {
	var v *coreversion.Version
	if v.VersionDisplay() != "" {
		t.Error("nil should return empty")
	}
}

func Test_Version_VersionDisplayMajor(t *testing.T) {
	v := coreversion.New.Create("5.2.3")
	d := v.VersionDisplayMajor()
	if d == "" {
		t.Error("should return non-empty")
	}
}

func Test_Version_VersionDisplayMajorMinor(t *testing.T) {
	v := coreversion.New.Create("5.2.3")
	d := v.VersionDisplayMajorMinor()
	if d == "" {
		t.Error("should return non-empty")
	}
}

func Test_Version_VersionDisplayMajorMinorPatch(t *testing.T) {
	v := coreversion.New.Create("5.2.3")
	d := v.VersionDisplayMajorMinorPatch()
	if d == "" {
		t.Error("should return non-empty")
	}
}

func Test_Version_CompiledVersion(t *testing.T) {
	v := coreversion.New.Create("1.2.3")
	if v.CompiledVersion() == "" {
		t.Error("should return non-empty")
	}
}

func Test_Version_CompiledVersion_Nil(t *testing.T) {
	var v *coreversion.Version
	if v.CompiledVersion() != "" {
		t.Error("nil should return empty")
	}
}

func Test_Version_String(t *testing.T) {
	v := coreversion.New.Create("1.2.3")
	if v.String() == "" {
		t.Error("should return non-empty")
	}
}

// ==========================================
// Validity checks
// ==========================================

func Test_Version_IsEmptyOrInvalid_Valid(t *testing.T) {
	v := coreversion.New.Create("1.2.3")
	if v.IsEmptyOrInvalid() {
		t.Error("valid version should not be empty/invalid")
	}
}

func Test_Version_IsEmptyOrInvalid_Nil(t *testing.T) {
	var v *coreversion.Version
	if !v.IsEmptyOrInvalid() {
		t.Error("nil should be empty/invalid")
	}
}

func Test_Version_HasAnyItem(t *testing.T) {
	v := coreversion.New.Create("1.2.3")
	if !v.HasAnyItem() {
		t.Error("should have item")
	}
}

func Test_Version_IsDefined(t *testing.T) {
	v := coreversion.New.Create("1.2.3")
	if !v.IsDefined() {
		t.Error("should be defined")
	}
}

// ==========================================
// Version comparison
// ==========================================

func Test_Version_IsVersionCompareEqual(t *testing.T) {
	v := coreversion.New.Create("1.2.3")
	if !v.IsVersionCompareEqual("1.2.3") {
		t.Error("should be equal")
	}
}

func Test_Version_IsVersionCompareNotEqual(t *testing.T) {
	v := coreversion.New.Create("1.2.3")
	if !v.IsVersionCompareNotEqual("2.0.0") {
		t.Error("should not be equal")
	}
}

func Test_Version_IsVersionCompareEqual_Nil(t *testing.T) {
	var v *coreversion.Version
	if !v.IsVersionCompareEqual("") {
		t.Error("nil with empty should be equal")
	}
	if v.IsVersionCompareEqual("1.0.0") {
		t.Error("nil with value should not be equal")
	}
}

// ==========================================
// ValueByIndex / AllVersionValues
// ==========================================

func Test_Version_AllVersionValues(t *testing.T) {
	v := coreversion.New.Create("1.2.3.4")
	vals := v.AllVersionValues()
	if len(vals) == 0 {
		t.Error("should return non-empty")
	}
}

func Test_Version_AllValidVersionValues(t *testing.T) {
	v := coreversion.New.Create("1.2.3")
	vals := v.AllValidVersionValues()
	if len(vals) == 0 {
		t.Error("should return non-empty")
	}
}

// ==========================================
// IsMajorAtLeast / IsMajorMinorAtLeast etc.
// ==========================================

func Test_Version_IsMajorAtLeast(t *testing.T) {
	v := coreversion.New.Create("3.0.0")
	if !v.IsMajorAtLeast(2) {
		t.Error("3 should be at least 2")
	}
	if v.IsMajorAtLeast(5) {
		t.Error("3 should not be at least 5")
	}
}

func Test_Version_IsMajorMinorAtLeast(t *testing.T) {
	v := coreversion.New.Create("3.2.0")
	if !v.IsMajorMinorAtLeast(3, 1) {
		t.Error("3.2 should be at least 3.1")
	}
}

func Test_Version_IsMajorMinorPatchAtLeast(t *testing.T) {
	v := coreversion.New.Create("3.2.1")
	if !v.IsMajorMinorPatchAtLeast(3, 2, 0) {
		t.Error("3.2.1 should be at least 3.2.0")
	}
}

func Test_Version_IsMajorStringAtLeast(t *testing.T) {
	v := coreversion.New.Create("3.0.0")
	if !v.IsMajorStringAtLeast("2") {
		t.Error("3 should be at least 2")
	}
}

// ==========================================
// Compare functions
// ==========================================

func Test_CompareVersionString(t *testing.T) {
	cmp := coreversion.CompareVersionString("2.0.0", "1.0.0")
	if !cmp.IsLeftGreaterEqualLogically() {
		t.Error("2.0.0 should be greater than 1.0.0")
	}
}

func Test_IsAtLeast(t *testing.T) {
	if !coreversion.IsAtLeast("2.0.0", "1.0.0") {
		t.Error("2.0.0 should be at least 1.0.0")
	}
	if coreversion.IsAtLeast("1.0.0", "2.0.0") {
		t.Error("1.0.0 should not be at least 2.0.0")
	}
}

func Test_IsLower(t *testing.T) {
	if !coreversion.IsLower("1.0.0", "2.0.0") {
		t.Error("1.0.0 should be lower than 2.0.0")
	}
}

func Test_IsLowerOrEqual(t *testing.T) {
	if !coreversion.IsLowerOrEqual("1.0.0", "1.0.0") {
		t.Error("equal should be lower or equal")
	}
	if !coreversion.IsLowerOrEqual("1.0.0", "2.0.0") {
		t.Error("1.0.0 should be lower or equal to 2.0.0")
	}
}

func Test_IsExpectedVersion(t *testing.T) {
	// just exercise the function
	_ = coreversion.IsExpectedVersion(
		coreversion.CompareVersionString("1.0.0", "1.0.0"),
		"1.0.0",
		"1.0.0",
	)
}

// ==========================================
// VersionsCollection
// ==========================================

func Test_VersionsCollection_Basic(t *testing.T) {
	vc := &coreversion.VersionsCollection{}
	vc.Add("1.0.0")
	vc.Add("2.0.0")
	if vc.Length() != 2 {
		t.Errorf("expected 2, got %d", vc.Length())
	}
	if vc.Count() != 2 {
		t.Error("Count should equal Length")
	}
	if vc.IsEmpty() {
		t.Error("should not be empty")
	}
	if !vc.HasAnyItem() {
		t.Error("should have items")
	}
	if vc.LastIndex() != 1 {
		t.Errorf("expected last index 1, got %d", vc.LastIndex())
	}
	if !vc.HasIndex(1) {
		t.Error("should have index 1")
	}
	if vc.HasIndex(5) {
		t.Error("should not have index 5")
	}
}

func Test_VersionsCollection_AddSkipInvalid(t *testing.T) {
	vc := &coreversion.VersionsCollection{}
	vc.AddSkipInvalid("1.0.0")
	vc.AddSkipInvalid("")
	if vc.Length() != 1 {
		t.Errorf("expected 1, got %d", vc.Length())
	}
}

func Test_VersionsCollection_AddVersionsRaw(t *testing.T) {
	vc := &coreversion.VersionsCollection{}
	vc.AddVersionsRaw("1.0.0", "2.0.0")
	if vc.Length() != 2 {
		t.Errorf("expected 2, got %d", vc.Length())
	}
}

func Test_VersionsCollection_Strings(t *testing.T) {
	vc := &coreversion.VersionsCollection{}
	vc.Add("1.0.0")
	compactStrs := vc.VersionCompactStrings()
	if len(compactStrs) != 1 {
		t.Errorf("expected 1, got %d", len(compactStrs))
	}
	verStrs := vc.VersionsStrings()
	if len(verStrs) != 1 {
		t.Errorf("expected 1, got %d", len(verStrs))
	}
}

func Test_VersionsCollection_Strings_Empty(t *testing.T) {
	vc := &coreversion.VersionsCollection{}
	if len(vc.VersionCompactStrings()) != 0 {
		t.Error("empty should return empty")
	}
	if len(vc.VersionsStrings()) != 0 {
		t.Error("empty should return empty")
	}
}

func Test_VersionsCollection_IndexOf(t *testing.T) {
	vc := &coreversion.VersionsCollection{}
	vc.Add("1.0.0")
	vc.Add("2.0.0")
	if vc.IndexOf("2.0.0") < 0 {
		t.Error("should find version")
	}
	if vc.IndexOf("3.0.0") >= 0 {
		t.Error("should not find version")
	}
}

func Test_VersionsCollection_IsContainsVersion(t *testing.T) {
	vc := &coreversion.VersionsCollection{}
	vc.Add("1.0.0")
	if !vc.IsContainsVersion("1.0.0") {
		t.Error("should contain version")
	}
}

func Test_VersionsCollection_IsEqual(t *testing.T) {
	vc1 := &coreversion.VersionsCollection{}
	vc1.Add("1.0.0")
	vc2 := &coreversion.VersionsCollection{}
	vc2.Add("1.0.0")
	if !vc1.IsEqual(vc2) {
		t.Error("should be equal")
	}
}

func Test_VersionsCollection_IsEqual_DifferentLength(t *testing.T) {
	vc1 := &coreversion.VersionsCollection{}
	vc1.Add("1.0.0")
	vc2 := &coreversion.VersionsCollection{}
	if vc1.IsEqual(vc2) {
		t.Error("different lengths should not be equal")
	}
}

func Test_VersionsCollection_IsEqual_BothNil(t *testing.T) {
	var vc1 *coreversion.VersionsCollection
	var vc2 *coreversion.VersionsCollection
	if !vc1.IsEqual(vc2) {
		t.Error("both nil should be equal")
	}
}

func Test_VersionsCollection_IsEqual_OneNil(t *testing.T) {
	var vc1 *coreversion.VersionsCollection
	vc2 := &coreversion.VersionsCollection{}
	if vc1.IsEqual(vc2) {
		t.Error("one nil should not be equal")
	}
}

func Test_VersionsCollection_String(t *testing.T) {
	vc := coreversion.VersionsCollection{}
	vc.Add("1.0.0")
	if vc.String() == "" {
		t.Error("should return non-empty string")
	}
}

func Test_VersionsCollection_Length_Nil(t *testing.T) {
	var vc *coreversion.VersionsCollection
	if vc.Length() != 0 {
		t.Error("nil should return 0")
	}
}

// ==========================================
// Version comparison methods
// ==========================================

func Test_Version_MajorBuild(t *testing.T) {
	v := coreversion.New.Create("3.0.0.5")
	cmp := v.MajorBuild(3, 5)
	if !cmp.IsEqual() {
		t.Error("should be equal")
	}
}

func Test_Version_IsMajorBuildAtLeast(t *testing.T) {
	v := coreversion.New.Create("3.0.0.5")
	if !v.IsMajorBuildAtLeast(3, 4) {
		t.Error("3.0.0.5 should be at least 3.0.0.4")
	}
}

func Test_Version_MajorBuildString(t *testing.T) {
	v := coreversion.New.Create("3.0.0.5")
	cmp := v.MajorBuildString("3", "5")
	if !cmp.IsEqual() {
		t.Error("should be equal")
	}
}

func Test_Version_MajorMinorPatchBuildString(t *testing.T) {
	v := coreversion.New.Create("3.2.1.5")
	cmp := v.MajorMinorPatchBuildString("3", "2", "5", "1")
	if !cmp.IsEqual() {
		t.Error("should be equal")
	}
}

func Test_Version_Patch(t *testing.T) {
	v := coreversion.New.Create("1.2.3")
	cmp := v.Patch(3)
	if !cmp.IsEqual() {
		t.Error("should be equal")
	}
}

func Test_Version_MajorPatch(t *testing.T) {
	v := coreversion.New.Create("1.2.3")
	cmp := v.MajorPatch(1, 3)
	if !cmp.IsEqual() {
		t.Error("should be equal")
	}
}

func Test_Version_Build(t *testing.T) {
	v := coreversion.New.Create("1.2.3.4")
	cmp := v.Build(4)
	if !cmp.IsEqual() {
		t.Error("should be equal")
	}
}

func Test_Version_MajorMinorPatchBuild(t *testing.T) {
	v := coreversion.New.Create("1.2.3.4")
	cmp := v.MajorMinorPatchBuild(1, 2, 3, 4)
	if !cmp.IsEqual() {
		t.Error("should be equal")
	}
}

func Test_Version_Compare(t *testing.T) {
	v1 := coreversion.New.Create("1.2.3")
	v2 := coreversion.New.Create("1.2.3")
	cmp := v1.Compare(&v2)
	if !cmp.IsEqual() {
		t.Error("should be equal")
	}
}

func Test_Version_IsEqual(t *testing.T) {
	v1 := coreversion.New.Create("1.2.3")
	v2 := coreversion.New.Create("1.2.3")
	if !v1.IsEqual(&v2) {
		t.Error("should be equal")
	}
}
