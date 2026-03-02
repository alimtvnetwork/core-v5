package corepayloadtests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coredata/corepayload"
)

// =============================================================================
// Tests: GetPagedCollection edge cases
// =============================================================================

func Test_TypedPayloadCollection_GetPagedCollection_EdgeCases(t *testing.T) {
	for caseIndex, testCase := range typedCollectionPagingEdgeCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count := input.GetDirectLower("count").(int)
		pageSize := input.GetDirectLower("pageSize").(int)

		collection := createNumberedUsers(count)

		// Act
		pages := collection.GetPagedCollection(pageSize)

		results := []string{fmt.Sprintf("%d", len(pages))}
		for _, page := range pages {
			results = append(results, fmt.Sprintf("%d", page.Length()))
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, results...)
	}
}

// =============================================================================
// Tests: GetSinglePageCollection edge cases
// =============================================================================

func Test_TypedPayloadCollection_GetSinglePageCollection_EdgeCases(t *testing.T) {
	for caseIndex, testCase := range typedCollectionSinglePageEdgeCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count := input.GetDirectLower("count").(int)
		pageSize := input.GetDirectLower("pageSize").(int)
		pageIndex := input.GetDirectLower("pageIndex").(int)

		collection := createNumberedUsers(count)

		// Act
		page := collection.GetSinglePageCollection(pageSize, pageIndex)

		results := []string{fmt.Sprintf("%d", page.Length())}
		for _, item := range page.Items() {
			results = append(results, item.Identifier())
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, results...)
	}
}

// =============================================================================
// Tests: GetPagedCollectionWithInfo edge cases
// =============================================================================

func Test_TypedPayloadCollection_GetPagedCollectionWithInfo_EdgeCases(t *testing.T) {
	for caseIndex, testCase := range typedCollectionPagedWithInfoEdgeCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count := input.GetDirectLower("count").(int)
		pageSize := input.GetDirectLower("pageSize").(int)

		collection := createNumberedUsers(count)

		// Act
		pages := collection.GetPagedCollectionWithInfo(pageSize)

		results := []string{fmt.Sprintf("%d", len(pages))}

		for i := 0; i < len(pages) && i < 2; i++ {
			page := pages[i]
			results = append(results,
				fmt.Sprintf("%d", page.Paging.CurrentPageIndex),
				fmt.Sprintf("%d", page.Paging.TotalPages),
				fmt.Sprintf("%d", page.Paging.PerPageItems),
				fmt.Sprintf("%d", page.Paging.TotalItems),
			)
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, results...)
	}
}

// =============================================================================
// Tests: GetPagesSize edge cases
// =============================================================================

func Test_TypedPayloadCollection_GetPagesSize_EdgeCases(t *testing.T) {
	for caseIndex, testCase := range typedCollectionPagesSizeEdgeCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count := input.GetDirectLower("count").(int)
		pageSize := input.GetDirectLower("pageSize").(int)

		collection := createNumberedUsers(count)

		// Act
		pagesSize := collection.GetPagesSize(pageSize)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			fmt.Sprintf("%d", pagesSize),
		)
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
