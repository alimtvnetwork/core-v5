package main

import (
	"fmt"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
)

func jsonResultUnmarshallingTest01() {
	mapAnyItems := getMapAnyItems()
	jsonResult := mapAnyItems.JsonPtr()
	emptyMapResult := coredynamic.EmptyMapAnyItems()

	err := jsonResult.Unmarshal(emptyMapResult)
	fmt.Println("err:", err)
	fmt.Println(emptyMapResult)
}
