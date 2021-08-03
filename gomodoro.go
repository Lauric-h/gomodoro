package main

import (
	"fmt"
	"github.com/mattn/go-tty"
	"log"
	"os"
	"time"
)

var (
	sessionCount = 0
	w WorkSession
)

func main() {
	fmt.Println("Welcome to Gomodoro!")
	fmt.Println("Press p to pause the timer")
	fmt.Println("Press s to stop the timer")
	fmt.Println("Press c to exit the program")

	// ------------------------
	// Get current date
	date := time.Now()
	formattedDate := fmt.Sprintf(date.Format("02-01-2006 15:04:05"))

	// Log date to file
	if checkForLogFile() {
		// Append to existing file
		appendLineToFile(formattedDate)
	} else {
		// Create and write to new file
		writeLineToFile(formattedDate)
	}

	// --------------------

	// Open keypress listener
	tity, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tity.Close()

	// ------------------
	// Main loop
	for {
		go func() {
			r, err := tity.ReadRune()
			fmt.Println(r)
			if err != nil {
				fmt.Println(err)
			}
			// exit
			switch r {
			// c - saving to log & exit program
			case 99:
				appendLineToFile(formatLogInfo(w.count, w.workTime, w.shortBreak))
				fmt.Println("Exiting the program, see you later!")
				os.Exit(1)
			// p - pausing timer
			case 112:
				fmt.Println("Timer paused")
				fmt.Println("Stopping the clock, what do you want to do?")
				// Todo - add menu to start over
				// Todo - add menu to continue
			}
		}()

		// todo - put this into func
		breakTime := w.shortBreak
		sessionCount++
		w.count++
		fmt.Println("Session number ", w.count)
		w.WorkTimer()
		if sessionCount == 4 {
			sessionCount = 0
			breakTime = w.longBreak
		}
		w.BreakTimer(breakTime)
	}
}
