package enumimpl

func UnsupportedNames(
	allNames []string,
	supportedNames ...string,
) []string {
	unsupportedNames := make(
		[]string,
		0,
		len(allNames)-len(supportedNames)+1,
	)

	supportedNamesHashset := toHashset(supportedNames...)

	for _, name := range allNames {
		if !supportedNamesHashset[name] {
			unsupportedNames = append(
				unsupportedNames,
				name,
			)
		}
	}

	return unsupportedNames
}
