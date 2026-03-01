package coreinterface

//goland:noinspection SpellCheckingInspection
type JsonByter interface {
	JsonBytesPtr() (jsonBytes []byte, err error)
	JsonBytes() (jsonBytes []byte, err error)
}
