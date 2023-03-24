package repository

import (
	"context"
	"fmt"
	"weather-microservice/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type WeatherRepository struct {
	db *pgxpool.Pool
}

func NewWeatherRepository(db *pgxpool.Pool) *WeatherRepository {
	return &WeatherRepository{db: db}
}

func (r *WeatherRepository) GetPositions(ctx context.Context) ([]*model.City, error) {
	const query = "SELECT id, city, lat, lng FROM public.positions LIMIT 10"
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cities []*model.City
	for rows.Next() {
		var city model.City
		err = rows.Scan(
			&city.Id,
			&city.City,
			&city.Coordinates.Lat,
			&city.Coordinates.Lng)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		cities = append(cities, &city)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return cities, nil
}

func (r *WeatherRepository) GetAll(ctx context.Context) ([]*model.Weather, error) {
	const query = "SELECT * FROM public.weather"
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cities []*model.Weather
	for rows.Next() {
		var city model.Weather
		err = rows.Scan(
			&city.City,
			&city.AverageSpringTemp,
			&city.AverageSummerTemp,
			&city.AverageAutumnTemp,
			&city.AverageWinterTemp,
			&city.AirPollution,
			&city.Humidity)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		cities = append(cities, &city)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return cities, nil
}

func (r *WeatherRepository) GetByName(ctx context.Context, name string) (*model.Weather, error) {
	const query = "SELECT * FROM public.weather WHERE city = $1"

	var city model.Weather
	err := r.db.QueryRow(ctx, query, name).Scan(
		&city.City,
		&city.AverageSpringTemp,
		&city.AverageSummerTemp,
		&city.AverageAutumnTemp,
		&city.AverageWinterTemp,
		&city.AirPollution,
		&city.Humidity)
	return &city, err
}

func (r *WeatherRepository) FindAllByFilter(ctx context.Context, filter *model.WeatherFilter) ([]*model.Weather, error) {
	const query = "SELECT * FROM public.weather WHERE " +
		"average_spring_temp BETWEEN $1 AND $2 AND " +
		"average_summer_temp BETWEEN $3 AND $4 AND " +
		"average_autumn_temp BETWEEN $5 AND $6 AND " +
		"average_winter_temp BETWEEN $7 AND $8 AND " +
		"air_pollution BETWEEN $9 AND $10 AND " +
		"humidity BETWEEN $11 AND $12"

	rows, err := r.db.Query(ctx, query,
		filter.SpringTempMore,
		filter.SpringTempLess,
		filter.SummerTempMore,
		filter.SummerTempLess,
		filter.AutumnTempMore,
		filter.AutumnTempLess,
		filter.WinterTempMore,
		filter.WinterTempLess,
		filter.AirPollutionMore,
		filter.AirPollutionLess,
		filter.HumidityMore,
		filter.HumidityLess)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cities []*model.Weather
	for rows.Next() {
		var city model.Weather
		err = rows.Scan(
			&city.City,
			&city.AverageSpringTemp,
			&city.AverageSummerTemp,
			&city.AverageAutumnTemp,
			&city.AverageWinterTemp,
			&city.AirPollution,
			&city.Humidity)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		cities = append(cities, &city)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return cities, nil

}
