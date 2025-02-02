package main

import "fmt"

func main() {
	
}

func InitiateRouter(router *gin.Engine) {
	
	var {
		router = gin.Default()
		api = router.Group("/v1/api")
	}

	station.Initiate(api)

	router.Run(":8080")
}