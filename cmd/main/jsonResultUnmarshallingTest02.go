package main

import (
	"fmt"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
)

func jsonResultUnmarshallingTest02() {
	mapAnyItems := getMapAnyItems()
	jsonResult := mapAnyItems.JsonPtr()
	var emptyMapResult *coredynamic.MapAnyItems

	err := jsonResult.Unmarshal(emptyMapResult)
	fmt.Println("err:", err)
	fmt.Println(emptyMapResult)

	var emptyMapResult2 coredynamic.MapAnyItems

	unmarshalValueErr := jsonResult.Unmarshal(emptyMapResult2)
	fmt.Println("unmarshalValueErr:", unmarshalValueErr)
	fmt.Println(emptyMapResult2)

	var emptyMapResult3 coredynamic.MapAnyItems

	unmarshalPtrErr := jsonResult.Unmarshal(&emptyMapResult3)
	fmt.Println("unmarshalPtrErr:", unmarshalPtrErr)
	fmt.Println(emptyMapResult3)

	selfInjectErr := emptyMapResult3.JsonParseSelfInject(jsonResult)
	fmt.Println("selfInjectErr:", selfInjectErr)
	fmt.Println("emptyMapResult3:", emptyMapResult3)

	selfInjectNilErr := emptyMapResult3.JsonParseSelfInject(nil)
	fmt.Println("selfInjectNilErr:", selfInjectNilErr)
	fmt.Println("emptyMapResult3:", emptyMapResult3)

	nilMapInjectErr := emptyMapResult.JsonParseSelfInject(jsonResult)
	fmt.Println("nil emptyMapResult nilMapInjectErr:", nilMapInjectErr)

	jsonResult = nil
	nilJsonUnmarshalErr := jsonResult.Unmarshal(&emptyMapResult3)
	fmt.Println("json Result nil, nilJsonUnmarshalErr:", nilJsonUnmarshalErr)

	jsonResult = nil
	nilBothUnmarshalErr := jsonResult.Unmarshal(emptyMapResult)
	fmt.Println("json Result, object nil, nilBothUnmarshalErr:", nilBothUnmarshalErr)
}
