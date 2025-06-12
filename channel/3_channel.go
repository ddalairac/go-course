package channel

import (
	"fmt"
	"go-course/helpers"
)

func Channel(intChan chan int) {

	go ChannelChange(intChan)
	num := <-intChan
	fmt.Println("intChan", num)

	go ChannelChange(intChan)
	num = <-intChan
	fmt.Println("intChan", num)

	go ChannelChange(intChan)
	num = <-intChan
	fmt.Println("intChan", num)
}

func ChannelChange(intChan chan int) {
	someNumber := helpers.GetRandomNumber(0, 100)
	intChan <- someNumber
}
