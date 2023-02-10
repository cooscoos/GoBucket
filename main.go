package main

import (
	"fmt"
	"gobucket/src/readlinux"
	"log"
	"net/http"
)

func main() {

	fmt.Printf("Starting server at http://localhost:8000\n")
	http.HandleFunc("/", indexPage)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

}

// Defines the html of the index page for GoBucket
func indexPage(w http.ResponseWriter, r *http.Request) {

	// Fetch cpu temperatures and memory usage
	cpu_temp, err := readlinux.ReadTemp()
	if err != nil {
		log.Panic(err)
	}

	memory, err := readlinux.ReadMemory()
	if err != nil {
		log.Panic(err)
	}

	fmt.Fprintf(w, `<h1>GoBucket</h1>
	<p> Avg. CPU temperature (degC): %v<p>
	<p> Memory total / used / available (GB): %v / %v / %v`, cpu_temp, memory.Total, memory.Used, memory.Available)

}
