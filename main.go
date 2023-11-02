package main

import (
	"context"
	"github.com/weather/internal"
)

func main() {

	ctx := context.Background()
	internal.InitWeatherForecast(ctx)
}
