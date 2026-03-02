package coreinstructiontests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coreinstruction"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/errcore"
)

func Test_BaseIdentifier_Verification(t *testing.T) {
	for caseIndex, testCase := range identifierTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		id, _ := input.GetAsString("id")

		// Act
		identifier := coreinstruction.NewIdentifier(id)
		idStr := identifier.IdString()
		isEmpty := fmt.Sprintf("%v", identifier.IsIdEmpty())
		isWhitespace := fmt.Sprintf("%v", identifier.IsIdWhitespace())

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			idStr,
			isEmpty,
			isWhitespace,
		)
	}
}

func Test_Identifiers_Length_Verification(t *testing.T) {
	for caseIndex, testCase := range identifiersLengthTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		ids, err := input.GetAsStrings("ids")
		errcore.HandleErrMessage("ids required", err)

		// Act
		identifiers := coreinstruction.NewIdentifiers(ids...)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			fmt.Sprintf("%v", identifiers.Length()),
			fmt.Sprintf("%v", identifiers.IsEmpty()),
			fmt.Sprintf("%v", identifiers.HasAnyItem()),
		)
	}
}

func Test_Identifiers_GetById_Verification(t *testing.T) {
	for caseIndex, testCase := range identifiersGetByIdTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		ids, err := input.GetAsStrings("ids")
		errcore.HandleErrMessage("ids required", err)
		searchId, _ := input.GetAsString("searchId")

		// Act
		identifiers := coreinstruction.NewIdentifiers(ids...)
		result := identifiers.GetById(searchId)

		// Assert
		found := result != nil
		foundId := ""
		if found {
			foundId = result.Id
		}

		testCase.ShouldBeEqual(
			t,
			caseIndex,
			fmt.Sprintf("%v", found),
			foundId,
		)
	}
}

func Test_Identifiers_IndexOf_Verification(t *testing.T) {
	for caseIndex, testCase := range identifiersIndexOfTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		ids, err := input.GetAsStrings("ids")
		errcore.HandleErrMessage("ids required", err)
		searchId, _ := input.GetAsString("searchId")

		// Act
		identifiers := coreinstruction.NewIdentifiers(ids...)
		index := identifiers.IndexOf(searchId)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			fmt.Sprintf("%v", index),
		)
	}
}

func Test_Identifiers_Clone_Verification(t *testing.T) {
	for caseIndex, testCase := range identifiersCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		ids, err := input.GetAsStrings("ids")
		errcore.HandleErrMessage("ids required", err)

		// Act
		original := coreinstruction.NewIdentifiers(ids...)
		cloned := original.Clone()

		// Assert
		results := []string{fmt.Sprintf("%v", cloned.Length())}
		for _, baseId := range cloned.Ids {
			results = append(results, baseId.Id)
		}

		testCase.ShouldBeEqual(
			t,
			caseIndex,
			results...,
		)
	}
}

func Test_Identifiers_Add_Verification(t *testing.T) {
	for caseIndex, testCase := range identifiersAddTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		ids, err := input.GetAsStrings("ids")
		errcore.HandleErrMessage("ids required", err)
		addId, _ := input.GetAsString("addId")

		// Act
		identifiers := coreinstruction.NewIdentifiers(ids...)
		identifiers.Add(addId)

		// Assert
		results := []string{fmt.Sprintf("%v", identifiers.Length())}
		for _, baseId := range identifiers.Ids {
			results = append(results, baseId.Id)
		}

		testCase.ShouldBeEqual(
			t,
			caseIndex,
			results...,
		)
	}
}

func Test_Specification_Clone_Verification(t *testing.T) {
	for caseIndex, testCase := range specificationCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		id, _ := input.GetAsString("id")
		display, _ := input.GetAsString("display")
		typeName, _ := input.GetAsString("typeName")
		tags, err := input.GetAsStrings("tags")
		errcore.HandleErrMessage("tags required", err)
		isGlobal, _ := input.GetAsBool("isGlobal")

		// Act
		spec := coreinstruction.NewSpecification(id, display, typeName, tags, isGlobal)
		cloned := spec.Clone()

		// Assert
		results := []string{
			cloned.Id,
			cloned.Display,
			cloned.Type,
			fmt.Sprintf("%v", cloned.TagsLength()),
		}

		for _, tag := range cloned.Tags {
			results = append(results, tag)
		}

		if cloned.TagsLength() == 0 {
			results = append(results, fmt.Sprintf("%v", cloned.IsGlobal))
		} else {
			results = append(results, fmt.Sprintf("%v", cloned.IsGlobal))
		}

		testCase.ShouldBeEqual(
			t,
			caseIndex,
			results...,
		)
	}
}

func Test_BaseTags_Verification(t *testing.T) {
	for caseIndex, testCase := range baseTagsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		tags, err := input.GetAsStrings("tags")
		errcore.HandleErrMessage("tags required", err)
		searchTags, err := input.GetAsStrings("searchTags")
		errcore.HandleErrMessage("searchTags required", err)

		// Act
		baseTags := coreinstruction.NewTags(tags)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			fmt.Sprintf("%v", baseTags.TagsLength()),
			fmt.Sprintf("%v", baseTags.IsTagsEmpty()),
			fmt.Sprintf("%v", baseTags.HasAllTags(searchTags...)),
			fmt.Sprintf("%v", baseTags.HasAnyTags(searchTags...)),
		)
	}
}

func Test_Specification_Clone_NilSafety(t *testing.T) {
	// Verify nil receiver returns nil without panic
	var spec *coreinstruction.Specification
	result := spec.Clone()

	if result != nil {
		t.Error("Expected nil from nil Specification.Clone()")
	}
}

func Test_Specification_Clone_DeepCopy_Tags(t *testing.T) {
	// Verify tags are deep-copied (mutating clone doesn't affect original)
	original := coreinstruction.NewSpecification("id", "display", "type", []string{"a", "b"}, false)
	cloned := original.Clone()

	// Mutate clone's tags
	cloned.Tags[0] = "MUTATED"

	if original.Tags[0] == "MUTATED" {
		t.Error("Clone shares Tags backing array with original — shallow copy detected")
	}
}
