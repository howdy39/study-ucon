package main

import (
	"net/http"

	"fmt"

	"github.com/favclip/ucon"
)

func main() {
	ucon.Orthodox()

	ucon.Middleware(Logger)

	ucon.HandleFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	ucon.ListenAndServe(":8080")
}

func Logger(b *ucon.Bubble) error {
	fmt.Printf("Received: %s %s\n", b.R.Method, b.R.URL.String())
	return b.Next()
}
