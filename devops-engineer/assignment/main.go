package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	hostname, _ := os.Hostname()
	response := fmt.Sprintf("hello, my hostname is %v\n", hostname)
	fmt.Printf(response)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(response))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("serving: %v\n", err)
		os.Exit(1)
	}
}
