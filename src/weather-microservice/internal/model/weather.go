package model

type Weather struct {
	City              string `json:"city"`
	AverageSpringTemp int    `json:"average_spring_temp"`
	AverageSummerTemp int    `json:"average_summer_temp"`
	AverageAutumnTemp int    `json:"average_autumn_temp"`
	AverageWinterTemp int    `json:"average_winter_temp"`
	AirPollution      int    `json:"air_pollution"`
	Humidity          int    `json:"humidity"`
}
type coordinates struct {
	Lat float32 `json:"lat"`
	Lng float32 `json:"lng"`
}

type City struct {
	Id          int         `json:"id"`
	City        string      `json:"city"`
	Coordinates coordinates `json:"coordinates"`
}

type WeatherFilter struct {
	SpringTempMore   int `json:"spring_temp_more"`
	SpringTempLess   int `json:"spring_temp_less"`
	SummerTempMore   int `json:"summer_temp_more"`
	SummerTempLess   int `json:"summer_temp_less"`
	AutumnTempMore   int `json:"autumn_temp_more"`
	AutumnTempLess   int `json:"autumn_temp_less"`
	WinterTempMore   int `json:"winter_temp_more"`
	WinterTempLess   int `json:"winter_temp_less"`
	AirPollutionMore int `json:"air_pollution_more"`
	AirPollutionLess int `json:"air_pollution_less"`
	HumidityMore     int `json:"humidity_more"`
	HumidityLess     int `json:"humidity_less"`
}
