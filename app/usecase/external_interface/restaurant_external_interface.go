package external_interface

import "github.com/HiroyukiNakatsuma/adviser-go/app/domain/model"

type RestaurantExternalInterface interface {
    GetRestaurants(latitude float64, longitude float64, isNoSmoking bool) []*model.Restaurant
}