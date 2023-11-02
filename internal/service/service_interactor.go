package service

import "context"

type WeatherForecast interface {
	GetWeatherReport(ctx context.Context, city string) (err error)
}
