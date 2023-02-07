package readlinux

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Stores info about system memory usage.
type Memory struct {
	total     float64
	free      float64
	available float64
	used      float64
}

// Read average CPU temps, return degrees C.
func ReadTemp() (float64, error) {
	// Cpu temps stored here on a linux system; read them in.
	filename := "/sys/class/thermal/thermal_zone0/temp"
	reading, err := readFile(filename)

	if err != nil {
		return .0, errors.New(fmt.Sprintf("Failed to open %v because: %v.", filename, err))
	}

	// Average cpu temp log is always one line: a temperature value in units 1e-3 degC.
	temps, err := strconv.ParseFloat(reading[0], 32)
	if err != nil {
		return .0, errors.New(fmt.Sprintf("Error parsing cpu temperatures to float because: %v", err))
	}

	// Convert 1e-3 degC to degC.
	return temps / 1000, nil
}

// Read and return memory usage.
func ReadMemory() (Memory, error) {
	filename := "/proc/meminfo"
	reading, err := readFile(filename)

	if err != nil {
		return Memory{}, errors.New(fmt.Sprintf("Failed to open %v because: %v.", filename, err))
	}

	// Search through the lines for patterns of interest
	// Can link to regexp with struct tags????
	for _, line := range reading {
		fmt.Println(line)
	}

	return Memory{}, nil

}

// Read in file and return array of lines.
func readFile(filename string) ([]string, error) {

	file, err := os.Open(filename)

	if err != nil {
		return []string{}, err
	}

	// Scan through all lines
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	return text, nil
}
