package pathinternal

import (
	"path"
	"strings"

	"gitlab.com/auk-go/core/osconsts"
)

func Clean(curPath string) string {
	if len(curPath) == 0 {
		return curPath
	}

	v := path.Clean(curPath)
	v = replaceFix(v)

	if osconsts.IsWindows {
		v = strings.ReplaceAll(v, "/", "\\")
	}

	return v
}
