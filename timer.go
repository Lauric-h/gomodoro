package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"time"
)


func breakTimer(breakTime int, s tcell.Screen, style tcell.Style) {
	fmt.Printf("Starting %d minutes break...\n", breakTime)
	fmt.Printf("\033[2K\r%02d:00", breakTime)
	time.Sleep(time.Second)
	remainingTime := breakTime - 1
	for remainingTime >= 0 {
		for i := 2; i >= 0; i-- {
			clock := fmt.Sprintf("%02d:%02d", remainingTime, i)
			drawText(s, 1, 1, 6, 2, style, clock)
			fmt.Println("t ", clock)
			time.Sleep(time.Second)
		}
		remainingTime--
	}
	fmt.Println()
	fmt.Println("Break is over, get back to work")
}

//func workTimer(workTime int) {
//	fmt.Println("Starting work session...")
//	fmt.Printf("\033[2K\r%02d:00", workTime)
//	time.Sleep(time.Second)
//	remainingTime := workTime - 1
//	for remainingTime >= 0 {
//		for i := 10; i >= 0; i-- {
//			fmt.Printf("\033[2K\r%02d:%02d", remainingTime, i)
//			time.Sleep(time.Second)
//		}
//		remainingTime--
//	}
//	fmt.Println()
//	fmt.Println("Work session is over, time for a break!")
//}
