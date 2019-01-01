package controller

import (
    "github.com/HiroyukiNakatsuma/adviser-go/app/infrastructure/external_service"
    "github.com/HiroyukiNakatsuma/adviser-go/app/usecase/service"

    "github.com/line/line-bot-sdk-go/linebot"
)

func Reply(event *linebot.Event, profile *linebot.UserProfileResponse) (replyContent string) {
    if event.Type == linebot.EventTypeMessage {
        switch message := event.Message.(type) {
        case *linebot.TextMessage:
            replyContent = service.ReplyContent4PlaneMessage(message.Text, profile.DisplayName)
        case *linebot.LocationMessage:
            restServ := service.NewRestaurantService(external_service.NewGnavi())
            replyContent = restServ.ReplyContent4Location(message.Latitude, message.Longitude)
        }
    }

    return replyContent
}
