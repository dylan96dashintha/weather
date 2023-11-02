package usecase

import (
	"context"
	"github.com/weather/internal/domain"
)

type WeatherReport interface {
	GetCoordinates(ctx context.Context,
		city string) (resp domain.AddressValidatorResponse, err error)
	ValidateAddress(ctx context.Context,
		addressValidationReq domain.AddressValidatorResponse) (isValid bool)
	GetWeatherReport(ctx context.Context,
		data domain.AddressValidatorResponse) (response domain.WeatherForecastResponse, err error)
	PrintFormattedReport(ctx context.Context, reportData domain.WeatherForecastResponse)
}
