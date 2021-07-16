package main

import (
  "fmt"
  "time"
)

func main() {
	workTimer(2)
	breakTimer(2)
}

func breakTimer(breakTime int) {
	fmt.Println("Starting break...")
	for breakTime >= 0 {
		for i := 10; i >= 0; i-- {
			fmt.Printf("\033[2K\r%02d:%02d", breakTime, i)
			time.Sleep(time.Second)
		}
		breakTime--
	}
	fmt.Println()
	fmt.Println("Break is over, get back to work")
}

func workTimer(workTime int) {
	fmt.Println("Starting work session...")
	for workTime >= 0 {
		for i := 10; i >= 0; i-- {
			fmt.Printf("\033[2K\r%02d:%02d", workTime, i)
			time.Sleep(time.Second)
		}
		workTime--
	}
	fmt.Println()
	fmt.Println("Work session is over, time for a break!")
}