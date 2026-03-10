package chmodhelpertestwrappers

import "github.com/alimtvnetwork/core/chmodhelper/chmodins"

type VerifyRwxChmodUsingRwxInstructionsWrapper struct {
	Header string
	chmodins.RwxInstruction
	Locations            []string
	ExpectedErrorMessage string
}
