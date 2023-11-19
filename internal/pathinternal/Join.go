package pathinternal

import (
	"path"
	"strings"

	"gitlab.com/auk-go/core/osconsts"
)

func Join(joiningPaths ...string) string {
	v := path.Join(joiningPaths...)

	return Clean(v)
}

func Clean(curPath string) string {
	if len(curPath) == 0 {
		return curPath
	}

	v := path.Clean(curPath)

	if osconsts.IsWindows {
		return strings.
			ReplaceAll(
				v,
				"\\\\",
				"\\",
			)
	}
}
