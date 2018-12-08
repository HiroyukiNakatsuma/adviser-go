package api

import (
    "os"
    "log"
    "net/http"

    "github.com/HiroyukiNakatsuma/adviser-go/app/infrastructure/api/handler"
)

func Run() {
    port := os.Getenv("PORT")
    if port == "" {
        log.Fatal("PORT must be set")
    }

    http.HandleFunc("/", handler.HelloHandler)
    http.ListenAndServe(":"+port, nil)

}
