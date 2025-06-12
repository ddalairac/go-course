package main

import (
	"go-course/helpers"
	"go-course/operations"
	"go-course/types"
)

func main() {
	helpers.PrintTitle("Package types")
	types.BasicsTypes()
	types.StructsTypes()
	types.TypeMap()
	types.TypeSlice()
	types.TypeInterface()

	helpers.PrintTitle("Package operations")
	operations.OperationCondition()
	operations.OperationLoops()

	// fmt.Println(" ************************************************** ")
}
