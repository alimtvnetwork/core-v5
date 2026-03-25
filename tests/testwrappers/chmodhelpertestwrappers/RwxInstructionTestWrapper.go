package chmodhelpertestwrappers

import (
	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core/coretests"
)

type RwxInstructionTestWrapper struct {
	RwxInstructions []chmodins.RwxInstruction
	DefaultRwx      *chmodins.RwxOwnerGroupOther
	IsErrorExpected bool
	CreatePaths     []chmodhelper.DirFilesWithRwxPermission
	TestFuncName    coretests.TestFuncName
	WhatIsExpected  chmodins.RwxOwnerGroupOther
	actual          any
}

func (it *RwxInstructionTestWrapper) Actual() any {
	return it.actual
}

func (it *RwxInstructionTestWrapper) SetActual(actual any) {
	it.actual = actual
}

func (it *RwxInstructionTestWrapper) GetFuncName() string {
	return it.TestFuncName.Value()
}

func (it *RwxInstructionTestWrapper) Value() any {
	return it
}

func (it *RwxInstructionTestWrapper) Expected() any {
	return it.WhatIsExpected
}

func (it *RwxInstructionTestWrapper) ExpectedAsRwxOwnerGroupOtherInstruction() chmodins.RwxOwnerGroupOther {
	return it.WhatIsExpected
}

func (it *RwxInstructionTestWrapper) AsTestCaseMessenger() coretests.TestCaseMessenger {
	var testCaseMessenger coretests.TestCaseMessenger = it

	return testCaseMessenger
}
