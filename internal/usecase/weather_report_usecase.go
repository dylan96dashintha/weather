package usecase

import (
	"context"
	"fmt"
	"github.com/weather/internal/adapter"
	"github.com/weather/internal/config"
	"github.com/weather/internal/domain"
	"sync"
	"time"
)

var (
	weatherReportObj  WeatherReport
	weatherReportOnce sync.Once
)

type weatherReport struct {
	mapAdapter     adapter.MapAdapter
	geoJsonAdapter adapter.GeoJsonAdapter
	conf           *config.Config
}

func newWeatherReportObj(conf *config.Config) WeatherReport {
	weatherReportObject := new(weatherReport)
	weatherReportObject.conf = conf
	weatherReportObject.mapAdapter = adapter.GetMapAdapterObject(conf)
	weatherReportObject.geoJsonAdapter = adapter.GetGeoJsonAdapterObject(conf)
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

func (w weatherReport) GetWeatherReport(ctx context.Context,
	data domain.AddressValidatorResponse) (response domain.WeatherForecastResponse, err error) {

	locationData := data.Result.Geocode.Location
	resp, err := w.geoJsonAdapter.GetWeatherForecast(ctx, locationData.Latitude, locationData.Longitude)
	if err != nil {
		fmt.Println("Error in getting weather report")
		return resp, err
	}

	return resp, nil
}

func (w weatherReport) PrintFormattedReport(ctx context.Context, reportData domain.WeatherForecastResponse) {

	unitMap := getUnitRelatedToTheWeatherReport(ctx, reportData)
	currentTime := time.Now().UTC()
	fmt.Println(fmt.Sprintf("current time is, %v\n", currentTime))
	currentTimeYear, currentTimeMonth, currentTimeDay := currentTime.Date()
	currentTimeHour := currentTime.Hour()

	for _, timeSeriesData := range reportData.Properties.Timeseries {
		year, month, day := timeSeriesData.Time.Date()
		hour := timeSeriesData.Time.Hour()
		instantData := timeSeriesData.Data.Instant.Details
		if year == currentTimeYear && month ==
			currentTimeMonth && day ==
			currentTimeDay && hour == currentTimeHour {
			fmt.Println(fmt.Sprintf("Temperature in this hour : %v %s",
				instantData.AirTemperature, unitMap[domain.Temperature]))
			fmt.Println(fmt.Sprintf("wind speed in this hour : %v %s",
				instantData.WindSpeed, unitMap[domain.WindSpeed]))
			fmt.Println(fmt.Sprintf("wind from direction in this hour : %v %s",
				instantData.WindFromDirection, unitMap[domain.WindDirection]))
			fmt.Println(fmt.Sprintf("relative humidity in this hour : %v %s",
				instantData.RelativeHumidity, unitMap[domain.RelativeHumidity]))
			fmt.Println(fmt.Sprintf("precipitation amount in next 1 hour : %v %s",
				timeSeriesData.Data.Next1Hours.Details.PrecipitationAmount, unitMap[domain.Precipitation]))
			fmt.Println(fmt.Sprintf("Next weather update will be in %s\n", w.conf.TimeOutConfig.UpdateInterval))
		}
	}
}

func getUnitRelatedToTheWeatherReport(ctx context.Context,
	data domain.WeatherForecastResponse) map[string]string {

	unitMap := make(map[string]string)
	unitData := data.Properties.Meta.Units
	unitMap[domain.Temperature] = unitData.AirTemperature
	unitMap[domain.WindSpeed] = unitData.WindSpeed
	unitMap[domain.WindDirection] = unitData.WindFromDirection
	unitMap[domain.Precipitation] = unitData.PrecipitationAmount
	unitMap[domain.RelativeHumidity] = unitData.RelativeHumidity
	return unitMap
}
