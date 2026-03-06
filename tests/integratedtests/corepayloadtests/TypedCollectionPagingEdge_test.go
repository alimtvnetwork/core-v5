package corepayloadtests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/corepayload"
	"gitlab.com/auk-go/core/coretests/args"
)

// =============================================================================
// Tests: GetPagedCollection edge cases
// =============================================================================

func Test_TypedPayloadCollection_GetPagedCollection_EdgeCases(t *testing.T) {
	for caseIndex, testCase := range typedCollectionPagingEdgeCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count, _ := input.GetAsInt("count")
		pageSize, _ := input.GetAsInt("pageSize")

		collection := createNumberedUsers(count)

		// Act
		pages := collection.GetPagedCollection(pageSize)

		actual := args.Map{
			"pageCount": len(pages),
		}
		for i, page := range pages {
			actual[fmt.Sprintf("page%dItems", i+1)] = page.Length()
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// Tests: GetSinglePageCollection edge cases
// =============================================================================

func Test_TypedPayloadCollection_GetSinglePageCollection_EdgeCases(t *testing.T) {
	for caseIndex, testCase := range typedCollectionSinglePageEdgeCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count, _ := input.GetAsInt("count")
		pageSize, _ := input.GetAsInt("pageSize")
		pageIndex, _ := input.GetAsInt("pageIndex")

		collection := createNumberedUsers(count)

		// Act
		page := collection.GetSinglePageCollection(pageSize, pageIndex)

		actual := args.Map{
			"pageItemCount": page.Length(),
		}
		for i, item := range page.Items() {
			actual[fmt.Sprintf("item%d", i)] = item.Identifier()
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// Tests: GetPagedCollectionWithInfo edge cases
// =============================================================================

func Test_TypedPayloadCollection_GetPagedCollectionWithInfo_EdgeCases(t *testing.T) {
	for caseIndex, testCase := range typedCollectionPagedWithInfoEdgeCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count, _ := input.GetAsInt("count")
		pageSize, _ := input.GetAsInt("pageSize")

		collection := createNumberedUsers(count)

		// Act
		pages := collection.GetPagedCollectionWithInfo(pageSize)

		actual := args.Map{
			"pageCount": len(pages),
		}

		for i := 0; i < len(pages) && i < 2; i++ {
			page := pages[i]
			prefix := fmt.Sprintf("p%d", i+1)
			actual[prefix+"CurrentPageIndex"] = page.Paging.CurrentPageIndex
			actual[prefix+"TotalPages"] = page.Paging.TotalPages
			actual[prefix+"PerPageItems"] = page.Paging.PerPageItems
			actual[prefix+"TotalItems"] = page.Paging.TotalItems
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// Tests: GetPagesSize edge cases
// =============================================================================

func Test_TypedPayloadCollection_GetPagesSize_EdgeCases(t *testing.T) {
	for caseIndex, testCase := range typedCollectionPagesSizeEdgeCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count, _ := input.GetAsInt("count")
		pageSize, _ := input.GetAsInt("pageSize")

		collection := createNumberedUsers(count)

		// Act
		pagesSize := collection.GetPagesSize(pageSize)

		// Assert
		actual := args.Map{
			"pagesSize": pagesSize,
		}
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// Tests: Paging empty with GetPagedCollectionWithInfo
// =============================================================================

func Test_TypedPayloadCollection_PagingWithInfo_Empty(t *testing.T) {
	// Arrange
	collection := corepayload.EmptyTypedPayloadCollection[testUser]()

	// Act
	pages := collection.GetPagedCollectionWithInfo(5)

	// Assert
	if len(pages) != 1 {
		t.Errorf("Expected 1 page for empty collection, got %d", len(pages))
	}

	if pages[0].Collection.Length() != 0 {
		t.Errorf("Expected 0 items on empty page, got %d", pages[0].Collection.Length())
	}
}
