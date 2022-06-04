package main

import (
    "github.com/gin-gonic/gin"
)

func main(){
    r := gin.Default()
    r.GET("/healthz",HandleGet)
    r.Run(":9000")
}

func HandleGet(c *gin.Context){
    c.JSON(200,gin.H{
    "Header":c.Request.Header,
    })
}