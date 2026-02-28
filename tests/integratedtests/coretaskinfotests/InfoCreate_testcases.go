package coretaskinfotests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================
// Info Creation — Default, Examples
// ==========================================

var infoCreateTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Info.Default creates info with name, desc, url",
		ArrangeInput: args.Map{
			"when": "given default info creation",
			"name": "some name",
			"desc": "some desc",
			"url":  "some url",
		},
		ExpectedInput: []string{
			"some name",
			"some desc",
			"some url",
			"false",
			"true",
		},
	},
	{
		Title: "New.Info.Examples creates info with examples",
		ArrangeInput: args.Map{
			"when":     "given info with examples",
			"name":     "example name",
			"desc":     "example desc",
			"url":      "example url",
			"examples": []string{"ex1", "ex2"},
		},
		ExpectedInput: []string{
			"example name",
			"example desc",
			"example url",
			"false",
			"true",
			"true",
			"2",
		},
	},
	{
		Title: "New.Info.Examples with no examples has empty slice",
		ArrangeInput: args.Map{
			"when":       "given info with no examples",
			"name":       "no-ex name",
			"desc":       "no-ex desc",
			"url":        "no-ex url",
			"noExamples": true,
		},
		ExpectedInput: []string{
			"no-ex name",
			"no-ex desc",
			"no-ex url",
			"false",
			"true",
			"false",
			"0",
		},
	},
}

// ==========================================
// Nil Safety
// ==========================================

var infoNilSafetyTestCases = []coretestcases.CaseV1{
	{
		Title: "Nil info SafeName returns empty",
		ArrangeInput: args.Map{
			"when":   "given nil info",
			"method": "SafeName",
		},
		ExpectedInput: []string{""},
	},
	{
		Title: "Nil info SafeDescription returns empty",
		ArrangeInput: args.Map{
			"when":   "given nil info",
			"method": "SafeDescription",
		},
		ExpectedInput: []string{""},
	},
	{
		Title: "Nil info SafeUrl returns empty",
		ArrangeInput: args.Map{
			"when":   "given nil info",
			"method": "SafeUrl",
		},
		ExpectedInput: []string{""},
	},
	{
		Title: "Nil info SafeHintUrl returns empty",
		ArrangeInput: args.Map{
			"when":   "given nil info",
			"method": "SafeHintUrl",
		},
		ExpectedInput: []string{""},
	},
	{
		Title: "Nil info SafeErrorUrl returns empty",
		ArrangeInput: args.Map{
			"when":   "given nil info",
			"method": "SafeErrorUrl",
		},
		ExpectedInput: []string{""},
	},
	{
		Title: "Nil info SafeExampleUrl returns empty",
		ArrangeInput: args.Map{
			"when":   "given nil info",
			"method": "SafeExampleUrl",
		},
		ExpectedInput: []string{""},
	},
	{
		Title: "Nil info IsNull returns true, IsDefined returns false",
		ArrangeInput: args.Map{
			"when":   "given nil info for null check",
			"method": "NullCheck",
		},
		ExpectedInput: []string{
			"true",
			"false",
		},
	},
	{
		Title: "Nil info IsEmpty returns true, HasAnyItem returns false",
		ArrangeInput: args.Map{
			"when":   "given nil info for empty check",
			"method": "EmptyCheck",
		},
		ExpectedInput: []string{
			"true",
			"false",
		},
	},
	{
		Title: "Nil info ClonePtr returns nil",
		ArrangeInput: args.Map{
			"when":   "given nil info for clone",
			"method": "ClonePtr",
		},
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "Nil info PrettyJsonString returns empty",
		ArrangeInput: args.Map{
			"when":   "given nil info for json",
			"method": "PrettyJsonString",
		},
		ExpectedInput: []string{
			"",
		},
	},
}

// ==========================================
// Secure Mode
// ==========================================

var infoSecureModeTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Info.Secure.Default creates secure info",
		ArrangeInput: args.Map{
			"when": "given secure default creation",
			"name": "secure-name",
			"desc": "secure-desc",
			"url":  "secure-url",
		},
		ExpectedInput: []string{
			"secure-name",
			"secure-desc",
			"secure-url",
			"true",
			"false",
			"true",
		},
	},
	{
		Title: "New.Info.Secure.NameDescUrlExamples has secure flag and examples",
		ArrangeInput: args.Map{
			"when":     "given secure with examples",
			"name":     "sec-ex-name",
			"desc":     "sec-ex-desc",
			"url":      "sec-ex-url",
			"examples": []string{"ex1", "ex2", "ex3"},
		},
		ExpectedInput: []string{
			"sec-ex-name",
			"true",
			"false",
			"true",
			"3",
		},
	},
	{
		Title: "SetSecure on nil returns new secure info",
		ArrangeInput: args.Map{
			"when":      "given nil info with SetSecure",
			"setSecure": true,
		},
		ExpectedInput: []string{
			"true",
			"false",
		},
	},
	{
		Title: "SetSecure on existing info mutates to secure",
		ArrangeInput: args.Map{
			"when":      "given plain info then SetSecure",
			"name":      "was-plain",
			"setSecure": true,
			"existing":  true,
		},
		ExpectedInput: []string{
			"true",
			"false",
			"was-plain",
		},
	},
}

// ==========================================
// Plain Mode
// ==========================================

var infoPlainModeTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Info.Plain.Default creates plain info",
		ArrangeInput: args.Map{
			"when": "given plain default creation",
			"name": "plain-name",
			"desc": "plain-desc",
			"url":  "plain-url",
		},
		ExpectedInput: []string{
			"plain-name",
			"plain-desc",
			"plain-url",
			"false",
			"true",
			"true",
		},
	},
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
		ExpectedInput: []string{
			"all-name",
			"all-desc",
			"all-url",
			"all-hint",
			"all-err",
			"false",
			"true",
			"1",
		},
	},
	{
		Title: "SetPlain on nil returns new plain info",
		ArrangeInput: args.Map{
			"when":     "given nil info with SetPlain",
			"setPlain": true,
		},
		ExpectedInput: []string{
			"false",
			"true",
		},
	},
}

// ==========================================
// JSON Serialization Round-Trip
// ==========================================

var infoSerializeTestCases = []coretestcases.CaseV1{
	{
		Title: "Default info serializes and deserializes correctly",
		ArrangeInput: args.Map{
			"when": "given round-trip serialize/deserialize",
			"name": "round-trip name",
			"desc": "round-trip desc",
			"url":  "round-trip url",
		},
		ExpectedInput: []string{
			"round-trip name",
			"round-trip desc",
			"round-trip url",
			"true",
			"false",
		},
	},
	{
		Title: "Secure info preserves secure flag through serialization",
		ArrangeInput: args.Map{
			"when":     "given secure info round-trip",
			"name":     "sec-rt-name",
			"desc":     "sec-rt-desc",
			"url":      "sec-rt-url",
			"isSecure": true,
		},
		ExpectedInput: []string{
			"sec-rt-name",
			"sec-rt-desc",
			"sec-rt-url",
			"true",
			"true",
		},
	},
	{
		Title: "Info with examples preserves examples through serialization",
		ArrangeInput: args.Map{
			"when":     "given info with examples round-trip",
			"name":     "ex-rt-name",
			"desc":     "ex-rt-desc",
			"url":      "ex-rt-url",
			"examples": []string{"cmd1 --flag", "cmd2 --other"},
		},
		ExpectedInput: []string{
			"ex-rt-name",
			"true",
			"2",
			"cmd1 --flag",
			"cmd2 --other",
		},
	},
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
		ExpectedInput: []string{
			"full-name",
			"full-url",
			"full-hint",
			"full-error",
			"true",
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
			"when":     "given cloned info with mutation",
			"name":     "original",
			"desc":     "original-desc",
			"url":      "original-url",
			"newName":  "mutated",
		},
		ExpectedInput: []string{
			"original",
			"mutated",
			"original-desc",
		},
	},
}

// ==========================================
// Has/IsEmpty field checks
// ==========================================

var infoFieldCheckTestCases = []coretestcases.CaseV1{
	{
		Title: "Info with all fields populated — Has checks return true",
		ArrangeInput: args.Map{
			"when": "given fully populated info",
		},
		ExpectedInput: []string{
			"true",
			"true",
			"true",
			"true",
			"true",
			"true",
			"true",
		},
	},
	{
		Title: "Empty info — Has checks return false, IsEmpty checks return true",
		ArrangeInput: args.Map{
			"when":  "given empty info",
			"empty": true,
		},
		ExpectedInput: []string{
			"false",
			"false",
			"false",
			"false",
			"false",
			"false",
			"false",
		},
	},
}
