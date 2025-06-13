package data

import (
	"fmt"
	"net/http"
)

func HttpHandler() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, World!")
		if err != nil {
			fmt.Println("Error writing to response:", err)
		}
		fmt.Println("Number of bytes written:", n)
	})
	_ = http.ListenAndServe(":8080", nil)
	fmt.Println("Server is running on port 8080")
}
