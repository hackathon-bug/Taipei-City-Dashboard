package controllers

import (
	"TaipeiCityDashboardBE/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFirstWeatherForecast(c *gin.Context){
	var weatherForecast models.WeatherForecast;
	var err error
	if weatherForecast, err = models.GetFirstWeatherForecast(); err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": weatherForecast})
}