package registry

import (
    "github.com/HiroyukiNakatsuma/adviser-go/app/infrastructure/api/handler"
    "github.com/HiroyukiNakatsuma/adviser-go/app/infrastructure/external_service"
    "github.com/HiroyukiNakatsuma/adviser-go/app/interface/controller"
    "github.com/HiroyukiNakatsuma/adviser-go/app/interface/presenter"
    "github.com/HiroyukiNakatsuma/adviser-go/app/usecase/service"
)

func ResolveDependencies() *handler.AppHandler {
    return handler.NewAppHandler(
        *handler.NewLinebotHandler(
            *controller.NewLinebotController(
                *service.NewTextService(
                    presenter.NewTextPresenter()),
                *service.NewRestaurantService(
                    external_service.NewGnavi(),
                    presenter.NewRestaurantPresenter()))),
        *handler.NewHelloHandler())
}
