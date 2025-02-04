package main

import (
	"gin-rest-api/modules/station" // Sesuaikan path dengan struktur folder kamu

	"github.com/gin-gonic/gin"
)

func main() {
	InitiateRouter()
}

func InitiateRouter() {

	var (
		router = gin.Default()           // Membuat router baru
		api    = router.Group("/v1/api") // Membuat grup /v1/api
	)

	station.Initiate(api)

	router.Run(":8080")
}
