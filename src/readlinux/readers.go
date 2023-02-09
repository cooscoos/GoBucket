package readlinux

import (
	"bufio"
	"errors"
	"fmt"

	"os"
	"strconv"
)

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

	// Read in memory usage from here
	filename := "/proc/meminfo"
	reading, err := readFile(filename)

	if err != nil {
		return Memory{}, errors.New(fmt.Sprintf("Failed to open %v because: %v.", filename, err))
	}

	// Store info in Memory struct
	m, err := Memory{}.New(reading)

	if err != nil {
		return Memory{}, err
	}

	return m, nil

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
