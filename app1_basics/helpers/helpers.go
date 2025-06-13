package helpers

import (
	"fmt"
	"math/rand"
)

/*
* PrintTitle prints a title with a separator
Parameters:
  - message: the title to print
*/
func PrintTitle(message string) {
	fmt.Println("")
	fmt.Println("**************************************************")
	fmt.Println(message)
	fmt.Println("**************************************************")
	fmt.Println("")
}

/*
* PrintSubTitle prints a subtitle with a separator
Parameters:
  - message: the subtitle to print
*/
func PrintSubTitle(message string) {
	fmt.Println(" ")
	fmt.Println(message)
	fmt.Println("-------------------------------------------------")
}

/*
* GetRandomNumber generates a random integer between min and max (inclusive)
Parameters:
  - min: the minimum value in the range
  - max: the maximum value in the range

Returns:
  - a random integer between min and max
*/
func GetRandomNumber(min, max int) int {
	return rand.Intn(max-min+1) + min
}

/*
* GetRandomString generates a random string of specified length
Parameters:
  - length: the desired length of the random string

Returns:
  - a random string containing only letters (a-z, A-Z)
*/
func GetRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
