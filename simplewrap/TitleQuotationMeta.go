package simplewrap

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

// TitleQuotationMeta
//
// Example :
//   - constants.QuotationTitleMetaWrapFormat
//   - "%v: \"%v\" (%v)"
func TitleQuotationMeta(
	title,
	value,
	meta any,
) string {
	return fmt.Sprintf(
		constants.QuotationTitleMetaWrapFormat,
		toString(title),
		toString(value),
		meta)
}
