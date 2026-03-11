package simplewrap

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

func TitleCurlyWrap(
	title, value any,
) string {
	return fmt.Sprintf(
		constants.CurlyTitleWrapFormat,
		toString(title),
		toString(value))
}
