package keymktests

import (
	"testing"

	"github.com/alimtvnetwork/core/keymk"
)

func Test_Cov_Key_CompileKeys(t *testing.T) {
	k := keymk.NewKey("root", "a")
	k2 := keymk.NewKey("sub", "b")
	result := k.CompileKeys(k2)
	if result == "" {
		t.Error("expected non-empty")
	}
	// nil key in list
	result2 := k.CompileKeys(nil, k2)
	if result2 == "" {
		t.Error("expected non-empty")
	}
}

func Test_Cov_Key_CompileKeys_Empty(t *testing.T) {
	k := keymk.NewKey("root", "a")
	result := k.CompileKeys()
	if result == "" {
		t.Error("expected compiled")
	}
}

func Test_Cov_Key_Finalized(t *testing.T) {
	k := keymk.NewKey("root", "a")
	k.Finalized("extra")
	// now IsComplete should be true, subsequent compiles use cached
	r1 := k.Compile()
	r2 := k.Compile("more")
	if r1 == "" || r2 == "" {
		t.Error("expected non-empty")
	}
}

func Test_Cov_Key_CompileStrings_Finalized(t *testing.T) {
	k := keymk.NewKey("root")
	k.Finalized()
	r := k.CompileStrings("a", "b")
	if r == "" {
		t.Error("expected non-empty")
	}
	// empty additional
	r2 := k.CompileStrings()
	if r2 == "" {
		t.Error("expected non-empty")
	}
}

func Test_Cov_Key_CompileReplaceCurlyKeyMap(t *testing.T) {
	k := keymk.NewKey("root", "{name}")
	result := k.CompileReplaceCurlyKeyMap(map[string]string{"name": "world"})
	if result == "" {
		t.Error("expected non-empty")
	}
}

func Test_Cov_Key_CompileReplaceMapUsingItemsOption_NoCurly(t *testing.T) {
	k := keymk.NewKey("root", "KEY")
	result := k.CompileReplaceMapUsingItemsOption(false, map[string]string{"KEY": "val"})
	if result == "" {
		t.Error("expected non-empty")
	}
}

func Test_Cov_Key_CompileReplaceMapUsingItemsOption_EmptyMap(t *testing.T) {
	k := keymk.NewKey("root", "a")
	result := k.CompileReplaceMapUsingItemsOption(true, nil)
	if result == "" {
		t.Error("expected non-empty")
	}
}

func Test_Cov_Key_IntRange(t *testing.T) {
	k := keymk.NewKey("item")
	result := k.IntRange(0, 2)
	if len(result) != 3 {
		t.Errorf("expected 3 got %d", len(result))
	}
}

func Test_Cov_Key_IntRangeEnding(t *testing.T) {
	k := keymk.NewKey("item")
	result := k.IntRangeEnding(2)
	if len(result) != 3 {
		t.Errorf("expected 3 got %d", len(result))
	}
}

func Test_Cov_Key_JoinUsingOption(t *testing.T) {
	k := keymk.NewKey("root", "a")
	opt := keymk.NewOption("-")
	result := k.JoinUsingOption(opt, "b")
	if result == "" {
		t.Error("expected non-empty")
	}
}

func Test_Cov_KeyJson_Serialize_Unmarshal(t *testing.T) {
	k := keymk.NewKey("root", "a")
	data, err := k.Serialize()
	if err != nil || len(data) == 0 {
		t.Error("expected serialize success")
	}

	k2 := &keymk.Key{}
	err = k2.UnmarshalJSON(data)
	if err != nil {
		t.Error("expected unmarshal success")
	}
}

func Test_Cov_KeyJson_ParseInjectUsingJson(t *testing.T) {
	k := keymk.NewKey("root", "a")
	jr := k.JsonPtr()
	k2 := &keymk.Key{}
	result, err := k2.ParseInjectUsingJson(jr)
	if err != nil || result == nil {
		t.Error("expected success")
	}
}

func Test_Cov_KeyJson_ParseInjectUsingJsonMust(t *testing.T) {
	k := keymk.NewKey("root")
	jr := k.JsonPtr()
	k2 := &keymk.Key{}
	result := k2.ParseInjectUsingJsonMust(jr)
	if result == nil {
		t.Error("expected non-nil")
	}
}

func Test_Cov_KeyJson_AsJsonContractsBinder(t *testing.T) {
	k := keymk.NewKey("root")
	if k.AsJsonContractsBinder() == nil {
		t.Error("expected non-nil")
	}
	if k.AsJsoner() == nil {
		t.Error("expected non-nil")
	}
	if k.AsJsonParseSelfInjector() == nil {
		t.Error("expected non-nil")
	}
	if k.AsJsonMarshaller() == nil {
		t.Error("expected non-nil")
	}
}

func Test_Cov_KeyJson_JsonParseSelfInject(t *testing.T) {
	k := keymk.NewKey("root")
	jr := k.JsonPtr()
	k2 := &keymk.Key{}
	err := k2.JsonParseSelfInject(jr)
	if err != nil {
		t.Error("expected no error")
	}
}

func Test_Cov_KeyJson_TemplateReplacer(t *testing.T) {
	k := keymk.NewKey("root", "{name}")
	tr := k.TemplateReplacer()
	_ = tr
}
