package corestr

import (
	"strings"

	"gitlab.com/evatix-go/core/internal/stringutil"
)

type ValidValue struct {
	Value          string
	valueBytes     *[]byte
	IsValid        bool
	InvalidMessage string
}

func (receiver *ValidValue) ValueBytesOnce() []byte {
	return *receiver.ValueBytesOncePtr()
}

func (receiver *ValidValue) ValueBytesOncePtr() *[]byte {
	if receiver.valueBytes == nil {
		valueBytes := []byte(receiver.Value)

		receiver.valueBytes = &valueBytes
	}

	return receiver.valueBytes
}

func (receiver *ValidValue) IsEmpty() bool {
	return receiver.Value == ""
}

func (receiver *ValidValue) IsWhitespace() bool {
	return stringutil.IsEmptyOrWhitespace(receiver.Value)
}

func (receiver *ValidValue) HasValidNonEmpty() bool {
	return receiver.IsValid && !receiver.IsEmpty()
}

func (receiver *ValidValue) HasValidNonWhitespace() bool {
	return receiver.IsValid && !receiver.IsWhitespace()
}

// HasSafeNonEmpty receiver.IsValid &&
//		!receiver.IsLeftEmpty() &&
//		!receiver.IsMiddleEmpty() &&
//		!receiver.IsRightEmpty()
func (receiver *ValidValue) HasSafeNonEmpty() bool {
	return receiver.IsValid &&
		!receiver.IsEmpty()
}

func (receiver *ValidValue) Is(val string) bool {
	return receiver.Value == val
}

// IsAnyOf if length of values are 0 then returns true
func (receiver *ValidValue) IsAnyOf(values ...string) bool {
	if len(values) == 0 {
		return true
	}

	for _, value := range values {
		if receiver.Value == value {
			return true
		}
	}

	return false
}

// IsAllOf if length of values are 0 then returns true
func (receiver *ValidValue) IsAllOf(values ...string) bool {
	if len(values) == 0 {
		return true
	}

	for _, value := range values {
		if receiver.Value != value {
			return false
		}
	}

	return true
}

func (receiver *ValidValue) IsContains(val string) bool {
	return strings.Contains(receiver.Value, val)
}

// IsAnyContains if length of values are 0 then returns true
func (receiver *ValidValue) IsAnyContains(values ...string) bool {
	if len(values) == 0 {
		return true
	}

	for _, value := range values {
		if receiver.IsContains(value) {
			return true
		}
	}

	return false
}

func (receiver *ValidValue) IsAllContains(values ...string) bool {
	for _, value := range values {
		if !receiver.IsContains(value) {
			return false
		}
	}

	return true
}

func (receiver *ValidValue) IsEqualNonSensitive(val string) bool {
	return strings.EqualFold(receiver.Value, val)
}
