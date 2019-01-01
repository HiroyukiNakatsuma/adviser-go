package external_service

import "github.com/HiroyukiNakatsuma/adviser-go/app/domain/model"

type RestaurantExternalService interface {
    GetRestaurants(latitude float64, longitude float64, isNoSmoking bool) []*model.Restaurant
}
