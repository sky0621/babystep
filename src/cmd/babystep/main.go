package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, "<h1>Hello, World</h1>"); err != nil {
			panic(err)
		}
	})
	if err := http.ListenAndServe(":9090", nil); err != nil {
		panic(err)
	}
}
