package timewriter

import (
	"log"
	"testing"
)

func TestWrite(t *testing.T) {
	timeWriter := &TimeWriter{
		Dir:           "./log",
		Compress:      true,
		ReserveDay:    30,
		LogFilePrefix: "timewriter_test", // default is process name
	}
	logDebug := log.New(timeWriter, " [Debug] ", log.LstdFlags)
	logDebug.Println("this is debug")
}
