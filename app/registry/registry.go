package registry

import (
    "github.com/HiroyukiNakatsuma/adviser-go/app/infrastructure/http/handler"
    "github.com/HiroyukiNakatsuma/adviser-go/app/infrastructure/http/external_interfaces"
    "github.com/HiroyukiNakatsuma/adviser-go/app/interface/controller"
    "github.com/HiroyukiNakatsuma/adviser-go/app/interface/presenters"
    "github.com/HiroyukiNakatsuma/adviser-go/app/usecase/external_interface"
    "github.com/HiroyukiNakatsuma/adviser-go/app/usecase/presenter"
    "github.com/HiroyukiNakatsuma/adviser-go/app/usecase/service"
)

func ResolveDependencies() *handler.AppHandler {
    return NewAppHandler()
}

func NewAppHandler() *handler.AppHandler {
    return handler.NewAppHandler(*NewLinebotHandler(), *NewHelloHandler(), *NewImageHandler())
}

func NewLinebotHandler() *handler.LinebotHandler {
    return handler.NewLinebotHandler(*NewLinebotController())
}

func NewHelloHandler() *handler.HelloHandler {
    return handler.NewHelloHandler()
}

func NewImageHandler() *handler.ImageHandler {
    return handler.NewImageHandler()
}

func NewLinebotController() *controller.LinebotController {
    return controller.NewLinebotController(*NewTextService(), *NewRestaurantService())
}

func NewTextService() *service.TextService {
    return service.NewTextService(NewTextPresenter())
}

func NewRestaurantService() *service.RestaurantService {
    return service.NewRestaurantService(NewGnavi())
}

func NewTextPresenter() presenter.TextPresenter {
    return presenters.NewTextPresenter()
}

func NewGnavi() external_interface.RestaurantExternalInterface {
    return external_interfaces.NewGnavi()
}
