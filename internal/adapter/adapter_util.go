package adapter

import (
	"github.com/weather/internal/config"
	"net/http"
)

func GetServiceDetailsByName(name string, config *config.Config) (url string, timeoutPriority string) {

	if len(name) == 0 {
		return "", ""
	}

	for _, s := range config.ServiceConfig.Services {
		if s.Name == name {
			return s.URL, s.TimeoutPriority
		}
	}

	return "", ""
}

func getHttpClient(timeoutPriority string, config *config.Config) http.Client {

	client := config.AppConfig.HttpClient
	if timeoutPriority == "LOW" {
		client = config.AppConfig.HttpClientLow
	} else if timeoutPriority == "HIGH" {
		client = config.AppConfig.HttpClientHigh
	}
	return client
}

var (
	successStatuses = []int{http.StatusOK, http.StatusCreated, http.StatusNoContent}
)

func isSuccessful(code int) bool {
	for _, v := range successStatuses {
		if v == code {
			return true
		}
	}
	return false
}
