package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

func init() {
	workPtr := flag.Int("work",  25, "Duration of the work session")
	shortPtr := flag.Int("short", 5, "Duration of the short breaks")
	longPtr := flag.Int("long", 10, "Duration of the long breaks")

	flag.Parse()

	// Timer cannot be set for more than 60 minutes
	if *workPtr > 60 || *shortPtr > 60 || *longPtr > 60 {
		log.Fatal("Invalid time input. Cannot be more than 60 minutes.")
	}

	// Timer cannot be set for less than 1 minute
	if *workPtr < 1 || *shortPtr < 1 || *longPtr < 1 {
		log.Fatal("Invalid time input. Cannot be less than 1 minute.")
	}

	// Refuse other command line arguments
	if len(flag.Args()) >= 1 {
		for _, value := range flag.Args() {
			fmt.Printf("Invalid argument: %v\n", value)
		}
		log.Fatal("Exiting... Too many arguments")
	}

	w.shortBreak = time.Duration(*shortPtr)
	w.longBreak = time.Duration(*longPtr)
	w.workTime = time.Duration(*workPtr)
}