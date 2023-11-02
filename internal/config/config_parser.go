package config

import (
	"fmt"
	yaml "gopkg.in/yaml.v3"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

const (
	filePathService = "config/service_config.yaml"
	filePathMap     = "config/map_config.yaml"
	filePathTimeout = "config/timeout_config.yaml"
)

var (
	ServiceConf OutBoundServiceConfig
	MapConf     MapConfig
	TimeoutConf TimeoutConfig
)

func parseOutBoundServiceConfig() {

	byt, err := os.ReadFile(filePathService)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error in loading service_config.yaml, err %s", err))
	}

	err = yaml.Unmarshal(byt, &ServiceConf)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error in unmarshalling service_config.yaml, err %s", err))
	}

}

func parseMapConfig() {

	byt, err := os.ReadFile(filePathMap)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error in loading service_config.yaml, err %s", err))
	}

	err = yaml.Unmarshal(byt, &MapConf)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error in unmarshalling service_config.yaml, err %s", err))
	}

}

func parseTimeoutConfig() {

	byt, err := os.ReadFile(filePathTimeout)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error in loading service_config.yaml, err %s", err))
	}

	err = yaml.Unmarshal(byt, &TimeoutConf)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error in unmarshalling service_config.yaml, err %s", err))
	}

}

func ConfigurationParser() (conf *Config) {
	parseOutBoundServiceConfig()
	parseMapConfig()
	parseTimeoutConfig()

	appConf := AppConfig{
		HttpClient:     getHttpClient(TimeoutConf),
		HttpClientLow:  getHttpClientLow(TimeoutConf),
		HttpClientHigh: getHttpClientHigh(TimeoutConf),
	}
	return &Config{
		ServiceConfig: ServiceConf,
		MapConfig:     MapConf,
		AppConfig:     appConf,
		TimeOutConfig: TimeoutConf,
	}
}

func getHttpClient(to TimeoutConfig) http.Client {
	return http.Client{
		Timeout: time.Millisecond * time.Duration(to.RemoteCall),
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   time.Millisecond * time.Duration(to.Dial),
				KeepAlive: time.Millisecond * time.Duration(to.KeepAlive),
			}).DialContext,
		},
	}
}

func getHttpClientLow(to TimeoutConfig) http.Client {
	return http.Client{
		Timeout: time.Millisecond * time.Duration(to.RemoteCallLow),
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   time.Millisecond * time.Duration(to.DialLow),
				KeepAlive: time.Millisecond * time.Duration(to.KeepAliveLow),
			}).DialContext,
		},
	}
}

func getHttpClientHigh(to TimeoutConfig) http.Client {
	return http.Client{
		Timeout: time.Millisecond * time.Duration(to.RemoteCallHigh),
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   time.Millisecond * time.Duration(to.DialHigh),
				KeepAlive: time.Millisecond * time.Duration(to.KeepAliveHigh),
			}).DialContext,
		},
	}
}
