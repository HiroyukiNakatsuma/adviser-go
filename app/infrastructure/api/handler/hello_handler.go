package handler

import (
    "net/http"
    "log"
    "fmt"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    log.Printf("Start \"/%s\"", r.URL.Path[1:])
    fmt.Fprintf(w, "Welcome to LINE BOT app!!")
}
