package models

import "time"

type WeatherForecast struct {
	ID         int     `json:"id" gorm:"column:id;autoincrement;primaryKey"`
	City       string  `json:"city" gorm:"column:city;"`
	District   string  `json:"district" gorm:"column:district"`
	Longitude  float64 `json:"longitude" gorm:"column:longitude"`
	Latitude   float64 `json:"latitude" gorm:"column:latitude"`
	StartTime time.Time `json:"start_time" gorm:"column:start_time"`
	EndTime time.Time `json:"end_time" gorm:"column:end_time"`
	Temperature string `json:"temperature" gorm:"column:temperature"`
	MaxTemperature string `json:"max_temperature" gorm:"column:max_temperature"`
	MinTemperature string `json:"min_temperature" gorm:"column:min_temperature"`
	DewPoint string `json:"dew_point" gorm:"column:dew_point"`
	RelativeHumidity string `json:"relative_humidity" gorm:"column:relative_humidity"`
	MaxApparentTemperature string `json:"max_apparent_temperature" gorm:"column:max_apparent_temperature"`
	MinApparentTemperature string `json:"min_apparent_temperature" gorm:"column:min_apparent_temperature"`
	MaxComfortIndex string `json:"max_comfort_index" gorm:"column:max_comfort_index"`
	MinComfortIndex string `json:"min_comfort_index" gorm:"column:min_comfort_index"`
	WindSpeed int `json:"wind_speed" gorm:"column:wind_speed"`
	WindDirection string `json:"wind_direction" gorm:"column:wind_direction"`
	ProbabilityOfPrecipitation string `json:"probability_of_precipitation" gorm:"column:probability_of_precipitation"`
	Weather string `json:"weather" gorm:"column:weather"`
	UvIndex int `json:"uv_index" gorm:"column:uv_index"`
	WeatherDescription string `json:"weather_description" gorm:"column:weather_description"`
	Year int `json:"year" gorm:"column:year"`
	Month int `json:"month" gorm:"column:month"`
	Day int `json:"day" gorm:"column:day"`
	Hour int `json:"hour" gorm:"column:hour"`
}

func (WeatherForecast) TableName() string {
	return "weather_forecast"
}
func GetFirstWeatherForecast() (WeatherForecast,error){
	var weatherForecast WeatherForecast
	if err := DBDashboard.Find(&weatherForecast).First(&weatherForecast).Error; err!=nil{
		return weatherForecast, err
	}
	return weatherForecast, nil
}