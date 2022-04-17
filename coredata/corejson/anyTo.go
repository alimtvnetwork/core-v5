package corejson

import (
	"errors"

	"gitlab.com/evatix-go/core/internal/reflectinternal"
)

type anyTo struct{}

// SerializedJsonResult
//
//  Casting happens:
//  - self or self pointer returns directly
//  - []Bytes to Result
//  - string (json) to Result
//  - Jsoner to Result
//  - bytesSerializer to Result
//  - error to Result
//  - AnyItem
func (it anyTo) SerializedJsonResult(
	fromAny interface{},
) *Result {
	if reflectinternal.IsNull(fromAny) {
		return &Result{
			Error:    errors.New("nil object given"),
			TypeName: reflectinternal.SafeTypeName(fromAny),
		}
	}

	switch castedTo := fromAny.(type) {
	case Result:
		return castedTo.Ptr()
	case *Result:
		return castedTo
	case []byte:
		return NewResult.UsingBytesTypePtr(
			castedTo,
			"RawBytes",
		)
	case string:
		return NewResult.UsingBytesTypePtr(
			[]byte(castedTo),
			"RawString",
		)
	case Jsoner:
		return castedTo.JsonPtr()
	case bytesSerializer:
		return NewResult.UsingSerializer(castedTo)
	case error:
		if castedTo == nil || castedTo.Error() == "" {
			// empty err
			return NewResult.UsingBytesTypePtr(
				[]byte{},
				errTypeString,
			)
		}

		return NewResult.UsingTypePlusString(
			errTypeString, // type
			castedTo.Error()) // json string
	}

	return Serialize.Apply(
		fromAny)
}

func (it anyTo) SerializedRaw(
	fromAny interface{},
) (allBytes []byte, err error) {
	return it.SerializedJsonResult(fromAny).Raw()
}

// SerializedString
//
// accepted types (usages SerializedJsonResult):
//  - Result, *Result
//  - []byte
//  - string
//  - jsoner
//  - bytesSerializer
//  - anyItem
func (it anyTo) SerializedString(
	fromAny interface{},
) (serializedString string, err error) {
	jsonResult := it.SerializedJsonResult(fromAny)

	if jsonResult.HasError() {
		return "", jsonResult.MeaningfulError()
	}

	return jsonResult.JsonString(), nil
}

// SerializedSafeString
//
// accepted types (usages SerializedJsonResult):
//  - Result, *Result
//  - []byte
//  - string
//  - jsoner
//  - bytesSerializer
//  - anyItem
//
// Warning:
//  swallows error, important data convert must not go into this.
func (it anyTo) SerializedSafeString(
	fromAny interface{},
) (serializedString string) {
	jsonResult := it.SerializedJsonResult(fromAny)

	if jsonResult.HasError() {
		return ""
	}

	return jsonResult.JsonString()
}

func (it anyTo) SerializedStringMust(
	fromAny interface{},
) (serializedString string) {
	jsonResult := it.SerializedJsonResult(fromAny)
	jsonResult.MustBeSafe()

	return jsonResult.JsonString()
}

// SafeJsonString
//
//  warning : swallows error
func (it anyTo) SafeJsonString(
	anyItem interface{},
) string {
	jsonResult := New(anyItem)

	return jsonResult.JsonString()
}

func (it anyTo) PrettyStringWithError(
	anyItem interface{},
) (string, error) {
	switch casted := anyItem.(type) {
	case string:
		return casted, nil
	case []byte:
		return BytesToPrettyString(casted), nil
	case Result:
		if casted.HasError() {
			return casted.PrettyJsonString(), casted.MeaningfulError()
		}

		return casted.PrettyJsonString(), nil
	case *Result:
		if casted.HasError() {
			return casted.PrettyJsonString(), casted.MeaningfulError()
		}

		return casted.PrettyJsonString(), nil
	}

	jsonResult := New(anyItem)

	return jsonResult.PrettyJsonString(), jsonResult.MeaningfulError()
}

// SafeJsonPrettyString
//
//  warning : swallows error
func (it anyTo) SafeJsonPrettyString(
	anyItem interface{},
) string {
	switch casted := anyItem.(type) {
	case string:
		return casted
	case []byte:
		return BytesToPrettyString(casted)
	case Result:
		return casted.PrettyJsonString()
	case *Result:
		return casted.PrettyJsonString()
	}

	jsonResult := New(anyItem)

	return jsonResult.PrettyJsonString()
}

func (it anyTo) JsonString(
	anyItem interface{},
) string {
	switch casted := anyItem.(type) {
	case string:
		return casted
	case []byte:
		return BytesToString(casted)
	case Result:
		return casted.JsonString()
	case *Result:
		return casted.JsonString()
	}

	jsonResult := New(anyItem)

	return jsonResult.JsonString()
}

func (it anyTo) JsonStringWithErr(
	anyItem interface{},
) (jsonString string, parsingErr error) {
	switch casted := anyItem.(type) {
	case string:
		return casted, nil
	case []byte:
		return BytesToString(casted), nil
	case Result:
		if casted.HasError() {
			return casted.JsonString(), casted.MeaningfulError()
		}

		return casted.JsonString(), nil
	case *Result:
		if casted.HasError() {
			return casted.JsonString(), casted.MeaningfulError()
		}

		return casted.JsonString(), nil
	}

	jsonResult := New(anyItem)

	return jsonResult.JsonString(), jsonResult.MeaningfulError()
}

func (it anyTo) JsonStringMust(
	anyItem interface{},
) string {
	jsonStr, err := it.JsonStringWithErr(anyItem)

	if err != nil {
		panic(err)
	}

	return jsonStr
}

func (it anyTo) PrettyStringMust(
	anyItem interface{},
) string {
	jsonPretty, err := it.JsonStringWithErr(
		anyItem)

	if err != nil {
		panic(err)
	}

	return jsonPretty
}

func (it anyTo) UsingSerializer(
	serializer bytesSerializer,
) *Result {
	return NewResult.UsingSerializer(
		serializer)
}

// SerializedFieldsMap
//
//  usages json to bytes then use json to create fields map
func (it anyTo) SerializedFieldsMap(
	anyItem interface{},
) (fieldsMap map[string]interface{}, parsingErr error) {
	return it.
		SerializedJsonResult(anyItem).
		DeserializedFieldsToMap()
}

// SerializedFieldsMapFilter
//
//  usages json to bytes then use json to create fields map
func (it anyTo) SerializedFieldsMapFilter(
	isSkipInvalid bool,
	anyItem interface{},
	selectedFields ...string,
) (fieldsMap map[string]interface{}, parsingErr error) {
	currentMap, parsingErr := it.
		SerializedJsonResult(anyItem).
		DeserializedFieldsToMap()

	if parsingErr != nil || len(currentMap) == 0 {
		return currentMap, parsingErr
	}

	finalMap := make(map[string]interface{}, len(selectedFields))

	for _, fieldName := range selectedFields {
		currentVal, has := currentMap[fieldName]

		if isSkipInvalid && !has {
			continue
		}

		finalMap[fieldName] = currentVal
	}

	return finalMap, parsingErr
}

// SerializedFieldsMapSkipFilter
//
//  usages json to bytes then use json to create fields map
func (it anyTo) SerializedFieldsMapSkipFilter(
	anyItem interface{},
	skipFieldNames ...string,
) (fieldsMap map[string]interface{}, parsingErr error) {
	currentMap, parsingErr := it.
		SerializedJsonResult(anyItem).
		DeserializedFieldsToMap()

	if parsingErr != nil || len(currentMap) == 0 {
		return currentMap, parsingErr
	}

	for _, skipFieldName := range skipFieldNames {
		_, has := currentMap[skipFieldName]

		if has {
			delete(currentMap, skipFieldName)
		}
	}

	return currentMap, parsingErr
}



// SerializedMappedFields
//
//  usages json to bytes then use json to create fields map
func (it anyTo) SerializedMappedFields(
	anyItem interface{},
) (mappedFields MappedFields, parsingErr error) {
	currentMap, parsingErr := it.
		SerializedJsonResult(anyItem).
		DeserializedFieldsToMap()

	return MappedFields{
		TypeName:  reflectinternal.TypeName(anyItem),
		FieldsMap: currentMap,
	}, parsingErr
}
