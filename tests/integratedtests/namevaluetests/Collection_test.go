package namevaluetests

import (
	"fmt"
	"strings"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/namevalue"
)

// region StringStringCollection tests

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
		testCase.ShouldBeEqual(t, caseIndex, length, isEmpty, hasAny)
	}
}

// endregion

// region StringIntCollection tests

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
		// For count > 1, join contains comma separator
		containsComma := fmt.Sprintf("%v", strings.Contains(joined, ","))

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, length, containsItem0, containsComma)
	}
}

// endregion

// region Collection Prepend/Append tests

func Test_CollectionPrependAppend_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionPrependAppendTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		op, _ := input.GetAsString("op")

		col := namevalue.NewGenericCollectionDefault[string, string]()
		col.Add(namevalue.StringString{Name: "original-0", Value: "v0"})
		col.Add(namevalue.StringString{Name: "original-1", Value: "v1"})

		var length string
		var firstItemName string

		// Act
		switch op {
		case "prepend":
			col.Prepend(namevalue.StringString{Name: "prepended", Value: "vp"})
			length = fmt.Sprintf("%d", col.Length())
			firstItemName = col.Items[0].Name
		case "append":
			col.Append(namevalue.StringString{Name: "appended", Value: "va"})
			length = fmt.Sprintf("%d", col.Length())
			firstItemName = col.Items[col.LastIndex()].Name
		case "prependif-false":
			col.PrependIf(false, namevalue.StringString{Name: "skipped", Value: "vs"})
			length = fmt.Sprintf("%d", col.Length())
			firstItemName = col.Items[0].Name
		case "appendif-false":
			col.AppendIf(false, namevalue.StringString{Name: "skipped", Value: "vs"})
			length = fmt.Sprintf("%d", col.Length())
			firstItemName = col.Items[0].Name
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, length, firstItemName)
	}
}

// endregion

// region Collection Clone tests

func Test_CollectionClone_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count, _ := input.Get("count")
		countInt := count.(int)

		if countInt < 0 {
			// nil clone test
			var nilCol *namevalue.StringStringCollection
			result := nilCol.ClonePtr()
			isNil := fmt.Sprintf("%v", result == nil)

			// Assert
			testCase.ShouldBeEqual(t, caseIndex, isNil)
			continue
		}

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

		// Modify original, verify clone is independent
		col.Add(namevalue.StringString{Name: "extra", Value: "x"})
		cloneUnchanged := fmt.Sprintf("%v", cloned.Length() == countInt)
		isEqual := fmt.Sprintf("%v", cloned.IsEqualByString(
			namevalue.NewGenericCollectionUsing[string, string](true, cloned.Items...),
		))

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, sameLength, cloneUnchanged, isEqual)
	}
}

// endregion

// region Collection IsEqualByString tests

func Test_CollectionIsEqual_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionIsEqualTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		caseType, _ := input.GetAsString("case")

		var result string

		// Act
		switch caseType {
		case "equal":
			a := namevalue.NewGenericCollectionDefault[string, int]()
			a.Add(namevalue.StringInt{Name: "x", Value: 1})
			b := namevalue.NewGenericCollectionDefault[string, int]()
			b.Add(namevalue.StringInt{Name: "x", Value: 1})
			result = fmt.Sprintf("%v", a.IsEqualByString(b))
		case "notequal":
			a := namevalue.NewGenericCollectionDefault[string, int]()
			a.Add(namevalue.StringInt{Name: "x", Value: 1})
			b := namevalue.NewGenericCollectionDefault[string, int]()
			b.Add(namevalue.StringInt{Name: "x", Value: 99})
			result = fmt.Sprintf("%v", a.IsEqualByString(b))
		case "difflength":
			a := namevalue.NewGenericCollectionDefault[string, int]()
			a.Add(namevalue.StringInt{Name: "x", Value: 1})
			b := namevalue.NewGenericCollectionDefault[string, int]()
			b.Add(namevalue.StringInt{Name: "x", Value: 1})
			b.Add(namevalue.StringInt{Name: "y", Value: 2})
			result = fmt.Sprintf("%v", a.IsEqualByString(b))
		case "bothnils":
			var a *namevalue.StringIntCollection
			var b *namevalue.StringIntCollection
			result = fmt.Sprintf("%v", a.IsEqualByString(b))
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

// endregion

// region Collection Error tests

func Test_CollectionError_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionErrorTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count, _ := input.Get("count")
		countInt := count.(int)

		// Act
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

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, hasError, hasMsgError)
	}
}

// endregion

// region Collection Dispose tests

func Test_CollectionDispose_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionDisposeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count, _ := input.Get("count")
		countInt := count.(int)

		// Act
		col := namevalue.NewGenericCollectionDefault[string, string]()
		for i := 0; i < countInt; i++ {
			col.Add(namevalue.StringString{
				Name:  fmt.Sprintf("d%d", i),
				Value: fmt.Sprintf("v%d", i),
			})
		}
		col.Dispose()
		isNilItems := fmt.Sprintf("%v", col.Items == nil)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, isNilItems)
	}
}

// endregion

// region Collection ConcatNew tests

func Test_CollectionConcatNew_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionConcatNewTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		original, _ := input.Get("original")
		extra, _ := input.Get("extra")
		originalInt := original.(int)
		extraInt := extra.(int)

		// Act
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

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, newLength, originalUnchanged)
	}
}

// endregion

// region StringMapAnyCollection tests

func Test_StringMapAnyCollection_Verification(t *testing.T) {
	for caseIndex, testCase := range stringMapAnyCollectionTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count, _ := input.Get("count")
		countInt := count.(int)

		// Act
		col := namevalue.NewGenericCollectionDefault[string, map[string]any]()
		for i := 0; i < countInt; i++ {
			var mapVal map[string]any
			if i == 0 && countInt == 1 {
				// nil map case
				mapVal = nil
			} else {
				mapVal = map[string]any{"key": i}
			}
			col.Add(namevalue.StringMapAny{
				Name:  fmt.Sprintf("map%d", i),
				Value: mapVal,
			})
		}
		length := fmt.Sprintf("%d", col.Length())
		hasItems := fmt.Sprintf("%v", col.HasAnyItem())

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, length, hasItems)
	}
}

// endregion
