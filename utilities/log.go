package utilities

import (
	"log"
	"os"
)

var Log *log.Logger

func Init() {
	Log = log.New(os.Stdout, "GoShips: ", log.Ldate|log.Ltime|log.Lshortfile)
}
