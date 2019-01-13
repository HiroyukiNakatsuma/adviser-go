package handler

type AppHandler struct {
    LinebotHandler
    HelloHandler
}

func NewAppHandler(linebotHandler LinebotHandler, helloHandler HelloHandler) *AppHandler {
    return &AppHandler{linebotHandler, helloHandler}
}
