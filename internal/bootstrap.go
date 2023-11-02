package internal

import (
	"context"
	"flag"
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/weather/internal/config"
	"github.com/weather/internal/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func InitWeatherForecast(ctx context.Context, conf *config.Config) {

	city := readCityName()
	weatherServiceObj := service.GetWeatherServiceObject(conf)

	weatherServiceObj.GetWeatherReport(ctx, city)

	sigs := make(chan os.Signal, 1)

	signal.Notify(
		sigs,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGTERM,
	)

	// cron job for updating data background
	c := cron.New()
	_, err := c.AddFunc(conf.TimeOutConfig.CronJobTimeInterval, func() {
		weatherServiceObj.GetWeatherReport(ctx, city)
	})
	if err != nil {
		log.Fatal("error in updating weather")
	}

	c.Start()

	<-sigs
}

func readCityName() (city string) {
	flag.Parse()
	nonFlagArgs := flag.Args()
	if len(nonFlagArgs) > 0 {
		city = nonFlagArgs[0]
		err := stringValidator(city)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Printf("weather forecast for the city %s \n\n", city)
	} else {
		log.Fatal(fmt.Sprintf("cannot find the city name"))
	}

	return city
}
