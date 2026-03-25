package coredynamic

import (
	"encoding/json"
	"reflect"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
	"github.com/alimtvnetwork/core/isany"
)

// ReflectSetFromTo
//
// # Set any object from to toPointer object
//
// Valid Inputs or Supported (https://t.ly/SGWUx):
//   - From, To: (null, null)                          -- do nothing
//   - From, To: (sameTypePointer, sameTypePointer)    -- try reflection
//   - From, To: (sameTypeNonPointer, sameTypePointer) -- try reflection
//   - From, To: ([]byte, otherType)                   -- try unmarshal, reflect
//   - From, To: (otherType, *[]byte)                  -- try marshal, reflect
//
// Validations:
//   - Check null, if both null no error return quickly.
//   - NotSupported returns as error.
//   - NotSupported: (from, to) - (..., not pointer)
//   - NotSupported: (from, to) - (null, notNull)
//   - NotSupported: (from, to) - (notNull, null)
//   - NotSupported: (from, to) - not same type and not bytes on any
//   - `From` null or nil is not supported and will return error.
//
// Reference:
//   - Reflection String Set Example : https://go.dev/play/p/fySLYuOvoRK.go?download=true
//   - Method document screenshot    : https://prnt.sc/26dmf5g
func ReflectSetFromTo(
	from,
	toPointer any,
) error {
	isLeftNull, isRightNull := isany.NullLeftRight(from, toPointer)

	if isLeftNull && isRightNull {
		return nil
	}

	leftRfType := reflect.TypeOf(from)
	rightRfType := reflect.TypeOf(toPointer)

	if err := validateReflectSetInputs(isLeftNull, isRightNull, leftRfType, rightRfType); err != nil {
		return err
	}

	leftRv := reflect.ValueOf(from)
	rightRv := reflect.ValueOf(toPointer)

	if err := validateLeftNotNil(leftRv, leftRfType, rightRfType); err != nil {
		return err
	}

	return reflectSetByType(from, toPointer, leftRfType, rightRfType, leftRv, rightRv)
}

// validateReflectSetInputs checks that the destination is non-nil and a pointer.
func validateReflectSetInputs(
	isLeftNull, isRightNull bool,
	leftRfType, rightRfType reflect.Type,
) error {
	if isRightNull {
		return errcore.
			InvalidNullPointerType.
			MsgCsvRefError(
				"\"destination pointer is null, cannot proceed further!\""+supportedTypesMessageReference,
				"FromType", leftRfType, "ToType", rightRfType,
			)
	}

	if rightRfType.Kind() != reflect.Ptr {
		return errcore.UnexpectedType.
			MsgCsvRefError(
				"\"destination or toPointer must be a pointer to set!\""+supportedTypesMessageReference,
				"FromType", leftRfType, "ToType", rightRfType,
			)
	}

	return nil
}

// validateLeftNotNil checks the source value is not nil.
func validateLeftNotNil(
	leftRv reflect.Value,
	leftRfType, rightRfType reflect.Type,
) error {
	isLeftAnyNull := reflectinternal.Is.NullRv(leftRv) ||
		reflectinternal.Is.Null(leftRfType)

	if isLeftAnyNull {
		return errcore.
			InvalidValueType.
			SrcDestinationErr(
				"`from` is nil, cannot set null or nil to destination.\"!"+supportedTypesMessageReference,
				"FromType", leftRfType,
				"ToType", rightRfType,
			)
	}

	return nil
}

// reflectSetByType dispatches to the appropriate set strategy based on types.
func reflectSetByType(
	from, toPointer any,
	leftRfType, rightRfType reflect.Type,
	leftRv, rightRv reflect.Value,
) error {
	// case: same pointer types — direct set
	if leftRfType == rightRfType {
		rightRv.Elem().Set(leftRv.Elem())
		return nil
	}

	// case: non-pointer source, pointer destination of same base type
	if leftRfType.Kind() != reflect.Ptr && leftRfType == rightRfType.Elem() {
		rightRv.Elem().Set(leftRv)
		return nil
	}

	isLeftBytes := leftRfType == emptyBytesType
	isRightBytesPointer := rightRfType == emptyBytesPointerType

	if !(leftRfType == rightRfType || isLeftBytes || isRightBytesPointer) {
		return errcore.
			TypeMismatchType.
			SrcDestinationErr(
				"supported: \"types are same pointer or any bytes or destination is pointer.\"!"+supportedTypesMessageReference,
				"FromType", leftRfType,
				"ToType", rightRfType,
			)
	}

	return reflectSetBytes(from, toPointer, isLeftBytes, isRightBytesPointer, leftRfType, rightRfType)
}

// reflectSetBytes handles byte-based serialization/deserialization.
func reflectSetBytes(
	from, toPointer any,
	isLeftBytes, isRightBytesPointer bool,
	leftRfType, rightRfType reflect.Type,
) error {
	// case: []byte → other type (unmarshal)
	if isLeftBytes {
		return corejson.
			Deserialize.
			UsingBytes(from.([]byte), toPointer)
	}

	// case: other type → *[]byte (marshal)
	if isRightBytesPointer {
		rawBytes, err := json.Marshal(from)
		if err != nil {
			return errcore.
				MarshallingFailedType.
				SrcDestinationErr(
					err.Error(),
					"FromType", leftRfType,
					"ToType", rightRfType,
				)
		}

		bytesPtr := toPointer.(*[]byte)
		*bytesPtr = rawBytes
		return nil
	}

	return errcore.
		UnMarshallingFailedType.
		SrcDestinationErr(
			"unexpected state in byte conversion",
			"FromType", leftRfType,
			"ToType", rightRfType,
		)
}
