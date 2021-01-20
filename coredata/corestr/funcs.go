package corestr

type OnComplete func(stringsMap *CharCollectionMap)
type IsStringFilter func(str string) (result string, isKeep bool)
type IsStringPointerFilter func(stringPointer *string) (result *string, isKeep bool)
