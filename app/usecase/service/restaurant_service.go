package service

import (
    "github.com/HiroyukiNakatsuma/adviser-go/app/usecase/external_interface"
    "github.com/HiroyukiNakatsuma/adviser-go/app/usecase/presenter"
    "github.com/HiroyukiNakatsuma/adviser-go/app/domain/model"
    "math/rand"
)

const isLunch = true
const isNoSmoking = true
const restaurantSearchAmount = 20

type RestaurantService struct {
    restExServ external_interface.RestaurantExternalInterface
    restPres   presenter.RestaurantPresenter
}

func NewRestaurantService(restExServ external_interface.RestaurantExternalInterface, restPres presenter.RestaurantPresenter) *RestaurantService {
    return &RestaurantService{restExServ, restPres}
}

func (restServ *RestaurantService) BuildReplyContent(rests []*model.Restaurant) string {
    return restServ.restPres.BuildReplyContent(rests)
}

func (restServ *RestaurantService) ReplyContentByLocation(latitude float64, longitude float64) string {
    return restServ.getRestaurants(latitude, longitude, isLunch, isNoSmoking, restaurantSearchAmount)
}

func (restServ *RestaurantService) getRestaurants(latitude float64, longitude float64, isLunch bool, isNoSmoking bool, amount int) string {
    externalRestaurants := restServ.restExServ.GetRestaurants(latitude, longitude, isLunch, isNoSmoking, amount)
    rests := restServ.extractThreeRandomly(externalRestaurants)
    return restServ.BuildReplyContent(rests)
}

func (restServ *RestaurantService) extractThreeRandomly(rests []*model.Restaurant) []*model.Restaurant {
    shuffle(rests)
    return rests[:2]
}

/**
Fisher–Yates shuffle
refs: https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
 */
func shuffle(data []*model.Restaurant) {
    n := len(data)
    for i := n - 1; i >= 0; i-- {
        j := rand.Intn(i + 1)
        data[i], data[j] = data[j], data[i]
    }
}
