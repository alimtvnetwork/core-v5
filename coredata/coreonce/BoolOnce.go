package coreonce

import (
	"encoding/json"
	"strconv"
	"sync"
)

type BoolOnce struct {
	innerData       bool
	initializerFunc func() bool
	once            sync.Once
}

func NewBoolOnce(initializerFunc func() bool) BoolOnce {
	return BoolOnce{
		initializerFunc: initializerFunc,
	}
}

func NewBoolOncePtr(initializerFunc func() bool) *BoolOnce {
	return &BoolOnce{
		initializerFunc: initializerFunc,
	}
}

func (it *BoolOnce) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.Value())
}

func (it *BoolOnce) UnmarshalJSON(data []byte) error {
	it.once.Do(func() {})
	return json.Unmarshal(data, &it.innerData)
}

func (it *BoolOnce) Value() bool {
	it.once.Do(func() {
		it.innerData = it.initializerFunc()
	})

	return it.innerData
}

func (it *BoolOnce) Execute() bool {
	return it.Value()
}

func (it *BoolOnce) String() string {
	return strconv.FormatBool(it.Value())
}

func (it *BoolOnce) Serialize() ([]byte, error) {
	value := it.Value()

	return json.Marshal(value)
}
