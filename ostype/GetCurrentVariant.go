package ostype

import "github.com/alimtvnetwork/core/osconsts"

func GetCurrentVariant() Variation {
	return GetVariant(osconsts.CurrentOperatingSystem)
}
