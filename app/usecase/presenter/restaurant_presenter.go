package presenter

import "github.com/HiroyukiNakatsuma/adviser-go/app/domain/model"

type RestaurantPresenter interface {
    BuildReplyContent(rests []*model.Restaurant) string
}
