package usecase

import (
	"context"
	"sync"
)

var (
	weatherReportObj  WeatherReport
	weatherReportOnce sync.Once
)

type weatherReport struct{}

func newWeatherReportObj() WeatherReport {
	weatherReportObject := new(weatherReport)
	return weatherReportObject
}

func GetWeatherServiceObject() WeatherReport {
	weatherReportOnce.Do(func() {
		weatherReportObj = newWeatherReportObj()
	})

	return weatherReportObj
}

func (w weatherReport) GetCoordinates(ctx context.Context, city string) {

}
