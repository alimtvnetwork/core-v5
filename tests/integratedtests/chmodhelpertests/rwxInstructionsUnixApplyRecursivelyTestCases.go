package chmodhelpertests

import (
	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core/tests/testwrappers/chmodhelpertestwrappers"
)

// rwxInstructionsUnixApplyRecursivelyTestCase https://ss64.com/bash/chmod.html
var rwxInstructionsUnixApplyRecursivelyTestCase = chmodhelpertestwrappers.RwxInstructionTestWrapper{
	RwxInstructions: []chmodins.RwxInstruction{
		{
			Condition: chmodins.Condition{
				IsSkipOnInvalid:   false,
				IsContinueOnError: false,
				IsRecursive:       true,
			},
			RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
				Owner: "*-x",
				Group: "**x",
				Other: "-w-",
			},
		},
	},
	DefaultRwx:      &chmodhelpertestwrappers.DefaultRwx,
	IsErrorExpected: false,
	CreatePaths:     pathInstructionsV2,
	TestFuncName:    chmodhelpertestwrappers.RwxApplyOnPath,
	WhatIsExpected:  chmodhelpertestwrappers.DefaultExpected,
}
