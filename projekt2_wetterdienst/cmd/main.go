package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/risingwavelabs/eris"

	"weather-service/internal/config"
	"weather-service/internal/server"
	"weather-service/internal/services"
	"weather-service/internal/station"
)

func main() {
	fmt.Println("Wetterdienst")

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	err := run(ctx)
	if err != nil {
		fmt.Println(eris.ToString(err, true))
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	//
	// Load and print config.

	var configPath string
	flag.StringVar(&configPath, "config", "", "Path to a file with configurations.")
	flag.Parse()

	err := config.C.Load(configPath)
	if err != nil {
		return eris.Wrap(err, "error while loading config")
	}
	config.C.Print()

	//
	// Run services.

	svcList := []services.Service{
		// List services here.
		&server.Server{},
		&server.Streamer{},
	}
	for _, city := range config.C.Cities {
		svcList = append(svcList, station.City(city))
	}

	err = services.Run(ctx, svcList)
	if err != nil {
		return eris.Wrap(err, "error while running services")
	}

	return nil
}
