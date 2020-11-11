package common

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type requestLogger struct {
	logger *log.Logger
}

var logger = log.New(os.Stdout, "", log.LstdFlags|log.Llongfile)

func GetRequestLogger() *requestLogger {
	return &requestLogger{logger}
}

func GetLogger() *log.Logger {
	return logger
}

func (l *requestLogger) requestText(req *http.Request) string {
	if req != nil {
		return fmt.Sprintf("Request([%v] %v)", req.Method, req.URL.Path)
	} else {
		return "Request(nil)"
	}
}

func (l *requestLogger) output(log string) {
	err := l.logger.Output(3, log)
	if err != nil {
		logger.Printf("write log fail - %v", err)
	}
}

func (l *requestLogger) PrintJson(req *http.Request, label string, val interface{}) {
	out, err := json.Marshal(val)
	if err != nil {
		logger.Printf("to json fail. %v", err)
	}
	l.output(fmt.Sprint("%v - %s : %s", l.requestText(req), label, string(out)))
}

func (l *requestLogger) Print(req *http.Request, args ...interface{}) {
	l.output(fmt.Sprintf("%v - %s", l.requestText(req), fmt.Sprint(args...)))
}

func (l *requestLogger) Println(req *http.Request, args ...interface{}) {
	l.output(fmt.Sprintf("%v - %s", l.requestText(req), fmt.Sprintln(args...)))
}

func (l *requestLogger) Printf(req *http.Request, format string, args ...interface{}) {
	l.output(fmt.Sprintf("%v - %s", l.requestText(req), fmt.Sprintf(format, args...)))
}

func (l *requestLogger) Fatal(req *http.Request, args ...interface{}) {
	l.output(fmt.Sprint("%v - %s", l.requestText(req), fmt.Sprint(args...)))
	os.Exit(1)
}

func (l *requestLogger) Fatalln(req *http.Request, args ...interface{}) {
	l.output(fmt.Sprintf("%v - %s", l.requestText(req), fmt.Sprintln(args...)))
	os.Exit(1)
}

func (l *requestLogger) Fatalf(req *http.Request, format string, args ...interface{}) {
	l.output(fmt.Sprintf("%v - %s", l.requestText(req), fmt.Sprintf(format, args...)))
	os.Exit(1)
}

func (l *requestLogger) Panic(req *http.Request, args ...interface{}) {
	msg := fmt.Sprintf("%v - %s", l.requestText(req), fmt.Sprint(args...))
	l.output(msg)
	panic(msg)
}

func (l *requestLogger) Panicln(req *http.Request, args ...interface{}) {
	msg := fmt.Sprintf("%v - %s", l.requestText(req), fmt.Sprintln(args...))
	l.output(msg)
	panic(msg)
}

func (l *requestLogger) Panicf(req *http.Request, format string, args ...interface{}) {
	msg := fmt.Sprintf("%v - %s", l.requestText(req), fmt.Sprintf(format, args...))
	l.output(msg)
	panic(msg)
}
