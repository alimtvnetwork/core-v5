package chmodhelpertests

import (
	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ── RwxWrapper creation test cases ──

var ext2RwxWrapperCreateTestCases = []coretestcases.CaseV1{
	{
		Title: "Create RwxWrapper using mode 755",
		ArrangeInput: args.Map{
			"mode": "755",
		},
		ExpectedInput: args.Map{
			"rwxFull":   "-rwxr-xr-x",
			"fileMode":  "0755",
			"rwx3":      "755",
			"hasError":  false,
			"isDefined": true,
		},
	},
	{
		Title: "Create RwxWrapper using mode 644",
		ArrangeInput: args.Map{
			"mode": "644",
		},
		ExpectedInput: args.Map{
			"rwxFull":   "-rw-r--r--",
			"fileMode":  "0644",
			"rwx3":      "644",
			"hasError":  false,
			"isDefined": true,
		},
	},
	{
		Title: "Create RwxWrapper using mode 777",
		ArrangeInput: args.Map{
			"mode": "777",
		},
		ExpectedInput: args.Map{
			"rwxFull":   "-rwxrwxrwx",
			"fileMode":  "0777",
			"rwx3":      "777",
			"hasError":  false,
			"isDefined": true,
		},
	},
}

// ── RwxFullString parse test cases ──

var ext2RwxFullStringParseTestCases = []coretestcases.CaseV1{
	{
		Title: "Parse -rwxr-xr-x",
		ArrangeInput: args.Map{
			"rwxFull": "-rwxr-xr-x",
		},
		ExpectedInput: args.Map{
			"hasError":  false,
			"rwx3":      "755",
			"isDefined": true,
		},
	},
	{
		Title: "Parse invalid length string",
		ArrangeInput: args.Map{
			"rwxFull": "rwx",
		},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
}

// ── Attribute creation test cases ──

var ext2AttributeTestCases = []coretestcases.CaseV1{
	{
		Title: "Attribute rwx -- all true",
		ArrangeInput: args.Map{
			"rwx": "rwx",
		},
		ExpectedInput: args.Map{
			"isRead":    true,
			"isWrite":   true,
			"isExecute": true,
			"toByte":    byte(7),
			"rwxStr":    "rwx",
			"isEmpty":   false,
		},
	},
	{
		Title: "Attribute r-- -- read only",
		ArrangeInput: args.Map{
			"rwx": "r--",
		},
		ExpectedInput: args.Map{
			"isRead":    true,
			"isWrite":   false,
			"isExecute": false,
			"toByte":    byte(4),
			"rwxStr":    "r--",
			"isEmpty":   false,
		},
	},
	{
		Title: "Attribute --- -- all false",
		ArrangeInput: args.Map{
			"rwx": "---",
		},
		ExpectedInput: args.Map{
			"isRead":    false,
			"isWrite":   false,
			"isExecute": false,
			"toByte":    byte(0),
			"rwxStr":    "---",
			"isEmpty":   true,
		},
	},
}

// ── Variant test cases ──

var ext2VariantTestCases = []coretestcases.CaseV1{
	{
		Title: "Variant X755 creates rwxr-xr-x",
		ArrangeInput: args.Map{
			"variant": chmodhelper.X755,
		},
		ExpectedInput: args.Map{
			"rwxFull":  "-rwxr-xr-x",
			"hasError": false,
		},
	},
	{
		Title: "Variant X644 creates rw-r--r--",
		ArrangeInput: args.Map{
			"variant": chmodhelper.X644,
		},
		ExpectedInput: args.Map{
			"rwxFull":  "-rw-r--r--",
			"hasError": false,
		},
	},
}

// ── AttrVariant test cases ──

var ext2AttrVariantTestCases = []coretestcases.CaseV1{
	{
		Title: "AttrVariant ReadWriteExecute = 7",
		ArrangeInput: args.Map{
			"variant": chmodhelper.ReadWriteExecute,
		},
		ExpectedInput: args.Map{
			"value":     byte(7),
			"isGreater": false, // 5 > 7 is false
		},
	},
	{
		Title: "AttrVariant Read = 4",
		ArrangeInput: args.Map{
			"variant": chmodhelper.Read,
		},
		ExpectedInput: args.Map{
			"value":     byte(4),
			"isGreater": true, // 5 > 4 is true
		},
	},
}

// ── ParseRwxToVarAttribute test cases ──

var ext2ParseRwxToVarAttrTestCases = []coretestcases.CaseV1{
	{
		Title: "Parse fixed rwx",
		ArrangeInput: args.Map{
			"rwx": "rwx",
		},
		ExpectedInput: args.Map{
			"hasError":    false,
			"isFixedType": true,
		},
	},
	{
		Title: "Parse wildcard r*x",
		ArrangeInput: args.Map{
			"rwx": "r*x",
		},
		ExpectedInput: args.Map{
			"hasError":    false,
			"isFixedType": false,
		},
	},
	{
		Title: "Parse invalid length rw",
		ArrangeInput: args.Map{
			"rwx": "rw",
		},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
}

// ── MergeRwxWildcardWithFixedRwx test cases ──

var ext2MergeRwxTestCases = []coretestcases.CaseV1{
	{
		Title: "Merge r-x with r*- gives r--",
		ArrangeInput: args.Map{
			"existing": "r-x",
			"wildcard": "r*-",
		},
		ExpectedInput: args.Map{
			"hasError": false,
			"result":   "r--",
		},
	},
	{
		Title: "Merge rwx with *** gives rwx",
		ArrangeInput: args.Map{
			"existing": "rwx",
			"wildcard": "***",
		},
		ExpectedInput: args.Map{
			"hasError": false,
			"result":   "rwx",
		},
	},
	{
		Title: "Invalid existing length",
		ArrangeInput: args.Map{
			"existing": "rw",
			"wildcard": "rwx",
		},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
	{
		Title: "Invalid wildcard length",
		ArrangeInput: args.Map{
			"existing": "rwx",
			"wildcard": "rw",
		},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
}

// ── NewRwxVariableWrapper test cases ──

var ext2NewRwxVarWrapperTestCases = []coretestcases.CaseV1{
	{
		Title: "Full partial -rwx creates fixed",
		ArrangeInput: args.Map{
			"partial": "-rwxr-xr--",
		},
		ExpectedInput: args.Map{
			"hasError":    false,
			"isFixedType": true,
		},
	},
	{
		Title: "Short partial -rwx pads wildcards",
		ArrangeInput: args.Map{
			"partial": "-rwx",
		},
		ExpectedInput: args.Map{
			"hasError":    false,
			"isFixedType": false,
		},
	},
}

// ── SingleRwx test cases ──

var ext2SingleRwxAllTestCase = coretestcases.CaseV1{
	Title: "SingleRwx All classtype rwx",
	ArrangeInput: args.Map{
		"rwx":       "rwx",
		"classType": "all",
	},
	ExpectedInput: args.Map{
		"owner": "rwx",
		"group": "rwx",
		"other": "rwx",
	},
}

// ── ParseRwxInstructionToStringRwx test cases ──

var ext2ParseRwxToStringTestCases = []coretestcases.CaseV1{
	{
		Title: "With hyphen",
		ArrangeInput: args.Map{
			"rwx": &chmodins.RwxOwnerGroupOther{
				Owner: "rwx",
				Group: "r-x",
				Other: "r--",
			},
			"includeHyphen": true,
		},
		ExpectedInput: args.Map{
			"result": "-rwxr-xr--",
		},
	},
	{
		Title: "Without hyphen",
		ArrangeInput: args.Map{
			"rwx": &chmodins.RwxOwnerGroupOther{
				Owner: "rwx",
				Group: "r-x",
				Other: "r--",
			},
			"includeHyphen": false,
		},
		ExpectedInput: args.Map{
			"result": "rwxr-xr--",
		},
	},
	{
		Title: "Nil input",
		ArrangeInput: args.Map{
			"rwx":           (*chmodins.RwxOwnerGroupOther)(nil),
			"includeHyphen": true,
		},
		ExpectedInput: args.Map{
			"result": "",
		},
	},
}

// ── GetRwxLengthError test cases ──

var ext2GetRwxLengthErrTestCases = []coretestcases.CaseV1{
	{
		Title: "Valid rwx length 3",
		ArrangeInput: args.Map{
			"rwx": "rwx",
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
	{
		Title: "Invalid rwx length 2",
		ArrangeInput: args.Map{
			"rwx": "rw",
		},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
	{
		Title: "Invalid rwx length 4",
		ArrangeInput: args.Map{
			"rwx": "rwxr",
		},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
}
