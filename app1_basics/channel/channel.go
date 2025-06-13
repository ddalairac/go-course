/*
Goroutines and Channels in Go

Goroutines:
- Lightweight threads managed by Go runtime
- Much more efficient than OS threads (few KB vs 1MB per thread)
- Can run thousands or millions simultaneously
- Managed by Go's scheduler which multiplexes them onto OS threads

Channels:
- Primary mechanism for communication between goroutines
- Provide synchronization and data sharing
- Can be buffered or unbuffered
- Blocking operations: send blocks until receive, receive blocks until send
- Thread-safe by design

This example demonstrates:
- Creating goroutines with the 'go' keyword
- Using channels 'chan' for communication between goroutines
- Basic channel operations (send and receive)
- Buffered channels
- Channel closing and range iteration
- Select statement for handling multiple channels

Best Practices:
- Always use defer to close channels
- Benefits of using defer:
  1. Guarantees channel closure even if the function panics
  2. Ensures proper cleanup in case of early returns
  3. Makes the code more maintainable by keeping cleanup close to initialization
  4. Follows the principle of "defensive programming"
  5. Prevents resource leaks in error scenarios
*/

// Package channel demonstrates the usage of Go channels and goroutines
package channel

import (
	"fmt"
	"time"

	"github.com/diegodalairac/go-course/app1_basics/helpers"
)

/*
* Demonstrates basic channel operations with goroutines
It creates a channel, spawns multiple goroutines to send data,
and receives the data from the channel
*/
func Channel(intChan chan int) {
	// Spawn a goroutine to send a random number through the channel
	go ChannelChange(intChan)
	// Block and wait to receive the number from the channel
	num := <-intChan
	fmt.Println("intChan", num)

	// Spawn another goroutine to send a random number
	go ChannelChange(intChan)
	// Receive the second number
	num = <-intChan
	fmt.Println("intChan", num)

	// Spawn a third goroutine to send a random number
	go ChannelChange(intChan)
	num = <-intChan
	fmt.Println("intChan", num)

	// Demonstrate buffered channel
	fmt.Println("\nBuffered Channel Example:")
	bufferedChan := make(chan int, 2) // Channel with buffer size of 2
	bufferedChan <- 1                 // This won't block
	bufferedChan <- 2                 // This won't block
	fmt.Println(<-bufferedChan)       // Receive 1
	fmt.Println(<-bufferedChan)       // Receive 2

	// Demonstrate channel closing and range
	fmt.Println("\nChannel Closing Example:")
	numbersChan := make(chan int)

	// Using defer to ensure channel is closed when the goroutine exits
	// This is better than closing at the end because:
	// 1. It will close even if the goroutine panics
	// 2. It will close if there's an early return
	// 3. The cleanup code is right next to where the channel is used
	go func() {
		defer close(numbersChan) // Will be called when the goroutine exits
		for i := 1; i <= 3; i++ {
			numbersChan <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Range over channel until it's closed
	// The range loop will automatically exit when the channel is closed
	// This is a common pattern for consuming values from a channel
	for num := range numbersChan {
		fmt.Println("Received:", num)
	}

	// Demonstrate select with multiple channels
	fmt.Println("\nSelect Example:")
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Using defer to ensure channels are closed
	// This is important because:
	// 1. Multiple channels need to be closed
	// 2. We want to ensure cleanup even if the function exits early
	// 3. It's cleaner than closing channels at the end of the function
	defer close(ch1)
	defer close(ch2)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- 2
	}()

	// Select will choose the first channel that's ready
	// It's like a switch statement but for channels
	for i := 0; i < 2; i++ {
		select {
		case x := <-ch1:
			fmt.Println("Received from ch1:", x)
		case x := <-ch2:
			fmt.Println("Received from ch2:", x)
		case <-time.After(300 * time.Millisecond):
			fmt.Println("Timeout!")
		}
	}
}

// ChannelChange is a goroutine that generates a random number
// and sends it through the provided channel
func ChannelChange(intChan chan int) {
	// Generate a random number between 0 and 100
	someNumber := helpers.GetRandomNumber(0, 100)
	// Send the random number through the channel
	intChan <- someNumber
}
