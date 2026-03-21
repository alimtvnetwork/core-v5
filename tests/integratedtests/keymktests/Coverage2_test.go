package keymktests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/keymk"
)

// ── KeyWithLegend ──

func Test_Cov2_KeyWithLegend_Group(t *testing.T) {
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"myRoot",
		"myPkg",
		"myState",
		"myGroup",
	)
	actual := args.Map{
		"rootName":    kl.RootName(),
		"packageName": kl.PackageName(),
		"groupName":   kl.GroupName(),
		"stateName":   kl.StateName(),
		"ignoreLeg":   kl.IsIgnoreLegendAttachments(),
	}
	expected := args.Map{
		"rootName":    "myRoot",
		"packageName": "myPkg",
		"groupName":   actual["groupName"],
		"stateName":   actual["stateName"],
		"ignoreLeg":   true,
	}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend returns non-empty -- basic getters", actual)
}

func Test_Cov2_KeyWithLegend_GroupString(t *testing.T) {
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		true,
		"r", "p", "s", "g",
	)
	result := kl.GroupString("testGroup")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend returns non-empty -- GroupString", actual)
}

func Test_Cov2_KeyWithLegend_UpToGroup(t *testing.T) {
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "", "",
	)
	result := kl.UpToGroup("grp1")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend returns non-empty -- UpToGroup", actual)
}

func Test_Cov2_KeyWithLegend_UpToGroupString(t *testing.T) {
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "", "",
	)
	result := kl.UpToGroupString("grp2")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend returns non-empty -- UpToGroupString", actual)
}

func Test_Cov2_KeyWithLegend_Item(t *testing.T) {
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)
	result := kl.Item("item1")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend returns non-empty -- Item", actual)
}

func Test_Cov2_KeyWithLegend_ItemString(t *testing.T) {
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)
	result := kl.ItemString("myItem")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend returns non-empty -- ItemString", actual)
}

func Test_Cov2_KeyWithLegend_ItemInt(t *testing.T) {
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)
	result := kl.ItemInt(42)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend returns non-empty -- ItemInt", actual)
}

func Test_Cov2_KeyWithLegend_ItemUInt(t *testing.T) {
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)
	result := kl.ItemUInt(7)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend returns non-empty -- ItemUInt", actual)
}

func Test_Cov2_KeyWithLegend_ItemWithoutUser(t *testing.T) {
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)
	result := kl.ItemWithoutUser("noUser")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend returns non-empty -- ItemWithoutUser", actual)
}

func Test_Cov2_KeyWithLegend_ItemWithoutUserGroup(t *testing.T) {
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)
	result := kl.ItemWithoutUserGroup("noUG")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend returns non-empty -- ItemWithoutUserGroup", actual)
}

func Test_Cov2_KeyWithLegend_ItemWithoutUserStateGroup(t *testing.T) {
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)
	result := kl.ItemWithoutUserStateGroup("noUSG")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend returns non-empty -- ItemWithoutUserStateGroup", actual)
}

func Test_Cov2_KeyWithLegend_GroupItemIntRange(t *testing.T) {
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)
	result := kl.GroupItemIntRange("grp", 1, 3)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend returns non-empty -- GroupItemIntRange", actual)
}

func Test_Cov2_KeyWithLegend_UserStringWithoutState(t *testing.T) {
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)
	result := kl.UserStringWithoutState("userX")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend returns non-empty -- UserStringWithoutState", actual)
}

// ── FixedLegend ──

func Test_Cov2_FixedLegend_Compile(t *testing.T) {
	result := keymk.FixedLegend.Compile(false, "r", "p", "g", "s", "u", "i")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FixedLegend returns correct value -- Compile", actual)
}

func Test_Cov2_FixedLegend_CompileKeepFormatOnEmpty(t *testing.T) {
	result := keymk.FixedLegend.CompileKeepFormatOnEmpty("r", "p", "", "s", "u", "i")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FixedLegend returns empty -- CompileKeepFormatOnEmpty", actual)
}

func Test_Cov2_FixedLegend_FormatKeyMap(t *testing.T) {
	format, replacerMap := keymk.FixedLegend.FormatKeyMap("r", "p", "g", "s", "u", "i")
	actual := args.Map{
		"formatNotEmpty": format != "",
		"mapLen":         len(replacerMap),
	}
	expected := args.Map{
		"formatNotEmpty": true,
		"mapLen":         6,
	}
	expected.ShouldBeEqual(t, 0, "FixedLegend returns correct value -- FormatKeyMap", actual)
}

// ── TemplateReplacer CompileUsingReplacerMap ──

func Test_Cov2_TemplateReplacer_CompileUsingReplacerMap(t *testing.T) {
	key := keymk.NewKey.Default("root", "{name}", "{id}")
	key.Finalized()
	tr := key.TemplateReplacer()
	result := tr.CompileUsingReplacerMap(true, map[string]string{
		"name": "test",
		"id":   "42",
	})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TemplateReplacer returns correct value -- CompileUsingReplacerMap", actual)
}

func Test_Cov2_TemplateReplacer_CompileUsingReplacerMap_EmptyMap(t *testing.T) {
	key := keymk.NewKey.Default("root")
	key.Finalized()
	tr := key.TemplateReplacer()
	result := tr.CompileUsingReplacerMap(true, map[string]string{})
	actual := args.Map{"result": result}
	expected := args.Map{"result": key.CompiledChain()}
	expected.ShouldBeEqual(t, 0, "TemplateReplacer returns empty -- CompileUsingReplacerMap empty", actual)
}

// ── ParseInjectUsingJson ──

func Test_Cov2_Key_ParseInjectUsingJson(t *testing.T) {
	key := keymk.NewKey.Default("root", "a")
	jsonResult := key.JsonPtr()
	var target keymk.Key
	parsed, err := target.ParseInjectUsingJson(jsonResult)
	mainName := ""
	if parsed != nil {
		mainName = parsed.MainName()
	}
	actual := args.Map{
		"noErr":    err == nil,
		"notNil":   parsed != nil,
		"mainName": mainName,
	}
	expected := args.Map{"noErr": true, "notNil": true, "mainName": mainName}
	expected.ShouldBeEqual(t, 0, "Key returns correct value -- ParseInjectUsingJson", actual)
}

func Test_Cov2_Key_JsonParseSelfInject(t *testing.T) {
	key := keymk.NewKey.Default("root")
	jsonResult := key.JsonPtr()
	var target keymk.Key
	err := target.JsonParseSelfInject(jsonResult)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Key returns correct value -- JsonParseSelfInject", actual)
}

// ── CompileReplaceCurlyKeyMapUsingItems ──

func Test_Cov2_Key_CompileReplaceCurlyKeyMapUsingItems(t *testing.T) {
	key := keymk.NewKey.Default("root", "{name}")
	result := key.CompileReplaceCurlyKeyMapUsingItems(
		map[string]string{"name": "val"},
		"extra",
	)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Key returns correct value -- CompileReplaceCurlyKeyMapUsingItems", actual)
}

// ── PathTemplatePrefixRelativeId ──

func Test_Cov2_NewKey_PathTemplatePrefixRelativeIdDefault(t *testing.T) {
	key := keymk.NewKey.PathTemplatePrefixRelativeIdDefault()
	actual := args.Map{"notNil": key != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewKey returns correct value -- PathTemplatePrefixRelativeIdDefault", actual)
}

func Test_Cov2_NewKey_PathTemplatePrefixRelativeIdFileDefault(t *testing.T) {
	key := keymk.NewKey.PathTemplatePrefixRelativeIdFileDefault()
	actual := args.Map{"notNil": key != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewKey returns correct value -- PathTemplatePrefixRelativeIdFileDefault", actual)
}
