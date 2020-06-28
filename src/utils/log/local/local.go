package local

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/sofyan48/boilerplate/src/utils/log/entity"
)

// LogLocal ...
type LogLocal struct{}

// LogLocalHandler ...
func LogLocalHandler() *LogLocal {
	return &LogLocal{}
}

// LogLocalInterface ...
type LogLocalInterface interface {
	Exception(err error)
}

// CreateLog ...
func logStorage() *log.Logger {
	logFile, err := os.OpenFile("./log/event.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	// defer logFile.Close()
	logger := log.New(logFile, "", log.LstdFlags)
	return logger
}

func logsLocal(name string, desc interface{}) {
	logs := logStorage()
	logging := &entity.Logging{}
	logging.Name = name
	logging.Description = desc
	logging.TimeAt = time.Now()
	data, err := json.Marshal(logging)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	logs.Println(string(data))
}

// Exception ...
func (logs *LogLocal) Exception(err error) {
	go logsLocal("Exception", err)
}
