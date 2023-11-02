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
	"strings"
	"syscall"
)

func InitWeatherForecast(ctx context.Context, conf *config.Config) {

	city := readCityName()
	weatherServiceObj := service.GetWeatherServiceObject(conf)

	err := weatherServiceObj.GetWeatherReport(ctx, city)
	if err != nil {
		log.Fatal(err)
	}
	startCronJob(ctx, city, conf, weatherServiceObj)
}

func readCityName() (city string) {
	flag.Parse()
	nonFlagArgs := flag.Args()
	if len(nonFlagArgs) > 0 {
		city = strings.Join(nonFlagArgs, " ")
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

func startCronJob(ctx context.Context, city string, conf *config.Config,
	weatherSvcObj service.WeatherForecast) {
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
		err := weatherSvcObj.GetWeatherReport(ctx, city)
		if err != nil {
			log.Fatal(err)
		}
	})
	if err != nil {
		log.Fatal("error in updating weather")
	}

	c.Start()

	<-sigs
}
