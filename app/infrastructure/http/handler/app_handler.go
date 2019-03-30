package handler

type AppHandler struct {
    LinebotHandler
    HelloHandler
    ImageHandler
}

func NewAppHandler(linebotHandler LinebotHandler, helloHandler HelloHandler, imageHandler ImageHandler) *AppHandler {
    return &AppHandler{linebotHandler, helloHandler, imageHandler}
}
