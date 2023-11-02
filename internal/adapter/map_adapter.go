package adapter

import (
	"bytes"
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
	mapAdapterObj  MapAdapter
	mapAdapterOnce sync.Once
)

const (
	mapBaseUrl = "googleMapBaseUrl"
)

type mapAdapter struct {
	conf *config.Config
}

func newMapAdapterObj(conf *config.Config) MapAdapter {
	mapAdapterObject := new(mapAdapter)
	mapAdapterObject.conf = conf
	return mapAdapterObject
}

func GetWeatherServiceObject(conf *config.Config) MapAdapter {
	mapAdapterOnce.Do(func() {
		mapAdapterObj = newMapAdapterObj(conf)
	})

	return mapAdapterObj
}

func (m *mapAdapter) GetCoordinatesFromCity(ctx context.Context,
	city string) (resp domain.AddressValidatorResponse, err error) {

	var (
		response domain.AddressValidatorResponse
	)

	baseUrl, timeOutPriority := GetServiceDetailsByName(mapBaseUrl, m.conf)
	addressValidatorUrl := baseUrl + "/v1:validateAddress?key=" + m.conf.MapConfig.ApiKey

	reqBody := domain.AddressValidatorReq{}
	reqBody.Address.AddressLines = city

	jsonBody := new(bytes.Buffer)
	err = json.NewEncoder(jsonBody).Encode(reqBody)
	if err != nil {
		log.Fatal("Error in encoding request")
	}

	req, err := http.NewRequest(http.MethodPost, addressValidatorUrl, jsonBody)
	if err != nil {
		log.Fatal("Error in making request")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "close")

	httpClient := getHttpClient(timeOutPriority, m.conf)
	res, err := httpClient.Do(req)
	statusCode := http.StatusInternalServerError
	if res != nil {
		if res.Body != nil {
			statusCode = res.StatusCode
			defer res.Body.Close()
		}
	}

	if err != nil {
		fmt.Printf("Error in getting address validations")
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
