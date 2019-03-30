package presenters

import (
    "github.com/HiroyukiNakatsuma/adviser-go/app/domain/model"

    "github.com/line/line-bot-sdk-go/linebot"
)

const noContentMessage = "ごめんなさい。該当するレストランがありませんでした。。"
const noImageUrl = "https://adviser-go.herokuapp.com/public/images/noImage.jpg"
const gnaviCreditText = "Supported by ぐるなびWebService"
const altText = "レストラン情報を送信しました"
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
        columns = append(columns, linebot.NewCarouselColumn(restaurantPresenter.imageUrl(rest.ImageUrls), rest.Name, gnaviCreditText, actions))
    }

    return linebot.NewTemplateMessage(altText, linebot.NewCarouselTemplate(columns...))
}

func (restaurantPresenter *RestaurantPresenter) imageUrl(urls []string) string {
    for _, url := range urls {
        if url != "" {
            return url
        }
    }
    return noImageUrl
}
