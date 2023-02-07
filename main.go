package main

import (
	"fmt"
	"gobucket/src/readlinux"
)

func main() {
	cpu_temp := readlinux.ReadTemp()
	fmt.Println(cpu_temp)
}
