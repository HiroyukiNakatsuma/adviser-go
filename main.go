package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", r.URL.Path[1:])
}

func main() {
    port := os.Getenv("ADVISER_GO_PORT")
    if port == "" {
        log.Fatal("ADVISER_GO_PORT must be set")
    }

    http.HandleFunc("/", handler)
    http.ListenAndServe(":"+port, nil)
}
