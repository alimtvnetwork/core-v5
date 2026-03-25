package corevalidatortestwrappers

import (
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/corevalidator"
)

type SliceValidatorWrapper struct {
	Case      coretestcases.CaseV1
	Validator corevalidator.SliceValidator
}
