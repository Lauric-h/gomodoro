package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"log"
	"os"
	"time"
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
		sessionCount++
		totalSessionCount++
		printSession := fmt.Sprintf("Session number %v", totalSessionCount)
		//fmt.Println("Session number ", totalSessionCount)
		workTimer(1)
		if sessionCount == 4 {
		sessionCount = 0
		breakTime = longBreak
		}
		breakTimer(1)
	//}

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
	s.EnableMouse()
	s.EnablePaste()
	s.Clear()

	// Draw initial boxes
	drawBox(s, 1, 1, 42, 7, boxStyle, printSession)
	time.Sleep(time.Second * 3)
	//drawBox(s, 1, 1, 42, 7, boxStyle, "coucou")
	time.Sleep(time.Second * 3)
	drawText(s, 2, 2, 32, 14, boxStyle, "Cuciui")

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

	}
}
