package handler

import (
    "net/http"
    "log"
    "os"

    "github.com/HiroyukiNakatsuma/adviser-go/app/interface/controller"

    "github.com/line/line-bot-sdk-go/linebot"
)

var Cli *linebot.Client

func init() {
    bot, err := linebot.New(
        os.Getenv("CHANNEL_SECRET_ADVISER"),
        os.Getenv("CHANNEL_TOKEN_ADVISER"),
    )
    if err != nil {
        log.Fatal(err)
    }
    Cli = bot
}

func LinebotHandler(w http.ResponseWriter, r *http.Request) {
    events, err := Cli.ParseRequest(r)
    if err != nil {
        if err == linebot.ErrInvalidSignature {
            w.WriteHeader(400)
        } else {
            w.WriteHeader(500)
        }
        return
    }

    controller.Reply(events)
}
