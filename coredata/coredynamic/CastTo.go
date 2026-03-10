package coredynamic

import (
	"reflect"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

func CastTo(
	isOutputPointer bool,
	input any,
	acceptedTypes ...reflect.Type,
) CastedResult {
	currentRfType := reflect.TypeOf(input)
	rv := reflect.ValueOf(input)
	kind := rv.Kind()
	var sliceErr []string

	isMatchingAcceptedType := IsAnyTypesOf(
		currentRfType,
		acceptedTypes...,
	)

	isNoMatch := !isMatchingAcceptedType

	if isNoMatch {
		// not matching
		sliceErr = append(
			sliceErr,
			errcore.UnsupportedType.Combine(
				"none matches, current type:"+currentRfType.String(),
				getTypeNamesUsingReflectFunc(true, acceptedTypes...),
			),
		)
	}

	isNull := input == nil || reflectinternal.Is.NullRv(
		rv,
	)
	isOutNonPointer := !isOutputPointer
	hasNonPointerIssue := isNull && isOutNonPointer

	if hasNonPointerIssue {
		sliceErr = append(
			sliceErr,
			errcore.
				InvalidNullPointerType.
				SrcDestination(
					"cannot output non pointer if pointer is null",
					"Value", constants.NilAngelBracket,
					"Type", currentRfType.String(),
				),
		)

		return CastedResult{
			Casted:                 nil,
			SourceReflectType:      currentRfType,
			SourceKind:             kind,
			Error:                  errcore.SliceToError(sliceErr),
			IsNull:                 isNull,
			IsMatchingAcceptedType: isMatchingAcceptedType,
			IsPointer:              isOutNonPointer,
			IsSourcePointer:        kind == reflect.Ptr,
			IsValid:                rv.IsValid(),
		}
	}

	val, _ := PointerOrNonPointerUsingReflectValue(
		isOutputPointer,
		rv,
	)

	return CastedResult{
		Casted:                 val,
		SourceReflectType:      currentRfType,
		SourceKind:             kind,
		Error:                  errcore.SliceToError(sliceErr),
		IsNull:                 isNull,
		IsMatchingAcceptedType: isMatchingAcceptedType,
		IsPointer:              isOutNonPointer,
		IsSourcePointer:        kind == reflect.Ptr,
		IsValid:                rv.IsValid(),
	}
}
