package simplewrap

import (
	"fmt"
	
	"github.com/alimtvnetwork/core/constants"
)

func TitleSquareMetaUsingFmt(
	title,
	value,
	meta fmt.Stringer,
) string {
	return fmt.Sprintf(
		constants.SquareTitleMetaWrapFormat,
		toString(title),
		toString(value),
		meta)
}
