package presenters

import (
    "github.com/HiroyukiNakatsuma/adviser-go/app/domain/model"
    "github.com/HiroyukiNakatsuma/adviser-go/app/usecase/presenter"
    "fmt"
)

const noContentMessage = "ごめんなさい。該当するコンテンツがありませんでした。。"
const gnaviCreditText = "Supported by ぐるなびWebService : https://api.gnavi.co.jp/api/scope/"

type restaurantPresenter struct{}

func NewRestaurantPresenter() presenter.RestaurantPresenter {
    return &restaurantPresenter{}
}

func (restaurantPresenter *restaurantPresenter) BuildReplyContent(rests []*model.Restaurant) (reply string) {
    if len(rests) == 0 {
        return fmt.Sprintf("%s\n\n%s", noContentMessage, gnaviCreditText)
    }

    for _, rest := range rests {
        reply += fmt.Sprintf("%s\n%s\n\n", rest.Name, rest.Url)
    }
    return reply + gnaviCreditText
}
