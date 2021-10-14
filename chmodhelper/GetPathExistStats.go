package chmodhelper

func GetPathExistStats(
	isCollectValidOnly bool,
	locations ...string,
) []*PathExistStat {
	if len(locations) == 0 {
		return []*PathExistStat{}
	}

	if isCollectValidOnly {
		slice := make([]*PathExistStat, 0, len(locations))

		for _, location := range locations {
			stat := GetPathExistStat(location)

			if stat.IsInvalid() {
				continue
			}

			slice = append(slice, stat)
		}

		return slice
	}

	allSlice := make([]*PathExistStat, len(locations))

	for i, location := range locations {
		stat := GetPathExistStat(location)
		allSlice[i] = stat
	}

	return allSlice
}
