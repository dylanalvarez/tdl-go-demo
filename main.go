package main

import (
    "example/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    engine := gin.Default()
    routes.Setup(engine)
    engine.Run()
}
