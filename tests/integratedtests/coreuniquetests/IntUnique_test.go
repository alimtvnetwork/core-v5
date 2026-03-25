package coreuniquetests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coreunique/intunique"
)

func Test_IntUnique_Get_RemovesDuplicates(t *testing.T) {
	tc := intUniqueGetRemovesDuplicatesTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	inputVal, _ := input.Get("input")
	slice := inputVal.([]int)
	clone := make([]int, len(slice))
	copy(clone, slice)

	// Act
	result := intunique.Get(&clone)
	actLines := []string{
		fmt.Sprintf("%v", len(*result)),
	}

	// Assert

	tc.ShouldBeEqualFirst(t, actLines...)
}

func Test_IntUnique_Get_AlreadyUnique(t *testing.T) {
	tc := intUniqueGetAlreadyUniqueTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	inputVal, _ := input.Get("input")
	slice := inputVal.([]int)
	clone := make([]int, len(slice))
	copy(clone, slice)

	// Act
	result := intunique.Get(&clone)
	actLines := []string{
		fmt.Sprintf("%v", len(*result)),
	}

	// Assert

	tc.ShouldBeEqualFirst(t, actLines...)
}

func Test_IntUnique_Get_Nil(t *testing.T) {
	tc := intUniqueGetNilTestCase

	// Act
	result := intunique.Get(nil)
	actLines := []string{
		fmt.Sprintf("%v", result == nil),
	}

	// Assert

	tc.ShouldBeEqualFirst(t, actLines...)
}
