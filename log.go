package glog

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

type Logger interface {
	Info(output string, args ...any)
	Warn(output string, args ...any)
	Error(output string, args ...any)
	Debug(output string, args ...any)
}

type LogType string
type Color string

const (
	Green   Color = "\033[32m"
	Magenta Color = "\033[95m"
	Red     Color = "\033[31m"
	Yellow  Color = "\033[33m"
	Cyan    Color = "\033[36m"
	Reset   Color = "\033[0m"

	info  LogType = "INFO"
	warn  LogType = "WARN"
	error LogType = "ERROR"
	debug LogType = "DEBUG"

	space = " "
)

type GLogger struct {
	appName string
}

func NewGLogger(appName string) Logger {
	return &GLogger{appName: appName}
}
func (log *GLogger) Info(output string, args ...any) {
	log.log(info, fmt.Sprintf(output, args...))
}

func (log *GLogger) Warn(output string, args ...any) {
	log.log(warn, fmt.Sprintf(output, args...))
}

func (log *GLogger) Debug(output string, args ...any) {
	log.log(debug, fmt.Sprintf(output, args...))
}

func (log *GLogger) Error(output string, args ...any) {
	log.log(error, fmt.Sprintf(output, args...))
}

func (log *GLogger) log(logType LogType, output string) {
	switch logType {
	case info:
		log.print(string(logType)+space, Green, output)
	case warn:
		log.print(string(logType)+space, Yellow, output)
	case debug:
		log.print(string(logType), Green, output)
	case error:
		log.print(string(logType), Red, output)
		printStackTrace()
	}

}

func (log *GLogger) print(logType string, color Color, output string) {
	fmt.Printf("%s %s %s %s %d %s [%s] %s \n",
		time.Now().Format(time.RFC3339),
		color,
		logType,
		Magenta,
		os.Getpid(),
		Reset,
		log.appName,
		output,
	)
}

func printStackTrace() {
	trace := make([]byte, 4096)
	n := runtime.Stack(trace, false)
	fmt.Println(string(trace[:n]))
}
