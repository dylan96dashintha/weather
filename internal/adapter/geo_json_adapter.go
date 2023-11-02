package adapter

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/weather/internal/config"
	"github.com/weather/internal/domain"
	"io"
	"log"
	"net/http"
	"sync"
)

var (
	geoJsonAdapterObj  GeoJsonAdapter
	geoJsonAdapterOnce sync.Once
)

const (
	geoJsonBaseUrl = "geoJsonBaseUrl"
)

type geoJsonAdapter struct {
	conf *config.Config
}

func newGeoJsonAdapterObj(conf *config.Config) GeoJsonAdapter {
	geoJsonAdapterObject := new(geoJsonAdapter)
	geoJsonAdapterObject.conf = conf
	return geoJsonAdapterObject
}

func GetGeoJsonAdapterObject(conf *config.Config) GeoJsonAdapter {
	geoJsonAdapterOnce.Do(func() {
		geoJsonAdapterObj = newGeoJsonAdapterObj(conf)
	})

	return geoJsonAdapterObj
}

func (g geoJsonAdapter) GetWeatherForecast(ctx context.Context,
	lat, lon float64) (resp domain.WeatherForecastResponse, err error) {
	var (
		response domain.WeatherForecastResponse
	)

	baseUrl, timeOutPriority := GetServiceDetailsByName(geoJsonBaseUrl, g.conf)
	weatherForecastUrl := baseUrl + "/2.0/complete?lat=" + fmt.Sprintf("%f", lat) + "&lon=" + fmt.Sprintf("%f", lon)

	httpClient := getHttpClient(timeOutPriority, g.conf)

	req, err := http.NewRequest(http.MethodGet, weatherForecastUrl, nil)
	if err != nil {
		log.Fatal("Error in making request")
	}
	req.Header.Set("User-Agent", "locationforecast")

	res, err := httpClient.Do(req)
	statusCode := http.StatusInternalServerError
	if res != nil {
		if res.Body != nil {
			statusCode = res.StatusCode
			defer res.Body.Close()
		}
	}

	if err != nil {
		fmt.Printf("Error in getting weather forecast")
		return resp, err
	}

	if !isSuccessful(statusCode) {
		return resp, errors.New("bad request")
	}

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error in reading response")
		return resp, err
	}

	if respBody != nil {
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			fmt.Printf("Error in unmarshalling response")
			return resp, err
		}
	}

	return response, nil
}
