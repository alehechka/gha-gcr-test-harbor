package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
    http.HandleFunc("/", HelloServer)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

    http.ListenAndServe(":" + PORT, nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}