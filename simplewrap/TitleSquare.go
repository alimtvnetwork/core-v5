package simplewrap

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

func TitleSquare(
	title, value any,
) string {
	return fmt.Sprintf(
		constants.SquareTitleWrapFormat,
		toString(title),
		toString(value))
}
