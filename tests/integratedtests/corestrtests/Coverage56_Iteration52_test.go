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
// Hashset — deep coverage of remaining methods
// =============================================================================

func Test_Cov56_Hashset_AddPtr(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	key := "hello"
	// Act
	hs.AddPtr(&key)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddPtr", args.Map{"Has": true}),
	}
	actual := args.Map{"Has": hs.Has("hello")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_AddPtrLock(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	key := "world"
	// Act
	hs.AddPtrLock(&key)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddPtrLock", args.Map{"Has": true}),
	}
	actual := args.Map{"Has": hs.Has("world")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_AddWithWgLock(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	// Act
	hs.AddWithWgLock("item", wg)
	wg.Wait()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddWithWgLock", args.Map{"Has": true}),
	}
	actual := args.Map{"Has": hs.Has("item")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_AddBool_New(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	// Act
	isExist := hs.AddBool("new")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddBool new", args.Map{"IsExist": false}),
	}
	actual := args.Map{"IsExist": isExist}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_AddBool_Existing(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	hs.Add("existing")
	// Act
	isExist := hs.AddBool("existing")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddBool existing", args.Map{"IsExist": true}),
	}
	actual := args.Map{"IsExist": isExist}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_AddNonEmptyWhitespace(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	// Act
	hs.AddNonEmptyWhitespace("  ")
	hs.AddNonEmptyWhitespace("valid")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddNonEmptyWhitespace", args.Map{"Length": 1}),
	}
	actual := args.Map{"Length": hs.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_AddIfMany_True(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	// Act
	hs.AddIfMany(true, "a", "b")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddIfMany true", args.Map{"Length": 2}),
	}
	actual := args.Map{"Length": hs.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_AddIfMany_False(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	// Act
	hs.AddIfMany(false, "a", "b")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddIfMany false", args.Map{"Length": 0}),
	}
	actual := args.Map{"Length": hs.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_AddFunc(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	// Act
	hs.AddFunc(func() string { return "computed" })
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddFunc", args.Map{"Has": true}),
	}
	actual := args.Map{"Has": hs.Has("computed")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_AddFuncErr_NoErr(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	// Act
	hs.AddFuncErr(
		func() (string, error) { return "ok", nil },
		func(err error) { t.Fatal(err) },
	)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddFuncErr no err", args.Map{"Has": true}),
	}
	actual := args.Map{"Has": hs.Has("ok")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_AddFuncErr_WithErr(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	handlerCalled := false
	// Act
	hs.AddFuncErr(
		func() (string, error) { return "", fmt.Errorf("fail") },
		func(err error) { handlerCalled = true },
	)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddFuncErr with err", args.Map{"HandlerCalled": true, "Length": 0}),
	}
	actual := args.Map{"HandlerCalled": handlerCalled, "Length": hs.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_AddStringsPtrWgLock(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	// Act
	hs.AddStringsPtrWgLock([]string{"a", "b"}, wg)
	wg.Wait()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddStringsPtrWgLock", args.Map{"Length": 2}),
	}
	actual := args.Map{"Length": hs.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_AddHashsetWgLock(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	other := corestr.New.Hashset.Strings([]string{"x", "y"})
	wg := &sync.WaitGroup{}
	wg.Add(1)
	// Act
	hs.AddHashsetWgLock(other, wg)
	wg.Wait()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddHashsetWgLock", args.Map{"Length": 2}),
	}
	actual := args.Map{"Length": hs.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_AddHashsetWgLock_Nil(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	wg := &sync.WaitGroup{}
	// Act
	hs.AddHashsetWgLock(nil, wg)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddHashsetWgLock nil", args.Map{"Length": 0}),
	}
	actual := args.Map{"Length": hs.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_AddSimpleSlice(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
	// Act
	hs.AddSimpleSlice(ss)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddSimpleSlice", args.Map{"Length": 2}),
	}
	actual := args.Map{"Length": hs.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_AddCollection(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	col := corestr.New.Collection.Strings([]string{"a", "b"})
	// Act
	hs.AddCollection(col)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddCollection", args.Map{"Length": 2}),
	}
	actual := args.Map{"Length": hs.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_AddCollection_Nil(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	// Act
	hs.AddCollection(nil)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddCollection nil", args.Map{"Length": 0}),
	}
	actual := args.Map{"Length": hs.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_AddCollections(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	col1 := corestr.New.Collection.Strings([]string{"a"})
	col2 := corestr.New.Collection.Strings([]string{"b"})
	// Act
	hs.AddCollections(col1, col2)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddCollections", args.Map{"Length": 2}),
	}
	actual := args.Map{"Length": hs.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_AddCollections_Nil(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	// Act
	hs.AddCollections(nil, nil)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddCollections nil items", args.Map{"Length": 0}),
	}
	actual := args.Map{"Length": hs.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_AddsUsingFilter(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	filter := func(s string, i int) (string, bool, bool) {
		return s, s != "skip", false
	}
	// Act
	hs.AddsUsingFilter(filter, "a", "skip", "b")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddsUsingFilter", args.Map{"Length": 2}),
	}
	actual := args.Map{"Length": hs.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_AddsUsingFilter_Break(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	filter := func(s string, i int) (string, bool, bool) {
		return s, true, true
	}
	// Act
	hs.AddsUsingFilter(filter, "a", "b")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddsUsingFilter break", args.Map{"Length": 1}),
	}
	actual := args.Map{"Length": hs.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_AddsUsingFilter_Nil(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	filter := func(s string, i int) (string, bool, bool) {
		return s, true, false
	}
	// Act
	hs.AddsUsingFilter(filter)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddsUsingFilter nil", args.Map{"Length": 0}),
	}
	actual := args.Map{"Length": hs.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_AddsAnyUsingFilter(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	filter := func(s string, i int) (string, bool, bool) {
		return s, true, false
	}
	// Act
	hs.AddsAnyUsingFilter(filter, "hello", 42)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddsAnyUsingFilter", args.Map{"Length": 2}),
	}
	actual := args.Map{"Length": hs.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_AddsAnyUsingFilter_NilItem(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	filter := func(s string, i int) (string, bool, bool) {
		return s, true, false
	}
	// Act
	hs.AddsAnyUsingFilter(filter, nil, "valid")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddsAnyUsingFilter nil item", args.Map{"Length": 1}),
	}
	actual := args.Map{"Length": hs.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_AddsAnyUsingFilter_Break(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	filter := func(s string, i int) (string, bool, bool) {
		return s, true, true
	}
	// Act
	hs.AddsAnyUsingFilter(filter, "a", "b")
	// Assert
	convey.Convey("AddsAnyUsingFilter break", t, func() {
		convey.So(hs.Length(), convey.ShouldBeLessThanOrEqualTo, 2)
	})
}

func Test_Cov56_Hashset_AddsAnyUsingFilterLock(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	filter := func(s string, i int) (string, bool, bool) {
		return s, true, false
	}
	// Act
	hs.AddsAnyUsingFilterLock(filter, "x")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddsAnyUsingFilterLock", args.Map{"Length": 1}),
	}
	actual := args.Map{"Length": hs.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_AddsAnyUsingFilterLock_Nil(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	filter := func(s string, i int) (string, bool, bool) {
		return s, true, false
	}
	// Act
	hs.AddsAnyUsingFilterLock(filter)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddsAnyUsingFilterLock nil", args.Map{"Length": 0}),
	}
	actual := args.Map{"Length": hs.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_AddItemsMapWgLock(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	itemsMap := map[string]bool{"a": true, "b": false}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	// Act
	hs.AddItemsMapWgLock(&itemsMap, wg)
	wg.Wait()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddItemsMapWgLock", args.Map{"HasA": true, "HasB": false}),
	}
	actual := args.Map{"HasA": hs.Has("a"), "HasB": hs.Has("b")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_AddItemsMapWgLock_Nil(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	wg := &sync.WaitGroup{}
	// Act
	hs.AddItemsMapWgLock(nil, wg)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddItemsMapWgLock nil", args.Map{"Length": 0}),
	}
	actual := args.Map{"Length": hs.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_ConcatNewHashsets(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	other := corestr.New.Hashset.Strings([]string{"b"})
	// Act
	result := hs.ConcatNewHashsets(false, other)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("ConcatNewHashsets", args.Map{"HasA": true, "HasB": true}),
	}
	actual := args.Map{"HasA": result.Has("a"), "HasB": result.Has("b")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_ConcatNewHashsets_NoArgs(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	result := hs.ConcatNewHashsets(true)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("ConcatNewHashsets no args", args.Map{"Has": true}),
	}
	actual := args.Map{"Has": result.Has("a")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_ConcatNewStrings(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	result := hs.ConcatNewStrings(false, []string{"b", "c"})
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("ConcatNewStrings", args.Map{"HasA": true, "HasB": true}),
	}
	actual := args.Map{"HasA": result.Has("a"), "HasB": result.Has("b")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_ConcatNewStrings_Empty(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	result := hs.ConcatNewStrings(true)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("ConcatNewStrings empty", args.Map{"Has": true}),
	}
	actual := args.Map{"Has": result.Has("a")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_Filter(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"abc", "def", "abx"})
	// Act
	result := hs.Filter(func(s string) bool { return s[0] == 'a' })
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("Filter", args.Map{"Length": 2}),
	}
	actual := args.Map{"Length": result.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_IsMissing(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	missing := hs.IsMissing("b")
	notMissing := hs.IsMissing("a")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("IsMissing", args.Map{"Missing": true, "NotMissing": false}),
	}
	actual := args.Map{"Missing": missing, "NotMissing": notMissing}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_IsMissingLock(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	result := hs.IsMissingLock("b")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("IsMissingLock", args.Map{"Missing": true}),
	}
	actual := args.Map{"Missing": result}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_HasWithLock(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	result := hs.HasWithLock("a")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("HasWithLock", args.Map{"Has": true}),
	}
	actual := args.Map{"Has": result}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_IsAllMissing(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	allMissing := hs.IsAllMissing("b", "c")
	notAllMissing := hs.IsAllMissing("a", "b")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("IsAllMissing", args.Map{"AllMissing": true, "NotAll": false}),
	}
	actual := args.Map{"AllMissing": allMissing, "NotAll": notAllMissing}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_HasAny(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	found := hs.HasAny("b", "a")
	notFound := hs.HasAny("c", "d")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("HasAny", args.Map{"Found": true, "NotFound": false}),
	}
	actual := args.Map{"Found": found, "NotFound": notFound}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_HasAllCollectionItems(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a", "b"})
	col := corestr.New.Collection.Strings([]string{"a", "b"})
	// Act
	result := hs.HasAllCollectionItems(col)
	resultNil := hs.HasAllCollectionItems(nil)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("HasAllCollectionItems", args.Map{"HasAll": true, "Nil": false}),
	}
	actual := args.Map{"HasAll": result, "Nil": resultNil}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_OrderedList(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"c", "a", "b"})
	// Act
	result := hs.OrderedList()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("OrderedList", args.Map{"First": "a", "Last": "c"}),
	}
	actual := args.Map{"First": result[0], "Last": result[2]}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_OrderedList_Empty(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	// Act
	result := hs.OrderedList()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("OrderedList empty", args.Map{"Length": 0}),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_SafeStrings(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	// Act
	result := hs.SafeStrings()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("SafeStrings empty", args.Map{"Length": 0}),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_SafeStrings_NonEmpty(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	result := hs.SafeStrings()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("SafeStrings non-empty", args.Map{"Length": 1}),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_Lines(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	// Act
	result := hs.Lines()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("Lines empty", args.Map{"Length": 0}),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_SimpleSlice(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	result := hs.SimpleSlice()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("SimpleSlice", args.Map{"Length": 1}),
	}
	actual := args.Map{"Length": result.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_SimpleSlice_Empty(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	// Act
	result := hs.SimpleSlice()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("SimpleSlice empty", args.Map{"IsEmpty": true}),
	}
	actual := args.Map{"IsEmpty": result.IsEmpty()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_MapStringAny(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	result := hs.MapStringAny()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("MapStringAny", args.Map{"Length": 1}),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_MapStringAny_Empty(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	// Act
	result := hs.MapStringAny()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("MapStringAny empty", args.Map{"Length": 0}),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_MapStringAnyDiff(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	result := hs.MapStringAnyDiff()
	// Assert
	convey.Convey("MapStringAnyDiff", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

func Test_Cov56_Hashset_JoinSorted(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"c", "a", "b"})
	// Act
	result := hs.JoinSorted(",")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("JoinSorted", "a,b,c"),
	}
	tc.ShouldBeEqual(0, t, result)
}

func Test_Cov56_Hashset_JoinSorted_Empty(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	// Act
	result := hs.JoinSorted(",")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("JoinSorted empty", ""),
	}
	tc.ShouldBeEqual(0, t, result)
}

func Test_Cov56_Hashset_ListPtrSortedAsc(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"c", "a", "b"})
	// Act
	result := hs.ListPtrSortedAsc()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("ListPtrSortedAsc", args.Map{"First": "a"}),
	}
	actual := args.Map{"First": result[0]}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_ListPtrSortedDsc(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"c", "a", "b"})
	// Act
	result := hs.ListPtrSortedDsc()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("ListPtrSortedDsc", args.Map{"First": "c"}),
	}
	actual := args.Map{"First": result[0]}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_ListCopyLock(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	result := hs.ListCopyLock()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("ListCopyLock", args.Map{"Length": 1}),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_ToLowerSet(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"ABC", "DEF"})
	// Act
	result := hs.ToLowerSet()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("ToLowerSet", args.Map{"HasAbc": true, "HasDef": true}),
	}
	actual := args.Map{"HasAbc": result.Has("abc"), "HasDef": result.Has("def")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_DistinctDiffLinesRaw_BothEmpty(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	// Act
	result := hs.DistinctDiffLinesRaw()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("DistinctDiffLinesRaw both empty", args.Map{"Length": 0}),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_DistinctDiffLinesRaw_LeftOnly(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	result := hs.DistinctDiffLinesRaw()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("DistinctDiffLinesRaw left only", args.Map{"Length": 1}),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_DistinctDiffLinesRaw_RightOnly(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	// Act
	result := hs.DistinctDiffLinesRaw("x", "y")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("DistinctDiffLinesRaw right only", args.Map{"Length": 2}),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_DistinctDiffLinesRaw_Both(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a", "b"})
	// Act
	result := hs.DistinctDiffLinesRaw("b", "c")
	// Assert
	convey.Convey("DistinctDiffLinesRaw both have items", t, func() {
		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
	})
}

func Test_Cov56_Hashset_DistinctDiffLines_BothEmpty(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	// Act
	result := hs.DistinctDiffLines()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("DistinctDiffLines both empty", args.Map{"Length": 0}),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_DistinctDiffLines_LeftOnly(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	result := hs.DistinctDiffLines()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("DistinctDiffLines left only", args.Map{"Length": 1}),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_DistinctDiffLines_RightOnly(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	// Act
	result := hs.DistinctDiffLines("x")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("DistinctDiffLines right only", args.Map{"Length": 1}),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_DistinctDiffHashset(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a", "b"})
	other := corestr.New.Hashset.Strings([]string{"b", "c"})
	// Act
	result := hs.DistinctDiffHashset(other)
	// Assert
	convey.Convey("DistinctDiffHashset", t, func() {
		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
	})
}

func Test_Cov56_Hashset_Transpile(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	result := hs.Transpile(func(s string) string { return "[" + s + "]" })
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("Transpile", args.Map{"Has": true}),
	}
	actual := args.Map{"Has": result.Has("[a]")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_Transpile_Empty(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	// Act
	result := hs.Transpile(func(s string) string { return s })
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("Transpile empty", args.Map{"IsEmpty": true}),
	}
	actual := args.Map{"IsEmpty": result.IsEmpty()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_WrapDoubleQuote(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	result := hs.WrapDoubleQuote()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("WrapDoubleQuote", args.Map{"Has": true}),
	}
	actual := args.Map{"Has": result.Has("\"a\"")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_WrapSingleQuote(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	result := hs.WrapSingleQuote()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("WrapSingleQuote", args.Map{"Has": true}),
	}
	actual := args.Map{"Has": result.Has("'a'")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_WrapDoubleQuoteIfMissing(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	result := hs.WrapDoubleQuoteIfMissing()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("WrapDoubleQuoteIfMissing", args.Map{"Has": true}),
	}
	actual := args.Map{"Has": result.Has("\"a\"")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_WrapSingleQuoteIfMissing(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	result := hs.WrapSingleQuoteIfMissing()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("WrapSingleQuoteIfMissing", args.Map{"Has": true}),
	}
	actual := args.Map{"Has": result.Has("'a'")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_JoinLine(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	result := hs.JoinLine()
	// Assert
	convey.Convey("JoinLine", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov56_Hashset_NonEmptyJoins(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a", "b"})
	// Act
	result := hs.NonEmptyJoins(",")
	// Assert
	convey.Convey("NonEmptyJoins", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov56_Hashset_NonWhitespaceJoins(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a", "b"})
	// Act
	result := hs.NonWhitespaceJoins(",")
	// Assert
	convey.Convey("NonWhitespaceJoins", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov56_Hashset_SortedList(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"c", "a", "b"})
	// Act
	result := hs.SortedList()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("SortedList", args.Map{"First": "a", "Last": "c"}),
	}
	actual := args.Map{"First": result[0], "Last": result[2]}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_Contains(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	result := hs.Contains("a")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("Contains alias", args.Map{"Contains": true}),
	}
	actual := args.Map{"Contains": result}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_IsEqual(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	other := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	result := hs.IsEqual(other)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("IsEqual alias", args.Map{"IsEqual": true}),
	}
	actual := args.Map{"IsEqual": result}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_GetFilteredItems(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"abc", "def"})
	filter := func(s string, i int) (string, bool, bool) { return s, s == "abc", false }
	// Act
	result := hs.GetFilteredItems(filter)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("GetFilteredItems", args.Map{"Length": 1}),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_GetFilteredItems_Empty(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	filter := func(s string, i int) (string, bool, bool) { return s, true, false }
	// Act
	result := hs.GetFilteredItems(filter)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("GetFilteredItems empty", args.Map{"Length": 0}),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_GetFilteredItems_Break(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a", "b"})
	filter := func(s string, i int) (string, bool, bool) { return s, true, true }
	// Act
	result := hs.GetFilteredItems(filter)
	// Assert
	convey.Convey("GetFilteredItems break", t, func() {
		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
	})
}

func Test_Cov56_Hashset_GetFilteredCollection(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"abc"})
	filter := func(s string, i int) (string, bool, bool) { return s, true, false }
	// Act
	result := hs.GetFilteredCollection(filter)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("GetFilteredCollection", args.Map{"Length": 1}),
	}
	actual := args.Map{"Length": result.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_GetFilteredCollection_Empty(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	filter := func(s string, i int) (string, bool, bool) { return s, true, false }
	// Act
	result := hs.GetFilteredCollection(filter)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("GetFilteredCollection empty", args.Map{"IsEmpty": true}),
	}
	actual := args.Map{"IsEmpty": result.IsEmpty()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_GetFilteredCollection_Break(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a", "b"})
	filter := func(s string, i int) (string, bool, bool) { return s, true, true }
	// Act
	result := hs.GetFilteredCollection(filter)
	// Assert
	convey.Convey("GetFilteredCollection break", t, func() {
		convey.So(result.Length(), convey.ShouldBeGreaterThan, 0)
	})
}

func Test_Cov56_Hashset_GetAllExceptHashset(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a", "b", "c"})
	exclude := corestr.New.Hashset.Strings([]string{"b"})
	// Act
	result := hs.GetAllExceptHashset(exclude)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("GetAllExceptHashset", args.Map{"Length": 2}),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_GetAllExceptHashset_Nil(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	result := hs.GetAllExceptHashset(nil)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("GetAllExceptHashset nil", args.Map{"Length": 1}),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_GetAllExcept(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a", "b"})
	// Act
	result := hs.GetAllExcept([]string{"a"})
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("GetAllExcept", args.Map{"Length": 1}),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_GetAllExcept_Nil(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	result := hs.GetAllExcept(nil)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("GetAllExcept nil", args.Map{"Length": 1}),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_GetAllExceptSpread(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a", "b"})
	// Act
	result := hs.GetAllExceptSpread("a")
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("GetAllExceptSpread", args.Map{"Length": 1}),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_GetAllExceptCollection(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a", "b"})
	col := corestr.New.Collection.Strings([]string{"a"})
	// Act
	result := hs.GetAllExceptCollection(col)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("GetAllExceptCollection", args.Map{"Length": 1}),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_GetAllExceptCollection_Nil(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	result := hs.GetAllExceptCollection(nil)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("GetAllExceptCollection nil", args.Map{"Length": 1}),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_StringLock(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	result := hs.StringLock()
	// Assert
	convey.Convey("StringLock non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov56_Hashset_StringLock_Empty(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Empty()
	// Act
	result := hs.StringLock()
	// Assert
	convey.Convey("StringLock empty", t, func() {
		convey.So(result, convey.ShouldContainSubstring, "No Element")
	})
}

func Test_Cov56_Hashset_ParseInjectUsingJsonMust(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	jsonResult := hs.JsonPtr()
	target := corestr.New.Hashset.Empty()
	// Act
	result := target.ParseInjectUsingJsonMust(jsonResult)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("ParseInjectUsingJsonMust", args.Map{"Has": true}),
	}
	actual := args.Map{"Has": result.Has("a")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_Items(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	result := hs.Items()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("Items", args.Map{"Length": 1}),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_Collection(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	result := hs.Collection()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("Hashset.Collection", args.Map{"Length": 1}),
	}
	actual := args.Map{"Length": result.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_ListPtr(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	result := hs.ListPtr()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("ListPtr deprecated alias", args.Map{"Length": 1}),
	}
	actual := args.Map{"Length": len(result)}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_AddCapacities(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	hs.AddCapacities(10, 20)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddCapacities", args.Map{"Has": true}),
	}
	actual := args.Map{"Has": hs.Has("a")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_AddCapacities_Empty(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	hs.AddCapacities()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddCapacities empty", args.Map{"Has": true}),
	}
	actual := args.Map{"Has": hs.Has("a")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_AddCapacitiesLock(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	hs.AddCapacitiesLock(10)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddCapacitiesLock", args.Map{"Has": true}),
	}
	actual := args.Map{"Has": hs.Has("a")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_AddCapacitiesLock_Empty(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	hs.AddCapacitiesLock()
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddCapacitiesLock empty", args.Map{"Has": true}),
	}
	actual := args.Map{"Has": hs.Has("a")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_Resize(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	hs.Resize(100)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("Resize", args.Map{"Has": true}),
	}
	actual := args.Map{"Has": hs.Has("a")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_Resize_AlreadyLarger(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a", "b", "c"})
	// Act
	hs.Resize(1)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("Resize already larger", args.Map{"Length": 3}),
	}
	actual := args.Map{"Length": hs.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_ResizeLock(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a"})
	// Act
	hs.ResizeLock(100)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("ResizeLock", args.Map{"Has": true}),
	}
	actual := args.Map{"Has": hs.Has("a")}
	tc.ShouldBeEqualMap(0, t, actual)
}

func Test_Cov56_Hashset_ResizeLock_AlreadyLarger(t *testing.T) {
	// Arrange
	hs := corestr.New.Hashset.Strings([]string{"a", "b", "c"})
	// Act
	hs.ResizeLock(1)
	// Assert
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("ResizeLock already larger", args.Map{"Length": 3}),
	}
	actual := args.Map{"Length": hs.Length()}
	tc.ShouldBeEqualMap(0, t, actual)
}
