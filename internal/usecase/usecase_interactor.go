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
}
