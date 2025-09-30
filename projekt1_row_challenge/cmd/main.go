package main

import (
	"fmt"
	"os"
	"time"

	"row-challenge/internal/generate"
	"row-challenge/internal/process"
	"row-challenge/internal/procon"
)

// go run cmd/main.go generate ./files/cities.txt ./files/city_temps.txt
// go run cmd/main.go process ./files/city_temps.txt
// go run cmd/main.go procon ./files/city_temps.txt

func main() {
	fmt.Println("Row Challenge")

	if len(os.Args) <= 1 {
		fmt.Println("Please specify a task!")
		os.Exit(1)
	}

	start := time.Now()

	switch os.Args[1] {
	case "generate":
		if len(os.Args[2:]) < 2 {
			fmt.Println("Please specify input and output path.")
			os.Exit(1)
		}

		err := generate.CityTemps(os.Args[2], os.Args[3])
		if err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(1)
		}

	case "process":
		err := process.CityTemps(os.Args[2])
		if err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(1)
		}

	case "procon":
		err := procon.CityTemps(os.Args[2])
		if err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(1)
		}

	default:
		fmt.Println("Unknown task.")
		os.Exit(1)
	}

	fmt.Println("done after", time.Since(start))
}
