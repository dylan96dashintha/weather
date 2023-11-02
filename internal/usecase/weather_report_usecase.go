package usecase

import (
	"context"
	"fmt"
	"github.com/weather/internal/adapter"
	"github.com/weather/internal/config"
	"github.com/weather/internal/domain"
	"sync"
)

var (
	weatherReportObj  WeatherReport
	weatherReportOnce sync.Once
)

type weatherReport struct {
	mapAdapter adapter.MapAdapter
}

func newWeatherReportObj(conf *config.Config) WeatherReport {
	weatherReportObject := new(weatherReport)
	weatherReportObject.mapAdapter = adapter.GetWeatherServiceObject(conf)
	return weatherReportObject
}

func GetWeatherServiceObject(conf *config.Config) WeatherReport {
	weatherReportOnce.Do(func() {
		weatherReportObj = newWeatherReportObj(conf)
	})

	return weatherReportObj
}

func (w weatherReport) GetCoordinates(ctx context.Context,
	city string) (resp domain.AddressValidatorResponse, err error) {

	response, err := w.mapAdapter.GetCoordinatesFromCity(ctx, city)
	if err != nil {
		fmt.Println("Error in getting coordinates")
		return resp, err
	}

	return response, nil
}

func (w weatherReport) ValidateAddress(ctx context.Context,
	addressValidationReq domain.AddressValidatorResponse) (isValid bool) {

	for _, addressComp := range addressValidationReq.Result.Address.AddressComponents {
		if addressComp.ComponentType == "country" &&
			addressComp.ConfirmationLevel == "CONFIRMED" {
			return true
		}
	}

	return false
}
