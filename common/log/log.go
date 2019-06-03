package log

import (
	"log"
	"os"
	"sync"
)

// Simple logger :))

// todo move the parameter to the config file
const logFileName = "./logFile.txt"

func Info(msg string) {
	writeToFile("[INFO]: " + msg)
}

func Error(msg string) {
	writeToFile("[ERROR]: " + msg)
}

func writeToFile(msg string) error {
	var mutex sync.Mutex

	mutex.Lock()
	defer mutex.Unlock()

	file, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	log.SetOutput(file)
	log.Println(msg)

	return nil
}
