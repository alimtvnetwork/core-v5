package stringslice

import "strings"

func JoinWith(joiner string, items ...string) string {
	if len(items) == 0 {
		return ""
	}

	return joiner + strings.Join(items, joiner)
}
