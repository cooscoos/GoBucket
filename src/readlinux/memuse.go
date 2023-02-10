package readlinux

import (
	"fmt"
	"math"
	"strconv"

	"github.com/helmbold/richgo/regexp"
)

// Stores info about system memory usage.
type Memory struct {
	Total     float64
	Available float64
	Used      float64
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
			fmt.Println(match)
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

// Convert array of memory values into a Memory struct (in GB)
func (m Memory) FromArray(mem_array []float64) Memory {

	// The order of memory values extracted from linux system is:
	// [MemTotal, MemFree, MemAvailable, Buffers, Cached]
	// Convert all to GB

	total := mem_array[0]
	free := mem_array[1]
	avail := mem_array[2]
	buffers := mem_array[3]
	cached := mem_array[4]

	// Used memory defined as follows:
	used := total - free - buffers - cached

	m.Total = kbToGb(total)
	m.Available = kbToGb(avail)
	m.Used = kbToGb(used)

	return m

}

// Round a kb value to a gb value returning to 1 decimal point
func kbToGb(m float64) float64 {
	return math.Round(m*10/math.Pow(1024, 2)) / 10

}
