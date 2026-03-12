package msgcreatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/msgcreator"
)

// Cover Assert.SortedArray with isPrint=true
func Test_Cov2_Assert_SortedArray_Print(t *testing.T) {
	result := msgcreator.Assert.SortedArray(true, true, "c b a")
	actual := args.Map{"first": result[0], "len": len(result)}
	expected := args.Map{"first": "a", "len": 3}
	expected.ShouldBeEqual(t, 0, "SortedArray isPrint=true", actual)
}

func Test_Cov2_Assert_SortedArray_NoSort(t *testing.T) {
	result := msgcreator.Assert.SortedArray(false, false, "c b a")
	actual := args.Map{"first": result[0], "len": len(result)}
	expected := args.Map{"first": "c", "len": 3}
	expected.ShouldBeEqual(t, 0, "SortedArray isSort=false", actual)
}

func Test_Cov2_Assert_SortedMessage_Print(t *testing.T) {
	result := msgcreator.Assert.SortedMessage(true, "c b a", ",")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SortedMessage isPrint=true", actual)
}

// Cover ToStrings with various types
func Test_Cov2_Assert_ToStrings_Slice(t *testing.T) {
	result := msgcreator.Assert.ToStrings([]string{"a", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ToStrings slice", actual)
}

func Test_Cov2_Assert_ToStrings_Int(t *testing.T) {
	result := msgcreator.Assert.ToStrings(42)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ToStrings int", actual)
}

func Test_Cov2_Assert_ToStrings_Bool(t *testing.T) {
	result := msgcreator.Assert.ToStrings(true)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ToStrings bool", actual)
}

func Test_Cov2_Assert_ToStrings_MapStringAny(t *testing.T) {
	result := msgcreator.Assert.ToStrings(map[string]any{"k": "v"})
	actual := args.Map{"hasItems": len(result) > 0}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "ToStrings map[string]any", actual)
}

func Test_Cov2_Assert_ToStringsWithSpace_Empty(t *testing.T) {
	result := msgcreator.Assert.ToStringsWithSpace(4, []string{})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ToStringsWithSpace empty", actual)
}
