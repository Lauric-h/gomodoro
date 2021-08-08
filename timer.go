package main

import (
	"fmt"
	"github.com/pterm/pterm"
	"time"
)

type WorkSession struct {
	shortBreak int
	longBreak  int
	workTime   int
	count      int
	start      string
	end        string
}

/**
timer launches timer with big letters
 */
func (w WorkSession) timer(area pterm.AreaPrinter, usedTimer int, color pterm.Color) {
	timerStr := fmt.Sprintf("%02d:00", usedTimer)
	area.Update(createBigLetters(timerStr, color))
	time.Sleep(time.Second)

	remainingTime := usedTimer - 1
	for remainingTime >= 0 {
		for i := 10; i >= 0; i-- {
			timerStr := fmt.Sprintf("%02d:%02d", remainingTime, i)
			area.Update(createBigLetters(timerStr, color))
			time.Sleep(time.Second)
		}
		remainingTime--
	}
}

/**
timerSession full session with displays depending on work or break
 */
func (w WorkSession) timerSession(timerType string, area pterm.AreaPrinter, breakTime int) {
	// Default values
	usedTimer := 0
	fgColor := pterm.FgBlack
	bgColor := pterm.BgBlack

	if timerType == "break" {
		usedTimer = breakTime
		fgColor = pterm.FgGreen
		bgColor = pterm.BgGreen
		w.start = fmt.Sprintf("Starting %d minutes break...\n", breakTime)
		w.end = "Break is over, get back to work"
	}
	if timerType == "work" {
		usedTimer = w.workTime
		fgColor = pterm.FgRed
		bgColor = pterm.BgRed
		w.start = fmt.Sprintf("Starting %d minutes work session...\n", w.workTime)
		w.end = "Work session is over, time to take a break"
	}

	displayHeader(area, bgColor, w.start)

	w.timer(area, usedTimer, fgColor)

	displayHeader(area, bgColor, w.end)
}

