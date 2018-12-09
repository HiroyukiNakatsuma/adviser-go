package controller

import (
    "log"

    "github.com/line/line-bot-sdk-go/linebot"
)

func Reply(cli *linebot.Client, events []*linebot.Event) {
    for _, event := range events {
        var replyMessage string
        if event.Type == linebot.EventTypeMessage {
            switch message := event.Message.(type) {
            case *linebot.TextMessage:
                replyMessage = message.Text
            }
        }
        if _, err := cli.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
            log.Print(err)
        }
    }
}
