package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Серыер запушен на поту: 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
