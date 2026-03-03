package coregenerictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coregeneric"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/errcore"
)

var hashmapTestCases = []coretestcases.CaseV1{
	{Title: "EmptyHashmap creates empty", ArrangeInput: args.Map{"case": "empty"}, ExpectedInput: []string{"true", "0", "false"}},
	{Title: "NewHashmap with capacity", ArrangeInput: args.Map{"case": "new"}, ExpectedInput: []string{"true"}},
	{Title: "HashmapFrom wraps map", ArrangeInput: args.Map{"case": "from"}, ExpectedInput: []string{"2", "true"}},
	{Title: "HashmapClone independence", ArrangeInput: args.Map{"case": "clone-func"}, ExpectedInput: []string{"1", "99"}},
	{Title: "Set new key returns true", ArrangeInput: args.Map{"case": "set-new"}, ExpectedInput: []string{"true", "1"}},
	{Title: "Set existing key returns false", ArrangeInput: args.Map{"case": "set-existing"}, ExpectedInput: []string{"false", "2"}},
	{Title: "Get found", ArrangeInput: args.Map{"case": "get-found"}, ExpectedInput: []string{"true", "42"}},
	{Title: "Get not found", ArrangeInput: args.Map{"case": "get-notfound"}, ExpectedInput: []string{"false", "0"}},
	{Title: "GetOrDefault missing returns default", ArrangeInput: args.Map{"case": "getordefault-missing"}, ExpectedInput: []string{"99"}},
	{Title: "GetOrDefault found returns value", ArrangeInput: args.Map{"case": "getordefault-found"}, ExpectedInput: []string{"5"}},
	{Title: "Has/Contains/IsKeyMissing", ArrangeInput: args.Map{"case": "has"}, ExpectedInput: []string{"true", "true", "false"}},
	{Title: "IsKeyMissing true", ArrangeInput: args.Map{"case": "iskeymissing"}, ExpectedInput: []string{"true"}},
	{Title: "Remove existing", ArrangeInput: args.Map{"case": "remove-existing"}, ExpectedInput: []string{"true", "true"}},
	{Title: "Remove missing", ArrangeInput: args.Map{"case": "remove-missing"}, ExpectedInput: []string{"false"}},
	{Title: "Keys returns all", ArrangeInput: args.Map{"case": "keys"}, ExpectedInput: []string{"2"}},
	{Title: "Keys empty", ArrangeInput: args.Map{"case": "keys-empty"}, ExpectedInput: []string{"0"}},
	{Title: "Values returns all", ArrangeInput: args.Map{"case": "values"}, ExpectedInput: []string{"1", "1"}},
	{Title: "Values empty", ArrangeInput: args.Map{"case": "values-empty"}, ExpectedInput: []string{"0"}},
	{Title: "AddOrUpdateMap merges", ArrangeInput: args.Map{"case": "addorupdatemap"}, ExpectedInput: []string{"2", "10"}},
	{Title: "AddOrUpdateMap empty noop", ArrangeInput: args.Map{"case": "addorupdatemap-empty"}, ExpectedInput: []string{"1"}},
	{Title: "AddOrUpdateHashmap merges", ArrangeInput: args.Map{"case": "addorupdatehashmap"}, ExpectedInput: []string{"2"}},
	{Title: "AddOrUpdateHashmap nil noop", ArrangeInput: args.Map{"case": "addorupdatehashmap-nil"}, ExpectedInput: []string{"1"}},
	{Title: "ConcatNew merged copy", ArrangeInput: args.Map{"case": "concatnew"}, ExpectedInput: []string{"2", "1"}},
	{Title: "ConcatNew nil", ArrangeInput: args.Map{"case": "concatnew-nil"}, ExpectedInput: []string{"1"}},
	{Title: "Clone method independence", ArrangeInput: args.Map{"case": "clone-method"}, ExpectedInput: []string{"1"}},
	{Title: "IsEquals same content", ArrangeInput: args.Map{"case": "isequals-same"}, ExpectedInput: []string{"true"}},
	{Title: "IsEquals different keys", ArrangeInput: args.Map{"case": "isequals-diffkeys"}, ExpectedInput: []string{"false"}},
	{Title: "IsEquals different length", ArrangeInput: args.Map{"case": "isequals-difflen"}, ExpectedInput: []string{"false"}},
	{Title: "IsEquals both nil", ArrangeInput: args.Map{"case": "isequals-bothnil"}, ExpectedInput: []string{"true"}},
	{Title: "IsEquals one nil", ArrangeInput: args.Map{"case": "isequals-onenil"}, ExpectedInput: []string{"false"}},
	{Title: "IsEquals same pointer", ArrangeInput: args.Map{"case": "isequals-sameptr"}, ExpectedInput: []string{"true"}},
	{Title: "ForEach visits all", ArrangeInput: args.Map{"case": "foreach"}, ExpectedInput: []string{"2"}},
	{Title: "ForEachBreak stops early", ArrangeInput: args.Map{"case": "foreachbreak"}, ExpectedInput: []string{"2"}},
	{Title: "String not empty", ArrangeInput: args.Map{"case": "string"}, ExpectedInput: []string{"true"}},
	{Title: "IsEmpty nil receiver", ArrangeInput: args.Map{"case": "isempty-nil"}, ExpectedInput: []string{"true"}},
	{Title: "Length nil receiver", ArrangeInput: args.Map{"case": "length-nil"}, ExpectedInput: []string{"0"}},
	{Title: "HasItems nil receiver", ArrangeInput: args.Map{"case": "hasitems-nil"}, ExpectedInput: []string{"false"}},
}

func Test_Hashmap_Verification(t *testing.T) {
	for caseIndex, tc := range hashmapTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		caseType := input["case"].(string)

		var actLines []string

		// Act
		switch caseType {
		case "empty":
			hm := coregeneric.EmptyHashmap[string, int]()
			actLines = []string{fmt.Sprintf("%v", hm.IsEmpty()), fmt.Sprintf("%v", hm.Length()), fmt.Sprintf("%v", hm.HasItems())}
		case "new":
			hm := coregeneric.NewHashmap[string, int](10)
			actLines = []string{fmt.Sprintf("%v", hm.IsEmpty())}
		case "from":
			hm := coregeneric.HashmapFrom(map[string]int{"a": 1, "b": 2})
			actLines = []string{fmt.Sprintf("%v", hm.Length()), fmt.Sprintf("%v", hm.Has("a"))}
		case "clone-func":
			orig := coregeneric.HashmapFrom(map[string]int{"k": 1})
			cloned := coregeneric.HashmapClone(orig.Map())
			cloned.Set("k", 99)
			origVal, _ := orig.Get("k")
			clonedVal, _ := cloned.Get("k")
			actLines = []string{fmt.Sprintf("%v", origVal), fmt.Sprintf("%v", clonedVal)}
		case "set-new":
			hm := coregeneric.EmptyHashmap[string, int]()
			isNew := hm.Set("key", 42)
			actLines = []string{fmt.Sprintf("%v", isNew), fmt.Sprintf("%v", hm.Length())}
		case "set-existing":
			hm := coregeneric.HashmapFrom(map[string]int{"key": 1})
			isNew := hm.Set("key", 2)
			val, _ := hm.Get("key")
			actLines = []string{fmt.Sprintf("%v", isNew), fmt.Sprintf("%v", val)}
		case "get-found":
			hm := coregeneric.HashmapFrom(map[string]int{"k": 42})
			val, found := hm.Get("k")
			actLines = []string{fmt.Sprintf("%v", found), fmt.Sprintf("%v", val)}
		case "get-notfound":
			hm := coregeneric.EmptyHashmap[string, int]()
			val, found := hm.Get("missing")
			actLines = []string{fmt.Sprintf("%v", found), fmt.Sprintf("%v", val)}
		case "getordefault-missing":
			hm := coregeneric.EmptyHashmap[string, int]()
			actLines = []string{fmt.Sprintf("%v", hm.GetOrDefault("x", 99))}
		case "getordefault-found":
			hm := coregeneric.HashmapFrom(map[string]int{"x": 5})
			actLines = []string{fmt.Sprintf("%v", hm.GetOrDefault("x", 99))}
		case "has":
			hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
			actLines = []string{fmt.Sprintf("%v", hm.Has("a")), fmt.Sprintf("%v", hm.Contains("a")), fmt.Sprintf("%v", hm.IsKeyMissing("a"))}
		case "iskeymissing":
			hm := coregeneric.EmptyHashmap[string, int]()
			actLines = []string{fmt.Sprintf("%v", hm.IsKeyMissing("x"))}
		case "remove-existing":
			hm := coregeneric.HashmapFrom(map[string]int{"k": 1})
			existed := hm.Remove("k")
			actLines = []string{fmt.Sprintf("%v", existed), fmt.Sprintf("%v", hm.IsEmpty())}
		case "remove-missing":
			hm := coregeneric.EmptyHashmap[string, int]()
			actLines = []string{fmt.Sprintf("%v", hm.Remove("x"))}
		case "keys":
			hm := coregeneric.HashmapFrom(map[int]string{1: "a", 2: "b"})
			actLines = []string{fmt.Sprintf("%v", len(hm.Keys()))}
		case "keys-empty":
			hm := coregeneric.EmptyHashmap[int, string]()
			actLines = []string{fmt.Sprintf("%v", len(hm.Keys()))}
		case "values":
			hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
			vals := hm.Values()
			actLines = []string{fmt.Sprintf("%v", len(vals)), fmt.Sprintf("%v", vals[0])}
		case "values-empty":
			hm := coregeneric.EmptyHashmap[string, int]()
			actLines = []string{fmt.Sprintf("%v", len(hm.Values()))}
		case "addorupdatemap":
			hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
			hm.AddOrUpdateMap(map[string]int{"b": 2, "a": 10})
			val, _ := hm.Get("a")
			actLines = []string{fmt.Sprintf("%v", hm.Length()), fmt.Sprintf("%v", val)}
		case "addorupdatemap-empty":
			hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
			hm.AddOrUpdateMap(map[string]int{})
			actLines = []string{fmt.Sprintf("%v", hm.Length())}
		case "addorupdatehashmap":
			hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
			hm.AddOrUpdateHashmap(coregeneric.HashmapFrom(map[string]int{"b": 2}))
			actLines = []string{fmt.Sprintf("%v", hm.Length())}
		case "addorupdatehashmap-nil":
			hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
			hm.AddOrUpdateHashmap(nil)
			actLines = []string{fmt.Sprintf("%v", hm.Length())}
		case "concatnew":
			hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1})
			hm2 := coregeneric.HashmapFrom(map[string]int{"b": 2})
			result := hm1.ConcatNew(hm2)
			actLines = []string{fmt.Sprintf("%v", result.Length()), fmt.Sprintf("%v", hm1.Length())}
		case "concatnew-nil":
			hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
			result := hm.ConcatNew(nil)
			actLines = []string{fmt.Sprintf("%v", result.Length())}
		case "clone-method":
			hm := coregeneric.HashmapFrom(map[string]int{"k": 1})
			cloned := hm.Clone()
			cloned.Set("k", 99)
			origVal, _ := hm.Get("k")
			actLines = []string{fmt.Sprintf("%v", origVal)}
		case "isequals-same":
			hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1, "b": 2})
			hm2 := coregeneric.HashmapFrom(map[string]int{"a": 1, "b": 2})
			actLines = []string{fmt.Sprintf("%v", hm1.IsEquals(hm2))}
		case "isequals-diffkeys":
			hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1})
			hm2 := coregeneric.HashmapFrom(map[string]int{"b": 1})
			actLines = []string{fmt.Sprintf("%v", hm1.IsEquals(hm2))}
		case "isequals-difflen":
			hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1})
			hm2 := coregeneric.HashmapFrom(map[string]int{"a": 1, "b": 2})
			actLines = []string{fmt.Sprintf("%v", hm1.IsEquals(hm2))}
		case "isequals-bothnil":
			var hm1, hm2 *coregeneric.Hashmap[string, int]
			actLines = []string{fmt.Sprintf("%v", hm1.IsEquals(hm2))}
		case "isequals-onenil":
			hm := coregeneric.EmptyHashmap[string, int]()
			actLines = []string{fmt.Sprintf("%v", hm.IsEquals(nil))}
		case "isequals-sameptr":
			hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
			actLines = []string{fmt.Sprintf("%v", hm.IsEquals(hm))}
		case "foreach":
			hm := coregeneric.HashmapFrom(map[string]int{"a": 1, "b": 2})
			count := 0
			hm.ForEach(func(_ string, _ int) { count++ })
			actLines = []string{fmt.Sprintf("%v", count)}
		case "foreachbreak":
			hm := coregeneric.HashmapFrom(map[int]int{1: 1, 2: 2, 3: 3})
			count := 0
			hm.ForEachBreak(func(_ int, _ int) bool { count++; return count >= 2 })
			actLines = []string{fmt.Sprintf("%v", count)}
		case "string":
			hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
			actLines = []string{fmt.Sprintf("%v", hm.String() != "")}
		case "isempty-nil":
			var hm *coregeneric.Hashmap[string, int]
			actLines = []string{fmt.Sprintf("%v", hm.IsEmpty())}
		case "length-nil":
			var hm *coregeneric.Hashmap[string, int]
			actLines = []string{fmt.Sprintf("%v", hm.Length())}
		case "hasitems-nil":
			var hm *coregeneric.Hashmap[string, int]
			actLines = []string{fmt.Sprintf("%v", hm.HasItems())}
		}

		expectedLines := tc.ExpectedInput.([]string)

		// Assert
		errcore.PrintLineDiff(caseIndex, tc.Title, actLines, expectedLines)
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
