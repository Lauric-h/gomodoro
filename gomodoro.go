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
						os.Exit(1)
					case 112:
						fmt.Println(string(r))
						fmt.Println("pause")
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
