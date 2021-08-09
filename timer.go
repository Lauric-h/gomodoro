package main

import (
	"fmt"
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
			fmt.Println("stop")
			break
		}

		o := time.Time{}.Add(difference)
		fmt.Println(o.Format("04:05"))
	}



	//timerStr := fmt.Sprintf("%02d:00", usedTimer)
	//area.Update(createBigLetters(timerStr, color))
	//time.Sleep(time.Second)
	//
	//remainingTime := usedTimer - 1
	//for remainingTime >= 0 {
	//	for i := 10; i >= 0; i-- {
	//		timerStr := fmt.Sprintf("%02d:%02d", remainingTime, i)
	//		area.Update(createBigLetters(timerStr, color))
	//		time.Sleep(time.Second)
	//	}
	//	remainingTime--
	//}

	//timerStr := fmt.Sprintf("%02d:00", usedTimer)
	//area.Update(createBigLetters(timerStr, color))
	//time.Sleep(time.Second)
	//
	//remainingTime := usedTimer - 1
	//for remainingTime >= 0 {
	//	for i := 10; i >= 0; i-- {
	//		timerStr := fmt.Sprintf("%02d:%02d", remainingTime, i)
	//		area.Update(createBigLetters(timerStr, color))
	//		time.Sleep(time.Second)
	//	}
	//	remainingTime--
	//}



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

	w.timer(area, usedTimer, fgColor)

	displayHeader(area, bgColor, w.endPhrase)
}

func waitDuration()  {
	start := time.Now()
	minutes := time.Minute * 2

	finish := start.Add(minutes)

	diff := finish.Sub(start)
	out := time.Time{}.Add(diff)
	fmt.Println(out.Format("04:05"))

	for range time.Tick(1 * time.Second) {
		currentTime := time.Now()
		difference := finish.Sub(currentTime)
		if difference <= 0 {
			fmt.Println("stop")
			break
		}

		o := time.Time{}.Add(difference)
		fmt.Println(o.Format("04:05"))
	}


}


