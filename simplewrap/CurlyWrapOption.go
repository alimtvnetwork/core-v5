package simplewrap

import (
	"fmt"
	
	"gitlab.com/auk-go/core/constants"
)

func CurlyWrapOption(
	isSkipIfExists bool,
	source interface{},
) string {
	toStr := fmt.Sprintf(
		constants.CurlyWrapFormat,
		source)
	
	return ConditionalWrapWith(
		'{',
		toStr,
		'}')
}
