package controller

import (
    "github.com/HiroyukiNakatsuma/adviser-go/app/usecase/service"
    "github.com/HiroyukiNakatsuma/adviser-go/app/interface/presenters"

    "github.com/line/line-bot-sdk-go/linebot"
)

type LinebotController struct {
    txtServ  service.TextService
    restServ service.RestaurantService
}

func NewLinebotController(txtServ service.TextService, restServ service.RestaurantService) *LinebotController {
    return &LinebotController{txtServ, restServ}
}

func (linebotController *LinebotController) Reply(event *linebot.Event, profile *linebot.UserProfileResponse) (reply linebot.SendingMessage) {
    if event.Type == linebot.EventTypeMessage {
        switch message := event.Message.(type) {
        case *linebot.TextMessage:
            replyText := linebotController.txtServ.ReplyContent4PlaneMessage(message.Text, profile.DisplayName)
            reply = linebot.NewTextMessage(replyText)
        case *linebot.LocationMessage:
            restaurants := linebotController.restServ.GetRestaurantsByLocation(message.Latitude, message.Longitude)
            reply = presenters.NewRestaurantPresenter().BuildReplyContent(restaurants)
        }
    }
    return reply
}
