package chmodhelper

import (
	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core/constants"
)

func ParseRwxInstructionToStringRwx(
	rwxInstruction *chmodins.RwxOwnerGroupOther,
	isIncludeHyphen bool,
) string {
	if rwxInstruction == nil {
		return constants.EmptyString
	}

	var hyphen string

	if isIncludeHyphen {
		hyphen = constants.Hyphen
	}

	compiled := hyphen +
		rwxInstruction.Owner +
		rwxInstruction.Group +
		rwxInstruction.Other

	return compiled
}
