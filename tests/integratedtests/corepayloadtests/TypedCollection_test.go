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

func createTestUsers() []*corepayload.TypedPayloadWrapper[testUser] {
	users := []testUser{
		{Name: "Alice", Email: "alice@test.com", Age: 30},
		{Name: "Bob", Email: "bob@test.com", Age: 25},
		{Name: "Carol", Email: "carol@test.com", Age: 35},
	}

	wrappers := make([]*corepayload.TypedPayloadWrapper[testUser], 0, len(users))

	for i, user := range users {
		category := "senior"
		if user.Age < 30 {
			category = "junior"
		}

		typed, err := corepayload.TypedPayloadWrapperNameIdCategory[testUser](
			user.Name,
			fmt.Sprintf("usr-%d", i),
			category,
			user,
		)
		errcore.HandleErr(err)

		wrappers = append(wrappers, typed)
	}

	return wrappers
}

func createTestCollection() *corepayload.TypedPayloadCollection[testUser] {
	return corepayload.TypedPayloadCollectionFrom[testUser](createTestUsers())
}

// =============================================================================
// Tests
// =============================================================================

func Test_TypedPayloadCollection_Creation(t *testing.T) {
	for caseIndex, testCase := range typedCollectionCreationTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		capacity := input.GetDirectLower("capacity").(int)

		// Act
		var collection *corepayload.TypedPayloadCollection[testUser]
		if capacity == 0 {
			collection = corepayload.EmptyTypedPayloadCollection[testUser]()
		} else {
			collection = corepayload.NewTypedPayloadCollection[testUser](capacity)
		}

		length := fmt.Sprintf("%d", collection.Length())
		isEmpty := fmt.Sprintf("%v", collection.IsEmpty())

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			length,
			isEmpty,
		)
	}
}

func Test_TypedPayloadCollection_Add(t *testing.T) {
	for caseIndex, testCase := range typedCollectionAddTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		name, _ := input.GetAsString("name")
		email, _ := input.GetAsString("email")
		age := input.GetDirectLower("age").(int)
		collection := corepayload.EmptyTypedPayloadCollection[testUser]()

		firstUser := testUser{Name: name, Email: email, Age: age}
		firstTyped, firstErr := corepayload.TypedPayloadWrapperNameIdRecord[testUser](
			name, "id-1", firstUser,
		)
		errcore.HandleErr(firstErr)

		// Act
		collection.Add(firstTyped)

		results := []string{
			fmt.Sprintf("%d", collection.Length()),
			fmt.Sprintf("%v", collection.IsEmpty()),
			collection.First().Data().Name,
		}

		name2, hasSecond := input.GetAsString("name2")
		if hasSecond {
			email2, _ := input.GetAsString("email2")
			age2 := input.GetDirectLower("age2").(int)
			secondUser := testUser{Name: name2, Email: email2, Age: age2}
			secondTyped, secondErr := corepayload.TypedPayloadWrapperNameIdRecord[testUser](
				name2, "id-2", secondUser,
			)
			errcore.HandleErr(secondErr)
			collection.Add(secondTyped)

			results = []string{
				fmt.Sprintf("%d", collection.Length()),
				fmt.Sprintf("%v", collection.IsEmpty()),
				collection.First().Data().Name,
				collection.Last().Data().Name,
			}
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, results...)
	}
}

func Test_TypedPayloadCollection_FilterByData(t *testing.T) {
	for caseIndex, testCase := range typedCollectionFilterTestCases {
		// Arrange
		collection := createTestCollection()

		// Act
		filtered := collection.FilterByData(func(user testUser) bool {
			return user.Age >= 30
		})

		results := []string{
			fmt.Sprintf("%d", filtered.Length()),
		}

		filtered.ForEachData(func(index int, data testUser) {
			results = append(results, data.Name)
		})

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, results...)
	}
}

func Test_TypedPayloadCollection_MapData(t *testing.T) {
	for caseIndex, testCase := range typedCollectionMapTestCases {
		// Arrange
		collection := createTestCollection()

		// Act
		names := corepayload.MapTypedPayloadData[testUser, string](
			collection,
			func(user testUser) string { return user.Name },
		)

		results := []string{fmt.Sprintf("%d", len(names))}
		results = append(results, names...)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, results...)
	}
}

func Test_TypedPayloadCollection_ReduceData(t *testing.T) {
	for caseIndex, testCase := range typedCollectionReduceTestCases {
		// Arrange
		collection := createTestCollection()

		// Act
		totalAge := corepayload.ReduceTypedPayloadData[testUser, int](
			collection,
			0,
			func(acc int, user testUser) int { return acc + user.Age },
		)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%d", totalAge))
	}
}

func Test_TypedPayloadCollection_GroupByCategory(t *testing.T) {
	for caseIndex, testCase := range typedCollectionGroupTestCases {
		// Arrange
		collection := createTestCollection()

		// Act
		groups := corepayload.GroupTypedPayloads[testUser, string](
			collection,
			func(item *corepayload.TypedPayloadWrapper[testUser]) string {
				return item.CategoryName()
			},
		)

		groupCount := fmt.Sprintf("%d", len(groups))
		juniorCount := fmt.Sprintf("%d", groups["junior"].Length())
		seniorCount := fmt.Sprintf("%d", groups["senior"].Length())

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			groupCount,
			juniorCount,
			seniorCount,
		)
	}
}

func Test_TypedPayloadCollection_Partition(t *testing.T) {
	for caseIndex, testCase := range typedCollectionPartitionTestCases {
		// Arrange
		collection := createTestCollection()

		// Act
		senior, junior := corepayload.PartitionTypedPayloads[testUser](
			collection,
			func(item *corepayload.TypedPayloadWrapper[testUser]) bool {
				return item.Data().Age >= 30
			},
		)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			fmt.Sprintf("%d", senior.Length()),
			fmt.Sprintf("%d", junior.Length()),
		)
	}
}

func Test_TypedPayloadCollection_AllData(t *testing.T) {
	for caseIndex, testCase := range typedCollectionAllDataTestCases {
		// Arrange
		collection := createTestCollection()

		// Act
		allData := collection.AllData()

		results := []string{fmt.Sprintf("%d", len(allData))}
		for _, user := range allData {
			results = append(results, user.Name)
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, results...)
	}
}

func Test_TypedPayloadCollection_ElementAccess(t *testing.T) {
	for caseIndex, testCase := range typedCollectionElementAccessTestCases {
		// Arrange
		collection := createTestCollection()

		// Act
		firstName := collection.First().Data().Name
		lastName := collection.Last().Data().Name

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			firstName,
			lastName,
		)
	}
}

func Test_TypedPayloadCollection_AnyAll(t *testing.T) {
	for caseIndex, testCase := range typedCollectionAnyAllTestCases {
		// Arrange
		collection := createTestCollection()

		// Act
		hasBob := corepayload.AnyTypedPayload[testUser](
			collection,
			func(item *corepayload.TypedPayloadWrapper[testUser]) bool {
				return item.Data().Name == "Bob"
			},
		)

		hasNonExistent := corepayload.AnyTypedPayload[testUser](
			collection,
			func(item *corepayload.TypedPayloadWrapper[testUser]) bool {
				return item.Data().Name == "Nonexistent"
			},
		)

		allParsed := corepayload.AllTypedPayloads[testUser](
			collection,
			func(item *corepayload.TypedPayloadWrapper[testUser]) bool {
				return item.IsParsed()
			},
		)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			fmt.Sprintf("%v", hasBob),
			fmt.Sprintf("%v", hasNonExistent),
			fmt.Sprintf("%v", allParsed),
		)
	}
}

func Test_TypedPayloadCollection_EmptyBehavior(t *testing.T) {
	collection := corepayload.EmptyTypedPayloadCollection[testUser]()

	// Verify empty collection operations don't panic
	allData := collection.AllData()
	names := corepayload.MapTypedPayloadData[testUser, string](
		collection, func(u testUser) string { return u.Name },
	)
	filtered := collection.FilterByData(func(u testUser) bool { return true })
	totalAge := corepayload.ReduceTypedPayloadData[testUser, int](
		collection, 0, func(acc int, u testUser) int { return acc + u.Age },
	)

	if len(allData) != 0 || len(names) != 0 || filtered.Length() != 0 || totalAge != 0 {
		t.Error("Empty collection operations should return empty/zero results")
	}
}

func Test_TypedPayloadCollection_FirstByName(t *testing.T) {
	collection := createTestCollection()

	found := collection.FirstByName("Bob")
	if found == nil || found.Data().Name != "Bob" {
		t.Error("FirstByName should find Bob")
	}

	notFound := collection.FirstByName("Nonexistent")
	if notFound != nil {
		t.Error("FirstByName should return nil for nonexistent")
	}
}

func Test_TypedPayloadCollection_RemoveAt(t *testing.T) {
	collection := createTestCollection()

	removed := collection.RemoveAt(1) // Remove Bob
	if !removed || collection.Length() != 2 {
		t.Error("RemoveAt should remove item and reduce length")
	}

	if collection.First().Data().Name != "Alice" || collection.Last().Data().Name != "Carol" {
		t.Error("Remaining items should be Alice and Carol")
	}

	invalidRemove := collection.RemoveAt(99)
	if invalidRemove {
		t.Error("RemoveAt with invalid index should return false")
	}
}

func Test_TypedPayloadCollection_ToPayloadsCollection(t *testing.T) {
	collection := createTestCollection()
	payloads := collection.ToPayloadsCollection()

	if payloads.Length() != 3 {
		t.Errorf("Expected 3 payloads, got %d", payloads.Length())
	}

	if payloads.First().Name != "Alice" {
		t.Errorf("Expected first payload name Alice, got %s", payloads.First().Name)
	}
}
