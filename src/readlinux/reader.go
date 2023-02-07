package readlinux

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Read CPU temps and return float.
func ReadTemp() (float64, error) {
	// Cpu temps stored here on a linux system; read them in.
	filename := "/sys/class/thermal/thermal_zone0/temp"
	reading, err := readFile(filename)

	if err != nil {
		return .0, errors.New(fmt.Sprintf("Failed to open %v because: %v.", filename, err))
	}

	temps, err := strconv.ParseFloat(reading, 32)
	if err != nil {
		return .0, errors.New(fmt.Sprintf("Error parsing cpu temperatures to float because: %v", err))
	}

	// Linux temps reported in units of 1e-3 degC; convert to degC.
	return temps / 1000, nil
}

// Read in file and return the first line only.
func readFile(filename string) (string, error) {

	file, err := os.Open(filename)

	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(file)

	// There's always only one line in these linux log files
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

	return scanner.Text(), nil
}
