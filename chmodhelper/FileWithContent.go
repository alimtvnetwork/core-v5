package chmodhelper

import (
	"os"
	"strings"

	"gitlab.com/auk-go/core/constants"
)

type FileWithContent struct {
	RelativePath string
	FileMode     os.FileMode // default for file fileDefaultChmod
	Content      []string
}

func (it FileWithContent) ContentToString() string {
	return strings.Join(it.Content, constants.NewLineUnix)
}

func (it FileWithContent) ContentToBytes() []byte {
	return []byte(it.ContentToString())
}
