package main

import (
	"fmt"
	"time"
)

type WorkSession struct {
	shortBreak	int
	longBreak 	int
	workTime 		int
	count 			int
}

func (w WorkSession) BreakTimer(breakTime int) {
	fmt.Printf("Starting %d minutes break...\n", breakTime)
	fmt.Printf("\033[2K\r%02d:00", breakTime)
	time.Sleep(time.Second)
	remainingTime := breakTime - 1
	for remainingTime >= 0 {
		for i := 2; i >= 0; i-- {
			fmt.Printf("\033[2K\r%02d:%02d", remainingTime, i)
			time.Sleep(time.Second)
		}
		remainingTime--
	}
	fmt.Println()
	fmt.Println("Break is over, get back to work")
}

func (w WorkSession) WorkTimer() {
	fmt.Println("Starting work session...")
	fmt.Printf("\033[2K\r%02d:00", w.workTime)
	time.Sleep(time.Second)
	remainingTime := w.workTime - 1
	for remainingTime >= 0 {
		for i := 10; i >= 0; i-- {
			fmt.Printf("\033[2K\r%02d:%02d", remainingTime, i)
			time.Sleep(time.Second)
		}
		remainingTime--
	}
	fmt.Println()
	fmt.Println("Work session is over, time for a break!")
}
