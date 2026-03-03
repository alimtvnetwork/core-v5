package coreuniquetests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coreunique/intunique"
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
	tc.ShouldBeEqual(t, 0, actLines...)
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
	tc.ShouldBeEqual(t, 0, actLines...)
}

func Test_IntUnique_Get_Nil(t *testing.T) {
	tc := intUniqueGetNilTestCase

	// Act
	result := intunique.Get(nil)
	actLines := []string{
		fmt.Sprintf("%v", result == nil),
	}

	// Assert
	tc.ShouldBeEqual(t, 0, actLines...)
}
