package coretests

import (
	"fmt"

	"gitlab.com/evatix-go/core/internal/msgformats"
)

func GetAssertMessageQuick(
	when,
	actual,
	expected interface{},
	counter int,
) string {
	return fmt.Sprintf(msgformats.QuickActualExpectedMessageFormat,
		counter,
		when,
		actual,
		expected,
	)
}
