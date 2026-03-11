package coretaskinfotests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================
// Info.Default creation
// ==========================================

var infoDefaultTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Info.Default creates info with name, desc, url",
		ArrangeInput: args.Map{
			"when": "given default info creation",
			"name": "some name",
			"desc": "some desc",
			"url":  "some url",
		},
		ExpectedInput: args.Map{
			"name":      "some name",
			"desc":      "some desc",
			"url":       "some url",
			"isNull":    "false",
			"isDefined": "true",
		},
	},
}

// ==========================================
// Info.Examples with items
// ==========================================

var infoExamplesWithItemsTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Info.Examples creates info with examples",
		ArrangeInput: args.Map{
			"when":     "given info with examples",
			"name":     "example name",
			"desc":     "example desc",
			"url":      "example url",
			"examples": []string{"ex1", "ex2"},
		},
		ExpectedInput: args.Map{
			"name":         "example name",
			"desc":         "example desc",
			"url":          "example url",
			"isNull":       "false",
			"isDefined":    "true",
			"hasExamples":  "true",
			"exampleCount": "2",
		},
	},
}

// ==========================================
// Info.Examples with no examples
// ==========================================

var infoExamplesEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Info.Examples with no examples has empty slice",
		ArrangeInput: args.Map{
			"when": "given info with no examples",
			"name": "no-ex name",
			"desc": "no-ex desc",
			"url":  "no-ex url",
		},
		ExpectedInput: args.Map{
			"name":         "no-ex name",
			"desc":         "no-ex desc",
			"url":          "no-ex url",
			"isNull":       "false",
			"isDefined":    "true",
			"hasExamples":  "false",
			"exampleCount": "0",
		},
	},
}

// ==========================================
// Nil Safety — one case per method
// ==========================================

var infoNilSafeNameTestCases = []coretestcases.CaseV1{
	{
		Title: "Nil info SafeName returns empty",
		ArrangeInput: args.Map{
			"when": "given nil info",
		},
		ExpectedInput: "",
	},
}

var infoNilSafeDescriptionTestCases = []coretestcases.CaseV1{
	{
		Title: "Nil info SafeDescription returns empty",
		ArrangeInput: args.Map{
			"when": "given nil info",
		},
		ExpectedInput: "",
	},
}

var infoNilSafeUrlTestCases = []coretestcases.CaseV1{
	{
		Title: "Nil info SafeUrl returns empty",
		ArrangeInput: args.Map{
			"when": "given nil info",
		},
		ExpectedInput: "",
	},
}

var infoNilSafeHintUrlTestCases = []coretestcases.CaseV1{
	{
		Title: "Nil info SafeHintUrl returns empty",
		ArrangeInput: args.Map{
			"when": "given nil info",
		},
		ExpectedInput: "",
	},
}

var infoNilSafeErrorUrlTestCases = []coretestcases.CaseV1{
	{
		Title: "Nil info SafeErrorUrl returns empty",
		ArrangeInput: args.Map{
			"when": "given nil info",
		},
		ExpectedInput: "",
	},
}

var infoNilSafeExampleUrlTestCases = []coretestcases.CaseV1{
	{
		Title: "Nil info SafeExampleUrl returns empty",
		ArrangeInput: args.Map{
			"when": "given nil info",
		},
		ExpectedInput: "",
	},
}

var infoNilNullCheckTestCases = []coretestcases.CaseV1{
	{
		Title: "Nil info IsNull returns true, IsDefined returns false",
		ArrangeInput: args.Map{
			"when": "given nil info for null check",
		},
		ExpectedInput: args.Map{
			"isNull":    "true",
			"isDefined": "false",
		},
	},
}

var infoNilEmptyCheckTestCases = []coretestcases.CaseV1{
	{
		Title: "Nil info IsEmpty returns true, HasAnyItem returns false",
		ArrangeInput: args.Map{
			"when": "given nil info for empty check",
		},
		ExpectedInput: args.Map{
			"isEmpty":    "true",
			"hasAnyItem": "false",
		},
	},
}

var infoNilClonePtrTestCases = []coretestcases.CaseV1{
	{
		Title: "Nil info ClonePtr returns nil",
		ArrangeInput: args.Map{
			"when": "given nil info for clone",
		},
		ExpectedInput: "true",
	},
}

var infoNilPrettyJsonTestCases = []coretestcases.CaseV1{
	{
		Title: "Nil info PrettyJsonString returns empty",
		ArrangeInput: args.Map{
			"when": "given nil info for json",
		},
		ExpectedInput: "",
	},
}

// ==========================================
// Secure Mode — separate cases
// ==========================================

var infoSecureDefaultTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Info.Secure.Default creates secure info",
		ArrangeInput: args.Map{
			"when": "given secure default creation",
			"name": "secure-name",
			"desc": "secure-desc",
			"url":  "secure-url",
		},
		ExpectedInput: args.Map{
			"name":             "secure-name",
			"desc":             "secure-desc",
			"url":              "secure-url",
			"isSecure":         "true",
			"isPlainText":      "false",
			"isExcludePayload": "true",
		},
	},
}

var infoSecureExamplesTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Info.Secure.NameDescUrlExamples has secure flag and examples",
		ArrangeInput: args.Map{
			"when":     "given secure with examples",
			"name":     "sec-ex-name",
			"desc":     "sec-ex-desc",
			"url":      "sec-ex-url",
			"examples": []string{"ex1", "ex2", "ex3"},
		},
		ExpectedInput: args.Map{
			"name":             "sec-ex-name",
			"isSecure":         "true",
			"isPlainText":      "false",
			"isExcludePayload": "true",
			"exampleCount":     "3",
		},
	},
}

var infoSetSecureOnNilTestCases = []coretestcases.CaseV1{
	{
		Title: "SetSecure on nil returns new secure info",
		ArrangeInput: args.Map{
			"when": "given nil info with SetSecure",
		},
		ExpectedInput: args.Map{
			"isSecure":    "true",
			"isPlainText": "false",
		},
	},
}

var infoSetSecureOnExistingTestCases = []coretestcases.CaseV1{
	{
		Title: "SetSecure on existing info mutates to secure",
		ArrangeInput: args.Map{
			"when": "given plain info then SetSecure",
			"name": "was-plain",
		},
		ExpectedInput: args.Map{
			"isSecure":    "true",
			"isPlainText": "false",
			"name":        "was-plain",
		},
	},
}

// ==========================================
// Plain Mode — separate cases
// ==========================================

var infoPlainDefaultTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Info.Plain.Default creates plain info",
		ArrangeInput: args.Map{
			"when": "given plain default creation",
			"name": "plain-name",
			"desc": "plain-desc",
			"url":  "plain-url",
		},
		ExpectedInput: args.Map{
			"name":              "plain-name",
			"desc":              "plain-desc",
			"url":               "plain-url",
			"isSecure":          "false",
			"isPlainText":       "true",
			"isIncludePayloads": "true",
		},
	},
}

var infoPlainAllUrlExamplesTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Info.Plain.AllUrlExamples populates all fields",
		ArrangeInput: args.Map{
			"when":     "given plain with all urls",
			"name":     "all-name",
			"desc":     "all-desc",
			"url":      "all-url",
			"hintUrl":  "all-hint",
			"errorUrl": "all-err",
			"examples": []string{"ex1"},
		},
		ExpectedInput: args.Map{
			"name":         "all-name",
			"desc":         "all-desc",
			"url":          "all-url",
			"hintUrl":      "all-hint",
			"errorUrl":     "all-err",
			"isSecure":     "false",
			"isPlainText":  "true",
			"exampleCount": "1",
		},
	},
}

var infoSetPlainOnNilTestCases = []coretestcases.CaseV1{
	{
		Title: "SetPlain on nil returns new plain info",
		ArrangeInput: args.Map{
			"when": "given nil info with SetPlain",
		},
		ExpectedInput: args.Map{
			"isSecure":    "false",
			"isPlainText": "true",
		},
	},
}

// ==========================================
// JSON Serialization — separate cases
// ==========================================

var infoSerializeDefaultTestCases = []coretestcases.CaseV1{
	{
		Title: "Default info serializes and deserializes correctly",
		ArrangeInput: args.Map{
			"when": "given round-trip serialize/deserialize",
			"name": "round-trip name",
			"desc": "round-trip desc",
			"url":  "round-trip url",
		},
		ExpectedInput: args.Map{
			"name":     "round-trip name",
			"desc":     "round-trip desc",
			"url":      "round-trip url",
			"noError":  "true",
			"isSecure": "false",
		},
	},
}

var infoSerializeSecureTestCases = []coretestcases.CaseV1{
	{
		Title: "Secure info preserves secure flag through serialization",
		ArrangeInput: args.Map{
			"when": "given secure info round-trip",
			"name": "sec-rt-name",
			"desc": "sec-rt-desc",
			"url":  "sec-rt-url",
		},
		ExpectedInput: args.Map{
			"name":     "sec-rt-name",
			"desc":     "sec-rt-desc",
			"url":      "sec-rt-url",
			"noError":  "true",
			"isSecure": "true",
		},
	},
}

var infoSerializeExamplesTestCases = []coretestcases.CaseV1{
	{
		Title: "Info with examples preserves examples through serialization",
		ArrangeInput: args.Map{
			"when":     "given info with examples round-trip",
			"name":     "ex-rt-name",
			"desc":     "ex-rt-desc",
			"url":      "ex-rt-url",
			"examples": []string{"cmd1 --flag", "cmd2 --other"},
		},
		ExpectedInput: args.Map{
			"name":         "ex-rt-name",
			"noError":      "true",
			"exampleCount": "2",
			"example0":     "cmd1 --flag",
			"example1":     "cmd2 --other",
		},
	},
}

var infoSerializeAllUrlsTestCases = []coretestcases.CaseV1{
	{
		Title: "Info with all URLs preserves through serialization",
		ArrangeInput: args.Map{
			"when":     "given info with all URLs round-trip",
			"name":     "full-name",
			"desc":     "full-desc",
			"url":      "full-url",
			"hintUrl":  "full-hint",
			"errorUrl": "full-error",
		},
		ExpectedInput: args.Map{
			"name":     "full-name",
			"url":      "full-url",
			"hintUrl":  "full-hint",
			"errorUrl": "full-error",
			"noError":  "true",
		},
	},
}

// ==========================================
// Clone
// ==========================================

var infoCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "Clone preserves all fields independently",
		ArrangeInput: args.Map{
			"when":    "given cloned info with mutation",
			"name":    "original",
			"desc":    "original-desc",
			"url":     "original-url",
			"newName": "mutated",
		},
		ExpectedInput: args.Map{
			"originalName": "original",
			"clonedName":   "mutated",
			"clonedDesc":   "original-desc",
		},
	},
}

// ==========================================
// Field checks — populated vs empty
// ==========================================

var infoFieldCheckPopulatedTestCases = []coretestcases.CaseV1{
	{
		Title: "Info with all fields populated -- Has checks return true",
		ArrangeInput: args.Map{
			"when": "given fully populated info",
		},
		ExpectedInput: args.Map{
			"hasRootName":        "true",
			"hasDescription":     "true",
			"hasUrl":             "true",
			"hasHintUrl":         "true",
			"hasErrorUrl":        "true",
			"hasExamples":        "true",
			"hasChainingExample": "true",
		},
	},
}

var infoFieldCheckEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "Empty info — Has checks return false",
		ArrangeInput: args.Map{
			"when": "given empty info",
		},
		ExpectedInput: args.Map{
			"hasRootName":        "false",
			"hasDescription":     "false",
			"hasUrl":             "false",
			"hasHintUrl":         "false",
			"hasErrorUrl":        "false",
			"hasExamples":        "false",
			"hasChainingExample": "false",
		},
	},
}
