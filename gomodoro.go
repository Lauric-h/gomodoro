package main

import (
	"fmt"
	"github.com/mattn/go-tty"
	"log"
	"os"
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

	// Open keypress listener
	tity, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tity.Close()

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
						fmt.Println("Exiting the program, see you later!")
						os.Exit(1)
					case 112:
						fmt.Println(string(r))
						fmt.Println("Timer paused")
						fmt.Println("Stopping the clock, what do you want to do?")
						// Todo - add menu to start over
						// Todo - add menu to continue
				}
			}()

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
