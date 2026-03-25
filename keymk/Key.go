package keymk

import (
	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/errcore"
)

type Key struct {
	option        *Option
	mainName      string
	keyChains     []string
	compiledChain *string
}

func (it *Key) CompiledChain() string {
	if it.IsComplete() {
		return *it.compiledChain
	}

	return constants.EmptyString
}

func (it *Key) MainName() string {
	return it.mainName
}

func (it *Key) IsEmpty() bool {
	return it.Length() == 0 && it.MainName() == ""
}

func (it *Key) Length() int {
	if it == nil {
		return 0
	}

	return len(it.keyChains)
}

func (it *Key) AppendChain(items ...any) *Key {
	if it.IsComplete() {
		// panic
		errcore.CannotModifyCompleteResourceType.HandleUsingPanic(
			cannotModifyErrorMessage,
			items)
	}

	it.keyChains = appendAnyItemsWithBaseStrings(
		it.option.IsSkipEmptyEntry,
		it.keyChains,
		items)

	return it
}

func (it *Key) AppendChainKeys(
	keys ...*Key,
) *Key {
	if len(keys) == 0 {
		return it
	}

	for _, key := range keys {
		if key == nil {
			continue
		}

		it.AppendChainStrings(key.MainName())
		it.AppendChainStrings(key.keyChains...)
	}

	return it
}

func (it *Key) AppendChainStrings(
	items ...string,
) *Key {
	if it.IsComplete() {
		// panic
		errcore.CannotModifyCompleteResourceType.HandleUsingPanic(
			cannotModifyErrorMessage,
			items)
	}

	isSkipOnEmpty := it.option.IsSkipEmptyEntry

	for _, item := range items {
		if isSkipOnEmpty && item == "" {
			continue
		}

		it.keyChains = append(
			it.keyChains,
			item)
	}

	return it
}

func (it *Key) KeyChains() []string {
	if it == nil {
		return nil
	}

	return it.keyChains
}

// AllRawItems
//
// Returns main + whole chain (raw elements)
func (it *Key) AllRawItems() []string {
	if it == nil {
		return nil
	}

	return stringslice.PrependLineNew(
		it.MainName(),
		it.KeyChains())
}

func (it *Key) HasInChains(
	chainItem string,
) bool {
	if it == nil {
		return false
	}

	for _, chain := range it.keyChains {
		if chain == chainItem {
			return true
		}
	}

	return false
}

func (it *Key) IsComplete() bool {
	return it.compiledChain != nil
}

func (it *Key) ConcatNewUsingKeys(
	keys ...*Key,
) *Key {
	cloned := it.ClonePtr()

	return cloned.AppendChainKeys(keys...)
}

func (it *Key) ClonePtr(
	newAppendingChains ...any,
) *Key {
	if it == nil {
		return nil
	}

	key := NewKey.All(
		it.option.ClonePtr(),
		it.mainName,
	)

	key.AppendChainStrings(
		it.keyChains...)

	return key.AppendChain(
		newAppendingChains...)
}

func (it *Key) String() string {
	return it.Compile()
}

func (it *Key) Strings() []string {
	return it.AllRawItems()
}

func (it *Key) Name() string {
	return it.Compile()
}

func (it *Key) KeyCompiled() string {
	return it.Compile()
}
