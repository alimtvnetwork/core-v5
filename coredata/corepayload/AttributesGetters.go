package corepayload

import "gitlab.com/auk-go/core/coreinterface/errcoreinf"

// AttributesGetters.go — Read-only accessor and query methods extracted from Attributes.go

func (it *Attributes) IsNull() bool {
	return it == nil
}

func (it *Attributes) HasSafeItems() bool {
	return !it.HasIssuesOrEmpty()
}

func (it *Attributes) HasStringKey(key string) bool {
	if it.HasKeyValuePairs() {
		return it.KeyValuePairs.Has(key)
	}

	return false
}

func (it *Attributes) HasAnyKey(key string) bool {
	if it.HasAnyKeyValuePairs() {
		return it.AnyKeyValuePairs.HasKey(key)
	}

	return false
}

func (it *Attributes) Payloads() []byte {
	if it.IsEmpty() {
		return []byte{}
	}

	return it.DynamicPayloads
}

func (it *Attributes) PayloadsString() string {
	if it.IsEmpty() || len(it.DynamicPayloads) == 0 {
		return ""
	}

	return string(it.DynamicPayloads)
}

func (it *Attributes) AnyKeyValMap() map[string]any {
	if it.IsEmpty() {
		return map[string]any{}
	}

	return it.AnyKeyValuePairs.Items
}

func (it *Attributes) Hashmap() map[string]string {
	if it.IsEmpty() {
		return map[string]string{}
	}

	return it.KeyValuePairs.Items()
}

func (it *Attributes) CompiledError() error {
	return it.Error()
}

func (it *Attributes) HasIssuesOrEmpty() bool {
	return it.IsEmpty() ||
		!it.IsValid() ||
		it.BasicErrWrapper != nil &&
			it.BasicErrWrapper.HasError()
}

func (it *Attributes) IsSafeValid() bool {
	return it.HasIssuesOrEmpty()
}

func (it *Attributes) HasAnyItem() bool {
	return !it.IsEmpty()
}

func (it *Attributes) Count() int {
	return it.Length()
}

func (it *Attributes) Capacity() int {
	return it.Length()
}

func (it *Attributes) Length() int {
	if it == nil {
		return 0
	}

	return len(it.DynamicPayloads)
}

func (it *Attributes) HasPagingInfo() bool {
	return it != nil && it.PagingInfo != nil
}

func (it *Attributes) HasKeyValuePairs() bool {
	return it != nil && it.KeyValuePairs.HasAnyItem()
}

func (it *Attributes) HasFromTo() bool {
	return it != nil && it.FromTo != nil
}

func (it *Attributes) IsValid() bool {
	return it != nil &&
		it.IsEmptyError()
}

func (it *Attributes) IsInvalid() bool {
	return it == nil || it.HasIssuesOrEmpty()
}

func (it *Attributes) HasError() bool {
	return it != nil &&
		it.BasicErrWrapper != nil &&
		it.BasicErrWrapper.HasError()
}

func (it *Attributes) Error() error {
	if it.IsEmptyError() {
		return nil
	}

	return it.
		BasicErrWrapper.
		CompiledErrorWithStackTraces()
}

func (it *Attributes) IsEmptyError() bool {
	return it == nil ||
		it.BasicErrWrapper == nil ||
		it.BasicErrWrapper.IsEmpty()
}

func (it *Attributes) DynamicBytesLength() int {
	if it == nil {
		return 0
	}

	return len(it.DynamicPayloads)
}

func (it *Attributes) StringKeyValuePairsLength() int {
	if it == nil {
		return 0
	}

	return it.KeyValuePairs.Length()
}

func (it *Attributes) AnyKeyValuePairsLength() int {
	if it == nil {
		return 0
	}

	return it.AnyKeyValuePairs.Length()
}

func (it *Attributes) IsEmpty() bool {
	return it == nil ||
		it.DynamicBytesLength() == 0 &&
			it.StringKeyValuePairsLength() == 0 &&
			it.AnyKeyValuePairsLength() == 0
}

func (it *Attributes) HasItems() bool {
	return !it.IsEmpty()
}

func (it *Attributes) IsPagingInfoEmpty() bool {
	return it == nil ||
		it.PagingInfo.IsEmpty()
}

func (it *Attributes) IsKeyValuePairsEmpty() bool {
	return it == nil ||
		it.KeyValuePairs.IsEmpty()
}

func (it *Attributes) IsAnyKeyValuePairsEmpty() bool {
	return it == nil ||
		it.AnyKeyValuePairs.IsEmpty()
}

func (it *Attributes) IsUserInfoEmpty() bool {
	return it == nil ||
		it.AuthInfo.IsUserInfoEmpty()
}

func (it *Attributes) VirtualUser() *User {
	if it.IsUserInfoEmpty() {
		return nil
	}

	return it.AuthInfo.UserInfo.User
}

func (it *Attributes) SystemUser() *User {
	if it.IsUserInfoEmpty() {
		return nil
	}

	return it.AuthInfo.UserInfo.SystemUser
}

func (it *Attributes) SessionUser() *User {
	if it.IsSessionInfoEmpty() {
		return nil
	}

	return it.AuthInfo.SessionInfo.User
}

func (it *Attributes) IsAuthInfoEmpty() bool {
	return it == nil ||
		it.AuthInfo.IsEmpty()
}

func (it *Attributes) IsSessionInfoEmpty() bool {
	return it == nil ||
		it.AuthInfo.IsSessionInfoEmpty()
}

func (it *Attributes) HasUserInfo() bool {
	return !it.IsUserInfoEmpty()
}

func (it *Attributes) HasAuthInfo() bool {
	return !it.IsAuthInfoEmpty()
}

func (it *Attributes) HasSessionInfo() bool {
	return !it.IsSessionInfoEmpty()
}

func (it *Attributes) SessionInfo() *SessionInfo {
	if it.IsSessionInfoEmpty() {
		return nil
	}

	return it.AuthInfo.SessionInfo
}

func (it *Attributes) AuthType() string {
	if it.IsAuthInfoEmpty() {
		return ""
	}

	return it.AuthInfo.ActionType
}

func (it *Attributes) ResourceName() string {
	if it.IsAuthInfoEmpty() {
		return ""
	}

	return it.AuthInfo.ResourceName
}

func (it *Attributes) HasStringKeyValuePairs() bool {
	return it.StringKeyValuePairsLength() > 0
}

func (it *Attributes) HasAnyKeyValuePairs() bool {
	return it.AnyKeyValuePairsLength() > 0
}

func (it *Attributes) HasDynamicPayloads() bool {
	return it.DynamicBytesLength() > 0
}

func (it *Attributes) GetStringKeyValue(
	key string,
) (value string, isFound bool) {
	if it == nil || it.KeyValuePairs == nil {
		return "", false
	}

	return it.KeyValuePairs.Get(key)
}

func (it *Attributes) GetAnyKeyValue(
	key string,
) (valueAny any, isFound bool) {
	if it == nil || it.KeyValuePairs == nil {
		return nil, false
	}

	return it.AnyKeyValuePairs.Get(key)
}

func (it *Attributes) IsErrorDifferent(basicErr errcoreinf.BasicErrWrapper) bool {
	return !it.IsErrorEqual(basicErr)
}

func (it *Attributes) IsErrorEqual(basicErr errcoreinf.BasicErrWrapper) bool {
	if it.IsEmptyError() || basicErr == nil || basicErr.IsEmpty() {
		return true
	}

	return it.BasicErrWrapper.IsBasicErrEqual(basicErr)
}
