package defaultcapacity

import (
	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/constants/percentconst"
)

func PredictiveDefaultSmall(possibleLen int) int {
	return Predictive(
		possibleLen,
		percentconst.OnePointTwoPercentIncrement,
		constants.Capacity4)
}
