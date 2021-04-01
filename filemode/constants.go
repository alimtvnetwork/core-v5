package filemode

// Reference : https://ss64.com/bash/chmod.html
const (
	supportedLengthString      = "3"
	supportedLength            = 3
	readValue                  = 4
	writeValue                 = 2
	executeValue               = 1
	readWriteValue             = readValue + writeValue
	readExecuteValue           = readValue + executeValue
	writeExecuteValue          = writeValue + executeValue
	readWriteExecuteValue      = readValue + writeValue + executeValue
	ownerIndex                 = 0
	groupIndex                 = 1
	otherIndex                 = 2
	readChar              byte = 'r'
	writeChar             byte = 'w'
	executeChar           byte = 'x'
)
