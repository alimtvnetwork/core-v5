package corepayloadtests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coredata/corepayload"
	"gitlab.com/auk-go/core/errcore"
)

// =============================================================================
// Helpers
// =============================================================================

func createTypedProduct(input args.Map) *corepayload.TypedPayloadWrapper[testProduct] {
	name, _ := input.GetAsString("name")
	id, _ := input.GetAsString("id")
	sku, _ := input.GetAsString("sku")
	title, _ := input.GetAsString("title")
	price := input.GetDirectLower("price").(float64)

	product := testProduct{SKU: sku, Title: title, Price: price}
	typed, err := corepayload.TypedPayloadWrapperNameIdRecord[testProduct](name, id, product)
	errcore.HandleErr(err)

	return typed
}

// =============================================================================
// Tests
// =============================================================================

func Test_TypedPayloadWrapper_Deserialization(t *testing.T) {
	for caseIndex, testCase := range typedWrapperDeserializationTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		typed := createTypedProduct(input)

		// Act — serialize then deserialize
		serialized, serializeErr := typed.Serialize()
		errcore.HandleErr(serializeErr)

		deserialized, deserializeErr := corepayload.TypedPayloadWrapperDeserialize[testProduct](serialized)
		errcore.HandleErr(deserializeErr)

		data := deserialized.Data()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			deserialized.Name(),
			deserialized.Identifier(),
			data.SKU,
			data.Title,
			fmt.Sprintf("%.2f", data.Price),
		)
	}
}

func Test_TypedPayloadWrapper_RoundTrip(t *testing.T) {
	for caseIndex, testCase := range typedWrapperRoundTripTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		original := createTypedProduct(input)

		// Act — full round trip: create → serialize → deserialize → verify
		jsonResult := original.JsonPtr()
		restored, restoreErr := corepayload.TypedPayloadWrapperDeserializeUsingJsonResult[testProduct](jsonResult)
		errcore.HandleErr(restoreErr)

		restoredData := restored.Data()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			restored.Name(),
			restored.Identifier(),
			restoredData.SKU,
			restoredData.Title,
			fmt.Sprintf("%.2f", restoredData.Price),
		)
	}
}

func Test_TypedPayloadWrapper_DeepClone(t *testing.T) {
	for caseIndex, testCase := range typedWrapperCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		original := createTypedProduct(input)

		// Act — clone then mutate clone
		cloned, cloneErr := original.ClonePtr(true)
		errcore.HandleErr(cloneErr)

		mutatedProduct := testProduct{
			SKU:   cloned.Data().SKU,
			Title: "Modified",
			Price: cloned.Data().Price,
		}
		setErr := cloned.SetTypedData(mutatedProduct)
		errcore.HandleErr(setErr)

		originalData := original.Data()
		clonedData := cloned.Data()

		// Assert — original unchanged, clone mutated
		testCase.ShouldBeEqual(t, caseIndex,
			original.Name(),
			original.Identifier(),
			originalData.SKU,
			originalData.Title,
			fmt.Sprintf("%.2f", originalData.Price),
			clonedData.Title,
		)
	}
}

func Test_TypedPayloadWrapper_SetTypedData(t *testing.T) {
	for caseIndex, testCase := range typedWrapperSetDataTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		typed := createTypedProduct(input)

		newTitle, _ := input.GetAsString("new_title")
		newPrice := input.GetDirectLower("new_price").(float64)

		// Act — update data
		updatedProduct := testProduct{
			SKU:   typed.Data().SKU,
			Title: newTitle,
			Price: newPrice,
		}
		setErr := typed.SetTypedData(updatedProduct)
		errcore.HandleErr(setErr)

		// Verify raw payloads also updated by re-parsing
		reparsed, reparseErr := corepayload.NewTypedPayloadWrapper[testProduct](typed.ToPayloadWrapper())
		errcore.HandleErr(reparseErr)

		directData := typed.Data()
		reparsedData := reparsed.Data()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			directData.Title,
			fmt.Sprintf("%.2f", directData.Price),
			reparsedData.Title,
			fmt.Sprintf("%.2f", reparsedData.Price),
		)
	}
}

func Test_TypedPayloadWrapper_NilWrapper(t *testing.T) {
	tc := typedWrapperNilAndInvalidTestCases[0]

	// Act
	_, err := corepayload.NewTypedPayloadWrapper[testProduct](nil)

	actLines := []string{fmt.Sprintf("%v", err != nil)}
	expectedLines := tc.ExpectedInput.([]string)

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, expectedLines)
}

func Test_TypedPayloadWrapper_InvalidJson(t *testing.T) {
	tc := typedWrapperNilAndInvalidTestCases[1]

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	invalidBytes, _ := input.GetAsString("bytes")

	// Act
	_, err := corepayload.TypedPayloadWrapperDeserialize[testProduct]([]byte(invalidBytes))

	actLines := []string{fmt.Sprintf("%v", err != nil)}
	expectedLines := tc.ExpectedInput.([]string)

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, expectedLines)
}

func Test_TypedPayloadWrapper_DeserializeToMany(t *testing.T) {
	for caseIndex, testCase := range typedWrapperDeserializeToManyTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count := input.GetDirectLower("count").(int)

		wrappers := make([]*corepayload.TypedPayloadWrapper[testProduct], 0, count)

		for i := 0; i < count; i++ {
			product := testProduct{
				SKU:   fmt.Sprintf("SKU-%d", i),
				Title: fmt.Sprintf("item-%d", i),
				Price: float64(i) * 10.0,
			}

			typed, createErr := corepayload.TypedPayloadWrapperNameIdRecord[testProduct](
				fmt.Sprintf("item-%d", i),
				fmt.Sprintf("id-%d", i),
				product,
			)
			errcore.HandleErr(createErr)
			wrappers = append(wrappers, typed)
		}

		// Serialize all to JSON array
		payloadWrappers := make([]*corepayload.PayloadWrapper, len(wrappers))

		for i, w := range wrappers {
			payloadWrappers[i] = w.ToPayloadWrapper()
		}

		jsonSlice := corejson.Serialize.Apply(payloadWrappers)
		jsonSlice.HandleError()

		// Act
		deserialized, deserializeErr := corepayload.TypedPayloadWrapperDeserializeToMany[testProduct](jsonSlice.Bytes)
		errcore.HandleErr(deserializeErr)

		results := []string{fmt.Sprintf("%d", len(deserialized))}

		for _, item := range deserialized {
			results = append(results, item.Data().Title)
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, results...)
	}
}

func Test_TypedPayloadWrapper_MetadataAccessors(t *testing.T) {
	product := testProduct{SKU: "META-1", Title: "Meta Test", Price: 42.0}
	typed, err := corepayload.TypedPayloadWrapperNameIdCategory[testProduct](
		"meta-name", "meta-id", "meta-category", product,
	)
	errcore.HandleErr(err)

	if typed.Name() != "meta-name" {
		t.Errorf("Expected name 'meta-name', got '%s'", typed.Name())
	}

	if typed.Identifier() != "meta-id" {
		t.Errorf("Expected id 'meta-id', got '%s'", typed.Identifier())
	}

	if typed.CategoryName() != "meta-category" {
		t.Errorf("Expected category 'meta-category', got '%s'", typed.CategoryName())
	}

	if !typed.IsParsed() {
		t.Error("Expected IsParsed to be true")
	}

	if typed.IsEmpty() {
		t.Error("Expected IsEmpty to be false")
	}

	if typed.HasError() {
		t.Error("Expected HasError to be false")
	}

	if typed.HasSingleRecord() != true {
		t.Error("Expected HasSingleRecord to be true")
	}

	payloadsStr := typed.PayloadsString()

	if payloadsStr == "" {
		t.Error("Expected non-empty PayloadsString")
	}
}

func Test_TypedPayloadWrapper_TypedDataJson(t *testing.T) {
	product := testProduct{SKU: "JSON-1", Title: "Json Test", Price: 99.99}
	typed, err := corepayload.TypedPayloadWrapperNameIdRecord[testProduct](
		"json-test", "jt-1", product,
	)
	errcore.HandleErr(err)

	dataJson := typed.TypedDataJson()

	if dataJson.IsEmpty() {
		t.Error("Expected non-empty TypedDataJson result")
	}

	dataJsonPtr := typed.TypedDataJsonPtr()

	if dataJsonPtr == nil || dataJsonPtr.IsEmpty() {
		t.Error("Expected non-empty TypedDataJsonPtr result")
	}

	jsonBytes, jsonErr := typed.TypedDataJsonBytes()

	if jsonErr != nil || len(jsonBytes) == 0 {
		t.Error("Expected non-empty TypedDataJsonBytes")
	}
}
