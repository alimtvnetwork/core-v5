package keymktests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/keymk"
)

// ============================================================================
// KeyJson: Serialize, MarshalJSON, UnmarshalJSON, JsonModel, JsonString
// ============================================================================

func Test_KeyJson_Serialize_Ext2(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a")

	// Act
	bytes, err := key.Serialize()

	// Assert
	if err != nil {
		t.Errorf("Serialize should not error: %v", err)
	}
	if len(bytes) == 0 {
		t.Error("Serialize should return non-empty bytes")
	}
}

func Test_KeyJson_MarshalUnmarshal_Ext2(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a", "b")

	// Act
	data, err := key.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON error: %v", err)
	}

	var restored keymk.Key
	err = restored.UnmarshalJSON(data)
	if err != nil {
		t.Fatalf("UnmarshalJSON error: %v", err)
	}

	// Assert
	if restored.MainName() != "root" {
		t.Errorf("expected main 'root', got '%s'", restored.MainName())
	}
}

func Test_KeyJson_JsonModel_Ext2(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a")

	// Act
	model := key.JsonModel()

	// Assert
	if model.MainName != "root" {
		t.Errorf("expected 'root', got '%s'", model.MainName)
	}
}

func Test_KeyJson_JsonModelAny_Ext2(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root")

	// Act
	result := key.JsonModelAny()

	// Assert
	if result == nil {
		t.Error("JsonModelAny should not be nil")
	}
}

func Test_KeyJson_JsonString_Ext2(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root")

	// Act
	result := key.JsonString()

	// Assert
	if result == "" {
		t.Error("JsonString should return non-empty")
	}
}

func Test_KeyJson_Json_Ext2(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root")

	// Act
	result := key.Json()

	// Assert
	if result.HasError() {
		t.Errorf("Json() should not have error: %v", result.MeaningfulError())
	}
}

func Test_KeyJson_JsonPtr_Ext2(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root")

	// Act
	result := key.JsonPtr()

	// Assert
	if result == nil {
		t.Error("JsonPtr should not be nil")
	}
}

func Test_KeyJson_ParseInjectUsingJson_Ext2(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a")
	jsonResult := key.JsonPtr()

	// Act
	var target keymk.Key
	parsed, err := target.ParseInjectUsingJson(jsonResult)

	// Assert
	if err != nil {
		t.Errorf("ParseInjectUsingJson error: %v", err)
	}
	if parsed == nil {
		t.Error("should return non-nil")
	}
}

func Test_KeyJson_ParseInjectUsingJsonMust_Ext2(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a")
	jsonResult := key.JsonPtr()

	// Act
	var target keymk.Key
	parsed := target.ParseInjectUsingJsonMust(jsonResult)

	// Assert
	if parsed == nil {
		t.Error("ParseInjectUsingJsonMust should return non-nil")
	}
}

func Test_KeyJson_AsJsonContractsBinder_Ext2(t *testing.T) {
	key := keymk.NewKey.Default("root")
	binder := key.AsJsonContractsBinder()
	if binder == nil {
		t.Error("should not be nil")
	}
}

func Test_KeyJson_AsJsoner_Ext2(t *testing.T) {
	key := keymk.NewKey.Default("root")
	jsoner := key.AsJsoner()
	if jsoner == nil {
		t.Error("should not be nil")
	}
}

func Test_KeyJson_JsonParseSelfInject_Ext2(t *testing.T) {
	key := keymk.NewKey.Default("root", "a")
	jsonResult := key.JsonPtr()

	var target keymk.Key
	err := target.JsonParseSelfInject(jsonResult)
	if err != nil {
		t.Errorf("JsonParseSelfInject error: %v", err)
	}
}

func Test_KeyJson_AsJsonParseSelfInjector_Ext2(t *testing.T) {
	key := keymk.NewKey.Default("root")
	injector := key.AsJsonParseSelfInjector()
	if injector == nil {
		t.Error("should not be nil")
	}
}

func Test_KeyJson_AsJsonMarshaller_Ext2(t *testing.T) {
	key := keymk.NewKey.Default("root")
	m := key.AsJsonMarshaller()
	if m == nil {
		t.Error("should not be nil")
	}
}

// ============================================================================
// TemplateReplacer
// ============================================================================

func Test_TemplateReplacer_IntRange_Ext2(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Curly("root", "id")
	key.Finalized()
	tr := key.TemplateReplacer()

	// Act
	result := tr.IntRange(true, "id", 0, 2)

	// Assert
	if len(result) != 3 {
		t.Errorf("expected 3 items, got %d", len(result))
	}
}

func Test_TemplateReplacer_RequestIntRange_Ext2(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Curly("root", "id")
	key.Finalized()
	tr := key.TemplateReplacer()

	// Act
	result := tr.RequestIntRange(true, keymk.TempReplace{
		KeyName: "id",
		Range:   keymk.Range{Start: 1, End: 3},
	})

	// Assert
	if len(result) != 3 {
		t.Errorf("expected 3 items, got %d", len(result))
	}
}

func Test_TemplateReplacer_CompileUsingReplacerMap_Ext2(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Curly("root", "name")
	key.Finalized()
	tr := key.TemplateReplacer()

	// Act
	result := tr.CompileUsingReplacerMap(true, map[string]string{
		"root": "myRoot",
		"name": "myName",
	})

	// Assert
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_TemplateReplacer_CompileUsingReplacerMap_Empty_Ext2(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Curly("root")
	key.Finalized()
	tr := key.TemplateReplacer()

	// Act
	result := tr.CompileUsingReplacerMap(true, map[string]string{})

	// Assert
	if result == "" {
		t.Error("should return compiled chain")
	}
}

// ============================================================================
// FixedLegend
// ============================================================================

func Test_FixedLegend_FormatKeyMap_Ext2(t *testing.T) {
	// Act
	format, replacerMap := keymk.FixedLegend.FormatKeyMap(
		"r", "p", "g", "s", "u", "i",
	)

	// Assert
	if format == "" {
		t.Error("format should not be empty")
	}
	if len(replacerMap) != 6 {
		t.Errorf("expected 6 replacers, got %d", len(replacerMap))
	}
}

func Test_FixedLegend_Compile_Ext2(t *testing.T) {
	// Act
	result := keymk.FixedLegend.Compile(
		false, "r", "p", "g", "s", "u", "i",
	)

	// Assert
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_FixedLegend_CompileKeepFormatOnEmpty_Ext2(t *testing.T) {
	// Act
	result := keymk.FixedLegend.CompileKeepFormatOnEmpty(
		"r", "p", "", "s", "u", "i",
	)

	// Assert
	if result == "" {
		t.Error("should return non-empty")
	}
}

// ============================================================================
// KeyWithLegend additional methods
// ============================================================================

func Test_KeyWithLegend_NoLegend_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	if !k.IsIgnoreLegendAttachments() {
		t.Error("should ignore legend attachments")
	}
}

func Test_KeyWithLegend_Create_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.Create(keymk.JoinerOption, "r", "p", "g")
	if k.IsIgnoreLegendAttachments() {
		t.Error("should not ignore legend attachments")
	}
}

func Test_KeyWithLegend_ShortLegend_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.ShortLegend(keymk.JoinerOption, "r", "p", "g")
	if k.IsIgnoreLegendAttachments() {
		t.Error("should not ignore legend attachments")
	}
}

func Test_KeyWithLegend_NoLegendPackage_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegendPackage(false, keymk.JoinerOption, "r", "g")
	if !k.IsIgnoreLegendAttachments() {
		t.Error("should ignore legend")
	}
}

func Test_KeyWithLegend_Getters_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.All(keymk.JoinerOption, keymk.FullLegends, true, "r", "p", "g", "s")
	if k.RootName() != "r" {
		t.Error("RootName mismatch")
	}
	if k.PackageName() != "p" {
		t.Error("PackageName mismatch")
	}
	if k.GroupName() != "g" {
		t.Error("GroupName mismatch")
	}
	if k.StateName() != "s" {
		t.Error("StateName mismatch")
	}
}

func Test_KeyWithLegend_Item_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.Item("myitem")
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_ItemString_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.ItemString("myitem")
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_ItemInt_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.ItemInt(42)
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_ItemUInt_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.ItemUInt(42)
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_GroupItemIntRange_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupItemIntRange("grp", 0, 2)
	if len(result) != 3 {
		t.Errorf("expected 3, got %d", len(result))
	}
}

func Test_KeyWithLegend_GroupUIntRange_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupUIntRange(0, 2)
	if len(result) != 3 {
		t.Errorf("expected 3, got %d", len(result))
	}
}

func Test_KeyWithLegend_ItemIntRange_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.ItemIntRange(0, 2)
	if len(result) != 3 {
		t.Errorf("expected 3, got %d", len(result))
	}
}

func Test_KeyWithLegend_ItemUIntRange_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.ItemUIntRange(0, 2)
	if len(result) != 3 {
		t.Errorf("expected 3, got %d", len(result))
	}
}

func Test_KeyWithLegend_Group_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.Group("myg")
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_GroupString_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupString("myg")
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_UpToGroup_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.UpToGroup("myg")
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_UpToGroupString_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.UpToGroupString("myg")
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_ItemWithoutUser_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.ItemWithoutUser("item1")
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_ItemWithoutUserGroup_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.ItemWithoutUserGroup("item1")
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_ItemWithoutUserStateGroup_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.ItemWithoutUserStateGroup("item1")
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_GroupUser_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupUser("g1", "u1")
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_GroupUserString_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupUserString("g1", "u1")
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_GroupUInt_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupUInt(1)
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_GroupByte_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupByte(1)
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_GroupUserByte_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupUserByte(1, 2)
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_GroupUserItem_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupUserItem("g1", "u1", "i1")
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_GroupStateUserItem_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupStateUserItem("g1", "s1", "u1", "i1")
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_StateUserItem_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.StateUserItem("s1", "u1", "i1")
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_StateUser_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.StateUser("s1", "u1")
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_GroupStateUserItemString_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupStateUserItemString("g1", "s1", "u1", "i1")
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_GroupUserItemString_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupUserItemString("g1", "u1", "i1")
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_GroupUserItemUint_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupUserItemUint(1, 2, 3)
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_GroupUserItemInt_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupUserItemInt(1, 2, 3)
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_GroupItem_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupItem("g1", "i1")
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_StateItem_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.StateItem("s1", "i1")
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_GroupItemString_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupItemString("g1", "i1")
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_GroupStateItemString_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupStateItemString("g1", "s1", "i1")
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_StateItemString_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.StateItemString("s1", "i1")
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_Compile_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.Compile("i1")
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_CompileDefault_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.CompileDefault()
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_CompileUsingJoiner_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.CompileUsingJoiner("/")
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_CompileStrings_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.CompileStrings()
	if len(result) == 0 {
		t.Error("should return non-empty slice")
	}
}

func Test_KeyWithLegend_Strings_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.Strings()
	if len(result) == 0 {
		t.Error("should return non-empty slice")
	}
}

func Test_KeyWithLegend_CompileItemUsingJoiner_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.CompileItemUsingJoiner("/", "i1")
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_KeyWithLegend_Clone_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	cloned := k.Clone()
	if cloned == nil {
		t.Error("Clone should not be nil")
	}
}

func Test_KeyWithLegend_CloneUsing_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	cloned := k.CloneUsing("newGroup")
	if cloned == nil {
		t.Error("CloneUsing should not be nil")
	}
	if cloned.GroupName() != "newGroup" {
		t.Errorf("expected 'newGroup', got '%s'", cloned.GroupName())
	}
}

func Test_KeyWithLegend_NilCloneUsing_Ext2(t *testing.T) {
	var k *keymk.KeyWithLegend
	cloned := k.CloneUsing("newGroup")
	if cloned != nil {
		t.Error("nil CloneUsing should return nil")
	}
}

func Test_KeyWithLegend_OutputItemsArray_WithLegend_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.All(keymk.JoinerOption, keymk.FullLegends, true, "r", "p", "g", "s")
	request := keymk.KeyLegendCompileRequest{
		GroupId:   "g1",
		StateName: "s1",
		UserId:    "u1",
		ItemId:    "i1",
	}
	result := k.OutputItemsArray(request)
	if len(result) == 0 {
		t.Error("should return non-empty items")
	}
}

func Test_KeyWithLegend_FinalStrings_WithBrackets_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.BracketJoinerOption, "r", "p", "g")
	request := keymk.KeyLegendCompileRequest{
		GroupId: "g1",
		ItemId:  "i1",
	}
	result := k.FinalStrings(request)
	if len(result) == 0 {
		t.Error("should return non-empty items")
	}
}

// ============================================================================
// Key: JoinUsingOption, CompileReplaceCurlyKeyMapUsingItems, CompileReplaceMapUsingItemsOption
// ============================================================================

func Test_Key_JoinUsingOption_Ext2(t *testing.T) {
	key := keymk.NewKey.Default("root", "a")
	opt := &keymk.Option{
		Joiner:           "/",
		IsSkipEmptyEntry: true,
	}
	result := key.JoinUsingOption(opt, "b")
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_Key_CompileReplaceCurlyKeyMapUsingItems_Ext2(t *testing.T) {
	key := keymk.NewKey.Curly("root", "name")
	compiled := key.CompileReplaceCurlyKeyMapUsingItems(
		map[string]string{"root": "myroot", "name": "myname"},
		"extra",
	)
	if compiled == "" {
		t.Error("should return non-empty")
	}
}

func Test_Key_CompileReplaceMapUsingItemsOption_NoCurly_Ext2(t *testing.T) {
	key := keymk.NewKey.Default("root", "name")
	compiled := key.CompileReplaceMapUsingItemsOption(
		false,
		map[string]string{"root": "myroot"},
	)
	if compiled == "" {
		t.Error("should return non-empty")
	}
}

func Test_Key_CompileReplaceMapUsingItemsOption_EmptyMap_Ext2(t *testing.T) {
	key := keymk.NewKey.Default("root")
	compiled := key.CompileReplaceMapUsingItemsOption(
		true,
		map[string]string{},
	)
	if compiled != "root" {
		t.Errorf("expected 'root', got '%s'", compiled)
	}
}

// ============================================================================
// Key: Finalized then Compile with additional items
// ============================================================================

func Test_Key_Finalized_CompileWithAdditional_Ext2(t *testing.T) {
	key := keymk.NewKey.Default("root", "a")
	key.Finalized()

	// Compile with additional should append
	result := key.Compile("b")
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_Key_Finalized_CompileStringsWithAdditional_Ext2(t *testing.T) {
	key := keymk.NewKey.Default("root", "a")
	key.Finalized()

	result := key.CompileStrings("b")
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_Key_Finalized_CompileNoAdditional_Ext2(t *testing.T) {
	key := keymk.NewKey.Default("root", "a")
	key.Finalized()

	result := key.Compile()
	if result != "root-a" {
		t.Errorf("expected 'root-a', got '%s'", result)
	}
}

func Test_Key_Finalized_CompileStringsNoAdditional_Ext2(t *testing.T) {
	key := keymk.NewKey.Default("root", "a")
	key.Finalized()

	result := key.CompileStrings()
	if result != "root-a" {
		t.Errorf("expected 'root-a', got '%s'", result)
	}
}

// ============================================================================
// KeyLegendCompileRequest constructors
// ============================================================================

func Test_KeyLegendCompileRequest_NewKeyLegend_Ext2(t *testing.T) {
	req := keymk.KeyLegendCompileRequest{GroupId: "g1"}
	k := req.NewKeyLegend(keymk.JoinerOption, keymk.ShortLegends, false, "r", "p", "s")
	if k == nil {
		t.Error("should not be nil")
	}
}

func Test_KeyLegendCompileRequest_NewKeyLegendDefaults_Ext2(t *testing.T) {
	req := keymk.KeyLegendCompileRequest{GroupId: "g1"}
	k := req.NewKeyLegendDefaults("r", "p", "s")
	if k == nil {
		t.Error("should not be nil")
	}
}

// ============================================================================
// NewKey creators: PathTemplatePrefixRelativeIdDefault, PathTemplatePrefixRelativeIdFileDefault
// ============================================================================

func Test_NewKey_PathTemplatePrefixRelativeIdDefault_Ext2(t *testing.T) {
	key := keymk.NewKey.PathTemplatePrefixRelativeIdDefault()
	if key.Compile() == "" {
		t.Error("should return non-empty")
	}
}

func Test_NewKey_PathTemplatePrefixRelativeIdFileDefault_Ext2(t *testing.T) {
	key := keymk.NewKey.PathTemplatePrefixRelativeIdFileDefault()
	if key.Compile() == "" {
		t.Error("should return non-empty")
	}
}

func Test_NewKey_CurlyStrings_Ext2(t *testing.T) {
	key := keymk.NewKey.CurlyStrings("root", "a")
	if key.Compile() == "" {
		t.Error("should return non-empty")
	}
}

func Test_NewKey_SquareBracketsStrings_Ext2(t *testing.T) {
	key := keymk.NewKey.SquareBracketsStrings("root", "a")
	if key.Compile() == "" {
		t.Error("should return non-empty")
	}
}

func Test_NewKey_ParenthesisStrings_Ext2(t *testing.T) {
	key := keymk.NewKey.ParenthesisStrings("root", "a")
	if key.Compile() == "" {
		t.Error("should return non-empty")
	}
}

func Test_NewKey_StringsWithOptions_Ext2(t *testing.T) {
	key := keymk.NewKey.StringsWithOptions(keymk.JoinerOption, "root", "a")
	if key.Compile() == "" {
		t.Error("should return non-empty")
	}
}

func Test_NewKey_OptionMain_Ext2(t *testing.T) {
	key := keymk.NewKey.OptionMain(keymk.JoinerOption, "root")
	if key.Compile() != "root" {
		t.Errorf("expected 'root', got '%s'", key.Compile())
	}
}

// ============================================================================
// Key: ItemEnumByte via KeyWithLegend (need to pass a mock enuminf.ByteEnumNamer)
// ============================================================================

func Test_KeyWithLegend_ItemEnumByte_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.ItemEnumByte(mockByteEnumNamer{name: "test-item"})
	if result == "" {
		t.Error("should return non-empty")
	}
}

// ============================================================================
// Key: AppendChainStrings skip empty
// ============================================================================

func Test_Key_AppendChainStrings_SkipEmpty_Ext2(t *testing.T) {
	key := keymk.NewKey.Default("root")
	key.AppendChainStrings("", "a", "", "b")
	actual := args.Map{
		"length": key.Length(),
	}
	expected := args.Map{
		"length": 2,
	}
	if actual["length"] != expected["length"] {
		t.Errorf("expected 2 chains (empty skipped), got %d", key.Length())
	}
}
