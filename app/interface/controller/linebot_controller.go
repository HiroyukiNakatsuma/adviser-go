package controller

import (
    "github.com/HiroyukiNakatsuma/adviser-go/app/usecase"

    "github.com/line/line-bot-sdk-go/linebot"
)

func Reply(event *linebot.Event) (replyContent string) {
    if event.Type == linebot.EventTypeMessage {
        switch message := event.Message.(type) {
        case *linebot.TextMessage:
            replyContent = usecase.ReplyEcho(message.Text)
        }
    }

    return replyContent
}
