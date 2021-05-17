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
	log.SetOutput(f)
	f.Close()

	log.SetFlags(log.Lmicroseconds)

	started := time.Now()
	log.Println("Start")

	time.Sleep(5 * time.Second)

	log.Println("Finished")
	timeElapsed := time.Since(started)
	log.Printf("Duration: %f seconds.\n", timeElapsed.Seconds())
}
