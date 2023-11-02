package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/weather/internal/config"
	"github.com/weather/internal/usecase"
	"sync"
)

var (
	weatherObj  WeatherForecast
	weatherOnce sync.Once
)

type weather struct {
	conf                 *config.Config
	weatherReportUseCase usecase.WeatherReport
}

func newWeatherObj(conf *config.Config) WeatherForecast {
	weatherObject := new(weather)
	weatherObject.conf = conf
	weatherObject.weatherReportUseCase = usecase.GetWeatherServiceObject(conf)
	return weatherObject
}

func GetWeatherServiceObject(conf *config.Config) WeatherForecast {
	weatherOnce.Do(func() {
		weatherObj = newWeatherObj(conf)
	})

	return weatherObj
}

func (w weather) GetWeatherReport(ctx context.Context, city string) (err error) {

	// get the coordinates
	response, err := w.weatherReportUseCase.GetCoordinates(ctx, city)
	if err != nil {
		return err
	}

	// check whether a valid address
	isValid := w.weatherReportUseCase.ValidateAddress(ctx, response)
	if !isValid {
		return errors.New(fmt.Sprintf("City name validation failed, check spelling next time"))
	}

	return nil
}
