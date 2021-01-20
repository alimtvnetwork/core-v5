package corestr

const (
	emptyChar                              byte = 0
	eachItemDefaultCapacity                     = 100
	defaultHashsetItems                         = 10
	RegularCollectionEfficiencyLimit            = 1000
	DoubleLimit                                 = RegularCollectionEfficiencyLimit * 3
	commonJoiner                                = "\n\t- "
	NoElements                                  = " {No Elements}"
	charCollectionMapLengthFormat               = "\n## Items of `%s`"
	charCollectionMapSingleItemFormat           = "\n\t- `%s` has `%d` items."
	summaryOfCharCollectionMapLengthFormat      = "# Summary of `%T`, Length (\"%d\")"
)
