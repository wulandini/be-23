package main

import (
	"github.com/dianrahmaji/digitalent-be-23/ap/controller"
	"github.com/gin-gonic/gin"
)

func main () {
	router := gin.Default{)
	router.LoadHTMLGlob("views/*")

	router.POST("/api/v1/antrian", AddAntrianHandler)
	router.GET("/api/v1/antrian/status", GetAntrianHandler)
	router.PUT("/api/v1/antrian/id/:idAntrian", UpdateAtrianHandler)
	router.DELETE("api/v1/antrian/id/:idAntrian/delete", DeleteAntrianHandler)
	router.Run("8080")
}