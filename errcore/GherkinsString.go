package errcore

import (
	"fmt"

	"github.com/alimtvnetwork/core/internal/msgformats"
)

func GherkinsString(
	testCaseIndex int,
	feature,
	given,
	when,
	then any,
) string {
	return fmt.Sprintf(
		msgformats.SimpleGherkinsFormat,
		testCaseIndex,
		feature,
		testCaseIndex,
		given,
		when,
		then,
	)
}
