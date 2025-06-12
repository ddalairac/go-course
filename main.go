package main

import (
	"go-course/channel"
	"go-course/data"
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

	helpers.PrintTitle("Package channels")
	intChan := make(chan int)
	defer close(intChan)
	channel.Channel(intChan)

	helpers.PrintTitle("Package data")
	data.JsonReadWrite()
}
