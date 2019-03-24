package presenters

import (
    "github.com/HiroyukiNakatsuma/adviser-go/app/domain/model"

    "github.com/line/line-bot-sdk-go/linebot"
)

const noContentMessage = "ごめんなさい。該当するレストランがありませんでした。。"
const gnaviCreditText = "Supported by ぐるなびWebService"
const altText = "This is restaurant list."
const detailLabel = "詳細を見る"

type RestaurantPresenter struct{}

func NewRestaurantPresenter() *RestaurantPresenter {
    return &RestaurantPresenter{}
}

func (restaurantPresenter *RestaurantPresenter) BuildReplyContent(rests []*model.Restaurant) (reply linebot.SendingMessage) {
    if len(rests) == 0 {
        return linebot.NewTextMessage(noContentMessage)
    }

    var columns []*linebot.CarouselColumn
    for _, rest := range rests {
        actions := linebot.NewURIAction(detailLabel, rest.Url)
        columns = append(columns, linebot.NewCarouselColumn(rest.ImageUrl, rest.Name, gnaviCreditText, actions))
    }

    return linebot.NewTemplateMessage(altText, linebot.NewCarouselTemplate(columns...))
}
