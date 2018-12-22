package controller

import (
    "github.com/HiroyukiNakatsuma/adviser-go/app/usecase"

    "github.com/line/line-bot-sdk-go/linebot"
)

func Reply(event *linebot.Event, profile *linebot.UserProfileResponse) (replyContent string) {
    if event.Type == linebot.EventTypeMessage {
        switch message := event.Message.(type) {
        case *linebot.TextMessage:
            replyContent = usecase.ReplyContent(message.Text, profile.DisplayName)
        }
    }

    return replyContent
}
