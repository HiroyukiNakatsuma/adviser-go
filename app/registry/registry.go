package registry

import (
    "github.com/HiroyukiNakatsuma/adviser-go/app/infrastructure/api/handler"
    "github.com/HiroyukiNakatsuma/adviser-go/app/infrastructure/external_interfaces"
    "github.com/HiroyukiNakatsuma/adviser-go/app/interface/controller"
    "github.com/HiroyukiNakatsuma/adviser-go/app/interface/presenters"
    "github.com/HiroyukiNakatsuma/adviser-go/app/usecase/service"
)

func ResolveDependencies() *handler.AppHandler {
    return handler.NewAppHandler(
        *handler.NewLinebotHandler(
            *controller.NewLinebotController(
                *service.NewTextService(
                    presenters.NewTextPresenter()),
                *service.NewRestaurantService(
                    external_interfaces.NewGnavi(),
                    presenters.NewRestaurantPresenter()))),
        *handler.NewHelloHandler())
}
