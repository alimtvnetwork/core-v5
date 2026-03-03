package namevaluetests

import (
	"fmt"
	"strings"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/namevalue"
)

// ==========================================================================
// Test: StringStringCollection
// ==========================================================================

func Test_StringStringCollection_Verification(t *testing.T) {
	for caseIndex, testCase := range stringStringCollectionTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count, _ := input.Get("count")
		countInt := count.(int)

		// Act
		col := namevalue.NewGenericCollectionDefault[string, string]()
		for i := 0; i < countInt; i++ {
			col.Add(namevalue.StringString{
				Name:  fmt.Sprintf("key%d", i),
				Value: fmt.Sprintf("val%d", i),
			})
		}
		length := fmt.Sprintf("%d", col.Length())
		isEmpty := fmt.Sprintf("%v", col.IsEmpty())
		hasAny := fmt.Sprintf("%v", col.HasAnyItem())

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			length,
			isEmpty,
			hasAny,
		)
	}
}

// ==========================================================================
// Test: StringIntCollection
// ==========================================================================

func Test_StringIntCollection_Verification(t *testing.T) {
	for caseIndex, testCase := range stringIntCollectionTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count, _ := input.Get("count")
		countInt := count.(int)

		// Act
		col := namevalue.NewGenericCollectionDefault[string, int]()
		for i := 0; i < countInt; i++ {
			col.Add(namevalue.StringInt{
				Name:  fmt.Sprintf("item%d", i),
				Value: i * 10,
			})
		}
		length := fmt.Sprintf("%d", col.Length())
		joined := col.Join(", ")
		containsItem0 := fmt.Sprintf("%v", strings.Contains(joined, "item0"))
		containsComma := fmt.Sprintf("%v", strings.Contains(joined, ","))

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			length,
			containsItem0,
			containsComma,
		)
	}
}

// ==========================================================================
// Test: Prepend
// ==========================================================================

func Test_Collection_Prepend(t *testing.T) {
	tc := collectionPrependTestCase

	col := namevalue.NewGenericCollectionDefault[string, string]()
	col.Add(namevalue.StringString{Name: "original-0", Value: "v0"})
	col.Add(namevalue.StringString{Name: "original-1", Value: "v1"})

	// Act
	col.Prepend(namevalue.StringString{Name: "prepended", Value: "vp"})

	// Assert
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", col.Length()),
		col.Items[0].Name,
	)
}

// ==========================================================================
// Test: Append
// ==========================================================================

func Test_Collection_Append(t *testing.T) {
	tc := collectionAppendTestCase

	col := namevalue.NewGenericCollectionDefault[string, string]()
	col.Add(namevalue.StringString{Name: "original-0", Value: "v0"})
	col.Add(namevalue.StringString{Name: "original-1", Value: "v1"})

	// Act
	col.Append(namevalue.StringString{Name: "appended", Value: "va"})

	// Assert
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", col.Length()),
		col.Items[col.LastIndex()].Name,
	)
}

// ==========================================================================
// Test: PrependIf false
// ==========================================================================

func Test_Collection_PrependIfFalse(t *testing.T) {
	tc := collectionPrependIfFalseTestCase

	col := namevalue.NewGenericCollectionDefault[string, string]()
	col.Add(namevalue.StringString{Name: "original-0", Value: "v0"})
	col.Add(namevalue.StringString{Name: "original-1", Value: "v1"})

	// Act
	col.PrependIf(false, namevalue.StringString{Name: "skipped", Value: "vs"})

	// Assert
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", col.Length()),
		col.Items[0].Name,
	)
}

// ==========================================================================
// Test: AppendIf false
// ==========================================================================

func Test_Collection_AppendIfFalse(t *testing.T) {
	tc := collectionAppendIfFalseTestCase

	col := namevalue.NewGenericCollectionDefault[string, string]()
	col.Add(namevalue.StringString{Name: "original-0", Value: "v0"})
	col.Add(namevalue.StringString{Name: "original-1", Value: "v1"})

	// Act
	col.AppendIf(false, namevalue.StringString{Name: "skipped", Value: "vs"})

	// Assert
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", col.Length()),
		col.Items[0].Name,
	)
}

// ==========================================================================
// Test: Clone — valid collection
// ==========================================================================

func Test_CollectionClone_Valid(t *testing.T) {
	tc := collectionCloneValidTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	count, _ := input.Get("count")
	countInt := count.(int)

	// Act
	col := namevalue.NewGenericCollectionDefault[string, string]()

	for i := 0; i < countInt; i++ {
		col.Add(namevalue.StringString{
			Name:  fmt.Sprintf("k%d", i),
			Value: fmt.Sprintf("v%d", i),
		})
	}

	cloned := col.Clone()
	sameLength := fmt.Sprintf("%d", cloned.Length())

	col.Add(namevalue.StringString{Name: "extra", Value: "x"})
	cloneUnchanged := fmt.Sprintf("%v", cloned.Length() == countInt)
	isEqual := fmt.Sprintf("%v", cloned.IsEqualByString(
		namevalue.NewGenericCollectionUsing[string, string](true, cloned.Items...),
	))

	// Assert
	tc.ShouldBeEqual(t, 0,
		sameLength,
		cloneUnchanged,
		isEqual,
	)
}

// ==========================================================================
// Test: Clone — nil receiver
// ==========================================================================

func Test_CollectionClone_NilReceiver(t *testing.T) {
	tc := collectionCloneNilTestCase

	// Act
	var nilCol *namevalue.StringStringCollection
	result := nilCol.ClonePtr()
	isNil := fmt.Sprintf("%v", result == nil)

	// Assert
	tc.ShouldBeEqual(t, 0, isNil)
}

// ==========================================================================
// Test: IsEqualByString — Equal
// ==========================================================================

func Test_CollectionIsEqual_Equal(t *testing.T) {
	tc := collectionIsEqualEqualTestCase

	a := namevalue.NewGenericCollectionDefault[string, int]()
	a.Add(namevalue.StringInt{Name: "x", Value: 1})
	b := namevalue.NewGenericCollectionDefault[string, int]()
	b.Add(namevalue.StringInt{Name: "x", Value: 1})

	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", a.IsEqualByString(b)))
}

// ==========================================================================
// Test: IsEqualByString — NotEqual
// ==========================================================================

func Test_CollectionIsEqual_NotEqual(t *testing.T) {
	tc := collectionIsEqualNotEqualTestCase

	a := namevalue.NewGenericCollectionDefault[string, int]()
	a.Add(namevalue.StringInt{Name: "x", Value: 1})
	b := namevalue.NewGenericCollectionDefault[string, int]()
	b.Add(namevalue.StringInt{Name: "x", Value: 99})

	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", a.IsEqualByString(b)))
}

// ==========================================================================
// Test: IsEqualByString — DiffLength
// ==========================================================================

func Test_CollectionIsEqual_DiffLength(t *testing.T) {
	tc := collectionIsEqualDiffLengthTestCase

	a := namevalue.NewGenericCollectionDefault[string, int]()
	a.Add(namevalue.StringInt{Name: "x", Value: 1})
	b := namevalue.NewGenericCollectionDefault[string, int]()
	b.Add(namevalue.StringInt{Name: "x", Value: 1})
	b.Add(namevalue.StringInt{Name: "y", Value: 2})

	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", a.IsEqualByString(b)))
}

// ==========================================================================
// Test: IsEqualByString — BothNils
// ==========================================================================

func Test_CollectionIsEqual_BothNils(t *testing.T) {
	tc := collectionIsEqualBothNilsTestCase

	var a *namevalue.StringIntCollection
	var b *namevalue.StringIntCollection

	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", a.IsEqualByString(b)))
}

// ==========================================================================
// Test: Error
// ==========================================================================

func Test_CollectionError_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionErrorTestCases {
		input := testCase.ArrangeInput.(args.Map)
		count, _ := input.Get("count")
		countInt := count.(int)

		col := namevalue.NewGenericCollectionDefault[string, string]()
		for i := 0; i < countInt; i++ {
			col.Add(namevalue.StringString{
				Name:  fmt.Sprintf("err%d", i),
				Value: fmt.Sprintf("msg%d", i),
			})
		}

		err := col.Error()
		hasError := fmt.Sprintf("%v", err != nil)
		errMsg := col.ErrorUsingMessage("failed:")
		hasMsgError := fmt.Sprintf("%v", errMsg != nil)

		testCase.ShouldBeEqual(t, caseIndex,
			hasError,
			hasMsgError,
		)
	}
}

// ==========================================================================
// Test: Dispose
// ==========================================================================

func Test_CollectionDispose_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionDisposeTestCases {
		input := testCase.ArrangeInput.(args.Map)
		count, _ := input.Get("count")
		countInt := count.(int)

		col := namevalue.NewGenericCollectionDefault[string, string]()
		for i := 0; i < countInt; i++ {
			col.Add(namevalue.StringString{
				Name:  fmt.Sprintf("d%d", i),
				Value: fmt.Sprintf("v%d", i),
			})
		}

		col.Dispose()
		isNilItems := fmt.Sprintf("%v", col.Items == nil)

		testCase.ShouldBeEqual(t, caseIndex, isNilItems)
	}
}

// ==========================================================================
// Test: ConcatNew
// ==========================================================================

func Test_CollectionConcatNew_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionConcatNewTestCases {
		input := testCase.ArrangeInput.(args.Map)
		original, _ := input.Get("original")
		extra, _ := input.Get("extra")
		originalInt := original.(int)
		extraInt := extra.(int)

		col := namevalue.NewGenericCollectionDefault[string, int]()
		for i := 0; i < originalInt; i++ {
			col.Add(namevalue.StringInt{Name: fmt.Sprintf("o%d", i), Value: i})
		}

		extraItems := make([]namevalue.StringInt, extraInt)
		for i := 0; i < extraInt; i++ {
			extraItems[i] = namevalue.StringInt{Name: fmt.Sprintf("e%d", i), Value: i + 100}
		}

		newCol := col.ConcatNew(extraItems...)
		newLength := fmt.Sprintf("%d", newCol.Length())
		originalUnchanged := fmt.Sprintf("%d", col.Length())

		testCase.ShouldBeEqual(t, caseIndex,
			newLength,
			originalUnchanged,
		)
	}
}

// ==========================================================================
// Test: StringMapAnyCollection — with values
// ==========================================================================

func Test_StringMapAnyCollection_WithValues(t *testing.T) {
	tc := stringMapAnyCollectionWithValuesTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	mapValues := input["mapValues"].([]map[string]any)

	// Act
	col := namevalue.NewGenericCollectionDefault[string, map[string]any]()

	for i, mapVal := range mapValues {
		col.Add(namevalue.StringMapAny{
			Name:  fmt.Sprintf("map%d", i),
			Value: mapVal,
		})
	}

	length := fmt.Sprintf("%d", col.Length())
	hasItems := fmt.Sprintf("%v", col.HasAnyItem())

	// Assert
	tc.ShouldBeEqual(t, 0,
		length,
		hasItems,
	)
}

// ==========================================================================
// Test: StringMapAnyCollection — nil value
// ==========================================================================

func Test_StringMapAnyCollection_NilValue(t *testing.T) {
	tc := stringMapAnyCollectionNilValueTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	mapValues := input["mapValues"].([]map[string]any)

	// Act
	col := namevalue.NewGenericCollectionDefault[string, map[string]any]()

	for i, mapVal := range mapValues {
		col.Add(namevalue.StringMapAny{
			Name:  fmt.Sprintf("map%d", i),
			Value: mapVal,
		})
	}

	length := fmt.Sprintf("%d", col.Length())
	hasItems := fmt.Sprintf("%v", col.HasAnyItem())

	// Assert
	tc.ShouldBeEqual(t, 0,
		length,
		hasItems,
	)
}
