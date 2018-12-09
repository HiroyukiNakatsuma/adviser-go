package controller

import (
    "log"

    "github.com/HiroyukiNakatsuma/adviser-go/app/infrastructure/api/handler"

    "github.com/line/line-bot-sdk-go/linebot"
)

func Reply(events []*linebot.Event) {
    for _, event := range events {
        var replyMessage string
        if event.Type == linebot.EventTypeMessage {
            switch message := event.Message.(type) {
            case *linebot.TextMessage:
                replyMessage = message.Text
            }
        }
        if _, err := handler.Cli.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
            log.Print(err)
        }
    }
}
