package station
func Initiate(router *gin.Engine) {
	
	station := router.Group("/v1/api/station")
	station.GET("", func(c *gin.Context) {
		
	})

	router.Run(":8080")
}