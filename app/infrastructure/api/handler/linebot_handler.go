package handler

import (
    "net/http"
    "os"
    "log"

    "github.com/HiroyukiNakatsuma/adviser-go/app/interface/controller"

    "github.com/line/line-bot-sdk-go/linebot"
)

func LinebotHandler(w http.ResponseWriter, r *http.Request) {
    log.Printf("Start \"/%s\"", r.URL.Path[1:])

    bot, err := linebot.New(
        os.Getenv("CHANNEL_SECRET_ADVISER"),
        os.Getenv("CHANNEL_TOKEN_ADVISER"),
    )
    if err != nil {
        log.Fatal(err)
    }

    events, err := bot.ParseRequest(r)
    if err != nil {
        if err == linebot.ErrInvalidSignature {
            w.WriteHeader(400)
        } else {
            w.WriteHeader(500)
        }
        return
    }

    var replyContent string
    for _, event := range events {

        replyContent = controller.Reply(event)

        if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyContent)).Do(); err != nil {
            log.Print(err)
        }
    }
}
