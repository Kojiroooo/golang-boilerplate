package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	})
	fmt.Println("[INFO] Server listening")
	http.ListenAndServe(":3000", nil)
}
