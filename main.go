package main

import (
    "github.com/HiroyukiNakatsuma/adviser-go/app/infrastructure/api/router"
    "github.com/HiroyukiNakatsuma/adviser-go/app/registry"
)

func main() {
    r := *registry.ResolveDependencies()
    router.Run(r)
}
