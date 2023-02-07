package main

import (
	"fmt"
	"gobucket/src/readlinux"
	"log"
)

func main() {
	cpu_temp, err := readlinux.ReadTemp()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(cpu_temp)
}
