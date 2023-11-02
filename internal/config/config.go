package config

import "net/http"

type OutBoundServiceConfig struct {
	Services []ServiceConfig `yaml:"services" json:"services"`
}

type MapConfig struct {
	ApiKey string `yaml:"mapAPIKey"`
}

type ServiceConfig struct {
	Name            string `yaml:"name"`
	TimeoutPriority string `yaml:"timeout"`
	URL             string `yaml:"url"`
}

type Config struct {
	ServiceConfig OutBoundServiceConfig
	MapConfig     MapConfig
	AppConfig     AppConfig
}

type TimeoutConfig struct {
	RemoteCall int64 `yaml:"remote_call_timeout"`
	Dial       int64 `yaml:"dial_timeout"`
	KeepAlive  int64 `yaml:"keep_alive_timeout"`

	RemoteCallLow int64 `yaml:"remote_call_timeout_low"`
	DialLow       int64 `yaml:"dial_timeout_low"`
	KeepAliveLow  int64 `yaml:"keep_alive_timeout_low"`

	RemoteCallHigh int64 `yaml:"remote_call_timeout_high"`
	DialHigh       int64 `yaml:"dial_timeout_high"`
	KeepAliveHigh  int64 `yaml:"keep_alive_timeout_high"`
}

type AppConfig struct {
	HttpClient     http.Client
	HttpClientLow  http.Client
	HttpClientHigh http.Client
}
