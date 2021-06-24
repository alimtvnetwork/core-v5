package main

import (
	"fmt"

	"gitlab.com/evatix-go/core/coredata/corestr"
)

func main() {
	// rwx, err := chmodhelper.NewRwxVariableWrapper("-rwx-*-r*x")
	// fmt.Println(err)
	// fmt.Println(rwx.String())
	//
	// rs := rwx.IsEqualPartialRwxPartial("-rwx-w-r*x")
	// fmt.Println(rs)
	//
	// fmt.Println(msgtype.ExpectingSimpleNoType("Alim", "Rwx", "wrx"))
	//
	// rwxIns := chmodins.RwxInstruction{
	// 	RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
	// 		Owner: "rwx",
	// 		Group: "rwx",
	// 		Other: "r--",
	// 	},
	// 	Condition: chmodins.Condition{
	// 		IsSkipOnInvalid:  false,
	// 		IsContinueOnError: false,
	// 		IsRecursive:       false,
	// 	},
	// }
	//
	// executor, err := chmodhelper.ParseRwxInstructionToExecutor(&rwxIns)
	//
	// msgtype.SimpleHandleErr(err, "")
	//
	// locations := []string{
	// 	"/temp/core/test-cases-2",
	// 	"/temp/core/test-cases-3s",
	// 	"/temp/core/test-cases-3x",
	// 	"/temp/core/test-cases-3",
	// }
	//
	// err2 := chmodhelper.VerifyChmodLocationsUsingPartialRwx(
	// 	true, true,
	// 	"-rwxrwx",
	// 	locations)
	//
	// err3 := executor.VerifyRwxModifiersDirect(
	// 	false,
	// 	locations...)
	//
	// msgtype.SimpleHandleErrMany("", err2)
}

func PrintCollection(collection *corestr.Collection) {
	fmt.Println(collection.GetPagedCollection(3).String())

	fmt.Print("\n\nTake 5:\n\n")
	fmt.Println(collection.Take(5))
	fmt.Print("\n\n Skip 2:\n\n")
	fmt.Println(collection.Skip(2))
	fmt.Print("\n\n Skip 0:\n\n")
	fmt.Println(collection.Skip(0))
	fmt.Print("\n\n Take 0:\n\n")
	fmt.Println(collection.Take(0))
	fmt.Print("\n\n Skip(5).Take(2):\n\n")
	fmt.Println(collection.Skip(5).Take(2))

}

func PrintCollectionPtr(collectionPtr *corestr.CollectionPtr) {
	fmt.Println(collectionPtr.GetPagedCollection(3).String())

	fmt.Print("\n\nTake 5:\n\n")
	fmt.Println(collectionPtr.Take(5))
	fmt.Print("\n\n Skip 2:\n\n")
	fmt.Println(collectionPtr.Skip(2))
	fmt.Print("\n\n Skip 0:\n\n")
	fmt.Println(collectionPtr.Skip(0))
	fmt.Print("\n\n Take 0:\n\n")
	fmt.Println(collectionPtr.Take(0))
	fmt.Print("\n\n Skip(5).Take(2):\n\n")
	fmt.Println(collectionPtr.Skip(5).Take(2))
}
