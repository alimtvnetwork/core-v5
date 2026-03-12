package coreinstructiontests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coreinstruction"
)

// ==========================================
// BaseIdentifier
// ==========================================

func Test_BaseIdentifier_IdString(t *testing.T) {
	id := coreinstruction.NewIdentifier("test-id")
	if id.IdString() != "test-id" {
		t.Errorf("expected 'test-id', got '%s'", id.IdString())
	}
}

func Test_BaseIdentifier_IsIdEmpty(t *testing.T) {
	id := coreinstruction.NewIdentifier("")
	if !id.IsIdEmpty() {
		t.Error("empty id should be empty")
	}
	id2 := coreinstruction.NewIdentifier("x")
	if id2.IsIdEmpty() {
		t.Error("non-empty should not be empty")
	}
}

func Test_BaseIdentifier_IsIdWhitespace(t *testing.T) {
	id := coreinstruction.NewIdentifier("   ")
	if !id.IsIdWhitespace() {
		t.Error("whitespace id should be whitespace")
	}
}

func Test_BaseIdentifier_IsId(t *testing.T) {
	id := coreinstruction.NewIdentifier("test")
	if !id.IsId("test") {
		t.Error("should match")
	}
	if id.IsId("other") {
		t.Error("should not match")
	}
}

func Test_BaseIdentifier_IsIdCaseInsensitive(t *testing.T) {
	id := coreinstruction.NewIdentifier("Test")
	if !id.IsIdCaseInsensitive("test") {
		t.Error("should match case insensitive")
	}
}

func Test_BaseIdentifier_IsIdContains(t *testing.T) {
	id := coreinstruction.NewIdentifier("hello-world")
	if !id.IsIdContains("world") {
		t.Error("should contain 'world'")
	}
}

func Test_BaseIdentifier_IsIdRegexMatches(t *testing.T) {
	id := coreinstruction.NewIdentifier("test-123")
	re := regexp.MustCompile(`\d+`)
	if !id.IsIdRegexMatches(re) {
		t.Error("should match regex")
	}
}

func Test_BaseIdentifier_Clone(t *testing.T) {
	id := coreinstruction.NewIdentifier("orig")
	cloned := id.Clone()
	if cloned.IdString() != "orig" {
		t.Error("clone should have same id")
	}
}

// ==========================================
// BaseDisplay
// ==========================================

func Test_BaseDisplay_IsDisplay(t *testing.T) {
	spec := coreinstruction.NewSpecificationSimple("id", "MyDisplay", "type")
	if !spec.IsDisplay("MyDisplay") {
		t.Error("should match display")
	}
}

func Test_BaseDisplay_IsDisplayCaseInsensitive(t *testing.T) {
	spec := coreinstruction.NewSpecificationSimple("id", "MyDisplay", "type")
	if !spec.IsDisplayCaseInsensitive("mydisplay") {
		t.Error("should match case insensitive")
	}
}

func Test_BaseDisplay_IsDisplayContains(t *testing.T) {
	spec := coreinstruction.NewSpecificationSimple("id", "MyDisplay", "type")
	if !spec.IsDisplayContains("Disp") {
		t.Error("should contain 'Disp'")
	}
}

func Test_BaseDisplay_IsDisplayRegexMatches(t *testing.T) {
	spec := coreinstruction.NewSpecificationSimple("id", "display-123", "type")
	re := regexp.MustCompile(`\d+`)
	if !spec.IsDisplayRegexMatches(re) {
		t.Error("should match regex")
	}
}

// ==========================================
// BaseEnabler
// ==========================================

func Test_BaseEnabler_SetEnable(t *testing.T) {
	e := &coreinstruction.BaseEnabler{}
	e.SetEnable()
	if !e.IsEnabled {
		t.Error("should be enabled")
	}
}

func Test_BaseEnabler_SetDisable(t *testing.T) {
	e := &coreinstruction.BaseEnabler{IsEnabled: true}
	e.SetDisable()
	if e.IsEnabled {
		t.Error("should be disabled")
	}
}

func Test_BaseEnabler_SetEnableVal(t *testing.T) {
	e := &coreinstruction.BaseEnabler{}
	e.SetEnableVal(true)
	if !e.IsEnabled {
		t.Error("should be enabled")
	}
	e.SetEnableVal(false)
	if e.IsEnabled {
		t.Error("should be disabled")
	}
}

// ==========================================
// BaseFromTo
// ==========================================

func Test_BaseFromTo_Create(t *testing.T) {
	ft := coreinstruction.NewBaseFromTo("src", "dst")
	if ft.From != "src" || ft.To != "dst" {
		t.Error("from/to not set correctly")
	}
}

// ==========================================
// Specification
// ==========================================

func Test_Specification_Simple(t *testing.T) {
	spec := coreinstruction.NewSpecificationSimple("id1", "Display1", "Type1")
	if spec.Id != "id1" {
		t.Errorf("expected 'id1', got '%s'", spec.Id)
	}
	if spec.Display != "Display1" {
		t.Errorf("expected 'Display1'")
	}
	if spec.Type != "Type1" {
		t.Errorf("expected 'Type1'")
	}
}

func Test_Specification_SimpleGlobal(t *testing.T) {
	spec := coreinstruction.NewSpecificationSimpleGlobal("id1", "Display1", "Type1")
	if !spec.IsGlobal {
		t.Error("should be global")
	}
}

func Test_Specification_Full(t *testing.T) {
	spec := coreinstruction.NewSpecification("id1", "Display1", "Type1", []string{"tag1", "tag2"}, true)
	if len(spec.Tags) != 2 {
		t.Errorf("expected 2 tags, got %d", len(spec.Tags))
	}
	if !spec.IsGlobal {
		t.Error("should be global")
	}
}

func Test_Specification_Clone(t *testing.T) {
	spec := coreinstruction.NewSpecification("id1", "Display1", "Type1", []string{"tag1"}, true)
	cloned := spec.Clone()
	if cloned.Id != "id1" {
		t.Error("clone id mismatch")
	}
	cloned.Tags[0] = "modified"
	if spec.Tags[0] == "modified" {
		t.Error("clone should be independent")
	}
}

func Test_Specification_Clone_Nil(t *testing.T) {
	var spec *coreinstruction.Specification
	cloned := spec.Clone()
	if cloned != nil {
		t.Error("nil clone should return nil")
	}
}

func Test_Specification_FlatSpecification(t *testing.T) {
	spec := coreinstruction.NewSpecificationSimple("id1", "Display1", "Type1")
	flat := spec.FlatSpecification()
	if flat == nil {
		t.Error("should return non-nil flat spec")
	}
	// Second call should return cached
	flat2 := spec.FlatSpecification()
	if flat != flat2 {
		t.Error("should return same cached instance")
	}
}

func Test_Specification_FlatSpecification_Nil(t *testing.T) {
	var spec *coreinstruction.Specification
	flat := spec.FlatSpecification()
	if flat != nil {
		t.Error("nil spec should return nil flat")
	}
}

// ==========================================
// Rename
// ==========================================

func Test_Rename_Properties(t *testing.T) {
	r := &coreinstruction.Rename{Existing: "old", New: "new"}
	if r.FromName() != "old" || r.ToName() != "new" {
		t.Error("properties mismatch")
	}
	if r.ExistingName() != "old" || r.NewName() != "new" {
		t.Error("alias properties mismatch")
	}
}

func Test_Rename_IsNull(t *testing.T) {
	var r *coreinstruction.Rename
	if !r.IsNull() {
		t.Error("nil should be null")
	}
}

func Test_Rename_IsExistingEmpty(t *testing.T) {
	r := &coreinstruction.Rename{Existing: "", New: "new"}
	if !r.IsExistingEmpty() {
		t.Error("empty existing should be empty")
	}
}

func Test_Rename_IsNewEmpty(t *testing.T) {
	r := &coreinstruction.Rename{Existing: "old", New: ""}
	if !r.IsNewEmpty() {
		t.Error("empty new should be empty")
	}
}

func Test_Rename_String(t *testing.T) {
	r := coreinstruction.Rename{Existing: "old", New: "new"}
	s := r.String()
	if s == "" {
		t.Error("should return non-empty string")
	}
}

func Test_Rename_SourceDestination(t *testing.T) {
	r := &coreinstruction.Rename{Existing: "old", New: "new"}
	sd := r.SourceDestination()
	if sd == nil || sd.Source != "old" || sd.Destination != "new" {
		t.Error("source destination conversion failed")
	}
}

func Test_Rename_SourceDestination_Nil(t *testing.T) {
	var r *coreinstruction.Rename
	sd := r.SourceDestination()
	if sd != nil {
		t.Error("nil should return nil")
	}
}

func Test_Rename_FromTo(t *testing.T) {
	r := &coreinstruction.Rename{Existing: "old", New: "new"}
	ft := r.FromTo()
	if ft == nil || ft.From != "old" || ft.To != "new" {
		t.Error("from-to conversion failed")
	}
}

func Test_Rename_FromTo_Nil(t *testing.T) {
	var r *coreinstruction.Rename
	ft := r.FromTo()
	if ft != nil {
		t.Error("nil should return nil")
	}
}

func Test_Rename_SetFromToName(t *testing.T) {
	r := &coreinstruction.Rename{Existing: "old", New: "new"}
	r.SetFromName("newFrom")
	r.SetToName("newTo")
	if r.Existing != "newFrom" || r.New != "newTo" {
		t.Error("set methods failed")
	}
}

func Test_Rename_SetFromToName_Nil(t *testing.T) {
	var r *coreinstruction.Rename
	r.SetFromName("x") // should not panic
	r.SetToName("y")   // should not panic
}

func Test_Rename_Clone(t *testing.T) {
	r := &coreinstruction.Rename{Existing: "old", New: "new"}
	cloned := r.Clone()
	if cloned.Existing != "old" || cloned.New != "new" {
		t.Error("clone mismatch")
	}
}

func Test_Rename_Clone_Nil(t *testing.T) {
	var r *coreinstruction.Rename
	cloned := r.Clone()
	if cloned != nil {
		t.Error("nil clone should return nil")
	}
}

// ==========================================
// SourceDestination
// ==========================================

func Test_SourceDestination_Properties(t *testing.T) {
	sd := &coreinstruction.SourceDestination{Source: "src", Destination: "dst"}
	if sd.FromName() != "src" || sd.ToName() != "dst" {
		t.Error("properties mismatch")
	}
}

func Test_SourceDestination_IsNull(t *testing.T) {
	var sd *coreinstruction.SourceDestination
	if !sd.IsNull() {
		t.Error("nil should be null")
	}
}

func Test_SourceDestination_IsSourceEmpty(t *testing.T) {
	sd := &coreinstruction.SourceDestination{Source: ""}
	if !sd.IsSourceEmpty() {
		t.Error("empty source should be empty")
	}
}

func Test_SourceDestination_IsDestinationEmpty(t *testing.T) {
	sd := &coreinstruction.SourceDestination{Destination: ""}
	if !sd.IsDestinationEmpty() {
		t.Error("empty destination should be empty")
	}
}

func Test_SourceDestination_String(t *testing.T) {
	sd := coreinstruction.SourceDestination{Source: "src", Destination: "dst"}
	if sd.String() == "" {
		t.Error("should return non-empty string")
	}
}

func Test_SourceDestination_SetFromToName(t *testing.T) {
	sd := &coreinstruction.SourceDestination{Source: "src", Destination: "dst"}
	sd.SetFromName("newSrc")
	sd.SetToName("newDst")
	if sd.Source != "newSrc" || sd.Destination != "newDst" {
		t.Error("set methods failed")
	}
}

func Test_SourceDestination_SetFromToName_Nil(t *testing.T) {
	var sd *coreinstruction.SourceDestination
	sd.SetFromName("x")
	sd.SetToName("y")
}

func Test_SourceDestination_FromTo(t *testing.T) {
	sd := &coreinstruction.SourceDestination{Source: "src", Destination: "dst"}
	ft := sd.FromTo()
	if ft == nil || ft.From != "src" || ft.To != "dst" {
		t.Error("from-to conversion failed")
	}
}

func Test_SourceDestination_FromTo_Nil(t *testing.T) {
	var sd *coreinstruction.SourceDestination
	if sd.FromTo() != nil {
		t.Error("nil should return nil")
	}
}

func Test_SourceDestination_Rename(t *testing.T) {
	sd := &coreinstruction.SourceDestination{Source: "src", Destination: "dst"}
	r := sd.Rename()
	if r == nil || r.Existing != "src" || r.New != "dst" {
		t.Error("rename conversion failed")
	}
}

func Test_SourceDestination_Rename_Nil(t *testing.T) {
	var sd *coreinstruction.SourceDestination
	if sd.Rename() != nil {
		t.Error("nil should return nil")
	}
}

func Test_SourceDestination_Clone(t *testing.T) {
	sd := &coreinstruction.SourceDestination{Source: "src", Destination: "dst"}
	cloned := sd.Clone()
	if cloned.Source != "src" || cloned.Destination != "dst" {
		t.Error("clone mismatch")
	}
}

func Test_SourceDestination_Clone_Nil(t *testing.T) {
	var sd *coreinstruction.SourceDestination
	if sd.Clone() != nil {
		t.Error("nil clone should return nil")
	}
}

// ==========================================
// NameList
// ==========================================

func Test_NameList_IsNull(t *testing.T) {
	var nl *coreinstruction.NameList
	if !nl.IsNull() {
		t.Error("nil should be null")
	}
}

func Test_NameList_IsAnyNull_Nil(t *testing.T) {
	var nl *coreinstruction.NameList
	if !nl.IsAnyNull() {
		t.Error("nil should be any null")
	}
}

func Test_NameList_IsAnyNull_NilList(t *testing.T) {
	nl := &coreinstruction.NameList{Name: "test"}
	if !nl.IsAnyNull() {
		t.Error("nil list should be any null")
	}
}

func Test_NameList_IsNameEmpty(t *testing.T) {
	nl := &coreinstruction.NameList{Name: ""}
	if !nl.IsNameEmpty() {
		t.Error("empty name should be empty")
	}
}

func Test_NameList_HasName(t *testing.T) {
	nl := &coreinstruction.NameList{Name: "test"}
	if !nl.HasName() {
		t.Error("should have name")
	}
}

func Test_NameList_Clone_Nil(t *testing.T) {
	var nl *coreinstruction.NameList
	if nl.Clone(true) != nil {
		t.Error("nil clone should return nil")
	}
}
