package main

import (
	"flag"
	"fmt"
	"github.com/MarinX/keylogger"
	"log"
	"time"
)

var (
	sessionCount int = 0
	totalSessionCount int = 0
)

func main() {
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

	// ------------------------

	// find keyboard device, does not require a root permission
	keyboard := keylogger.FindKeyboardDevice()

	// check if we found a path to keyboard
	if len(keyboard) <= 0 {
		fmt.Println("No keyboard found...you will need to provide manual input path")
		return
	}

	fmt.Println("Found a keyboard at", keyboard)
	// init keylogger with keyboard
	k, err := keylogger.New(keyboard)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer k.Close()

	events := k.Read()

	// range of events
	for e := range events {
		switch e.Type {
		// EvKey is used to describe state changes of keyboards, buttons, or other key-like devices.
		// check the input_event.go for more events
		case keylogger.EvKey:

			// if the state of key is pressed
			if e.KeyPress() {
				fmt.Println("[event] press key ", e.KeyString())
			}

			// if the state of key is released
			if e.KeyRelease() {
				fmt.Println("[event] release key ", e.KeyString())
			}

			break
		}
		breakTime := *shortPtr
		sessionCount++
		totalSessionCount++
		fmt.Println("Session number ", totalSessionCount)
		workTimer(*workPtr)
		if sessionCount == 4 {
			sessionCount = 0
			breakTime = *longPtr
		}
		breakTimer(breakTime)
	}

	// ------------------------
	for
	{
		breakTime := *shortPtr
		sessionCount++
		totalSessionCount++
		fmt.Println("Session number ", totalSessionCount)
		workTimer(*workPtr)
		if sessionCount == 4 {
			sessionCount = 0
			breakTime = *longPtr
		}
		breakTimer(breakTime)
	}
}

func breakTimer(breakTime int) {
	fmt.Printf("Starting %d minutes break...\n", breakTime)
	fmt.Printf("\033[2K\r%02d:00", breakTime)
	time.Sleep(time.Second)
	remainingTime := breakTime - 1
	for remainingTime >= 0 {
		for i := 2; i >= 0; i-- {
			fmt.Printf("\033[2K\r%02d:%02d", remainingTime, i)
			time.Sleep(time.Second)
		}
		remainingTime--
	}
	fmt.Println()
	fmt.Println("Break is over, get back to work")
}

func workTimer(workTime int) {
	fmt.Println("Starting work session...")
	fmt.Printf("\033[2K\r%02d:00", workTime)
	time.Sleep(time.Second)
	remainingTime := workTime - 1
	for remainingTime >= 0 {
		for i := 10; i >= 0; i-- {
			fmt.Printf("\033[2K\r%02d:%02d", remainingTime, i)
			time.Sleep(time.Second)
		}
		remainingTime--
	}
	fmt.Println()
	fmt.Println("Work session is over, time for a break!")
}