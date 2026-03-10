package codegen

import (
	"github.com/alimtvnetwork/core/coretests/args"
)

type Map map[string]args.Map

//
// func (it Map) Set(k, k2, k3 string, v any) Map {
// 	finalMap := it.GetSetMap(k, k2)
//
// 	finalMap[k3] = v
// }
//
// func (it Map) IsEmpty() bool {
// 	return len(it) == 0
// }
//
// func (it Map) Length() int {
// 	return len(it)
// }
//
// func (it Map) Count() int {
// 	return len(it)
// }
//
// func (it Map) HasKey1(k string) bool {
// 	if it.IsEmpty() {
// 		return false
// 	}
//
// 	_, has := it[k]
//
// 	return has
// }
//
// func (it Map) HasKey2(k, k2 string) bool {
// 	if it.IsEmpty() {
// 		return false
// 	}
//
// 	if !it.HasKey1(k) {
// 		return false
// 	}
//
// 	_, has := it[k][k2]
//
// 	return has
// }
//
// func (it Map) HasKey3(k, k2, k3 string) bool {
// 	if it.IsEmpty() {
// 		return false
// 	}
//
// 	if !it.HasKey2(k, k2) {
// 		return false
// 	}
//
// 	x, _ := it[k][k2].(args.Map)
// 	_, has := x[k3]
//
// 	return has
// }
//
// func (it Map) GetSetMap(k, k2 string) args.Map {
// 	x, has := it[k]
//
// 	if has {
// 		y, isPresent := x[k2]
//
// 		if isPresent {
// 			return y.(args.Map)
// 		} else {
// 			it[k][k2] = new(args.Map)
//
// 			return x[k2].(args.Map)
// 		}
// 	}
//
// 	it[k] = args.Map{}
// 	it[k][k2] = args.Map{}
//
// 	return it[k][k2].(args.Map)
// }
