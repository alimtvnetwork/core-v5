package reflectinternal

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type utils struct{}

func (it utils) MaxLimit(currentLength, maxCapacity int) int {
	if maxCapacity <= -1 {
		return currentLength
	}

	if currentLength >= maxCapacity {
		return maxCapacity
	}

	return currentLength
}

func (it utils) AppendArgs(appendingItem interface{}, args []interface{}) []interface{} {
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

func (it utils) VerifyReflectTypesAny(left, right []interface{}) (isOkay bool, err error) {
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

func (it utils) errorMessageForTypeVerification(i int, errFirst error) string {
	return fmt.Sprintf("- Index {%d} - %dth : %s", i, i+1, errFirst.Error())
}

func (it utils) VerifyReflectTypes(left, right []reflect.Type) (isOkay bool, err error) {
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

	return false, errors.New(strings.Join(errSlice, "\n"))
}

func (it utils) IsReflectTypeMatch(left, right reflect.Type) (isOkay bool, err error) {
	if left == right {
		return true, nil
	}

	errMsg := fmt.Sprintf(
		"Left Type (%s) != Right Type (%s)",
		left.Name(),
		right.Name(),
	)

	return false, errors.New(errMsg)
}

func (it utils) IsReflectTypeMatchAny(left, right interface{}) (isOkay bool, err error) {
	l := reflect.TypeOf(left)
	r := reflect.TypeOf(right)

	return it.IsReflectTypeMatch(l, r)
}
