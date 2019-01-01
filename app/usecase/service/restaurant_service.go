package service

import (
    "github.com/HiroyukiNakatsuma/adviser-go/app/usecase/external_service"
    "github.com/HiroyukiNakatsuma/adviser-go/app/usecase/presenter"
    "github.com/HiroyukiNakatsuma/adviser-go/app/domain/model"
)

type restaurantService struct {
    restExServ external_service.RestaurantExternalService
    restPres   presenter.RestaurantPresenter
}

func NewRestaurantService(restRepo external_service.RestaurantExternalService, restPres presenter.RestaurantPresenter) *restaurantService {
    return &restaurantService{restRepo, restPres}
}

func (restServ *restaurantService) GetRestaurants(latitude float64, longitude float64, isNoSmoking bool) []*model.Restaurant {
    return restServ.restExServ.GetRestaurants(latitude, longitude, isNoSmoking)
}

func (restServ *restaurantService) BuildReplyContent(rests []*model.Restaurant) string {
    return restServ.restPres.BuildReplyContent(rests)
}

func (restServ *restaurantService) getRestaurants(latitude float64, longitude float64, isNoSmoking bool) string {
    restaurants := restServ.GetRestaurants(latitude, longitude, isNoSmoking)
    return restServ.BuildReplyContent(restaurants)
}

func (restServ *restaurantService) getNoSmokingRestaurants(latitude float64, longitude float64) string {
    return restServ.getRestaurants(latitude, longitude, true)
}

func (restServ *restaurantService) ReplyContent4Location(latitude float64, longitude float64) string {
    return restServ.getNoSmokingRestaurants(latitude, longitude)
}
