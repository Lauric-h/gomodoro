package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func checkForLogFile() bool {
	_, err := os.Stat("temp.txt")
	if os.IsNotExist(err) {
		fmt.Println("File does not exist. Creating one...")
		return false
	}
	fmt.Println("File exists. Adding log to it...")
	return true
}

func writeLineToFile(str string) {
	err := ioutil.WriteFile("temp.txt", []byte(str + "\n"), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func appendLineToFile(str string) {
	file, err := os.OpenFile("temp.txt", os.O_APPEND|os.O_WRONLY, 0644)
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