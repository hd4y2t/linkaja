package main

import (
	"github.com/gin-gonic/gin"
	accountcontroller "github.com/hd4y2t/go_linkaja/controllers"
	"github.com/hd4y2t/go_linkaja/models"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/account/", accountcontroller.Index)
	r.GET("/api/account/:account_number", accountcontroller.Show)
	r.POST("/api/account/:account_number/transfer", accountcontroller.Update)

	r.Run(":8001")
}
