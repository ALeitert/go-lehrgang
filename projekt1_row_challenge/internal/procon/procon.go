package procon

import (
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

type CityStats struct {
	sum, count int
	min, max   int
}

func NewCityStats() CityStats {
	return CityStats{
		sum:   0,
		count: 0,
		min:   math.MaxInt,
		max:   math.MinInt,
	}
}

func (cs *CityStats) add(temp int) {
	cs.sum += temp
	cs.count++
	cs.min = min(cs.min, temp)
	cs.max = max(cs.max, temp)
}

func (cs *CityStats) merge(other CityStats) {
	cs.sum += other.sum
	cs.count += other.count
	cs.min = min(cs.min, other.min)
	cs.max = max(cs.max, other.max)
}

func CityTemps(inFile string) error {
	const WorkerCount = 3

	wg := sync.WaitGroup{}

	dataChan := make(chan []byte, 32)
	statChan := make(chan map[string]CityStats, WorkerCount) // Should be an array.

	workerErrs := make([]error, WorkerCount+1)

	//
	// Reader "thread".

	wg.Add(1)
	go func() {
		defer wg.Done()
		workerErrs[WorkerCount] = readFile(inFile, dataChan)
	}()

	//
	// Worker "threads".

	wg.Add(WorkerCount)
	for id := range WorkerCount {
		go func(id int) {
			defer wg.Done()
			workerErrs[id] = processLines(dataChan, statChan)
		}(id)
	}

	wg.Wait()
	close(statChan)

	err := errors.Join(workerErrs...)
	if err != nil {
		return err
	}

	_ = mergeStats(statChan)

	return nil
}

func readFile(fileName string, dataChan chan<- []byte) error {
	//
	// Open file.

	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("failed to open file '%s': %w", fileName, err)
	}
	defer file.Close()
	defer close(dataChan)

	//
	// Read data chunks.

	const BufferSize = 1 << 12

	var (
		bufferArr [BufferSize]byte
		leftovers = bufferArr[:0]

		lenB int
	)

	for err == nil {
		nextChunk := make([]byte, 2*BufferSize)
		var lenC int

		// Invariant: There is enough space in `nextChunk` for all `leftovers`.
		for breakIdx := 0; err == nil && breakIdx < BufferSize; {
			lenC += copy(nextChunk[lenC:], leftovers)

			lenB, err = file.Read(bufferArr[:])
			newData := bufferArr[:lenB]

			maxSize := min(2*BufferSize-lenC, lenB)

			idx := lastIndexByte(newData[:maxSize], '\n')
			if idx < 0 {
				continue
			}

			breakIdx = lenC + idx

			lenC += copy(nextChunk[lenC:], newData[:idx+1])
			leftovers = newData[idx+1:]
		}

		dataChan <- nextChunk[:lenC]
	}

	if errors.Is(err, io.EOF) {
		err = nil
	}

	return err
}

func lastIndexByte(s []byte, c byte) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			return i
		}
	}
	return -1
}

func processLines(dataChan <-chan []byte, into chan<- map[string]CityStats) error {
	stats := map[string]CityStats{}

	totalLines := 0

	for data := range dataChan {
		for line := range strings.Lines(string(data)) {
			totalLines++
			line = line[:len(line)-1]

			name, temp, err := parseLine(line)
			if err != nil {
				return err
			}

			cStats, ok := stats[name]
			if !ok {
				cStats = NewCityStats()
			}

			cStats.add(temp)
			stats[name] = cStats
		}
	}

	into <- stats

	return nil
}

func parseLine(line string) (string, int, error) {
	info := strings.Split(line, ":")

	if len(info) < 2 {
		return "", 0, fmt.Errorf("failed to parse '%s'", line)
	}

	name := info[0]
	temp := info[1]

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

func mergeStats(from <-chan map[string]CityStats) map[string]CityStats {
	merged := map[string]CityStats{}

	for data := range from {
		for name, stats := range data {
			mStats, ok := merged[name]
			if ok {
				stats.merge(mStats)
			}
			merged[name] = stats
		}
	}

	return merged
}
