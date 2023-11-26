package reflectinternal

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"gitlab.com/auk-go/core/internal/convertinteranl"
)

type reflectUtils struct{}

func (it reflectUtils) MaxLimit(currentLength, maxCapacity int) int {
	if maxCapacity <= -1 {
		return currentLength
	}

	if currentLength >= maxCapacity {
		return maxCapacity
	}

	return currentLength
}

func (it reflectUtils) AppendArgs(appendingItem interface{}, args []interface{}) []interface{} {
	if len(args) == 0 {
		return []interface{}{appendingItem}
	}

	list := make(
		[]interface{},
		len(args)+1,
	)

	list[0] = appendingItem

	for i, arg := range args {
		list[i+1] = reflect.ValueOf(arg)
	}

	return list
}

func (it reflectUtils) VerifyReflectTypesAny(left, right []interface{}) (isOkay bool, err error) {
	leftLen := len(left)
	rightLen := len(right)

	if leftLen != rightLen {
		errMsg := fmt.Sprintf(
			"Left Len(%d) != Right Len (%d)",
			leftLen,
			rightLen,
		)

		return false, errors.New(errMsg)
	}

	var errSlice []string

	for i := 0; i < leftLen; i++ {
		l := left[i]
		r := right[i]

		isCurrTypeOkay, errFirst := it.IsReflectTypeMatchAny(l, r)

		if isCurrTypeOkay {
			continue
		}

		if errFirst != nil {
			errSlice = append(
				errSlice,
				it.errorMessageForTypeVerification(i, errFirst),
			)
		}
	}

	if len(errSlice) == 0 {
		return true, nil
	}

	return false, errors.New(strings.Join(errSlice, "\n"))
}

func (it reflectUtils) errorMessageForTypeVerification(i int, errFirst error) string {
	return fmt.Sprintf("- Index {%d} - %dth : %s", i, i+1, errFirst.Error())
}

func (it reflectUtils) VerifyReflectTypes(
	rootName string,
	expectedArgs,
	givenArgs []reflect.Type,
) (isOkay bool, err error) {
	leftLen := len(expectedArgs)
	rightLen := len(givenArgs)

	if leftLen != rightLen {
		errMsg := fmt.Sprintf(
			"Expected Length (%d) != (%d) Given Length",
			leftLen,
			rightLen,
		)

		return false, errors.New(errMsg)
	}

	var errSlice []string

	for i := 0; i < leftLen; i++ {
		l := expectedArgs[i]
		r := givenArgs[i]

		isCurrTypeOkay, errFirst := it.IsReflectTypeMatch(l, r)

		if isCurrTypeOkay {
			continue
		}

		if errFirst != nil {
			errSlice = append(
				errSlice,
				it.errorMessageForTypeVerification(i, errFirst),
			)
		}
	}

	if len(errSlice) == 0 {
		return true, nil
	}

	convertinteranl.Util.String

	return false, errors.New(finalErrMessage)
}

func (it reflectUtils) IsReflectTypeMatch(expectedType, givenType reflect.Type) (isOkay bool, err error) {
	if expectedType == givenType {
		return true, nil
	}

	errMsg := fmt.Sprintf(
		"Expected Type (%s) != (%s) Given Type",
		expectedType.Name(),
		givenType.Name(),
	)

	return false, errors.New(errMsg)
}

func (it reflectUtils) IsReflectTypeMatchAny(left, right interface{}) (isOkay bool, err error) {
	l := reflect.TypeOf(left)
	r := reflect.TypeOf(right)

	return it.IsReflectTypeMatch(l, r)
}
