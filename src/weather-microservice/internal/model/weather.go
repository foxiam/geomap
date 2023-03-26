package model

type Weather struct {
	Id                int         `json:"id"`
	City              string      `json:"city"`
	Coordinates       coordinates `json:"coordinates"`
	AverageSpringTemp float32     `json:"average_spring_temp"`
	AverageSummerTemp float32     `json:"average_summer_temp"`
	AverageAutumnTemp float32     `json:"average_autumn_temp"`
	AverageWinterTemp float32     `json:"average_winter_temp"`
	Humidity          int         `json:"humidity"`
	Score             int         `json:"score"`
	Green             int         `json:"green"`
	TypeCity          string      `json:"type_city"`
	Climate           string      `json:"climate"`
	Population        string      `json:"population"`
}

type coordinates struct {
	Lat float32 `json:"lat"`
	Lng float32 `json:"lng"`
}

type WeatherFilter struct {
	SpringTempMore float32 `json:"spring_temp_more"`
	SpringTempLess float32 `json:"spring_temp_less"`
	SummerTempMore float32 `json:"summer_temp_more"`
	SummerTempLess float32 `json:"summer_temp_less"`
	AutumnTempMore float32 `json:"autumn_temp_more"`
	AutumnTempLess float32 `json:"autumn_temp_less"`
	WinterTempMore float32 `json:"winter_temp_more"`
	WinterTempLess float32 `json:"winter_temp_less"`
	HumidityMore   int     `json:"humidity_more"`
	HumidityLess   int     `json:"humidity_less"`
	ScoreMore      int     `json:"score_more"`
	ScoreLess      int     `json:"score_less"`
	GreenMore      int     `json:"green_more"`
	GreenLess      int     `json:"green_less"`
}
