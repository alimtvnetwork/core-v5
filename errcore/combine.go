package errcore

import (
	"github.com/alimtvnetwork/core/constants"
)

func Combine(
	genericMsg,
	otherMsg,
	reference string,
) string {
	return genericMsg +
		constants.Space +
		otherMsg +
		constants.Space +
		ReferenceStart +
		reference +
		ReferenceEnd
}
