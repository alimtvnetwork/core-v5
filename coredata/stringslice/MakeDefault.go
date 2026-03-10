package stringslice

import "github.com/alimtvnetwork/core/constants"

func MakeDefault(capacity int) []string {
	return make([]string, constants.Zero, capacity)
}
