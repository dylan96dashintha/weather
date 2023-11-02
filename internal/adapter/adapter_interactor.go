package adapter

import (
	"context"
	"github.com/weather/internal/domain"
)

type MapAdapter interface {
	GetCoordinatesFromCity(ctx context.Context,
		city string) (resp domain.AddressValidatorResponse, err error)
}

type GeoJsonAdapter interface {
	GetWeatherForecast(ctx context.Context,
		lat, lon float64) (resp domain.WeatherForecastResponse, err error)
}
