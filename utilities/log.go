package utilities

import (
	"log"
	"os"
)

var Log *log.Logger

func Init(logFile string) {
	file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic("Cannot open " + logFile + " for log")
	}
	Log = log.New(file, "GoShips: ", log.Ldate|log.Ltime|log.Lshortfile)
}
