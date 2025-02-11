package station

import (
	"gin-rest-api/common/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initiate(router *gin.RouterGroup) {
	stationService := NewService()

	station := router.Group("/stations")
	station.GET("", func(c *gin.Context) {
		GetAllStation(c, stationService)
	})
}

func GetAllStation(c *gin.Context, service Service) {
	// layer servcie
	datas, err := service.GetAllStation()
	if err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	c.JSON(http.StatusOK, response.APIResponse{
		Success: true,
		Message: "Success get all station",
		Data:    datas,
	})

}
