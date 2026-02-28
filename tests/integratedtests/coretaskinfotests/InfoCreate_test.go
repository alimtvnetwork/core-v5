package coretaskinfotests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretaskinfo"
	"gitlab.com/auk-go/core/coretests/args"
)

// ==========================================
// Test: Info.Default creation
// ==========================================

func Test_Info_Default_Verification(t *testing.T) {
	for caseIndex, testCase := range infoDefaultTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nameStr, _ := input.GetAsString("name")
		descStr, _ := input.GetAsString("desc")
		urlStr, _ := input.GetAsString("url")

		// Act
		info := coretaskinfo.New.Info.Default(nameStr, descStr, urlStr)
		actLines := []string{
			info.SafeName(),
			info.SafeDescription(),
			info.SafeUrl(),
			fmt.Sprintf("%v", info.IsNull()),
			fmt.Sprintf("%v", info.IsDefined()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Info.Examples with examples
// ==========================================

func Test_Info_ExamplesWithItems_Verification(t *testing.T) {
	for caseIndex, testCase := range infoExamplesWithItemsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nameStr, _ := input.GetAsString("name")
		descStr, _ := input.GetAsString("desc")
		urlStr, _ := input.GetAsString("url")
		examples, _ := input.GetAsStrings("examples")

		// Act
		info := coretaskinfo.New.Info.Examples(nameStr, descStr, urlStr, examples...)
		actLines := []string{
			info.SafeName(),
			info.SafeDescription(),
			info.SafeUrl(),
			fmt.Sprintf("%v", info.IsNull()),
			fmt.Sprintf("%v", info.IsDefined()),
			fmt.Sprintf("%v", info.HasExamples()),
			fmt.Sprintf("%d", len(info.Examples)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Info.Examples with no examples
// ==========================================

func Test_Info_ExamplesEmpty_Verification(t *testing.T) {
	for caseIndex, testCase := range infoExamplesEmptyTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nameStr, _ := input.GetAsString("name")
		descStr, _ := input.GetAsString("desc")
		urlStr, _ := input.GetAsString("url")

		// Act
		info := coretaskinfo.New.Info.Examples(nameStr, descStr, urlStr)
		actLines := []string{
			info.SafeName(),
			info.SafeDescription(),
			info.SafeUrl(),
			fmt.Sprintf("%v", info.IsNull()),
			fmt.Sprintf("%v", info.IsDefined()),
			fmt.Sprintf("%v", info.HasExamples()),
			fmt.Sprintf("%d", len(info.Examples)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Nil info — SafeName
// ==========================================

func Test_Info_Nil_SafeName_Verification(t *testing.T) {
	for caseIndex, testCase := range infoNilSafeNameTestCases {
		// Arrange
		var info *coretaskinfo.Info

		// Act
		actLines := []string{info.SafeName()}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Nil info — SafeDescription
// ==========================================

func Test_Info_Nil_SafeDescription_Verification(t *testing.T) {
	for caseIndex, testCase := range infoNilSafeDescriptionTestCases {
		// Arrange
		var info *coretaskinfo.Info

		// Act
		actLines := []string{info.SafeDescription()}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Nil info — SafeUrl
// ==========================================

func Test_Info_Nil_SafeUrl_Verification(t *testing.T) {
	for caseIndex, testCase := range infoNilSafeUrlTestCases {
		// Arrange
		var info *coretaskinfo.Info

		// Act
		actLines := []string{info.SafeUrl()}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Nil info — SafeHintUrl
// ==========================================

func Test_Info_Nil_SafeHintUrl_Verification(t *testing.T) {
	for caseIndex, testCase := range infoNilSafeHintUrlTestCases {
		// Arrange
		var info *coretaskinfo.Info

		// Act
		actLines := []string{info.SafeHintUrl()}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Nil info — SafeErrorUrl
// ==========================================

func Test_Info_Nil_SafeErrorUrl_Verification(t *testing.T) {
	for caseIndex, testCase := range infoNilSafeErrorUrlTestCases {
		// Arrange
		var info *coretaskinfo.Info

		// Act
		actLines := []string{info.SafeErrorUrl()}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Nil info — SafeExampleUrl
// ==========================================

func Test_Info_Nil_SafeExampleUrl_Verification(t *testing.T) {
	for caseIndex, testCase := range infoNilSafeExampleUrlTestCases {
		// Arrange
		var info *coretaskinfo.Info

		// Act
		actLines := []string{info.SafeExampleUrl()}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Nil info — NullCheck
// ==========================================

func Test_Info_Nil_NullCheck_Verification(t *testing.T) {
	for caseIndex, testCase := range infoNilNullCheckTestCases {
		// Arrange
		var info *coretaskinfo.Info

		// Act
		actLines := []string{
			fmt.Sprintf("%v", info.IsNull()),
			fmt.Sprintf("%v", info.IsDefined()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Nil info — EmptyCheck
// ==========================================

func Test_Info_Nil_EmptyCheck_Verification(t *testing.T) {
	for caseIndex, testCase := range infoNilEmptyCheckTestCases {
		// Arrange
		var info *coretaskinfo.Info

		// Act
		actLines := []string{
			fmt.Sprintf("%v", info.IsEmpty()),
			fmt.Sprintf("%v", info.HasAnyItem()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Nil info — ClonePtr
// ==========================================

func Test_Info_Nil_ClonePtr_Verification(t *testing.T) {
	for caseIndex, testCase := range infoNilClonePtrTestCases {
		// Arrange
		var info *coretaskinfo.Info

		// Act
		cloned := info.ClonePtr()
		actLines := []string{fmt.Sprintf("%v", cloned == nil)}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Nil info — PrettyJsonString
// ==========================================

func Test_Info_Nil_PrettyJsonString_Verification(t *testing.T) {
	for caseIndex, testCase := range infoNilPrettyJsonTestCases {
		// Arrange
		var info *coretaskinfo.Info

		// Act
		actLines := []string{info.PrettyJsonString()}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Secure.Default creation
// ==========================================

func Test_Info_SecureDefault_Verification(t *testing.T) {
	for caseIndex, testCase := range infoSecureDefaultTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nameStr, _ := input.GetAsString("name")
		descStr, _ := input.GetAsString("desc")
		urlStr, _ := input.GetAsString("url")

		// Act
		info := coretaskinfo.New.Info.Secure.Default(nameStr, descStr, urlStr)
		actLines := []string{
			info.SafeName(),
			info.SafeDescription(),
			info.SafeUrl(),
			fmt.Sprintf("%v", info.IsSecure()),
			fmt.Sprintf("%v", info.IsPlainText()),
			fmt.Sprintf("%v", info.IsExcludePayload()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Secure.NameDescUrlExamples creation
// ==========================================

func Test_Info_SecureExamples_Verification(t *testing.T) {
	for caseIndex, testCase := range infoSecureExamplesTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nameStr, _ := input.GetAsString("name")
		descStr, _ := input.GetAsString("desc")
		urlStr, _ := input.GetAsString("url")
		examples, _ := input.GetAsStrings("examples")

		// Act
		info := coretaskinfo.New.Info.Secure.NameDescUrlExamples(
			nameStr, descStr, urlStr, examples...)
		actLines := []string{
			info.SafeName(),
			fmt.Sprintf("%v", info.IsSecure()),
			fmt.Sprintf("%v", info.IsPlainText()),
			fmt.Sprintf("%v", info.IsExcludePayload()),
			fmt.Sprintf("%d", len(info.Examples)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: SetSecure on nil
// ==========================================

func Test_Info_SetSecureOnNil_Verification(t *testing.T) {
	for caseIndex, testCase := range infoSetSecureOnNilTestCases {
		// Arrange
		var info *coretaskinfo.Info

		// Act
		result := info.SetSecure()
		actLines := []string{
			fmt.Sprintf("%v", result.IsSecure()),
			fmt.Sprintf("%v", result.IsPlainText()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: SetSecure on existing plain info
// ==========================================

func Test_Info_SetSecureOnExisting_Verification(t *testing.T) {
	for caseIndex, testCase := range infoSetSecureOnExistingTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nameStr, _ := input.GetAsString("name")

		// Act
		info := coretaskinfo.New.Info.Plain.Default(nameStr, "d", "u")
		info.SetSecure()
		actLines := []string{
			fmt.Sprintf("%v", info.IsSecure()),
			fmt.Sprintf("%v", info.IsPlainText()),
			info.SafeName(),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Plain.Default creation
// ==========================================

func Test_Info_PlainDefault_Verification(t *testing.T) {
	for caseIndex, testCase := range infoPlainDefaultTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nameStr, _ := input.GetAsString("name")
		descStr, _ := input.GetAsString("desc")
		urlStr, _ := input.GetAsString("url")

		// Act
		info := coretaskinfo.New.Info.Plain.Default(nameStr, descStr, urlStr)
		actLines := []string{
			info.SafeName(),
			info.SafeDescription(),
			info.SafeUrl(),
			fmt.Sprintf("%v", info.IsSecure()),
			fmt.Sprintf("%v", info.IsPlainText()),
			fmt.Sprintf("%v", info.IsIncludePayloads()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Plain.AllUrlExamples creation
// ==========================================

func Test_Info_PlainAllUrlExamples_Verification(t *testing.T) {
	for caseIndex, testCase := range infoPlainAllUrlExamplesTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nameStr, _ := input.GetAsString("name")
		descStr, _ := input.GetAsString("desc")
		urlStr, _ := input.GetAsString("url")
		hintUrl, _ := input.GetAsString("hintUrl")
		errorUrl, _ := input.GetAsString("errorUrl")
		examples, _ := input.GetAsStrings("examples")

		// Act
		info := coretaskinfo.New.Info.Plain.AllUrlExamples(
			nameStr, descStr, urlStr, hintUrl, errorUrl, examples...)
		actLines := []string{
			info.SafeName(),
			info.SafeDescription(),
			info.SafeUrl(),
			info.SafeHintUrl(),
			info.SafeErrorUrl(),
			fmt.Sprintf("%v", info.IsSecure()),
			fmt.Sprintf("%v", info.IsPlainText()),
			fmt.Sprintf("%d", len(info.Examples)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: SetPlain on nil
// ==========================================

func Test_Info_SetPlainOnNil_Verification(t *testing.T) {
	for caseIndex, testCase := range infoSetPlainOnNilTestCases {
		// Arrange
		var info *coretaskinfo.Info

		// Act
		result := info.SetPlain()
		actLines := []string{
			fmt.Sprintf("%v", result.IsSecure()),
			fmt.Sprintf("%v", result.IsPlainText()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Serialize Default round-trip
// ==========================================

func Test_Info_SerializeDefault_Verification(t *testing.T) {
	for caseIndex, testCase := range infoSerializeDefaultTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nameStr, _ := input.GetAsString("name")
		descStr, _ := input.GetAsString("desc")
		urlStr, _ := input.GetAsString("url")

		// Act
		original := coretaskinfo.New.Info.Default(nameStr, descStr, urlStr)
		jsonBytes := original.JsonPtr().Bytes
		deserialized, err := coretaskinfo.New.Info.Deserialized(jsonBytes)
		actLines := []string{
			deserialized.SafeName(),
			deserialized.SafeDescription(),
			deserialized.SafeUrl(),
			fmt.Sprintf("%v", err == nil),
			fmt.Sprintf("%v", deserialized.IsSecure()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Serialize Secure round-trip
// ==========================================

func Test_Info_SerializeSecure_Verification(t *testing.T) {
	for caseIndex, testCase := range infoSerializeSecureTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nameStr, _ := input.GetAsString("name")
		descStr, _ := input.GetAsString("desc")
		urlStr, _ := input.GetAsString("url")

		// Act
		original := coretaskinfo.New.Info.Secure.Default(nameStr, descStr, urlStr)
		jsonBytes := original.JsonPtr().Bytes
		deserialized, err := coretaskinfo.New.Info.Deserialized(jsonBytes)
		actLines := []string{
			deserialized.SafeName(),
			deserialized.SafeDescription(),
			deserialized.SafeUrl(),
			fmt.Sprintf("%v", err == nil),
			fmt.Sprintf("%v", deserialized.IsSecure()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Serialize with examples round-trip
// ==========================================

func Test_Info_SerializeExamples_Verification(t *testing.T) {
	for caseIndex, testCase := range infoSerializeExamplesTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nameStr, _ := input.GetAsString("name")
		descStr, _ := input.GetAsString("desc")
		urlStr, _ := input.GetAsString("url")
		examples, _ := input.GetAsStrings("examples")

		// Act
		original := coretaskinfo.New.Info.Examples(nameStr, descStr, urlStr, examples...)
		jsonBytes := original.JsonPtr().Bytes
		deserialized, err := coretaskinfo.New.Info.Deserialized(jsonBytes)
		actLines := []string{
			deserialized.SafeName(),
			fmt.Sprintf("%v", err == nil),
			fmt.Sprintf("%d", len(deserialized.Examples)),
		}
		for _, ex := range deserialized.Examples {
			actLines = append(actLines, ex)
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Serialize with all URLs round-trip
// ==========================================

func Test_Info_SerializeAllUrls_Verification(t *testing.T) {
	for caseIndex, testCase := range infoSerializeAllUrlsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nameStr, _ := input.GetAsString("name")
		descStr, _ := input.GetAsString("desc")
		urlStr, _ := input.GetAsString("url")
		hintUrl, _ := input.GetAsString("hintUrl")
		errorUrl, _ := input.GetAsString("errorUrl")

		// Act
		original := coretaskinfo.New.Info.Plain.AllUrl(nameStr, descStr, urlStr, hintUrl, errorUrl)
		jsonBytes := original.JsonPtr().Bytes
		deserialized, err := coretaskinfo.New.Info.Deserialized(jsonBytes)
		actLines := []string{
			deserialized.SafeName(),
			deserialized.SafeUrl(),
			deserialized.SafeHintUrl(),
			deserialized.SafeErrorUrl(),
			fmt.Sprintf("%v", err == nil),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Clone
// ==========================================

func Test_Info_Clone_Verification(t *testing.T) {
	for caseIndex, testCase := range infoCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nameStr, _ := input.GetAsString("name")
		descStr, _ := input.GetAsString("desc")
		urlStr, _ := input.GetAsString("url")
		newName, _ := input.GetAsString("newName")

		// Act
		original := coretaskinfo.New.Info.Default(nameStr, descStr, urlStr)
		cloned := original.Clone()
		cloned.RootName = newName

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			original.SafeName(),
			cloned.SafeName(),
			cloned.SafeDescription(),
		)
	}
}

// ==========================================
// Test: Field checks — populated
// ==========================================

func Test_Info_FieldChecks_Populated_Verification(t *testing.T) {
	for caseIndex, testCase := range infoFieldCheckPopulatedTestCases {
		// Arrange
		info := coretaskinfo.New.Info.Secure.AllUrlExamples(
			"name", "desc",
			"url", "hint", "err",
			"ex1", "ex2",
		)
		info.SingleExample = "single"

		// Act
		actLines := []string{
			fmt.Sprintf("%v", info.HasRootName()),
			fmt.Sprintf("%v", info.HasDescription()),
			fmt.Sprintf("%v", info.HasUrl()),
			fmt.Sprintf("%v", info.HasHintUrl()),
			fmt.Sprintf("%v", info.HasErrorUrl()),
			fmt.Sprintf("%v", info.HasExamples()),
			fmt.Sprintf("%v", info.HasChainingExample()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Field checks — empty
// ==========================================

func Test_Info_FieldChecks_Empty_Verification(t *testing.T) {
	for caseIndex, testCase := range infoFieldCheckEmptyTestCases {
		// Arrange
		info := &coretaskinfo.Info{}

		// Act
		actLines := []string{
			fmt.Sprintf("%v", info.HasRootName()),
			fmt.Sprintf("%v", info.HasDescription()),
			fmt.Sprintf("%v", info.HasUrl()),
			fmt.Sprintf("%v", info.HasHintUrl()),
			fmt.Sprintf("%v", info.HasErrorUrl()),
			fmt.Sprintf("%v", info.HasExamples()),
			fmt.Sprintf("%v", info.HasChainingExample()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
