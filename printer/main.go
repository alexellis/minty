package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
	fmt.Fprintf(os.Stdout, "Method: %s Path: %s \nHeaders:\n", r.Method, r.URL.Path)

	for k, v := range r.Header {
		fmt.Fprintf(os.Stdout, "%s=%s", k, v)
	}

	if r.Body != nil && r.ContentLength > 0 {
		defer r.Body.Close()
		fmt.Fprintf(os.Stdout, "\nBody (%d bytes):\n", r.ContentLength)

		io.Copy(os.Stdin, r.Body)
	}
	fmt.Fprintf(os.Stdout, "\n")

}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting server on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
