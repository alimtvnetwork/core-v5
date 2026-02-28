package coredynamic

import (
	"fmt"
	"reflect"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/coreonce"
	"gitlab.com/auk-go/core/issetter"
)

// Dynamic wraps an arbitrary value with cached reflection metadata.
//
// Getters and type checks are in DynamicGetters.go.
// Reflect-based operations and loops are in DynamicReflect.go.
// JSON serialization/deserialization is in DynamicJson.go.
type Dynamic struct {
	innerData       any
	isValid         bool
	reflectType     reflect.Type
	reflectVal      *reflect.Value
	innerDataString *string
	typeName        coreonce.StringOnce
	length          coreonce.IntegerOnce
	isPointer       issetter.Value
}

// =============================================================================
// Constructors
// =============================================================================

func InvalidDynamic() Dynamic {
	return *InvalidDynamicPtr()
}

func InvalidDynamicPtr() *Dynamic {
	return NewDynamicPtr(
		nil,
		false,
	)
}

func NewDynamicValid(
	data any,
) Dynamic {
	return *NewDynamicPtr(data, true)
}

func NewDynamic(
	data any,
	isValid bool,
) Dynamic {
	return *NewDynamicPtr(data, isValid)
}

func NewDynamicPtr(
	data any,
	isValid bool,
) *Dynamic {
	return &Dynamic{
		innerData: data,
		isValid:   isValid,
		typeName: coreonce.NewStringOnce(
			func() string {
				return fmt.Sprintf(constants.SprintTypeFormat, data)
			},
		),
		length: coreonce.NewIntegerOnce(
			func() int {
				if data == nil {
					return 0
				}

				return LengthOfReflect(reflect.ValueOf(data))
			},
		),
	}
}

// =============================================================================
// Clone
// =============================================================================

func (it Dynamic) Clone() Dynamic {
	return NewDynamic(
		it.innerData,
		it.isValid,
	)
}

func (it *Dynamic) ClonePtr() *Dynamic {
	if it == nil {
		return nil
	}

	return NewDynamicPtr(
		it.innerData,
		it.isValid,
	)
}

func (it Dynamic) NonPtr() Dynamic {
	return it
}

func (it *Dynamic) Ptr() *Dynamic {
	return it
}
