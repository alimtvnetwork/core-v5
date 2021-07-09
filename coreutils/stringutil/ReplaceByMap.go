package stringutil

import "strings"

func ReplaceByMap(
	contentTemplate string,
	compilingMap map[string]string,
) string {
	if len(compilingMap) == 0 || len(contentTemplate) == 0 {
		return contentTemplate
	}

	for key, value := range compilingMap {
		contentTemplate = strings.ReplaceAll(
			contentTemplate,
			key,
			value)
	}

	return contentTemplate
}
