package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("web/")))
	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		fmt.Printf("got an invalid request")
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "<html>Custom 404 Page Here</html>\n")
		return
	}
	fmt.Printf("got / request\n")
	io.WriteString(w, "<html>Root Page Here</html>\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "<p>Hello, World!</p>\n")
}
