package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func checkForLogFile() bool {
	_, err := os.Stat("sessions.txt")
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func writeLineToFile(str string) {
	err := ioutil.WriteFile("sessions.txt", []byte(str + "\n"), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func appendLineToFile(str string) {
	file, err := os.OpenFile("sessions.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	if _, err := file.WriteString(str + "\n"); err != nil {
		log.Fatal(err)
	}
}

func formatLogInfo(sessionNumber int, workTime int, breakTime int) string {
	return "Number of sessions: " +
		strconv.Itoa(sessionNumber) +
		"\n Work time: " + strconv.Itoa(workTime) +" minutes"+
		"\n Break time: " + strconv.Itoa(breakTime) +" minutes"+
		"\n------------------------------"
}