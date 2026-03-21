package corestrtests

import (
	"fmt"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/smartystreets/goconvey/convey"
)

// =============================================================================
// Hashmap — deeper coverage (AddOrUpdate variants, lock variants, keys/values)
// =============================================================================

func Test_Cov55_Hashmap_AddOrUpdateKeyStrValFloat(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	// Act
	h.AddOrUpdateKeyStrValFloat("pi", 3.14)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"AddOrUpdateKeyStrValFloat",
			args.Map{"HasKey": true, "IsEmpty": false},
		),
	}
	actual := args.Map{"HasKey": h.Has("pi"), "IsEmpty": h.IsEmpty()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_AddOrUpdateKeyStrValFloat64(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	// Act
	h.AddOrUpdateKeyStrValFloat64("e", 2.71828)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"AddOrUpdateKeyStrValFloat64",
			args.Map{"HasKey": true},
		),
	}
	actual := args.Map{"HasKey": h.Has("e")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_AddOrUpdateKeyStrValAny(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	// Act
	h.AddOrUpdateKeyStrValAny("val", 42)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"AddOrUpdateKeyStrValAny",
			args.Map{"HasKey": true},
		),
	}
	actual := args.Map{"HasKey": h.Has("val")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_AddOrUpdateKeyValueAny(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	pair := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	// Act
	h.AddOrUpdateKeyValueAny(pair)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"AddOrUpdateKeyValueAny",
			args.Map{"Has": true, "Length": 1},
		),
	}
	actual := args.Map{"Has": h.Has("k"), "Length": h.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_AddOrUpdateKeyVal(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	kv := corestr.KeyValuePair{Key: "a", Value: "b"}
	// Act
	isNew := h.AddOrUpdateKeyVal(kv)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"AddOrUpdateKeyVal new",
			args.Map{"IsNew": true},
		),
	}
	actual := args.Map{"IsNew": isNew}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_AddOrUpdateKeyVal_Existing(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "old")
	kv := corestr.KeyValuePair{Key: "a", Value: "new"}
	// Act
	isNew := h.AddOrUpdateKeyVal(kv)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"AddOrUpdateKeyVal existing",
			args.Map{"IsNew": false},
		),
	}
	actual := args.Map{"IsNew": isNew}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_AddOrUpdateKeyValues(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	// Act
	h.AddOrUpdateKeyValues(
		corestr.KeyValuePair{Key: "a", Value: "1"},
		corestr.KeyValuePair{Key: "b", Value: "2"},
	)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"AddOrUpdateKeyValues",
			args.Map{"Length": 2},
		),
	}
	actual := args.Map{"Length": h.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_AddOrUpdateKeyValues_Empty(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	// Act
	h.AddOrUpdateKeyValues()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"AddOrUpdateKeyValues empty",
			args.Map{"Length": 0},
		),
	}
	actual := args.Map{"Length": h.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_AddOrUpdateKeyAnyValues(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	// Act
	h.AddOrUpdateKeyAnyValues(
		corestr.KeyAnyValuePair{Key: "x", Value: 10},
	)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"AddOrUpdateKeyAnyValues",
			args.Map{"Has": true},
		),
	}
	actual := args.Map{"Has": h.Has("x")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_AddOrUpdateKeyAnyValues_Empty(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	// Act
	h.AddOrUpdateKeyAnyValues()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"AddOrUpdateKeyAnyValues empty",
			args.Map{"Length": 0},
		),
	}
	actual := args.Map{"Length": h.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_AddOrUpdateLock(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	// Act
	h.AddOrUpdateLock("k", "v")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"AddOrUpdateLock",
			args.Map{"Has": true},
		),
	}
	actual := args.Map{"Has": h.Has("k")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_AddOrUpdateWithWgLock(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	// Act
	h.AddOrUpdateWithWgLock("wg", "val", wg)
	wg.Wait()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"AddOrUpdateWithWgLock",
			args.Map{"Has": true},
		),
	}
	actual := args.Map{"Has": h.Has("wg")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_AddOrUpdateStringsPtrWgLock(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	keys := []string{"a", "b"}
	vals := []string{"1", "2"}
	// Act
	h.AddOrUpdateStringsPtrWgLock(wg, keys, vals)
	wg.Wait()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"AddOrUpdateStringsPtrWgLock",
			args.Map{"Length": 2},
		),
	}
	actual := args.Map{"Length": h.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_AddOrUpdateStringsPtrWgLock_Empty(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	// Act
	h.AddOrUpdateStringsPtrWgLock(wg, []string{}, []string{})
	wg.Wait()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"AddOrUpdateStringsPtrWgLock empty",
			args.Map{"Length": 0},
		),
	}
	actual := args.Map{"Length": h.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_AddOrUpdateHashmap(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	other := corestr.New.Hashmap.Empty()
	other.AddOrUpdate("x", "y")
	// Act
	h.AddOrUpdateHashmap(other)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"AddOrUpdateHashmap",
			args.Map{"Has": true},
		),
	}
	actual := args.Map{"Has": h.Has("x")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_AddOrUpdateHashmap_Nil(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	// Act
	h.AddOrUpdateHashmap(nil)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"AddOrUpdateHashmap nil",
			args.Map{"Length": 0},
		),
	}
	actual := args.Map{"Length": h.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_AddOrUpdateMap(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	// Act
	h.AddOrUpdateMap(map[string]string{"k": "v"})
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"AddOrUpdateMap",
			args.Map{"Has": true},
		),
	}
	actual := args.Map{"Has": h.Has("k")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_AddOrUpdateMap_Empty(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	// Act
	h.AddOrUpdateMap(map[string]string{})
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"AddOrUpdateMap empty",
			args.Map{"Length": 0},
		),
	}
	actual := args.Map{"Length": h.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_AddOrUpdateCollection(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	keys := corestr.New.Collection.Strings([]string{"a", "b"})
	vals := corestr.New.Collection.Strings([]string{"1", "2"})
	// Act
	h.AddOrUpdateCollection(keys, vals)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"AddOrUpdateCollection",
			args.Map{"Length": 2},
		),
	}
	actual := args.Map{"Length": h.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_AddOrUpdateCollection_NilKeys(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	vals := corestr.New.Collection.Strings([]string{"1"})
	// Act
	h.AddOrUpdateCollection(nil, vals)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"AddOrUpdateCollection nil keys",
			args.Map{"Length": 0},
		),
	}
	actual := args.Map{"Length": h.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_AddOrUpdateCollection_Mismatch(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	keys := corestr.New.Collection.Strings([]string{"a", "b"})
	vals := corestr.New.Collection.Strings([]string{"1"})
	// Act
	h.AddOrUpdateCollection(keys, vals)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"AddOrUpdateCollection mismatch",
			args.Map{"Length": 0},
		),
	}
	actual := args.Map{"Length": h.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_AddsOrUpdates(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	// Act
	h.AddsOrUpdates(
		corestr.KeyValuePair{Key: "a", Value: "1"},
		corestr.KeyValuePair{Key: "b", Value: "2"},
	)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"AddsOrUpdates",
			args.Map{"Length": 2},
		),
	}
	actual := args.Map{"Length": h.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_AddsOrUpdates_Nil(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	// Act
	h.AddsOrUpdates()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"AddsOrUpdates nil",
			args.Map{"Length": 0},
		),
	}
	actual := args.Map{"Length": h.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_AddsOrUpdatesAnyUsingFilter(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	filter := func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
		return fmt.Sprintf("%v", pair.Value), true, false
	}
	// Act
	h.AddsOrUpdatesAnyUsingFilter(filter, corestr.KeyAnyValuePair{Key: "k", Value: 1})
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"AddsOrUpdatesAnyUsingFilter",
			args.Map{"Has": true},
		),
	}
	actual := args.Map{"Has": h.Has("k")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_AddsOrUpdatesAnyUsingFilter_Break(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	filter := func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
		return "", false, true
	}
	// Act
	h.AddsOrUpdatesAnyUsingFilter(filter, corestr.KeyAnyValuePair{Key: "k", Value: 1})
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"AddsOrUpdatesAnyUsingFilter break",
			args.Map{"Length": 0},
		),
	}
	actual := args.Map{"Length": h.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_AddsOrUpdatesAnyUsingFilterLock(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	filter := func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
		return "val", true, false
	}
	// Act
	h.AddsOrUpdatesAnyUsingFilterLock(filter, corestr.KeyAnyValuePair{Key: "k", Value: 1})
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"AddsOrUpdatesAnyUsingFilterLock",
			args.Map{"Has": true},
		),
	}
	actual := args.Map{"Has": h.Has("k")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_AddsOrUpdatesUsingFilter(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	filter := func(pair corestr.KeyValuePair) (string, bool, bool) {
		return pair.Value, true, false
	}
	// Act
	h.AddsOrUpdatesUsingFilter(filter, corestr.KeyValuePair{Key: "k", Value: "v"})
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"AddsOrUpdatesUsingFilter",
			args.Map{"Has": true},
		),
	}
	actual := args.Map{"Has": h.Has("k")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_ConcatNew_WithArgs(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	other := corestr.New.Hashmap.Empty()
	other.AddOrUpdate("b", "2")
	// Act
	result := h.ConcatNew(false, other)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"ConcatNew with args",
			args.Map{"HasA": true, "HasB": true},
		),
	}
	actual := args.Map{"HasA": result.Has("a"), "HasB": result.Has("b")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_ConcatNew_NoArgs(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	// Act
	result := h.ConcatNew(true)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"ConcatNew no args",
			args.Map{"Has": true},
		),
	}
	actual := args.Map{"Has": result.Has("a")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_ConcatNewUsingMaps(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	// Act
	result := h.ConcatNewUsingMaps(false, map[string]string{"b": "2"})
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"ConcatNewUsingMaps",
			args.Map{"HasB": true},
		),
	}
	actual := args.Map{"HasB": result.Has("b")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_ConcatNewUsingMaps_NoArgs(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	// Act
	result := h.ConcatNewUsingMaps(true)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"ConcatNewUsingMaps no args",
			args.Map{"Has": true},
		),
	}
	actual := args.Map{"Has": result.Has("a")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_HasAny(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	// Act
	found := h.HasAny("b", "a")
	notFound := h.HasAny("c", "d")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"HasAny",
			args.Map{"Found": true, "NotFound": false},
		),
	}
	actual := args.Map{"Found": found, "NotFound": notFound}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_HasWithLock(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("k", "v")
	// Act
	result := h.HasWithLock("k")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"HasWithLock",
			args.Map{"Has": true},
		),
	}
	actual := args.Map{"Has": result}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_HasAllCollectionItems(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	h.AddOrUpdate("b", "2")
	col := corestr.New.Collection.Strings([]string{"a", "b"})
	// Act
	result := h.HasAllCollectionItems(col)
	resultNil := h.HasAllCollectionItems(nil)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"HasAllCollectionItems",
			args.Map{"HasAll": true, "Nil": false},
		),
	}
	actual := args.Map{"HasAll": result, "Nil": resultNil}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_ContainsLock(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("k", "v")
	// Act
	result := h.ContainsLock("k")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"ContainsLock",
			args.Map{"Contains": true},
		),
	}
	actual := args.Map{"Contains": result}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_IsKeyMissing(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("k", "v")
	// Act
	missing := h.IsKeyMissing("other")
	notMissing := h.IsKeyMissing("k")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"IsKeyMissing",
			args.Map{"Missing": true, "NotMissing": false},
		),
	}
	actual := args.Map{"Missing": missing, "NotMissing": notMissing}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_IsKeyMissingLock(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("k", "v")
	// Act
	result := h.IsKeyMissingLock("other")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"IsKeyMissingLock",
			args.Map{"Missing": true},
		),
	}
	actual := args.Map{"Missing": result}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_Diff(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	other := corestr.New.Hashmap.Empty()
	other.AddOrUpdate("a", "2")
	// Act
	diff := h.Diff(other)
	// Assert
	convey.Convey("Diff returns non-nil", t, func() {
		convey.So(diff, convey.ShouldNotBeNil)
	})
}

func Test_Cov55_Hashmap_KeysValuesCollection(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	// Act
	keys, values := h.KeysValuesCollection()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"KeysValuesCollection",
			args.Map{"KeysLen": 1, "ValsLen": 1},
		),
	}
	actual := args.Map{"KeysLen": keys.Length(), "ValsLen": values.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_KeysValuesList(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	// Act
	keys, values := h.KeysValuesList()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"KeysValuesList",
			args.Map{"KeysLen": 1, "ValsLen": 1},
		),
	}
	actual := args.Map{"KeysLen": len(keys), "ValsLen": len(values)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_KeysValuePairs(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	// Act
	pairs := h.KeysValuePairs()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"KeysValuePairs",
			args.Map{"Length": 1},
		),
	}
	actual := args.Map{"Length": len(pairs)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_KeysValuePairsCollection(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	// Act
	col := h.KeysValuePairsCollection()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"KeysValuePairsCollection",
			args.Map{"Length": 1},
		),
	}
	actual := args.Map{"Length": col.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_KeysValuesListLock(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	// Act
	keys, values := h.KeysValuesListLock()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"KeysValuesListLock",
			args.Map{"KeysLen": 1, "ValsLen": 1},
		),
	}
	actual := args.Map{"KeysLen": len(keys), "ValsLen": len(values)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_KeysLock(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	// Act
	keys := h.KeysLock()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"KeysLock",
			args.Map{"Length": 1},
		),
	}
	actual := args.Map{"Length": len(keys)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_ItemsCopyLock(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	// Act
	items := h.ItemsCopyLock()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"ItemsCopyLock",
			args.Map{"Length": 1},
		),
	}
	actual := args.Map{"Length": len(*items)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_ValuesCollection(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	// Act
	col := h.ValuesCollection()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"ValuesCollection",
			args.Map{"Length": 1},
		),
	}
	actual := args.Map{"Length": col.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_ValuesHashset(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	// Act
	hs := h.ValuesHashset()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"ValuesHashset",
			args.Map{"Length": 1},
		),
	}
	actual := args.Map{"Length": hs.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_ValuesCollectionLock(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	// Act
	col := h.ValuesCollectionLock()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"ValuesCollectionLock",
			args.Map{"Length": 1},
		),
	}
	actual := args.Map{"Length": col.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_ValuesHashsetLock(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	// Act
	hs := h.ValuesHashsetLock()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"ValuesHashsetLock",
			args.Map{"Length": 1},
		),
	}
	actual := args.Map{"Length": hs.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_ValuesToLower(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("KEY", "VAL")
	// Act
	result := h.ValuesToLower()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"ValuesToLower (deprecated alias)",
			args.Map{"HasLower": true},
		),
	}
	actual := args.Map{"HasLower": result.Has("key")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_StringLock(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	// Act
	result := h.StringLock()
	// Assert
	convey.Convey("StringLock non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov55_Hashmap_StringLock_Empty(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	// Act
	result := h.StringLock()
	// Assert
	convey.Convey("StringLock empty", t, func() {
		convey.So(result, convey.ShouldContainSubstring, "No Element")
	})
}

func Test_Cov55_Hashmap_GetKeysFilteredItems(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("abc", "1")
	h.AddOrUpdate("def", "2")
	filter := func(s string, i int) (string, bool, bool) {
		return s, s == "abc", false
	}
	// Act
	result := h.GetKeysFilteredItems(filter)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"GetKeysFilteredItems",
			args.Map{"Length": 1},
		),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_GetKeysFilteredItems_Empty(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	filter := func(s string, i int) (string, bool, bool) {
		return s, true, false
	}
	// Act
	result := h.GetKeysFilteredItems(filter)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"GetKeysFilteredItems empty",
			args.Map{"Length": 0},
		),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_GetKeysFilteredItems_Break(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	h.AddOrUpdate("b", "2")
	filter := func(s string, i int) (string, bool, bool) {
		return s, true, true
	}
	// Act
	result := h.GetKeysFilteredItems(filter)
	// Assert
	convey.Convey("GetKeysFilteredItems break", t, func() {
		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
	})
}

func Test_Cov55_Hashmap_GetKeysFilteredCollection(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("abc", "1")
	filter := func(s string, i int) (string, bool, bool) {
		return s, true, false
	}
	// Act
	result := h.GetKeysFilteredCollection(filter)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"GetKeysFilteredCollection",
			args.Map{"Length": 1},
		),
	}
	actual := args.Map{"Length": result.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_GetKeysFilteredCollection_Empty(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	filter := func(s string, i int) (string, bool, bool) {
		return s, true, false
	}
	// Act
	result := h.GetKeysFilteredCollection(filter)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"GetKeysFilteredCollection empty",
			args.Map{"IsEmpty": true},
		),
	}
	actual := args.Map{"IsEmpty": result.IsEmpty()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_GetKeysFilteredCollection_Break(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	h.AddOrUpdate("b", "2")
	filter := func(s string, i int) (string, bool, bool) {
		return s, true, true
	}
	// Act
	result := h.GetKeysFilteredCollection(filter)
	// Assert
	convey.Convey("GetKeysFilteredCollection break", t, func() {
		convey.So(result.Length(), convey.ShouldBeGreaterThan, 0)
	})
}

func Test_Cov55_Hashmap_SafeItems_Nil(t *testing.T) {
	// Arrange
	var h *corestr.Hashmap
	// Act
	result := h.SafeItems()
	// Assert
	convey.Convey("SafeItems nil", t, func() {
		convey.So(result, convey.ShouldBeNil)
	})
}

func Test_Cov55_Hashmap_GetValue(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("k", "v")
	// Act
	val, found := h.GetValue("k")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"GetValue",
			args.Map{"Val": "v", "Found": true},
		),
	}
	actual := args.Map{"Val": val, "Found": found}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_ToError(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("k", "v")
	// Act
	err := h.ToError(", ")
	// Assert
	convey.Convey("ToError", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func Test_Cov55_Hashmap_ToDefaultError(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("k", "v")
	// Act
	err := h.ToDefaultError()
	// Assert
	convey.Convey("ToDefaultError", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func Test_Cov55_Hashmap_KeyValStringLines(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("k", "v")
	// Act
	lines := h.KeyValStringLines()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"KeyValStringLines",
			args.Map{"Length": 1},
		),
	}
	actual := args.Map{"Length": len(lines)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_ToStringsUsingCompiler_Empty(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	compiler := func(k, v string) string { return k + "=" + v }
	// Act
	result := h.ToStringsUsingCompiler(compiler)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"ToStringsUsingCompiler empty",
			args.Map{"Length": 0},
		),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_GetValuesExceptKeysInHashset(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	h.AddOrUpdate("b", "2")
	exclude := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	result := h.GetValuesExceptKeysInHashset(exclude)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"GetValuesExceptKeysInHashset",
			args.Map{"Length": 1},
		),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_GetValuesExceptKeysInHashset_Nil(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	// Act
	result := h.GetValuesExceptKeysInHashset(nil)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"GetValuesExceptKeysInHashset nil",
			args.Map{"Length": 1},
		),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_GetValuesKeysExcept(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	h.AddOrUpdate("b", "2")
	// Act
	result := h.GetValuesKeysExcept([]string{"a"})
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"GetValuesKeysExcept",
			args.Map{"Length": 1},
		),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_GetValuesKeysExcept_Nil(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	// Act
	result := h.GetValuesKeysExcept(nil)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"GetValuesKeysExcept nil",
			args.Map{"Length": 1},
		),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_GetAllExceptCollection(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	h.AddOrUpdate("b", "2")
	col := corestr.New.Collection.Strings([]string{"a"})
	// Act
	result := h.GetAllExceptCollection(col)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"GetAllExceptCollection",
			args.Map{"Length": 1},
		),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_GetAllExceptCollection_Nil(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	// Act
	result := h.GetAllExceptCollection(nil)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"GetAllExceptCollection nil",
			args.Map{"Length": 1},
		),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var h *corestr.Hashmap
	// Act
	result := h.ClonePtr()
	// Assert
	convey.Convey("ClonePtr nil", t, func() {
		convey.So(result, convey.ShouldBeNil)
	})
}

func Test_Cov55_Hashmap_IsEqual(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	other := corestr.New.Hashmap.Empty()
	other.AddOrUpdate("a", "1")
	// Act
	result := h.IsEqual(*other)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"IsEqual",
			args.Map{"IsEqual": true},
		),
	}
	actual := args.Map{"IsEqual": result}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_IsEqualPtrLock(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	// Act
	result := h.IsEqualPtrLock(h)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"IsEqualPtrLock same ptr",
			args.Map{"IsEqual": true},
		),
	}
	actual := args.Map{"IsEqual": result}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_Remove(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	// Act
	h.Remove("a")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"Remove",
			args.Map{"IsEmpty": true},
		),
	}
	actual := args.Map{"IsEmpty": h.IsEmpty()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_RemoveWithLock(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	// Act
	h.RemoveWithLock("a")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields(
			"RemoveWithLock",
			args.Map{"IsEmpty": true},
		),
	}
	actual := args.Map{"IsEmpty": h.IsEmpty()}
	tc.ShouldBeEqualMap(0, t, actual)
}

// =============================================================================
// emptyCreator — all factory methods
// =============================================================================

func Test_Cov55_EmptyCreator_LinkedList(t *testing.T) {
	// Arrange // Act
	ll := corestr.Empty.LinkedList()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("Empty.LinkedList", args.Map{"IsEmpty": true}),
	}
	actual := args.Map{"IsEmpty": ll.IsEmpty()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_EmptyCreator_SimpleSlice(t *testing.T) {
	// Arrange // Act
	ss := corestr.Empty.SimpleSlice()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("Empty.SimpleSlice", args.Map{"IsEmpty": true}),
	}
	actual := args.Map{"IsEmpty": ss.IsEmpty()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_EmptyCreator_KeyAnyValuePair(t *testing.T) {
	// Arrange // Act
	p := corestr.Empty.KeyAnyValuePair()
	// Assert
	convey.Convey("Empty.KeyAnyValuePair", t, func() {
		convey.So(p, convey.ShouldNotBeNil)
		convey.So(p.Key, convey.ShouldBeEmpty)
	})
}

func Test_Cov55_EmptyCreator_KeyValuePair(t *testing.T) {
	// Arrange // Act
	p := corestr.Empty.KeyValuePair()
	// Assert
	convey.Convey("Empty.KeyValuePair", t, func() {
		convey.So(p, convey.ShouldNotBeNil)
		convey.So(p.Key, convey.ShouldBeEmpty)
	})
}

func Test_Cov55_EmptyCreator_KeyValueCollection(t *testing.T) {
	// Arrange // Act
	c := corestr.Empty.KeyValueCollection()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("Empty.KeyValueCollection", args.Map{"IsEmpty": true}),
	}
	actual := args.Map{"IsEmpty": c.IsEmpty()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_EmptyCreator_LinkedCollections(t *testing.T) {
	// Arrange // Act
	lc := corestr.Empty.LinkedCollections()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("Empty.LinkedCollections", args.Map{"IsEmpty": true}),
	}
	actual := args.Map{"IsEmpty": lc.IsEmpty()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_EmptyCreator_LeftRight(t *testing.T) {
	// Arrange // Act
	lr := corestr.Empty.LeftRight()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("Empty.LeftRight", args.Map{"IsEmpty": true}),
	}
	actual := args.Map{"IsEmpty": lr.IsEmpty()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_EmptyCreator_Hashset(t *testing.T) {
	// Arrange // Act
	hs := corestr.Empty.Hashset()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("Empty.Hashset", args.Map{"IsEmpty": true}),
	}
	actual := args.Map{"IsEmpty": hs.IsEmpty()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_EmptyCreator_HashsetsCollection(t *testing.T) {
	// Arrange // Act
	hc := corestr.Empty.HashsetsCollection()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("Empty.HashsetsCollection", args.Map{"IsEmpty": true}),
	}
	actual := args.Map{"IsEmpty": hc.IsEmpty()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_EmptyCreator_Hashmap(t *testing.T) {
	// Arrange // Act
	hm := corestr.Empty.Hashmap()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("Empty.Hashmap", args.Map{"IsEmpty": true}),
	}
	actual := args.Map{"IsEmpty": hm.IsEmpty()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_EmptyCreator_CharCollectionMap(t *testing.T) {
	// Arrange // Act
	ccm := corestr.Empty.CharCollectionMap()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("Empty.CharCollectionMap", args.Map{"IsEmpty": true}),
	}
	actual := args.Map{"IsEmpty": ccm.IsEmpty()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_EmptyCreator_KeyValuesCollection(t *testing.T) {
	// Arrange // Act
	kvc := corestr.Empty.KeyValuesCollection()
	// Assert
	convey.Convey("Empty.KeyValuesCollection", t, func() {
		convey.So(kvc, convey.ShouldNotBeNil)
	})
}

func Test_Cov55_EmptyCreator_CollectionsOfCollection(t *testing.T) {
	// Arrange // Act
	coc := corestr.Empty.CollectionsOfCollection()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("Empty.CollectionsOfCollection", args.Map{"IsEmpty": true}),
	}
	actual := args.Map{"IsEmpty": coc.IsEmpty()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_EmptyCreator_CharHashsetMap(t *testing.T) {
	// Arrange // Act
	chm := corestr.Empty.CharHashsetMap()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("Empty.CharHashsetMap", args.Map{"IsEmpty": true}),
	}
	actual := args.Map{"IsEmpty": chm.IsEmpty()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_EmptyCreator_SimpleStringOnce(t *testing.T) {
	// Arrange // Act
	sso := corestr.Empty.SimpleStringOnce()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("Empty.SimpleStringOnce", args.Map{"IsInit": false}),
	}
	actual := args.Map{"IsInit": sso.IsInitialized()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_EmptyCreator_SimpleStringOncePtr(t *testing.T) {
	// Arrange // Act
	sso := corestr.Empty.SimpleStringOncePtr()
	// Assert
	convey.Convey("Empty.SimpleStringOncePtr", t, func() {
		convey.So(sso, convey.ShouldNotBeNil)
		convey.So(sso.IsInitialized(), convey.ShouldBeFalse)
	})
}

// =============================================================================
// StringUtils — WrapDouble, WrapSingle, WrapTilda, WrapDoubleIfMissing, WrapSingleIfMissing
// =============================================================================

func Test_Cov55_StringUtils_WrapDouble(t *testing.T) {
	// Arrange // Act
	result := corestr.StringUtils.WrapDouble("hello")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("WrapDouble", "\"hello\""),
	}
	tc.ShouldBeEqual(0, t, result)
}

func Test_Cov55_StringUtils_WrapSingle(t *testing.T) {
	// Arrange // Act
	result := corestr.StringUtils.WrapSingle("hello")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("WrapSingle", "'hello'"),
	}
	tc.ShouldBeEqual(0, t, result)
}

func Test_Cov55_StringUtils_WrapTilda(t *testing.T) {
	// Arrange // Act
	result := corestr.StringUtils.WrapTilda("hello")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("WrapTilda", "`hello`"),
	}
	tc.ShouldBeEqual(0, t, result)
}

func Test_Cov55_StringUtils_WrapDoubleIfMissing_Empty(t *testing.T) {
	// Arrange // Act
	result := corestr.StringUtils.WrapDoubleIfMissing("")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("WrapDoubleIfMissing empty", "\"\""),
	}
	tc.ShouldBeEqual(0, t, result)
}

func Test_Cov55_StringUtils_WrapDoubleIfMissing_AlreadyWrapped(t *testing.T) {
	// Arrange // Act
	result := corestr.StringUtils.WrapDoubleIfMissing("\"hello\"")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("WrapDoubleIfMissing already", "\"hello\""),
	}
	tc.ShouldBeEqual(0, t, result)
}

func Test_Cov55_StringUtils_WrapDoubleIfMissing_NotWrapped(t *testing.T) {
	// Arrange // Act
	result := corestr.StringUtils.WrapDoubleIfMissing("hello")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("WrapDoubleIfMissing not wrapped", "\"hello\""),
	}
	tc.ShouldBeEqual(0, t, result)
}

func Test_Cov55_StringUtils_WrapSingleIfMissing_Empty(t *testing.T) {
	// Arrange // Act
	result := corestr.StringUtils.WrapSingleIfMissing("")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("WrapSingleIfMissing empty", "''"),
	}
	tc.ShouldBeEqual(0, t, result)
}

func Test_Cov55_StringUtils_WrapSingleIfMissing_AlreadyWrapped(t *testing.T) {
	// Arrange // Act
	result := corestr.StringUtils.WrapSingleIfMissing("'hello'")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("WrapSingleIfMissing already", "'hello'"),
	}
	tc.ShouldBeEqual(0, t, result)
}

func Test_Cov55_StringUtils_WrapSingleIfMissing_NotWrapped(t *testing.T) {
	// Arrange // Act
	result := corestr.StringUtils.WrapSingleIfMissing("hello")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("WrapSingleIfMissing not wrapped", "'hello'"),
	}
	tc.ShouldBeEqual(0, t, result)
}

func Test_Cov55_StringUtils_WrapDoubleIfMissing_DoubleEmpty(t *testing.T) {
	// Arrange // Act
	result := corestr.StringUtils.WrapDoubleIfMissing("\"\"")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("WrapDoubleIfMissing double-empty", "\"\""),
	}
	tc.ShouldBeEqual(0, t, result)
}

func Test_Cov55_StringUtils_WrapSingleIfMissing_SingleEmpty(t *testing.T) {
	// Arrange // Act
	result := corestr.StringUtils.WrapSingleIfMissing("''")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("WrapSingleIfMissing single-empty", "''"),
	}
	tc.ShouldBeEqual(0, t, result)
}

// =============================================================================
// Hashmap — DiffRaw, Collection, SetTrim, SetBySplitter
// =============================================================================

func Test_Cov55_Hashmap_DiffRaw(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	// Act
	result := h.DiffRaw(map[string]string{"a": "2"})
	// Assert
	convey.Convey("DiffRaw", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

func Test_Cov55_Hashmap_Collection(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	// Act
	col := h.Collection()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("Hashmap.Collection", args.Map{"Length": 1}),
	}
	actual := args.Map{"Length": col.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_SetTrim(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	// Act
	isNew := h.SetTrim("  key  ", "  val  ")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("SetTrim", args.Map{"IsNew": true, "Has": true}),
	}
	actual := args.Map{"IsNew": isNew, "Has": h.Has("key")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_SetBySplitter(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	// Act
	isNew := h.SetBySplitter("=", "key=value")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("SetBySplitter", args.Map{"IsNew": true, "Has": true}),
	}
	actual := args.Map{"IsNew": isNew, "Has": h.Has("key")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_SetBySplitter_NoSplit(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	// Act
	isNew := h.SetBySplitter("=", "nosplit")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("SetBySplitter no split", args.Map{"IsNew": true, "Has": true}),
	}
	actual := args.Map{"IsNew": isNew, "Has": h.Has("nosplit")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_KeysCollection(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("a", "1")
	// Act
	col := h.KeysCollection()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("KeysCollection", args.Map{"Length": 1}),
	}
	actual := args.Map{"Length": col.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_ParseInjectUsingJsonMust(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("k", "v")
	jsonResult := h.JsonPtr()
	target := corestr.New.Hashmap.Empty()
	// Act
	result := target.ParseInjectUsingJsonMust(jsonResult)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("ParseInjectUsingJsonMust", args.Map{"Has": true}),
	}
	actual := args.Map{"Has": result.Has("k")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov55_Hashmap_Clone_Empty(t *testing.T) {
	// Arrange
	h := corestr.New.Hashmap.Empty()
	// Act
	cloned := h.Clone()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("Clone empty", args.Map{"IsEmpty": true}),
	}
	actual := args.Map{"IsEmpty": cloned.IsEmpty()}
	tc.ShouldBeEqualMap(0, t, actual)
}
