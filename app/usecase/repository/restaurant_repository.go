package repository

import "github.com/HiroyukiNakatsuma/adviser-go/app/domain/model"

type RestaurantRepository interface {
    GetRestaurants(latitude float64, longitude float64, isNoSmoking bool) []*model.Restaurant
}
