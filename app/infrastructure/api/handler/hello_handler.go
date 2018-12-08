package handler

import (
    "net/http"
    "fmt"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", r.URL.Path[1:])
}
