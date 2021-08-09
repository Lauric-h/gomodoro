package main

import (
	"fmt"
	"github.com/0xAX/notificator"
	"github.com/mattn/go-tty"
	"github.com/pterm/pterm"
	"log"
	"os"
	"time"
)

var (
	sessionCount = 0
	w WorkSession
	notify *notificator.Notificator
)

func main() {
	introLogo()
	printInfo()

	notify = notificator.New(notificator.Options{
		DefaultIcon: "icon/default.png",
		AppName:     "Gomodoro",
	})

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

	// Open keypress listener
	tity, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tity.Close()

	// Main loop
	area, _ := pterm.DefaultArea.WithCenter().WithRemoveWhenDone(true).Start()

	for {
		go func() {
			r, err := tity.ReadRune()
			if err != nil {
				fmt.Println(err)
			}
			// exit
			switch r {
			// c - saving to log & exit program
			case 99:
				appendLineToFile(formatLogInfo(w.count, int(w.workTime), int(w.shortBreak)))
				pterm.Info.Println("Exiting the program, see you later!")
				os.Exit(1)
			}
		}()

		breakTime := w.shortBreak
		sessionCount++
		w.count++

		w.timerSession("work", *area, breakTime)

		if sessionCount == 4 {
			sessionCount = 0
			breakTime = w.longBreak
		}

		w.timerSession("break", *area, breakTime)

		area.Stop()
	}
}
