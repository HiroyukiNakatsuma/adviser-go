package router

import (
    "os"
    "log"
    "net/http"

    "github.com/HiroyukiNakatsuma/adviser-go/app/infrastructure/api/handler"
)

func Run(appHandler handler.AppHandler) {
    port := os.Getenv("PORT")
    if port == "" {
        log.Fatal("PORT must be set")
    }

    http.HandleFunc("/linebot/message", appHandler.LinebotHandler.Handle)
    http.HandleFunc("/", appHandler.HelloHandler.Handle)
    http.ListenAndServe(":"+port, nil)
}
