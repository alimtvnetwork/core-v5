package stringutil

import (
	"strings"

	"github.com/alimtvnetwork/core/constants"
)

func RemoveMany(
	content string,
	removeRequests ...string,
) string {
	if content == "" {
		return content
	}

	for _, remove := range removeRequests {
		content = strings.ReplaceAll(
			content,
			remove,
			constants.EmptyString)
	}

	return content
}
