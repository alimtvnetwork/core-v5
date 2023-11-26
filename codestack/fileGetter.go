package codestack

import (
	"path/filepath"
	"runtime"

	"gitlab.com/auk-go/core/constants"
)

type fileGetter struct{}

func FileName(skipStack int) string {
	_, file, _, isOkay := runtime.Caller(skipStack + defaultInternalSkip)

	if !isOkay && file == "" {
		return constants.EmptyString
	}

	_, fileName := filepath.Split(file)

	return fileName
}

func FilePath(skipStack int) string {
	_, file, _, isOkay := runtime.Caller(skipStack + defaultInternalSkip)

	if isOkay {
		return file
	}

	return constants.EmptyString
}

func FilePathWithLineSeparate(
	skipStack int,
) (
	filePath string, lineNumber int,
) {
	stack := New(Skip1 + skipStack)
	fileWithLine := stack.FileWithLine()
	filePath = fileWithLine.FullFilePath()
	lineNumber = fileWithLine.LineNumber()

	stack.Dispose()

	return filePath, lineNumber
}

func FilePathWithLineSeparateDefault() (
	filePath string, lineNumber int,
) {
	return FilePathWithLineSeparate(defaultInternalSkip)
}

func FilePathWithLineString(
	skipStack int,
) string {
	stack := New(Skip1 + skipStack)
	fileWithLine := stack.FileWithLineString()
	stack.Dispose()

	return fileWithLine
}

func FilePathWithLineStringDefault() string {
	stack := New(Skip1)
	fileWithLine := stack.FileWithLineString()
	stack.Dispose()

	return fileWithLine
}

func FullMethodNameOf(fullName string) (packageName string) {
	fullMethodNameOf, _, _ := MethodNamePackageName(
		fullName,
	)

	return fullMethodNameOf
}

func CurrentFilePath() string {
	_, filePath, _, isOkay := runtime.Caller(defaultInternalSkip)

	if isOkay {
		return filePath
	}

	return constants.EmptyString
}
func CurDir() string {
	_, filePath, _, isOkay := runtime.Caller(defaultInternalSkip)

	if isOkay {
		return filepath.Dir(filePath)
	}

	return constants.EmptyString
}

func Dir(skipStack int) string {
	_, filePath, _, isOkay := runtime.Caller(skipStack + defaultInternalSkip)

	if isOkay {
		return filepath.Dir(filePath)
	}

	return constants.EmptyString
}
