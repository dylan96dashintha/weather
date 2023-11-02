package adapter

import (
	"context"
	"github.com/weather/internal/domain"
)

type MapAdapter interface {
	GetCoordinatesFromCity(ctx context.Context,
		city string) (resp domain.AddressValidatorResponse, err error)
}
