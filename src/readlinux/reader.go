package readlinux

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Read CPU temps and return float.
func ReadTemp() float64 {
	// Cpu temps stored here on a linux system
	filename := "/sys/class/thermal/thermal_zone0/temp"

	file, err := os.Open(filename)

	if err != nil {
		log.Panic(fmt.Sprintf("Failed to open %v.", filename))
	}

	scanner := bufio.NewScanner(file)

	// There's only one line in the file and it's the cpu temp
	scanner.Scan()

	// Multi lines would need to be scanned like so:
	// scanner.Split(bufio.ScanLines)
	// var text []string
	// for scanner.Scan() {
	// 	text = append(text, scanner.Text())
	// }
	// for _, line := range text {
	// 	fmt.Println(line)
	// }

	file.Close()

	temps, err := strconv.ParseFloat(scanner.Text(), 32)

	if err != nil {
		log.Panic("Error parsing cpu temperatures to float.")
	}

	// Linux temps reported in units of 1e-3 degC; convert to degC
	return temps / 1000
}
