package converterstests

import (
	"testing"

	"github.com/alimtvnetwork/core/converters"
)

func Test_QW_ToStringsUsingProcessor_NilInput(t *testing.T) {
	result := converters.AnyTo.ToStringsUsingProcessor(
		false,
		func(index int, in any) (string, bool, bool) { return "", true, false },
		nil,
	)
	if len(result) != 0 {
		t.Fatal("expected empty slice for nil input")
	}
}
