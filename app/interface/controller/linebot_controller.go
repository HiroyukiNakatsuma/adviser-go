package controller

import (
    "github.com/HiroyukiNakatsuma/adviser-go/app/usecase/service"

    "github.com/line/line-bot-sdk-go/linebot"
)

type LinebotController struct {
    txtServ  service.TextService
    restServ service.RestaurantService
}

func NewLinebotController(txtServ service.TextService, restServ service.RestaurantService) *LinebotController {
    return &LinebotController{txtServ, restServ}
}

func (linebotController *LinebotController) Reply(event *linebot.Event, profile *linebot.UserProfileResponse) (replyContent string) {
    if event.Type == linebot.EventTypeMessage {
        switch message := event.Message.(type) {
        case *linebot.TextMessage:
            // txtServ := service.NewTextService(presenter.NewTextPresenter())
            replyContent = linebotController.txtServ.ReplyContent4PlaneMessage(message.Text, profile.DisplayName)
        case *linebot.LocationMessage:
            // restServ := service.NewRestaurantService(external_service.NewGnavi(), presenter.NewRestaurantPresenter())
            replyContent = linebotController.restServ.ReplyContent4Location(message.Latitude, message.Longitude)
        }
    }

    return replyContent
}
