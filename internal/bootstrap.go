package internal

import (
	"context"
	"flag"
	"fmt"
	"github.com/weather/internal/service"
)

func InitWeatherForecast(ctx context.Context) {

	city := readCityName()
	weatherServiceObj := service.GetWeatherServiceObject()
	weatherServiceObj.GetWeatherReport(ctx, city)

}

func readCityName() (city string) {
	flag.Parse()
	nonFlagArgs := flag.Args()
	if len(nonFlagArgs) > 0 {
		city = nonFlagArgs[0]
		err := stringValidator(city)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("weather forecast for the city %s", city)
	} else {
		panic(fmt.Sprintf("cannot find the city name"))
	}

	return city
}