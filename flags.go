package main

import (
	"flag"
	"log"
)

func init() {
	workPtr := flag.Int("work",  2, "Duration of the work session")
	shortPtr := flag.Int("short", 2, "Duration of the short breaks")
	longPtr := flag.Int("long", 4, "Duration of the long breaks")

	flag.Parse()

	if *workPtr > 60 || *shortPtr > 60 || *longPtr > 60 {
		log.Fatal("Invalid time input. Cannot be more than 60 minutes.")
	}

	if *workPtr < 0 || *shortPtr < 0 || *longPtr < 0 {
		log.Fatal("Invalid time input. Cannot be negative.")
	}

	breakTime = *shortPtr
	longBreak = *longPtr
	workTime = *workPtr
}