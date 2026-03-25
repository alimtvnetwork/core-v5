package coreversiontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreversion"
)

// TestVersion_Creation verifies version creation from various formats.
func TestVersion_Creation(t *testing.T) {
	for _, tc := range versionCreationCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			v := coreversion.New.Default(tc.input)

			// Assert
			if v.VersionMajor != tc.expectedMajor {
				t.Errorf("major: expected %d, got %d", tc.expectedMajor, v.VersionMajor)
			}
			if v.VersionMinor != tc.expectedMinor {
				t.Errorf("minor: expected %d, got %d", tc.expectedMinor, v.VersionMinor)
			}
			if v.VersionPatch != tc.expectedPatch {
				t.Errorf("patch: expected %d, got %d", tc.expectedPatch, v.VersionPatch)
			}
		})
	}
}

// TestVersion_Display verifies display methods.
func TestVersion_Display(t *testing.T) {
	v := coreversion.New.Default("v1.2.3")
	if v.VersionDisplay() == "" {
		t.Error("VersionDisplay should not be empty")
	}
	if v.CompiledVersion() == "" {
		t.Error("CompiledVersion should not be empty")
	}
	if v.String() == "" {
		t.Error("String should not be empty")
	}
	if v.VersionDisplayMajor() == "" {
		t.Error("VersionDisplayMajor should not be empty")
	}
	if v.VersionDisplayMajorMinor() == "" {
		t.Error("VersionDisplayMajorMinor should not be empty")
	}
	if v.VersionDisplayMajorMinorPatch() == "" {
		t.Error("VersionDisplayMajorMinorPatch should not be empty")
	}
}

// TestVersion_NilDisplay verifies nil pointer display.
func TestVersion_NilDisplay(t *testing.T) {
	var v *coreversion.Version
	if v.VersionDisplay() != "" {
		t.Error("nil VersionDisplay should be empty")
	}
	if v.CompiledVersion() != "" {
		t.Error("nil CompiledVersion should be empty")
	}
	if v.MajorString() != "" {
		t.Error("nil MajorString should be empty")
	}
	if v.MinorString() != "" {
		t.Error("nil MinorString should be empty")
	}
	if v.PatchString() != "" {
		t.Error("nil PatchString should be empty")
	}
	if v.BuildString() != "" {
		t.Error("nil BuildString should be empty")
	}
}

// TestVersion_HasMethods verifies Has/Invalid checks.
func TestVersion_HasMethods(t *testing.T) {
	v := coreversion.New.Default("v1.2.3")
	if !v.HasMajor() {
		t.Error("should have major")
	}
	if !v.HasMinor() {
		t.Error("should have minor")
	}
	if !v.HasPatch() {
		t.Error("should have patch")
	}
	if v.IsMajorInvalid() {
		t.Error("major should be valid")
	}
	if v.IsMinorInvalid() {
		t.Error("minor should be valid")
	}
	if v.IsPatchInvalid() {
		t.Error("patch should be valid")
	}
}

// TestVersion_NilHas verifies nil receiver for Has methods.
func TestVersion_NilHas(t *testing.T) {
	var v *coreversion.Version
	if v.HasMajor() {
		t.Error("nil should not have major")
	}
	if v.HasMinor() {
		t.Error("nil should not have minor")
	}
	if v.HasPatch() {
		t.Error("nil should not have patch")
	}
	if v.HasBuild() {
		t.Error("nil should not have build")
	}
	if !v.IsMajorInvalid() {
		t.Error("nil major should be invalid")
	}
}

// TestVersion_Empty verifies empty version.
func TestVersion_Empty(t *testing.T) {
	v := coreversion.New.Default("")
	if !v.IsEmptyOrInvalid() {
		t.Error("empty should be invalid")
	}
	if v.HasAnyItem() {
		t.Error("empty should not have items")
	}
	if v.IsDefined() {
		t.Error("empty should not be defined")
	}
}

// TestVersion_Comparison verifies comparison methods.
func TestVersion_Comparison(t *testing.T) {
	v1 := coreversion.New.Default("v1.2.3")
	v2 := coreversion.New.Default("v1.2.3")
	v3 := coreversion.New.Default("v2.0.0")

	if !v1.IsEqual(&v2) {
		t.Error("v1 should equal v2")
	}
	if v1.IsEqual(&v3) {
		t.Error("v1 should not equal v3")
	}
	if !v1.IsLeftLessThan(&v3) {
		t.Error("v1 should be less than v3")
	}
	if !v3.IsLeftGreaterThan(&v1) {
		t.Error("v3 should be greater than v1")
	}
	if !v1.IsLeftLessThanOrEqual(&v2) {
		t.Error("v1 should be <= v2")
	}
	if !v1.IsLeftGreaterThanOrEqual(&v2) {
		t.Error("v1 should be >= v2")
	}
}

// TestVersion_AtLeast verifies AtLeast.
func TestVersion_AtLeast(t *testing.T) {
	v := coreversion.New.Default("v2.1.0")
	if !v.IsAtLeast("v2.0.0") {
		t.Error("v2.1.0 should be at least v2.0.0")
	}
	if v.IsAtLeast("v3.0.0") {
		t.Error("v2.1.0 should not be at least v3.0.0")
	}
}

// TestVersion_IsEqualVersionString verifies string comparison.
func TestVersion_IsEqualVersionString(t *testing.T) {
	v := coreversion.New.Default("v1.0.0")
	if !v.IsEqualVersionString("v1.0.0") {
		t.Error("should be equal")
	}
	if !v.IsEqualVersionString("1.0.0") {
		t.Error("should be equal without v prefix")
	}
}

// TestVersion_IsLowerVersionString verifies lower version string.
func TestVersion_IsLowerVersionString(t *testing.T) {
	v := coreversion.New.Default("v1.0.0")
	if !v.IsLowerVersionString("v2.0.0") {
		t.Error("v1 should be lower than v2")
	}
}

// TestVersion_IsLowerEqualVersionString verifies lower or equal.
func TestVersion_IsLowerEqualVersionString(t *testing.T) {
	v := coreversion.New.Default("v1.0.0")
	if !v.IsLowerEqualVersionString("v1.0.0") {
		t.Error("v1 should be <= v1")
	}
}

// TestVersion_IsVersionCompareEqual verifies compact comparison.
func TestVersion_IsVersionCompareEqual(t *testing.T) {
	v := coreversion.New.Default("v1.2.3")
	if !v.IsVersionCompareEqual("1.2.3") {
		t.Error("should be equal")
	}
	if !v.IsVersionCompareNotEqual("2.0.0") {
		t.Error("should not be equal")
	}

	var nilV *coreversion.Version
	if !nilV.IsVersionCompareEqual("") {
		t.Error("nil with empty should be equal")
	}
	if nilV.IsVersionCompareEqual("1.0") {
		t.Error("nil with non-empty should not be equal")
	}
}

// TestVersion_Clone verifies clone.
func TestVersion_Clone(t *testing.T) {
	v := coreversion.New.Default("v1.2.3")
	cloned := v.Clone()
	if cloned.VersionMajor != v.VersionMajor {
		t.Error("clone should have same major")
	}
}

// TestVersion_ClonePtr verifies pointer clone.
func TestVersion_ClonePtr(t *testing.T) {
	var nilV *coreversion.Version
	if nilV.ClonePtr() != nil {
		t.Error("nil clone should be nil")
	}

	v := coreversion.New.Default("v1.0")
	ptr := v.ClonePtr()
	if ptr == nil {
		t.Error("clone should not be nil")
	}
}

// TestVersion_AllVersionValues verifies all version values.
func TestVersion_AllVersionValues(t *testing.T) {
	v := coreversion.New.Default("v1.2.3")
	vals := v.AllVersionValues()
	if len(vals) == 0 {
		t.Error("should have values")
	}
}

// TestVersion_AllValidVersionValues verifies valid version values.
func TestVersion_AllValidVersionValues(t *testing.T) {
	v := coreversion.New.Default("v1.2.3")
	vals := v.AllValidVersionValues()
	if len(vals) == 0 {
		t.Error("should have valid values")
	}
}

// TestVersion_IsMajorAtLeast verifies major at least.
func TestVersion_IsMajorAtLeast(t *testing.T) {
	v := coreversion.New.Default("v2.0.0")
	if !v.IsMajorAtLeast(1) {
		t.Error("2 should be at least 1")
	}
	if !v.IsMajorAtLeast(2) {
		t.Error("2 should be at least 2")
	}
	if v.IsMajorAtLeast(3) {
		t.Error("2 should not be at least 3")
	}
}

// TestVersion_IsMajorMinorAtLeast verifies major.minor at least.
func TestVersion_IsMajorMinorAtLeast(t *testing.T) {
	v := coreversion.New.Default("v2.3.0")
	if !v.IsMajorMinorAtLeast(2, 3) {
		t.Error("2.3 should be at least 2.3")
	}
	if !v.IsMajorMinorAtLeast(2, 2) {
		t.Error("2.3 should be at least 2.2")
	}
	if v.IsMajorMinorAtLeast(2, 4) {
		t.Error("2.3 should not be at least 2.4")
	}
}

// TestNewCreator_Spread verifies spread constructors.
func TestNewCreator_Spread(t *testing.T) {
	v := coreversion.New.SpreadIntegers(1, 2, 3)
	if v.VersionMajor != 1 {
		t.Error("major should be 1")
	}

	v = coreversion.New.SpreadStrings("1", "2")
	if v.VersionMajor != 1 {
		t.Error("major should be 1")
	}

	v = coreversion.New.SpreadBytes(1, 2, 3)
	if v.VersionMajor != 1 {
		t.Error("major should be 1")
	}

	v = coreversion.New.SpreadUnsignedIntegers(1, 2)
	if v.VersionMajor != 1 {
		t.Error("major should be 1")
	}
}

// TestNewCreator_AllVariants verifies all constructors.
func TestNewCreator_AllVariants(t *testing.T) {
	v := coreversion.New.MajorMinor("1", "2")
	if v.VersionMinor != 2 {
		t.Error("minor should be 2")
	}

	v = coreversion.New.MajorMinorPatch("1", "2", "3")
	if v.VersionPatch != 3 {
		t.Error("patch should be 3")
	}

	v = coreversion.New.MajorMinorPatchBuild("1", "2", "3", "4")
	if v.VersionBuild != 4 {
		t.Error("build should be 4")
	}

	v = coreversion.New.AllInt(1, 2, 3, 4)
	if v.VersionBuild != 4 {
		t.Error("AllInt build should be 4")
	}

	v = coreversion.New.AllByte(1, 2, 3, 4)
	if v.VersionBuild != 4 {
		t.Error("AllByte build should be 4")
	}

	v = coreversion.New.MajorMinorInt(1, 2)
	if v.VersionMinor != 2 {
		t.Error("MajorMinorInt minor should be 2")
	}

	v = coreversion.New.MajorMinorPatchInt(1, 2, 3)
	if v.VersionPatch != 3 {
		t.Error("MajorMinorPatchInt patch should be 3")
	}
}

// TestVersion_Json verifies JSON serialization.
func TestVersion_Json(t *testing.T) {
	v := coreversion.New.Default("v1.2.3")
	j := v.Json()
	if j.HasError() {
		t.Error("json should not have error")
	}

	jp := v.JsonPtr()
	if jp == nil {
		t.Error("json ptr should not be nil")
	}
}

// TestVersion_IsMajorStringAtLeast verifies string-based comparison.
func TestVersion_IsMajorStringAtLeast(t *testing.T) {
	v := coreversion.New.Default("v3.0.0")
	if !v.IsMajorStringAtLeast("2") {
		t.Error("3 should be at least 2")
	}
}

// TestVersion_InvalidOrZero verifies invalid-or-zero checks.
func TestVersion_InvalidOrZero(t *testing.T) {
	v := coreversion.New.Default("v0.0.0")
	if !v.IsMajorInvalidOrZero() {
		t.Error("0 should be invalid or zero")
	}
	if !v.IsMinorInvalidOrZero() {
		t.Error("0 should be invalid or zero")
	}
	if !v.IsPatchInvalidOrZero() {
		t.Error("0 should be invalid or zero")
	}
	if !v.IsBuildInvalidOrZero() {
		t.Error("0 should be invalid or zero")
	}
}

// TestVersion_NonPtrAndPtr verifies NonPtr/Ptr.
func TestVersion_NonPtrAndPtr(t *testing.T) {
	v := coreversion.New.Default("v1.0")
	np := v.NonPtr()
	if np.VersionMajor != 1 {
		t.Error("NonPtr major should be 1")
	}
	p := v.Ptr()
	if p == nil {
		t.Error("Ptr should not be nil")
	}
}
