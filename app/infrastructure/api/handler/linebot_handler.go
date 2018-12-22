package handler

import (
    "net/http"
    "os"
    "log"

    "github.com/HiroyukiNakatsuma/adviser-go/app/interface/controller"

    "github.com/line/line-bot-sdk-go/linebot"
)

func getClient() (bot *linebot.Client) {
    bot, err := linebot.New(
        os.Getenv("CHANNEL_SECRET_ADVISER"),
        os.Getenv("CHANNEL_TOKEN_ADVISER"),
    )
    if err != nil {
        log.Fatal(err)
    }
    return
}

func getUserProfile(src *linebot.EventSource) (res *linebot.UserProfileResponse) {
    bot := getClient()
    var err error
    if len(src.GroupID) != 0 {
        res, err = bot.GetGroupMemberProfile(src.GroupID, src.UserID).Do()
    } else if len(src.RoomID) != 0 {
        res, err = bot.GetRoomMemberProfile(src.RoomID, src.UserID).Do()
    } else {
        res, err = bot.GetProfile(src.UserID).Do()
    }
    if err != nil {
        log.Print(err)
    }
    log.Printf("Get User. userId: %s", res.UserID)
    return
}

func LinebotHandler(w http.ResponseWriter, r *http.Request) {
    log.Printf("Start \"/%s\"", r.URL.Path[1:])

    bot := getClient()
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
        profile := getUserProfile(event.Source)
        replyContent = controller.Reply(event, profile)

        if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyContent)).Do(); err != nil {
            log.Print(err)
        }
    }
}
