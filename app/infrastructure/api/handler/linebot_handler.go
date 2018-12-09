package handler

import (
    "net/http"
    "os"
    "log"

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

    for _, event := range events {
        if event.Type == linebot.EventTypeMessage {
            switch message := event.Message.(type) {
            case *linebot.TextMessage:
                if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
                    log.Print(err)
                }
            }
        }
    }
}
