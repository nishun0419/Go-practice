package main

import (
    "github.com/gin-gonic/gin"
    c "server/controllers"
)

func main() {
        r := gin.Default()
        r.LoadHTMLGlob("views/*")
        r.GET("/", c.GetUsers)
        r.POST("/user/post", c.PostUser)
        r.GET("/user/delete/:id", c.DeleteUser)
        r.Run(":8080") 
}