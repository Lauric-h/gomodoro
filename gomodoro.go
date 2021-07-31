package main

import (
	"github.com/gdamore/tcell/v2"
	"log"
	"os"
)

var (
	sessionCount int = 0
	totalSessionCount int = 0
	breakTime int
	workTime int
	longBreak int
)



func main() {
	//for
	//{

	//}

	// Initialize styles
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple)

	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	s.Clear()




	// Event loop
	quit := func() {
		s.Fini()
		os.Exit(0)
	}
	for {
		// Update screen
		s.Show()

		// Poll event
		ev := s.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				quit()
			} else if ev.Key() == tcell.KeyCtrlL {
				s.Sync()
			} else if ev.Rune() == 'C' || ev.Rune() == 'c' {
				s.Clear()
			}

		}

		// timer
		sessionCount++
		totalSessionCount++
		//printSession := fmt.Sprintf("Session number %v", totalSessionCount)
		//drawText(s, 2, 2, 32, 14, boxStyle, printSession)
		breakTimer(1, s, boxStyle)


		//fmt.Println("Session number ", totalSessionCount)
		//workTimer(1)
		//if sessionCount == 4 {
		//	sessionCount = 0
		//	breakTime = longBreak
		//}
		//breakTimer(1)


	}
}
