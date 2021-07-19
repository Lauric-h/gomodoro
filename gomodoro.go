package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	sessionCount int = 0
)

func main() {
	workPtr := flag.Int("work",  2, "Duration of the work session")
	shortPtr := flag.Int("short", 2, "Duration of the short breaks")
	longPtr := flag.Int("long", 4, "Duration of the long breaks")

	flag.Parse()

	breakTime := *shortPtr

	for
	{
		sessionCount++
		fmt.Println("Session number ", sessionCount)
		workTimer(*workPtr)
		if sessionCount % 4 == 0 {
			breakTime = *longPtr
		}
		breakTimer(breakTime)
	}
}

func breakTimer(breakTime int) {
	fmt.Printf("Starting %d minutes break...\n", breakTime)
	for breakTime >= 0 {
		for i := 2; i >= 0; i-- {
			fmt.Printf("\033[2K\r%02d:%02d", breakTime, i)
			time.Sleep(time.Second)
		}
		breakTime--
	}
	fmt.Println()
	fmt.Println("Break is over, get back to work")
}

func workTimer(workTime int) {
	fmt.Println("Starting work session...")
	for workTime >= 0 {
		for i := 10; i >= 0; i-- {
			fmt.Printf("\033[2K\r%02d:%02d", workTime, i)
			time.Sleep(time.Second)
		}
		workTime--
	}
	fmt.Println()
	fmt.Println("Work session is over, time for a break!")
}