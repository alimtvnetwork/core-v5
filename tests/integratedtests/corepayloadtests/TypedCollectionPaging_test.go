package corepayloadtests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coredata/corepayload"
	"gitlab.com/auk-go/core/errcore"
)

// =============================================================================
// Helpers
// =============================================================================

func createNumberedUsers(count int) *corepayload.TypedPayloadCollection[testUser] {
	wrappers := make([]*corepayload.TypedPayloadWrapper[testUser], 0, count)

	for i := 0; i < count; i++ {
		user := testUser{
			Name:  fmt.Sprintf("User%d", i),
			Email: fmt.Sprintf("user%d@test.com", i),
			Age:   20 + i,
		}

		typed, err := corepayload.TypedPayloadWrapperNameIdRecord[testUser](
			user.Name,
			fmt.Sprintf("user-%d", i),
			user,
		)
		errcore.HandleErr(err)

		wrappers = append(wrappers, typed)
	}

	return corepayload.TypedPayloadCollectionFrom[testUser](wrappers)
}

// =============================================================================
// Tests: GetPagesSize
// =============================================================================

func Test_TypedPayloadCollection_GetPagesSize(t *testing.T) {
	for caseIndex, testCase := range typedCollectionPagesSizeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pageSize := input.GetDirectLower("pageSize").(int)

		count := 3 // default from createTestCollection (Alice, Bob, Carol)
		if c, ok := input.GetDirectLower("count").(int); ok {
			count = c
		}

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
// Tests: GetSinglePageCollection
// =============================================================================

func Test_TypedPayloadCollection_GetSinglePageCollection(t *testing.T) {
	for caseIndex, testCase := range typedCollectionSinglePageTestCases {
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
// Tests: GetPagedCollection
// =============================================================================

func Test_TypedPayloadCollection_GetPagedCollection(t *testing.T) {
	for caseIndex, testCase := range typedCollectionPagedCollectionTestCases {
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
// Tests: GetPagedCollectionWithInfo
// =============================================================================

func Test_TypedPayloadCollection_GetPagedCollectionWithInfo(t *testing.T) {
	for caseIndex, testCase := range typedCollectionPagedWithInfoTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count := input.GetDirectLower("count").(int)
		pageSize := input.GetDirectLower("pageSize").(int)

		collection := createNumberedUsers(count)

		// Act
		pages := collection.GetPagedCollectionWithInfo(pageSize)

		results := []string{fmt.Sprintf("%d", len(pages))}

		// Emit first 2 pages' PagingInfo fields: CurrentPageIndex, TotalPages, PerPageItems, TotalItems
		for i := 0; i < 2 && i < len(pages); i++ {
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
// Tests: Paging on empty collection
// =============================================================================

func Test_TypedPayloadCollection_Paging_Empty(t *testing.T) {
	for caseIndex, testCase := range typedCollectionPagingEmptyTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pageSize := input.GetDirectLower("pageSize").(int)

		collection := corepayload.EmptyTypedPayloadCollection[testUser]()

		// Act
		pages := collection.GetPagedCollection(pageSize)

		results := []string{
			fmt.Sprintf("%d", len(pages)),
			fmt.Sprintf("%d", pages[0].Length()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, results...)
	}
}
