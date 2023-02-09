package readlinux

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/helmbold/richgo/regexp"
)

// Stores info about system memory usage.
type Memory struct {
	total     float64
	free      float64
	available float64
	used      float64
}

// Create a new memory struct based on input string array
func (m Memory) New(reading []string) (Memory, error) {

	// Search through each line of string array with regexp capturing to <number>.
	pattern := `(?m)^(?:Buffers|Cached|Mem(?:Total|Free|Available)):.*?(?P<number>\d+)`
	re := regexp.MustCompile(pattern)

	memory_vals := []float64{}

	// Match on re and convert to float
	for _, line := range reading {
		match := re.Match(line)

		if match != nil {
			s := match.NamedGroups["number"]
			val, err := strconv.ParseFloat(s, 64)

			if err != nil {
				return Memory{}, err
			}

			memory_vals = append(memory_vals, val)
		}

	}

	fmt.Println(memory_vals)

	// Convert the array to Memory struct
	m = m.FromArray(memory_vals)

	return m, nil
}

// Converts an array of memory values into a Memory struct (in GB)
func (m Memory) FromArray(mem_array []float64) Memory {

	// The order of memory values extracted from linux system is:
	// [MemTotal, MemFree, MemAvailable, Buffers, Cached]
	// Convert all to GB

	total := mem_array[0]
	free := mem_array[1]
	avail := mem_array[2]
	buffers := kbToGb(mem_array[3])
	cached := kbToGb(mem_array[4])

	m.total = kbToGb(total)
	m.free = kbToGb(free)
	m.available = kbToGb(avail)

	// Used memory defined as follows:
	m.used = kbToGb(total - free - buffers - cached)

	return m

}

// Round a kb value to a gb value returning to 1 decimal point
func kbToGb(m float64) float64 {
	return math.Round(m*10/math.Pow(1024, 2)) / 10

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
