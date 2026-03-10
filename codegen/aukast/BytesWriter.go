package aukast

import (
	"io"

	"github.com/alimtvnetwork/core/internal/jsoninternal"
)

type BytesWriter struct {
	rawBytes []byte
}

func (it *BytesWriter) IsEmpty() bool {
	return it == nil || len(it.rawBytes) == 0
}

func (it *BytesWriter) HasAny() bool {
	return it != nil && len(it.rawBytes) > 0
}

func (it *BytesWriter) Write(p []byte) (n int, err error) {
	it.rawBytes = append(it.rawBytes, p...)

	return len(p), nil
}

func (it *BytesWriter) Bytes() []byte {
	return it.rawBytes
}

func (it *BytesWriter) String() string {
	if it.IsEmpty() {
		return ""
	}

	return string(it.rawBytes)
}

func (it *BytesWriter) Pretty() string {
	if it.IsEmpty() {
		return ""
	}

	return jsoninternal.
		Pretty.
		Bytes.
		SafeDefault(it.rawBytes)
}

func (it BytesWriter) AsWriter() io.Writer {
	return &it
}
