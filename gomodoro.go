package main

import (
	"fmt"
	"time"
)

func main() {
	minutes := 2
	for minutes >= 0 {
		for i := 10; i >= 0; i-- {
			fmt.Printf("\033[2K\r%02d:%02d", minutes, i)
			time.Sleep(time.Second)
		}
		minutes--
	}


}

