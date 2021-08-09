package main

import (
	"github.com/pterm/pterm"
	"math/rand"
	"time"
)

func introLogo() {
	gomoLogo, _ := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("G", pterm.NewStyle(getRandomColor())),
		pterm.NewLettersFromStringWithStyle("o", pterm.NewStyle(getRandomColor())),
		pterm.NewLettersFromStringWithStyle("m", pterm.NewStyle(getRandomColor())),
		pterm.NewLettersFromStringWithStyle("o", pterm.NewStyle(getRandomColor())),
		pterm.NewLettersFromStringWithStyle("d", pterm.NewStyle(getRandomColor())),
		pterm.NewLettersFromStringWithStyle("o", pterm.NewStyle(getRandomColor())),
		pterm.NewLettersFromStringWithStyle("r", pterm.NewStyle(getRandomColor())),
		pterm.NewLettersFromStringWithStyle("o", pterm.NewStyle(getRandomColor()))).
		Srender()
	pterm.DefaultCenter.Print(gomoLogo)
	pterm.Info.Println("Welcome to Gomodoro!")
}

func printInfo() {
	pterm.Info.Println("Press c to exit the program")
	pterm.Error.Prefix = pterm.Prefix{
		Text: "WORK",
		Style: pterm.NewStyle(pterm.BgRed, pterm.FgBlack),
	}
	pterm.Error.Println("Red is for work")
	pterm.Success.Prefix = pterm.Prefix{
		Text: "BREAK",
		Style: pterm.NewStyle(pterm.BgGreen, pterm.FgBlack),
	}
	pterm.Success.Println("Green is for breaks")
	pterm.Print("\n\n")
}

func createBigLetters(str string, color pterm.Color) string {
	bigTimer, _ := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle(str, pterm.NewStyle(color))).
		Srender()
	return bigTimer
}

func createHeader(bgColor pterm.Color) pterm.HeaderPrinter {
	return pterm.HeaderPrinter{
		TextStyle: pterm.NewStyle(pterm.FgBlack),
		BackgroundStyle: pterm.NewStyle(bgColor),
		Margin: 15,
	}
}

func displayHeader(area pterm.AreaPrinter, bgColor pterm.Color, str string) {
	newHeader := createHeader(bgColor)
	h := newHeader.Sprintf(str)
	area.Update(h)
	time.Sleep(time.Second * 3)
}

func getRandomColor() pterm.Color {
	colors := []pterm.Color{
		pterm.FgBlack,
		pterm.FgRed,
		pterm.FgGreen,
		pterm.FgYellow,
		pterm.FgBlue,
		pterm.FgMagenta,
		pterm.FgCyan,
		pterm.FgWhite,
	}
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(colors)-1)
	return colors[r]
}



