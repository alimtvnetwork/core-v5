package coregenerictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coregeneric"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/errcore"
)

var funcsTestCases = []coretestcases.CaseV1{
	// MapCollection
	{Title: "MapCollection int to string", ArrangeInput: args.Map{"case": "mapcol"}, ExpectedInput: []string{"3", "v1", "v3"}},
	{Title: "MapCollection nil source", ArrangeInput: args.Map{"case": "mapcol-nil"}, ExpectedInput: []string{"true"}},
	{Title: "MapCollection empty source", ArrangeInput: args.Map{"case": "mapcol-empty"}, ExpectedInput: []string{"true"}},
	// FlatMapCollection
	{Title: "FlatMapCollection flattens", ArrangeInput: args.Map{"case": "flatmap"}, ExpectedInput: []string{"6"}},
	{Title: "FlatMapCollection nil", ArrangeInput: args.Map{"case": "flatmap-nil"}, ExpectedInput: []string{"true"}},
	// ReduceCollection
	{Title: "ReduceCollection sum", ArrangeInput: args.Map{"case": "reduce-sum"}, ExpectedInput: []string{"10"}},
	{Title: "ReduceCollection nil returns initial", ArrangeInput: args.Map{"case": "reduce-nil"}, ExpectedInput: []string{"99"}},
	{Title: "ReduceCollection string concat", ArrangeInput: args.Map{"case": "reduce-concat"}, ExpectedInput: []string{"abc"}},
	// GroupByCollection
	{Title: "GroupByCollection groups", ArrangeInput: args.Map{"case": "groupby"}, ExpectedInput: []string{"2", "3", "3"}},
	{Title: "GroupByCollection nil", ArrangeInput: args.Map{"case": "groupby-nil"}, ExpectedInput: []string{"0"}},
	// ContainsFunc/ContainsItem
	{Title: "ContainsFunc found", ArrangeInput: args.Map{"case": "containsfunc-found"}, ExpectedInput: []string{"true"}},
	{Title: "ContainsFunc not found", ArrangeInput: args.Map{"case": "containsfunc-notfound"}, ExpectedInput: []string{"false"}},
	{Title: "ContainsFunc nil", ArrangeInput: args.Map{"case": "containsfunc-nil"}, ExpectedInput: []string{"false"}},
	{Title: "ContainsItem found", ArrangeInput: args.Map{"case": "containsitem-found"}, ExpectedInput: []string{"true"}},
	{Title: "ContainsItem not found", ArrangeInput: args.Map{"case": "containsitem-notfound"}, ExpectedInput: []string{"false"}},
	{Title: "ContainsItem nil", ArrangeInput: args.Map{"case": "containsitem-nil"}, ExpectedInput: []string{"false"}},
	// IndexOfFunc/IndexOfItem
	{Title: "IndexOfFunc found", ArrangeInput: args.Map{"case": "indexoffunc-found"}, ExpectedInput: []string{"1"}},
	{Title: "IndexOfFunc not found", ArrangeInput: args.Map{"case": "indexoffunc-notfound"}, ExpectedInput: []string{"-1"}},
	{Title: "IndexOfFunc nil", ArrangeInput: args.Map{"case": "indexoffunc-nil"}, ExpectedInput: []string{"-1"}},
	{Title: "IndexOfItem found", ArrangeInput: args.Map{"case": "indexofitem-found"}, ExpectedInput: []string{"2"}},
	{Title: "IndexOfItem not found", ArrangeInput: args.Map{"case": "indexofitem-notfound"}, ExpectedInput: []string{"-1"}},
	// Distinct
	{Title: "Distinct removes duplicates", ArrangeInput: args.Map{"case": "distinct"}, ExpectedInput: []string{"3"}},
	{Title: "Distinct nil", ArrangeInput: args.Map{"case": "distinct-nil"}, ExpectedInput: []string{"true"}},
	{Title: "Distinct no duplicates", ArrangeInput: args.Map{"case": "distinct-unique"}, ExpectedInput: []string{"3"}},
	// MapSimpleSlice
	{Title: "MapSimpleSlice transforms", ArrangeInput: args.Map{"case": "mapsimple"}, ExpectedInput: []string{"3"}},
	{Title: "MapSimpleSlice nil", ArrangeInput: args.Map{"case": "mapsimple-nil"}, ExpectedInput: []string{"true"}},
}

func Test_GenericFuncs_Verification(t *testing.T) {
	for caseIndex, tc := range funcsTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		caseType := input["case"].(string)

		var actLines []string

		// Act
		switch caseType {
		case "mapcol":
			src := coregeneric.New.Collection.Int.Items(1, 2, 3)
			result := coregeneric.MapCollection(src, func(i int) string { return fmt.Sprintf("v%d", i) })
			actLines = []string{fmt.Sprintf("%v", result.Length()), result.First(), result.Last()}
		case "mapcol-nil":
			result := coregeneric.MapCollection[int, string](nil, func(i int) string { return "" })
			actLines = []string{fmt.Sprintf("%v", result.IsEmpty())}
		case "mapcol-empty":
			src := coregeneric.EmptyCollection[int]()
			result := coregeneric.MapCollection(src, func(i int) string { return "" })
			actLines = []string{fmt.Sprintf("%v", result.IsEmpty())}
		case "flatmap":
			src := coregeneric.New.Collection.Int.Items(1, 2, 3)
			result := coregeneric.FlatMapCollection(src, func(i int) []int { return []int{i, i * 10} })
			actLines = []string{fmt.Sprintf("%v", result.Length())}
		case "flatmap-nil":
			result := coregeneric.FlatMapCollection[int, int](nil, func(i int) []int { return nil })
			actLines = []string{fmt.Sprintf("%v", result.IsEmpty())}
		case "reduce-sum":
			src := coregeneric.New.Collection.Int.Items(1, 2, 3, 4)
			sum := coregeneric.ReduceCollection(src, 0, func(a, b int) int { return a + b })
			actLines = []string{fmt.Sprintf("%v", sum)}
		case "reduce-nil":
			result := coregeneric.ReduceCollection[int, int](nil, 99, func(a, b int) int { return a + b })
			actLines = []string{fmt.Sprintf("%v", result)}
		case "reduce-concat":
			src := coregeneric.New.Collection.String.Items("a", "b", "c")
			result := coregeneric.ReduceCollection(src, "", func(a, b string) string { return a + b })
			actLines = []string{result}
		case "groupby":
			src := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5, 6)
			groups := coregeneric.GroupByCollection(src, func(i int) string {
				if i%2 == 0 {
					return "even"
				}
				return "odd"
			})
			actLines = []string{fmt.Sprintf("%v", len(groups)), fmt.Sprintf("%v", groups["even"].Length()), fmt.Sprintf("%v", groups["odd"].Length())}
		case "groupby-nil":
			groups := coregeneric.GroupByCollection[int, string](nil, func(i int) string { return "" })
			actLines = []string{fmt.Sprintf("%v", len(groups))}
		case "containsfunc-found":
			src := coregeneric.New.Collection.Int.Items(1, 2, 3)
			actLines = []string{fmt.Sprintf("%v", coregeneric.ContainsFunc(src, func(i int) bool { return i == 2 }))}
		case "containsfunc-notfound":
			src := coregeneric.New.Collection.Int.Items(1, 2, 3)
			actLines = []string{fmt.Sprintf("%v", coregeneric.ContainsFunc(src, func(i int) bool { return i == 99 }))}
		case "containsfunc-nil":
			actLines = []string{fmt.Sprintf("%v", coregeneric.ContainsFunc[int](nil, func(i int) bool { return true }))}
		case "containsitem-found":
			src := coregeneric.New.Collection.String.Items("a", "b", "c")
			actLines = []string{fmt.Sprintf("%v", coregeneric.ContainsItem(src, "b"))}
		case "containsitem-notfound":
			src := coregeneric.New.Collection.String.Items("a", "b")
			actLines = []string{fmt.Sprintf("%v", coregeneric.ContainsItem(src, "z"))}
		case "containsitem-nil":
			actLines = []string{fmt.Sprintf("%v", coregeneric.ContainsItem[string](nil, "x"))}
		case "indexoffunc-found":
			src := coregeneric.New.Collection.Int.Items(10, 20, 30)
			actLines = []string{fmt.Sprintf("%v", coregeneric.IndexOfFunc(src, func(i int) bool { return i == 20 }))}
		case "indexoffunc-notfound":
			src := coregeneric.New.Collection.Int.Items(1, 2, 3)
			actLines = []string{fmt.Sprintf("%v", coregeneric.IndexOfFunc(src, func(i int) bool { return i == 99 }))}
		case "indexoffunc-nil":
			actLines = []string{fmt.Sprintf("%v", coregeneric.IndexOfFunc[int](nil, func(i int) bool { return true }))}
		case "indexofitem-found":
			src := coregeneric.New.Collection.String.Items("x", "y", "z")
			actLines = []string{fmt.Sprintf("%v", coregeneric.IndexOfItem(src, "z"))}
		case "indexofitem-notfound":
			src := coregeneric.New.Collection.String.Items("a")
			actLines = []string{fmt.Sprintf("%v", coregeneric.IndexOfItem(src, "q"))}
		case "distinct":
			src := coregeneric.New.Collection.Int.Items(1, 2, 2, 3, 1, 3)
			actLines = []string{fmt.Sprintf("%v", coregeneric.Distinct(src).Length())}
		case "distinct-nil":
			actLines = []string{fmt.Sprintf("%v", coregeneric.Distinct[int](nil).IsEmpty())}
		case "distinct-unique":
			src := coregeneric.New.Collection.String.Items("a", "b", "c")
			actLines = []string{fmt.Sprintf("%v", coregeneric.Distinct(src).Length())}
		case "mapsimple":
			src := coregeneric.SimpleSliceFrom([]int{1, 2, 3})
			result := coregeneric.MapSimpleSlice(src, func(i int) string { return fmt.Sprintf("%d", i) })
			actLines = []string{fmt.Sprintf("%v", result.Length())}
		case "mapsimple-nil":
			result := coregeneric.MapSimpleSlice[int, string](nil, func(i int) string { return "" })
			actLines = []string{fmt.Sprintf("%v", result.IsEmpty())}
		}

		expectedLines := tc.ExpectedInput.([]string)

		// Assert
		errcore.PrintDiffOnMismatch(caseIndex, tc.Title, actLines, expectedLines)
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
