package enumimpl

import (
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/internal/csvinternal"
)

func OnlySupportedErr(
	skipStack int,
	allNames []string,
	supportedNames ...string,
) error {
	if len(allNames) == 0 {
		return nil
	}

	unsupportedNames := UnsupportedNames(
		allNames,
		supportedNames...,
	)

	if len(unsupportedNames) == 0 {
		return nil
	}

	supportedCsv := csvinternal.StringsToStringDefault(
		supportedNames...,
	)

	unsupportedCsv := csvinternal.StringsToStringDefault(
		unsupportedNames...,
	)

	supportedMsg := "Only supported (" + supportedCsv + ")"
	unsupportedMsg := ". Unsupported (" + unsupportedCsv + ")"

	return errcore.
		RangesOnlySupportedType.
		ErrorNoRefsSkip(
			skipStack,
			supportedMsg+unsupportedMsg,
		)
}
