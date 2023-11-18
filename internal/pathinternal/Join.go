package pathinternal

import "path"

func Join(joiningPaths ...string) string {
	v := path.Join(joiningPaths...)

	return path.Clean(v)
}
