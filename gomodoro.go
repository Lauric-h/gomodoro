package main

import (
	"fmt"
	"github.com/mattn/go-tty"
	"github.com/pterm/pterm"
	"log"
	"os"
	"time"
)

var (
	sessionCount = 0
	w WorkSession
)

func main() {
	gomoLogo, _ := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("Gomodoro", pterm.NewStyle(pterm.FgLightCyan))).
		Srender()

	pterm.DefaultCenter.Print(gomoLogo)
	pterm.Info.Println("Welcome to Gomodoro!")
	pterm.Info.Println("Press p to pause the timer")
	pterm.Info.Println("Press s to stop the timer")
	pterm.Info.Println("Press c to exit the program")
	pterm.Error.Prefix = pterm.Prefix{
		Text: "WORK",
		Style: pterm.NewStyle(pterm.BgRed, pterm.FgBlack),
	}
	pterm.Error.Println("Red is for work")
	pterm.Success.Prefix = pterm.Prefix{
		Text: "BREAK",
		Style: pterm.NewStyle(pterm.BgGreen, pterm.FgBlack),
	}
	pterm.Success.Println("Green is for breaks")
	pterm.Print("\n\n")

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

	area, _ := pterm.DefaultArea.WithCenter().WithRemoveWhenDone(true).Start()

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
				pterm.Info.Println("Exiting the program, see you later!")
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
		w.WorkTimer(*area)
		if sessionCount == 4 {
			sessionCount = 0
			breakTime = w.longBreak
		}
		w.BreakTimer(breakTime, *area)
		area.Stop()
	}
}
