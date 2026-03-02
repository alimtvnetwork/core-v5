package coredynamic

import (
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type LeftRight struct {
	Left, Right any
}

func (it *LeftRight) IsEmpty() bool {
	return it == nil ||
		reflectinternal.Is.Null(it.Left) &&
			reflectinternal.Is.Null(it.Right)
}

func (it *LeftRight) HasAnyItem() bool {
	return !it.IsEmpty()
}

func (it *LeftRight) HasLeft() bool {
	return it != nil &&
		!reflectinternal.Is.Null(it.Left)
}

func (it *LeftRight) HasRight() bool {
	return it != nil &&
		!reflectinternal.Is.Null(it.Right)
}

func (it *LeftRight) IsLeftEmpty() bool {
	return it == nil ||
		reflectinternal.Is.Null(it.Left)
}

func (it *LeftRight) IsRightEmpty() bool {
	return it == nil ||
		reflectinternal.Is.Null(it.Right)
}

func (it *LeftRight) LeftReflectSet(
	toPointerOrBytesPointer any,
) error {
	if it == nil {
		return nil
	}

	return ReflectSetFromTo(it.Left, toPointerOrBytesPointer)
}

func (it *LeftRight) RightReflectSet(
	toPointerOrBytesPointer any,
) error {
	if it == nil {
		return nil
	}

	return ReflectSetFromTo(it.Right, toPointerOrBytesPointer)
}

func (it *LeftRight) DeserializeLeft() *corejson.Result {
	if it == nil {
		return nil
	}

	return corejson.NewPtr(it.Left)
}

func (it *LeftRight) DeserializeRight() *corejson.Result {
	if it == nil {
		return nil
	}

	return corejson.NewPtr(it.Right)
}

func (it *LeftRight) LeftToDynamic() *Dynamic {
	if it == nil {
		return nil
	}

	return NewDynamicPtr(it.Left, true)
}

func (it *LeftRight) RightToDynamic() *Dynamic {
	if it == nil {
		return nil
	}

	return NewDynamicPtr(it.Right, true)
}

func (it *LeftRight) TypeStatus() TypeStatus {
	if it == nil {
		return TypeSameStatus(nil, nil)
	}

	return TypeSameStatus(it.Left, it.Right)
}
