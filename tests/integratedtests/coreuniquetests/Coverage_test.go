package coreuniquetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coreunique/intunique"
)

// ═══════════════════════════════════════════
// GetMap
// ═══════════════════════════════════════════

func Test_Cov_GetMap_NilSlice(t *testing.T) {
	result := intunique.GetMap(nil)
	actual := args.Map{"isNil": result == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "GetMap returns nil -- nil input", actual)
}

func Test_Cov_GetMap_EmptySlice(t *testing.T) {
	input := []int{}
	result := intunique.GetMap(&input)
	actual := args.Map{"notNil": result != nil, "len": len(*result)}
	expected := args.Map{"notNil": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "GetMap returns empty map -- empty slice", actual)
}

func Test_Cov_GetMap_WithDuplicates(t *testing.T) {
	input := []int{1, 2, 2, 3, 3, 3}
	result := intunique.GetMap(&input)
	actual := args.Map{
		"notNil": result != nil,
		"len":    len(*result),
		"has1":   (*result)[1],
		"has2":   (*result)[2],
		"has3":   (*result)[3],
	}
	expected := args.Map{
		"notNil": true,
		"len":    3,
		"has1":   true,
		"has2":   true,
		"has3":   true,
	}
	expected.ShouldBeEqual(t, 0, "GetMap returns unique map -- with duplicates", actual)
}

// ═══════════════════════════════════════════
// Get — additional branch coverage
// ═══════════════════════════════════════════

func Test_Cov_Get_NilSlice(t *testing.T) {
	result := intunique.Get(nil)
	actual := args.Map{"isNil": result == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Get returns nil -- nil input", actual)
}

func Test_Cov_Get_SingleElement(t *testing.T) {
	input := []int{42}
	result := intunique.Get(&input)
	actual := args.Map{"len": len(*result), "first": (*result)[0]}
	expected := args.Map{"len": 1, "first": 42}
	expected.ShouldBeEqual(t, 0, "Get returns same -- single element", actual)
}
