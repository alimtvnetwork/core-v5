package corejson

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/internal/reflectinternal"
	"gitlab.com/evatix-go/core/isany"
)

type MappedFields struct {
	TypeName  string
	FieldsMap map[string]interface{}
}

func (it *MappedFields) IsAnyNull() bool {
	return it == nil || it.FieldsMap == nil
}

func (it *MappedFields) IsNull() bool {
	return it == nil
}

func (it *MappedFields) IsDefined() bool {
	return it != nil && it.TypeName != "" && it.FieldsMap != nil
}

func (it *MappedFields) HasTypeName() bool {
	return it != nil && it.TypeName != ""
}

func (it *MappedFields) HasFields() bool {
	return it != nil && it.Length() > 0
}

func (it *MappedFields) HasAnyFields() bool {
	return it != nil && it.Length() > 0
}

func (it *MappedFields) HasAnyItem() bool {
	return it != nil && it.Length() > 0
}

func (it *MappedFields) Length() int {
	if it == nil {
		return 0
	}

	return len(it.FieldsMap)
}

func (it *MappedFields) IsTypeMatch(anyItem interface{}) bool {
	if it == nil && isany.ReflectNull(anyItem) {
		return true
	}

	if it == nil || isany.ReflectNull(anyItem) {
		return false
	}

	rightTypeName := reflectinternal.TypeName(anyItem)

	if rightTypeName == it.TypeName {
		return true
	}

	if rightTypeName == "" || it.TypeName == "" {
		return false
	}

	if rightTypeName[1:] == it.TypeName && rightTypeName[0] == constants.WildcardChar {
		return true
	}

	if rightTypeName == it.TypeName[1:] && it.TypeName[0] == constants.WildcardChar {
		return true
	}

	return false
}

func (it *MappedFields) GetFieldVal(fieldName string) (val interface{}, isFound bool) {
	if it == nil {
		return nil, false
	}

	val, isFound = it.FieldsMap[fieldName]

	return val, isFound
}

func (it *MappedFields) GetFieldValOnly(
	fieldName string,
) (val interface{}) {
	if it == nil {
		return nil
	}

	val, isFound := it.FieldsMap[fieldName]

	if isFound {
		return val
	}

	return nil
}
