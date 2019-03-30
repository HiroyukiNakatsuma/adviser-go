package service

import (
    "github.com/HiroyukiNakatsuma/adviser-go/app/usecase/external_interface"
    "github.com/HiroyukiNakatsuma/adviser-go/app/domain/model"
    "math/rand"
    "time"
)

const isLunch = true
const isNoSmoking = true
const restaurantSearchAmount = 30

type RestaurantService struct {
    restExServ external_interface.RestaurantExternalInterface
}

func NewRestaurantService(restExServ external_interface.RestaurantExternalInterface) *RestaurantService {
    return &RestaurantService{restExServ}
}

func (restServ *RestaurantService) GetRestaurantsByLocation(latitude float64, longitude float64) []*model.Restaurant {
    return restServ.getRestaurants(latitude, longitude, isLunch, isNoSmoking, restaurantSearchAmount)
}

func (restServ *RestaurantService) getRestaurants(latitude float64, longitude float64, isLunch bool, isNoSmoking bool, amount int) []*model.Restaurant {
    rests := restServ.restExServ.GetRestaurants(latitude, longitude, isLunch, isNoSmoking, amount)
    if len(rests) > 3 {
        rests = restServ.extractThreeRandomly(rests)
    }
    return rests
}

func (restServ *RestaurantService) extractThreeRandomly(rests []*model.Restaurant) []*model.Restaurant {
    shuffle(rests)
    return rests[:3]
}

/**
Fisher–Yates shuffle
refs: https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
 */
func shuffle(data []*model.Restaurant) {
    n := len(data)
    for i := n - 1; i >= 0; i-- {
        rand.Seed(time.Now().UnixNano())
        j := rand.Intn(i + 1)
        data[i], data[j] = data[j], data[i]
    }
}
