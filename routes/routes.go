package routes

import (
    "example/handlers"
    "github.com/gin-gonic/gin"
)

func Setup(engine *gin.Engine) {
  engine.GET("/ping", handlers.PingHandler)
  engine.GET("/series", handlers.SeriesHandler)
}
