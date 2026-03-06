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

func createTaggedUsers() []*corepayload.TypedPayloadWrapper[testUserWithTags] {
	users := []testUserWithTags{
		{Name: "Alice", Age: 30, Tags: []string{"go", "rust"}},
		{Name: "Bob", Age: 25, Tags: []string{"python", "java"}},
		{Name: "Carol", Age: 35, Tags: []string{"ts", "js"}},
	}

	wrappers := make([]*corepayload.TypedPayloadWrapper[testUserWithTags], 0, len(users))

	for i, user := range users {
		typed, err := corepayload.TypedPayloadWrapperNameIdRecord[testUserWithTags](
			user.Name,
			fmt.Sprintf("usr-%d", i),
			user,
		)
		errcore.HandleErr(err)

		wrappers = append(wrappers, typed)
	}

	return wrappers
}

func createTaggedCollection() *corepayload.TypedPayloadCollection[testUserWithTags] {
	return corepayload.TypedPayloadCollectionFrom[testUserWithTags](createTaggedUsers())
}

func buildTagsMap(allTags []string) args.Map {
	actual := args.Map{
		"count": len(allTags),
	}

	for i, tag := range allTags {
		actual[fmt.Sprintf("tag%d", i)] = tag
	}

	return actual
}

// =============================================================================
// FlatMapTypedPayloads — wrapper-level
// =============================================================================

func Test_TypedPayloadCollection_FlatMapTypedPayloads(t *testing.T) {
	for caseIndex, testCase := range flatMapTypedPayloadsTestCases {
		// Arrange
		collection := createTaggedCollection()

		// Act
		allTags := corepayload.FlatMapTypedPayloads[testUserWithTags, string](
			collection,
			func(item *corepayload.TypedPayloadWrapper[testUserWithTags]) []string {
				return item.Data().Tags
			},
		)

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, buildTagsMap(allTags))
	}
}

// =============================================================================
// FlatMapTypedPayloadData — data-level
// =============================================================================

func Test_TypedPayloadCollection_FlatMapTypedPayloadData(t *testing.T) {
	for caseIndex, testCase := range flatMapTypedPayloadDataTestCases {
		// Arrange
		collection := createTaggedCollection()

		// Act
		allTags := corepayload.FlatMapTypedPayloadData[testUserWithTags, string](
			collection,
			func(user testUserWithTags) []string {
				return user.Tags
			},
		)

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, buildTagsMap(allTags))
	}
}

// =============================================================================
// FlatMap on empty collection
// =============================================================================

func Test_TypedPayloadCollection_FlatMap_Empty(t *testing.T) {
	for caseIndex, testCase := range flatMapEmptyCollectionTestCases {
		// Arrange
		collection := corepayload.EmptyTypedPayloadCollection[testUserWithTags]()

		// Act
		allTags := corepayload.FlatMapTypedPayloads[testUserWithTags, string](
			collection,
			func(item *corepayload.TypedPayloadWrapper[testUserWithTags]) []string {
				return item.Data().Tags
			},
		)

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, args.Map{
			"count": len(allTags),
		})
	}
}

// =============================================================================
// FlatMap with mapper returning nil slices
// =============================================================================

func Test_TypedPayloadCollection_FlatMap_NoOutput(t *testing.T) {
	for caseIndex, testCase := range flatMapNoOutputTestCases {
		// Arrange
		collection := createTaggedCollection()

		// Act
		result := corepayload.FlatMapTypedPayloadData[testUserWithTags, string](
			collection,
			func(user testUserWithTags) []string {
				return nil
			},
		)

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, args.Map{
			"count": len(result),
		})
	}
}

// =============================================================================
// Edge: nil wrappers in collection
// =============================================================================

func Test_TypedPayloadCollection_NilWrapperEdge(t *testing.T) {
	for caseIndex, testCase := range nilWrapperEdgeCaseTestCases {
		// Arrange
		wrappers := createTaggedUsers()
		wrappers = append(wrappers, nil) // inject nil wrapper
		collection := corepayload.TypedPayloadCollectionFrom[testUserWithTags](wrappers)

		// Act
		isValid := collection.IsValid()
		length := collection.Length()

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, args.Map{
			"isValid": isValid,
			"length":  length,
		})
	}
}

// =============================================================================
// Edge: deserialization failure — TypedPayloadCollectionFromPayloads
// =============================================================================

func Test_TypedPayloadCollection_DeserializationFailure(t *testing.T) {
	for caseIndex, testCase := range deserializationFailureTestCases {
		// Arrange — create PayloadsCollection with 2 valid + 1 invalid wrapper
		validUsers := []testUser{
			{Name: "Alice", Email: "alice@test.com", Age: 30},
			{Name: "Bob", Email: "bob@test.com", Age: 25},
		}

		payloadsCollection := &corepayload.PayloadsCollection{
			Items: make([]*corepayload.PayloadWrapper, 0, 3),
		}

		for i, user := range validUsers {
			typed, err := corepayload.TypedPayloadWrapperNameIdRecord[testUser](
				user.Name, fmt.Sprintf("id-%d", i), user,
			)
			errcore.HandleErr(err)
			payloadsCollection.Items = append(payloadsCollection.Items, typed.ToPayloadWrapper())
		}

		// Add invalid wrapper with garbage payloads
		invalidWrapper := &corepayload.PayloadWrapper{
			Name:     "invalid",
			Payloads: []byte("{{not-valid-json"),
		}
		payloadsCollection.Items = append(payloadsCollection.Items, invalidWrapper)

		// Act
		collection := corepayload.TypedPayloadCollectionFromPayloads[testUser](payloadsCollection)

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, args.Map{
			"count": collection.Length(),
		})
	}
}

// =============================================================================
// Edge: TypedPayloadCollectionDeserialize with invalid bytes
// =============================================================================

func Test_TypedPayloadCollection_DeserializeInvalidBytes(t *testing.T) {
	for caseIndex, testCase := range collectionDeserializeInvalidTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		invalidBytes, _ := input.GetAsString("bytes")

		// Act
		_, err := corepayload.TypedPayloadCollectionDeserialize[testUser]([]byte(invalidBytes))

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, args.Map{
			"hasError": err != nil,
		})
	}
}

// =============================================================================
// Edge: nil receiver safety
// =============================================================================

func Test_TypedPayloadCollection_NilReceiver(t *testing.T) {
	for caseIndex, testCase := range nilReceiverTestCases {
		// Arrange
		var collection *corepayload.TypedPayloadCollection[testUser]

		// Act
		length := collection.Length()
		isEmpty := collection.IsEmpty()
		hasItems := collection.HasItems()

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, args.Map{
			"length":  length,
			"isEmpty": isEmpty,
			"noItems": !hasItems,
		})
	}
}
