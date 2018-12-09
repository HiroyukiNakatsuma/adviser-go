package handler

import (
    "net/http"
    "fmt"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    fmt.Fprintf(w, "Welcome to LINE BOT app!!")
}
