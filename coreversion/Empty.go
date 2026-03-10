package coreversion

import "github.com/alimtvnetwork/core/constants"

func Empty() Version {
	return Version{
		VersionCompact: constants.EmptyString,
		Compiled:       "",
		IsInvalid:      true,
		VersionMajor:   InvalidVersionValue,
		VersionMinor:   InvalidVersionValue,
		VersionPatch:   InvalidVersionValue,
		VersionBuild:   InvalidVersionValue,
	}
}
