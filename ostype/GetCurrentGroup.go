package ostype

import "github.com/alimtvnetwork/core/osconsts"

func GetCurrentGroup() Group {
	return GetGroup(osconsts.CurrentOperatingSystem)
}
