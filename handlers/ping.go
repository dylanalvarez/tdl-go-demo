package handlers

import "github.com/gin-gonic/gin"

func PingHandler(context *gin.Context) {
    context.JSON(200, gin.H{
        "message": "pong",
    })
}
