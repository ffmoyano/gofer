package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

var timeFormat = time.Now().Format("02-01-2006")

type logger struct {
	file   *os.File
	level  int
	log    *log.Logger
	name   string
	prefix string
}

// when we call a log of level x all logs with level >= x are also called
var loggers = []logger{
	{file: new(os.File), level: 2, log: new(log.Logger), name: "infolog"},
	{file: new(os.File), level: 1, log: new(log.Logger), name: "warnlog"},
	{file: new(os.File), level: 0, log: new(log.Logger), name: "errorlog"},
	{file: nil, level: 999, log: new(log.Logger), name: "consolelog"},
}

var err error

// OpenLogs Open this function receives a path where the logs folder will be created
func OpenLogs(path string) {
	if _, err = os.Stat("logs"); os.IsNotExist(err) {
		err = os.Mkdir("logs", 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	for i, logger := range loggers {
		logger = logger.open()
		loggers[i] = logger
	}
}

// CloseLogs Close this function closes the logs
func CloseLogs() {
	for _, logger := range loggers {
		logger.close()
	}
}

// Error logs a message with level ERROR
func Error(message string) {
	write(newMessage(message), 0)
}

// Info logs a message with level INFO
func Info(message string, args ...any) {
	write(fmt.Sprintf(message, args), 2)
}

// Warn logs a message with level WARN
func Warn(message string) {
	write(newMessage(message), 1)
}

func (logger logger) close() {
	err = logger.file.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func (logger logger) open() logger {
	if logger.file != nil {

		logger.file, err = os.OpenFile(fmt.Sprintf("logs/%s_%s.log", logger.name, timeFormat),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

		logger.log = log.New(logger.file, logger.prefix, log.Ldate|log.Ltime)

	} else { // console log has no file associated
		logger.log = log.New(os.Stdout, "", log.LstdFlags)
	}
	return logger
}

func write(message string, level int) {
	var prefix string
	for _, logger := range loggers {
		if logger.level >= level {
			switch level {
			case 0:
				prefix = "[ERROR]"
			case 1:
				prefix = "[WARN]"
			case 2:
				prefix = "[INFO]"
			}
			logger.log.Print(prefix + " " + message)
		}
	}
}

func newMessage(message string) string {
	_, file, line, _ := runtime.Caller(2)
	return fmt.Sprintf("[%s][%d] : %s", file, line, message)
}
