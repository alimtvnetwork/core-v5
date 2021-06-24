package chmodhelpertestwrappers

import "gitlab.com/evatix-go/core/chmodhelper/chmodins"

type VerifyRwxChmodUsingRwxInstructionsWrapper struct {
	Header string
	chmodins.RwxInstruction
	Locations            []string
	ExpectedErrorMessage string
}
