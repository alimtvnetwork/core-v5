package simplewrap

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

// TitleSquareMeta
//
// Example :
//   - constants.SquareTitleMetaWrapFormat
//   - "%v: [%v] (%v)"
func TitleSquareMeta(
	title,
	value,
	meta any,
) string {
	return fmt.Sprintf(
		constants.SquareTitleMetaWrapFormat,
		toString(title),
		toString(value),
		meta)
}
