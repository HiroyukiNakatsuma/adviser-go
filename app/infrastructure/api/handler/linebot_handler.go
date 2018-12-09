package handler

import (
    "github.com/line/line-bot-sdk-go/linebot"
    "net/http"
    "fmt"
    "log"
)

var cli *linebot.Client

// func init() {
//     bot, err := linebot.New(
//         os.Getenv("CHANNEL_SECRET"),
//         os.Getenv("CHANNEL_TOKEN"),
//     )
//     if err != nil {
//         log.Fatal(err)
//     }
//     cli = bot
// }

func LinebotHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Received Line message!")
    log.Print("Received Line message!")
    // events, err := cli.ParseRequest(r)
    // if err != nil {
    //     if err == linebot.ErrInvalidSignature {
    //         w.WriteHeader(400)
    //     } else {
    //         w.WriteHeader(500)
    //     }
    //     return
    // }
    //
    // for _, event := range events {
    //     if event.Type == linebot.EventTypeMessage {
    //         switch message := event.Message.(type) {
    //         case *linebot.TextMessage:
    //             if _, err = cli.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
    //                 log.Print(err)
    //             }
    //         }
    //     }
    // }
}
