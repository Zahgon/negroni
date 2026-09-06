package negroni

import (
	"net/http"
	"text/template"
	"time"
)

type LoggerEntry struct {
	StartTime string
	Status    int
	Duration  time.Duration
	Hostname  string
	Method    string
	Path      string
	Request   *http.Request
}

var LoggerDefaultFormat = "{{.StartTime}} | {{.Status}} | \t {{.Duration}} | {{.Hostname}} | {{.Method}} {{.Path}}"

var LoggerDefaultDateFormat = time.RFC3339

type ALogger interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}

type Logger struct {
	ALogger
	dateFormat string
	template   *template.Template
}

func NewLogger() *Logger { _ = "STUB: not implemented"; return nil }

func (l *Logger) SetFormat(format string) { _ = "STUB: not implemented"; return }

func (l *Logger) SetDateFormat(format string) { _ = "STUB: not implemented"; return }

func (l *Logger) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	_ = "STUB: not implemented"
	return
}
