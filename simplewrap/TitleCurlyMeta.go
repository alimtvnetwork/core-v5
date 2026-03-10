package simplewrap

import (
	"fmt"
	
	"github.com/alimtvnetwork/core/constants"
)

// TitleCurlyMeta
//
// Example :
//   - constants.CurlyTitleMetaWrapFormat
//   - "%v: {%v} (%v)"
func TitleCurlyMeta(
	title,
	value,
	meta any,
) string {
	return fmt.Sprintf(
		constants.CurlyTitleMetaWrapFormat,
		toString(title),
		toString(value),
		meta)
}
