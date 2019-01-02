package handler

import (
    "net/http"
    "log"
    "fmt"
)

type HelloHandler struct{}

func NewHelloHandler() *HelloHandler {
    return &HelloHandler{}
}

func (helloHandler *HelloHandler) Handle(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    log.Printf("Start \"/%s\"", r.URL.Path[1:])
    fmt.Fprintf(w, "Welcome to LINE BOT app!!")
}
