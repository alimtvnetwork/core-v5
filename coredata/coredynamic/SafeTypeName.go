package coredynamic

import "reflect"

func SafeTypeName(item any) string {
	rf := reflect.TypeOf(item)

	if rf == nil {
		return ""
	}

	return rf.String()
}
