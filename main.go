package main

import (
	"log"
	"os"
	"time"
)

func main() {

	f, err := os.Create("run_output_" + time.Now().Format("02-01-2006-15:04:05") + ".txt")
	if err != nil {
		log.Fatalf("error creating file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	started := time.Now()
	log.Println("----------------------\nStarted at", started.Format("02-01-2006 - 15:04:05"))

	time.Sleep(5 * time.Second)

	log.Println("Finished at", time.Now().Format("02-01-2006 - 15:04:05"))
	timeElapsed := time.Since(started)
	log.Printf("Duration: %.2f seconds.\n----------------------\n", timeElapsed.Seconds())
}
