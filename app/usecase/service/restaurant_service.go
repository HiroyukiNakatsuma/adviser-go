package service

import (
    "fmt"

    "github.com/HiroyukiNakatsuma/adviser-go/app/usecase/external_service"
    "github.com/HiroyukiNakatsuma/adviser-go/app/domain/model"
)

const noContentMessage = "ごめんなさい。該当するコンテンツがありませんでした。。"
const gnaviCreditText = "Supported by ぐるなびWebService : https://api.gnavi.co.jp/api/scope/"

type restaurantService struct {
    restExServ external_service.RestaurantExternalService
}

func NewRestaurantService(restRepo external_service.RestaurantExternalService) *restaurantService {
    return &restaurantService{restRepo}
}

func (restServ *restaurantService) GetRestaurants(latitude float64, longitude float64, isNoSmoking bool) []*model.Restaurant {
    return restServ.restExServ.GetRestaurants(latitude, longitude, isNoSmoking)
}

func (restServ *restaurantService) getRestaurants(latitude float64, longitude float64, isNoSmoking bool) (reply string) {
    restaurants := restServ.GetRestaurants(latitude, longitude, isNoSmoking)

    for _, rest := range restaurants {
        reply += fmt.Sprintf("%s\n%s\n\n", rest.Name, rest.Url)
    }

    if len(restaurants) == 0 {
        reply = fmt.Sprintf("%s\n\n", noContentMessage)
    }

    return reply + gnaviCreditText
}

func (restServ *restaurantService) getNoSmokingRestaurants(latitude float64, longitude float64) (reply string) {
    return restServ.getRestaurants(latitude, longitude, true)
}

func (restServ *restaurantService) ReplyContent4Location(latitude float64, longitude float64) (reply string) {
    return restServ.getNoSmokingRestaurants(latitude, longitude)
}
