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
	str := fmt.Sprintf("Starting %d minutes break...\n", breakTime)
	displayHeader(area, pterm.FgBlack, pterm.BgGreen, str)

	timerStr := fmt.Sprintf("%02d:00", breakTime)
	area.Update(createBigLetters(timerStr, pterm.FgGreen))
	time.Sleep(time.Second)

	remainingTime := breakTime - 1
	for remainingTime >= 0 {
		for i := 2; i >= 0; i-- {
			timerStr := fmt.Sprintf("%02d:%02d", remainingTime, i)
			area.Update(createBigLetters(timerStr, pterm.FgGreen))
			time.Sleep(time.Second)
		}
		remainingTime--
	}
	displayHeader(area, pterm.FgBlack, pterm.BgGreen, "Break is over, get back to work")
}

func (w WorkSession) WorkTimer(area pterm.AreaPrinter) {
	str := fmt.Sprintf("Starting %d minutes work session...\n", w.workTime)
	displayHeader(area, pterm.FgBlack, pterm.BgRed, str)

	timerStr := fmt.Sprintf("%02d:00", w.workTime)
	area.Update(createBigLetters(timerStr, pterm.FgRed))
	time.Sleep(time.Second)

	remainingTime := w.workTime - 1
	for remainingTime >= 0 {
		for i := 10; i >= 0; i-- {
			timerStr := fmt.Sprintf("%02d:%02d", remainingTime, i)
			area.Update(createBigLetters(timerStr, pterm.FgRed))
			time.Sleep(time.Second)
		}
		remainingTime--
	}

	displayHeader(area, pterm.FgBlack, pterm.BgRed, "Work session is over, time for a break!")
}

func (w WorkSession) Pause() {
	// todo
}

func (w WorkSession) Stop() {
	//todo
}
