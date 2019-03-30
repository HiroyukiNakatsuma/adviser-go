package handler

import (
    "net/http"
    "flag"
)

type ImageHandler struct{}

func NewImageHandler() *ImageHandler {
    return &ImageHandler{}
}

func (imageHandler *ImageHandler) Handle() http.Handler {
    var root = flag.String("root", ".", "file system path")
    return http.FileServer(http.Dir(*root))
}
