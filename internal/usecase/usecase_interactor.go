package usecase

import "context"

type WeatherReport interface {
	GetCoordinates(ctx context.Context, city string)
}
