package main

import (
	"fmt"
	"go-course/operations"
	"go-course/types"
)

func main() {
	fmt.Println("************************************************** ")
	fmt.Println("Package types")
	types.BasicsTypes()
	types.StructsTypes()
	types.TypeMap()
	types.TypeSlice()
	types.TypeInterface()

	fmt.Println("")
	fmt.Println("************************************************** ")
	fmt.Println("Package operations")
	operations.OperationCondition()
	operations.OperationLoops()

	// fmt.Println(" ************************************************** ")
}
