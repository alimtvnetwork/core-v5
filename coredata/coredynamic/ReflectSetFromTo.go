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
	isLeftNull, isRightNull := isany.NullLeftRight(
		from,
		toPointer,
	)

	if isLeftNull == isRightNull && isLeftNull {
		return nil
	}

	leftRfType := reflect.TypeOf(from)
	rightRfType := reflect.TypeOf(toPointer)

	if isRightNull {
		return errcore.
			InvalidNullPointerType.
			MsgCsvRefError(
				"\"destination pointer is null, cannot proceed further!\""+supportedTypesMessageReference,
				"FromType", leftRfType, "ToType", rightRfType,
			)
	}

	rightKind := rightRfType.Kind()
	isRightKindNotPointer := rightKind != reflect.Ptr
	if isRightKindNotPointer {
		return errcore.UnexpectedType.
			MsgCsvRefError(
				"\"destination or toPointer must be a pointer to set!\""+supportedTypesMessageReference,
				"FromType", leftRfType, "ToType", rightRfType,
			)
	}

	isSameType := leftRfType == rightRfType
	leftRv := reflect.ValueOf(from)
	rightRv := reflect.ValueOf(toPointer) // right is pointer confirmed by previous validation
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

	isLeftBytes := leftRfType == emptyBytesType
	isRightBytesPointer := rightRfType == emptyBytesPointerType
	isAnyBytes := isLeftBytes || isRightBytesPointer

	// case : From, To  : (sameTypePointer, sameTypePointer)    -- try reflection
	if leftRfType == rightRfType {
		// reflect, both same
		rightRv.Elem().Set(leftRv.Elem())

		return nil
	}

	// case : From, To  : (sameTypeNonPointer, sameTypePointer) -- try reflection
	if leftRfType.Kind() != reflect.Ptr && !isLeftNull && leftRfType == rightRfType.Elem() {
		rightRv.Elem().Set(leftRv)

		return nil
	}

	isNotSupportedType := !(isSameType || isAnyBytes)

	if isNotSupportedType {
		return errcore.
			TypeMismatchType.
			SrcDestinationErr(
				"supported: \"types are same pointer or any bytes or destination is pointer.\"!"+supportedTypesMessageReference,
				"FromType", leftRfType,
				"ToType", rightRfType,
			)
	}

	// case : From, To  : ([]byte, otherType)  -- try unmarshal, reflect
	if isLeftBytes {
		return corejson.
			Deserialize.
			UsingBytes(
				from.([]byte),
				toPointer,
			)
	}

	// case : From, To: (otherType, *[]byte) -- try marshal, reflect
	var rawBytes []byte
	var finalErr error

	if isRightBytesPointer {
		rawBytes, finalErr = json.Marshal(from)
	}

	if finalErr != nil {
		return errcore.
			MarshallingFailedType.
			SrcDestinationErr(
				finalErr.Error(),
				"FromType", leftRfType,
				"ToType", rightRfType,
			)
	}

	// For *[]byte destination, set marshaled bytes directly
	if isRightBytesPointer {
		bytesPtr := toPointer.(*[]byte)
		*bytesPtr = rawBytes

		return nil
	}

	err := json.Unmarshal(
		rawBytes,
		toPointer,
	)

	if err == nil {
		return nil
	}

	// has error
	return errcore.
		UnMarshallingFailedType.
		SrcDestinationErr(
			err.Error(),
			"FromType", leftRfType,
			"ToType", rightRfType,
		)
}
