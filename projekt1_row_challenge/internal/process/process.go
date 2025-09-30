package process

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func CityTemps(inFile string) error {
	temps, err := readTemps(inFile)
	if err != nil {
		return err
	}

	_ = computeStats(temps)
	return nil
}

func readTemps(fileName string) (map[string][]int, error) {
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

	temps := map[string][]int{}

	for scanner.Scan() {
		name, temp, err := parseLine(scanner.Text())
		if err != nil {
			return nil, err
		}

		// Question: Why is the below correct?
		temps[name] = append(temps[name], temp)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan file: %w", err)
	}

	return temps, nil
}

func parseLine(line string) (string, int, error) {
	info := strings.Split(line, ":")

	if len(info) < 2 {
		return "", 0, fmt.Errorf("failed to parse '%s'", line)
	}

	name := info[0]
	temp := info[1]

	//
	// Parse temperature.

	wholeStr := temp[:len(temp)-2]
	digitStr := temp[len(temp)-1:]

	wholeInt, err := strconv.Atoi(wholeStr)
	if err != nil {
		return "", 0, fmt.Errorf("failed to parse '%s': %w", line, err)
	}

	digitInt, err := strconv.Atoi(digitStr)
	if err != nil {
		return "", 0, fmt.Errorf("failed to parse '%s': %w", line, err)
	}

	return name, wholeInt*10 + digitInt, nil
}

type CityStats struct {
	name       string
	sum, count int
	min, max   int
}

func (cs *CityStats) add(temp int) {
	cs.sum += temp
	cs.count++
	cs.min = min(cs.min, temp)
	cs.max = max(cs.max, temp)
}

func computeStats(temps map[string][]int) []CityStats {
	allStats := make([]CityStats, 0, len(temps))

	for city, tempList := range temps {
		stats := CityStats{
			name: city,
			min:  math.MaxInt,
			max:  math.MinInt,
		}

		for _, temp := range tempList {
			stats.add(temp)
		}

		allStats = append(allStats, stats)
	}

	slices.SortFunc(allStats, func(a, b CityStats) int {
		return cmp.Compare(a.name, b.name)
	})

	return allStats
}
