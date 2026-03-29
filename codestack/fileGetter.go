package codestack

import (
	"path/filepath"
	"runtime"

	"github.com/alimtvnetwork/core/constants"
)

type fileGetter struct{}

func (it fileGetter) Name(skipStack int) string {
	_, file, _, _ := runtime.Caller(skipStack + defaultInternalSkip)
	_, fileName := filepath.Split(file)

	return fileName
}

func (it fileGetter) Path(skipStack int) string {
	_, file, _, isOkay := runtime.Caller(skipStack + defaultInternalSkip)

	if !isOkay {
		return constants.EmptyString
	}

	return file
}

func (it fileGetter) PathLineSep(skipStack int) (
	filePath string, lineNumber int,
) {
	stack := New.Create(Skip1 + skipStack)
	fileWithLine := stack.FileWithLine()
	filePath = fileWithLine.FullFilePath()
	lineNumber = fileWithLine.LineNumber()

	stack.Dispose()

	return filePath, lineNumber
}

func (it fileGetter) PathLineSepDefault() (filePath string, lineNumber int) {
	return it.PathLineSep(defaultInternalSkip)
}

func (it fileGetter) FilePathWithLineString(skipStack int) string {
	stack := New.Create(Skip1 + skipStack)
	fileWithLine := stack.FileWithLineString()
	stack.Dispose()

	return fileWithLine
}

func (it fileGetter) PathLineStringDefault() string {
	stack := New.Create(Skip1)
	fileWithLine := stack.FileWithLineString()
	stack.Dispose()

	return fileWithLine
}

func (it fileGetter) CurrentFilePath() string {
	_, filePath, _, isOkay := runtime.Caller(defaultInternalSkip)

	if !isOkay {
		return constants.EmptyString
	}

	return filePath
}
