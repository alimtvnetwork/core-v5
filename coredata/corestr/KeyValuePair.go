package corestr

import (
	"fmt"
	"strings"
)

type KeyValuePair struct {
	Key, Value string
}

func NewKeyValuePairTrimmed(key, val string) *KeyValuePair {
	return &KeyValuePair{
		Key:   strings.TrimSpace(key),
		Value: strings.TrimSpace(val),
	}
}

func (it *KeyValuePair) IsKeyEmpty() bool {
	return it.Key == ""
}

func (it *KeyValuePair) IsValueEmpty() bool {
	return it.Value == ""
}

func (it *KeyValuePair) HasKey() bool {
	return it.Key != ""
}

func (it *KeyValuePair) HasValue() bool {
	return it.Value != ""
}

func (it *KeyValuePair) IsKeyValueEmpty() bool {
	return it.Key == "" && it.Value == ""
}

func (it *KeyValuePair) TrimKey() string {
	return strings.TrimSpace(it.Key)
}

func (it *KeyValuePair) TrimValue() string {
	return strings.TrimSpace(it.Value)
}

func (it *KeyValuePair) Is(key, val string) bool {
	return it.Key == key && it.Value == val
}

func (it *KeyValuePair) IsKey(key string) bool {
	return it.Key == key
}

func (it *KeyValuePair) IsVal(val string) bool {
	return it.Value == val
}

func (it *KeyValuePair) IsKeyValueAnyEmpty() bool {
	return it.Key == "" || it.Value == ""
}

func (it *KeyValuePair) FormatString(format string) string {
	return fmt.Sprintf(
		format,
		it.Key,
		it.Value)
}

func (it *KeyValuePair) String() string {
	return fmt.Sprintf(
		keyValuePrintFormat,
		it.Key,
		it.Value)
}
