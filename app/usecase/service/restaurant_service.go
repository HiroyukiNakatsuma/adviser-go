package service

import (
    "github.com/HiroyukiNakatsuma/adviser-go/app/usecase/external_interface"
    "github.com/HiroyukiNakatsuma/adviser-go/app/usecase/presenter"
    "github.com/HiroyukiNakatsuma/adviser-go/app/domain/model"
)

const isLunch = true
const isNoSmoking = true

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

func (restServ *RestaurantService) getRestaurants(latitude float64, longitude float64, isLunch bool, isNoSmoking bool) string {
    restaurants := restServ.restExServ.GetRestaurants(latitude, longitude, isLunch, isNoSmoking)
    return restServ.BuildReplyContent(restaurants)
}

func (restServ *RestaurantService) ReplyContent4Location(latitude float64, longitude float64) string {
    return restServ.getRestaurants(latitude, longitude, isLunch, isNoSmoking)
}
