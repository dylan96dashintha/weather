package main

import (
	"context"
	"github.com/weather/internal"
	"github.com/weather/internal/config"
)

func main() {

	conf := config.ConfigurationParser()

	ctx := context.Background()
	internal.InitWeatherForecast(ctx, conf)
}
