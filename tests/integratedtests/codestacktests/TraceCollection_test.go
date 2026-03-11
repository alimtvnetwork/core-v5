package codestacktests

import (
	"testing"

	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/coretests/args"
)

func newTestTrace(pkgName string, line int) codestack.Trace {
	return codestack.Trace{
		PackageName:       pkgName,
		MethodName:        "TestMethod",
		PackageMethodName: pkgName + ".TestMethod",
		FilePath:          "/src/" + pkgName + "/file.go",
		Line:              line,
		IsOkay:            true,
	}
}

func newTestCollection() *codestack.TraceCollection {
	t1 := newTestTrace("pkg1", 10)
	t2 := newTestTrace("pkg2", 20)
	t3 := newTestTrace("pkg3", 30)

	collection := codestack.New.TraceCollection.Default()
	collection.Add(t1).Add(t2).Add(t3)

	return collection
}

func Test_TraceCollection_Basic(t *testing.T) {
	for caseIndex, testCase := range traceCollectionBasicTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count, _ := input.GetAsInt("count")

		// Act
		collection := codestack.New.TraceCollection.Cap(count)
		for i := 0; i < count; i++ {
			collection.Add(newTestTrace("pkg", i))
		}

		actual := args.Map{
			"length":     collection.Length(),
			"isEmpty":    collection.IsEmpty(),
			"hasAnyItem": collection.HasAnyItem(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_FirstLast(t *testing.T) {
	for caseIndex, testCase := range traceCollectionFirstLastTestCases {
		// Arrange
		collection := newTestCollection()

		// Act
		actual := args.Map{
			"firstPkg": collection.First().PackageName,
			"lastPkg":  collection.Last().PackageName,
			"lastIdx":  collection.LastIndex(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_SkipTake(t *testing.T) {
	for caseIndex, testCase := range traceCollectionSkipTakeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		expected := testCase.ExpectedInput.(args.Map)
		collection := newTestCollection()

		skipVal, hasSkip := input.GetAsInt("skip")
		takeVal, hasTake := input.GetAsInt("take")

		// Act
		if hasSkip == nil {
			skipped := collection.SkipCollection(skipVal)
			actual := args.Map{
				"length":   skipped.Length(),
				"firstPkg": skipped.First().PackageName,
			}
			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		} else if hasTake == nil {
			taken := collection.TakeCollection(takeVal)
			expectedLastPkg, _ := expected.GetAsString("lastPkg")
			actual := args.Map{
				"length":  taken.Length(),
				"lastPkg": expectedLastPkg,
			}
			_ = taken.Last().PackageName
			actual["lastPkg"] = taken.Last().PackageName
			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		}
	}
}

func Test_TraceCollection_Reverse(t *testing.T) {
	for caseIndex, testCase := range traceCollectionReverseTestCases {
		// Arrange
		collection := newTestCollection()

		// Act
		collection.Reverse()

		actual := args.Map{
			"firstPkg": collection.First().PackageName,
			"lastPkg":  collection.Last().PackageName,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_FilterPackageName(t *testing.T) {
	for caseIndex, testCase := range traceCollectionFilterTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pkgName, _ := input.GetAsString("package")
		collection := newTestCollection()

		// Act
		filtered := collection.FilterPackageNameTraceCollection(pkgName)

		actual := args.Map{
			"length":   filtered.Length(),
			"firstPkg": filtered.First().PackageName,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
