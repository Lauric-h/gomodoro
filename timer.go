package main

import (
	"fmt"
	"github.com/0xAX/notificator"
	"github.com/pterm/pterm"
	"time"
)

type WorkSession struct {
	shortBreak  time.Duration
	longBreak   time.Duration
	workTime    time.Duration
	count       int
	startPhrase string
	endPhrase   string
}

/**
timer launches timer with big letters
 */
func (w WorkSession) timer(area pterm.AreaPrinter, usedTimer time.Duration, color pterm.Color) {
	d := time.Minute * usedTimer
	start := time.Now()
	finish := start.Add(d)

	for range time.Tick(1 * time.Second) {
		currentTime := time.Now()
		difference := finish.Sub(currentTime)
		if difference <= 0 {
			break
		}

		o := time.Time{}.Add(difference)
		timerStr := fmt.Sprintf(o.Format("04:05"))
		area.Update(createBigLetters(timerStr, color))
	}
}

/**
timerSession full session with displays depending on work or break
 */
func (w WorkSession) timerSession(timerType string, area pterm.AreaPrinter, breakTime time.Duration) {
	// Default values
	usedTimer := 0 * time.Second
	fgColor := pterm.FgBlack
	bgColor := pterm.BgBlack

	if timerType == "break" {
		usedTimer = breakTime
		fgColor = pterm.FgGreen
		bgColor = pterm.BgGreen
		w.startPhrase = fmt.Sprintf("Starting %d minutes break...\n", breakTime)
		w.endPhrase = "Break is over, get back to work"
	}
	if timerType == "work" {
		usedTimer = w.workTime
		fgColor = pterm.FgRed
		bgColor = pterm.BgRed
		w.startPhrase = fmt.Sprintf("Starting %d minutes work session...\n", w.workTime)
		w.endPhrase = "Work session is over, time to take a break"
	}

	displayHeader(area, bgColor, w.startPhrase)
	notify.Push(timerType, w.startPhrase, "", notificator.UR_CRITICAL)

	w.timer(area, usedTimer, fgColor)

	displayHeader(area, bgColor, w.endPhrase)
	notify.Push(timerType, w.endPhrase, "", notificator.UR_CRITICAL)
}



