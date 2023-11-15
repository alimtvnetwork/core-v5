package coreversion

func InvalidCompactVersion(compactVersion string) *Version {
	return &Version{
		VersionCompact:  compactVersion,
		compiledVersion: compactVersion,
		isInvalid:       true,
		VersionMajor:    InvalidVersionValue,
		VersionMinor:    InvalidVersionValue,
		VersionPatch:    InvalidVersionValue,
		VersionBuild:    InvalidVersionValue,
	}
}
