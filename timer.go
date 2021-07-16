package main

import (
	"fmt"
	"time"
)

type countdown struct {
	h int
	m int
	s int
}

func clock(workTimer int, breakTimer int) {
	for i:= workTimer; i>=0; i-- {
		fmt.Printf("\033[2K\r%d", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println()
	fmt.Println("time for a break")

	for i:= breakTimer; i>=0; i-- {
		fmt.Printf("\033[2K\r%d", i)
		time.Sleep(time.Second)
	}

	fmt.Println()
	fmt.Println("break is over")

	fmt.Println()
}
