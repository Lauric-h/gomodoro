package main

import (
	"flag"
	"fmt"
	"github.com/mattn/go-tty"
	"log"
	"os"
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

	tity, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tity.Close()

	//for {
	//	r, err := tity.ReadRune()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	// handle key event
	//	fmt.Println(string(r))
	//
	//}

	// ------------------

		for {
			go func() {
				r, err := tity.ReadRune()
				fmt.Println(r)
				if err != nil {
					fmt.Println(err)
				}
				// exit
				switch r {
					case 99:
						fmt.Println(string(r))
						os.Exit(1)
					case 112:
						fmt.Println(string(r))
						fmt.Println("pause")
				}
			}()


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