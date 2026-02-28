package coretaskinfotests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretaskinfo"
	"gitlab.com/auk-go/core/coretests/args"
)

// ==========================================
// Test: Info Creation
// ==========================================

func Test_Info_Create_Verification(t *testing.T) {
	for caseIndex, testCase := range infoCreateTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nameStr, _ := input.GetAsString("name")
		descStr, _ := input.GetAsString("desc")
		urlStr, _ := input.GetAsString("url")
		examples, _ := input.GetAsStrings("examples")
		noExamples := input.GetDirectLower("noExamples")

		// Act
		var info *coretaskinfo.Info
		if len(examples) > 0 {
			info = coretaskinfo.New.Info.Examples(nameStr, descStr, urlStr, examples...)
		} else if noExamples == true {
			info = coretaskinfo.New.Info.Examples(nameStr, descStr, urlStr)
		} else {
			info = coretaskinfo.New.Info.Default(nameStr, descStr, urlStr)
		}

		actLines := []string{
			info.SafeName(),
			info.SafeDescription(),
			info.SafeUrl(),
			fmt.Sprintf("%v", info.IsNull()),
			fmt.Sprintf("%v", info.IsDefined()),
		}

		if len(examples) > 0 || noExamples == true {
			actLines = append(actLines,
				fmt.Sprintf("%v", info.HasExamples()),
				fmt.Sprintf("%d", len(info.Examples)),
			)
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Nil Safety
// ==========================================

func Test_Info_NilSafety_Verification(t *testing.T) {
	for caseIndex, testCase := range infoNilSafetyTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		method, _ := input.GetAsString("method")
		var info *coretaskinfo.Info // nil

		// Act
		var actLines []string

		switch method {
		case "SafeName":
			actLines = []string{info.SafeName()}
		case "SafeDescription":
			actLines = []string{info.SafeDescription()}
		case "SafeUrl":
			actLines = []string{info.SafeUrl()}
		case "SafeHintUrl":
			actLines = []string{info.SafeHintUrl()}
		case "SafeErrorUrl":
			actLines = []string{info.SafeErrorUrl()}
		case "SafeExampleUrl":
			actLines = []string{info.SafeExampleUrl()}
		case "NullCheck":
			actLines = []string{
				fmt.Sprintf("%v", info.IsNull()),
				fmt.Sprintf("%v", info.IsDefined()),
			}
		case "EmptyCheck":
			actLines = []string{
				fmt.Sprintf("%v", info.IsEmpty()),
				fmt.Sprintf("%v", info.HasAnyItem()),
			}
		case "ClonePtr":
			cloned := info.ClonePtr()
			actLines = []string{fmt.Sprintf("%v", cloned == nil)}
		case "PrettyJsonString":
			actLines = []string{info.PrettyJsonString()}
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Secure Mode
// ==========================================

func Test_Info_SecureMode_Verification(t *testing.T) {
	for caseIndex, testCase := range infoSecureModeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		setSecureVal := input.GetDirectLower("setSecure")
		existingVal := input.GetDirectLower("existing")

		var actLines []string

		if setSecureVal == true && existingVal != true {
			// Act — SetSecure on nil
			var nilInfo *coretaskinfo.Info
			result := nilInfo.SetSecure()
			actLines = []string{
				fmt.Sprintf("%v", result.IsSecure()),
				fmt.Sprintf("%v", result.IsPlainText()),
			}
		} else if setSecureVal == true && existingVal == true {
			// Act — SetSecure on existing plain info
			nameStr, _ := input.GetAsString("name")
			info := coretaskinfo.New.Info.Plain.Default(nameStr, "d", "u")
			info.SetSecure()
			actLines = []string{
				fmt.Sprintf("%v", info.IsSecure()),
				fmt.Sprintf("%v", info.IsPlainText()),
				info.SafeName(),
			}
		} else {
			// Act — Secure creator
			nameStr, _ := input.GetAsString("name")
			descStr, _ := input.GetAsString("desc")
			urlStr, _ := input.GetAsString("url")
			examples, _ := input.GetAsStrings("examples")

			if len(examples) > 0 {
				info := coretaskinfo.New.Info.Secure.NameDescUrlExamples(
					nameStr, descStr, urlStr, examples...)
				actLines = []string{
					info.SafeName(),
					fmt.Sprintf("%v", info.IsSecure()),
					fmt.Sprintf("%v", info.IsPlainText()),
					fmt.Sprintf("%v", info.IsExcludePayload()),
					fmt.Sprintf("%d", len(info.Examples)),
				}
			} else {
				info := coretaskinfo.New.Info.Secure.Default(nameStr, descStr, urlStr)
				actLines = []string{
					info.SafeName(),
					info.SafeDescription(),
					info.SafeUrl(),
					fmt.Sprintf("%v", info.IsSecure()),
					fmt.Sprintf("%v", info.IsPlainText()),
					fmt.Sprintf("%v", info.IsExcludePayload()),
				}
			}
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Plain Mode
// ==========================================

func Test_Info_PlainMode_Verification(t *testing.T) {
	for caseIndex, testCase := range infoPlainModeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		setPlainVal := input.GetDirectLower("setPlain")

		var actLines []string

		if setPlainVal == true {
			// Act — SetPlain on nil
			var nilInfo *coretaskinfo.Info
			result := nilInfo.SetPlain()
			actLines = []string{
				fmt.Sprintf("%v", result.IsSecure()),
				fmt.Sprintf("%v", result.IsPlainText()),
			}
		} else {
			nameStr, _ := input.GetAsString("name")
			descStr, _ := input.GetAsString("desc")
			urlStr, _ := input.GetAsString("url")
			hintUrl, _ := input.GetAsString("hintUrl")
			errorUrl, _ := input.GetAsString("errorUrl")
			examples, _ := input.GetAsStrings("examples")

			if hintUrl != "" {
				// Act — AllUrlExamples
				info := coretaskinfo.New.Info.Plain.AllUrlExamples(
					nameStr, descStr, urlStr, hintUrl, errorUrl, examples...)
				actLines = []string{
					info.SafeName(),
					info.SafeDescription(),
					info.SafeUrl(),
					info.SafeHintUrl(),
					info.SafeErrorUrl(),
					fmt.Sprintf("%v", info.IsSecure()),
					fmt.Sprintf("%v", info.IsPlainText()),
					fmt.Sprintf("%d", len(info.Examples)),
				}
			} else {
				// Act — Default
				info := coretaskinfo.New.Info.Plain.Default(nameStr, descStr, urlStr)
				actLines = []string{
					info.SafeName(),
					info.SafeDescription(),
					info.SafeUrl(),
					fmt.Sprintf("%v", info.IsSecure()),
					fmt.Sprintf("%v", info.IsPlainText()),
					fmt.Sprintf("%v", info.IsIncludePayloads()),
				}
			}
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: JSON Serialization Round-Trip
// ==========================================

func Test_Info_Serialize_Verification(t *testing.T) {
	for caseIndex, testCase := range infoSerializeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nameStr, _ := input.GetAsString("name")
		descStr, _ := input.GetAsString("desc")
		urlStr, _ := input.GetAsString("url")
		hintUrl, _ := input.GetAsString("hintUrl")
		errorUrl, _ := input.GetAsString("errorUrl")
		isSecure := input.GetDirectLower("isSecure")
		examples, _ := input.GetAsStrings("examples")

		// Act — create original
		var original *coretaskinfo.Info
		switch {
		case isSecure == true:
			original = coretaskinfo.New.Info.Secure.Default(nameStr, descStr, urlStr)
		case len(examples) > 0:
			original = coretaskinfo.New.Info.Examples(nameStr, descStr, urlStr, examples...)
		case hintUrl != "":
			original = coretaskinfo.New.Info.Plain.AllUrl(nameStr, descStr, urlStr, hintUrl, errorUrl)
		default:
			original = coretaskinfo.New.Info.Default(nameStr, descStr, urlStr)
		}

		// Act — serialize then deserialize
		jsonBytes := original.JsonPtr().Bytes
		deserialized, err := coretaskinfo.New.Info.Deserialized(jsonBytes)
		noErr := err == nil

		var actLines []string

		switch {
		case len(examples) > 0:
			actLines = []string{
				deserialized.SafeName(),
				fmt.Sprintf("%v", noErr),
				fmt.Sprintf("%d", len(deserialized.Examples)),
			}
			for _, ex := range deserialized.Examples {
				actLines = append(actLines, ex)
			}
		case hintUrl != "":
			actLines = []string{
				deserialized.SafeName(),
				deserialized.SafeUrl(),
				deserialized.SafeHintUrl(),
				deserialized.SafeErrorUrl(),
				fmt.Sprintf("%v", noErr),
			}
		default:
			actLines = []string{
				deserialized.SafeName(),
				deserialized.SafeDescription(),
				deserialized.SafeUrl(),
				fmt.Sprintf("%v", noErr),
				fmt.Sprintf("%v", deserialized.IsSecure()),
			}
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

		// Assert — original should NOT be mutated
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
// Test: Field Has/IsEmpty Checks
// ==========================================

func Test_Info_FieldChecks_Verification(t *testing.T) {
	for caseIndex, testCase := range infoFieldCheckTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isEmpty := input.GetDirectLower("empty")

		var info *coretaskinfo.Info
		if isEmpty == true {
			info = &coretaskinfo.Info{}
		} else {
			info = coretaskinfo.New.Info.Secure.AllUrlExamples(
				"name", "desc",
				"url", "hint", "err",
				"ex1", "ex2",
			)
			info.SingleExample = "single"
		}

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
