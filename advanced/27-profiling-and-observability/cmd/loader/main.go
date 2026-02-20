package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // Blank import auto-registers /debug/pprof/* paths
	"strings"

	obs "go-playbook/advanced/27-profiling-and-observability"
)

func main() {
	// Start the pprof HTTP server in the background
	go func() {
		fmt.Println("Starting pprof on http://localhost:6060")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	fmt.Println("Running load generator... Press CTRL+C to stop.")

	// Create a synthetic payload
	payload := strings.Repeat("a", 1000) + "|" + strings.Repeat("b", 1000)

	// Blast the ProcessData function
	// We use 5 concurrent workers
	for i := 0; i < 5; i++ {
		go func(workerID int) {
			counter := 0
			for {
				id := fmt.Sprintf("req-%d-%d", workerID, counter)
				obs.ProcessData(id, payload)
				counter++
			}
		}(i)
	}

	// Block forever
	select {}
}
