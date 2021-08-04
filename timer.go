package main

import (
	"fmt"
	"github.com/pterm/pterm"
	"time"
)

type WorkSession struct {
	shortBreak	int
	longBreak 	int
	workTime 		int
	count 			int
}

func (w WorkSession) BreakTimer(breakTime int, area pterm.AreaPrinter) {
	//fmt.Printf("Starting %d minutes break...\n", breakTime)
	//fmt.Printf("\033[2K\r%02d:00", breakTime)

	timerStr := fmt.Sprintf("%02d:00", breakTime)
	bigTimer, _ := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle(timerStr, pterm.NewStyle(pterm.FgGreen))).
		Srender()

	area.Update(bigTimer)

	time.Sleep(time.Second)
	remainingTime := breakTime - 1
	for remainingTime >= 0 {
		for i := 2; i >= 0; i-- {
			timerStr := fmt.Sprintf("%02d:%02d", remainingTime, i)
			bigTimer, _ := pterm.DefaultBigText.WithLetters(
				pterm.NewLettersFromStringWithStyle(timerStr, pterm.NewStyle(pterm.FgGreen))).
				Srender()

			area.Update(bigTimer)


			//fmt.Printf("\033[2K\r%02d:%02d", remainingTime, i)
			time.Sleep(time.Second)
		}
		remainingTime--
	}
	//fmt.Println()
	//fmt.Println("Break is over, get back to work")
}

func (w WorkSession) WorkTimer(area pterm.AreaPrinter) {
	//fmt.Println("Starting work session...")
	//fmt.Printf("\033[2K\r%02d:00", w.workTime)

	timerStr := fmt.Sprintf("%02d:00", w.workTime)
	bigTimer, _ := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle(timerStr, pterm.NewStyle(pterm.FgRed))).
		Srender()

	area.Update(bigTimer)


	time.Sleep(time.Second)
	remainingTime := w.workTime - 1



	for remainingTime >= 0 {
		for i := 10; i >= 0; i-- {
			timerStr := fmt.Sprintf("%02d:%02d", remainingTime, i)
			bigTimer, _ := pterm.DefaultBigText.WithLetters(
				pterm.NewLettersFromStringWithStyle(timerStr, pterm.NewStyle(pterm.FgRed))).
				Srender()

			area.Update(bigTimer)
			//fmt.Printf("\033[2K\r%02d:%02d", remainingTime, i)
			time.Sleep(time.Second)
		}
		remainingTime--
	}
	//fmt.Println()
	//fmt.Println("Work session is over, time for a break!")
}

func (w WorkSession) Pause() {
	// todo
}

func (w WorkSession) Stop() {
	//todo
}
