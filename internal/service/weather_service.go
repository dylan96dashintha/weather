package service

import (
	"context"
	"sync"
)

var (
	weatherObj  WeatherForecast
	weatherOnce sync.Once
)

type weather struct{}

func newWeatherObj() WeatherForecast {
	weatherObject := new(weather)
	return weatherObject
}

func GetWeatherServiceObject() WeatherForecast {
	weatherOnce.Do(func() {
		weatherObj = newWeatherObj()
	})

	return weatherObj
}

func (w weather) GetWeatherReport(ctx context.Context, city string) {

}
