package main

import (
	"flag"
	"fmt"
	"github.com/gdamore/tcell/v2"
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

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	// Set default text style
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s.SetStyle(defStyle)

	// Clear screen
	s.Clear()

	quit := func() {
		s.Fini()
		os.Exit(0)
	}

	for {
		s.Show()
		// Poll event
		ev := s.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				quit()
			}
		}
	}


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