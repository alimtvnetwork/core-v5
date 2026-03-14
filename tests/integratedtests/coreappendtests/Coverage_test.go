package coreappendtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreappend"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── AppendAnyItemsToStringSkipOnNil ──

func Test_Cov_AppendAnyItems_Basic(t *testing.T) {
	result := coreappend.AppendAnyItemsToStringSkipOnNil(
		", ", "suffix", "a", "b",
	)
	actual := args.Map{"result": result}
	expected := args.Map{"result": "a, b, suffix"}
	expected.ShouldBeEqual(t, 0, "AppendAnyItems basic -- joined", actual)
}

func Test_Cov_AppendAnyItems_NilItems(t *testing.T) {
	result := coreappend.AppendAnyItemsToStringSkipOnNil(
		", ", "end", nil, "a", nil,
	)
	actual := args.Map{"result": result}
	expected := args.Map{"result": "a, end"}
	expected.ShouldBeEqual(t, 0, "AppendAnyItems nil items -- skipped", actual)
}

func Test_Cov_AppendAnyItems_NilAppend(t *testing.T) {
	result := coreappend.AppendAnyItemsToStringSkipOnNil(
		", ", nil, "a", "b",
	)
	actual := args.Map{"result": result}
	expected := args.Map{"result": "a, b"}
	expected.ShouldBeEqual(t, 0, "AppendAnyItems nil append -- no suffix", actual)
}

// ── PrependAnyItemsToStringSkipOnNil ──

func Test_Cov_PrependAnyItems_Basic(t *testing.T) {
	result := coreappend.PrependAnyItemsToStringSkipOnNil(
		", ", "prefix", "a", "b",
	)
	actual := args.Map{"result": result}
	expected := args.Map{"result": "prefix, a, b"}
	expected.ShouldBeEqual(t, 0, "PrependAnyItems basic -- joined", actual)
}

func Test_Cov_PrependAnyItems_NilPrepend(t *testing.T) {
	result := coreappend.PrependAnyItemsToStringSkipOnNil(
		", ", nil, "a", "b",
	)
	actual := args.Map{"result": result}
	expected := args.Map{"result": "a, b"}
	expected.ShouldBeEqual(t, 0, "PrependAnyItems nil prepend -- no prefix", actual)
}

// ── PrependAppendAnyItemsToStringSkipOnNil ──

func Test_Cov_PrependAppendAnyItems_Both(t *testing.T) {
	result := coreappend.PrependAppendAnyItemsToStringSkipOnNil(
		", ", "pre", "post", "a",
	)
	actual := args.Map{"result": result}
	expected := args.Map{"result": "pre, a, post"}
	expected.ShouldBeEqual(t, 0, "PrependAppend both -- joined", actual)
}

func Test_Cov_PrependAppendAnyItems_BothNil(t *testing.T) {
	result := coreappend.PrependAppendAnyItemsToStringSkipOnNil(
		", ", nil, nil, "a",
	)
	actual := args.Map{"result": result}
	expected := args.Map{"result": "a"}
	expected.ShouldBeEqual(t, 0, "PrependAppend both nil -- items only", actual)
}

// ── PrependAppendAnyItemsToStringsUsingFunc ──

func Test_Cov_PrependAppendUsingFunc_Basic(t *testing.T) {
	fn := func(item any) string {
		if item == nil {
			return ""
		}
		return item.(string)
	}
	result := coreappend.PrependAppendAnyItemsToStringsUsingFunc(
		true, fn, "pre", "post", "a", nil, "b",
	)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 4} // pre, a, b, post
	expected.ShouldBeEqual(t, 0, "UsingFunc basic -- skips empty", actual)
}

func Test_Cov_PrependAppendUsingFunc_NoSkipEmpty(t *testing.T) {
	fn := func(item any) string {
		if item == nil {
			return ""
		}
		return item.(string)
	}
	result := coreappend.PrependAppendAnyItemsToStringsUsingFunc(
		false, fn, "pre", "post", "a", nil,
	)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3} // pre, a, post (nil item skipped)
	expected.ShouldBeEqual(t, 0, "UsingFunc no skip -- nil item skipped", actual)
}

// ── MapStringStringAppendMapStringToAnyItems ──

func Test_Cov_MapStringStringAppend_Basic(t *testing.T) {
	mainMap := map[string]string{"a": "1"}
	appendMap := map[string]any{"b": 2, "c": "three"}
	result := coreappend.MapStringStringAppendMapStringToAnyItems(
		false, mainMap, appendMap,
	)
	actual := args.Map{
		"hasA": result["a"] == "1",
		"hasB": result["b"] != "",
		"hasC": result["c"] == "three",
	}
	expected := args.Map{"hasA": true, "hasB": true, "hasC": true}
	expected.ShouldBeEqual(t, 0, "MapAppend basic -- merged", actual)
}

func Test_Cov_MapStringStringAppend_EmptyAppend(t *testing.T) {
	mainMap := map[string]string{"a": "1"}
	result := coreappend.MapStringStringAppendMapStringToAnyItems(
		false, mainMap, nil,
	)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "MapAppend empty append -- unchanged", actual)
}

func Test_Cov_MapStringStringAppend_SkipEmpty(t *testing.T) {
	mainMap := map[string]string{}
	appendMap := map[string]any{"a": "", "b": "val"}
	result := coreappend.MapStringStringAppendMapStringToAnyItems(
		true, mainMap, appendMap,
	)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1} // "a" skipped because empty
	expected.ShouldBeEqual(t, 0, "MapAppend skip empty -- only b", actual)
}

// ── PrependAppendAnyItemsToStringsSkipOnNil (direct slice) ──

func Test_Cov_PrependAppendStrings_Empty(t *testing.T) {
	result := coreappend.PrependAppendAnyItemsToStringsSkipOnNil(
		nil, nil,
	)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "PrependAppendStrings empty -- zero items", actual)
}
