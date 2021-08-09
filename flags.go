package main

import (
	"flag"
	"fmt"
	"log"
)

func init() {
	workPtr := flag.Duration("work",  2, "Duration of the work session")
	shortPtr := flag.Duration("short", 2, "Duration of the short breaks")
	longPtr := flag.Duration("long", 4, "Duration of the long breaks")



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

	w.shortBreak = *shortPtr
	w.longBreak = *longPtr
	w.workTime = *workPtr
}