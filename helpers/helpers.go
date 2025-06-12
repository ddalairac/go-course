package helpers

import (
	"fmt"
	"math/rand"
)

func PrintTitle(message string) {
	fmt.Println("")
	fmt.Println("**************************************************")
	fmt.Println(message)
	fmt.Println("**************************************************")
	fmt.Println("")
}

func PrintSubTitle(message string) {
	fmt.Println(" ")
	fmt.Println(message)
	fmt.Println("-------------------------------------------------")
}

func GetRandomNumber(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func GetRandomString(length int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
