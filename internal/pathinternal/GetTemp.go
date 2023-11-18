package pathinternal

import "os"

func GetTemp() string {
	return os.TempDir()
}
