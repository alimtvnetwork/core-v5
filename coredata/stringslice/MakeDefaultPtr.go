package stringslice

import "gitlab.com/evatix-go/core/constants"

func MakeDefaultPtr(capacity int) *[]string {
	slice := make([]string, constants.Zero, capacity)

	return &slice
}
