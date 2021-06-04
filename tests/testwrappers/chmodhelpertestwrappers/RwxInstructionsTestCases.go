package chmodhelpertestwrappers

import "gitlab.com/evatix-go/core/chmodhelper/chmodins"

var defaultRwx = chmodins.RwxOwnerGroupOther{
	Owner: "rwx",
	Group: "r-x",
	Other: "r--",
}

// RwxInstructionsApplyTestCases https://ss64.com/bash/chmod.html
var RwxInstructionsApplyTestCases = []RwxInstructionTestWrapper{
	{
		RwxInstructions: []*chmodins.RwxInstruction{
			{
				IsSkipOnNonExist:  false,
				IsContinueOnError: false,
				IsRecursive:       false,
				RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
					Owner: "*-x",
					Group: "**x",
					Other: "-w-",
				},
			},
		},
		DefaultRwx:      &defaultRwx,
		IsErrorExpected: false,
		CreatePaths: []*CreatePathsInstruction{
			{
				Dir: "/temp/core/test-cases",
				Files: []string{
					"file-1.txt",
					"file-2.txt",
					"file-3.txt",
				},
				ApplyRwx: defaultRwx,
			},
			{
				Dir: "/temp/core/test-cases-2",
				Files: []string{
					"file-1.txt",
					"file-2.txt",
					"file-3.txt",
				},
				ApplyRwx: defaultRwx,
			},
		},
		funcName: RwxApplyOnPath,
		expected: chmodins.RwxOwnerGroupOther{
			Owner: "r-x",
			Group: "r-x",
			Other: "-w-",
		},
	},
}

// RwxInstructionsUnixApplyRecursivelyTestCases https://ss64.com/bash/chmod.html
var RwxInstructionsUnixApplyRecursivelyTestCases = []RwxInstructionTestWrapper{
	{
		RwxInstructions: []*chmodins.RwxInstruction{
			{
				IsSkipOnNonExist:  false,
				IsContinueOnError: false,
				IsRecursive:       true,
				RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
					Owner: "*-x",
					Group: "**x",
					Other: "-w-",
				},
			},
		},
		DefaultRwx:      &defaultRwx,
		IsErrorExpected: false,
		CreatePaths: []*CreatePathsInstruction{
			{
				Dir: "/temp/core/test-cases",
				Files: []string{
					"file-1.txt",
					"file-2.txt",
					"file-3.txt",
				},
				ApplyRwx: defaultRwx,
			},
			{
				Dir: "/temp/core/test-cases-2",
				Files: []string{
					"file-1.txt",
					"file-2.txt",
					"file-3.txt",
				},
				ApplyRwx: defaultRwx,
			},
			{
				Dir: "/temp/core/test-cases-3",
				Files: []string{
					"file-1.txt",
					"file-2.txt",
					"file-3.txt",
				},
				ApplyRwx: defaultRwx,
			},
		},
		funcName: RwxApplyOnPath,
		expected: chmodins.RwxOwnerGroupOther{
			Owner: "r-x",
			Group: "r-x",
			Other: "-w-",
		},
	},
}
