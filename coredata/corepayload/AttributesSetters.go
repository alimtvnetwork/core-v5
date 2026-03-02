package corepayload

import (
	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coreinterface/errcoreinf"
	"gitlab.com/auk-go/core/coreinterface/payloadinf"
	"gitlab.com/auk-go/core/errcore"
)

// AttributesSetters.go — Mutating methods extracted from Attributes.go

func (it *Attributes) HandleErr() {
	if it.HasError() {
		it.BasicErrWrapper.HandleError()
	}
}

func (it *Attributes) HandleError() {
	if it.HasError() {
		it.BasicErrWrapper.HandleError()
	}
}

func (it *Attributes) MustBeEmptyError() {
	if it.IsEmptyError() {
		return
	}

	panic(it.Error())
}

// SetAuthInfo
//
//	On nil create new attributes
func (it *Attributes) SetAuthInfo(authInfo *AuthInfo) *Attributes {
	if it == nil {
		return New.
			Attributes.
			UsingAuthInfo(authInfo)
	}

	it.AuthInfo = authInfo

	return it
}

func (it *Attributes) SetUserInfo(
	userInfo *UserInfo,
) *Attributes {
	if it == nil {
		return &Attributes{
			AuthInfo: &AuthInfo{
				UserInfo: userInfo,
			},
		}
	}

	it.AuthInfo.SetUserInfo(userInfo)

	return it
}

func (it *Attributes) AddNewStringKeyValueOnly(key, value string) (isAdded bool) {
	if it == nil || it.KeyValuePairs == nil {
		return false
	}

	it.KeyValuePairs.
		AddOrUpdate(key, value)

	return true
}

func (it *Attributes) AddNewAnyKeyValueOnly(
	key string, value any,
) (isAdded bool) {
	if it == nil || it.AnyKeyValuePairs == nil {
		return false
	}

	it.AnyKeyValuePairs.Add(key, value)

	return true
}

func (it *Attributes) AnyKeyReflectSetTo(
	key string,
	toPtr any,
) error {
	if it == nil || it.KeyValuePairs == nil {
		return errcore.
			CannotBeNilOrEmptyType.ErrorNoRefs(
			"KeyValuePairs is nil")
	}

	return it.AnyKeyValuePairs.ReflectSetTo(key, toPtr)
}

func (it *Attributes) ReflectSetTo(
	toPointer any,
) error {
	return coredynamic.ReflectSetFromTo(it, toPointer)
}

func (it *Attributes) AddOrUpdateString(
	key, value string,
) (isNewlyAdded bool) {
	if it == nil || it.KeyValuePairs == nil {
		return false
	}

	return it.
		KeyValuePairs.
		AddOrUpdate(key, value)
}

func (it *Attributes) AddOrUpdateAnyItem(
	key string,
	anyItem any,
) (isNewlyAdded bool) {
	if it == nil || it.AnyKeyValuePairs == nil {
		return false
	}

	return it.
		AnyKeyValuePairs.
		Add(key, anyItem)
}

// SetBasicErr
//
//	on nil creates and attach new error and returns the attributes
func (it *Attributes) SetBasicErr(
	basicErr errcoreinf.BasicErrWrapper,
) payloadinf.AttributesBinder {
	if it == nil {
		return New.Attributes.UsingBasicError(
			basicErr)
	}

	it.BasicErrWrapper = basicErr

	return it
}

func (it *Attributes) Clear() {
	if it == nil {
		return
	}

	it.KeyValuePairs.Clear()
	it.AnyKeyValuePairs.Clear()
	it.DynamicPayloads = []byte{}
}

func (it *Attributes) Dispose() {
	if it == nil {
		return
	}

	it.Clear()
}
