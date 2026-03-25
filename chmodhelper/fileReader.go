package chmodhelper

import (
	"fmt"
	"os"
)

type fileReader struct{}

func (it fileReader) Read(filePath string) (string, error) {
	b, err := it.ReadBytes(filePath)

	return string(b), err
}

func (it fileReader) ReadBytes(filePath string) ([]byte, error) {
	b, err := os.ReadFile(filePath)

	if err != nil {
		return []byte{}, fmt.Errorf(
			"%s, path : %s",
			err,
			filePath,
		)
	}

	return b, err
}
