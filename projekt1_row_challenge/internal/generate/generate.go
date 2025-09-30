package generate

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

type City struct {
	Name  string
	Value int
}

func CityTemps(inFile, outFile string) error {
	cities, err := readCities(inFile)
	if err != nil {
		return err
	}

	rndCities := genRandomCityTemps(cities, 1_000_000)

	err = writeCityTemps(outFile, rndCities)
	if err != nil {
		return err
	}

	return nil
}

func readCities(fileName string) ([]City, error) {
	//
	// Open file.

	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file '%s': %w", fileName, err)
	}
	defer file.Close()

	//
	// Read line by line.

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	cities := []City{}

	for scanner.Scan() {
		line := scanner.Text()
		info := strings.Split(line, "\t")

		if len(info) < 2 {
			return nil, fmt.Errorf("failed to parse '%s'", line)
		}

		name := info[0]

		// Parse population inti int.
		population, err := strconv.Atoi(info[1])
		if err != nil {
			return nil, fmt.Errorf("failed to parse '%s': %w", line, err)
		}

		cities = append(cities, City{Name: name, Value: population})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan file: %w", err)
	}

	return cities, nil
}

func genRandomCityTemps(cities []City, n int) []City {
	rndCities := make([]City, n)

	for i := range n {
		rndCities[i].Name = cities[rand.IntN(len(cities))].Name
		rndCities[i].Value = rand.IntN(350) // 0 - 35 degrees
	}

	return rndCities
}

func writeCityTemps(fileName string, cities []City) error {
	f, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file '%s': %w", fileName, err)
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	for _, city := range cities {
		_, err = writer.WriteString(fmt.Sprintf(
			"%s:%d.%d\n",
			city.Name, city.Value/10, city.Value%10,
		))
		if err != nil {
			return fmt.Errorf("failed to write into file: %w", err)
		}
	}

	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("failed to write into file: %w", err)
	}

	return nil
}
